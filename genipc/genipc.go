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

var enumTemplate = `// auto-generated
package xmmsclient

const (
{{- range $key, $values := .}}
{{ range $values}}
	{{.Name}} = {{.Value}}
{{- end}}
{{- end}}
)`

var methodTemplate = `// auto-generated
package xmmsclient

import (
	"bytes"
)

{{range .}}
// {{.Doc}}
func (c *Client) {{.Name}}(
	{{- range $index, $arg := .Args}}
		{{- if $index}}, {{end -}}
		{{- $arg.Name}} {{$arg.Type -}}
	{{end -}}) ({{.ReturnType}}, error) {
	__payload := <-c.dispatch({{.ObjectId}}, {{.CommandId}}, XmmsList{
	{{- range $index, $arg := .Args -}}
		{{- if $index}}, {{end -}}
		{{- if len $arg.XmmsType}}
			{{- $arg.XmmsType}}({{$arg.Name -}})
		{{- else -}}
			{{- $arg.Name -}}
		{{- end -}}
	{{- end -}}})
	__buffer := bytes.NewBuffer(__payload)
	{{ if len .Deserializer -}}
	return {{.Deserializer}}(__buffer)
	{{- else -}}
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return {{.DefaultValue}}, __err
	}
	return __value.({{.ReturnType}}), nil
	{{- end}}
}
{{end}}`

var broadcastTemplate = `// auto-generated
package xmmsclient

import (
	"bytes"
)

type Broadcast struct {
	result chan []byte
}

func (b *Broadcast) Next() (XmmsValue, error) {
	__payload := <- b.result
	__buffer := bytes.NewBuffer(__payload)
	return tryDeserialize(__buffer)
}

{{range .}}
// {{.Doc}}
func (c *Client) Broadcast{{.Name}}() Broadcast {
	__chan := c.dispatch(0, 33, XmmsList{XmmsInt({{- .SignalId -}})})
	return Broadcast{__chan}
}
{{end}}`

func collect(api *Query, template string) interface{} {
	switch template {
	case "enums":
		return collectEnums(api.Enums, api.Offset)
	case "methods":
		return collectFunctions(api.Objects, api.Offset)
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
