# Container Image Analyser (CIA)

[![PkgGoDev](https://pkg.go.dev/badge/github.com/moabukar/cia)](https://pkg.go.dev/mod/github.com/moabukar/cia)
[![Go Report Card](https://goreportcard.com/badge/github.com/moabukar/cia)](https://goreportcard.com/report/github.com/moabukar/cia)
[![Docker Pulls](https://img.shields.io/docker/pulls/moabukar/cia.svg)](https://hub.docker.com/r/moabukar/cia/)
[![test](https://github.com/moabukar/cia/actions/workflows/test.yml/badge.svg)](https://github.com/moabukar/cia/actions/workflows/test.yml)
[![Coverage Status](https://coveralls.io/repos/github/moabukar/cia/badge.svg?branch=main)](https://coveralls.io/github/moabukar/cia?branch=main)
[![release](https://github.com/moabukar/cia/actions/workflows/release.yml/badge.svg)](https://github.com/moabukar/cia/actions/workflows/release.yml)

Container Image Analyser (CIA) is a lightweight CLI tool designed to help DevOps engineers and developers analyse Docker container images for vulnerabilities and potential issues. It simplifies the process of ensuring the security and quality of containerized applications.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Installation](#installation)
  - [Usage](#usage)
- [Contributing](#contributing)
  - [Development Setup](#development-setup)
  - [Adding New Features](#adding-new-features)
- [License](#license)

## Features

- **Vulnerability Scanning**: CIA scans Docker container images for known vulnerabilities and rates their severity.
- **Detailed Reports**: Generate detailed reports about the vulnerabilities found, including their CVE IDs, severity levels, and package information.
- **Remediation Suggestions**: Get recommendations for remediation, such as updating specific packages or using alternative base images.
- **User-Friendly CLI**: An easy-to-use command-line interface that supports a variety of options for scanning and reporting.

## 📝 Usage

```bash
~$ cia

NAME:
   cia -  CIA is your go-to CLI tool for analyzing container images. It can pull images, scan for vulnerabilities, and output reports in multiple formats.

USAGE:
   cia [global options] command [command options] [arguments...]

COMMANDS:
   scan     Scans the given container image
   report   Generate a report of the last scan
   version  Show the version of CIA tool
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --format value  Report format (json, xml) (default: "json")
   --skip-pull     Skip pulling image before scanning
   --help, -h      show help

```

### Installation

You can install CIA via `go get` or by downloading the binary release for your platform from the [Releases](https://github.com/moabukar/cia/releases) page.

```bash
go get github.com/moabukar/cia

```

## Directory structure

```bash
.
├── CHANGELOG.md
├── Dockerfile
├── LICENSE
├── Makefile
├── README.md
├── cmd
│   └── cia
│       ├── cia
│       └── main.go
├── data
├── go.mod
├── go.sum
├── internal
│   ├── cli
│   │   └── main.go
│   ├── cmd
│   │   └── main.go
│   ├── report
│   │   └── report.go
│   └── scanner
│       └── scanner.go
└── tests                    

```
