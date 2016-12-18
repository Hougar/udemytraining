package main

import "fmt"

func main() {
	for i := 65; i <= 122; i++ {
		fmt.Println(i, " - ", string(i), " - ", i%12)
	}
}

// A is in bucket 5 since 65%12 = 5, 12 x 5 = 60 - 5 remainder from 60 - 65, 12 x 6 = 72 over price is wrong bitch
// 65  -  A  -  5
// 66  -  B  -  6
// 67  -  C  -  7
// 68  -  D  -  8
// 69  -  E  -  9
// 70  -  F  -  10
// 71  -  G  -  11
// 72  -  H  -  0
// 73  -  I  -  1
// 74  -  J  -  2
// 75  -  K  -  3
// 76  -  L  -  4
// 77  -  M  -  5
// 78  -  N  -  6
// 79  -  O  -  7
// 80  -  P  -  8
// 81  -  Q  -  9
// 82  -  R  -  10
// 83  -  S  -  11
// 84  -  T  -  0
// 85  -  U  -  1
// 86  -  V  -  2
// 87  -  W  -  3
// 88  -  X  -  4
// 89  -  Y  -  5
// 90  -  Z  -  6
// 91  -  [  -  7
// 92  -  \  -  8
// 93  -  ]  -  9
// 94  -  ^  -  10
// 95  -  _  -  11
// 96  -  `  -  0
// 97  -  a  -  1
// 98  -  b  -  2
// 99  -  c  -  3
// 100  -  d  -  4
// 101  -  e  -  5
// 102  -  f  -  6
// 103  -  g  -  7
// 104  -  h  -  8
// 105  -  i  -  9
// 106  -  j  -  10
// 107  -  k  -  11
// 108  -  l  -  0
// 109  -  m  -  1
// 110  -  n  -  2
// 111  -  o  -  3
// 112  -  p  -  4
// 113  -  q  -  5
// 114  -  r  -  6
// 115  -  s  -  7
// 116  -  t  -  8
// 117  -  u  -  9
// 118  -  v  -  10
// 119  -  w  -  11
// 120  -  x  -  0
// 121  -  y  -  1
// 122  -  z  -  2
