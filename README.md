# sitemaps

[![Go Reference](https://pkg.go.dev/badge/cattlecloud.net/go/sitemaps.svg)](https://pkg.go.dev/cattlecloud.net/go/sitemaps)
[![License](https://img.shields.io/github/license/cattlecloud/sitemaps?color=7C00D8&style=flat-square&label=License)](https://github.com/cattlecloud/sitemaps/blob/main/LICENSE)
[![Build](https://img.shields.io/github/actions/workflow/status/cattlecloud/sitemaps/ci.yaml?style=flat-square&color=0FAA07&label=Tests)](https://github.com/cattlecloud/sitemaps/actions/workflows/ci.yaml)

`sitemaps` is a Go library for generating sitemap.xml files

### Getting Started

The `sitemaps` package can be added to a Go project with `go get`.

```shell
go get cattlecloud.net/go/sitemaps@latest
```

### Examples

```go
sitemap := make(sitemaps.Site, 0, 1)
sitemap = append(sitemap, &sitemaps.URL{
  Location:  "/",
  Modified:  sitemaps.DayFrom(time.Now()),
  Frequency: sitemaps.Weekly,
  Priority:  sitemaps.Low,
})

// w is io.Writer (e.g. http.ResponseWriter)
_ = sitemap.Write(w)
```

### License

The `cattlecloud.net/go/sitemaps` module is open source under the [BSD-3-Clause](LICENSE) license.
