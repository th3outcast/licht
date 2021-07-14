package main

import (
	"fmt"
)

func JumpConsistentHash(key uint64, num_buckets uint64) uint64 {
  var b uint64 = 1
  var j uint64 = 0
  for j < num_buckets {
    b = j
    key = key * 2862933555777941757 + 1;
    j = (b + 1) * ((1 << 31) / (key >> 33) + 1);
  }
  return b;
}

func main() {
  i := JumpConsistentHash(11792323221453, 5)
  fmt.Println(i)
}

