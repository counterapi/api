# Counter API

<p align="center">
  <a href="https://counterapi.dev/" target="_blank">
    <img width="180" src="https://counterapi.dev/img/logo.png" alt="logo">
  </a>
</p>

<p align="center">
    <img src="https://img.shields.io/github/workflow/status/counterapi/api/Tests" alt="Check"></a>
    <img src="https://coveralls.io/repos/github/counterapi/api/badge.svg?branch=master" alt="Coverall"></a>
    <img src="https://goreportcard.com/badge/github.com/counterapi/api" alt="Report"></a>
    <a href="http://pkg.go.dev/github.com/counterapi/counter"><img src="https://img.shields.io/badge/pkg.go.dev-doc-blue" alt="Doc"></a>
    <a href="https://github.com/counterapi/api/blob/master/LICENSE"><img src="https://img.shields.io/github/license/counterapi/counter" alt="License"></a>
</p>

> 💬 **Have a Question or Need Help?**  
> We’d love to hear from you!  
> 👉 [Join the Community Discussions](https://community.counterapi.dev)
> 
> Ask questions, get support, or share your ideas with others.  
> Let’s build something great together! 🚀


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

[MIT](https://github.com/counterapi/api/blob/main/LICENSE)
