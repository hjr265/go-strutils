// Copyright 2014 The Go-strutils Authors. All rights reserved.

// Package z provides an implementation of Z-Algorithm (adopted from http://codeforces.com/blog/entry/3107).
//
// Given a string p, Z-Algorithm produces an array Z, where Z[i] is the length of the longest substring starting from p[i] which is also a prefix of p.
package z

type Z []int

func New(p string) Z {
	z := make(Z, len(p))
	for i, l, r := 1, 0, 0; i < len(p); i++ {
		if i > r {
			l = i
			r = i
			for r < len(p) && p[r-l] == p[r] {
				r++
			}
			z[i] = r - l
			r--
		} else {
			k := i - l
			if z[k] < r-i+1 {
				z[i] = z[k]
			} else {
				l = i
				for r < len(p) && p[r-l] == p[r] {
					r++
				}
				z[i] = r - l
				r--
			}
		}
	}
	return z
}

func Contains(s, substr string) bool {
	return Index(s, substr) >= 0
}

func Index(s, substr string) int {
	if substr == "" {
		return 0
	}
	r := rune(0)
	for _, c := range substr {
		if c > r {
			r = c
		}
	}
	r += 1
	p := substr + string(r) + s
	z := New(p)
	for i := len(substr); i < len(p); i++ {
		if z[i] == len(substr) {
			return i - len(substr) - 1
		}
	}
	return -1
}
