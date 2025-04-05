package mp4streamparse

import (
	"encoding/binary"
	"fmt"
	"strings"
)

// Flag Constants
// FirstSampleFlagsPresent: this over-rides the default flags for the first sample only.
//
//	This makes it possible ot record a group of frames where the first is a key and
//	the rest are difference frames, without supplying explicit flags for every sample.
//	If this flag and field are used, sample-flags shall not be present.
//
// SampleDurationPresent: indicates that each sample has its own duration, otherwise the
//
//	default is used.
//
// SampleSizePresent: each sample has its own size, otherwise the default is used.
// SampleFlagsPresent: each sample has its own flags, otherwise the default is used.
// SampleCompositionTimeOffsetsPresent: each sample has a composition time offset (e.g. as
//
//	used for I/P/B video in MPEG)
const (
	DataOffsetPresent                   uint32 = 0x00000001
	FirstSampleFlagsPresent             uint32 = 0x00000004
	SampleDurationPresent               uint32 = 0x00000100
	SampleSizePresent                   uint32 = 0x00000200
	SampleFlagsPresent                  uint32 = 0x00000400
	SampleCompositionTimeOffsetsPresent uint32 = 0x00000800
)

// TrunSample defines sample for trun atom
type TrunSample struct {
	sampleDuration              uint32
	sampleSize                  uint32
	sampleFlags                 uint32
	sampleCompositionTimeOffset uint32
}

func (b TrunSample) String() string {
	indentLevel := 3
	tab := strings.Repeat("\t", indentLevel)
	strMsg := fmt.Sprintf("%s<< TrunSample >> Duration(%d) Size(%d) Flags(%d) CompositionTimeOffset(%d)",
		tab, b.sampleDuration, b.sampleSize, b.sampleFlags, b.sampleCompositionTimeOffset)
	return strMsg
}

func (b TrunSample) SampleDuration() uint32 {
	return b.sampleDuration
}

func (b TrunSample) SampleSize() uint32 {
	return b.sampleSize
}

func (b TrunSample) SampleFlags() uint32 {
	return b.sampleFlags
}

func (b TrunSample) SampleCompositionTimeOffset() uint32 {
	return b.sampleCompositionTimeOffset
}

// TrunBox defines the tfhd box structure.
// If the data-offset is not present, the the data for this run starts immediately
// after the data of the previous run, or at the base-data-offset defined by the
// track fragment header if this is the first run in a track fragment.
// If the data-offset is present, it is relative to the base-data-offset established
// in the track fragment header.
type TrunBox struct {
	Box
	version     byte
	flags       uint32
	sampleCount uint32
	// Optional Fields
	dataOffsetPresentFlag                   bool
	firstSampleFlagsPresentFlag             bool
	sampleDurationPresentFlag               bool
	sampleSizePresentFlag                   bool
	sampleFlagsPresentFlag                  bool
	sampleCompositionTimeOffsetsPresentFlag bool
	dataOffset                              int32
	// First Sample Flags
	isLeadingFlag uint8
	// 0: Unknown
	// 1: Leading with dependency before the referenced I-picture and is not decodable
	// 2: Not a leading sample
	// 3: Leading with no dependency before the referenced I-picture and decodable
	sampleDependsOnFlag uint8
	// 0: Unknown
	// 1: Depends upon others - not an I-picture
	// 2: Does not depend upon others - I-picture
	// 3: Reserved
	sampleIsDependedOnFlag uint8
	// 0: Unknown
	// 1: Other samples may depend upon this one (not disposable)
	// 2: No other samples depend upon this one (disposable)
	// 3: Reserved
	sampleHasRedundancyFlag uint8
	// 0: Unknown
	// 1: Redundant coding in this sample
	// 2: No redundant coding in this sample
	// 3: Reserved
	samplePaddingValueFlag    bool
	sampleIsNonSyncSampleFlag bool
	sampleArray               []TrunSample
}

func NewTrunBox(name string, size uint32, buf []byte) *TrunBox {
	return &TrunBox{Box: NewBox(name, int64(size), buf)}
}

func (b TrunBox) String() string {
	indentLevel := 3
	tab := strings.Repeat("\t", indentLevel)
	strMsg := fmt.Sprintf("%s<< TrunBox >>", tab)
	strMsg = fmt.Sprint("%s\n%s\tSampleCount(%d) DataOffset(%d)", strMsg, tab, b.sampleCount, b.dataOffset)
	strMsg = fmt.Sprintf("%s\n%s\tSampleArray(%v)", strMsg, tab, b.sampleArray)
	return strMsg
}

func (b TrunBox) Version() byte {
	return b.version
}

func (b TrunBox) Flags() uint32 {
	return b.flags
}

func (b TrunBox) SampleCount() uint32 {
	return b.sampleCount
}

func (b TrunBox) DataOffsetPresentFlag() bool {
	return b.dataOffsetPresentFlag
}

func (b TrunBox) FirstSampleFlagsPresentFlag() bool {
	return b.firstSampleFlagsPresentFlag
}

func (b TrunBox) SampleDurationPresentFlag() bool {
	return b.sampleDurationPresentFlag
}

func (b TrunBox) SampleSizePresentFlag() bool {
	return b.sampleSizePresentFlag
}

func (b TrunBox) SampleFlagsPresentFlag() bool {
	return b.sampleFlagsPresentFlag
}

func (b TrunBox) SampleCompositionTimeOffsetsPresentFlag() bool {
	return b.sampleCompositionTimeOffsetsPresentFlag
}

func (b TrunBox) DataOffset() int32 {
	return b.dataOffset
}

func (b TrunBox) IsLeadingFlag() uint8 {
	return b.isLeadingFlag
}

func (b TrunBox) SampleDependsOnFlag() uint8 {
	return b.sampleDependsOnFlag
}

func (b TrunBox) SampleIsDependedOnFlag() uint8 {
	return b.sampleIsDependedOnFlag
}

func (b TrunBox) SampleHasRedundancyFlag() uint8 {
	return b.sampleHasRedundancyFlag
}

func (b TrunBox) SamplePaddingValueFlag() bool {
	return b.samplePaddingValueFlag
}

func (b TrunBox) SampleIsNonSyncSampleFlag() bool {
	return b.sampleIsNonSyncSampleFlag
}

func (b TrunBox) SampleArray() []TrunSample {
	return b.sampleArray
}

func (b *TrunBox) Parse() {
	buffer := b.ParseData()
	if buffer == nil {
		return
	}
	b.version = buffer[0]
	b.flags = u32(buffer[0:4])
	b.sampleCount = u32(buffer[4:8])

	// Extract Flags
	flagbyte1 := uint8(b.flags)
	flagbyte2 := uint8(b.flags >> 8)
	b.dataOffsetPresentFlag = TestBitMask(flagbyte1, uint8(DataOffsetPresent))
	b.firstSampleFlagsPresentFlag = TestBitMask(flagbyte1, uint8(FirstSampleFlagsPresent))
	b.sampleDurationPresentFlag = TestBitMask(flagbyte2, uint8(SampleDurationPresent>>8))
	b.sampleSizePresentFlag = TestBitMask(flagbyte2, uint8(SampleSizePresent>>8))
	b.sampleFlagsPresentFlag = TestBitMask(flagbyte2, uint8(SampleFlagsPresent>>8))
	b.sampleCompositionTimeOffsetsPresentFlag = TestBitMask(flagbyte2, uint8(SampleCompositionTimeOffsetsPresent>>8))

	ibyte := 8
	// Populate Optional Fields
	if b.dataOffsetPresentFlag {
		x, _ := binary.Varint(buffer[ibyte : ibyte+4])
		b.dataOffset = int32(x)
		ibyte += 4
	}

	if b.firstSampleFlagsPresentFlag {
		// Extract First Sample Flags (Big Endian)
		flagbyte1 := uint8(buffer[ibyte])
		flagbyte2 := uint8(buffer[ibyte+1])

		// SKIP: Four padding bits
		// is_leading Mask: 0000 1100
		// 0: Unknown
		// 1: Leading with dependency before the referenced I-picture and is not decodable
		// 2: Not a leading sample
		// 3: Leading with no dependency before the referenced I-picture and decodable
		b.isLeadingFlag = flagbyte1 & 0x0C

		// sample_depends_on Mask:0000 0011
		// 0: Unknown
		// 1: Depends upon others - not an I-picture
		// 2: Does not depend upon others - I-picture
		// 3: Reserved
		b.sampleDependsOnFlag = flagbyte1 & 0x03

		// sample_is_depended_on Mask: 1100 0000
		// 0: Unknown
		// 1: Other samples may depend upon this one (not disposable)
		// 2: No other samples depend upon this one (disposable)
		// 3: Reserved
		b.sampleIsDependedOnFlag = flagbyte2 & 0xC0

		// sample_has_redundancy Mask: 0011 0000
		// 0: Unknown
		// 1: Redundant coding in this sample
		// 2: No redundant coding in this sample
		// 3: Reserved
		b.sampleHasRedundancyFlag = flagbyte2 & 0x30

		// sample_padding_value (not implemented)
		//b.samplePaddingValueFlag =

		// sample_is_non_sync_sample
		b.sampleIsNonSyncSampleFlag = TestBitMask(flagbyte2, uint8(0x01))
		ibyte += 4
	}

	if b.sampleCount == 0 {
		return
	}
	b.sampleArray = make([]TrunSample, b.sampleCount)
	for i := 0; i < int(b.sampleCount); i++ {
		if b.sampleDurationPresentFlag {
			b.sampleArray[i].sampleDuration = u32(buffer[ibyte : ibyte+4])
			ibyte += 4
		}
		if b.sampleSizePresentFlag {
			b.sampleArray[i].sampleSize = u32(buffer[ibyte : ibyte+4])
			ibyte += 4
		}
		if b.sampleFlagsPresentFlag {
			b.sampleArray[i].sampleFlags = u32(buffer[ibyte : ibyte+4])
			ibyte += 4
		}
		if b.sampleCompositionTimeOffsetsPresentFlag {
			b.sampleArray[i].sampleCompositionTimeOffset = u32(buffer[ibyte : ibyte+4])
			ibyte += 4
		}
	}
}
