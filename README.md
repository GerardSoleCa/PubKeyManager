# PubKeyManager
[![Build Status](https://travis-ci.org/GerardSoleCa/PubKeyManager.svg?branch=master)](https://travis-ci.org/GerardSoleCa/PubKeyManager) [![Go Report Card](https://goreportcard.com/badge/github.com/GerardSoleCa/PubKeyManager)](https://goreportcard.com/report/github.com/GerardSoleCa/PubKeyManager)

PubKeyManager provides a simple way to manage public keys for SSH Connections on multiple servers. 
It makes use of **AuthorizedKeysCommand** of OpenSSH.

It makes use of sqlcipher to have an encrypted database which cannot be opened nor modified without its valid password.


## Configuration

It is possible to configure different parameters of PubKeyManager. The config file should be located along with the binary.

Store a file named **pubkeymanager.conf** where you have the binary.
```toml
port=8080
db_password="ratata"
```
## Disclaimer

On production systems configure HTTPs or your servers can be exposed to untrusted clients.

## Docker

A Docker image exists. 

```bash
$ docker run -d -p 8080:8080 gerardsoleca/pubkeymanager
``` 

## Development
Developed under *go1.6 linux/amd64*. Should work on other *nix systems.

### Install dependencies

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