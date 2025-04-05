package mp4streamparse

import (
	"fmt"
	"strings"
)

// TrafBox defines the traf box structure.
type TrafBox struct {
	Box
	tfhd *TfhdBox
	tfdt *TfdtBox
	trun *TrunBox
}

func NewTrafBox(name string, size uint32, buf []byte) *TrafBox {
	return &TrafBox{Box: NewBox(name, int64(size), buf)}
}

func (b TrafBox) String() string {
	indentLevel := 2
	tab := strings.Repeat("    ", indentLevel)
	tab1 := strings.Repeat("    ", indentLevel+1)
	strMsg := fmt.Sprintf("%s<< TrafBox >>", tab)
	strMsg = fmt.Sprintf("%s\n%s%s", strMsg, tab1, b.tfhd)
	strMsg = fmt.Sprintf("%s\n%s%s", strMsg, tab1, b.tfdt)
	strMsg = fmt.Sprintf("%s\n%s%s", strMsg, tab1, b.trun)
	return strMsg
}

func (b TrafBox) Tfhd() *TfhdBox {
	return b.tfhd
}

func (b TrafBox) Tfdt() *TfdtBox {
	return b.tfdt
}

func (b TrafBox) Trun() *TrunBox {
	return b.trun
}

func (b *TrafBox) Parse() {
	buffer := b.ParseData()
	for offset := 0; offset < len(buffer); {
		buf := buffer[offset:]
		if buffer == nil {
			return
		}
		size, name := ParseHeader(buf)
		if size == 0 {
			return
		}

		switch name {
		case "tfhd": // Mandatory
			b.tfhd = NewTfhdBox(string(name), size, buf[:size])
			b.tfhd.Parse()

		case "tfdt":
			b.tfdt = NewTfdtBox(string(name), size, buf[:size])
			b.tfdt.Parse()

		case "trun":
			b.trun = NewTrunBox(string(name), size, buf[:size])
			b.trun.Parse()
		}
		offset += int(size)
	}
	if b.trun.sampleDurationPresentFlag {
		for _, item := range b.trun.sampleArray {
			GSampleDurationAccum(item.sampleDuration)
		}
	} else {
		GSampleDurationAccum(b.tfhd.defaultSampleDuration * b.trun.sampleCount)
	}
}
