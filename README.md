# Currency converter

It's simple CLI currency converter written in Go, which gets exchange rates from [exchangeratesapi.io](https://exchangeratesapi.io/) ([GitHub repository](https://github.com/exchangeratesapi/exchangeratesapi))

## Install

Installing is as simple as typing this:

```bash
go get gitlab.com/koralowiec/currconv
```

## Usage

currconv gets 3/4 arguments, where the (1) amount of money (2) in base currency (3) word *to* (3/4) target currency, examples below:

```bash
currconv 10 usd to pln
currconv 222 usd eur
currconv 333.5 USD EUR
```
