package main

import (
	"algosGo/utils"
	"testing"
)

func TestFindLinearSearch(t *testing.T) {
  randomNumberArray := utils.GetRandomNumberArray()
  sortedRandomNumberArray := SelectionSort(randomNumberArray)

  needleIndex := len(sortedRandomNumberArray) / 2

  foundNeedleIndex, isNeedleFound := FindLinearSearch(sortedRandomNumberArray[needleIndex], sortedRandomNumberArray)

  if (!isNeedleFound || foundNeedleIndex != needleIndex) {
    t.Fatalf("Linear search algo not found seeked needle")
  }

  _, isNeedleFound = FindLinearSearch(-1, make([]int, 0))

  if (isNeedleFound) {
    t.Fatalf("Linear search finds needle in empty array")
  }
}
