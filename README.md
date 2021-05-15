<h1 align="center">
GO CLI Application Template
</h1>

<p align="center">
  <a href="https://omegion.dev" target="_blank">
    <img width="180" src="https://cdn.logo.com/hotlink-ok/logo-social-sq.png" alt="logo">
  </a>
</p>

<p align="center">
    <img src="https://img.shields.io/github/workflow/status/omegion/go-ddclient/Test" alt="Test"></a>
    <img src="https://coveralls.io/repos/github/omegion/go-ddclient/badge.svg?branch=master" alt="Coverall"></a>
    <img src="https://goreportcard.com/badge/github.com/omegion/go-ddclient" alt="Report"></a>
    <a href="http://pkg.go.dev/github.com/omegion/go-ddclient"><img src="https://img.shields.io/badge/pkg.go.dev-doc-blue" alt="Doc"></a>
    <a href="https://github.com/omegion/go-ddclient/blob/master/LICENSE"><img src="https://img.shields.io/github/license/omegion/go-ddclient" alt="License"></a>
</p>

```shell
Dynamic DNS Client CLI application to keep DNS record updated.

Usage:
  ddclient [command]

Available Commands:
  help        Help about any command
  set         Sets DNS record to current IP address.
  version     Print the version/build number

Flags:
  -h, --help   help for ddclient

Use "ddclient [command] --help" for more information about a command.
```

## Run with Docker

```shell
export CF_API_KEY=<YOUR_KEY>
docker run -e "CF_API_KEY=${CF_API_KEY}" ghcr.io/omegion/ddclient:latest \
 set \
 --record=pi-1.omegion.dev \
 --zone=omegion.dev  \
 --dns-provider=cloudflare \
 --logLevel debug
```

## Requirements

* Req 1
* Req 2

## What does it do?

A template for Go CLI application.

## How to use it

* how 1
* how 2

## Improvements to be made

* 100% test coverage.
* Better covering for other features.

