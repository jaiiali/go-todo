# Todo

Todo with Hexagonal Architecture

## Installation

### Export environment variables
```shell
export $(grep -v '^#' ./.env | xargs -d '\n')
```

### Run database migrations
```shell
make migrate-up
```

### Build the project

```shell
make build
```

### Run the project

Run the project on `ip` and `port` that are sets in `.env` file:
```shell
make run
```
or
```shell
./bin/go-todo
```

## APIs

### Create Todo

Request:
```shell
curl -X POST http://0.0.0.0:8000/api/todos --data '{"title": "todo", "desc": "description"}'
```

### Todo List

Request:
```shell
curl -X GET http://0.0.0.0:8000/api/todos
```

### Get Todo

Request:
```shell
curl -X GET http://0.0.0.0:8000/api/todos/{id}
```
