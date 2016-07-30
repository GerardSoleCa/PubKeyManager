# PubKeyManager
[![Build Status](https://travis-ci.org/GerardSoleCa/PubKeyManager.svg?branch=master)](https://travis-ci.org/GerardSoleCa/PubKeyManager) [![Go Report Card](https://goreportcard.com/badge/github.com/GerardSoleCa/PubKeyManager)](https://goreportcard.com/report/github.com/GerardSoleCa/PubKeyManager)

PubKeyManger **WILL** provide a simple way to manage public keys for SSH Connections on multiple servers by using *AuthorizedKeysCommand*.

## Disclaimer

**Still under active development. Stay tuned for changes and releases**

## Requirements
Developed under *go1.6 linux/amd64* but should be compatible with other operating systems.

**Install dependencies**
```bash
# sudo apt-get install libssl-dev
```

## Development

**Installation**
```bash
$ git clone https://github.com/GerardSoleCa/PubKeyManager
$ cd PublicKeyManager
$ go get .../.
```
**Run**
```bash
$ go run main.go
```

## License
[MIT Licensed](https://github.com/GerardSoleCa/PubKeyManager/blob/master/LICENSE.md).