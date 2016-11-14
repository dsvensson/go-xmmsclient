package main

import (
	"fmt"
	"github.com/dsvensson/go-xmmsclient/xmmsclient"
	"strings"
	"time"
)

func repeat(client *xmmsclient.Client) {
	for {
		value, err := client.PlaylistListEntries(xmmsclient.ActivePlaylist)
		if err != nil {
			fmt.Println("Fail!", err)
			return
		}
		for position, mid := range value {
			propDict, err := client.MedialibGetInfo(mid)
			if err != nil {
				fmt.Println("Fail!", err)
				return
			}

			dict, err := xmmsclient.PropDictToDictDefault(propDict)
			if err != nil {
				return
			}

			fmt.Printf("repeat(): [%2d] %s // %s // %s\n",
				position, dict["artist"], dict["album"], dict["title"])
		}
		time.Sleep(time.Millisecond * 1500)
	}
}

func main() {
	client := xmmsclient.NewClient("hello-from-go")

	client.Dial("localhost:xmms2")

	go repeat(client)

	time.Sleep(time.Millisecond * 5)

	value, err := client.CollectionList(xmmsclient.NamespacePlaylists)
	if err != nil {
		fmt.Println(err)
		return
	}

	for index, name := range value {
		if !strings.HasPrefix(name, "_") {
			fmt.Printf("  main(): [%2v] %v::%v\n", index, xmmsclient.NamespacePlaylists, name)
		}
	}

	time.Sleep(time.Second * 2)
	fmt.Println(" close():")
	client.Close()
	fmt.Println(" sleep():")
	time.Sleep(time.Second * 1)
}
