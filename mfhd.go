package mp4streamparse

import (
	"fmt"
	"strings"
)

// MfhdBox defines the mfhd box structure.
type MfhdBox struct {
	Box
	version        byte
	flags          uint32
	sequenceNumber uint32
}

func NewMfhdBox(name string, size uint32, buf []byte) *MfhdBox {
	return &MfhdBox{Box: NewBox(name, int64(size), buf)}
}

func (b MfhdBox) String() string {
	indentLevel := 2
	tab := strings.Repeat("    ", indentLevel)
	strMsg := fmt.Sprintf("%s<< MfhdBox >>", tab)
	strMsg = fmt.Sprintf("%s\tSequenceNumber(%d)", strMsg, b.sequenceNumber)
	return strMsg
}

func (b MfhdBox) Version() byte {
	return b.version
}

func (b MfhdBox) Flags() uint32 {
	return b.flags
}

func (b MfhdBox) SequenceNumber() uint32 {
	return b.sequenceNumber
}

func (b *MfhdBox) Parse() {
	buffer := b.ParseData()
	if buffer == nil {
		return
	}
	b.version = buffer[0]
	// SKIP: flags
	b.sequenceNumber = u32(buffer[4:8])
}
