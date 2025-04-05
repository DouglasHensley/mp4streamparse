package mp4streamparse

import (
	"fmt"
	"strings"
)

// MinfBox defines the minf box structure.
type MinfBox struct {
	Box
	vmhd *VmhdBox
	stbl *StblBox
	hmhd *HmhdBox
}

func NewMinfBox(name string, size uint32, buf []byte) *MinfBox {
	return &MinfBox{Box: NewBox(name, int64(size), buf)}
}

func (b MinfBox) String() string {
	indentLevel := 4
	tab := strings.Repeat("\t", indentLevel)
	strMsg := fmt.Sprintf("%s<< MinfBox >>", tab)
	strMsg = fmt.Sprintf("%s\n%s\t%s", strMsg, tab, b.vmhd)
	strMsg = fmt.Sprintf("%s\n%s\t%s", strMsg, tab, b.stbl)
	strMsg = fmt.Sprintf("%s\n%s\t%s", strMsg, tab, b.hmhd)
	return strMsg
}

func (b MinfBox) Vmhd() *VmhdBox {
	return b.vmhd
}

func (b MinfBox) Stbl() *StblBox {
	return b.stbl
}

func (b MinfBox) Hmhd() *HmhdBox {
	return b.hmhd
}

func (b *MinfBox) Parse() {
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
		case "vmhd":
			b.vmhd = NewVmhdBox(string(name), size, buf[:size])
			b.vmhd.Parse()

		case "stbl":
			b.stbl = NewStblBox(string(name), size, buf[:size])
			b.stbl.Parse()

		case "hmhd":
			b.hmhd = NewHmhdBox(string(name), size, buf[:size])
			b.hmhd.Parse()
		}
		offset += int(size)
	}
}
