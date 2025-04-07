// Copyright 2025 Douglas Hensley. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mp4streamparse

import (
	"fmt"
	"strings"
)

// SttsBox defines the stts box structure.
type SttsBox struct {
	Box
	version      byte
	flags        uint32
	entryCount   uint32
	sampleCounts []uint32
	sampleDeltas []uint32
}

func NewSttsBox(name string, size uint32, buf []byte) *SttsBox {
	return &SttsBox{Box: NewBox(name, int64(size), buf[:size])}
}

func (b SttsBox) String() string {
	indentLevel := 6
	tab := strings.Repeat("    ", indentLevel)
	strMsg := fmt.Sprintf("%s<< SttsBox >>", tab)
	strMsg = fmt.Sprintf("%s\tEntryCount(%d)", strMsg, b.entryCount)
	strMsg = fmt.Sprintf("%s\tSampleCounts(%v)", strMsg, b.sampleCounts)
	strMsg = fmt.Sprintf("%s\tSampleDeltas(%v)", strMsg, b.sampleDeltas)
	return strMsg
}

func (b SttsBox) Version() byte {
	return b.version
}

func (b SttsBox) Flags() uint32 {
	return b.flags
}

func (b SttsBox) EntryCount() uint32 {
	return b.entryCount
}

func (b SttsBox) SampleCounts() []uint32 {
	return b.sampleCounts
}

func (b SttsBox) SampleDeltas() []uint32 {
	return b.sampleDeltas
}

func (b *SttsBox) Parse() {
	buffer := b.ParseData()
	if buffer == nil {
		return
	}
	b.version = buffer[0]
	b.flags = u32(buffer[0:4])

	b.entryCount = u32(buffer[4:8])
	b.sampleCounts = make([]uint32, b.entryCount)
	b.sampleDeltas = make([]uint32, b.entryCount)

	for i := 0; i < int(b.entryCount); i++ {
		b.sampleCounts[i] = u32(buffer[(8 + 8*i):(12 + 8*i)])
		b.sampleDeltas[i] = u32(buffer[(12 + 8*i):(16 + 8*i)])
	}
}
