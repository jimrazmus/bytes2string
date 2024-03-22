# bytes2string

[![Go Report Card](https://goreportcard.com/badge/github.com/jimrazmus/bytes2string)](https://goreportcard.com/report/github.com/jimrazmus/bytes2string)
[![Go Reference](https://pkg.go.dev/badge/github.com/jimrazmus/bytes2string.svg)](https://pkg.go.dev/github.com/jimrazmus/bytes2string)
[![License](http://img.shields.io/:license-mit-blue.svg?style=flat-square)](http://badges.mit-license.org)

## Summary

Go module for converting int64 byte counts to a shortened SI, or IEC, string format. For example, 1000000 to '1.0 MB' or 1048576 to '1.0 MiB'.

## Installation

`go get github.com/jimrazmus/bytes2string`

## Usage

First, import the module.

```golang
import github.com/jimrazmus/bytes2string
```

Then, use the provided functions.

```golang
log.Println(ByteCountSI(12345678))

log.Println(ByteCountIEC(12345678))
```

## Contributing

Please follow the [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) specification for your commit messages. Commit type options include: feat, fix, build, chore, ci, docs, style, refactor, perf, and test.

Open a pull request when you have completed your work and want it reviewed for inclusion.

## Author

Functions originate here: https://yourbasic.org/golang/formatting-byte-size-to-human-readable-format/

Jim Razmus II - Minor modifications and formatting, along with adding testing and packaging.

## License

This project is licensed under the MIT License - see the [LICENSE.txt](LICENSE.txt) file for details.
