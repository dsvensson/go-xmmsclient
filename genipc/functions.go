package main

type Arg struct {
	Name     string
	Doc      string
	Type     string
	XmmsType string
}

type Return struct {
	Type         string
	Default      string
	Deserializer string
}

type Function struct {
	ObjectId       int
	CommandId      int
	Name           string
	Doc            string
	Args           []Arg
	ResultConsumer string
	Return         Return
}

type Broadcast struct {
	ObjectId int
	SignalId int
	Name     string
	Doc      string
	Return   Return
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
			arg.Type = "XmmsColl"
		default:
			panic("Unexpected type")
		}
		result = append(result, arg)
	}

	return result
}

func collectResultConsumer(signature XmlReturnValue) Return {
	if len(signature.Type) == 0 {
		// TODO: Deal with void functions.
		return Return{Type: "XmmsValue", Default: DefaultPtr}
	}
	switch signature.Type[0] {
	case "enum-value":
		return Return{Type: "XmmsInt", Default: DefaultInt}
	case "int":
		return Return{Type: "XmmsInt", Default: DefaultInt}
	case "string":
		return Return{Type: "XmmsString", Default: DefaultString}
	case "list":
		if len(signature.Type) > 1 {
			switch signature.Type[1] {
			case "int":
				return Return{
					Type: "[]int", Default: DefaultPtr, Deserializer: "tryDeserializeIntList"}
			case "string":
				return Return{
					Type: "[]string", Default: DefaultPtr, Deserializer: "tryDeserializeStringList"}
			case "dictionary":
				return Return{
					Type: "[]XmmsDict", Default: DefaultPtr, Deserializer: "tryDeserializeDictList"}
			}
		}
		return Return{Type: "XmmsList", Default: DefaultList}
	case "dictionary":
		return Return{Type: "XmmsDict", Default: DefaultDict}
	default:
		return Return{Type: "XmmsValue", Default: DefaultPtr}
	}
}

func collectFunctions(objects []XmlObject, offset int) []Function {
	var functions []Function
	for objectId, obj := range objects {
		for commandId, method := range obj.Methods {
			if skip(obj.Name, method.Name) {
				continue
			}
			functions = append(functions, Function{
				ObjectId:  objectId + 1,
				CommandId: commandId + offset,
				Name:      toCamelCase(obj.Name+"_"+method.Name, true),
				Doc:       method.Doc,
				Args:      collectArguments(method.Arguments),
				Return:    collectResultConsumer(method.ReturnValue),
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
			if skip(obj.Name, broadcast.Name) {
				continue
			}
			broadcasts = append(broadcasts, Broadcast{
				ObjectId: offset,
				SignalId: signalId,
				Name:     toCamelCase(obj.Name+"_"+broadcast.Name, true),
				Doc:      broadcast.Doc,
				Return:   collectResultConsumer(broadcast.ReturnValue),
			})
			signalId += 1
		}
	}

	return broadcasts
}
