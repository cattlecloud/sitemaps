// Copyright (c) CattleCloud LLC
// SPDX-License-Identifier: BSD-3-Clause

package sitemaps

import (
	"testing"
	"time"

	"github.com/shoenig/test/must"
)

func Test_DayFrom(t *testing.T) {
	t.Parallel()

	now := time.Date(2025, 11, 1, 6, 15, 0, 0, time.UTC)
	cd := DayFrom(now)
	must.Eq(t, 2025, cd.Year)
	must.Eq(t, 11, cd.Month)
	must.Eq(t, 1, cd.Day)
}
