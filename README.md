# scorekeeper
TDD example of a library that talks to Redis to keep a score

### Setup
Run redis docker container
```
docker run --name some-redis -d redis
```

You can stop it later with
```
docker stop some-redis
```
and you can delete `some-redis` with
```
docker rm some-redis
```
