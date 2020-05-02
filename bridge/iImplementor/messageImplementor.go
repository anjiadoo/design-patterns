package implementor

//MessageImplementor 定义业务实现类的接口
type MessageImplementor interface {
	Send(msg string, user string)
	Receive(msg string, user string)
}
