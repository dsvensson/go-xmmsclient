// auto-generated
package xmmsclient

func (c *Client) MainHello(protocolVersion int, client string) (XmmsInt, error) {
	consumer := newIntConsumer()
	c.dispatch(&consumer, 1, 32, NewXmmsList(XmmsInt(protocolVersion), XmmsString(client)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) MainQuit() (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 1, 33, NewXmmsList())
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) MainListPlugins(pluginType int) (XmmsList, error) {
	consumer := newListConsumer()
	c.dispatch(&consumer, 1, 34, NewXmmsList(XmmsInt(pluginType)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) MainStats() (XmmsDict, error) {
	consumer := newDictConsumer()
	c.dispatch(&consumer, 1, 35, NewXmmsList())
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaylistReplace(name string, replacement XmmsValue, action int) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 2, 32, NewXmmsList(XmmsString(name), replacement, XmmsInt(action)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaylistSetNext(position int) (XmmsInt, error) {
	consumer := newIntConsumer()
	c.dispatch(&consumer, 2, 33, NewXmmsList(XmmsInt(position)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaylistSetNextRel(positionDelta int) (XmmsInt, error) {
	consumer := newIntConsumer()
	c.dispatch(&consumer, 2, 34, NewXmmsList(XmmsInt(positionDelta)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaylistAddUrl(name string, url string) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 2, 35, NewXmmsList(XmmsString(name), XmmsString(url)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaylistAddCollection(name string, collection XmmsValue) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 2, 36, NewXmmsList(XmmsString(name), collection))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaylistRemoveEntry(name string, position int) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 2, 37, NewXmmsList(XmmsString(name), XmmsInt(position)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaylistMoveEntry(name string, position int, newPosition int) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 2, 38, NewXmmsList(XmmsString(name), XmmsInt(position), XmmsInt(newPosition)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaylistListEntries(name string) (XmmsList, error) {
	consumer := newListConsumer()
	c.dispatch(&consumer, 2, 39, NewXmmsList(XmmsString(name)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaylistCurrentPos(name string) (XmmsDict, error) {
	consumer := newDictConsumer()
	c.dispatch(&consumer, 2, 40, NewXmmsList(XmmsString(name)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaylistCurrentActive() (XmmsString, error) {
	consumer := newStringConsumer()
	c.dispatch(&consumer, 2, 41, NewXmmsList())
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaylistInsertUrl(name string, position int, url string) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 2, 42, NewXmmsList(XmmsString(name), XmmsInt(position), XmmsString(url)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaylistInsertCollection(name string, position int, collection XmmsValue) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 2, 43, NewXmmsList(XmmsString(name), XmmsInt(position), collection))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaylistLoad(name string) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 2, 44, NewXmmsList(XmmsString(name)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaylistRadd(name string, url string) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 2, 45, NewXmmsList(XmmsString(name), XmmsString(url)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaylistRinsert(name string, position int, url string) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 2, 46, NewXmmsList(XmmsString(name), XmmsInt(position), XmmsString(url)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) ConfigGetValue(key string) (XmmsString, error) {
	consumer := newStringConsumer()
	c.dispatch(&consumer, 3, 32, NewXmmsList(XmmsString(key)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) ConfigSetValue(key string, value string) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 3, 33, NewXmmsList(XmmsString(key), XmmsString(value)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) ConfigRegisterValue(key string, value string) (XmmsString, error) {
	consumer := newStringConsumer()
	c.dispatch(&consumer, 3, 34, NewXmmsList(XmmsString(key), XmmsString(value)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) ConfigListValues() (XmmsDict, error) {
	consumer := newDictConsumer()
	c.dispatch(&consumer, 3, 35, NewXmmsList())
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaybackStart() (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 4, 32, NewXmmsList())
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaybackStop() (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 4, 33, NewXmmsList())
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaybackPause() (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 4, 34, NewXmmsList())
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaybackTickle() (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 4, 35, NewXmmsList())
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaybackPlaytime() (XmmsInt, error) {
	consumer := newIntConsumer()
	c.dispatch(&consumer, 4, 36, NewXmmsList())
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaybackSeekMs(offset int, whence int) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 4, 37, NewXmmsList(XmmsInt(offset), XmmsInt(whence)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaybackSeekSamples(offset int, whence int) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 4, 38, NewXmmsList(XmmsInt(offset), XmmsInt(whence)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaybackStatus() (XmmsInt, error) {
	consumer := newIntConsumer()
	c.dispatch(&consumer, 4, 39, NewXmmsList())
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaybackCurrentId() (XmmsInt, error) {
	consumer := newIntConsumer()
	c.dispatch(&consumer, 4, 40, NewXmmsList())
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaybackVolumeSet(channel string, volume int) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 4, 41, NewXmmsList(XmmsString(channel), XmmsInt(volume)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) PlaybackVolumeGet() (XmmsDict, error) {
	consumer := newDictConsumer()
	c.dispatch(&consumer, 4, 42, NewXmmsList())
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) MedialibGetInfo(id int) (XmmsDict, error) {
	consumer := newDictConsumer()
	c.dispatch(&consumer, 5, 32, NewXmmsList(XmmsInt(id)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) MedialibImportPath(directory string) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 5, 33, NewXmmsList(XmmsString(directory)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) MedialibRehash(id int) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 5, 34, NewXmmsList(XmmsInt(id)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) MedialibGetId(url string) (XmmsInt, error) {
	consumer := newIntConsumer()
	c.dispatch(&consumer, 5, 35, NewXmmsList(XmmsString(url)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) MedialibRemoveEntry(id int) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 5, 36, NewXmmsList(XmmsInt(id)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) MedialibSetPropertyString(id int, source string, key string, value string) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 5, 37, NewXmmsList(XmmsInt(id), XmmsString(source), XmmsString(key), XmmsString(value)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) MedialibSetPropertyInt(id int, source string, key string, value int) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 5, 38, NewXmmsList(XmmsInt(id), XmmsString(source), XmmsString(key), XmmsInt(value)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) MedialibRemoveProperty(id int, source string, key string) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 5, 39, NewXmmsList(XmmsInt(id), XmmsString(source), XmmsString(key)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) MedialibMoveEntry(id int, url string) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 5, 40, NewXmmsList(XmmsInt(id), XmmsString(url)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) MedialibAddEntry(url string) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 5, 41, NewXmmsList(XmmsString(url)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) CollectionGet(name string, namespace string) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 6, 32, NewXmmsList(XmmsString(name), XmmsString(namespace)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) CollectionList(namespace string) (XmmsList, error) {
	consumer := newListConsumer()
	c.dispatch(&consumer, 6, 33, NewXmmsList(XmmsString(namespace)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) CollectionSave(name string, namespace string, collection XmmsValue) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 6, 34, NewXmmsList(XmmsString(name), XmmsString(namespace), collection))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) CollectionRemove(name string, namespace string) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 6, 35, NewXmmsList(XmmsString(name), XmmsString(namespace)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) CollectionFind(id int, namespace string) (XmmsList, error) {
	consumer := newListConsumer()
	c.dispatch(&consumer, 6, 36, NewXmmsList(XmmsInt(id), XmmsString(namespace)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) CollectionRename(name string, newName string, namespace string) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 6, 37, NewXmmsList(XmmsString(name), XmmsString(newName), XmmsString(namespace)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) CollectionQuery(collection XmmsValue, fetch XmmsDict) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 6, 38, NewXmmsList(collection, fetch))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) CollectionQueryInfos(collection XmmsValue, limitStart int, limitLength int, properties XmmsList, groupBy XmmsList) (XmmsList, error) {
	consumer := newListConsumer()
	c.dispatch(&consumer, 6, 39, NewXmmsList(collection, XmmsInt(limitStart), XmmsInt(limitLength), properties, groupBy))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) CollectionIdlistFromPlaylist(url string) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 6, 40, NewXmmsList(XmmsString(url)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) VisualizationQueryVersion() (XmmsInt, error) {
	consumer := newIntConsumer()
	c.dispatch(&consumer, 7, 32, NewXmmsList())
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) VisualizationRegister() (XmmsInt, error) {
	consumer := newIntConsumer()
	c.dispatch(&consumer, 7, 33, NewXmmsList())
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) VisualizationInitShm(id int, shmId string) (XmmsInt, error) {
	consumer := newIntConsumer()
	c.dispatch(&consumer, 7, 34, NewXmmsList(XmmsInt(id), XmmsString(shmId)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) VisualizationInitUdp(id int) (XmmsInt, error) {
	consumer := newIntConsumer()
	c.dispatch(&consumer, 7, 35, NewXmmsList(XmmsInt(id)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) VisualizationSetProperty(id int, key string, value string) (XmmsInt, error) {
	consumer := newIntConsumer()
	c.dispatch(&consumer, 7, 36, NewXmmsList(XmmsInt(id), XmmsString(key), XmmsString(value)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) VisualizationSetProperties(id int, properties XmmsDict) (XmmsInt, error) {
	consumer := newIntConsumer()
	c.dispatch(&consumer, 7, 37, NewXmmsList(XmmsInt(id), properties))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) VisualizationShutdown(id int) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 7, 38, NewXmmsList(XmmsInt(id)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) XformBrowse(url string) (XmmsList, error) {
	consumer := newListConsumer()
	c.dispatch(&consumer, 9, 32, NewXmmsList(XmmsString(url)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) BindataRetrieve(hash string) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 10, 32, NewXmmsList(XmmsString(hash)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) BindataAdd(rawData XmmsValue) (XmmsString, error) {
	consumer := newStringConsumer()
	c.dispatch(&consumer, 10, 33, NewXmmsList(rawData))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) BindataRemove(hash string) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 10, 34, NewXmmsList(XmmsString(hash)))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) BindataList() (XmmsList, error) {
	consumer := newListConsumer()
	c.dispatch(&consumer, 10, 35, NewXmmsList())
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) CollSyncSync() (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 11, 32, NewXmmsList())
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) CourierSendMessage(toClient int, replyPolicy int, payload XmmsDict) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 12, 32, NewXmmsList(XmmsInt(toClient), XmmsInt(replyPolicy), payload))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) CourierReply(messageId int, replyPolicy int, payload XmmsDict) (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 12, 33, NewXmmsList(XmmsInt(messageId), XmmsInt(replyPolicy), payload))
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) CourierGetConnectedClients() (XmmsList, error) {
	consumer := newListConsumer()
	c.dispatch(&consumer, 12, 34, NewXmmsList())
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) CourierReady() (XmmsValue, error) {
	consumer := newGenericConsumer()
	c.dispatch(&consumer, 12, 35, NewXmmsList())
	result := <-consumer.result
	return result.value, result.err
}

func (c *Client) CourierGetReadyClients() (XmmsList, error) {
	consumer := newListConsumer()
	c.dispatch(&consumer, 12, 36, NewXmmsList())
	result := <-consumer.result
	return result.value, result.err
}
