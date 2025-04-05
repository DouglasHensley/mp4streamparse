package mp4streamparse

import (
	"fmt"
	"math"
	"strings"
)

// Avc1Box defines the avc1 box structure.
// References
// QuickTime File Format Specification
// developer.apple.com/library/archive/documentation/QuickTime/QTFF/QTFFChap3/
//	qtff3.heml#//apple_ref/doc/uid/TP40000939-CH205-74522
// Width: 16-bit integer that specifies the width of the source image in pixels
// Height: 16-bit integer that specifies the height of the source iamge in pixels
// Horizontal resolution: 32-bit fixed-point number containing the horizontal resolution of the image in pixels per inch
//   The high order 16 bits contain the integer portion of the value, the low-order 8 bits contain the fractional part
// Vertical resolution: 32-bit fixed-point number containing the vertical resolution of the image in pixels per inch
//   The high order 16 bits contain the integer portion of the value, the low-order 8 bits contain the fractional part
// Data Size: 32-bit integer that must be set to 0.
// Frame Count: 16-bit integer that indicates how many frames of compressed data are stored in each sample. Usually set to 1
// Compressor Name: 32-Byte Pascal string containing the name of the compressor that created the image, such as "jpeg"
//   A Pascal-style string has one leading byte(length), followed by 'length' bytes of character data
// Color Depth: 16-bit integer that indicates the pixel depth of the compressed image.
//   Values of 1, 2, 4, 8, 16, 24, and 32 indicate the depth of color images.
//   The value of 32 should be used only if the image contains an alpha channel.
//   Values of 34, 36 and 40 indicate 2-, 4-, and 8-bit grayscale, respectively, for grayscale images.
// Color Table ID: 16-bit integer that indicates which color table to use. If this field is set to -1 (0xFF FF),
//   The default color table should be used for the specified depth. For all depths below
//   16 bits per pixel, ths indicates a standard Macintosh color table for the specified depth.
//   Depths of 16, 24, and 32 have no color table.
//   If the color table ID is set to 0, a color table is contained within the sample description itself.
//   The color table immediately follows the color table ID field in the sample description

type Avc1Box struct {
	Box
	version        byte
	dataRefIndex   uint16
	width          uint16
	height         uint16
	hResolution    float32
	vResolution    float32
	frameCount     uint16
	compressorName string
	colorDepth     uint16

	avcC *AvcCBox
	pasp *PaspBox
}

func NewAvc1Box(name string, size uint32, buf []byte) *Avc1Box {
	return &Avc1Box{Box: NewBox(name, int64(size), buf)}
}

func (b Avc1Box) String() string {
	indentLevel := 7
	tab := strings.Repeat("    ", indentLevel)
	tab1 := strings.Repeat("    ", indentLevel+1)
	strMsg := fmt.Sprintf("%s<< Avc1Box >>", tab)
	strMsg = fmt.Sprintf("%s\tDataRefIndex(%d)", strMsg, b.dataRefIndex)
	strMsg = fmt.Sprintf("%s  Width(%d) Height(%d) H-Resolution(%f) V-Resolution(%f)",
		strMsg, b.width, b.height, b.hResolution, b.vResolution)
	strMsg = fmt.Sprintf("%s\n%s%s", strMsg, tab1, b.avcC)
	strMsg = fmt.Sprintf("%s\n%s%s", strMsg, tab1, b.pasp)
	return strMsg
}

func (b Avc1Box) Version() byte {
	return b.version
}

func (b Avc1Box) DataRefIndex() uint16 {
	return b.dataRefIndex
}

func (b Avc1Box) Width() uint16 {
	return b.width
}

func (b Avc1Box) Height() uint16 {
	return b.height
}

func (b Avc1Box) HResolution() float32 {
	return b.hResolution
}

func (b Avc1Box) VResolution() float32 {
	return b.vResolution
}

func (b Avc1Box) FrameCount() uint16 {
	return b.frameCount
}

func (b Avc1Box) CompressorName() string {
	return b.compressorName
}

func (b Avc1Box) ColorDepth() uint16 {
	return b.colorDepth
}

func (b Avc1Box) AvcC() *AvcCBox {
	return b.avcC
}

func (b Avc1Box) Pasp() *PaspBox {
	return b.pasp
}

func (b *Avc1Box) Parse() {
	buffer := b.ParseData()
	if buffer == nil {
		return
	}
	offset := 6 // Skip Reserved Bytes (6)
	b.dataRefIndex = u16(buffer[offset : offset+2])
	offset += 2
	offset += 16                             // Skip Predefines and Reserved Bytes
	b.width = u16(buffer[offset : offset+2]) // width of the source image in pixels
	offset += 2
	b.height = u16(buffer[offset : offset+2]) // height of the source image in pixels
	offset += 2
	// horizontal resolution of the image in pixels per inch
	var valInt uint16 = u16(buffer[offset : offset+2])    // high order 16 bits contain the integer portion of the value
	var valFrac uint16 = u16(buffer[offset+2 : offset+4]) // the low-order 8 bits contain the fractional part
	b.hResolution = float32(valInt) + float32(valFrac)/float32(math.MaxUint16)
	offset += 4
	// vertical resolution of the image in pixels per inch
	valInt = u16(buffer[offset : offset+2])    // high order 16 bits contain the integer portion of the value
	valFrac = u16(buffer[offset+2 : offset+4]) // the low-order 8 bits contain the fractional part
	b.vResolution = float32(valInt) + float32(valFrac)/float32(math.MaxUint16)
	offset += 4
	offset += 4                                   // Skip data size
	b.frameCount = u16(buffer[offset : offset+2]) // Number of frames stored in each sample
	offset += 2
	var compressorLength uint8 = buffer[offset] // the name of the compressor that created the image, such as "jpeg"
	b.compressorName = string(buffer[offset+1 : offset+1+int(compressorLength)])
	offset += 32
	// pixel depth of the compressed image.
	b.colorDepth = u16(buffer[offset : offset+2])
	offset += 2
	// indicates which color table to use
	// NOTE COLOR TABLE IS NOT CURRENTLY SUPPORTED
	offset += 2
	// for ; offset<int(b.Size); {
	// 	buf := buffer[offset:]
	// 	size, name := ParseHeader(buf)
	// 	switch name {
	// 	case "avcC":
	// 		b.avcC = NewAvcCBox(string(name), size, buf[:size])
	// 		b.avcC.Parse(indent+1)
	// 	case "pasp":
	// 		b.pasp = NewPaspBox(string(name), size, buf[:size])
	// 		b.pasp.Parse(indent+1)
	// 	}
	// 	offset += int(size)
	// }
}
