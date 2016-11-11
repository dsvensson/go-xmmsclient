package xmmsclient

func (c *Client) MainHello(protocolVersion int, client string) (XmmsInt, error) {
	result := <-c.dispatch(1, 32, NewXmmsList(XmmsInt(protocolVersion), XmmsString(client)))
	if result.err != nil {
		return 0, result.err
	}
	return valueAsInt(result.value)
}

func (c *Client) MainQuit() (XmmsValue, error) {
	result := <-c.dispatch(1, 33, NewXmmsList())
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) MainListPlugins(pluginType int) (XmmsList, error) {
	result := <-c.dispatch(1, 34, NewXmmsList(XmmsInt(pluginType)))
	if result.err != nil {
		return XmmsList{}, result.err
	}
	return valueAsList(result.value)
}

func (c *Client) MainStats() (XmmsDict, error) {
	result := <-c.dispatch(1, 35, NewXmmsList())
	if result.err != nil {
		return XmmsDict{}, result.err
	}
	return valueAsDict(result.value)
}

func (c *Client) PlaylistReplace(name string, replacement XmmsValue, action int) (XmmsValue, error) {
	result := <-c.dispatch(2, 32, NewXmmsList(XmmsString(name), replacement, XmmsInt(action)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) PlaylistSetNext(position int) (XmmsInt, error) {
	result := <-c.dispatch(2, 33, NewXmmsList(XmmsInt(position)))
	if result.err != nil {
		return 0, result.err
	}
	return valueAsInt(result.value)
}

func (c *Client) PlaylistSetNextRel(positionDelta int) (XmmsInt, error) {
	result := <-c.dispatch(2, 34, NewXmmsList(XmmsInt(positionDelta)))
	if result.err != nil {
		return 0, result.err
	}
	return valueAsInt(result.value)
}

func (c *Client) PlaylistAddUrl(name string, url string) (XmmsValue, error) {
	result := <-c.dispatch(2, 35, NewXmmsList(XmmsString(name), XmmsString(url)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) PlaylistAddCollection(name string, collection XmmsValue) (XmmsValue, error) {
	result := <-c.dispatch(2, 36, NewXmmsList(XmmsString(name), collection))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) PlaylistRemoveEntry(name string, position int) (XmmsValue, error) {
	result := <-c.dispatch(2, 37, NewXmmsList(XmmsString(name), XmmsInt(position)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) PlaylistMoveEntry(name string, position int, newPosition int) (XmmsValue, error) {
	result := <-c.dispatch(2, 38, NewXmmsList(XmmsString(name), XmmsInt(position), XmmsInt(newPosition)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) PlaylistListEntries(name string) (XmmsList, error) {
	result := <-c.dispatch(2, 39, NewXmmsList(XmmsString(name)))
	if result.err != nil {
		return XmmsList{}, result.err
	}
	return valueAsList(result.value)
}

func (c *Client) PlaylistCurrentPos(name string) (XmmsDict, error) {
	result := <-c.dispatch(2, 40, NewXmmsList(XmmsString(name)))
	if result.err != nil {
		return XmmsDict{}, result.err
	}
	return valueAsDict(result.value)
}

func (c *Client) PlaylistCurrentActive() (XmmsString, error) {
	result := <-c.dispatch(2, 41, NewXmmsList())
	if result.err != nil {
		return "", result.err
	}
	return valueAsString(result.value)
}

func (c *Client) PlaylistInsertUrl(name string, position int, url string) (XmmsValue, error) {
	result := <-c.dispatch(2, 42, NewXmmsList(XmmsString(name), XmmsInt(position), XmmsString(url)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) PlaylistInsertCollection(name string, position int, collection XmmsValue) (XmmsValue, error) {
	result := <-c.dispatch(2, 43, NewXmmsList(XmmsString(name), XmmsInt(position), collection))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) PlaylistLoad(name string) (XmmsValue, error) {
	result := <-c.dispatch(2, 44, NewXmmsList(XmmsString(name)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) PlaylistRadd(name string, url string) (XmmsValue, error) {
	result := <-c.dispatch(2, 45, NewXmmsList(XmmsString(name), XmmsString(url)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) PlaylistRinsert(name string, position int, url string) (XmmsValue, error) {
	result := <-c.dispatch(2, 46, NewXmmsList(XmmsString(name), XmmsInt(position), XmmsString(url)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) ConfigGetValue(key string) (XmmsString, error) {
	result := <-c.dispatch(3, 32, NewXmmsList(XmmsString(key)))
	if result.err != nil {
		return "", result.err
	}
	return valueAsString(result.value)
}

func (c *Client) ConfigSetValue(key string, value string) (XmmsValue, error) {
	result := <-c.dispatch(3, 33, NewXmmsList(XmmsString(key), XmmsString(value)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) ConfigRegisterValue(key string, value string) (XmmsString, error) {
	result := <-c.dispatch(3, 34, NewXmmsList(XmmsString(key), XmmsString(value)))
	if result.err != nil {
		return "", result.err
	}
	return valueAsString(result.value)
}

func (c *Client) ConfigListValues() (XmmsDict, error) {
	result := <-c.dispatch(3, 35, NewXmmsList())
	if result.err != nil {
		return XmmsDict{}, result.err
	}
	return valueAsDict(result.value)
}

func (c *Client) PlaybackStart() (XmmsValue, error) {
	result := <-c.dispatch(4, 32, NewXmmsList())
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) PlaybackStop() (XmmsValue, error) {
	result := <-c.dispatch(4, 33, NewXmmsList())
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) PlaybackPause() (XmmsValue, error) {
	result := <-c.dispatch(4, 34, NewXmmsList())
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) PlaybackTickle() (XmmsValue, error) {
	result := <-c.dispatch(4, 35, NewXmmsList())
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) PlaybackPlaytime() (XmmsInt, error) {
	result := <-c.dispatch(4, 36, NewXmmsList())
	if result.err != nil {
		return 0, result.err
	}
	return valueAsInt(result.value)
}

func (c *Client) PlaybackSeekMs(offset int, whence int) (XmmsValue, error) {
	result := <-c.dispatch(4, 37, NewXmmsList(XmmsInt(offset), XmmsInt(whence)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) PlaybackSeekSamples(offset int, whence int) (XmmsValue, error) {
	result := <-c.dispatch(4, 38, NewXmmsList(XmmsInt(offset), XmmsInt(whence)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) PlaybackStatus() (XmmsInt, error) {
	result := <-c.dispatch(4, 39, NewXmmsList())
	if result.err != nil {
		return 0, result.err
	}
	return valueAsInt(result.value)
}

func (c *Client) PlaybackCurrentId() (XmmsInt, error) {
	result := <-c.dispatch(4, 40, NewXmmsList())
	if result.err != nil {
		return 0, result.err
	}
	return valueAsInt(result.value)
}

func (c *Client) PlaybackVolumeSet(channel string, volume int) (XmmsValue, error) {
	result := <-c.dispatch(4, 41, NewXmmsList(XmmsString(channel), XmmsInt(volume)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) PlaybackVolumeGet() (XmmsDict, error) {
	result := <-c.dispatch(4, 42, NewXmmsList())
	if result.err != nil {
		return XmmsDict{}, result.err
	}
	return valueAsDict(result.value)
}

func (c *Client) MedialibGetInfo(id int) (XmmsDict, error) {
	result := <-c.dispatch(5, 32, NewXmmsList(XmmsInt(id)))
	if result.err != nil {
		return XmmsDict{}, result.err
	}
	return valueAsDict(result.value)
}

func (c *Client) MedialibImportPath(directory string) (XmmsValue, error) {
	result := <-c.dispatch(5, 33, NewXmmsList(XmmsString(directory)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) MedialibRehash(id int) (XmmsValue, error) {
	result := <-c.dispatch(5, 34, NewXmmsList(XmmsInt(id)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) MedialibGetId(url string) (XmmsInt, error) {
	result := <-c.dispatch(5, 35, NewXmmsList(XmmsString(url)))
	if result.err != nil {
		return 0, result.err
	}
	return valueAsInt(result.value)
}

func (c *Client) MedialibRemoveEntry(id int) (XmmsValue, error) {
	result := <-c.dispatch(5, 36, NewXmmsList(XmmsInt(id)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) MedialibSetPropertyString(id int, source string, key string, value string) (XmmsValue, error) {
	result := <-c.dispatch(5, 37, NewXmmsList(XmmsInt(id), XmmsString(source), XmmsString(key), XmmsString(value)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) MedialibSetPropertyInt(id int, source string, key string, value int) (XmmsValue, error) {
	result := <-c.dispatch(5, 38, NewXmmsList(XmmsInt(id), XmmsString(source), XmmsString(key), XmmsInt(value)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) MedialibRemoveProperty(id int, source string, key string) (XmmsValue, error) {
	result := <-c.dispatch(5, 39, NewXmmsList(XmmsInt(id), XmmsString(source), XmmsString(key)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) MedialibMoveEntry(id int, url string) (XmmsValue, error) {
	result := <-c.dispatch(5, 40, NewXmmsList(XmmsInt(id), XmmsString(url)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) MedialibAddEntry(url string) (XmmsValue, error) {
	result := <-c.dispatch(5, 41, NewXmmsList(XmmsString(url)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) CollectionGet(name string, namespace string) (XmmsValue, error) {
	result := <-c.dispatch(6, 32, NewXmmsList(XmmsString(name), XmmsString(namespace)))
	if result.err != nil {
		return nil, result.err
	}
	return result.value, result.err
}

func (c *Client) CollectionList(namespace string) (XmmsList, error) {
	result := <-c.dispatch(6, 33, NewXmmsList(XmmsString(namespace)))
	if result.err != nil {
		return XmmsList{}, result.err
	}
	return valueAsList(result.value)
}

func (c *Client) CollectionSave(name string, namespace string, collection XmmsValue) (XmmsValue, error) {
	result := <-c.dispatch(6, 34, NewXmmsList(XmmsString(name), XmmsString(namespace), collection))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) CollectionRemove(name string, namespace string) (XmmsValue, error) {
	result := <-c.dispatch(6, 35, NewXmmsList(XmmsString(name), XmmsString(namespace)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) CollectionFind(id int, namespace string) (XmmsList, error) {
	result := <-c.dispatch(6, 36, NewXmmsList(XmmsInt(id), XmmsString(namespace)))
	if result.err != nil {
		return XmmsList{}, result.err
	}
	return valueAsList(result.value)
}

func (c *Client) CollectionRename(name string, newName string, namespace string) (XmmsValue, error) {
	result := <-c.dispatch(6, 37, NewXmmsList(XmmsString(name), XmmsString(newName), XmmsString(namespace)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) CollectionQuery(collection XmmsValue, fetch XmmsDict) (XmmsValue, error) {
	result := <-c.dispatch(6, 38, NewXmmsList(collection, fetch))
	if result.err != nil {
		return nil, result.err
	}
	return result.value, result.err
}

func (c *Client) CollectionQueryInfos(collection XmmsValue, limitStart int, limitLength int, properties XmmsList, groupBy XmmsList) (XmmsList, error) {
	result := <-c.dispatch(6, 39, NewXmmsList(collection, XmmsInt(limitStart), XmmsInt(limitLength), properties, groupBy))
	if result.err != nil {
		return XmmsList{}, result.err
	}
	return valueAsList(result.value)
}

func (c *Client) CollectionIdlistFromPlaylist(url string) (XmmsValue, error) {
	result := <-c.dispatch(6, 40, NewXmmsList(XmmsString(url)))
	if result.err != nil {
		return nil, result.err
	}
	return result.value, result.err
}

func (c *Client) VisualizationQueryVersion() (XmmsInt, error) {
	result := <-c.dispatch(7, 32, NewXmmsList())
	if result.err != nil {
		return 0, result.err
	}
	return valueAsInt(result.value)
}

func (c *Client) VisualizationRegister() (XmmsInt, error) {
	result := <-c.dispatch(7, 33, NewXmmsList())
	if result.err != nil {
		return 0, result.err
	}
	return valueAsInt(result.value)
}

func (c *Client) VisualizationInitShm(id int, shmId string) (XmmsInt, error) {
	result := <-c.dispatch(7, 34, NewXmmsList(XmmsInt(id), XmmsString(shmId)))
	if result.err != nil {
		return 0, result.err
	}
	return valueAsInt(result.value)
}

func (c *Client) VisualizationInitUdp(id int) (XmmsInt, error) {
	result := <-c.dispatch(7, 35, NewXmmsList(XmmsInt(id)))
	if result.err != nil {
		return 0, result.err
	}
	return valueAsInt(result.value)
}

func (c *Client) VisualizationSetProperty(id int, key string, value string) (XmmsInt, error) {
	result := <-c.dispatch(7, 36, NewXmmsList(XmmsInt(id), XmmsString(key), XmmsString(value)))
	if result.err != nil {
		return 0, result.err
	}
	return valueAsInt(result.value)
}

func (c *Client) VisualizationSetProperties(id int, properties XmmsDict) (XmmsInt, error) {
	result := <-c.dispatch(7, 37, NewXmmsList(XmmsInt(id), properties))
	if result.err != nil {
		return 0, result.err
	}
	return valueAsInt(result.value)
}

func (c *Client) VisualizationShutdown(id int) (XmmsValue, error) {
	result := <-c.dispatch(7, 38, NewXmmsList(XmmsInt(id)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) XformBrowse(url string) (XmmsList, error) {
	result := <-c.dispatch(9, 32, NewXmmsList(XmmsString(url)))
	if result.err != nil {
		return XmmsList{}, result.err
	}
	return valueAsList(result.value)
}

func (c *Client) BindataRetrieve(hash string) (XmmsValue, error) {
	result := <-c.dispatch(10, 32, NewXmmsList(XmmsString(hash)))
	if result.err != nil {
		return nil, result.err
	}
	return result.value, result.err
}

func (c *Client) BindataAdd(rawData XmmsValue) (XmmsString, error) {
	result := <-c.dispatch(10, 33, NewXmmsList(rawData))
	if result.err != nil {
		return "", result.err
	}
	return valueAsString(result.value)
}

func (c *Client) BindataRemove(hash string) (XmmsValue, error) {
	result := <-c.dispatch(10, 34, NewXmmsList(XmmsString(hash)))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) BindataList() (XmmsList, error) {
	result := <-c.dispatch(10, 35, NewXmmsList())
	if result.err != nil {
		return XmmsList{}, result.err
	}
	return valueAsList(result.value)
}

func (c *Client) CollSyncSync() (XmmsValue, error) {
	result := <-c.dispatch(11, 32, NewXmmsList())
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) CourierSendMessage(toClient int, replyPolicy int, payload XmmsDict) (XmmsValue, error) {
	result := <-c.dispatch(12, 32, NewXmmsList(XmmsInt(toClient), XmmsInt(replyPolicy), payload))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) CourierReply(messageId int, replyPolicy int, payload XmmsDict) (XmmsValue, error) {
	result := <-c.dispatch(12, 33, NewXmmsList(XmmsInt(messageId), XmmsInt(replyPolicy), payload))
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) CourierGetConnectedClients() (XmmsList, error) {
	result := <-c.dispatch(12, 34, NewXmmsList())
	if result.err != nil {
		return XmmsList{}, result.err
	}
	return valueAsList(result.value)
}

func (c *Client) CourierReady() (XmmsValue, error) {
	result := <-c.dispatch(12, 35, NewXmmsList())
	if result.err != nil {
		return result.value, result.err
	}
	return result.value, result.err
}

func (c *Client) CourierGetReadyClients() (XmmsList, error) {
	result := <-c.dispatch(12, 36, NewXmmsList())
	if result.err != nil {
		return XmmsList{}, result.err
	}
	return valueAsList(result.value)
}
