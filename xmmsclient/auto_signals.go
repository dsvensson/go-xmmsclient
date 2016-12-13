package xmmsclient

// auto-generated

import (
	"bytes"
)


// Emits the current playtime.
func (c *Client) SignalPlaybackPlaytime() (XmmsInt, error) {
	__reply := <-c.dispatch(0, 32, XmmsList{XmmsInt(8)})
	if __reply.err != nil {
		return 0, __reply.err
	}
	__buffer := bytes.NewBuffer(__reply.payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return 0, __err
	}
	return __value.(XmmsInt), nil
}

// Emits the number of unresolved medialib entries.
func (c *Client) SignalMediainfoReaderUnindexed() (XmmsInt, error) {
	__reply := <-c.dispatch(0, 32, XmmsList{XmmsInt(14)})
	if __reply.err != nil {
		return 0, __reply.err
	}
	__buffer := bytes.NewBuffer(__reply.payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return 0, __err
	}
	return __value.(XmmsInt), nil
}
