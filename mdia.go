// Copyright 2025 Douglas Hensley. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mp4streamparse

import (
	"fmt"
	"strings"
)

// MdiaBox defines the mdia box structure.
type MdiaBox struct {
	Box
	hdlr *HdlrBox
	mdhd *MdhdBox
	minf *MinfBox
}

func NewMdiaBox(name string, size uint32, buf []byte) *MdiaBox {
	return &MdiaBox{Box: NewBox(name, int64(size), buf)}
}

func (b MdiaBox) String() string {
	indentLevel := 3
	tab := strings.Repeat("    ", indentLevel)
	tab1 := strings.Repeat("    ", indentLevel+1)
	strMsg := fmt.Sprintf("%s<< MdiaBox >>", tab)
	strMsg = fmt.Sprintf("%s\n%s%s", strMsg, tab1, b.hdlr)
	strMsg = fmt.Sprintf("%s\n%s%s", strMsg, tab1, b.mdhd)
	strMsg = fmt.Sprintf("%s\n%s%s", strMsg, tab1, b.minf)
	return strMsg
}

func (b MdiaBox) Hdlr() *HdlrBox {
	return b.hdlr
}

func (b MdiaBox) Mdhd() *MdhdBox {
	return b.mdhd
}

func (b MdiaBox) Minf() *MinfBox {
	return b.minf
}

func (b *MdiaBox) Parse() {
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
		case "hdlr":
			b.hdlr = NewHdlrBox(name, size, buf[:size])
			b.hdlr.Parse()

		case "mdhd":
			b.mdhd = NewMdhdBox(name, size, buf[:size])
			b.mdhd.Parse()

		case "minf":
			b.minf = NewMinfBox(name, size, buf[:size])
			b.minf.Parse()
		}
		offset += int(size)
	}
}
