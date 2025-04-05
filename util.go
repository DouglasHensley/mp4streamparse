package mp4streamparse

import (
	"encoding/binary"
	"fmt"
	"math"
)

// GetDurationString Helper function to print a duration value in the form H:MM:SS.MS
func GetDurationString(duration uint32, timescale uint32) string {
	durationSec := float64(duration) / float64(timescale)

	hours := math.Floor(durationSec / 3600)
	durationSec -= hours * 3600

	minutes := math.Floor(durationSec / 60)
	durationSec -= minutes * 60

	//sec, msec := math.Modf(durationSec)

	str := fmt.Sprintf("%02.0f:%02.0f:%06.3f", hours, minutes, durationSec)
	return str
}

// TestBitMask return bool result of (val & mask)
func TestBitMask(val uint8, mask uint8) bool {
	rc := false
	if val&mask != 0 {
		rc = true
	}
	return rc
}

func u16(bytes []byte) uint16 {
	return binary.BigEndian.Uint16(bytes)
}

func u32(bytes []byte) uint32 {
	return binary.BigEndian.Uint32(bytes)
}

func u64(bytes []byte) uint64 {
	return binary.BigEndian.Uint64(bytes)
}

// Fixed16 is an 8.8 Fixed Point Decimal notation
type Fixed16 uint16

func (f Fixed16) String() string {
	return fmt.Sprintf("%v", uint16(f)>>8)
}

func fixed16(bytes []byte) Fixed16 {
	return Fixed16(binary.BigEndian.Uint16(bytes))
}

// Fixed32 is a 16.16 Fixed Point Decimal notation
type Fixed32 uint32

func fixed32(bytes []byte) Fixed32 {
	return Fixed32(binary.BigEndian.Uint32(bytes))
}
