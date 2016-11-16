package xmmsclient

const (
	ActivePlaylist = "_active"
)

const (
	NamespaceAll         = "*"
	NamespaceCollections = "Collections"
	NamespacePlaylists   = "Playlists"
)

const (
	IpcVersion int64 = 24
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
