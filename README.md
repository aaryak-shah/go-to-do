# Go-To-Do

A basic, but complete to-do list web application written in golang.

> Learning project to understand Go, Gin and GORM

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
> /api/v1/          (GET)   ->  Online Check
    > auth/         (NIL)   <-  404
        register/   (POST)  ->  Register user
        login/      (POST)  ->  Login user
        logout/     (GET)   ->  Logout user
    create/         (POST)  ->  Create a todo list
    edit/           (PATCH) ->  Modify a todo list
    > view/         (GET)   ->  View all todo lists belonging to user
        :id/        (GET)   ->  View todo list with <id>
```