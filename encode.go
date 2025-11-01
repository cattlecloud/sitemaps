// Copyright (c) CattleCloud LLC
// SPDX-License-Identifier: BSD-3-Clause

package sitemaps

import (
	_ "embed"
	"io"
	"text/template"
)

//go:embed layout.xml
var layout string

func (s Site) Write(w io.Writer) error {
	t := template.Must(template.New("sitemap").Parse(layout))
	return t.ExecuteTemplate(w, "sitemap", s)
}
