package mp4streamparse

import (
	"fmt"
	"strings"
)

// MehdBox defines the mehd box.
type MehdBox struct {
	Box
	fragmentDuration uint32
}

func NewMehdBox(name string, size uint32, buf []byte) *MehdBox {
	return &MehdBox{Box: NewBox(name, int64(size), buf)}
}

func (b MehdBox) String() string {
	indentLevel := 3
	tab := strings.Repeat("    ", indentLevel)
	strMsg := fmt.Sprintf("%s<< MehdBox >>", tab)
	strMsg = fmt.Sprintf("%s\tFragmentDuration(%d)", strMsg, b.fragmentDuration)
	return strMsg
}

func (b MehdBox) FragmentDuration() uint32 {
	return b.fragmentDuration
}

func (b *MehdBox) Parse() {
	buffer := b.ParseData()
	if buffer == nil {
		return
	}
	b.fragmentDuration = u32(buffer[4:8])
}
