package flv

import (
	"errors"
	"io"
	"time"
)

var errUnsupportedAudio = errors.New("flv: unsupported audio")

var audioTypes = map[uint8]string{
	0:  "audio/pcm",
	1:  "audio/adpcm",
	2:  "audio/mpeg",
	3:  "audio/pcm",
	4:  "audio/nellymoser",
	5:  "audio/nellymoser",
	6:  "audio/nellymoser",
	7:  "audio/pcma",
	8:  "audio/pcmu",
	10: "audio/mpeg",
	11: "audio/speex",
	14: "audio/mpeg",
}

var audioRates = []int{5512, 11025, 22050, 44100}

type AudioFormat struct {
	typ         uint8
	Type        string `json:"type,omitempty"`
	Rate        int    `json:"rate,omitempty"`
	Format      string `json:"format,omitempty"`
	Channels    int    `json:"channels,omitempty"`
	MPEGVersion int    `json:"mpeg-version,omitempty"`
	MPEGLayer   int    `json:"mpeg-layer,omitempty"`
}

func (r *AudioFormat) Equal(b []byte) bool {
	if len(b) < 1 {
		return false
	}
	return b[0] == r.typ
}

func ParseAudioFormat(b []byte) (*AudioFormat, error) {
	if len(b) < 1 {
		return nil, io.EOF
	}
	t := b[0]
	m, ok := audioTypes[t>>4]
	if !ok {
		return nil, errUnsupportedAudio
	}
	c := &AudioFormat{
		typ:      t,
		Type:     m,
		Rate:     audioRates[t>>2&3],
		Channels: int(t&1) + 1,
	}
	s := t >> 1 & 1
	switch t >> 4 {
	case 0, 1:
		c.Format = []string{"u8", "s16"}[s]
	case 2:
		c.MPEGVersion = 1
		c.MPEGLayer = 3
	case 3:
		c.Format = []string{"u8", "s16le"}[s]
	case 4, 11:
		c.Rate = 16000
		c.Channels = 1
	case 5, 7, 8:
		c.Rate = 8000
		c.Channels = 1
	case 10:
		c.MPEGVersion = 4
	case 14:
		c.Rate = 8000
		c.MPEGVersion = 1
		c.MPEGLayer = 3
	}
	return c, nil
}

type AudioFrame struct {
	format  *AudioFrame
	time    time.Duration
	payload []byte
}
