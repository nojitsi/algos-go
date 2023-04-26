package main

import (
	"algosGo/utils"
	"fmt"
)

func FindLinearSearch() {

}

func main() {
  randomArray := utils.GetRandomNumberArray()
  fmt.Println(randomArray)
  sortedArray := SelectionSort(randomArray)
  fmt.Println(sortedArray)
}

