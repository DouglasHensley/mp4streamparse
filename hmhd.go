// Copyright 2025 Douglas Hensley. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mp4streamparse

import (
	"fmt"
	"strings"
)

// HmhdBox contains general information, independent of the protocol, for hint tracks.
type HmhdBox struct {
	Box
	version    byte
	maxPDUSize uint16
	avgPDUSize uint16
	maxBitrate uint32
	avgBitrate uint32
}

func NewHmhdBox(name string, size uint32, buf []byte) *HmhdBox {
	return &HmhdBox{Box: NewBox(name, int64(size), buf[:size])}
}

func (b HmhdBox) String() string {
	indentLevel := 5
	tab := strings.Repeat("    ", indentLevel)
	strMsg := fmt.Sprintf("%s<< HmhdBox >>", tab)
	strMsg = fmt.Sprintf("%s\tPDU Size: Max(%d) Avg(%d);", strMsg, b.maxPDUSize, b.avgPDUSize)
	strMsg = fmt.Sprintf("%s\tBitRate: Max(%d) Avg(%d);", strMsg, b.maxBitrate, b.avgBitrate)
	return strMsg
}

func (b HmhdBox) Version() byte {
	return b.version
}

func (b HmhdBox) MaxPDUSize() uint16 {
	return b.maxPDUSize
}

func (b HmhdBox) AvgPDUSize() uint16 {
	return b.avgPDUSize
}

func (b HmhdBox) MaxBitrate() uint32 {
	return b.maxBitrate
}

func (b HmhdBox) AvgBitrate() uint32 {
	return b.avgBitrate
}

func (b *HmhdBox) Parse() {
	buffer := b.ParseData()
	if buffer == nil {
		return
	}
	b.version = buffer[0]
	b.maxPDUSize = u16(buffer[0:2])
	b.avgPDUSize = u16(buffer[2:4])
	b.maxBitrate = u32(buffer[4:8])
	b.avgBitrate = u32(buffer[8:12])
}
