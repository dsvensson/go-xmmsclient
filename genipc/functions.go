package main

type Arg struct {
	Name        string
	Type        string
	XmmsType    string
	HasXmmsType bool
}

type Function struct {
	ObjectId      int
	CommandId     int
	Name          string
	Args          []Arg
	ReturnType    string
	ReturnCast    string
	ReturnFail    string
	HasReturnCast bool
}

func collectArguments(arguments []XmlArgument) []Arg {
	var result []Arg

	for _, arg := range arguments {
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
			panic("Unexpected type")
		}
		result = append(result, Arg{
			Name:        toCamelCase(arg.Name, false),
			Type:        argType,
			XmmsType:    xmmsType,
			HasXmmsType: len(xmmsType) > 0,
		})
	}

	return result
}

func collectReturnType(signature XmlReturnValue) (string, string, string) {
	if len(signature.Type) == 0 {
		// TODO: Deal with void functions.
		return "XmmsValue", "", "result.value"
	}
	switch signature.Type[0] {
	case "enum-value":
		return "XmmsInt", "valueAsInt", "0"
	case "int":
		return "XmmsInt", "valueAsInt", "0"
	case "string":
		return "XmmsString", "valueAsString", "\"\""
	case "list":
		return "XmmsList", "valueAsList", "XmmsList{}"
	case "dictionary":
		return "XmmsDict", "valueAsDict", "XmmsDict{}"
	default:
		return "XmmsValue", "", "nil"
	}
}

func collectFunctions(objects []XmlObject, offset int) []Function {
	var functions []Function
	for objectId, obj := range objects {
		for commandId, method := range obj.Methods {
			returnType, returnCast, returnFail := collectReturnType(method.ReturnValue)
			functions = append(functions, Function{
				ObjectId:      objectId + 1,
				CommandId:     commandId + offset,
				Name:          toCamelCase(obj.Name+"_"+method.Name, true),
				Args:          collectArguments(method.Arguments),
				ReturnType:    returnType,
				ReturnCast:    returnCast,
				ReturnFail:    returnFail,
				HasReturnCast: returnCast != "",
			})
		}
		objectId += 1
	}

	return functions
}
