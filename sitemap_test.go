// Copyright (c) CattleCloud LLC
// SPDX-License-Identifier: BSD-3-Clause

package sitemaps

import (
	"testing"

	"github.com/shoenig/test/must"
)

func TestEncodeLocation(t *testing.T) {
	t.Parallel()

	cases := []struct {
		input string
		exp   string
	}{
		{
			input: "https://example.org",
			exp:   "https://example.org",
		},
		{
			input: "http://www.example.com/ümlat.php&q=name",
			exp:   "http://www.example.com/ümlat.php&amp;q=name",
		},
	}

	for _, tc := range cases {
		t.Run(tc.input, func(t *testing.T) {
			result := EncodeLocation(tc.input)
			must.Eq(t, tc.exp, result)
		})
	}
}

func TestEncodeUnix(t *testing.T) {
	t.Parallel()

	result := EncodeUnix(1750036057)
	must.Eq(t, 2025, result.Year)
	must.Eq(t, 6, result.Month)
	must.Eq(t, 15, result.Day)
}
