[English](README.md) | [日本語](README_ja.md)

# percentile

`percentile` calculates the percentile of the entered number.

## Description

`percentile` calculates and displays the percentile of a number as it is entered.

To enter a number into `percentile`, specify the file from which you want to enter the number, or enter the number from standard input.

## Usage

```bash
# Read numbers from a file
$ percentile <(seq 1 100)
p25: 25
p50: 50
p75: 75
p80: 80
p90: 90
p95: 95
p99: 99

# Read numbers from a standard input
$ seq 1 100 | percentile -
p25: 25
p50: 50
p75: 75
p80: 80
p90: 90
p95: 95
p99: 99

# Calculate the specified percentile value
$ seq 1 100 | percentile -p25,50 -
p25: 25
p50: 50

# Do not round percentile values
$ seq 1 100 | percentile -r -
p25: 25.5
p50: 50.5
p75: 75.5
p80: 80.5
p90: 90.5
p95: 95.5
p99: 99.5

# Show help
$ percentile -h
Usage of percentile:
  -p string
    	Specify percentiles (comma-separated list of integers) (default "25,50,75,80,90,95,99")
  -r	Don't Round percentile values
  -v	Show version
```

## Installation

### Developer

```bash
go install github.com/ebc-2in2crc/percentile/cmd/percentile@latest
```

### User

Download from the following url.

- [https://github.com/ebc-2in2crc/percentile/releases](https://github.com/ebc-2in2crc/percentile/releases)

## Contribution

1. Fork this repository
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Rebase your local changes against the master branch
5. Run test suite with the `make test` command and confirm that it passes
6. Run `make fmt` and `make lint`
7. Create new Pull Request

## Licence

[MIT](https://github.com/ebc-2in2crc/percentile/blob/master/LICENSE)

## Author

[ebc-2in2crc](https://github.com/ebc-2in2crc)
