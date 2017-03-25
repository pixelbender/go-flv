package flv

type Header struct {
	Flags     uint8
}

func NewHeader(flags uint8) *Header {
	return &Header{ flags}
}

type Tag struct {
	Type   uint8
	Size   int
	Time   int64
	Stream uint32
}

const (
	TypeAudio uint8 = 8
	TypeVideo uint8 = 9
	TypeData  uint8 = 18
)

const signature uint32 = 0x464C56
