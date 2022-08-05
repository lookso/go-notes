package main

import "fmt"

// 观察者模式
// 观察者模式也叫发布订阅模式，定义了对象之间一对多依赖,
// 当一个对象改变状态时，这个对象的所有依赖者都会收到通知并按照自己的方式进行更新。
// 实际例子当观察到登录报错,既会收到邮件,又会发送短信

type Subject struct {
	observers []Observer
	context   string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

// 添加观察者或者订阅者
func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) notifyAll(context string) {
	for _, o := range s.observers {
		o.Update(s)
	}
	s.context = context
}

// Observer 观察者接口
type Observer interface {
	Update(*Subject)
}

// sms 具体逻辑
type SmsSend struct {
	name string
}

func NewSmsSend(name string) *SmsSend {
	return &SmsSend{
		name: name,
	}
}
func (r *SmsSend) Update(s *Subject) {
	fmt.Printf("%s %s send success,account unsafe\n", r.name, s.context)
}

// mail 具体逻辑
type MailSend struct {
	name string
}

func NewMailSend(name string) *MailSend {
	return &MailSend{
		name: name,
	}
}
func (r *MailSend) Update(s *Subject) {
	fmt.Printf("%s %s send success,account unsafe\n", r.name, s.context)
}

func main() {
	subject := NewSubject()
	mailSend := NewMailSend("mail")
	smsSend := NewSmsSend("sms")
	subject.Attach(mailSend)
	subject.Attach(smsSend)

	subject.notifyAll("observer mode")
	for i := 1; i <= 10; i++ {
		subject.notifyAll(fmt.Sprintf("更新了%d", i))
		fmt.Println("+++++++++++++++++++++++++++++++++")
	}
}
