[English](README.md) | [日本語](README_ja.md)
[English](README.md) | [日本語](README_ja.md)

# percentile

`percentile` は入力された数値のパーセンタイルを計算します。

## Description

`percentile` は入力された数値のパーセンタイルを計算して表示します。

`percentile` への数値の入力は、入力元のファイルを指定するか、あるいは標準入力から入力します。

## Usage

```bash
# ファイルから数値を読み込む
$ percentile <(seq 1 100)
p25: 25
p50: 50
p75: 75
p80: 80
p90: 90
p95: 95
p99: 99

# 標準入力から数値を読み込む
$ seq 1 100 | percentile -
p25: 25
p50: 50
p75: 75
p80: 80
p90: 90
p95: 95
p99: 99

# 指定したパーセンタイル値を計算する
$ seq 1 100 | percentile -p25,50 -
p25: 25
p50: 50

# パーセンタイル値を四捨五入しない
$ seq 1 100 | percentile -r -
p25: 25.5
p50: 50.5
p75: 75.5
p80: 80.5
p90: 90.5
p95: 95.5
p99: 99.5

# ヘルプ
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

次の URL からダウンロードします。

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
