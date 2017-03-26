package flv

import "testing"

func TestVideoFormat(t *testing.T) {
	for _, it := range []struct {
		b      uint8
		format VideoFormat
	}{
		{0x15, VideoFormat{Type: "video/vp6-alpha"}},
		{0x14, VideoFormat{Type: "video/vp6"}},
		{0x17, VideoFormat{Type: "video/h264"}},
		{0x12, VideoFormat{Type: "video/x-flash-h263"}},
		{0x13, VideoFormat{Type: "video/x-flash-screen"}},
	} {
		format, err := ParseVideoFormat([]byte{it.b})
		if err != nil {
			t.Errorf("%v: %#v", err, it)
		}
		if f := it.format; format.Type != f.Type {
			t.Errorf("got: %#v, expected: %#v", format, f)
		}
	}
}
