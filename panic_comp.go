package main

func safeValue(vals []int, index int) int {
  defer func() {
    if err := recover(); err := nil {
      fmt.Printf("ERROR: %s\n", err)
    }
  }
}
