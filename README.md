# sshchecker

[![Build Status](https://travis-ci.com/lazytools/sshchecker.svg?token=S9wbQbp5C4dcPWszHpyt&branch=master)](https://travis-ci.com/lazytools/sshchecker)

sshchecker is a fast dedicated ssh brute-forcing tool to check ssh login on the giving IP list.

## Install

```bash
▶ go get -v github.com/lazytools/sshchecker/cmd/sshchecker
```
## From Github

```bash
git clone https://github.com/lazytools/sshchecker.git
cd sshchecker/cmd/sshchecker
go build .
mv sshchecker /usr/local/bin/
sshchecker -h
```
## Upgrading

```bash
go get -u -v github.com/lazytools/sshchecker/cmd/sshchecker
```
## usage

```bash
▶ cat testfiles/ips.txt | sshchecker -U testfiles/testuser -P testfiles/testpass
```
## Flags
```bash
sshchecker -h
```
## License
sshchecker is coded with :heart: by lazytools. See **[Thanks.md](https://github.com/lazytools/sshchecker/blob/master/Thanks.md)** for more details.
