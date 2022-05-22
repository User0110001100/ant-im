package chat

import (
	"Ant/server/message"
	"sync"
	"time"
)

// 需要保存在Server中
type chatClients struct {
	Clients map[string]User
}

type ChatServer struct {
	mu      sync.Mutex
	clients chatClients
	Stop    bool
}

func (s *ChatServer) closeClient(u *User) {

}

func (s *ChatServer) handleError(e error, u *User) {
	s.closeClient(u)
}

func (s *ChatServer) checkMessage(m *message.Message) bool {
	checkOk := false

	return checkOk
}

// 注册消息处理
func (s *ChatServer) registerUser(m *message.Message) bool {
	registerOk := false

	// push to db

	return registerOk
}

// 登陆消息
func (s *ChatServer) userlogIn(m *message.Message) bool {
	loginOk := false

	return loginOk
}

// 下线消息
func (s *ChatServer) userOffline(m *message.Message) bool {
	offlineOk := true

	return offlineOk
}

// 接收到交互消息
func (s *ChatServer) recvData(m *message.Message) bool {
	sendOk := false

	return sendOk
}

func (s *ChatServer) sendMessageToUser(u *User, m *message.Message) error {

}

// 负责定时检查客户是否在线，发送心跳来检查
func (s *ChatServer) checkAlive() {
	for n, u := range s.clients.Clients {

		m := message.Message{
			message.MessageHeartsBeats,
			ServerId,
			Non,
			"",
			"",
			"",
			"",
		}

		go func(n *string, u *User) {
			err := s.sendMessageToUser(u, &m)

			s.mu.Lock()
			defer s.mu.UnLock()

			if err != nil {
				s.handleError(err, u)
			}

		}(&n, &u)
	}
}

func (s *ChatServer) checkAliveLoop() {
	for !s.Stop {
		time.Sleep(10 * time.Second)
		s.checkAlive()
	}
}

func MakeServer() *ChatServer {
	var s ChatServer

	s.Stop = false

	go s.checkAliveLoop()

	return &s
}
