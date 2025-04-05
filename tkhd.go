package mp4streamparse

import (
	"fmt"
	"strings"
)

// TkhdBox defines the track header box structure.
type TkhdBox struct {
	Box
	version          byte
	flags            uint32
	creationTime     uint32
	modificationTime uint32
	trackID          uint32
	duration         uint32
	layer            uint16
	alternateGroup   uint16
	volume           Fixed16
	matrix           []byte
	width            Fixed32
	height           Fixed32
}

func NewTkhdBox(name string, size uint32, buf []byte) *TkhdBox {
	return &TkhdBox{Box: NewBox(name, int64(size), buf)}
}

func (b TkhdBox) String() string {
	indentLevel := 3
	tab := strings.Repeat("\t", indentLevel)
	strMsg := fmt.Sprintf("%s<< TkhdBox >>", tab)
	strMsg = fmt.Sprintf("%s\n%s\tCreationTime(%d) ModificationTime(%d) TrackID(%d) Duration(%d)",
		strMsg, tab, b.creationTime, b.modificationTime, b.trackID, b.duration)
	strMsg = fmt.Sprintf("%s\n%s\tLayer(%d) AlternateGroup(%d) Width(%d) Height(%d)",
		strMsg, tab, b.layer, b.alternateGroup, b.width, b.height)
	return strMsg
}

func (b TkhdBox) Version() byte {
	return b.version
}

func (b TkhdBox) Flags() uint32 {
	return b.flags
}

func (b TkhdBox) CreationTime() uint32 {
	return b.creationTime
}

func (b TkhdBox) ModificationTime() uint32 {
	return b.modificationTime
}

func (b TkhdBox) TrackID() uint32 {
	return b.trackID
}

func (b TkhdBox) Duration() uint32 {
	return b.duration
}

func (b TkhdBox) Layer() uint16 {
	return b.layer
}

func (b TkhdBox) AlternateGroup() uint16 {
	return b.alternateGroup
}

func (b TkhdBox) Volume() Fixed16 {
	return b.volume
}

func (b TkhdBox) Matrix() []byte {
	return b.matrix
}

func (b TkhdBox) Width() Fixed32 {
	return b.width / (1 << 16)
}

func (b TkhdBox) Height() Fixed32 {
	return b.height / (1 << 16)
}

func (b *TkhdBox) Parse() {
	buffer := b.ParseData()
	if buffer == nil {
		return
	}
	b.version = buffer[0]
	b.flags = u32(buffer[0:4])
	b.creationTime = u32(buffer[4:8])
	b.modificationTime = u32(buffer[8:12])
	b.trackID = u32(buffer[12:16])
	b.duration = u32(buffer[20:24])
	b.layer = u16(buffer[32:34])
	b.alternateGroup = u16(buffer[34:36])
	b.volume = fixed16(buffer[36:38])
	b.matrix = buffer[40:76]
	b.width = fixed32(buffer[76:80])
	b.height = fixed32(buffer[80:84])
}
