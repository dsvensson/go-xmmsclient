package main

import (
	"fmt"
)

type Arg struct {
	Name        string
	Type        string
	XmmsType    string
	HasXmmsType bool
}

type Function struct {
	ObjectId   int
	CommandId  int
	Name       string
	Args       []Arg
	ReturnType string
}

func collectFunctions(objects []XmlObject, offset int) []Function {
	var functions []Function
	for objectId, obj := range objects {
		for commandId, meth := range obj.Methods {
			var args []Arg
			skip := false
			for _, arg := range meth.Arguments {
				var argType string
				var xmmsType string
				switch arg.Type[0] {
				case "enum-value":
					argType = "int"
					xmmsType = "XmmsInt"
				case "int":
					argType = "int"
					xmmsType = "XmmsInt"
				case "string":
					argType = "string"
					xmmsType = "XmmsString"
				case "binary":
					argType = "XmmsValue"
				case "list":
					argType = "XmmsList" // TODO: Convert to array or vararg
				case "dictionary":
					argType = "XmmsDict"
				case "collection":
					argType = "XmmsValue" // TODO: Implement Collections
				default:
					fmt.Println(arg.Type[0])
				}
				args = append(args, Arg{
					Name:        toCamelCase(arg.Name, false),
					Type:        argType,
					XmmsType:    xmmsType,
					HasXmmsType: len(xmmsType) > 0,
				})

			}
			if !skip {
				functions = append(functions, Function{
					ObjectId:  objectId + 1,
					CommandId: commandId + offset,
					Name:      toCamelCase(obj.Name+"_"+meth.Name, true),
					Args:      args,
				})
			}
		}
		objectId += 1
	}

	return functions
}
