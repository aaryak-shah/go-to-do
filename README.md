# Go-To-Do

A basic, but complete to-do list web application written in golang.

> Learning project to understand Go, Gin and GORM

## How To Run

### Backend

```sh
go run server.go
```

### Frontend

```sh
cd client
npm run dev
```

## Project Structure

```sh
> go-to-do/
    > client/           # front-end code (svelte)
    > db/               # database management utils
    > initializers/     # project initialization utils
    > migrations/       # database migration utils
    > models/           # data models
    > mw/               # middleware definitions
    > routes/           # route controllers definitions
    server.go           # entrypoint
```

## API

```
> /api/v1/          (GET)   ->  Online check
    > auth/         (NIL)   <-  404
        register/   (POST)  ->  Register user
        login/      (POST)  ->  Login user
        logout/     (GET)   ->  Logout user
    create/         (POST)  ->  Create a todo list
    edit/           (PATCH) ->  Modify a todo list
    > view/         (GET)   ->  View all todo lists belonging to user
        :id/        (GET)   ->  View todo list having <id>
```

## Important Notes

### gin

- [Potential Bug] In some cases, `/foo/bar/route/` might not be considered the same as `/foo/bar/route`. Needs further investigation.\
To be kept in mind if/when a route has been defined but calls to the endpoint return in error (probably 404).

<!-- ### gorm -->

### golang-jwt

- always import `github.com/golang-jwt/jwt/v5` instead of `github.com/golang-jwt/jwt`. Both are valid imports but are incompatible with each other.

<!-- ### svelte -->