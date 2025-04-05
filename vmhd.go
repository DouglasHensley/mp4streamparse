package mp4streamparse

import (
	"fmt"
	"strings"
)

// VmhdBox defines the vmhd box structure.
type VmhdBox struct {
	Box
	version      byte
	flags        uint32
	graphicsMode uint16
	opColor      uint16
}

func NewVmhdBox(name string, size uint32, buf []byte) *VmhdBox {
	return &VmhdBox{Box: NewBox(name, int64(size), buf)}
}

func (b VmhdBox) String() string {
	indentLevel := 5
	tab := strings.Repeat("\t", indentLevel)
	strMsg := fmt.Sprintf("%s<< VmhdBox >>", tab)
	strMsg = fmt.Sprintf("%s\n%s\tGraphicsMode(%d) OpColor(%d)", strMsg, tab, b.graphicsMode, b.opColor)
	return strMsg
}

func (b VmhdBox) Version() byte {
	return b.version
}

func (b VmhdBox) Flags() uint32 {
	return b.flags
}

func (b VmhdBox) GraphicsMode() uint16 {
	return b.graphicsMode
}

func (b VmhdBox) OpColor() uint16 {
	return b.opColor
}

func (b *VmhdBox) Parse() {
	buffer := b.ParseData()
	if buffer == nil {
		return
	}
	b.version = buffer[0]
	b.flags = u32(buffer[0:4])
	b.graphicsMode = u16(buffer[4:6])
	b.opColor = u16(buffer[6:8])
}
