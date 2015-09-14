pushovercli [![Build Status](https://drone.io/github.com/blang/pushcli/status.png)](https://drone.io/github.com/blang/pushcli/latest)
======

pushcli is a commandline tool for the [pushover.net](https://pushover.net) api.

Installation
-----
Get one of the pre-compiled binaries from the releases pages.

Build
-----

Requirements: [glide](https://github.com/Masterminds/glide)

```bash
$ git clone https://github.com/blang/pushover-tools
$ cd pushover-tools/pushcli
$ export GO15VENDOREXPERIMENT=1
$ glide install
$ go build
```
