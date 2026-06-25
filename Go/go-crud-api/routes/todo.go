package routes

import "github.com/go-chi/chi/v5"

import "middlewares"


r :=chi.NewRouter()

r.Get("/profile", Profile)

// golbal middleware

r.Use(middlewares.TodoMiddleware)

r.Get("/todos", GetTodos)

// specific route with middleware

r.With(middlewares.TodoMiddleware).Get("/todos/{id}", GetTodoByID)