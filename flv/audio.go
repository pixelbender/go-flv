package flv

import "io"

const (
	formatPCMBE        uint8 = 0
	formatADPCM        uint8 = 0
	formatMP3          uint8 = 0
	formatPCMLE        uint8 = 0
	formatNellymoser16 uint8 = 0
	formatNellymoser8  uint8 = 0
	formatNellymoser   uint8 = 0
	formatPCMA         uint8 = 0
	formatPCMU         uint8 = 0
	formatAAC          uint8 = 0
	formatSpeex        uint8 = 0
	formatMP38         uint8 = 0
)

const (
	SampleRate5  uint8 = 0
	SampleRate11 uint8 = 1
	SampleRate22 uint8 = 2
	SampleRate44 uint8 = 3
)

const (
	SampleSize8  uint8 = 0
	SampleSize16 uint8 = 1
)

type Audio struct {
	Format     uint8
	SampleRate uint8
	SampleSize uint8
	Channels   uint8
	Payload    []byte
}

func ParseAudio(b []byte) (*Audio, error) {
	if len(b) < 1 {
		return nil, io.ErrUnexpectedEOF
	}
	f := b[0]
	a := &Audio{
		Format:     f >> 4,
		SampleRate: f >> 2 & 3,
		SampleSize: f >> 1 & 1,
		Channels:   f & 1,
		Payload: b[1:],
	}
	return a, nil
}
