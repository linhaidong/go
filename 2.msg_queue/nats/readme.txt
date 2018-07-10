nats的server为gnatsd

发送消息模型：
publish:发布订阅模型
request:请求相应模型
queue:队列模型

接收消息模型：
subscribe 订阅消息
subscribesync  同步订阅，用nextmsg接收

install:
Install and run the NATS server:
    go get github.com/nats-io/gnatsd
    gnatsd
Install the Go NATS client:
    go get github.com/nats-io/go-nats

get source and build:
    1.Run go version to verify that you are running Go 1.5+. (Run go help for more guidance.)
    2.Clone the https://github.com/nats-io/gnatsd repository.
    3,Run go build inside the /nats-io/gnatsd directory. A successful build produces no messages and creates the server executable gnatsd in the directory.
    4.Run go test ./... to run the unit regression tests.
(from gnatsd 目录的readme文档)
type Msg struct {
    Subject string
    Reply   string
    Data    []byte
    Sub     *Subscription
    // contains filtered or unexported fields
}
type MsgHandler func(msg *Msg)
MsgHandler is a callback function that processes messages delivered to asynchronous subscribers.

[发送接收模型]
func (nc *Conn) Request(subj string, data []byte, timeout time.Duration) (*Msg, error)
sub: 主题
data:消息
timeout:超时时间
发送消息,并在timeout时间内,接收消息

[消息队列模型]
func (nc *Conn) QueueSubscribe(subj, queue string, cb MsgHandler) (*Subscription, error)
sub:消息的主题
queue:消息队列
cb:回调函数
QueueSubscribe creates an asynchronous queue subscriber on the given subject. 
All subscribers with the same queue name will form the queue group and only one
 member of the group will be selected to receive any given message asynchronously.
一个消息主题,可以有多个队列名称,每个队列都会接受到消息.
每个队列具有多个处理进程,但只有一个处理进程可以接收到消息.
ex:
//订阅消息队列hello
$ go run nats-qsub.go  foo hello
$ go run nats-qsub.go  foo hello
//订阅消息队列msg
$go run nats-qsub.go foo msg
$go run nats-qsub.go foo msg
//发布消息
$go run nats-pub.go foo 22222222222222222

hello ,msg 为队列的名称,都可以收到消息,但队列中只有一个进程可以收到消息

订阅消息
func (*Conn) ChanSubscribe
func (nc *Conn) ChanSubscribe(subj string, ch chan *Msg) (*Subscription, error)
ChanSubscribe will place all messages received on the channel. 
You should not close the channel until sub.Unsubscribe() has been called.


func (*Subscription) AutoUnsubscribe
func (s *Subscription) AutoUnsubscribe(max int) error
AutoUnsubscribe will issue an automatic Unsubscribe that is processed by the server when max messages have been received.
This can be useful when sending a request to an unknown number of subscribers. Request() uses this functionality.

Example
Code:

nc, _ := nats.Connect(nats.DefaultURL)
defer nc.Close()

received, wanted, total := 0, 10, 100

sub, _ := nc.Subscribe("foo", func(_ *nats.Msg) {
    received++
})
sub.AutoUnsubscribe(wanted)

for i := 0; i < total; i++ {
    nc.Publish("foo", []byte("Hello"))
}
nc.Flush()

fmt.Printf("Received = %d", received)



同步任务接收消息
func (*Subscription) NextMsg 
func (s *Subscription) NextMsg(timeout time.Duration) (*Msg, error)
NextMsg will return the next message available to a synchronous subscriber or block until one is available. 
A timeout can be used to return when no message has been delivered.

Example
Code:

nc, _ := nats.Connect(nats.DefaultURL)
defer nc.Close()

sub, _ := nc.SubscribeSync("foo")
m, err := sub.NextMsg(1 * time.Second)
if err == nil {
    fmt.Printf("Received a message: %s\n", string(m.Data))
} else {
    fmt.Println("NextMsg timed out.")
}


发布消息
func (*Conn) PublishRequest ¶ Uses
参数：发布主题和接受者返回的主题
func (nc *Conn) PublishRequest(subj, reply string, data []byte) error
PublishRequest will perform a Publish() excpecting a response on the reply subject. Use Request() for automatically waiting for a response inline.


func (*Conn) Subscribe
func (nc *Conn) Subscribe(subj string, cb MsgHandler) (*Subscription, error)
Subscribe will express interest in the given subject. The subject can have wildcards (partial:*, full:>).
Messages will be delivered to the associated MsgHandler. If no MsgHandler is given, 
the subscription is a synchronous subscription and can be polled via Subscription.NextMsg().


生成唯一的标识
func NewInbox ¶ Uses
func NewInbox() string
NewInbox will return an inbox string which can be used for directed replies from subscribers. These are guaranteed to be unique, but can be shared and subscribed to by others.
