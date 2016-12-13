package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func toCamelCase(name string, initialUpper bool) string {
	var parts []string
	for _, part := range strings.Split(name, "_") {
		parts = append(parts, strings.ToUpper(part[:1])+strings.ToLower(part[1:]))
	}
	if initialUpper {
		return strings.Join(parts, "")
	}
	return strings.ToLower(parts[0]) + strings.Join(parts[1:], "")
}

var enumTemplate = `package xmmsclient

// auto-generated

const (
{{- range $key, $values := .}}
{{ range $values}}
	{{.Name}} = {{.Value}}
{{- end}}
{{- end}}
)`

var methodTemplate = `package xmmsclient

// auto-generated

import (
	"bytes"
)

{{range .}}
// {{.Doc}}
func (c *Client) {{.Name}}(
	{{- range $index, $arg := .Args}}
		{{- if $index}}, {{end -}}
		{{- $arg.Name}} {{$arg.Type -}}
	{{end -}}) ({{.Return.Type}}, error) {
	__reply := <-c.dispatch({{.ObjectId}}, {{.CommandId}}, XmmsList{
	{{- range $index, $arg := .Args -}}
		{{- if $index}}, {{end -}}
		{{- if len $arg.XmmsType}}
			{{- $arg.XmmsType}}({{$arg.Name -}})
		{{- else -}}
			{{- $arg.Name -}}
		{{- end -}}
	{{- end -}}})
	if __reply.err != nil {
		return {{.Return.Default}}, __reply.err
	}
	__buffer := bytes.NewBuffer(__reply.payload)
	{{ if len .Return.Deserializer -}}
	return {{.Return.Deserializer}}(__buffer)
	{{- else -}}
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return {{.Return.Default}}, __err
	}
	return __value.({{.Return.Type}}), nil
	{{- end}}
}
{{end}}`

var signalTemplate = `package xmmsclient

// auto-generated

import (
	"bytes"
)

{{range .}}
// {{.Doc}}
func (c *Client) {{.Name}}() ({{.Return.Type}}, error) {
	__reply := <-c.dispatch(0, {{.ObjectId}}, XmmsList{XmmsInt({{- .CommandId -}})})
	if __reply.err != nil {
		return {{.Return.Default}}, __reply.err
	}
	__buffer := bytes.NewBuffer(__reply.payload)
	{{ if len .Return.Deserializer -}}
	return {{.Return.Deserializer}}(__buffer)
	{{- else -}}
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return {{.Return.Default}}, __err
	}
	return __value.({{.Return.Type}}), nil
	{{- end}}
}
{{end}}`

var broadcastTemplate = `package xmmsclient

// auto-generated

import (
	"bytes"
)

type IntBroadcast struct {
	result chan reply
}

type StringBroadcast struct {
	result chan reply
}

type DictBroadcast struct {
	result chan reply
}

func (b *IntBroadcast) Next() (XmmsInt, error) {
	__reply := <- b.result
	if __reply.err != nil {
		return -1, __reply.err
	}
	__buffer := bytes.NewBuffer(__reply.payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return -1, __err
	}
	return __value.(XmmsInt), nil
}

func (b *StringBroadcast) Next() (XmmsString, error) {
	__reply := <- b.result
	if __reply.err != nil {
		return "", __reply.err
	}
	__buffer := bytes.NewBuffer(__reply.payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return "", __err
	}
	return __value.(XmmsString), nil
}

func (b *DictBroadcast) Next() (XmmsDict, error) {
	__reply := <- b.result
	if __reply.err != nil {
		return XmmsDict{}, __reply.err
	}
	__buffer := bytes.NewBuffer(__reply.payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return XmmsDict{}, __err
	}
	return __value.(XmmsDict), nil
}

{{range .}}
// {{.Doc}}
func (c *Client) {{.Name}}() {{.Return.Name}}Broadcast {
	__chan := c.dispatch(0, {{.ObjectId}}, XmmsList{XmmsInt({{- .CommandId -}})})
	return {{.Return.Name}}Broadcast{__chan}
}
{{end}}`

func collect(api *Query, template string) interface{} {
	switch template {
	case "enums":
		return collectEnums(api.Enums, api.Offset)
	case "methods":
		return collectFunctions(api.Objects, api.Offset)
	case "signals":
		return collectSignals(api.Objects, api.Offset)
	case "broadcasts":
		return collectBroadcasts(api.Objects, api.Offset)
	default:
		panic("unknown template target")
	}
}

func main() {
	// TODO: flags
	if len(os.Args) != 4 {
		fmt.Println("Missing ipc.xml argument")
		os.Exit(1)
	}

	api, err := parseAPI(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	target := os.Args[2]

	funcMap := template.FuncMap{
		"title": strings.Title,
	}

	tpl := template.New("").Funcs(funcMap)
	tpl = template.Must(tpl.New("enums").Parse(enumTemplate))
	tpl = template.Must(tpl.New("methods").Parse(methodTemplate))
	tpl = template.Must(tpl.New("signals").Parse(signalTemplate))
	tpl = template.Must(tpl.New("broadcasts").Parse(broadcastTemplate))

	data := collect(api, target)

	f, err := os.Create(os.Args[3])
	if err != nil {
		fmt.Println("Fail!", err)
		os.Exit(1)
		return
	}

	err = tpl.ExecuteTemplate(f, target, data)
	if err != nil {
		fmt.Println("Fail!", err)
		os.Exit(1)
		return
	}

	f.Close()
}
