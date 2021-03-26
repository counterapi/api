# counter

Counter Go application

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