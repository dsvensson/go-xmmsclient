package main

import (
	"fmt"
	"github.com/dsvensson/go-xmmsclient/xmmsclient"
	"strings"
	"time"
)

func playlistChanges(client *xmmsclient.Client) {
	bcast := client.BroadcastPlaylistChanged()
	for {
		value, err := bcast.Next()
		if err != nil {
			fmt.Println("Error(PlaylistChanged):", err)
			return
		}
		fmt.Println(value)
	}
}

func repeat(client *xmmsclient.Client) {
	for {
		value, err := client.PlaylistListEntries(xmmsclient.ActivePlaylist)
		if err != nil {
			fmt.Println("Error(PlaylistListEntries):", err)
			return
		}
		for position, mid := range value {
			propDict, err := client.MedialibGetInfo(mid)
			if err != nil {
				fmt.Println("Error(GetInfo):", err)
				return
			}

			dict, err := xmmsclient.PropDictToDictDefault(propDict)
			if err != nil {
				fmt.Println("Error(PropDict->Dict):", err)
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

	clientId, err := client.Dial("localhost:xmms2")
	if err != nil {
		fmt.Println("Error(Dial):", err)
		return
	}

	go repeat(client)
	go playlistChanges(client)

	time.Sleep(time.Millisecond * 5)

	value, err := client.CollectionList(xmmsclient.NamespacePlaylists)
	if err != nil {
		fmt.Println("Error(CollectionList):", err)
		return
	}

	for index, name := range value {
		if !strings.HasPrefix(name, "_") {
			fmt.Printf("  main(): [%2v] %v::%v\n", index, xmmsclient.NamespacePlaylists, name)
		}
	}

	clients, err := client.CourierGetConnectedClients()
	if err != nil {
		fmt.Println("Error(CourierGetConnectedClients):", err)
		return
	}
	fmt.Println("Connected clients:", clients, "self:", clientId)

	coll := xmmsclient.XmmsColl{Type: xmmsclient.CollectionTypeUniverse}
	fetch := xmmsclient.XmmsList{xmmsclient.XmmsString("artist"), xmmsclient.XmmsString("album")}
	group := xmmsclient.XmmsList{xmmsclient.XmmsString("album")}

	matches, err := client.CollectionQueryInfos(coll, 0, 0, fetch, group)
	if err != nil {
		fmt.Println("Error(CollectionQueryInfos):", err)
		return
	}
	fmt.Println("All albums:", matches)

	time.Sleep(time.Second * 3)
	fmt.Println(" close():")
	client.Close()
	fmt.Println(" sleep():")
	time.Sleep(time.Second * 2)
}
