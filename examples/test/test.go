package main

import (
	"fmt"
	"github.com/dsvensson/go-xmmsclient/xmmsclient"
)

func main() {
	client, _ := xmmsclient.Dial("localhost:xmms2")

	fmt.Println("Plugins:")
	fmt.Println(client.MainListPlugins())

	select {}
}
