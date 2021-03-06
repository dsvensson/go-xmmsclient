package main

type EnumValue struct {
	Name  string
	Value int
}

func collectEnums(enums []XMLEnum, offset int) map[string][]EnumValue {
	result := make(map[string][](EnumValue))

	for _, enum := range enums {
		index := 0
		for _, member := range enum.Members {
			var name string
			if len(enum.Hint) > 0 {
				name = enum.Hint + "_" + member
			} else {
				name = enum.Name + "_" + member
			}
			if enum.Name == "IPC_COMMAND" {
				index = offset
			}
			result[enum.Name] = append(result[enum.Name], EnumValue{toCamelCase(name, true), index})
			index++
		}
	}

	return result
}
