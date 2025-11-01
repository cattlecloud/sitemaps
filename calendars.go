// Copyright (c) CattleCloud LLC
// SPDX-License-Identifier: BSD-3-Clause

package sitemaps

import (
	"fmt"
	"time"
)

type CalendarDay struct {
	Year  int
	Month int
	Day   int
}

func (cd CalendarDay) ISO() string {
	return fmt.Sprintf("%d-%02d-%02d", cd.Year, cd.Month, cd.Day)
}

func DayFrom(t time.Time) CalendarDay {
	return CalendarDay{
		Year:  t.Year(),
		Month: int(t.Month()),
		Day:   t.Day(),
	}
}
