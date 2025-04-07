// Copyright 2025 Douglas Hensley. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mp4streamparse

import (
	"fmt"
	"strings"
)

// TfdtBox defines the tfdt box.
type TfdtBox struct {
	Box
	baseMediaDecodeTime uint32
}

func NewTfdtBox(name string, size uint32, buf []byte) *TfdtBox {
	return &TfdtBox{Box: NewBox(name, int64(size), buf)}
}

func (b TfdtBox) String() string {
	indentLevel := 3
	tab := strings.Repeat("    ", indentLevel)
	strMsg := fmt.Sprintf("%s<< TfdtBox >>", tab)
	strMsg = fmt.Sprintf("%s\tBaseMediaDecodeTime(%d)", strMsg, b.baseMediaDecodeTime)
	return strMsg
}

func (b TfdtBox) BaseMediaDecodeTime() uint32 {
	return b.baseMediaDecodeTime
}

func (b *TfdtBox) Parse() {
	buffer := b.ParseData()
	if buffer == nil {
		return
	}
	// SKIP: Version and Flags are first 4 bytes
	b.baseMediaDecodeTime = u32(buffer[4:8])
}
