// auto-generated
package xmmsclient

const (

	C2cReplyPolicyNoReply = 0
	C2cReplyPolicySingleReply = 1
	C2cReplyPolicyMultiReply = 2

	CollectionChangedAdd = 0
	CollectionChangedUpdate = 1
	CollectionChangedRename = 2
	CollectionChangedRemove = 3

	CollectionTypeReference = 0
	CollectionTypeUniverse = 1
	CollectionTypeUnion = 2
	CollectionTypeIntersection = 3
	CollectionTypeComplement = 4
	CollectionTypeHas = 5
	CollectionTypeMatch = 6
	CollectionTypeToken = 7
	CollectionTypeEquals = 8
	CollectionTypeNotequal = 9
	CollectionTypeSmaller = 10
	CollectionTypeSmallereq = 11
	CollectionTypeGreater = 12
	CollectionTypeGreatereq = 13
	CollectionTypeOrder = 14
	CollectionTypeLimit = 15
	CollectionTypeMediaset = 16
	CollectionTypeIdlist = 17
	CollectionTypeLast = 18

	IpcCommandSignal = 0
	IpcCommandBroadcast = 1

	IpcCommandReply = 0
	IpcCommandError = 1

	LogLevelUnknown = 0
	LogLevelFatal = 1
	LogLevelFail = 2
	LogLevelError = 3
	LogLevelInfo = 4
	LogLevelDebug = 5
	LogLevelCount = 6

	MediainfoReaderStatusIdle = 0
	MediainfoReaderStatusRunning = 1

	MedialibEntryStatusNew = 0
	MedialibEntryStatusOk = 1
	MedialibEntryStatusResolving = 2
	MedialibEntryStatusNotAvailable = 3
	MedialibEntryStatusRehash = 4

	PlaybackSeekCur = 0
	PlaybackSeekSet = 1

	PlaybackStatusStop = 0
	PlaybackStatusPlay = 1
	PlaybackStatusPause = 2

	PlaylistChangedAdd = 0
	PlaylistChangedInsert = 1
	PlaylistChangedShuffle = 2
	PlaylistChangedRemove = 3
	PlaylistChangedClear = 4
	PlaylistChangedMove = 5
	PlaylistChangedSort = 6
	PlaylistChangedUpdate = 7
	PlaylistChangedReplace = 8

	PlaylistCurrentIdForget = 0
	PlaylistCurrentIdKeep = 1
	PlaylistCurrentIdMoveToFront = 2

	PluginTypeAll = 0
	PluginTypeOutput = 1
	PluginTypeXform = 2
)