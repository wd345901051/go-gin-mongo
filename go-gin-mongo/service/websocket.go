package service

import (
	"im/define"
	helper "im/helpper"
	"im/models"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

// [string]*websocket.Conn
var wc sync.Map

func WebsocketMessage(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "系统错误" + err.Error(),
		})
		return
	}
	defer conn.Close()
	uc := c.MustGet("user_claims").(helper.UserClaims)
	wc.Store(uc.Identity, conn)
	for {
		ms := new(define.MessageStruct)
		err := conn.ReadJSON(ms)
		if err != nil {
			log.Printf("Read Error:%v\n", err)
			return
		}
		// TODO: 判断用户是否属于消息体的房间
		_, err = models.GetUserRoomByUserIdentityRoomIdentity(uc.Identity, ms.RoomIdentity)
		if err != nil {
			log.Printf("UserIdentity:%v RoomIdentity:%v Not Exits\n", uc.Identity, ms.RoomIdentity)
			return
		}
		// TODO: 保存消息
		mb := &models.MessageBasic{
			UserIdentity: uc.Identity,
			RoomIdentity: ms.RoomIdentity,
			Data:         ms.Message,
			CreatedAt:    int(time.Now().Unix()),
			UpdatedAt:    int(time.Now().Unix()),
		}
		err = models.InsertOneMessageBasic(mb)
		if err != nil {
			log.Printf("DB ERROR:%v\n", err)
			return
		}
		// TODO: 获取在特定房间的在线用户
		userRooms, err := models.GetUserRoomByRoomIdentity(ms.RoomIdentity)
		if err != nil {
			log.Printf("DB ERROR :%v\n", err)
		}
		for _, room := range userRooms {
			if cc, ok := wc.Load(room.UserIdentity); ok {
				err := cc.(*websocket.Conn).WriteMessage(websocket.TextMessage, []byte(ms.Message))
				if err != nil {
					log.Printf("Write Error:%v\n", err)
					return
				}
			}
		}
	}
}
