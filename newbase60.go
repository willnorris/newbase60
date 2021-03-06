// Copyright 2014 Google Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or at
// https://developers.google.com/open-source/licenses/bsd
//
// Based on Cassis by Tantek Çelik, released under CC0 license.

// Package newbase60 implements NewBase60 encoding and decoding as specified by
// http://ttk.me/w/NewBase60.
package newbase60 // import "willnorris.com/go/newbase60"

const table = "0123456789ABCDEFGHJKLMNPQRSTUVWXYZ_abcdefghijkmnopqrstuvwxyz"

// DecodeToInt decodes the sexagesimal number s to an int.
func DecodeToInt(s string) int {
	var n int
	for _, c := range s {
		switch {
		case 48 <= c && c <= 57: // 0-9
			c -= 48
		case 65 <= c && c <= 72: // A-H
			c -= 55
		case c == 73: // capital I => number 1
			c = 1
		case 74 <= c && c <= 78: // J-N
			c -= 56
		case c == 79: // capital O => number 0
			c = 0
		case 80 <= c && c <= 90: // P-Z
			c -= 57
		case c == 95: // underscore
			c = 34
		case 97 <= c && c <= 107: // a-k
			c -= 62
		case c == 108: // lowercase l => number 1
			c = 1
		case 109 <= c && c <= 122: // m-z
			c -= 63
		default:
			c = 0 // treat all other noise as 0
		}
		n = (60 * n) + int(c)
	}
	return n
}

// EncodeInt encodes the positive integer n to a sexagesimal
// string.
func EncodeInt(n int) string {
	var s string
	for n > 0 {
		d := n % 60
		s = string(table[d]) + s
		n = (n - d) / 60
	}

	return s
}
