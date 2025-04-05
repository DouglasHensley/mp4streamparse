package mp4streamparse

import (
	"fmt"
	"strings"
)

// MoofBox defines the moof box structure.
type MoofBox struct {
	Box
	mfhd        *MfhdBox
	traf        *TrafBox
	prevSeqNo   uint32
	elapsedTime float64
}

func NewMoofBox(name string, size uint32, buf []byte) *MoofBox {
	return &MoofBox{Box: NewBox(name, int64(size), buf[:size])}
}

func (b MoofBox) String() string {
	indentLevel := 1
	tab := strings.Repeat("\t", indentLevel)
	strMsg := fmt.Sprintf("%s<< MoofBox >>", tab)
	strMsg = fmt.Sprintf("%s\n%s\tPrevSeqNo(%d) ElapsedTime(%f)", strMsg, tab, b.prevSeqNo, b.elapsedTime)
	strMsg = fmt.Sprintf("%s\n%s\t%s", strMsg, tab, b.mfhd)
	strMsg = fmt.Sprintf("%s\n%s\t%s", strMsg, tab, b.traf)
	return strMsg
}

// IsIFrame returns I-Frame indicator from child data
func (b *MoofBox) IsIFrame() bool {
	return b.traf.trun.sampleDependsOnFlag == 2
}

func (b MoofBox) Mfhd() *MfhdBox {
	return b.mfhd
}

func (b MoofBox) Traf() *TrafBox {
	return b.traf
}

func (b MoofBox) PrevSeqNo() uint32 {
	return b.prevSeqNo
}

func (b MoofBox) ElapsedTime() float64 {
	return b.elapsedTime
}

func (b *MoofBox) Parse() {
	buffer := b.ParseData()
	if buffer == nil {
		return
	}
	for offset := 0; offset < len(buffer); {
		buf := buffer[offset:]
		size, name := ParseHeader(buf)
		if size == 0 {
			return
		}

		switch name {
		case "mfhd": // Mandatory
			b.mfhd = NewMfhdBox(string(name), size, buf[:size])
			b.mfhd.Parse()

		case "traf": // Mandatory
			b.traf = NewTrafBox(string(name), size, buf[:size])
			b.traf.Parse()
		}
		offset += int(size)
	}
	b.elapsedTime = float64(G_SampleDuration_accum) / float64(G_Timescale)
}
