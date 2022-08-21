package stream

const SET_BITS = 0b01
const CLEAR_BITS = 0b10
const PRESERVE_BITS = 0b11

// todo:
const STREAM_MASK = 1
const FLAG_MASK_IS = STREAM_MASK

func toStreamFlags(combOpFlags int) int {
	// By flipping the nibbles 0x11 become 0x00 and 0x01 become 0x10
	// Shift left 1 to restore set flags and mask off anything other than the set flags
	return ((^combOpFlags) >> 1) & FLAG_MASK_IS & combOpFlags
}
