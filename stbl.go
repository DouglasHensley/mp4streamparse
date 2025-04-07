// Copyright 2025 Douglas Hensley. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mp4streamparse

import (
	"fmt"
	"strings"
)

// StblBox defines the stbl box structure.
type StblBox struct {
	Box
	stts *SttsBox
	stsd *StsdBox
}

func NewStblBox(name string, size uint32, buf []byte) *StblBox {
	return &StblBox{Box: NewBox(name, int64(size), buf[:size])}
}

func (b StblBox) String() string {
	indentLevel := 5
	tab := strings.Repeat("    ", indentLevel)
	tab1 := strings.Repeat("    ", indentLevel+1)
	strMsg := fmt.Sprintf("%s<< StblBox >>", tab)
	strMsg = fmt.Sprintf("%s\n%s%s", strMsg, tab1, b.stts)
	strMsg = fmt.Sprintf("%s\n%s%s", strMsg, tab1, b.stsd)
	return strMsg
}

func (b StblBox) Stts() *SttsBox {
	return b.stts
}

func (b StblBox) Stsd() *StsdBox {
	return b.stsd
}

func (b *StblBox) Parse() {
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
		case "stts":
			b.stts = NewSttsBox(string(name), size, buf[:size])
			b.stts.Parse()

		case "stsd":
			b.stsd = NewStsdBox(string(name), size, buf[:size])
			b.stsd.Parse()
		}
		offset += int(size)
	}
}
