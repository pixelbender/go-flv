package flv

import (
	"errors"
	"io"
	"time"
)

var errUnsupportedVideo = errors.New("flv: unsupported video")

var videoTypes = map[uint8]string{
	1: "video/jpeg",
	2: "video/x-flash-h263",
	3: "video/x-flash-screen",
	4: "video/vp6",
	5: "video/vp6-alpha",
	6: "video/x-flash-screen2",
	7: "video/h264",
}

type VideoFormat struct {
	typ  uint8
	Type string `json:"type,omitempty"`
}

func (r *VideoFormat) Equal(b []byte) bool {
	if len(b) < 1 {
		return false
	}
	return b[0]&0xf == r.typ
}

func ParseVideoFormat(b []byte) (*VideoFormat, error) {
	if len(b) < 1 {
		return nil, io.EOF
	}
	t := b[0] & 0xf
	m, ok := videoTypes[t]
	if !ok {
		return nil, errUnsupportedVideo
	}
	c := &VideoFormat{
		typ:  t,
		Type: m,
	}
	return c, nil
}

type VideoFrame struct {
	format  *VideoFormat
	time    time.Duration
	key     bool
	payload []byte
}
