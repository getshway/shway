# Shway

Shway becomes more faster that your Mac' setup and your team's on-boarding.

This manages you and your teams

**Shway is super franken' Schway!!**

## install

default:
```bash
curl -sfL https://github.com/getshway/shway/releases/download/v0.0.1-rc/shway_v0.0.1-rc_darwin_amd64.tar.gz | tar zx -C 
```

or

build binary on your self
```bash
git clone git@github.com/go-shway/shway.git
cd shway
go build -ldflags '-w -s' -o /usr/local/bin/shway
```

## Usage

shway needs to be wrapped into your ~/.zshrc. To do that, run:

```
# ~/.zshrc
source <(shway init)
```

## Commands

### Set

```bash

```

### Update

```bash
$ shway update (project name) (git repository)
```

### List

You can list project directory names if you want to see them

```bash
$ shway list
default
reverseflash
zoom
savitar
thinker
cicada
```

## How to manage zsh configs

### 


## 優先順位

.zshrc > shway project configs > shway default configs

## Notice

Shway is open alpha.
