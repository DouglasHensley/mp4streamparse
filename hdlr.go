// Copyright 2025 Douglas Hensley. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mp4streamparse

import (
	"fmt"
	"strings"
)

// HdlrBox defines the hdlr box structure.
type HdlrBox struct {
	Box
	version byte
	flags   uint32
	handler string
	name    string
}

func NewHdlrBox(name string, size uint32, buf []byte) *HdlrBox {
	return &HdlrBox{Box: NewBox(name, int64(size), buf[:size])}
}

func (b HdlrBox) String() string {
	indentLevel := 4
	tab := strings.Repeat("    ", indentLevel)
	strMsg := fmt.Sprintf("%s<< HdlrBox >>", tab)
	strMsg = fmt.Sprintf("%s\tHandler(%s) Name(%s)", strMsg, b.handler, b.name)
	return strMsg
}

func (b HdlrBox) Version() byte {
	return b.version
}

func (b HdlrBox) Flags() uint32 {
	return b.flags
}

func (b HdlrBox) Handler() string {
	return b.handler
}

func (b HdlrBox) Name() string {
	return b.name
}

func (b *HdlrBox) Parse() {
	buffer := b.ParseData()
	if buffer == nil {
		return
	}
	b.version = buffer[0]
	b.flags = u32(buffer[0:4])
	b.handler = string(buffer[8:12])
	b.name = string(buffer[24 : b.size-BoxHeaderSize])
}
