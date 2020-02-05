# Mintscan Backend for Binance DEX 

This repository is for Mintscan Block Explorer's Backend for Binance DEX. This will be developed over the next three months and is still a work-in-progress.

## # Structure

- chain-exporter: watches a full node of Binance Chain and export data.

- mintscan: is where all custom APIs are located.


## # Install

**Note:** Requires [Go 1.13+](https://golang.org/dl/)

Both `api` and `chain-exporter` use [Tom's Obvious Minimal Language](https://github.com/toml-lang/toml) configuration file format.

Reference `example.toml` inside both `api` and `chain-exporter`

## # Configuration

The configuration needs to be passed in via `config.yaml` file inside each folder.

Example configuration file is provided in `example.yaml`

## # Database Schemas

- BlockInfo

- TransactionInfo

## # Contributing

We encourage and support an active, healthy community of contributors — any contribution, improvements, and suggestions are always welcome! Details are in the [contribution guide](https://github.com/cosmostation/mintscan-binance-dex-backend/docs/CONTRIBUTING.md)

## # Cosmostation's Services and Community 

- [Official Website](https://www.cosmostation.io)
- [Mintscan Explorer](https://www.mintscan.io)
- [Web Wallet](https://wallet.cosmostation.io)
- [Android Wallet](https://bit.ly/2BWex9D)
- [iOS Wallet](https://apple.co/2IAM3Xm)
- [Telegram - International](https://t.me/cosmostation)

## # License

Released under the [MIT License](https://github.com/cosmostation/mintscan-binance-dex-backend/LICENSE)