# ランダムな名前を生成してくれるプログラム（動物編）

[![Build Status](https://travis-ci.org/yokogawa-k/animal-names-generator.svg?branch=master)](https://travis-ci.org/yokogawa-k/animal-names-generator)
[![Coverage Status](https://coveralls.io/repos/github/yokogawa-k/animal-names-generator/badge.svg?branch=master)](https://coveralls.io/github/yokogawa-k/animal-names-generator?branch=master)

### 特徴

[moby/pkg/namesgenerator](https://github.com/moby/moby/tree/master/pkg/namesgenerator)を元にしているが人名ではなく動物の名前でローマ字にしている。

### 使い方

```console
$ make build
go build -ldflags="-s -w -X main.COMMIT=70574cb" -o animal-names-generator ./cmd/...
$ ./animal-names-generator
utsukushii_kitsune
$ ./animal-names-generator
ririshii_nezumi
$ ./animal-names-generator
menkoi_ushi
```
