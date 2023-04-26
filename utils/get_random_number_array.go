package utils

import (
	"math/rand"
	"time"
)

func GetRandomNumberArray() []int {
  randomNumberArray := []int{}

  for i := 0; i < MaxArrayLength; i++ {
    rand.Seed(time.Now().UnixNano())
    randomNumberArray = append(randomNumberArray, rand.Intn(MaxArrayValue - MinArrayValue) + MinArrayValue)
  }

  return randomNumberArray
}
