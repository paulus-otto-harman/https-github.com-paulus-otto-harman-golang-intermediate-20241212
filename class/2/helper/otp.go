package helper

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateOTP() string {
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	otp := fmt.Sprintf("%06d", rng.Intn(1000000)) // Generate 6 digit OTP
	return otp
}
