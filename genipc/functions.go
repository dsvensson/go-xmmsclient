package main

type Arg struct {
	Name     string
	Doc      string
	Type     string
	XmmsType string
}

type Return struct {
	Name         string
	Type         string
	Default      string
	Deserializer string
}

type Function struct {
	ObjectID       int
	CommandID      int
	Name           string
	Doc            string
	Args           []Arg
	ResultConsumer string
	Return         Return
}

const (
	DefaultInt    = "0"
	DefaultString = "\"\""
	DefaultPtr    = "nil"
	DefaultList   = "XmmsList{}"
	DefaultDict   = "XmmsDict{}"
)

func skip(object string, method string) bool {
	if object == "visualization" {
		return true
	}
	if object == "coll_sync" {
		return true
	}
	if object == "main" && method == "hello" {
		return true
	}
	return false
}

func collectArguments(arguments []XMLArgument) []Arg {
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
			if len(entry.Type) > 1 {
				switch entry.Type[1] {
				case "string":
					arg.Type = "[]string"
					arg.XmmsType = "XmmsStrings"
				default:
					panic("Unsupported list-type: " + entry.Type[1])
				}
			}
		case "dictionary":
			arg.Type = "XmmsDict"
		case "collection":
			arg.Type = "XmmsColl"
		default:
			panic("Unexpected type: " + entry.Type[0])
		}
		result = append(result, arg)
	}

	return result
}

func collectResultConsumer(signature XMLReturnValue) Return {
	if len(signature.Type) == 0 {
		// TODO: Deal with void functions.
		return Return{Name: "Void", Type: "XmmsValue", Default: DefaultPtr}
	}
	switch signature.Type[0] {
	case "enum-value":
		return Return{Name: "Int", Type: "XmmsInt", Default: DefaultInt}
	case "int":
		return Return{Name: "Int", Type: "XmmsInt", Default: DefaultInt}
	case "string":
		return Return{Name: "String", Type: "XmmsString", Default: DefaultString}
	case "list":
		if len(signature.Type) > 1 {
			switch signature.Type[1] {
			case "int":
				return Return{
					Name: "IntList", Type: "[]int", Default: DefaultPtr, Deserializer: "tryDeserializeIntList",
				}
			case "string":
				return Return{
					Name: "StringList", Type: "[]string", Default: DefaultPtr, Deserializer: "tryDeserializeStringList",
				}
			case "dictionary":
				return Return{
					Name: "DictList", Type: "[]XmmsDict", Default: DefaultPtr, Deserializer: "tryDeserializeDictList",
				}
			}
		}
		return Return{Name: "List", Type: "XmmsList", Default: DefaultList}
	case "dictionary":
		return Return{Name: "Dict", Type: "XmmsDict", Default: DefaultDict}
	default:
		return Return{Name: "Value", Type: "XmmsValue", Default: DefaultPtr}
	}
}

func collectFunctions(objects []XMLObject, offset int) []Function {
	var functions []Function
	for objectID, obj := range objects {
		for commandID, method := range obj.Methods {
			if skip(obj.Name, method.Name) {
				continue
			}
			functions = append(functions, Function{
				ObjectID:  objectID + 1,
				CommandID: commandID + offset,
				Name:      toCamelCase(obj.Name+"_"+method.Name, true),
				Doc:       method.Doc,
				Args:      collectArguments(method.Arguments),
				Return:    collectResultConsumer(method.ReturnValue),
			})
		}
		objectID += 1
	}

	return functions
}

func collectRepeatables(objects []XMLObject, offset int, class int, prefix string) []Function {
	var broadcasts []Function

	signalID := 0
	for _, obj := range objects {
		for _, method := range obj.Broadcasts {
			if !skip(obj.Name, method.Name) && method.ResultClass == class {
				broadcasts = append(broadcasts, Function{
					ObjectID:  offset,
					CommandID: signalID,
					Name:      prefix + toCamelCase(obj.Name+"_"+method.Name, true),
					Doc:       method.Doc,
					Return:    collectResultConsumer(method.ReturnValue),
				})
			}
			signalID += 1
		}
	}

	return broadcasts
}

func collectSignals(objects []XMLObject, offset int) []Function {
	return collectRepeatables(objects, offset, ResultClassSignal, "Signal")
}

func collectBroadcasts(objects []XMLObject, offset int) []Function {
	return collectRepeatables(objects, offset+1, ResultClassBroadcast, "Broadcast")
}
