package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/yokogawa-k/animal-names-generator"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(animalnamesgenerator.GetRandomName(0))
}
