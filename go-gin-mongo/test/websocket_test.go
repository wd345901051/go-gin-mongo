package test

import (
	"flag"
	"log"
	"net/http"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{}

var ws sync.Map

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
	}
	defer c.Close()
	ws.Store(c, struct{}{})
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv:%s", message)
		ws.Range(func(key, value any) bool {
			err = value.(*websocket.Conn).WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				return false
			}
			return true
		})

	}
}

func TestServer(t *testing.T) {
	http.HandleFunc("/echo", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func TestGinwebsocketServer(t *testing.T) {
	r := gin.Default()
	r.GET("/echo", func(ctx *gin.Context) {
		echo(ctx.Writer, ctx.Request)
	})
	r.Run(":8080")
}
