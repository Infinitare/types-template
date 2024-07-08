package requests

import (
	"os"
)

type domainValues struct {
	Website string
	Api     string
}

var Domains domainValues

func SetDomain() {
	env := os.Getenv("ENV")

	if env == "production" {
		Domains.Website = "https://" // INSERT website domain
		Domains.Api = "https://"     // INSERT api domain
	} else {
		Domains.Website = "http://localhost:" // INSERT website port
		Domains.Api = "http://localhost:"     // INSERT api port
	}
}
