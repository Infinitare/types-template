package cookies

import (
	"net/http"
	"os"
	"time"
)

func Add(w http.ResponseWriter, key string, value string) {

	expire := time.Now().Add(21900 * time.Hour)
	cookie := http.Cookie{
		Name:  key,
		Value: value,
		Domain: func() string {
			if os.Getenv("ENV") == "development" {
				return ".localhost"
			} else {
				return "." // INSERT domain here like: ".infinitare.com"
			}
		}(),
		Path:    "/",
		Expires: expire,
	}
	http.SetCookie(w, &cookie)

}
