# Counter API

[![GithubBuild](https://img.shields.io/github/workflow/status/counterapi/counterapi/Code%20Check)](http://pkg.go.dev/github.com/counterapi/counterapi)
[![Coverage Status](https://coveralls.io/repos/github/counterapi/counterapi/badge.svg?branch=master)](https://coveralls.io/github/counterapi/counterapi?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/counterapi/counterapi)](https://goreportcard.com/report/github.com/counterapi/counterapi)
[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/counterapi/counterapi)

Free counter API for your static website or application.

## Local Development

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