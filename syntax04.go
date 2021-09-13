package main

import (
	"fmt"
	"math/rand"
	"time" // seed 생성용 패키지
)

func main() {
	// seed 설정
	seed := time.Now().Unix()
	rand.Seed(seed)

	dice := rand.Intn(6) + 1
	fmt.Println(dice)
}
