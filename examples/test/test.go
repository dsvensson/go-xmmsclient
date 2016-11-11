package main

import (
	"fmt"
	"github.com/dsvensson/go-xmmsclient/xmmsclient"
	"time"
)

func repeat(client *xmmsclient.Client) {
	for {
		value, err := client.PlaylistListEntries(xmmsclient.ActivePlaylist)
		if err != nil {
			fmt.Println("Fail!", err)
			return
		}
		for position, v := range value.Entries {
			mid := int(v.(xmmsclient.XmmsInt))

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

	value, err := client.PlaylistListEntries("_active")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("  main():", value)

	time.Sleep(time.Second * 2)
	fmt.Println(" close():")
	client.Close()
	fmt.Println(" sleep():")
	time.Sleep(time.Second * 1)
}
