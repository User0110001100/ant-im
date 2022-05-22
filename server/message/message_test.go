package message

import (
	"testing"
	"time"
)

// Util test

func isSame(m *Message, n *Message) bool {
	return m.MessageBody == n.MessageBody && m.MessageOwnerId == n.MessageOwnerId && m.MessageType == n.MessageType && m.SendTime == n.SendTime
}

func TestMessageSerelize(t *testing.T) {
	now := time.Now()
	m := Message{
		MessageType:    MessageHeartsBeats,
		MessageOwnerId: 996234439,
		MessageBody:    "hello world",
	}
	m.SetTime(now)

	jsonM, ok := m.Serelize()

	//fmt.Print(string(jsonM))

	if !ok {
		t.Error("TestMessageSerelize : m.Serelize fail!")
	}

	n := Message{}
	ok = n.Deserelize(jsonM)

	if !ok {
		t.Error("TestMessageSerelize : m.Deserelize fail!")
	}

	ok = isSame(&m, &n)

	if !ok {
		t.Error("TestMessageSerelize : Is Not Same!")
	}
}
