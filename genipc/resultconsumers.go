package main

type Consumer struct {
	Name         string
	ResultType   string
	DefaultValue string
	Cast         string
}

func collectResultConsumers() []Consumer {
	return []Consumer{
		Consumer{
			Name:         "generic",
			ResultType:   "XmmsValue",
			DefaultValue: "value",
			Cast:         "value",
		},
		Consumer{
			Name:         "int",
			ResultType:   "XmmsInt",
			DefaultValue: "0",
			Cast:         "value.(XmmsInt)",
		},
		Consumer{
			Name:         "string",
			ResultType:   "XmmsString",
			DefaultValue: "\"\"",
			Cast:         "value.(XmmsString)",
		},
		Consumer{
			Name:         "list",
			ResultType:   "XmmsList",
			DefaultValue: "XmmsList{}",
			Cast:         "value.(XmmsList)",
		},
		Consumer{
			Name:         "dict",
			ResultType:   "XmmsDict",
			DefaultValue: "XmmsDict{}",
			Cast:         "value.(XmmsDict)",
		},
		Consumer{
			Name:         "coll",
			ResultType:   "XmmsColl",
			DefaultValue: "XmmsColl{}",
			Cast:         "value.(XmmsColl)",
		},
	}
}
