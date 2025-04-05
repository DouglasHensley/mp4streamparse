package mp4streamparse

import (
	"fmt"
	"strings"
)

// MvexBox defines the mvex box structure.
type MvexBox struct {
	Box
	mehd *MehdBox
}

func NewMvexBox(name string, size uint32, buf []byte) *MvexBox {
	return &MvexBox{Box: NewBox(name, int64(size), buf[:size])}
}

func (b MvexBox) String() string {
	indentLevel := 2
	tab := strings.Repeat("    ", indentLevel)
	strMsg := fmt.Sprintf("%s<< MvexBox >>", tab)
	strMsg = fmt.Sprintf("%s\t%s", strMsg, b.mehd)
	return strMsg
}

func (b MvexBox) Mehd() *MehdBox {
	return b.mehd
}

func (b *MvexBox) Parse() {
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
		case "mehd":
			b.mehd = NewMehdBox(string(name), size, buf[:size])
			b.mehd.Parse()
		}
		offset += int(size)
	}
}
