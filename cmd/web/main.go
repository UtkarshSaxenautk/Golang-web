package main

import (
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"

	"github.com/utkarshsaxenautk/pkg/config"
	"github.com/utkarshsaxenautk/pkg/handlers"
	"github.com/utkarshsaxenautk/pkg/render"
)

const portNumber = ":3030"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache : ", err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)
	/*http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)*/
	log.Println("Server running on : ", 3030)
	//err = http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	//if err != nil {
	//log.Fatal("Error in starting server :", err)
	//return
	//}
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	erro := srv.ListenAndServe()
	if erro != nil {
		log.Fatal("Error in starting server : ", erro)
		return
	}

}
