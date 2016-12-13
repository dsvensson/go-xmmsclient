package xmmsclient

//go:generate go run ../genipc/enums.go ../genipc/functions.go ../genipc/genipc.go ../genipc/xml.go ../data/ipc.xml enums auto_consts.go

//go:generate go run ../genipc/enums.go ../genipc/functions.go ../genipc/genipc.go ../genipc/xml.go ../data/ipc.xml methods auto_commands.go

//go:generate go run ../genipc/enums.go ../genipc/functions.go ../genipc/genipc.go ../genipc/xml.go ../data/ipc.xml broadcasts auto_broadcasts.go

//go:generate go run ../genipc/enums.go ../genipc/functions.go ../genipc/genipc.go ../genipc/xml.go ../data/ipc.xml signals auto_signals.go
