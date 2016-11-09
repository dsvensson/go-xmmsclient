package main

import (
	"fmt"
	"github.com/dsvensson/go-xmmsclient/xmmsclient"
	"time"
)

func repeat(c *xmmsclient.Client) {
	for {
		value, err := c.PlaylistListEntries("_active")
		if err != nil {
			return
		}
		fmt.Println("repeat():", value)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	client := xmmsclient.NewClient("hello-from-go")

	client.Dial("localhost:xmms2")

	go repeat(client)

	value, err := client.PlaylistListEntries("_active")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("  main():", value)

	time.Sleep(time.Second * 2)
	client.Close()
	time.Sleep(time.Second * 1)
}
