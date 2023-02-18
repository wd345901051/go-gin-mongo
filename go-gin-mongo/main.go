// gin
// go get -u github.com/gin-gonic/gin
// websocket
// go get github.com/gorilla/websocket
//	mongo-driver
// go get go.mongodb.org/mongo-driver/mongo
// jwt
// go get github.com/dgrijalva/jwt-go
package main

import "im/router"

func main() {
	e := router.Router()
	e.Run(":8080")
}
