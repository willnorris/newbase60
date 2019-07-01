// Copyright 2014 Google Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file or at
// https://developers.google.com/open-source/licenses/bsd

package newbase60

import "testing"

func TestEncodeInt(t *testing.T) {
	tests := []struct {
		n int
		s string
	}{
		{0, ""},
		{60, "10"},
		{1152, "KC"},
		{-1, ""}, // negative numbers are ignored
	}

	for _, tt := range tests {
		if got := EncodeInt(tt.n); got != tt.s {
			t.Errorf("EncodeInt(%d) = %v, want %v", tt.n, got, tt.s)
		}
	}
}

func TestDecodeToInt(t *testing.T) {
	tests := []struct {
		s string
		n int
	}{
		{"", 0},
		{"10", 60},
		{"KC", 1152},
		{"2BIK", 471679},
		{"OT_bln", 357393707},
		{"!/*#", 0},    // invalid chars are ignored
		{".*KC", 1152}, // invalid chars at beginning have no effect
		{"1*", 60},     // invalid chars at end are treated as 0
	}

	for _, tt := range tests {
		if got := DecodeToInt(tt.s); got != tt.n {
			t.Errorf("DecodeToInt(%s) = %v, want %v", tt.s, got, tt.n)
		}
	}
}
