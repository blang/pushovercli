pushovercli [![wercker status](https://app.wercker.com/status/048b1e883fbb86db44ebba610c498be2/s/master "wercker status")](https://app.wercker.com/project/bykey/048b1e883fbb86db44ebba610c498be2)
======

pushcli is a commandline tool for the [pushover.net](https://pushover.net) api.

Installation
-----
Get one of the pre-compiled binaries from the releases pages.

Usage
-----
```bash
./pushovercli --help
Usage of ./pushovercli:
  -devices string
    	Notification devices comma separated
  -priority int
    	Notification priority
  -sound string
    	Notification sound
  -timestamp string
    	Notification timestamp
  -timestamp_format string
    	Notification timestamp format, e.g. Mon Jan 2 15:04:05 -0700 MST 2006
  -title string
    	Notification title
  -token string
    	Pushover application token
  -url string
    	Notification url
  -url_title string
    	Notification url title
  -user string
    	Pushover user token
```
The last argument is the send message.

Example
-----

```
PUSHOVER_TITLE="My Title" ./pushovercli -user "usertoken" -token "apptoken" "My message"
```

Environment
--------
Each flag overrides the corresponding environment flag. Therefor it's possible to pre-configure using the environment variables.
- PUSHOVER_TOKEN
- PUSHOVER_USER
- PUSHOVER_TITLE
- PUSHOVER_URL
- PUSHOVER_URL_TITLE
- PUSHOVER_DEVICES
- PUSHOVER_PRIORITY
- PUSHOVER_SOUND
- PUSHOVER_TIMESTAMP
- PUSHOVER_TIMESTAMP_FORMAT

Exit Codes
-----
- 0 Successful delivery to pushover.net
- 1 Error creating the request (e.g. parsing timestamp)
- 2 Error returned by the pushover api

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
