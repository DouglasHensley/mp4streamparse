package mp4streamparse

import (
	"bytes"
	"fmt"
	"time"
)

// Package Namespace Global Values
var G_Timescale uint32 = 30000

func Set_G_Timescale(timescale uint32) {
	G_Timescale = timescale // Reset with MOOV->TRAK->MDIA->MDHD.Timescale
}

var G_SampleDuration_accum uint32 = 0

func GSampleDurationAccum(sampleDuration uint32) {
	G_SampleDuration_accum += sampleDuration
}

const (
	BoxHeaderSize = int64(8) // BoxHeaderSize Size of box header
)

var (
	ftyp = []byte("ftyp")
	styp = []byte("styp")
	moov = []byte("moov")
	moof = []byte("moof")
	mdat = []byte("mdat")
)

// TimestampBox mp4 atom with timestamp
type TimestampBox struct {
	TimestampUnixNano int64
	Mp4Box            interface{}
}

func (tsBox *TimestampBox) String() string {
	strMsg := fmt.Sprintf("\n## TimestampBox ##")
	strMsg = fmt.Sprintf("%s%s", strMsg, tsBox.Mp4Box)
	return strMsg
}

// ParseHeader (size, type) from an offset.
func ParseHeader(buffer []byte) (uint32, string) {
	if int64(len(buffer)) <= BoxHeaderSize {
		return 0, ""
	}
	return u32(buffer[0:4]), string(buffer[4:8])
}

// FindNextBox: ftyp, styp, moov, moof, mdat
func FindNextBox(buffer []byte) (int, uint32, bool) {
	if int64(len(buffer)) < BoxHeaderSize {
		return 0, 0, false
	}

	var i int = 0
	for ; i < len(buffer)-8; i++ {
		matchBuf := buffer[i+4 : i+8]
		// Arranged for hi frequency matches first
		if bytes.Equal(moof, matchBuf) || bytes.Equal(mdat, matchBuf) ||
			bytes.Equal(styp, matchBuf) || bytes.Equal(ftyp, matchBuf) ||
			bytes.Equal(moov, matchBuf) {
			var boxSize uint32 = u32(buffer[i : i+4])
			// Box Size larger than available buffer: FAIL
			if boxSize > uint32(len(buffer[i:])) {
				return i, boxSize, false
			}
			return i, boxSize, true
		}
	}
	// Buffer Searched, No Boxes Located
	return 0, 0, false
}

// ReadBoxes Search buffer for top level boxes; recurse for sub-boxes
func ReadBoxes(buffer []byte, mp4BoxTree []string) (uint32, []string) {
	var offset uint32
	// var prevSeqNo uint32
	for offset = 0; offset < (uint32)(len(buffer)); {
		buf := buffer[offset:]
		size, name := ParseHeader(buf)
		if size == 0 {
			return offset, mp4BoxTree
		}

		// Recurse for sub-boxes
		switch string(name) {
		// Initialization Boxes
		case "ftyp":
			mp4box := NewFtypBox(string(name), size, buf[:size])
			mp4box.Parse()
			mp4BoxTree = append(mp4BoxTree, mp4box.String())
		case "styp":
			mp4box := NewStypBox(string(name), size, buf[:size])
			mp4box.Parse()
			mp4BoxTree = append(mp4BoxTree, mp4box.String())
		case "moov":
			mp4box := NewMoovBox(string(name), size, buf[:size])
			mp4box.Parse()
			mp4BoxTree = append(mp4BoxTree, mp4box.String())

		// Data Boxes
		case "moof":
			mp4box := NewMoofBox(string(name), size, buf[:size])
			// if mp4box != nil {
			// 	prevSeqNo = mp4box.mfhd.sequenceNumber
			// } else {
			// 	prevSeqNo = 0
			// }
			mp4box.Parse()
			mp4BoxTree = append(mp4BoxTree, mp4box.String())
		case "mdat":
			boxSize := size
			if boxSize == 0 {
				boxSize = uint32(len(buf))
			}
			mp4box := NewMdatBox(string(name), boxSize, buf[:boxSize])
			mp4box.Parse()
			mp4BoxTree = append(mp4BoxTree, mp4box.String())
		}
		offset += size
	}
	return offset, mp4BoxTree
}

// ReadBox Search buffer for a top level box; recurse for sub-boxes
func ReadBox(buffer []byte) (uint32, *TimestampBox) {
	buf := buffer[:]
	size, name := ParseHeader(buf)
	if size == 0 {
		return 0, nil
	}
	timestampUnixNano := time.Now().UnixNano()

	// Recurse for sub-boxes
	switch name {
	// Initialization Boxes
	case "ftyp":
		mp4Box := NewFtypBox(name, size, buf)
		mp4Box.Parse()
		return size, &TimestampBox{timestampUnixNano, mp4Box}
	case "styp":
		mp4Box := NewStypBox(name, size, buf)
		mp4Box.Parse()
		return size, &TimestampBox{timestampUnixNano, mp4Box}
	case "moov":
		mp4Box := NewMoovBox(name, size, buf)
		mp4Box.Parse()
		return size, &TimestampBox{timestampUnixNano, mp4Box}
	// Data Boxes
	case "moof":
		mp4Box := NewMoofBox(name, size, buf)
		mp4Box.Parse()
		return size, &TimestampBox{timestampUnixNano, mp4Box}
	case "mdat":
		mp4Box := NewMdatBox(name, size, buf)
		mp4Box.Parse()
		return size, &TimestampBox{timestampUnixNano, mp4Box}
	}
	return size, &TimestampBox{timestampUnixNano, nil}
}
