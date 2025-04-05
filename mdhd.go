package mp4streamparse

import (
	"fmt"
	"strings"
)

// MdhdBox defines the mdhd box structure.
type MdhdBox struct {
	Box
	version          byte
	flags            uint32
	creationTime     uint32
	modificationTime uint32
	timescale        uint32
	duration         uint32
	language         uint16
	languageString   string
}

func NewMdhdBox(name string, size uint32, buf []byte) *MdhdBox {
	return &MdhdBox{Box: NewBox(name, int64(size), buf)}
}

func (b MdhdBox) String() string {
	indentLevel := 4
	tab := strings.Repeat("\t", indentLevel)
	strMsg := fmt.Sprintf("%s<< MdhdBox >>", tab)
	strMsg = fmt.Sprintf("%s\n%s\tCreationTime(%d) ModificationTime(%d)", strMsg, tab, b.creationTime, b.modificationTime)
	strMsg = fmt.Sprintf("%s\n%s\tTimescale(%d) Duration(%d) Language(%d, %s)", strMsg, tab, b.timescale, b.duration, b.language, b.languageString)
	return strMsg
}

func (b MdhdBox) Version() byte {
	return b.version
}

func (b MdhdBox) Flags() uint32 {
	return b.flags
}

func (b MdhdBox) CreationTime() uint32 {
	return b.creationTime
}

func (b MdhdBox) ModificationTime() uint32 {
	return b.modificationTime
}

func (b MdhdBox) Timescale() uint32 {
	return b.timescale
}

func (b MdhdBox) Duration() uint32 {
	return b.duration
}

func (b MdhdBox) Language() uint16 {
	return b.language
}

func (b MdhdBox) LanguageString() string {
	return b.languageString
}

func (b *MdhdBox) Parse() {
	buffer := b.ParseData()
	if buffer == nil {
		return
	}
	b.version = buffer[0]
	b.flags = u32(buffer[0:4])
	b.creationTime = u32(buffer[4:8])
	b.modificationTime = u32(buffer[8:12])
	b.timescale = u32(buffer[12:16])
	b.duration = u32(buffer[16:20])
	b.language = u16(buffer[20:22])
	b.languageString = getLanguageString(b.language)

	Set_G_Timescale(b.timescale)
}

func getLanguageString(language uint16) string {
	var lang [3]uint16
	lang[0] = (language >> 10) & 0x1F
	lang[1] = (language >> 5) & 0x1F
	lang[2] = (language) & 0x1F

	return fmt.Sprintf("%s%s%s",
		string(lang[0]+0x60),
		string(lang[1]+0x60),
		string(lang[2]+0x60))
}
