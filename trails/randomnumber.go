package trails

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomNumber() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%04d", rand.Intn(10000))
}
