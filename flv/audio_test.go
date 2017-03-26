package flv

import "testing"

func TestAudioFormat(t *testing.T) {
	for _, it := range []struct {
		b      uint8
		format AudioFormat
	}{
		{0x2b, AudioFormat{Type: "audio/mp3", Rate: 22050, Channels: 2}},
		{0x2a, AudioFormat{Type: "audio/mp3", Rate: 22050, Channels: 1}},
		{0x2e, AudioFormat{Type: "audio/mp3", Rate: 44100, Channels: 1}},
		{0x2f, AudioFormat{Type: "audio/mp3", Rate: 44100, Channels: 2}},
		{0x16, AudioFormat{Type: "audio/adpcm", Rate: 11025, Format: "s16", Channels: 1}},
		{0x1a, AudioFormat{Type: "audio/adpcm", Rate: 22050, Format: "s16", Channels: 1}},
		{0x1e, AudioFormat{Type: "audio/adpcm", Rate: 44100, Format: "s16", Channels: 1}},
		{0x17, AudioFormat{Type: "audio/adpcm", Rate: 11025, Format: "s16", Channels: 2}},
		{0x1b, AudioFormat{Type: "audio/adpcm", Rate: 22050, Format: "s16", Channels: 2}},
		{0x1f, AudioFormat{Type: "audio/adpcm", Rate: 44100, Format: "s16", Channels: 2}},
		{0xaf, AudioFormat{Type: "audio/aac", Channels: 0}},
		{0xaf, AudioFormat{Type: "audio/aac", Channels: 0}},
	} {
		format, err := ParseAudioFormat([]byte{it.b})
		if err != nil {
			t.Errorf("%v: %#v", err, it)
		}
		if f := it.format; format.Type != f.Type || format.Rate != f.Rate || format.Channels != f.Channels || format.Format != f.Format {
			t.Errorf("got: %#v, expected: %#v", format, f)
		}
	}
}
