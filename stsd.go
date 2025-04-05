package mp4streamparse

import (
	"fmt"
	"strings"
)

// StsdBox defines the stsd box structure.
type StsdBox struct {
	Box
	version     byte
	flags       uint32
	entryCount  uint32
	sampleEntry []SampleEntry
	avc1        *Avc1Box
}

// BtrtBox for stsd atom
type BtrtBox struct {
	Box
	bufferSizeDB uint32
	maxBitrate   uint32
	avgBitrate   uint32
}

// SampleEntry for stsd atom
type SampleEntry struct {
	dataReferenceIndex uint16
}

func NewStsdBox(name string, size uint32, buf []byte) *StsdBox {
	return &StsdBox{Box: NewBox(name, int64(size), buf[:size])}
}

func (b StsdBox) String() string {
	indentLevel := 6
	tab := strings.Repeat("\t", indentLevel)
	strMsg := fmt.Sprintf("%s<< StsdBox >>", tab)
	strMsg = fmt.Sprintf("%s\n%s\tEntryCount(%d)", strMsg, tab, b.entryCount)
	strMsg = fmt.Sprintf("%s\n%s\tSampleEntry(%v)", strMsg, tab, b.sampleEntry)
	strMsg = fmt.Sprintf("%s\n%s\tSampleDeltas(%v)", strMsg, tab, b.avc1)
	return strMsg
}

func (b StsdBox) Version() byte {
	return b.version
}

func (b StsdBox) Flags() uint32 {
	return b.flags
}

func (b StsdBox) EntryCount() uint32 {
	return b.entryCount
}

func (b StsdBox) SampleEntry() []SampleEntry {
	return b.sampleEntry
}

func (b StsdBox) Avc1() *Avc1Box {
	return b.avc1
}

func (b *StsdBox) Parse() {
	buffer := b.ParseData()
	if buffer == nil {
		return
	}
	// FullBox Data
	b.version = buffer[0]
	b.flags = u32(buffer[0:4])

	b.entryCount = u32(buffer[4:8])
	b.sampleEntry = make([]SampleEntry, b.entryCount)

	// NOTE:
	// THIS DOES NOT LOOK RIGHT
	offset := 8
	for i := 0; i < int(b.entryCount); i++ {
		buf := buffer[offset:]
		size, name := ParseHeader(buf)
		if size == 0 {
			return
		}

		switch name {
		case "avc1":
			b.avc1 = NewAvc1Box(string(name), size, buf[:size])
			b.avc1.Parse()
		}
	}
}
