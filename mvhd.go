package mp4streamparse

import (
	"fmt"
	"strings"
)

// MvhdBox defines the mvhd box structure.
// This box defines overall information which is media-independent, and relevant
// to the entire presentation considered as a whole.
//
// Timescale: is an integer that specifies the time-scale for the entire presentation;
//
//	this is the number of time units that pass in one second.
//
// Duration: is an integer that decleares length of the presentation (in the indicated timescale).
//
//	This property is derived from the presentation's tracks; the value of this field
//	corresponds to the duration of the longest track in the presentation. If the
//	duration cannot be determined then duration is set to all 1s.
type MvhdBox struct {
	Box
	version byte
	// flags            uint32
	// creationTime     uint32
	// modificationTime uint32
	timescale uint32
	duration  uint32
	rate      Fixed32
	volume    Fixed16
}

func NewMvhdBox(name string, size uint32, buf []byte) *MvhdBox {
	return &MvhdBox{Box: NewBox(name, int64(size), buf[:size])}
}

func (b MvhdBox) String() string {
	indentLevel := 2
	tab := strings.Repeat("    ", indentLevel)
	strMsg := fmt.Sprintf("%s<< MvhdBox >>", tab)
	strMsg = fmt.Sprintf("%s\tTimescale(%d) Duration(%d)", strMsg, b.timescale, b.duration)
	return strMsg
}

func (b MvhdBox) Version() byte {
	return b.version
}

func (b MvhdBox) Timescale() uint32 {
	return b.timescale
}

func (b MvhdBox) Duration() uint32 {
	return b.duration
}

func (b MvhdBox) Rate() Fixed32 {
	return b.rate
}

func (b MvhdBox) Volume() Fixed16 {
	return b.volume
}

func (b *MvhdBox) Parse() {
	buffer := b.ParseData()
	if buffer == nil {
		return
	}
	b.version = buffer[0]
	b.timescale = u32(buffer[12:16])
	b.duration = u32(buffer[16:20])
	b.rate = fixed32(buffer[20:24])
	b.volume = fixed16(buffer[24:26])
}
