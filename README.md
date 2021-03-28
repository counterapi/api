# Counter API

<p align="center">
  <a href="https://counterapi.dev/" target="_blank">
    <img width="180" src="https://raw.githubusercontent.com/counterapi/docs/master/src/.vuepress/public/favicons/apple-icon-180x180.png" alt="logo">
  </a>
</p>

<p align="center">
    <img src="https://img.shields.io/github/workflow/status/counterapi/counter/Code%20Check" alt="Check"></a>
    <img src="https://coveralls.io/repos/github/counterapi/counter/badge.svg?branch=master" alt="Coverall"></a>
    <img src="https://goreportcard.com/badge/github.com/counterapi/counter" alt="Report"></a>
    <a href="http://pkg.go.dev/github.com/counterapi/counter"><img src="https://img.shields.io/badge/pkg.go.dev-doc-blue" alt="Doc"></a>
    <a href="https://github.com/counterapi/counter/blob/master/LICENSE"><img src="https://img.shields.io/github/license/counterapi/counter" alt="License"></a>
</p>

Go Library for Counter API.

## Requirements

* Go installed.
* Postgresql database.

## What does it do?

Free counter API for your static website or application.

## How to use it

### Local Development

```shell
docker run -ti \
  --network host \
  -e POSTGRES_HOST=localhost \
  -e POSTGRES_PORT=5432 \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_DB=counter_api \
  -e POSTGRES_PASSWORD=root \
  counter
```

## Improvements to be made

* 100% test coverage.
* Better covering for other features.

## License

[MIT](https://github.com/counterapi/counterapi/blob/master/LICENSE)
