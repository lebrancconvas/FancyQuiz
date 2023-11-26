package server

import "os"

func Start() {
	router := NewRouter()
	router.Run(":", os.Getenv("PORT"))
}
