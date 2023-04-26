package main


func FindLinearSearch(needle int, haystack []int) (int, bool) {
  for i := 0; i < len(haystack); i++ {
    if haystack[i] == needle {
      return i, true
    }
  }

  return -1, false;
}

