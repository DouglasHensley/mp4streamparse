package mp4streamparse

import (
	"fmt"
	"strings"
)

// StypBox defines the styp box.
type StypBox struct {
	Box
	majorBrand       string
	minorVersion     uint32
	compatibleBrands []string
}

func NewStypBox(name string, size uint32, buf []byte) *StypBox {
	return &StypBox{Box: NewBox(name, int64(size), buf)}
}

func (b StypBox) String() string {
	indentLevel := 1
	tab := strings.Repeat("\t", indentLevel)
	strMsg := fmt.Sprintf("%s<< StypBox >>", tab)
	strMsg = fmt.Sprintf("%s\n%s\tMajorBrand(%s) MinorVersion(%d) CompatibleBrands(%v)", strMsg, tab, b.majorBrand, b.minorVersion, b.compatibleBrands)
	return strMsg
}

func (b StypBox) MajorBrand() string {
	return b.majorBrand
}

func (b StypBox) MinorVersion() uint32 {
	return b.minorVersion
}

func (b StypBox) CompatibleBrands() []string {
	return b.compatibleBrands
}

func (b *StypBox) Parse() {
	buffer := b.ParseData()
	if buffer == nil {
		return
	}
	b.majorBrand, b.minorVersion = string(buffer[0:4]), u32(buffer[4:8])
	if len(buffer) > 8 {
		for i := 8; i < len(buffer); i += 4 {
			b.compatibleBrands = append(b.compatibleBrands, string(buffer[i:i+4]))
		}
	}
}
