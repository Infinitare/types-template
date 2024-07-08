package secrets

import "os"

func Load() {

	JwtSecret = os.Getenv("JWT_SECRET")

}
