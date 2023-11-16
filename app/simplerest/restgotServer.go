package simplerest

import (
	"github.com/rest-go/rest/pkg/auth"
	"github.com/rest-go/rest/pkg/server"
	"log"
	"net/http"
	"strconv"
)

var (
	DNS       = "sqlite://my.db"
	jwtSecret = "let's hack"
)

func RunRestGoAdminSqliteServer() {
	h := server.New(&server.DBConfig{URL: DNS}, server.Prefix("/admin"))
	http.Handle("/admin/", h)
	log.Fatal(http.ListenAndServe(":3001", nil))
}

func RunRestGoAuthSqliteServer(port int) {
	server := server.New(&server.DBConfig{URL: DNS}, server.EnableAuth(true))
	authHandler, err := auth.NewHandler(DNS, []byte(jwtSecret))
	if err != nil {
		log.Fatal("initialize auth error ", err)
	}
	http.Handle("/auth/", authHandler)
	middleware := auth.NewMiddleware([]byte(jwtSecret))
	http.Handle("/", middleware(server))
	log.Fatal(func() error {
		server := &http.Server{Addr: string(":" + strconv.Itoa(port)), Handler: http.Handler(nil)}
		return server.ListenAndServe()
	}())
}
