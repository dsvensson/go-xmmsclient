package xmmsclient

const (
	IpcVersion int64 = 24
)

const (
	ObjectSignal   uint32 = 0
	ObjectMain     uint32 = 1
	ObjectPlaylist uint32 = 2
	ObjectConfig   uint32 = 3
	ObjectPlayback uint32 = 4
	ObjectMedialib uint32 = 5
	/* ... */
)
const (
	CommandMainHello       uint32 = 32
	CommandMainQuit        uint32 = 33
	CommandMainListPlugins uint32 = 34
	CommandMainStats       uint32 = 35
	/* ... */
)

const (
	TypeNone      uint32 = 0
	TypeError     uint32 = 1
	TypeInt64     uint32 = 2
	TypeString    uint32 = 3
	TypeColl      uint32 = 4
	TypeBin       uint32 = 5
	TypeList      uint32 = 6
	TypeDict      uint32 = 7
	TypeBitBuffer uint32 = 8
	TypeFloat     uint32 = 9
)
