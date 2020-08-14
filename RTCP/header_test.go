package rtcp

import (
	"bytes"
	"testing"
)

func TestHeaderEncodeDecode(t *testing.T) {
	header := &Header{
		Padding:              true,
		ReceptionReportCount: 12,
		PacketType:           TypeSendReport,
		Length:               0xfff,
	}
	buf := new(bytes.Buffer)
	err := header.Encode(buf)
	if err != nil {
		t.Error(err)
	}

	decHeader := &Header{}
	err = decHeader.Decode(buf.Bytes())
	if err != nil {
		t.Error(err)
	}
	if header.Padding != decHeader.Padding ||
		header.ReceptionReportCount != decHeader.ReceptionReportCount ||
		header.PacketType != decHeader.PacketType ||
		header.Length != decHeader.Length {
		t.Fail()
	}
}
