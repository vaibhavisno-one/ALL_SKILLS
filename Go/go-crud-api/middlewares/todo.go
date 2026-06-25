package middlewares

import ("fmt"
"net/http"
)


func TodoMiddleware(next http.Handler) http.Handler{
	return http.HandleFunc(func(w http.ResponseWriter, r *http.request){
		fmt.Println("Middleware executed before the handler")

		next.ServeHTTP(w,r)
	}
)
}
