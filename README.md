# sshchecker

[![Build Status](https://travis-ci.com/lazytools/sshchecker.svg?token=S9wbQbp5C4dcPWszHpyt&branch=master)](https://travis-ci.com/lazytools/sshchecker)
[![License](https://img.shields.io/badge/license-MIT-_red.svg)](https://opensource.org/licenses/MIT)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/lazytools/sshchecker/issues)

sshchecker is a fast dedicated ssh brute-forcing tool to check ssh login on the giving IP list.

# Installation Instruction
### From Command Line
```bash
▶ go get -v github.com/lazytools/sshchecker/cmd/sshchecker
```
### From Github

```bash
git clone https://github.com/lazytools/sshchecker.git
cd sshchecker/cmd/sshchecker
go build .
mv sshchecker /usr/local/bin/
sshchecker -h
```
### Upgrading

```bash
go get -u -v github.com/lazytools/sshchecker/cmd/sshchecker
```
# Usage

```bash
▶ cat testfiles/ips.txt | sshchecker -U testfiles/testuser -P testfiles/testpass
```
# Flags
```bash
sshchecker -h
```
# License
sshchecker is coded with :heart: in India :india:. See **[Thanks.md](https://github.com/lazytools/sshchecker/blob/master/Thanks.md)** for more details.
