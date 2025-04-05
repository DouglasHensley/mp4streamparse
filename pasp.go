package mp4streamparse

import (
	"fmt"
	"strings"
)

// PaspBox defines the pasp box structure.
// References
// QuickTime File Format Specification
// developer.apple.com/library/archive/documentation/QuickTime/QTFF/QTFFChap3/
//
//	qtff3.heml#//apple_ref/doc/uid/TP40000939-CH205-74522
//
// Video Sample Description Extension
// 'pasp' - Pixel Aspect Ratio
//
// This extension specifies the height-to-width ratio of pixels found in the video sample.
// This is a required extension for MPEG-4 and uncompressed Y'CbCr video formats when
// non-square pixels are used.
// It is optional when square pixels are used.
// The units of measure for the 'hSpacing' and 'vSpacing' parameters are not specified, as on ly the ration matters.
//
//	hSpacing: 32-bit integer specifying the horizontal spacing of pixels, such as luma sampling instants for Y'CbCr or YUV video
//	vSpacing: 32-bit integer specifying the vertical spacing of pixels, such as video picture lines
//
// The units of measure for height and width must be the same, however.
//
// Common Pixel Aspect Ratios
// Description							 hSpacing	vSpacing
// 4:3 square (composite NTSC or PAL)	 1			1
// 4:3 non-square 525 (NTSC)			 10			11
// 4:3 non-square 625 (PAL)				 59			54
// 16:9 analog (composite NTSC or PAL)	 4			3
// 16:9 digital 525 (NTSC)				 40			33
// 16:9 digital 625 (PAL)				 118		81
// 1920x1035 HDTV (per SMPTE 260M-1992)	 113		118
// 1920x1035 HDTV (per SMPTE RP 187-1995)1018		1062
// 1920x1080 HDTV or 1280X720 HDTV		 1			1
type PaspBox struct {
	Box
	// version  byte
	// hSpacing uint32
	// vSpacing uint32
}

func NewPaspBox(name string, size uint32, buf []byte) *PaspBox {
	return &PaspBox{Box: NewBox(name, int64(size), buf[:size])}
}

func (b PaspBox) String() string {
	indentLevel := 8
	tab := strings.Repeat("    ", indentLevel)
	strMsg := fmt.Sprintf("%s<< PaspBox >>", tab)
	strMsg = fmt.Sprintf("%s\tNot Parsed", strMsg)
	return strMsg
}

func (b *PaspBox) Parse() {
	buffer := b.ParseData()
	if buffer == nil {
		return
	}
	// var offset int = 0
	// var hSpacing uint32 = u32(buffer[offset : offset+4])
	// offset += 4
	// var vSpacing uint32 = u32(buffer[offset : offset+4])
}
