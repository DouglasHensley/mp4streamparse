package mp4streamparse

import (
	"fmt"
	"strings"
)

// MdatBox defines the mdat box structure.
type MdatBox struct {
	Box
}

func NewMdatBox(name string, size uint32, buf []byte) *MdatBox {
	return &MdatBox{Box: NewBox(name, int64(size), buf[:size])}
}

func (b MdatBox) String() string {
	indentLevel := 1
	tab := strings.Repeat("    ", indentLevel)
	strMsg := fmt.Sprintf("\n%s<< MdatBox >>", tab)
	strMsg = fmt.Sprintf("%s\tData Length(%d)", strMsg, len(b.buffer))
	return strMsg
}

func (m *MdatBox) Parse() {
	// This Space Intentionally Left Blank
}
