package sendmode

import "fmt"

//MessageSMS 具体实现类型
type MessageSMS struct {
}

//Send 实现Implementor接口
func (m *MessageSMS) Send(msg string, user string) {
	fmt.Println("使用SMS短消息发送:" + msg + "给" + user)
}

//Receive 实现Implementor接口
func (m *MessageSMS) Receive(msg string, user string) {
	fmt.Println("接收到来至" + user + "的SMS信息:" + msg)
}
