// Copyright (c) CattleCloud LLC
// SPDX-License-Identifier: BSD-3-Clause

package sitemaps

import (
	"bytes"
	"encoding/xml"
	"time"
)

type Site []*URL

type Frequency string

const (
	Always  Frequency = "always"
	Hourly  Frequency = "hourly"
	Daily   Frequency = "daily"
	Weekly  Frequency = "weekly"
	Monthly Frequency = "monthly"
	Yearly  Frequency = "yearly"
	Never   Frequency = "never"
)

type Priority float64

const (
	Top    = 0.9
	High   = 0.8
	Bump   = 0.6
	Normal = 0.5
	Low    = 0.2
	None   = 0.1
)

type URL struct {
	Location  string
	Modified  CalendarDay
	Frequency Frequency
	Priority  Priority
}

func EncodeLocation(original string) string {
	b := new(bytes.Buffer)
	xml.Escape(b, []byte(original))
	return b.String()
}

func EncodeUnix(unix int64) CalendarDay {
	t := time.Unix(unix, 0).In(time.UTC)
	return DayFrom(t)
}
