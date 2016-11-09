package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
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

var methodTemplate = `package xmmsclient
{{range .}}
func (c *Client) {{.Name}}(
	{{- range $index, $arg := .Args}}
		{{- if $index}}, {{end -}}
		{{- $arg.Name}} {{$arg.Type -}}
	{{end -}}) (XmmsValue, error) {
	result := <-c.dispatch({{.ObjectId}}, {{.CommandId}}, NewXmmsList(
	{{- range $index, $arg := .Args -}}
		{{- if $index}}, {{end -}}
		{{- if $arg.HasXmmsType}}
			{{- $arg.XmmsType}}({{$arg.Name -}})
		{{- else -}}
			{{- $arg.Name -}}
		{{- end -}}
	{{- end -}}))
	return result.value, result.err
}
{{end}}`

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Missing ipc.xml argument")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		os.Exit(1)
		return
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		os.Exit(1)
		return
	}

	var q Query
	err = xml.Unmarshal(data, &q)
	if err != nil {
		os.Exit(1)
		return
	}

	// var enums = collectEnums(q.Enums)

	tpl, err := template.New("method").Parse(methodTemplate)
	if err != nil {
		os.Exit(1)
		return
	}

	err = tpl.Execute(os.Stdout, collectFunctions(q.Objects, q.Offset))
	if err != nil {
		os.Exit(1)
		return
	}

}
