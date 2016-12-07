package main

type Arg struct {
	Name     string
	Doc      string
	Type     string
	XmmsType string
}

type Function struct {
	ObjectId       int
	CommandId      int
	Name           string
	Doc            string
	Args           []Arg
	ResultConsumer string
	ReturnType     string
	DefaultValue   string
	Deserializer   string
}

type Broadcast struct {
	ObjectId   int
	SignalId   int
	Name       string
	Doc        string
	ReturnType string
}

func collectArguments(arguments []XmlArgument) []Arg {
	var result []Arg

	for _, entry := range arguments {
		arg := Arg{
			Name: toCamelCase(entry.Name, false),
			Doc:  entry.Doc,
		}
		switch entry.Type[0] {
		case "enum-value":
			arg.Type = "int"
			arg.XmmsType = "XmmsInt"
		case "int":
			arg.Type = "int"
			arg.XmmsType = "XmmsInt"
		case "string":
			arg.Type = "string"
			arg.XmmsType = "XmmsString"
		case "binary":
			arg.Type = "XmmsValue"
		case "list":
			arg.Type = "XmmsList" // TODO: Convert to array or vararg
		case "dictionary":
			arg.Type = "XmmsDict"
		case "collection":
			arg.Type = "XmmsValue" // TODO: Implement Collections
		default:
			panic("Unexpected type")
		}
		result = append(result, arg)
	}

	return result
}

func collectResultConsumer(signature XmlReturnValue) (string, string, string) {
	if len(signature.Type) == 0 {
		// TODO: Deal with void functions.
		return "XmmsValue", "nil", ""
	}
	switch signature.Type[0] {
	case "enum-value":
		return "XmmsInt", "0", ""
	case "int":
		return "XmmsInt", "0", ""
	case "string":
		return "XmmsString", "\"\"", ""
	case "list":
		if len(signature.Type) > 1 {
			switch signature.Type[1] {
			case "int":
				return "[]int", "nil", "tryDeserializeIntList"
			case "string":
				return "[]string", "nil", "tryDeserializeStringList"
			case "dictionary":
				return "[]XmmsDict", "nil", "tryDeserializeDictList"
			}
		}
		return "XmmsList", "XmmsList{}", ""
	case "dictionary":
		return "XmmsDict", "XmmsDict{}", ""
	default:
		return "XmmsValue", "nil", ""
	}
}

func collectFunctions(objects []XmlObject, offset int) []Function {
	var functions []Function
	for objectId, obj := range objects {
		for commandId, method := range obj.Methods {
			returnType, defaultValue, deserializer := collectResultConsumer(method.ReturnValue)
			functions = append(functions, Function{
				ObjectId:     objectId + 1,
				CommandId:    commandId + offset,
				Name:         toCamelCase(obj.Name+"_"+method.Name, true),
				Doc:          method.Doc,
				Args:         collectArguments(method.Arguments),
				DefaultValue: defaultValue,
				ReturnType:   returnType,
				Deserializer: deserializer,
			})
		}
		objectId += 1
	}

	return functions
}

func collectBroadcasts(objects []XmlObject, offset int) []Broadcast {
	var broadcasts []Broadcast

	signalId := 0
	for _, obj := range objects {
		for _, broadcast := range obj.Broadcasts {
			returnType, _, _ := collectResultConsumer(broadcast.ReturnValue)
			broadcasts = append(broadcasts, Broadcast{
				ObjectId:   offset,
				SignalId:   signalId,
				Name:       toCamelCase(obj.Name+"_"+broadcast.Name, true),
				Doc:        broadcast.Doc,
				ReturnType: returnType,
			})
			signalId += 1
		}
	}

	return broadcasts
}
