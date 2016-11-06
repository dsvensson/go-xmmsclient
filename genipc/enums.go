package main

type EnumValue struct {
	name  string
	value int
}

func collectEnums(enums []XmlEnum) map[string][]EnumValue {
	var result = make(map[string][](EnumValue))

	for _, enum := range enums {
		index := 0
		for _, member := range enum.Members {
			var name string
			if len(enum.Hint) > 0 {
				name = enum.Hint + "_" + member
			} else {
				name = enum.Name + "_" + member
			}
			result[enum.Name] = append(result[enum.Name], EnumValue{toCamelCase(name, true), index})
			index += 1
		}
	}

	return result
}
