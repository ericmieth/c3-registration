package main

import (
	"github.com/ericmieth/c3-registration/config"
	"github.com/ericmieth/c3-registration/user"
	bv "github.com/ericmieth/c3-registration/views/backend"
	fv "github.com/ericmieth/c3-registration/views/frontend"

	"log"
	"net/http"

	"database/sql"
	_ "github.com/lib/pq"
)

func DBOpen() *sql.DB {
	configMap := config.ReturnDBParams()
	databaseConfiguration := "host=" + configMap["DBHost"] +
		" user=" + configMap["DBUser"] +
		" password=" + configMap["DBPass"] +
		" dbname=" + configMap["DBName"] +
		" sslmode=disable"

	db, err := sql.Open("postgres", databaseConfiguration)
	if err != nil {
		log.Print(err)
	}
	return db
}

type HandlerFuncDB func(http.ResponseWriter, *http.Request, *sql.DB)

func DBWrapper(f HandlerFuncDB, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		f(w, r, db)
	}
}

func main() {
	db := DBOpen()
	defer db.Close()

	http.NewServeMux()
	http.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("static/css/"))))
	http.Handle("/static/fonts/", http.StripPrefix("/static/fonts/", http.FileServer(http.Dir("static/fonts/"))))
	http.Handle("/static/js/", http.StripPrefix("/static/js/", http.FileServer(http.Dir("static/js/"))))

	// backend pages
	http.HandleFunc("/admin/auth/", DBWrapper(bv.AuthHandler, db))
	http.HandleFunc("/admin/login/", bv.LoginHandler)
	http.HandleFunc("/admin/logout/", bv.LogoutHandler)

	http.HandleFunc("/admin/userlist/", user.CheckAuth(bv.UserListHandler, db))
	http.HandleFunc("/admin/settings/", user.CheckAuth(bv.SettingsHandler, db))
	http.HandleFunc("/admin/delete/", user.CheckAuth(bv.DeleteHandler, db))
	http.HandleFunc("/admin/transfer/", user.CheckAuth(bv.TransferHandler, db))
	http.HandleFunc("/uploads/", user.CheckAuth(bv.ViewTicket, db))
	http.HandleFunc("/admin", user.CheckAuth(bv.IndexHandler, db))

	// frontend pages
	http.HandleFunc("/", DBWrapper(fv.IndexHandler, db))

	http.ListenAndServe(":9003", nil)
}
