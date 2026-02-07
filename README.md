# velog

[![Tests](https://github.com/NoOl01/velog/actions/workflows/go.yml/badge.svg)](https://github.com/NoOl01/velog/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/NoOl01/velog)](https://goreportcard.com/report/github.com/NoOl01/velog)
[![Go Reference](https://pkg.go.dev/badge/github.com/NoOl01/velog.svg)](https://pkg.go.dev/github.com/NoOl01/velog)

Lightweight and fast logger for Go with flexible formatting support

## Features

- Flexible configuration of the log output format
- High productivity and speed of work
- Support for multiple levels of logs (Info, Warn, Error, Debug, Fatal)
- Without third-party dependencies

## Installation

```
go get github.com/nool01/velog
```

## Quick start

### Simple example

```
package main

import (
    "github.com/nool01/velog/pkg/velog"
	"github.com/nool01/velog/pkg/velog/velog_config"
)

func main() {
    config := &velog_config.Config{
        Format: "${name} ${s} ${content} ${s} ${level} ${s} ${timestamp} ${s} ${caller}",
        Separator: " | ",
    }
    log := velog.Start(config)
    defer velog.Stop()
    
    log.Info("test", "hello world")
}
```

## Benchmark

|       Operation        | Iterations |   Speed   | Memory | Allocations |
|:----------------------:|:----------:|:---------:|:------:|:-----------:|
| `BenchmarkLogToConsole |  7000000   | 150 ns/op | 0 B/op | 0 allocs/op |

## Format placeholders

- `${level}` - Log level
- `${name}` - Log name
- `${content}` - Log message
- `${timestamp}` - Time of the log call
- `${caller}` - The place where the log was called
- `${s}` - Separator