# AES

# Table of Contents

- [Overview](#overview)
- [Installing](#installing)
- [Usage](#usage)
  - [GenerateKey](#generatekey)
  - [Encrypt](#encrypt)
  - [Decrypt](#decrypt)

# Overview

AES is a command-line tool based on the AES cbc encryption algorithm.

Based on the well-known Cobra library.

# Installing

# Usage

## GenerateKey

By default, a key file is generate in the current directory, the default file name like 20060102150405.key.

```shell script
aes generate [/full/path/secret.key]
aes gen
aes ge
aes g
```

## Encrypt

```shell script
aes e ~/secret.key /data/photos/1.png -o /data/photos/1.encrypt
aes en secret.key 1.png
```

## Decrypt

```shell script
aes d ~/secret.key /data/photos/1.png.encrypt -o /data/photos/1.png
aes de secret.key 1.png.encrypt
```
