package handler

import (
	"encoding/json"
	"github.com/cheng1005/XxCh-Common/common/response"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Node struct {
	Data chan []byte     `json:"data"`
	Conn *websocket.Conn `json:"conn"`
}
type Message struct {
	UserId      int    `json:"user_id"`      //用户id
	DistId      int    `json:"dist_id"`      //接收者id
	Cmd         int    `json:"cmd"`          //聊天类型   1:私聊 2:群聊 3:异常
	MessageType int    `json:"message_type"` //消息类型
	Content     string `json:"content"`      //消息内容
}

var OnlineUser map[int]Node = make(map[int]Node)

func Websocket(c *gin.Context) {
	userId := int(c.GetUint("userId"))
	log.Println(userId)
	//提供一个websocket连接
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	//1. 创建用户连接
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		response.Error(c, 500, err.Error())
	}
	//2. 保存用户信息到全局变量
	node := Node{
		Data: make(chan []byte, 50),
		Conn: conn,
	}
	conn.WriteMessage(websocket.TextMessage, []byte("欢迎使用websocket"))

	OnlineUser[userId] = node
	log.Println(OnlineUser)
	//3. 收发信息
	go Read(node)
	go Send(node)
}

// Read 读取信息
func Read(node Node) {

	for {
		_, message, err := node.Conn.ReadMessage()
		if err != nil {
			return
		}

		var msg Message
		json.Unmarshal(message, &msg)
		if msg.Cmd == 1 {
			if _, ok := OnlineUser[msg.DistId]; ok {
				OnlineUser[msg.DistId].Data <- message
			}
		} else if msg.Cmd == 2 {

		} else {
			content := Message{
				Cmd:         3,
				MessageType: 1,
				Content:     "消息类型错误",
			}
			marshal, _ := json.Marshal(content)

			OnlineUser[msg.UserId].Data <- marshal
		}
	}
}

// Send 发送信息
func Send(node Node) {

	for {
		v, ok := <-node.Data
		if ok {
			node.Conn.WriteMessage(websocket.TextMessage, v)
		}
	}
}
