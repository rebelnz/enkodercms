package middleware

import (
	"bitbucket.org/enkdr/enkoder/config"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"time"
)

// gets search/replaced by enkdr-script
var Store = sessions.NewCookieStore([]byte("haighaezieneidiizeiheivaepheisha"))

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}

func Auth(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, config.Conf().SessionName)
		userid, _ := session.Values["UserID"]
		if userid == nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		inner.ServeHTTP(w, r)
	})
}
