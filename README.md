# Bech32 Convert

Small go utility to convert cosmos bech32 addresses to other chains.

## Installation

```sh
go install
```

## Usage

You can convert either convert an individual address or a file containing a list of addresses.

### Convert an address

```sh
bech32-convert address chain
```

Example:

```sh
$ bech32-convert cosmos1z6mj02l2s8v0vsxfsark5v7t076ds8pu309hza stars
cosmos1z6mj02l2s8v0vsxfsark5v7t076ds8pu309hza,stars1z6mj02l2s8v0vsxfsark5v7t076ds8pu9nj2fv
```

### Convert a file of addresses

You can convert any text file containing a list of addresses:

```sh
bech32-convert addrs_file chain
```
