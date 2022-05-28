# iso8601_duration

[![Go Reference](https://pkg.go.dev/badge/github.com/alfonmga/iso8601_duration.svg)](https://pkg.go.dev/github.com/alfonmga/iso8601_duration)

[ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#Durations) duration Go utilities.

## Install

Via go get tool

```bash
$ go get github.com/alfonmga/iso8601_duration
```

## Usage

[Browse the docs for the available functions](https://pkg.go.dev/github.com/alfonmga/iso8601_duration#pkg-functions).

## Disclaimer

I built this because I wanted to convert YouTube ISO 8601 durations to seconds. [Check out the tests](iso8601_duration_test.go) and use it at your own risk. The current implementation may panic if an invalid ISO 8601 duration string is provided.
