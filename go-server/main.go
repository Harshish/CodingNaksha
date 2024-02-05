package main

import (
	"fmt"
	"go-server/middleware"
	"go-server/middleware/codeforces"
	"go-server/router"
	"log"
	"net/http"
)

func registerPlatforms() {
	middleware.PlatformMap = make(map[string]middleware.ContestInterface)
	middleware.PlatformMap["codeforces"] = codeforces.Codeforces{}
}

func main() {
	r := router.Router()
	registerPlatforms()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
