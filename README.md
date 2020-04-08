<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-2-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->
<p align="center">    
  <a href="https://www.cosmostation.io" target="_blank" rel="noopener noreferrer"><img width="400" src="https://user-images.githubusercontent.com/31615341/78533120-614f5900-7823-11ea-901a-b745880594cf.png" alt="Cosmostation logo"></a>    
</p>

<h2 align="center">
    Mintscan Explorer's Backend for Binance Chain 
</h2>

*:star: Developed / Developing by [Cosmostation](https://www.cosmostation.io/)*

## Overview

This project is sponsored by [Binance X Fellowship Program](https://binancex.dev/fellowship.html). The program supports talented developers and researchers in creating free and open-source software that would enable new innovations and businesses in the crypto community.

This repository provides backend code for [Mintscan Block Explorer for Binance Chain](https://binance.mintscan.io/), and you can find frontend code in [this repository](https://github.com/cosmostation/mintscan-binance-dex-frontend).

**_Note that this repository has just transitioned from the actively developing phase to the maintaining phase starting from the first official version `v1.0.0`. All intended functionality is implemented; however, it can always go back when there is a reason to go back._**

## Prerequisite

- Requires [Go 1.14+](https://golang.org/dl/)

- Prepare endpoints for [Binance Chain Node RPC](https://docs.binance.org/api-reference/node-rpc.html) and [API Server](https://docs.binance.org/api-reference/api-server.html)

- Prepare PostgreSQL Database

## Folder Structure

    /
    |- chain-exporter
    |- mintscan
    |- stats-exporter

#### Chain Exporter

[chain-exporter](https://github.com/cosmostation/mintscan-binance-dex-backend/chain-exporter) watches a full node of Binance Chain and export chain data into PostgreSQL database.

#### Mintscan

[mintscan](https://github.com/cosmostation/mintscan-binance-dex-backend/mintscan) provides any necesarry custom APIs.

#### Stats Exporter

[stats-exporter](https://github.com/cosmostation/mintscan-binance-dex-backend/stats-exporter) creates cron jobs to export market data to build chart history API.

## Configuration

For configuration, it uses human readable data-serialization configuration file format called [YAML](https://en.wikipedia.org/wiki/YAML).

To configure `chain-exporter` | `mintscan` | `stats-exporter`, you need to configure  `config.yaml` file in each folder. Reference `example.yaml`.

**_Note that the configuration needs to be passed in via `config.yaml` file, so make sure to change the name to `config.yaml`._**

## Install

#### Git clone this repo
```shell
git clone https://github.com/cosmostation/mintscan-binance-dex-backend.git
```

#### Build by Makefile
```shell
cd mintscan-binance-dex-backend/chain-exporter
make build

cd mintscan-binance-dex-backend/mintscan
make build

cd mintscan-binance-dex-backend/stats-exporter
make build
```

## Database 

This project uses [Golang ORM with focus on PostgreSQL features and performance](https://github.com/go-pg/pg). Once `chain-exporter` begins to run, it creates the following database tables if not exist already.

- Block
- PreCommit
- Transaction
- Validator

## Contributing

We encourage and support an active, healthy community of contributors — any contribution, improvements, and suggestions are always welcome! Details are in the [contribution guide](https://github.com/cosmostation/mintscan-binance-dex-backend/docs/CONTRIBUTING.md)

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://jaybdev.net"><img src="https://avatars1.githubusercontent.com/u/20435620?v=4" width="100px;" alt=""/><br /><sub><b>JayB</b></sub></a><br /><a href="https://github.com/cosmostation/mintscan-binance-dex-backend/commits?author=kogisin" title="Code">💻</a> <a href="https://github.com/cosmostation/mintscan-binance-dex-backend/commits?author=kogisin" title="Documentation">📖</a> <a href="#projectManagement-kogisin" title="Project Management">📆</a> <a href="https://github.com/cosmostation/mintscan-binance-dex-backend/commits?author=kogisin" title="Tests">⚠️</a> <a href="#maintenance-kogisin" title="Maintenance">🚧</a></td>
    <td align="center"><a href="http://dev.to/fly"><img src="https://avatars3.githubusercontent.com/u/31615341?v=4" width="100px;" alt=""/><br /><sub><b>fl-y</b></sub></a><br /><a href="https://github.com/cosmostation/mintscan-binance-dex-backend/commits?author=fl-y" title="Code">💻</a> <a href="#ideas-fl-y" title="Ideas, Planning, & Feedback">🤔</a></td>
    <td align="center"><a href="https://github.com/hyeryeong-lim"><img src="https://avatars1.githubusercontent.com/u/63229379?v=4" width="100px;" alt=""/><br /><sub><b>hyeryeong-lim</b></sub></a><br /><a href="#design-hyeryeong-lim" title="Design">🎨</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!

## Our Services and Community 

- [Official Website](https://www.cosmostation.io)
- [Mintscan Explorer](https://www.mintscan.io)
- [Web Wallet](https://wallet.cosmostation.io)
- [Android Wallet](https://bit.ly/2BWex9D)
- [iOS Wallet](https://apple.co/2IAM3Xm)
- [Telegram - International](https://t.me/cosmostation)

## License

Released under the [Apache 2.0 License](https://github.com/cosmostation/mintscan-binance-dex-backend/LICENSE).
