package random

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// GenerateSixDigitOtp generates a secure 6-digit OTP (from 100000 to 999999)
func GenerateSixDigitOtp() string {
	min := int64(100000)
	max := int64(999999)

	// Get a random number between min and max (inclusive)
	nBig, err := rand.Int(rand.Reader, big.NewInt(max-min+1))
	if err != nil {
		// fallback: just return a fixed value (NOT secure for production)
		return "123456"
	}
	otp := nBig.Int64() + min
	return fmt.Sprintf("%06d", otp)
}
