// Copyright 2025 Douglas Hensley. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mp4streamparse

import (
	"fmt"
	"strings"
)

// TrakBox defines the trak box structure.
type TrakBox struct {
	Box
	tkhd *TkhdBox
	mdia *MdiaBox
}

func NewTrakBox(name string, size uint32, buf []byte) *TrakBox {
	return &TrakBox{Box: NewBox(name, int64(size), buf)}
}

func (b TrakBox) String() string {
	indentLevel := 2
	tab := strings.Repeat("    ", indentLevel)
	tab1 := strings.Repeat("    ", indentLevel+1)
	strMsg := fmt.Sprintf("%s<< TrakBox >>", tab)
	strMsg = fmt.Sprintf("%s\n%s%s", strMsg, tab1, b.tkhd)
	strMsg = fmt.Sprintf("%s\n%s%s", strMsg, tab1, b.mdia)
	return strMsg
}

func (b TrakBox) Tkhd() *TkhdBox {
	return b.tkhd
}

func (b TrakBox) Mdia() *MdiaBox {
	return b.mdia
}

func (b *TrakBox) Parse() {
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
		case "tkhd":
			b.tkhd = NewTkhdBox(string(name), size, buf[:size])
			b.tkhd.Parse()

		case "mdia":
			b.mdia = NewMdiaBox(string(name), size, buf[:size])
			b.mdia.Parse()
		}
		offset += int(size)
	}

}
