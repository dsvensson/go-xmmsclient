// auto-generated
package xmmsclient

import (
	"bytes"
)

type Broadcast struct {
	result chan []byte
}

func (b *Broadcast) Next() (XmmsValue, error) {
	__payload := <- b.result
	__buffer := bytes.NewBuffer(__payload)
	return tryDeserialize(__buffer)
}


func (c *Client) BroadcastMainQuit() Broadcast {
	__chan := c.dispatch(0, 33, XmmsList{XmmsInt(0)})
	return Broadcast{__chan}
}

func (c *Client) BroadcastPlaylistChanged() Broadcast {
	__chan := c.dispatch(0, 33, XmmsList{XmmsInt(1)})
	return Broadcast{__chan}
}

func (c *Client) BroadcastPlaylistCurrentPos() Broadcast {
	__chan := c.dispatch(0, 33, XmmsList{XmmsInt(2)})
	return Broadcast{__chan}
}

func (c *Client) BroadcastPlaylistLoaded() Broadcast {
	__chan := c.dispatch(0, 33, XmmsList{XmmsInt(3)})
	return Broadcast{__chan}
}

func (c *Client) BroadcastConfigValueChanged() Broadcast {
	__chan := c.dispatch(0, 33, XmmsList{XmmsInt(4)})
	return Broadcast{__chan}
}

func (c *Client) BroadcastPlaybackStatus() Broadcast {
	__chan := c.dispatch(0, 33, XmmsList{XmmsInt(5)})
	return Broadcast{__chan}
}

func (c *Client) BroadcastPlaybackVolumeChanged() Broadcast {
	__chan := c.dispatch(0, 33, XmmsList{XmmsInt(6)})
	return Broadcast{__chan}
}

func (c *Client) BroadcastPlaybackCurrentId() Broadcast {
	__chan := c.dispatch(0, 33, XmmsList{XmmsInt(7)})
	return Broadcast{__chan}
}

func (c *Client) BroadcastMedialibEntryAdded() Broadcast {
	__chan := c.dispatch(0, 33, XmmsList{XmmsInt(8)})
	return Broadcast{__chan}
}

func (c *Client) BroadcastMedialibEntryChanged() Broadcast {
	__chan := c.dispatch(0, 33, XmmsList{XmmsInt(9)})
	return Broadcast{__chan}
}

func (c *Client) BroadcastMedialibEntryRemoved() Broadcast {
	__chan := c.dispatch(0, 33, XmmsList{XmmsInt(10)})
	return Broadcast{__chan}
}

func (c *Client) BroadcastCollectionChanged() Broadcast {
	__chan := c.dispatch(0, 33, XmmsList{XmmsInt(11)})
	return Broadcast{__chan}
}

func (c *Client) BroadcastMediainfoReaderStatus() Broadcast {
	__chan := c.dispatch(0, 33, XmmsList{XmmsInt(12)})
	return Broadcast{__chan}
}

func (c *Client) BroadcastCourierMessage() Broadcast {
	__chan := c.dispatch(0, 33, XmmsList{XmmsInt(13)})
	return Broadcast{__chan}
}

func (c *Client) BroadcastCourierReady() Broadcast {
	__chan := c.dispatch(0, 33, XmmsList{XmmsInt(14)})
	return Broadcast{__chan}
}

func (c *Client) BroadcastIpcManagerClientConnected() Broadcast {
	__chan := c.dispatch(0, 33, XmmsList{XmmsInt(15)})
	return Broadcast{__chan}
}

func (c *Client) BroadcastIpcManagerClientDisconnected() Broadcast {
	__chan := c.dispatch(0, 33, XmmsList{XmmsInt(16)})
	return Broadcast{__chan}
}
