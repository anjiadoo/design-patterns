package sendmode

import "fmt"

//MessageMobile 具体实现类型
type MessageMobile struct {
}

//Send 实现Implementor接口
func (m *MessageMobile) Send(msg string, user string) {
	fmt.Println("使用手机发送:" + msg + "给" + user)
}

//Receive 实现Implementor接口
func (m *MessageMobile) Receive(msg string, user string) {
	fmt.Println("接收到来至" + user + "的mobile信息:" + msg)
}
