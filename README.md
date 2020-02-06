<p align="center">
  <a href="https://www.cosmostation.io" target="_blank" rel="noopener noreferrer"><img width="100" src="https://user-images.githubusercontent.com/20435620/55696624-d7df2e00-59f8-11e9-9126-edf9a40b11a8.png" alt="Cosmostation logo"></a>
</p>

<h2 align="center">
    Mintscan Explorer's Backend for Binance Chain 
</h2>

*:star: Developed / Developing by [Cosmostation](https://www.cosmostation.io/)*

## Overview

This project is sponsored by [Binance X Fellowship Program](https://binancex.dev/fellowship.html).

This repository provides backend code for [Mintscan Block Explorer for Binance Chain](https://binance.mintscan.io/).

- [chain-exporter](https://github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter) watches a full node of Binance Chain and export data into PostgreSQL database.

- [mintscan](https://github.com/cosmostation/mintscan-binance-dex-backend/mintscan) is where all custom APIs are located.

**_Note that this repository is currently being developed meaning that most likely there will be many breaking changes._**

## Prerequisite

- Endpoints for [Binance Chain Node RPC](https://docs.binance.org/api-reference/node-rpc.html) and [API Server](https://docs.binance.org/api-reference/api-server.html)

- PostgreSQL Database

## Install

**Note:** Requires [Go 1.13+](https://golang.org/dl/)

Git clone this repo
```shell
git clone https://github.com/cosmostation/mintscan-binance-dex-backend.git
```

Chain Exporter
```shell
cd mintscan-binance-dex-backend/chain-exporter
go run main.go
```

Mintscan API
```shell
cd mintscan-binance-dex-backend/chain-exporter
go run application.go
```

**_Makefile will be supported soon._**

## Configuration

It uses human readable data-serialization configuration file format, [YAML](https://en.wikipedia.org/wiki/YAML).

Reference `example.yaml` inside both `chain-exporter` and `mintscan`.

The configuration needs to be passed in via `config.yaml` file, so make sure to change the name to `config.yaml`.

## Database 

This project uses [Golang ORM with focus on PostgreSQL features and performance](https://github.com/go-pg/pg).

Database tables are:

- Block
- PreCommit
- Transaction
- Validator
- More to add...

## Contributing

We encourage and support an active, healthy community of contributors — any contribution, improvements, and suggestions are always welcome! Details are in the [contribution guide](https://github.com/cosmostation/mintscan-binance-dex-backend/docs/CONTRIBUTING.md)

## Our Services and Community 

- [Official Website](https://www.cosmostation.io)
- [Mintscan Explorer](https://www.mintscan.io)
- [Web Wallet](https://wallet.cosmostation.io)
- [Android Wallet](https://bit.ly/2BWex9D)
- [iOS Wallet](https://apple.co/2IAM3Xm)
- [Telegram - International](https://t.me/cosmostation)

## License

Released under the [MIT License](https://github.com/cosmostation/mintscan-binance-dex-backend/LICENSE).