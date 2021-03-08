package bridge

type IMsgSender interface {
	Send(msg string) error
}

//EmailMsgSender 发送邮件
//可能还有电话、短信等各种实现
type EmailMsgSender struct {
	emails []string
}

// NewEmailMsgSender NewEmailMsgSender
func NewEmailMsgSender(emails []string) *EmailMsgSender {
	return &EmailMsgSender{emails: emails}
}

//Send Send
func (s *EmailMsgSender) Send(msg string) error {
	return nil
}

//INotification 通知接口
type INotification interface {
	Notity(msg string) error
}

//ErrorNotification错误通知
//后面可能还有warning 各种级别
type ErrorNotification struct {
	sender IMsgSender
}

//NewErrorNotification NewErrorNotification
func NewErrorNotification(sender IMsgSender) *ErrorNotification {
	return &ErrorNotification{sender: sender}
}

//Notify发送通知
func (n *ErrorNotification) Notify(msg string) error {
	return n.sender.Send(msg)
}
