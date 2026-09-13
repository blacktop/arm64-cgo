package instructions

// Pseudo pointer authentication shared by the SP-modifier PAC forms (system.go) and the
// FEAT_PAuth_LR forms (pauth_lr.go). The top byte of a signed pointer carries a tag derived
// from the key, the modifier, and the canonical pointer payload, so a modified pointer no
// longer authenticates. Bit 55 selects the address half, and authentication or stripping
// restores bits 63:56 from it as the architecture does, so kernel (0xffff...) and user
// (0x0000...) addresses both survive a sign/authenticate round trip.

const (
	pseudoPACKeyA uint8 = 0xA5
	pseudoPACKeyB uint8 = 0xB3

	pointerPayloadMask uint64 = 0x00FF_FFFF_FFFF_FFFF
	pointerHighHalfBit uint64 = 1 << 55
)

// canonicalPointer returns the pointer with bits 63:56 replaced by copies of bit 55.
func canonicalPointer(value uint64) uint64 {
	if value&pointerHighHalfBit != 0 {
		return value | ^pointerPayloadMask
	}
	return value & pointerPayloadMask
}

// computePseudoPAC folds a modifier value into a single byte and mixes in a key tag.
func computePseudoPAC(modifier uint64, keyTag uint8) uint8 {
	var pacByte uint8
	for shift := 0; shift < 64; shift += 8 {
		pacByte ^= uint8(modifier >> shift)
	}
	return pacByte ^ keyTag
}

// combineModifiers folds two modifier values into one pseudo-PAC modifier.
func combineModifiers(a, b uint64) uint64 {
	return a ^ (b<<17 | b>>47)
}

// pseudoPACTag derives the tag for a pointer from its canonical payload, the modifier and key.
func pseudoPACTag(value, modifier uint64, key uint8) uint8 {
	return computePseudoPAC(combineModifiers(modifier, canonicalPointer(value)), key)
}

// signPointer writes the pseudo-PAC tag into bits 63:56 of the canonical pointer.
func signPointer(value, modifier uint64, key uint8) uint64 {
	tag := pseudoPACTag(value, modifier, key)
	return (canonicalPointer(value) & pointerPayloadMask) | uint64(tag)<<56
}

// authenticatePointer returns the canonical pointer when the tag matches, and zero on a
// mismatch so a later use of the poisoned pointer is obvious.
func authenticatePointer(value, modifier uint64, key uint8) uint64 {
	if uint8(value>>56) != pseudoPACTag(value, modifier, key) {
		return 0
	}
	return canonicalPointer(value)
}
