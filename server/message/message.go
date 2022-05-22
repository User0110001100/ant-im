package message

import (
	"encoding/json"
	"time"
)

// 消息的类型
const (
	// 心跳消息
	MessageHeartsBeats = 0
	// 数据交互消息
	MessageData = 1
	// 登陆消息
	MessageLogin = 2
	// 下线消息
	MessageOffline = 3
	// 注册消息
	MessageRegister = 4

	TimeType = time.ANSIC
)

var mimeMap map[string]string

func Init() {
	mimeMap[""] = "text/plain"
	mimeMap[".txt"] = "text/plain"
}

func GetMsgMimeType(m *Message) string {
	return getMimeType(getMimeSuffix(m.MessageName))
}

// TODO
func getMimeSuffix(s string) string {
	suffix := ""

	return suffix
}

func getMimeType(s string) string {
	return mimeMap[s]
}

type Message struct {
	MessageType      int    // 消息类型
	MessageOwnerId   int    // 消息发送方
	MessageSessionId int    // 消息接收的会话ID
	SendTime         string // 发送时间, 注意时间的格式是 time.ANSCI	Mon Jan 2 15:04:05 MST 2006
	Mime             string // 消息body类型
	MessageName      string // 消息名称, 当消息是文件的时候使用的
	MessageBody      string // 消息主体
}

func (m *Message) Deserelize(b []byte) bool {
	err := json.Unmarshal(b, m)
	return err == nil
}

func (m *Message) Serelize() ([]byte, bool) {
	b, err := json.Marshal(m)
	return b, err == nil
}

func (m *Message) SetOwnerId(id int) {
	m.MessageOwnerId = id
}

func (m *Message) GetOwnerId() int {
	return m.MessageOwnerId
}

func (m *Message) SetSessionId(id int) {
	m.MessageSessionId = id
}

func (m *Message) GetSessionId() int {
	return m.MessageSessionId
}

func (m *Message) SetType(t int) {
	m.MessageType = t
}

func (m *Message) GetType() int {
	return m.MessageType
}

func (m *Message) SetMime(t string) {
	m.Mime = t
}

func (m *Message) GetMime() string {
	return m.Mime
}

func (m *Message) GetMessageName() string {
	return m.MessageName
}

func (m *Message) SetMessageName(name string) {
	m.MessageName = name
}

func (m *Message) GetTime() (time.Time, bool) {
	t, err := time.Parse(time.ANSIC, m.SendTime)
	return t, err == nil
}

func (m *Message) SetTime(t time.Time) {
	m.SendTime = t.Format(time.ANSIC)
}

func (m *Message) GetBodyByByte() []byte {
	return []byte(m.MessageBody)
}

func (m *Message) GetBodyByString() string {
	return m.MessageBody
}
