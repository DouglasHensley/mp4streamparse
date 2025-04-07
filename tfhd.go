// Copyright 2025 Douglas Hensley. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mp4streamparse

import (
	"encoding/binary"
	"fmt"
	"strings"
)

// Flag constants.
// BaseDataOffsetPresent: indicates the presence of the BaseDataOffset field.
//
//	This provides an explicit anchor for the data offsets in each track run.
//	If not provided and if the default-base-is-moof flag is not set, the
//	base-data-offset for the first track in the movie fragment is the position
//	of the first byte of the enclosing Movie Fragment Box, and for second and
//	subsequent track fragments, the default is the end of the data defined by
//	the preceding track fragment. Fragments 'inheriting' the offset in this way
//	must all use the same data-reference (i.e. the data for these tracks must be
//	in the same file)
//
// SampleDescriptionIndexPresent: indicates the rpesence of this field, which over-rides,
//
//	in this fragment, the default set up in the Track Extends Box.
//
// DurationIsEmpty: indicates the duration provided in either default-sample-duration
//
//	or by the default-duration in the Track Extends Box, is empty, i.e. that there
//	are no samples for this time interval. It is an error to make a presentation that
//	has both edit lists in the Movie Box, and empty-duration fragments.
//
// DefaultBaseIsMoof: if base-data-offset-present is 1, this flag is ignored.
//
//	If base-data-offset-present is zero, this indicates that the base-data-offset for
//	this track fragment is the position of the first byte of the enclosing Movie
//	Fragment Box. Support for the default-base-is-moof flag is required under the
//	'iso5' brand, and it shall not be used in brands or compatible brands earlier
//	than 'iso5'.
//
// IMPLEMENTATION NOTE:
// The Flags field is actually 24-bits (3 bytes).
// This implementation uses UINT32 (4 bytes) an includes the Version byte.
// All of the const fields for interpreting the Flags field are padded with
// an additional byte to screen out the Version byte.
const (
	BaseDataOffsetPresent         uint32 = 0x00000001
	SampleDescriptionIndexPresent uint32 = 0x00000002
	DefaultSampleDurationPresent  uint32 = 0x00000008
	DefaultSampleSizePresent      uint32 = 0x00000010
	DefaultSampleFlagsPresent     uint32 = 0x00000020
	DurationIsEmpty               uint32 = 0x00010000
	DefaultBaseIsMoof             uint32 = 0x00020000
)

// TfhdBox defines the tfhd box structure.
// The BaseDataOffset, if explicitly provided, is a data offset that is identical to a
// chunk offset in the Chunk Offset Box, i.e. applying to the complete file (e.g. starting
// with a file-type box and movie box). In circumstances when the complete file does not
// exist or its size is unknown, it may be impossible to use an explicit base-data-offset;
// then, offsets need to be established relative to the movie fragment.
type TfhdBox struct {
	Box
	version byte
	flags   uint32
	trackID uint32
	// Optional Fields
	baseDataOffsetPresentFlag         bool
	sampleDescriptionIndexPresentFlag bool
	defaultSampleDurationPresentFlag  bool
	defaultSampleSizePresentFlag      bool
	defaultSampleFlagsPresentFlag     bool
	durationIsEmptyFlag               bool
	defaultBaseIsMoofFlag             bool
	baseDataOffset                    uint64
	sampleDescriptionIndex            uint32
	defaultSampleDuration             uint32
	defaultSampleSize                 uint32
	defaultSampleFlags                uint32
}

func NewTfhdBox(name string, size uint32, buf []byte) *TfhdBox {
	return &TfhdBox{Box: NewBox(name, int64(size), buf)}
}

func (b TfhdBox) String() string {
	indentLevel := 3
	tab := strings.Repeat("    ", indentLevel)
	strMsg := fmt.Sprintf("%s<< TfhdBox >>", tab)
	strMsg = fmt.Sprintf("%s\tTrackID(%d) BaseDataOffset(%d) SampleDescriptionIndex(%d) Default Sample: Duration(%d) Size(%d) Flags(%d)",
		strMsg, b.trackID, b.baseDataOffset, b.sampleDescriptionIndex, b.defaultSampleDuration, b.defaultSampleSize, b.defaultSampleFlags)
	return strMsg
}

func (b TfhdBox) Version() byte {
	return b.version
}

func (b TfhdBox) Flags() uint32 {
	return b.flags
}

func (b TfhdBox) TrackID() uint32 {
	return b.trackID
}

func (b TfhdBox) BaseDataOffsetPresentFlag() bool {
	return b.baseDataOffsetPresentFlag
}

func (b TfhdBox) SampleDescriptionIndexPresentFlag() bool {
	return b.sampleDescriptionIndexPresentFlag
}

func (b TfhdBox) DefaultSampleDurationPresentFlag() bool {
	return b.defaultSampleDurationPresentFlag
}

func (b TfhdBox) DefaultSampleSizePresentFlag() bool {
	return b.defaultSampleSizePresentFlag
}

func (b TfhdBox) DefaultSampleFlagsPresentFlag() bool {
	return b.defaultSampleFlagsPresentFlag
}

func (b TfhdBox) DurationIsEmptyFlag() bool {
	return b.durationIsEmptyFlag
}

func (b TfhdBox) DefaultBaseIsMoofFlag() bool {
	return b.defaultBaseIsMoofFlag
}

func (b TfhdBox) BaseDataOffset() uint64 {
	return b.baseDataOffset
}

func (b TfhdBox) SampleDescriptionIndex() uint32 {
	return b.sampleDescriptionIndex
}

func (b TfhdBox) DefaultSampleDuration() uint32 {
	return b.defaultSampleDuration
}

func (b TfhdBox) DefaultSampleSize() uint32 {
	return b.defaultSampleSize
}

func (b TfhdBox) DefaultSampleFlags() uint32 {
	return b.defaultSampleFlags
}

func (b *TfhdBox) Parse() {
	buffer := b.ParseData()
	if buffer == nil {
		return
	}
	b.version = buffer[0]
	b.flags = binary.BigEndian.Uint32(buffer[0:4])
	b.trackID = binary.BigEndian.Uint32(buffer[4:8])
	// Extract Flags
	flagbyte1 := uint8(b.flags)
	//flagbyte2 := uint8(b.flags>>8)
	flagbyte3 := uint8(b.flags >> 16)

	b.baseDataOffsetPresentFlag = TestBitMask(flagbyte1, uint8(BaseDataOffsetPresent))
	b.sampleDescriptionIndexPresentFlag = TestBitMask(flagbyte1, uint8(SampleDescriptionIndexPresent))
	b.defaultSampleDurationPresentFlag = TestBitMask(flagbyte1, uint8(DefaultSampleDurationPresent))
	b.defaultSampleSizePresentFlag = TestBitMask(flagbyte1, uint8(DefaultSampleSizePresent))
	b.defaultSampleFlagsPresentFlag = TestBitMask(flagbyte1, uint8(DefaultSampleFlagsPresent))
	b.durationIsEmptyFlag = TestBitMask(flagbyte3, uint8(DurationIsEmpty>>16))
	b.defaultBaseIsMoofFlag = TestBitMask(flagbyte3, uint8(DefaultBaseIsMoof>>16))

	// Conditional support for Optional Fields
	ibyte := 8
	if b.baseDataOffsetPresentFlag {
		b.baseDataOffset = u64(buffer[ibyte : ibyte+8])
		ibyte += 8
	}
	if b.sampleDescriptionIndexPresentFlag {
		b.sampleDescriptionIndex = u32(buffer[ibyte : ibyte+4])
		ibyte += 4
	}
	if b.defaultSampleDurationPresentFlag {
		b.defaultSampleDuration = u32(buffer[ibyte : ibyte+4])
		ibyte += 4
	}
	if b.defaultSampleSizePresentFlag {
		b.defaultSampleSize = u32(buffer[ibyte : ibyte+4])
		ibyte += 4
	}
	if b.defaultSampleFlagsPresentFlag {
		b.defaultSampleFlags = u32(buffer[ibyte : ibyte+4])
		ibyte += 4
	}

}
