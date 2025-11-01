// Copyright (c) CattleCloud LLC
// SPDX-License-Identifier: BSD-3-Clause

package sitemaps

import "time"

type CalendarDay struct {
	Year  int
	Month int
	Day   int
}

func DayFrom(t time.Time) CalendarDay {
	return CalendarDay{
		Year:  t.Year(),
		Month: int(t.Month()),
		Day:   t.Day(),
	}
}
