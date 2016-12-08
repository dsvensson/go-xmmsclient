// auto-generated
package xmmsclient

import (
	"bytes"
)


// Shuts down the daemon.
func (c *Client) MainQuit() (XmmsValue, error) {
	__payload := <-c.dispatch(1, 33, XmmsList{})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Retrieves the list of available plugins.
func (c *Client) MainListPlugins(pluginType int) ([]XmmsDict, error) {
	__payload := <-c.dispatch(1, 34, XmmsList{XmmsInt(pluginType)})
	__buffer := bytes.NewBuffer(__payload)
	return tryDeserializeDictList(__buffer)
}

// Retrieves statistics from the server.
func (c *Client) MainStats() (XmmsDict, error) {
	__payload := <-c.dispatch(1, 35, XmmsList{})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return XmmsDict{}, __err
	}
	return __value.(XmmsDict), nil
}

// Queries ids from a collection and replaces the playlist with the result.
func (c *Client) PlaylistReplace(name string, replacement XmmsColl, action int) (XmmsValue, error) {
	__payload := <-c.dispatch(2, 32, XmmsList{XmmsString(name), replacement, XmmsInt(action)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Sets the playlist entry that will be played next.
func (c *Client) PlaylistSetNext(position int) (XmmsInt, error) {
	__payload := <-c.dispatch(2, 33, XmmsList{XmmsInt(position)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return 0, __err
	}
	return __value.(XmmsInt), nil
}

// Sets the playlist entry that will be played next.
func (c *Client) PlaylistSetNextRel(positionDelta int) (XmmsInt, error) {
	__payload := <-c.dispatch(2, 34, XmmsList{XmmsInt(positionDelta)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return 0, __err
	}
	return __value.(XmmsInt), nil
}

// Adds an URL to the given playlist.
func (c *Client) PlaylistAddUrl(name string, url string) (XmmsValue, error) {
	__payload := <-c.dispatch(2, 35, XmmsList{XmmsString(name), XmmsString(url)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Adds the contents of a collection to the given playlist.
func (c *Client) PlaylistAddCollection(name string, collection XmmsColl) (XmmsValue, error) {
	__payload := <-c.dispatch(2, 36, XmmsList{XmmsString(name), collection})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Removes an entry from the given playlist.
func (c *Client) PlaylistRemoveEntry(name string, position int) (XmmsValue, error) {
	__payload := <-c.dispatch(2, 37, XmmsList{XmmsString(name), XmmsInt(position)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Moves a playlist entry to a new position (absolute move).
func (c *Client) PlaylistMoveEntry(name string, position int, newPosition int) (XmmsValue, error) {
	__payload := <-c.dispatch(2, 38, XmmsList{XmmsString(name), XmmsInt(position), XmmsInt(newPosition)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Lists the contents of the given playlist.
func (c *Client) PlaylistListEntries(name string) ([]int, error) {
	__payload := <-c.dispatch(2, 39, XmmsList{XmmsString(name)})
	__buffer := bytes.NewBuffer(__payload)
	return tryDeserializeIntList(__buffer)
}

// Retrieves the current position in the playlist with the given name.
func (c *Client) PlaylistCurrentPos(name string) (XmmsDict, error) {
	__payload := <-c.dispatch(2, 40, XmmsList{XmmsString(name)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return XmmsDict{}, __err
	}
	return __value.(XmmsDict), nil
}

// Retrieves the name of the currently active playlist.
func (c *Client) PlaylistCurrentActive() (XmmsString, error) {
	__payload := <-c.dispatch(2, 41, XmmsList{})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return "", __err
	}
	return __value.(XmmsString), nil
}

// Inserts an URL into the given playlist.
func (c *Client) PlaylistInsertUrl(name string, position int, url string) (XmmsValue, error) {
	__payload := <-c.dispatch(2, 42, XmmsList{XmmsString(name), XmmsInt(position), XmmsString(url)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Inserts the contents of a collection into the given playlist.
func (c *Client) PlaylistInsertCollection(name string, position int, collection XmmsColl) (XmmsValue, error) {
	__payload := <-c.dispatch(2, 43, XmmsList{XmmsString(name), XmmsInt(position), collection})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Loads the playlist with the given name.
func (c *Client) PlaylistLoad(name string) (XmmsValue, error) {
	__payload := <-c.dispatch(2, 44, XmmsList{XmmsString(name)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Adds a directory recursively to the playlist with the given name.
func (c *Client) PlaylistRadd(name string, url string) (XmmsValue, error) {
	__payload := <-c.dispatch(2, 45, XmmsList{XmmsString(name), XmmsString(url)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Insert a directory recursively into the playlist with the given name at the given position.
func (c *Client) PlaylistRinsert(name string, position int, url string) (XmmsValue, error) {
	__payload := <-c.dispatch(2, 46, XmmsList{XmmsString(name), XmmsInt(position), XmmsString(url)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Retrieves the value of the config property with the given key.
func (c *Client) ConfigGetValue(key string) (XmmsString, error) {
	__payload := <-c.dispatch(3, 32, XmmsList{XmmsString(key)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return "", __err
	}
	return __value.(XmmsString), nil
}

// Sets the value of the config property with the given key.
func (c *Client) ConfigSetValue(key string, value string) (XmmsValue, error) {
	__payload := <-c.dispatch(3, 33, XmmsList{XmmsString(key), XmmsString(value)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Registers a new config property for the connected client.
func (c *Client) ConfigRegisterValue(key string, value string) (XmmsString, error) {
	__payload := <-c.dispatch(3, 34, XmmsList{XmmsString(key), XmmsString(value)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return "", __err
	}
	return __value.(XmmsString), nil
}

// Retrieves the list of known config properties.
func (c *Client) ConfigListValues() (XmmsDict, error) {
	__payload := <-c.dispatch(3, 35, XmmsList{})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return XmmsDict{}, __err
	}
	return __value.(XmmsDict), nil
}

// Starts playback.
func (c *Client) PlaybackStart() (XmmsValue, error) {
	__payload := <-c.dispatch(4, 32, XmmsList{})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Stops playback.
func (c *Client) PlaybackStop() (XmmsValue, error) {
	__payload := <-c.dispatch(4, 33, XmmsList{})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Pauses playback.
func (c *Client) PlaybackPause() (XmmsValue, error) {
	__payload := <-c.dispatch(4, 34, XmmsList{})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Stops decoding of the current song. This will start decoding of the song set with the playlist_set_next command or the current song again if the playlist_set_next command wasn't executed.
func (c *Client) PlaybackTickle() (XmmsValue, error) {
	__payload := <-c.dispatch(4, 35, XmmsList{})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Retrieves the current playtime.
func (c *Client) PlaybackPlaytime() (XmmsInt, error) {
	__payload := <-c.dispatch(4, 36, XmmsList{})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return 0, __err
	}
	return __value.(XmmsInt), nil
}

// Seeks to a position in the currently played song (given in milliseconds).
func (c *Client) PlaybackSeekMs(offset int, whence int) (XmmsValue, error) {
	__payload := <-c.dispatch(4, 37, XmmsList{XmmsInt(offset), XmmsInt(whence)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Seeks to a position in the currently played song (given in samples).
func (c *Client) PlaybackSeekSamples(offset int, whence int) (XmmsValue, error) {
	__payload := <-c.dispatch(4, 38, XmmsList{XmmsInt(offset), XmmsInt(whence)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Retrieves the current playback status.
func (c *Client) PlaybackStatus() (XmmsInt, error) {
	__payload := <-c.dispatch(4, 39, XmmsList{})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return 0, __err
	}
	return __value.(XmmsInt), nil
}

// Retrieves the ID of the song that's currently being played.
func (c *Client) PlaybackCurrentId() (XmmsInt, error) {
	__payload := <-c.dispatch(4, 40, XmmsList{})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return 0, __err
	}
	return __value.(XmmsInt), nil
}

// Changes the volume for the given channel.
func (c *Client) PlaybackVolumeSet(channel string, volume int) (XmmsValue, error) {
	__payload := <-c.dispatch(4, 41, XmmsList{XmmsString(channel), XmmsInt(volume)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Retrieves the volume of all available channels.
func (c *Client) PlaybackVolumeGet() (XmmsDict, error) {
	__payload := <-c.dispatch(4, 42, XmmsList{})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return XmmsDict{}, __err
	}
	return __value.(XmmsDict), nil
}

// Retrieves information about a medialib entry.
func (c *Client) MedialibGetInfo(id int) (XmmsDict, error) {
	__payload := <-c.dispatch(5, 32, XmmsList{XmmsInt(id)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return XmmsDict{}, __err
	}
	return __value.(XmmsDict), nil
}

// Adds a directory recursively to the medialib.
func (c *Client) MedialibImportPath(directory string) (XmmsValue, error) {
	__payload := <-c.dispatch(5, 33, XmmsList{XmmsString(directory)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Rehashes the medialib. This will make sure that the data in the medialib is the same as the data in the files. 
func (c *Client) MedialibRehash(id int) (XmmsValue, error) {
	__payload := <-c.dispatch(5, 34, XmmsList{XmmsInt(id)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Retrieves the medialib ID that belongs to the given URL.
func (c *Client) MedialibGetId(url string) (XmmsInt, error) {
	__payload := <-c.dispatch(5, 35, XmmsList{XmmsString(url)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return 0, __err
	}
	return __value.(XmmsInt), nil
}

// Removes an entry from the medialib.
func (c *Client) MedialibRemoveEntry(id int) (XmmsValue, error) {
	__payload := <-c.dispatch(5, 36, XmmsList{XmmsInt(id)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Sets a medialib property to a string value.
func (c *Client) MedialibSetPropertyString(id int, source string, key string, value string) (XmmsValue, error) {
	__payload := <-c.dispatch(5, 37, XmmsList{XmmsInt(id), XmmsString(source), XmmsString(key), XmmsString(value)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Sets a medialib property to an integer value.
func (c *Client) MedialibSetPropertyInt(id int, source string, key string, value int) (XmmsValue, error) {
	__payload := <-c.dispatch(5, 38, XmmsList{XmmsInt(id), XmmsString(source), XmmsString(key), XmmsInt(value)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Removes a propert from a medialib entry.
func (c *Client) MedialibRemoveProperty(id int, source string, key string) (XmmsValue, error) {
	__payload := <-c.dispatch(5, 39, XmmsList{XmmsInt(id), XmmsString(source), XmmsString(key)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Updates the URL of a medialib entry that has been moved to a new location.
func (c *Client) MedialibMoveEntry(id int, url string) (XmmsValue, error) {
	__payload := <-c.dispatch(5, 40, XmmsList{XmmsInt(id), XmmsString(url)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Add the given URL to the medialib.
func (c *Client) MedialibAddEntry(url string) (XmmsValue, error) {
	__payload := <-c.dispatch(5, 41, XmmsList{XmmsString(url)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Retrieves the structure of a given collection.
func (c *Client) CollectionGet(name string, namespace string) (XmmsValue, error) {
	__payload := <-c.dispatch(6, 32, XmmsList{XmmsString(name), XmmsString(namespace)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Lists the collections in the given namespace.
func (c *Client) CollectionList(namespace string) ([]string, error) {
	__payload := <-c.dispatch(6, 33, XmmsList{XmmsString(namespace)})
	__buffer := bytes.NewBuffer(__payload)
	return tryDeserializeStringList(__buffer)
}

// Save the given collection in the DAG under the given name in the given namespace.
func (c *Client) CollectionSave(name string, namespace string, collection XmmsColl) (XmmsValue, error) {
	__payload := <-c.dispatch(6, 34, XmmsList{XmmsString(name), XmmsString(namespace), collection})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Remove the given collection from the DAG.
func (c *Client) CollectionRemove(name string, namespace string) (XmmsValue, error) {
	__payload := <-c.dispatch(6, 35, XmmsList{XmmsString(name), XmmsString(namespace)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Find all collections in the given namespace that contain a given media.
func (c *Client) CollectionFind(id int, namespace string) ([]string, error) {
	__payload := <-c.dispatch(6, 36, XmmsList{XmmsInt(id), XmmsString(namespace)})
	__buffer := bytes.NewBuffer(__payload)
	return tryDeserializeStringList(__buffer)
}

// Rename a collection in the given namespace.
func (c *Client) CollectionRename(name string, newName string, namespace string) (XmmsValue, error) {
	__payload := <-c.dispatch(6, 37, XmmsList{XmmsString(name), XmmsString(newName), XmmsString(namespace)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// FIXME.
func (c *Client) CollectionQuery(collection XmmsColl, fetch XmmsDict) (XmmsValue, error) {
	__payload := <-c.dispatch(6, 38, XmmsList{collection, fetch})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// FIXME.
func (c *Client) CollectionQueryInfos(collection XmmsColl, limitStart int, limitLength int, properties XmmsList, groupBy XmmsList) ([]XmmsDict, error) {
	__payload := <-c.dispatch(6, 39, XmmsList{collection, XmmsInt(limitStart), XmmsInt(limitLength), properties, groupBy})
	__buffer := bytes.NewBuffer(__payload)
	return tryDeserializeDictList(__buffer)
}

// FIXME.
func (c *Client) CollectionIdlistFromPlaylist(url string) (XmmsValue, error) {
	__payload := <-c.dispatch(6, 40, XmmsList{XmmsString(url)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Retrieves a list of paths available (directly) under the given path.
func (c *Client) XformBrowse(url string) ([]XmmsDict, error) {
	__payload := <-c.dispatch(9, 32, XmmsList{XmmsString(url)})
	__buffer := bytes.NewBuffer(__payload)
	return tryDeserializeDictList(__buffer)
}

// Retrieves a file from the server's bindata directory given the file's hash.
func (c *Client) BindataRetrieve(hash string) (XmmsValue, error) {
	__payload := <-c.dispatch(10, 32, XmmsList{XmmsString(hash)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Adds binary data to the server's bindata directory.
func (c *Client) BindataAdd(rawData XmmsValue) (XmmsString, error) {
	__payload := <-c.dispatch(10, 33, XmmsList{rawData})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return "", __err
	}
	return __value.(XmmsString), nil
}

// Removes binary data from the server's bindata directory.
func (c *Client) BindataRemove(hash string) (XmmsValue, error) {
	__payload := <-c.dispatch(10, 34, XmmsList{XmmsString(hash)})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Retrieves a list of binary data hashes from the server's bindata directory.
func (c *Client) BindataList() ([]string, error) {
	__payload := <-c.dispatch(10, 35, XmmsList{})
	__buffer := bytes.NewBuffer(__payload)
	return tryDeserializeStringList(__buffer)
}

// Assemble and send a client-to-client message.
func (c *Client) CourierSendMessage(toClient int, replyPolicy int, payload XmmsDict) (XmmsValue, error) {
	__payload := <-c.dispatch(12, 32, XmmsList{XmmsInt(toClient), XmmsInt(replyPolicy), payload})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Assemble and send a reply to a client-to-client message
func (c *Client) CourierReply(messageId int, replyPolicy int, payload XmmsDict) (XmmsValue, error) {
	__payload := <-c.dispatch(12, 33, XmmsList{XmmsInt(messageId), XmmsInt(replyPolicy), payload})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Return a list of connected clients.
func (c *Client) CourierGetConnectedClients() ([]int, error) {
	__payload := <-c.dispatch(12, 34, XmmsList{})
	__buffer := bytes.NewBuffer(__payload)
	return tryDeserializeIntList(__buffer)
}

// Notify the server that the client's api is ready for query.
func (c *Client) CourierReady() (XmmsValue, error) {
	__payload := <-c.dispatch(12, 35, XmmsList{})
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

// Return a list of clients ready for c2c communication
func (c *Client) CourierGetReadyClients() ([]int, error) {
	__payload := <-c.dispatch(12, 36, XmmsList{})
	__buffer := bytes.NewBuffer(__payload)
	return tryDeserializeIntList(__buffer)
}
