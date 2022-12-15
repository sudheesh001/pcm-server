package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

type BitValueRange struct {
	min int
	max int
}

func (b *BitValueRange) GetEncodedRandom(shouldPad bool) string {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(b.max-b.min+1) + b.min
	// Encode the result into a 2 character sequence.
	if shouldPad {
		return fmt.Sprintf("%02d", value)
	}
	return fmt.Sprintf("%d", value)
}

func GenerateRandom4BitValue() string {
	b := BitValueRange{min: 0, max: 15}
	return b.GetEncodedRandom(true)
}

func GenerateRandom6BitValue() string {
	b := BitValueRange{min: 0, max: 63}
	return b.GetEncodedRandom(true)
}

func GenerateRandom8BitValue() string {
	b := BitValueRange{min: 0, max: 255}
	return b.GetEncodedRandom(false)
}

func GenerateRandom16BitValue() string {
	b := BitValueRange{min: 0, max: 65535}
	return b.GetEncodedRandom(false)
}
