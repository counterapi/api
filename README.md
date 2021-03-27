# Counter API

[![GithubBuild](https://img.shields.io/github/workflow/status/counterapi/counter/Code%20Check)](http://pkg.go.dev/github.com/counterapi/counter)
[![Coverage Status](https://coveralls.io/repos/github/counterapi/counter/badge.svg?branch=master)](https://coveralls.io/github/counterapi/counter?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/counterapi/counter)](https://goreportcard.com/report/github.com/counterapi/counter)
[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/counterapi/counter)

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