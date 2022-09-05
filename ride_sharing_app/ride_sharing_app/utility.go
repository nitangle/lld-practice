package ride_sharing_app

import "math/rand"

func GenerateId() uint64 {
	return rand.Uint64() + 1
}
