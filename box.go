// Copyright 2025 Douglas Hensley. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mp4streamparse

import (
	"fmt"
)

// Box defines an Atom Box structure.
type Box struct {
	name   string
	size   int64
	buffer []byte
}

func NewBox(name string, size int64, buffer []byte) Box {
	return Box{
		name:   name,
		size:   size,
		buffer: buffer,
	}
}

func (b *Box) String() string {
	strMsg := "<< Box >>"
	strMsg = fmt.Sprintf("%s\nName(%s) Size(%d)", strMsg, b.name, b.size)
	return strMsg
}

func (b *Box) Name() string {
	return b.name
}

func (b *Box) Size() int64 {
	return b.size
}

func (b *Box) Buffer() []byte {
	return b.buffer
}

// ParseData
func (b *Box) ParseData() []byte {
	if b.size <= BoxHeaderSize {
		return nil
	}
	boxSize := u32(b.buffer[0:4])
	if uint32(len(b.buffer)) < boxSize {
		return nil
	}
	return b.buffer[BoxHeaderSize:]
}

// GetDataSize
func (b *Box) GetDataSize() int64 {
	return b.size - BoxHeaderSize
}
