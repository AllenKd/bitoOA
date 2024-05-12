# Bito OA By AllenKao

## Build & run
```shell
$ make all
```

## Build docker image
```shell
$ docker-compose -f ./build/docker-compose.yaml build
```

## Run as docker
```shell
$ docker-compose -f ./build/docker-compose.yaml up
```

## API documentation

Once server start, visit http://localhost:8080/swagger/index.html#/

<p >
  <img src="resource/apiDocumentation.png">
</p>

## TBD

- Remove User
    - Remove the user from DB (current approach)
    - Set user date limit to 0
    - Block the user