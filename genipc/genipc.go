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
{{range .}}
func (c *Client) {{.Name}}(
	{{- range $index, $arg := .Args}}
		{{- if $index}}, {{end -}}
		{{- $arg.Name}} {{$arg.Type -}}
	{{end -}}) ({{.ReturnType}}, error) {
	consumer := new{{title .ResultConsumer}}Consumer()
	c.dispatch(&consumer, {{.ObjectId}}, {{.CommandId}}, NewXmmsList(
	{{- range $index, $arg := .Args -}}
		{{- if $index}}, {{end -}}
		{{- if len $arg.XmmsType}}
			{{- $arg.XmmsType}}({{$arg.Name -}})
		{{- else -}}
			{{- $arg.Name -}}
		{{- end -}}
	{{- end -}}))
	result := <-consumer.result
	return result.value, result.err
}
{{end}}`

var resultConsumerTemplate = `// auto-generated
package xmmsclient
{{range .}}
type {{.Name}}ConsumerType struct {
	value {{.ResultType}}
	err   error
}

type {{.Name}}Consumer struct {
	result chan {{.Name}}ConsumerType
}

func new{{title .Name}}Consumer() {{.Name}}Consumer {
	return {{.Name}}Consumer{make(chan {{.Name}}ConsumerType)}
}

func (r *{{.Name}}Consumer) post(value XmmsValue, err error) {
	if err != nil {
		r.result <- {{.Name}}ConsumerType{ {{- .DefaultValue}}, err}
	} else {
		r.result <- {{.Name}}ConsumerType{ {{- .Cast}}, err}
	}
}
{{end}}`

func collect(api *Query, template string) interface{} {
	switch template {
	case "enums":
		return collectEnums(api.Enums)
	case "methods":
		return collectFunctions(api.Objects, api.Offset)
	case "consumers":
		return collectResultConsumers()
	default:
		panic("unknown template target")
	}
}

func main() {
	// TODO: flags
	if len(os.Args) != 3 {
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
	tpl = template.Must(tpl.New("consumers").Parse(resultConsumerTemplate))

	data := collect(api, target)

	err = tpl.ExecuteTemplate(os.Stdout, target, data)
	if err != nil {
		fmt.Println("Fail!", err)
		os.Exit(1)
		return
	}
}
