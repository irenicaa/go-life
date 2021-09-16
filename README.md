# go-life

[![GoDoc](https://godoc.org/github.com/irenicaa/go-life?status.svg)](https://godoc.org/github.com/irenicaa/go-life)
[![Go Report Card](https://goreportcard.com/badge/github.com/irenicaa/go-life)](https://goreportcard.com/report/github.com/irenicaa/go-life)
[![Build Status](https://app.travis-ci.com/irenicaa/go-life.svg?branch=master)](https://app.travis-ci.com/irenicaa/go-life)
[![codecov](https://codecov.io/gh/irenicaa/go-life/branch/master/graph/badge.svg)](https://codecov.io/gh/irenicaa/go-life)

[Conway's Game of Life](https://en.wikipedia.org/wiki/Conway's_Game_of_Life).

## Installation

```
$ go get github.com/irenicaa/go-life/...
```

## Usage

```
$ go-life -h | -help | --help
$ go-life [options]
```

Stdin: grid in the [plaintext](https://www.conwaylife.com/wiki/Plaintext) format.

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit;
- `-grid-x INTEGER` &mdash; grid shift along the x-axis (default: `0`);
- `-grid-y INTEGER` &mdash; grid shift along the y-axis (default: `0`);
- `-x-min INTEGER` &mdash; x-coordinate of a top-left point of a viewport (default: `0`);
- `-x-max INTEGER` &mdash; x-coordinate of a bottom-right point of a viewport (default: `79`);
- `-y-min INTEGER` &mdash; y-coordinate of a top-left point of a viewport (default: `0`);
- `-y-max INTEGER` &mdash; y-coordinate of a bottom-right point of a viewport (default: `23`);
- `-outDelay DURATION` &mdash; delay between frames (e.g. `72h3m0.5s`; default: `100ms`).

## Output Example

```
........................O.............
......................O.O.............
............OO......OO............OO..
...........O...O....OO............OO..
OO........O.....O...OO................
OO........O...O.OO....O.O.............
..........O.....O.......O.............
...........O...O......................
............OO........................
.......................O..............
........................OO............
.......................OO.............
......................................
......................................
......................................
......................................
......................................
..............................O.O.....
...............................OO.....
...............................O......
```

## License

The MIT License (MIT)

Copyright &copy; 2020 irenica
