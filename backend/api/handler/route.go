package handler

import "net/http"

func (h *handler) RegisterRoutes() http.Handler {
	r := http.NewServeMux()

	stack := CreateStack(Logger)

	r.HandleFunc("GET /", handleFuncWrapper(h.generateSeed))
	r.HandleFunc("GET /sessions", handleFuncWrapper(h.handleListSessionByID))

	return stack(r)
}

type mw func(http.Handler) http.Handler

func CreateStack(xs ...mw) mw {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}
		return next
	}
}
