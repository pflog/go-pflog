package container

var messageContainerKind = []byte("message")

func NewMessage(b []byte) *Message {
	return &Message{
		Text: NewText(b),
	}
}

type Message struct {
	*Text
}

func (_ Message) Enclosed() bool {
	return false
}

func (_ Message) Kind() []byte {
	return messageContainerKind
}
