package sendmode

import "fmt"

//MessageEmail 具体实现类型
type MessageEmail struct {
}

//Send 实现Implementor接口
func (m *MessageEmail) Send(msg string, user string) {
	fmt.Println("使用Email发送:" + msg + "给" + user)
}

//Receive 实现Implementor接口
func (m *MessageEmail) Receive(msg string, user string) {
	fmt.Println("接收到来至" + user + "的e-mail信息:" + msg)
}
