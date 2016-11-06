package xmmsclient

func (c *Client) MainHello(protocolVersion int, client string) XmmsValue {
	return <-c.dispatch(1, 32, NewXmmsList(XmmsInt(protocolVersion), XmmsString(client)))
}

func (c *Client) MainQuit() XmmsValue {
	return <-c.dispatch(1, 33, NewXmmsList())
}

func (c *Client) MainListPlugins(pluginType int) XmmsValue {
	return <-c.dispatch(1, 34, NewXmmsList(XmmsInt(pluginType)))
}

func (c *Client) MainStats() XmmsValue {
	return <-c.dispatch(1, 35, NewXmmsList())
}

func (c *Client) PlaylistReplace(name string, replacement XmmsValue, action int) XmmsValue {
	return <-c.dispatch(2, 32, NewXmmsList(XmmsString(name), replacement, XmmsInt(action)))
}

func (c *Client) PlaylistSetNext(position int) XmmsValue {
	return <-c.dispatch(2, 33, NewXmmsList(XmmsInt(position)))
}

func (c *Client) PlaylistSetNextRel(positionDelta int) XmmsValue {
	return <-c.dispatch(2, 34, NewXmmsList(XmmsInt(positionDelta)))
}

func (c *Client) PlaylistAddUrl(name string, url string) XmmsValue {
	return <-c.dispatch(2, 35, NewXmmsList(XmmsString(name), XmmsString(url)))
}

func (c *Client) PlaylistAddCollection(name string, collection XmmsValue) XmmsValue {
	return <-c.dispatch(2, 36, NewXmmsList(XmmsString(name), collection))
}

func (c *Client) PlaylistRemoveEntry(name string, position int) XmmsValue {
	return <-c.dispatch(2, 37, NewXmmsList(XmmsString(name), XmmsInt(position)))
}

func (c *Client) PlaylistMoveEntry(name string, position int, newPosition int) XmmsValue {
	return <-c.dispatch(2, 38, NewXmmsList(XmmsString(name), XmmsInt(position), XmmsInt(newPosition)))
}

func (c *Client) PlaylistListEntries(name string) XmmsValue {
	return <-c.dispatch(2, 39, NewXmmsList(XmmsString(name)))
}

func (c *Client) PlaylistCurrentPos(name string) XmmsValue {
	return <-c.dispatch(2, 40, NewXmmsList(XmmsString(name)))
}

func (c *Client) PlaylistCurrentActive() XmmsValue {
	return <-c.dispatch(2, 41, NewXmmsList())
}

func (c *Client) PlaylistInsertUrl(name string, position int, url string) XmmsValue {
	return <-c.dispatch(2, 42, NewXmmsList(XmmsString(name), XmmsInt(position), XmmsString(url)))
}

func (c *Client) PlaylistInsertCollection(name string, position int, collection XmmsValue) XmmsValue {
	return <-c.dispatch(2, 43, NewXmmsList(XmmsString(name), XmmsInt(position), collection))
}

func (c *Client) PlaylistLoad(name string) XmmsValue {
	return <-c.dispatch(2, 44, NewXmmsList(XmmsString(name)))
}

func (c *Client) PlaylistRadd(name string, url string) XmmsValue {
	return <-c.dispatch(2, 45, NewXmmsList(XmmsString(name), XmmsString(url)))
}

func (c *Client) PlaylistRinsert(name string, position int, url string) XmmsValue {
	return <-c.dispatch(2, 46, NewXmmsList(XmmsString(name), XmmsInt(position), XmmsString(url)))
}

func (c *Client) ConfigGetValue(key string) XmmsValue {
	return <-c.dispatch(3, 32, NewXmmsList(XmmsString(key)))
}

func (c *Client) ConfigSetValue(key string, value string) XmmsValue {
	return <-c.dispatch(3, 33, NewXmmsList(XmmsString(key), XmmsString(value)))
}

func (c *Client) ConfigRegisterValue(key string, value string) XmmsValue {
	return <-c.dispatch(3, 34, NewXmmsList(XmmsString(key), XmmsString(value)))
}

func (c *Client) ConfigListValues() XmmsValue {
	return <-c.dispatch(3, 35, NewXmmsList())
}

func (c *Client) PlaybackStart() XmmsValue {
	return <-c.dispatch(4, 32, NewXmmsList())
}

func (c *Client) PlaybackStop() XmmsValue {
	return <-c.dispatch(4, 33, NewXmmsList())
}

func (c *Client) PlaybackPause() XmmsValue {
	return <-c.dispatch(4, 34, NewXmmsList())
}

func (c *Client) PlaybackTickle() XmmsValue {
	return <-c.dispatch(4, 35, NewXmmsList())
}

func (c *Client) PlaybackPlaytime() XmmsValue {
	return <-c.dispatch(4, 36, NewXmmsList())
}

func (c *Client) PlaybackSeekMs(offset int, whence int) XmmsValue {
	return <-c.dispatch(4, 37, NewXmmsList(XmmsInt(offset), XmmsInt(whence)))
}

func (c *Client) PlaybackSeekSamples(offset int, whence int) XmmsValue {
	return <-c.dispatch(4, 38, NewXmmsList(XmmsInt(offset), XmmsInt(whence)))
}

func (c *Client) PlaybackStatus() XmmsValue {
	return <-c.dispatch(4, 39, NewXmmsList())
}

func (c *Client) PlaybackCurrentId() XmmsValue {
	return <-c.dispatch(4, 40, NewXmmsList())
}

func (c *Client) PlaybackVolumeSet(channel string, volume int) XmmsValue {
	return <-c.dispatch(4, 41, NewXmmsList(XmmsString(channel), XmmsInt(volume)))
}

func (c *Client) PlaybackVolumeGet() XmmsValue {
	return <-c.dispatch(4, 42, NewXmmsList())
}

func (c *Client) MedialibGetInfo(id int) XmmsValue {
	return <-c.dispatch(5, 32, NewXmmsList(XmmsInt(id)))
}

func (c *Client) MedialibImportPath(directory string) XmmsValue {
	return <-c.dispatch(5, 33, NewXmmsList(XmmsString(directory)))
}

func (c *Client) MedialibRehash(id int) XmmsValue {
	return <-c.dispatch(5, 34, NewXmmsList(XmmsInt(id)))
}

func (c *Client) MedialibGetId(url string) XmmsValue {
	return <-c.dispatch(5, 35, NewXmmsList(XmmsString(url)))
}

func (c *Client) MedialibRemoveEntry(id int) XmmsValue {
	return <-c.dispatch(5, 36, NewXmmsList(XmmsInt(id)))
}

func (c *Client) MedialibSetPropertyString(id int, source string, key string, value string) XmmsValue {
	return <-c.dispatch(5, 37, NewXmmsList(XmmsInt(id), XmmsString(source), XmmsString(key), XmmsString(value)))
}

func (c *Client) MedialibSetPropertyInt(id int, source string, key string, value int) XmmsValue {
	return <-c.dispatch(5, 38, NewXmmsList(XmmsInt(id), XmmsString(source), XmmsString(key), XmmsInt(value)))
}

func (c *Client) MedialibRemoveProperty(id int, source string, key string) XmmsValue {
	return <-c.dispatch(5, 39, NewXmmsList(XmmsInt(id), XmmsString(source), XmmsString(key)))
}

func (c *Client) MedialibMoveEntry(id int, url string) XmmsValue {
	return <-c.dispatch(5, 40, NewXmmsList(XmmsInt(id), XmmsString(url)))
}

func (c *Client) MedialibAddEntry(url string) XmmsValue {
	return <-c.dispatch(5, 41, NewXmmsList(XmmsString(url)))
}

func (c *Client) CollectionGet(name string, namespace string) XmmsValue {
	return <-c.dispatch(6, 32, NewXmmsList(XmmsString(name), XmmsString(namespace)))
}

func (c *Client) CollectionList(namespace string) XmmsValue {
	return <-c.dispatch(6, 33, NewXmmsList(XmmsString(namespace)))
}

func (c *Client) CollectionSave(name string, namespace string, collection XmmsValue) XmmsValue {
	return <-c.dispatch(6, 34, NewXmmsList(XmmsString(name), XmmsString(namespace), collection))
}

func (c *Client) CollectionRemove(name string, namespace string) XmmsValue {
	return <-c.dispatch(6, 35, NewXmmsList(XmmsString(name), XmmsString(namespace)))
}

func (c *Client) CollectionFind(id int, namespace string) XmmsValue {
	return <-c.dispatch(6, 36, NewXmmsList(XmmsInt(id), XmmsString(namespace)))
}

func (c *Client) CollectionRename(name string, newName string, namespace string) XmmsValue {
	return <-c.dispatch(6, 37, NewXmmsList(XmmsString(name), XmmsString(newName), XmmsString(namespace)))
}

func (c *Client) CollectionQuery(collection XmmsValue, fetch XmmsDict) XmmsValue {
	return <-c.dispatch(6, 38, NewXmmsList(collection, fetch))
}

func (c *Client) CollectionQueryInfos(collection XmmsValue, limitStart int, limitLength int, properties XmmsList, groupBy XmmsList) XmmsValue {
	return <-c.dispatch(6, 39, NewXmmsList(collection, XmmsInt(limitStart), XmmsInt(limitLength), properties, groupBy))
}

func (c *Client) CollectionIdlistFromPlaylist(url string) XmmsValue {
	return <-c.dispatch(6, 40, NewXmmsList(XmmsString(url)))
}

func (c *Client) VisualizationQueryVersion() XmmsValue {
	return <-c.dispatch(7, 32, NewXmmsList())
}

func (c *Client) VisualizationRegister() XmmsValue {
	return <-c.dispatch(7, 33, NewXmmsList())
}

func (c *Client) VisualizationInitShm(id int, shmId string) XmmsValue {
	return <-c.dispatch(7, 34, NewXmmsList(XmmsInt(id), XmmsString(shmId)))
}

func (c *Client) VisualizationInitUdp(id int) XmmsValue {
	return <-c.dispatch(7, 35, NewXmmsList(XmmsInt(id)))
}

func (c *Client) VisualizationSetProperty(id int, key string, value string) XmmsValue {
	return <-c.dispatch(7, 36, NewXmmsList(XmmsInt(id), XmmsString(key), XmmsString(value)))
}

func (c *Client) VisualizationSetProperties(id int, properties XmmsDict) XmmsValue {
	return <-c.dispatch(7, 37, NewXmmsList(XmmsInt(id), properties))
}

func (c *Client) VisualizationShutdown(id int) XmmsValue {
	return <-c.dispatch(7, 38, NewXmmsList(XmmsInt(id)))
}

func (c *Client) XformBrowse(url string) XmmsValue {
	return <-c.dispatch(9, 32, NewXmmsList(XmmsString(url)))
}

func (c *Client) BindataRetrieve(hash string) XmmsValue {
	return <-c.dispatch(10, 32, NewXmmsList(XmmsString(hash)))
}

func (c *Client) BindataAdd(rawData XmmsValue) XmmsValue {
	return <-c.dispatch(10, 33, NewXmmsList(rawData))
}

func (c *Client) BindataRemove(hash string) XmmsValue {
	return <-c.dispatch(10, 34, NewXmmsList(XmmsString(hash)))
}

func (c *Client) BindataList() XmmsValue {
	return <-c.dispatch(10, 35, NewXmmsList())
}

func (c *Client) CollSyncSync() XmmsValue {
	return <-c.dispatch(11, 32, NewXmmsList())
}

func (c *Client) CourierSendMessage(toClient int, replyPolicy int, payload XmmsDict) XmmsValue {
	return <-c.dispatch(12, 32, NewXmmsList(XmmsInt(toClient), XmmsInt(replyPolicy), payload))
}

func (c *Client) CourierReply(messageId int, replyPolicy int, payload XmmsDict) XmmsValue {
	return <-c.dispatch(12, 33, NewXmmsList(XmmsInt(messageId), XmmsInt(replyPolicy), payload))
}

func (c *Client) CourierGetConnectedClients() XmmsValue {
	return <-c.dispatch(12, 34, NewXmmsList())
}

func (c *Client) CourierReady() XmmsValue {
	return <-c.dispatch(12, 35, NewXmmsList())
}

func (c *Client) CourierGetReadyClients() XmmsValue {
	return <-c.dispatch(12, 36, NewXmmsList())
}
