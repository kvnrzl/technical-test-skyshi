# Todo List API Backend Golang

This project has 3 layer :

- Repository Layer
- Service Layer
- Controller Layer

### How To Run This Project

#### Run the Testing

Go to test file in every layer and run the test

```bash
$ go test
```

#### Run the Project

The steps to run it with `docker-compose`

```bash
#clone the project
$ git clone https://github.com/kvnrzl/todo-list-api-backend-golang.git

#move to project
$ cd todo-list-api-backend-golang

#build the docker
$ docker build -t todo-list-api-backend-golang .

#run the docker
$ docker-compose up -d

#check the docker
$ docker ps

#check the logs
$ docker logs todo-list-api-backend-golang

#stop and remove the docker
$ docker-compose down
```

### API Documentation

This repository contains a router file with endpoints for managing `activity-groups` and `todo-items`.

## Endpoints

Each group has the following endpoints:

### `activity-groups`

- `POST /activity-groups`: create a new activity group
- `GET /activity-groups`: retrieve a list of all activity groups
- `GET /activity-groups/:id`: retrieve a single activity group by its ID
- `PATCH /activity-groups/:id`: update an activity group
- `DELETE /activity-groups/:id`: delete an activity group

### `todo-items`

- `POST /todo-items`: create a new todo item
- `GET /todo-items`: retrieve a list of all todo items
- `GET /todo-items/:id`: retrieve a single todo item by its ID
- `PATCH /todo-items/:id`: update a todo item
- `DELETE /todo-items/:id`: delete a todo item

Note: the `:id` parameter in the endpoints represents a dynamic value that will be replaced with a specific resource ID when the endpoint is called.

### Tools Used:

In this project, I use some tools listed below. But you can use any simmilar library that have the same purposes. But, well, different library will have different implementation type. Just be creative and use anything that you really need.

- All libraries listed in [`go.mod`](https://github.com/kvnrzl/backend-engineer-test-privy/blob/main/go.mod)
- To Generate Mocks for testing needs ["github.com/vektra/mockery"](https://github.com/vektra/mockery)
