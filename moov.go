// Copyright 2025 Douglas Hensley. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mp4streamparse

import (
	"fmt"
	"strings"
)

// Flag constants.
const (
	TrackFlagEnabled   = 0x0001
	TrackFlagInMovie   = 0x0002
	TrackFlagInPreview = 0x0004
)

// MoovBox defines the moov box structure.
type MoovBox struct {
	Box
	mvex *MvexBox
	mvhd *MvhdBox
	trak *TrakBox
}

func NewMoovBox(name string, size uint32, buf []byte) *MoovBox {
	return &MoovBox{Box: NewBox(name, int64(size), buf[:size])}
}

func (b MoovBox) String() string {
	indentLevel := 1
	tab := strings.Repeat("    ", indentLevel)
	tab1 := strings.Repeat("    ", indentLevel+1)
	strMsg := fmt.Sprintf("\n%s<< MoovBox >>", tab)
	strMsg = fmt.Sprintf("%s\n%s%s", strMsg, tab1, b.mvex)
	strMsg = fmt.Sprintf("%s\n%s%s", strMsg, tab1, b.mvhd)
	strMsg = fmt.Sprintf("%s\n%s%s", strMsg, tab1, b.trak)
	return strMsg
}

// GetTimescale return timescale from moov atom child
func (b *MoovBox) GetTimescale() uint32 {
	return b.trak.mdia.mdhd.timescale
}

func (b MoovBox) Mvex() *MvexBox {
	return b.mvex
}

func (b MoovBox) Mvhd() *MvhdBox {
	return b.mvhd
}

func (b MoovBox) Trak() *TrakBox {
	return b.trak
}

func (b *MoovBox) Parse() {
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
		case "mvhd":
			b.mvhd = NewMvhdBox(string(name), size, buf[:size])
			b.mvhd.Parse()

		case "trak":
			b.trak = NewTrakBox(string(name), size, buf[:size])
			b.trak.Parse()
		case "mvex":
			b.mvex = NewMvexBox(string(name), size, buf[:size])
			b.mvex.Parse()
		}
		offset += int(size)
	}
}
