// auto-generated
package xmmsclient

import (
	"bytes"
)


func (c *Client) MainHello(protocolVersion int, client string) (XmmsInt, error) {
	__payload := <-c.dispatch(1, 32, NewXmmsList(XmmsInt(protocolVersion), XmmsString(client)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return 0, __err
	}
	return __value.(XmmsInt), nil
}

func (c *Client) MainQuit() (XmmsValue, error) {
	__payload := <-c.dispatch(1, 33, NewXmmsList())
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) MainListPlugins(pluginType int) (XmmsList, error) {
	__payload := <-c.dispatch(1, 34, NewXmmsList(XmmsInt(pluginType)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return XmmsList{}, __err
	}
	return __value.(XmmsList), nil
}

func (c *Client) MainStats() (XmmsDict, error) {
	__payload := <-c.dispatch(1, 35, NewXmmsList())
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return XmmsDict{}, __err
	}
	return __value.(XmmsDict), nil
}

func (c *Client) PlaylistReplace(name string, replacement XmmsValue, action int) (XmmsValue, error) {
	__payload := <-c.dispatch(2, 32, NewXmmsList(XmmsString(name), replacement, XmmsInt(action)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) PlaylistSetNext(position int) (XmmsInt, error) {
	__payload := <-c.dispatch(2, 33, NewXmmsList(XmmsInt(position)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return 0, __err
	}
	return __value.(XmmsInt), nil
}

func (c *Client) PlaylistSetNextRel(positionDelta int) (XmmsInt, error) {
	__payload := <-c.dispatch(2, 34, NewXmmsList(XmmsInt(positionDelta)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return 0, __err
	}
	return __value.(XmmsInt), nil
}

func (c *Client) PlaylistAddUrl(name string, url string) (XmmsValue, error) {
	__payload := <-c.dispatch(2, 35, NewXmmsList(XmmsString(name), XmmsString(url)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) PlaylistAddCollection(name string, collection XmmsValue) (XmmsValue, error) {
	__payload := <-c.dispatch(2, 36, NewXmmsList(XmmsString(name), collection))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) PlaylistRemoveEntry(name string, position int) (XmmsValue, error) {
	__payload := <-c.dispatch(2, 37, NewXmmsList(XmmsString(name), XmmsInt(position)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) PlaylistMoveEntry(name string, position int, newPosition int) (XmmsValue, error) {
	__payload := <-c.dispatch(2, 38, NewXmmsList(XmmsString(name), XmmsInt(position), XmmsInt(newPosition)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) PlaylistListEntries(name string) (XmmsList, error) {
	__payload := <-c.dispatch(2, 39, NewXmmsList(XmmsString(name)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return XmmsList{}, __err
	}
	return __value.(XmmsList), nil
}

func (c *Client) PlaylistCurrentPos(name string) (XmmsDict, error) {
	__payload := <-c.dispatch(2, 40, NewXmmsList(XmmsString(name)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return XmmsDict{}, __err
	}
	return __value.(XmmsDict), nil
}

func (c *Client) PlaylistCurrentActive() (XmmsString, error) {
	__payload := <-c.dispatch(2, 41, NewXmmsList())
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return "", __err
	}
	return __value.(XmmsString), nil
}

func (c *Client) PlaylistInsertUrl(name string, position int, url string) (XmmsValue, error) {
	__payload := <-c.dispatch(2, 42, NewXmmsList(XmmsString(name), XmmsInt(position), XmmsString(url)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) PlaylistInsertCollection(name string, position int, collection XmmsValue) (XmmsValue, error) {
	__payload := <-c.dispatch(2, 43, NewXmmsList(XmmsString(name), XmmsInt(position), collection))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) PlaylistLoad(name string) (XmmsValue, error) {
	__payload := <-c.dispatch(2, 44, NewXmmsList(XmmsString(name)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) PlaylistRadd(name string, url string) (XmmsValue, error) {
	__payload := <-c.dispatch(2, 45, NewXmmsList(XmmsString(name), XmmsString(url)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) PlaylistRinsert(name string, position int, url string) (XmmsValue, error) {
	__payload := <-c.dispatch(2, 46, NewXmmsList(XmmsString(name), XmmsInt(position), XmmsString(url)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) ConfigGetValue(key string) (XmmsString, error) {
	__payload := <-c.dispatch(3, 32, NewXmmsList(XmmsString(key)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return "", __err
	}
	return __value.(XmmsString), nil
}

func (c *Client) ConfigSetValue(key string, value string) (XmmsValue, error) {
	__payload := <-c.dispatch(3, 33, NewXmmsList(XmmsString(key), XmmsString(value)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) ConfigRegisterValue(key string, value string) (XmmsString, error) {
	__payload := <-c.dispatch(3, 34, NewXmmsList(XmmsString(key), XmmsString(value)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return "", __err
	}
	return __value.(XmmsString), nil
}

func (c *Client) ConfigListValues() (XmmsDict, error) {
	__payload := <-c.dispatch(3, 35, NewXmmsList())
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return XmmsDict{}, __err
	}
	return __value.(XmmsDict), nil
}

func (c *Client) PlaybackStart() (XmmsValue, error) {
	__payload := <-c.dispatch(4, 32, NewXmmsList())
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) PlaybackStop() (XmmsValue, error) {
	__payload := <-c.dispatch(4, 33, NewXmmsList())
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) PlaybackPause() (XmmsValue, error) {
	__payload := <-c.dispatch(4, 34, NewXmmsList())
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) PlaybackTickle() (XmmsValue, error) {
	__payload := <-c.dispatch(4, 35, NewXmmsList())
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) PlaybackPlaytime() (XmmsInt, error) {
	__payload := <-c.dispatch(4, 36, NewXmmsList())
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return 0, __err
	}
	return __value.(XmmsInt), nil
}

func (c *Client) PlaybackSeekMs(offset int, whence int) (XmmsValue, error) {
	__payload := <-c.dispatch(4, 37, NewXmmsList(XmmsInt(offset), XmmsInt(whence)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) PlaybackSeekSamples(offset int, whence int) (XmmsValue, error) {
	__payload := <-c.dispatch(4, 38, NewXmmsList(XmmsInt(offset), XmmsInt(whence)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) PlaybackStatus() (XmmsInt, error) {
	__payload := <-c.dispatch(4, 39, NewXmmsList())
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return 0, __err
	}
	return __value.(XmmsInt), nil
}

func (c *Client) PlaybackCurrentId() (XmmsInt, error) {
	__payload := <-c.dispatch(4, 40, NewXmmsList())
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return 0, __err
	}
	return __value.(XmmsInt), nil
}

func (c *Client) PlaybackVolumeSet(channel string, volume int) (XmmsValue, error) {
	__payload := <-c.dispatch(4, 41, NewXmmsList(XmmsString(channel), XmmsInt(volume)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) PlaybackVolumeGet() (XmmsDict, error) {
	__payload := <-c.dispatch(4, 42, NewXmmsList())
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return XmmsDict{}, __err
	}
	return __value.(XmmsDict), nil
}

func (c *Client) MedialibGetInfo(id int) (XmmsDict, error) {
	__payload := <-c.dispatch(5, 32, NewXmmsList(XmmsInt(id)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return XmmsDict{}, __err
	}
	return __value.(XmmsDict), nil
}

func (c *Client) MedialibImportPath(directory string) (XmmsValue, error) {
	__payload := <-c.dispatch(5, 33, NewXmmsList(XmmsString(directory)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) MedialibRehash(id int) (XmmsValue, error) {
	__payload := <-c.dispatch(5, 34, NewXmmsList(XmmsInt(id)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) MedialibGetId(url string) (XmmsInt, error) {
	__payload := <-c.dispatch(5, 35, NewXmmsList(XmmsString(url)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return 0, __err
	}
	return __value.(XmmsInt), nil
}

func (c *Client) MedialibRemoveEntry(id int) (XmmsValue, error) {
	__payload := <-c.dispatch(5, 36, NewXmmsList(XmmsInt(id)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) MedialibSetPropertyString(id int, source string, key string, value string) (XmmsValue, error) {
	__payload := <-c.dispatch(5, 37, NewXmmsList(XmmsInt(id), XmmsString(source), XmmsString(key), XmmsString(value)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) MedialibSetPropertyInt(id int, source string, key string, value int) (XmmsValue, error) {
	__payload := <-c.dispatch(5, 38, NewXmmsList(XmmsInt(id), XmmsString(source), XmmsString(key), XmmsInt(value)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) MedialibRemoveProperty(id int, source string, key string) (XmmsValue, error) {
	__payload := <-c.dispatch(5, 39, NewXmmsList(XmmsInt(id), XmmsString(source), XmmsString(key)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) MedialibMoveEntry(id int, url string) (XmmsValue, error) {
	__payload := <-c.dispatch(5, 40, NewXmmsList(XmmsInt(id), XmmsString(url)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) MedialibAddEntry(url string) (XmmsValue, error) {
	__payload := <-c.dispatch(5, 41, NewXmmsList(XmmsString(url)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) CollectionGet(name string, namespace string) (XmmsValue, error) {
	__payload := <-c.dispatch(6, 32, NewXmmsList(XmmsString(name), XmmsString(namespace)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) CollectionList(namespace string) (XmmsList, error) {
	__payload := <-c.dispatch(6, 33, NewXmmsList(XmmsString(namespace)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return XmmsList{}, __err
	}
	return __value.(XmmsList), nil
}

func (c *Client) CollectionSave(name string, namespace string, collection XmmsValue) (XmmsValue, error) {
	__payload := <-c.dispatch(6, 34, NewXmmsList(XmmsString(name), XmmsString(namespace), collection))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) CollectionRemove(name string, namespace string) (XmmsValue, error) {
	__payload := <-c.dispatch(6, 35, NewXmmsList(XmmsString(name), XmmsString(namespace)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) CollectionFind(id int, namespace string) (XmmsList, error) {
	__payload := <-c.dispatch(6, 36, NewXmmsList(XmmsInt(id), XmmsString(namespace)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return XmmsList{}, __err
	}
	return __value.(XmmsList), nil
}

func (c *Client) CollectionRename(name string, newName string, namespace string) (XmmsValue, error) {
	__payload := <-c.dispatch(6, 37, NewXmmsList(XmmsString(name), XmmsString(newName), XmmsString(namespace)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) CollectionQuery(collection XmmsValue, fetch XmmsDict) (XmmsValue, error) {
	__payload := <-c.dispatch(6, 38, NewXmmsList(collection, fetch))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) CollectionQueryInfos(collection XmmsValue, limitStart int, limitLength int, properties XmmsList, groupBy XmmsList) (XmmsList, error) {
	__payload := <-c.dispatch(6, 39, NewXmmsList(collection, XmmsInt(limitStart), XmmsInt(limitLength), properties, groupBy))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return XmmsList{}, __err
	}
	return __value.(XmmsList), nil
}

func (c *Client) CollectionIdlistFromPlaylist(url string) (XmmsValue, error) {
	__payload := <-c.dispatch(6, 40, NewXmmsList(XmmsString(url)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) VisualizationQueryVersion() (XmmsInt, error) {
	__payload := <-c.dispatch(7, 32, NewXmmsList())
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return 0, __err
	}
	return __value.(XmmsInt), nil
}

func (c *Client) VisualizationRegister() (XmmsInt, error) {
	__payload := <-c.dispatch(7, 33, NewXmmsList())
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return 0, __err
	}
	return __value.(XmmsInt), nil
}

func (c *Client) VisualizationInitShm(id int, shmId string) (XmmsInt, error) {
	__payload := <-c.dispatch(7, 34, NewXmmsList(XmmsInt(id), XmmsString(shmId)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return 0, __err
	}
	return __value.(XmmsInt), nil
}

func (c *Client) VisualizationInitUdp(id int) (XmmsInt, error) {
	__payload := <-c.dispatch(7, 35, NewXmmsList(XmmsInt(id)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return 0, __err
	}
	return __value.(XmmsInt), nil
}

func (c *Client) VisualizationSetProperty(id int, key string, value string) (XmmsInt, error) {
	__payload := <-c.dispatch(7, 36, NewXmmsList(XmmsInt(id), XmmsString(key), XmmsString(value)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return 0, __err
	}
	return __value.(XmmsInt), nil
}

func (c *Client) VisualizationSetProperties(id int, properties XmmsDict) (XmmsInt, error) {
	__payload := <-c.dispatch(7, 37, NewXmmsList(XmmsInt(id), properties))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return 0, __err
	}
	return __value.(XmmsInt), nil
}

func (c *Client) VisualizationShutdown(id int) (XmmsValue, error) {
	__payload := <-c.dispatch(7, 38, NewXmmsList(XmmsInt(id)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) XformBrowse(url string) (XmmsList, error) {
	__payload := <-c.dispatch(9, 32, NewXmmsList(XmmsString(url)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return XmmsList{}, __err
	}
	return __value.(XmmsList), nil
}

func (c *Client) BindataRetrieve(hash string) (XmmsValue, error) {
	__payload := <-c.dispatch(10, 32, NewXmmsList(XmmsString(hash)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) BindataAdd(rawData XmmsValue) (XmmsString, error) {
	__payload := <-c.dispatch(10, 33, NewXmmsList(rawData))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return "", __err
	}
	return __value.(XmmsString), nil
}

func (c *Client) BindataRemove(hash string) (XmmsValue, error) {
	__payload := <-c.dispatch(10, 34, NewXmmsList(XmmsString(hash)))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) BindataList() (XmmsList, error) {
	__payload := <-c.dispatch(10, 35, NewXmmsList())
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return XmmsList{}, __err
	}
	return __value.(XmmsList), nil
}

func (c *Client) CollSyncSync() (XmmsValue, error) {
	__payload := <-c.dispatch(11, 32, NewXmmsList())
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) CourierSendMessage(toClient int, replyPolicy int, payload XmmsDict) (XmmsValue, error) {
	__payload := <-c.dispatch(12, 32, NewXmmsList(XmmsInt(toClient), XmmsInt(replyPolicy), payload))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) CourierReply(messageId int, replyPolicy int, payload XmmsDict) (XmmsValue, error) {
	__payload := <-c.dispatch(12, 33, NewXmmsList(XmmsInt(messageId), XmmsInt(replyPolicy), payload))
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) CourierGetConnectedClients() (XmmsList, error) {
	__payload := <-c.dispatch(12, 34, NewXmmsList())
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return XmmsList{}, __err
	}
	return __value.(XmmsList), nil
}

func (c *Client) CourierReady() (XmmsValue, error) {
	__payload := <-c.dispatch(12, 35, NewXmmsList())
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return nil, __err
	}
	return __value.(XmmsValue), nil
}

func (c *Client) CourierGetReadyClients() (XmmsList, error) {
	__payload := <-c.dispatch(12, 36, NewXmmsList())
	__buffer := bytes.NewBuffer(__payload)
	__value, __err := tryDeserialize(__buffer, DeserializeXmmsValue)
	if __err != nil {
		return XmmsList{}, __err
	}
	return __value.(XmmsList), nil
}
