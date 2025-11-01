// Copyright (c) CattleCloud LLC
// SPDX-License-Identifier: BSD-3-Clause

package sitemaps

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/shoenig/test/must"
)

func TestSite_Write(t *testing.T) {
	t.Parallel()

	unix := time.Date(2025, 10, 31, 9, 15, 0, 0, time.UTC).Unix()
	unix2 := time.Date(2025, 2, 21, 16, 30, 0, 0, time.UTC).Unix()

	site := make(Site, 0, 2)
	site = append(site, &URL{
		Location:  "/foo/bar",
		Modified:  EncodeUnix(unix),
		Frequency: Daily,
		Priority:  Bump,
	})
	site = append(site, &URL{
		Location:  "/other",
		Modified:  EncodeUnix(unix2),
		Frequency: Weekly,
		Priority:  None,
	})

	buf := new(bytes.Buffer)
	err := site.Write(buf)
	must.NoError(t, err)

	exp := strings.TrimSpace(`
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
    <url>
      <loc>/foo/bar</loc>
      <lastmod>2025-10-31</lastmod>
      <changefreq>daily</changefreq>
      <priority>0.6</priority>
    </url>
    <url>
      <loc>/other</loc>
      <lastmod>2025-02-21</lastmod>
      <changefreq>weekly</changefreq>
      <priority>0.1</priority>
    </url>
</urlset>`)

	result := strings.TrimSpace(buf.String())
	must.Eq(t, exp, result)
}
