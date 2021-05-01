package utils

import (
	"fmt"
	"time"
)

func ExectutionTime(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Time execution : %d\n", elapsed)

}
