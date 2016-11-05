package main

import (
	"fmt"
	"github.com/dsvensson/go-xmmsclient/xmmsclient"
	"time"
)

func RepeatSomething(c *xmmsclient.Client) {
	for {
		fmt.Println(c.MainListPlugins())
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	client := xmmsclient.NewClient("hello-from-go")

	client.Dial("localhost:xmms2")

	go RepeatSomething(client)

	fmt.Println("Plugins:")

	fmt.Println(client.MainListPlugins())

	select {}
}
