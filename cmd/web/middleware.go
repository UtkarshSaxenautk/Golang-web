package main

import (
	"github.com/justinas/nosurf"
	"log"
	"net/http"
)

func Writetoconsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got a Page")
		next.ServeHTTP(w, r)
	})
}

func Nosurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler

}

func Sessionload(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
