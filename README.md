<h1 align="center">
Dynamic DNS Client Tool
</h1>

<p align="center">
  <a href="https://omegion.dev" target="_blank">
    <img width="180" src="https://ssh-manager.omegion.dev/img/logo.png" alt="logo">
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
export CF_API_KEY=<YOUR_API_KEY>
docker run -e "CF_API_KEY=${CF_API_KEY}" ghcr.io/omegion/ddclient:latest \
 set \
 --record=test.example.com \
 --zone=example.com  \
 --dns-provider=cloudflare \
 --logLevel debug
```

## What does it do?

Dynamic DNS Client is a tool to update a domain based on machine internet address.

## Improvements to be made

* 100% test coverage.
* Better covering for other features.
