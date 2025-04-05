package mp4streamparse

import (
	"fmt"
	"strings"
)

// AvcCBox defines the avcC box structure.
// References
// QuickTime File Format Specification
// developer.apple.com/library/archive/documentation/QuickTime/QTFF/QTFFChap3/
//
//	qtff3.heml#//apple_ref/doc/uid/TP40000939-CH205-74522
//
// Video Sample Description Extension
// 'avcC' - AVC Decode Config Box
//
// This extension specifies the height-to-width ratio of pixels found in the video sample.
// This is a required extension for MPEG-4 and uncompressed Y'CbCr video formats when
// non-square pixels are used.
// It is optional when square pixels are used.
// The units of measure for the 'hSpacing' and 'vSpacing' parameters are not specified, as on ly the ration matters.
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
type AvcCBox struct {
	Box
	version                byte
	h264Profile            byte
	h264CompatibleProfiles byte
	h264Level              byte
}

func NewAvcCBox(name string, size uint32, buf []byte) *AvcCBox {
	return &AvcCBox{Box: NewBox(name, int64(size), buf)}
}

func (b AvcCBox) String() string {
	indentLevel := 8
	tab := strings.Repeat("    ", indentLevel)
	strMsg := fmt.Sprintf("%s<< AvcCBox >>", tab)
	strMsg = fmt.Sprintf("%s\tH264: Profile(%v) CompatibleProfiles(%v) Level(%v)",
		strMsg, b.h264Profile, b.h264CompatibleProfiles, b.h264Level)
	return strMsg
}

func (b AvcCBox) Version() byte {
	return b.version
}

func (b AvcCBox) H264Profile() byte {
	return b.h264Profile
}

func (b AvcCBox) H264CompatibleProfiles() byte {
	return b.h264CompatibleProfiles
}

func (b AvcCBox) H264Level() byte {
	return b.h264Level
}

func (b *AvcCBox) Parse() {
	buffer := b.ParseData()
	if buffer == nil {
		return
	}
	var offset int = 0
	b.version = buffer[offset]
	offset++
	b.h264Profile = buffer[offset]
	offset++
	b.h264CompatibleProfiles = buffer[offset]
	offset++
	b.h264Level = buffer[offset]
	offset++
	// SKIP: Reserved (6 bits - set to '1b')
	// SKIP: Num NAL Units minus 1 (2 bits)
	// SKIP: var nalLength byte = (buffer[offset] & 0x03) + 1
	offset++
	var spsNum byte = buffer[offset] & 0x1F // Number of SPS entries (5 bits)
	offset++
	for i := 0; i < int(spsNum); i++ {
		spsLength := u16(buffer[offset : offset+2])
		offset += 2
		sps := make([]byte, 0)
		for ii := 0; ii < int(spsLength); ii++ {
			sps = append(sps, buffer[offset+ii])
		}
		offset += int(spsLength)
	}
	var ppsNum byte = buffer[offset]
	offset++
	for i := 0; i < int(ppsNum); i++ {
		ppsLength := u16(buffer[offset : offset+2])
		offset += 2
		pps := make([]byte, 0)
		for ii := 0; ii < int(ppsLength); ii++ {
			pps = append(pps, buffer[offset+ii])
		}
		offset += int(ppsLength)
	}
}
