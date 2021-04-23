package main

import (
	"fmt"

	"time"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

//var nc *nats.Conn

func main() {
	//nc, _ = nats.Connect(nats.DefaultURL)
	//defer nc.Close()
	go subChan()
	go subFunc()

	pub()
}

func pub() {
	nc, _ := nats.Connect(nats.DefaultURL)
	var i int64
	for {
		i++
		nc.Publish("foo", []byte(fmt.Sprintf("Hi! This is # %d", i)))
		time.Sleep(time.Second)
	}
}

func subFunc() {
	nc, _ := nats.Connect(nats.DefaultURL)
	nc.Subscribe("foo", func(m *nats.Msg) {
		logrus.WithField("t", "func").Info(string(m.Data))
		//nc.Publish(m.Reply, []byte("I can help!"))
	})
}
func subChan() {
	nc, _ := nats.Connect(nats.DefaultURL)
	ch := make(chan *nats.Msg, 64)
	sub, err := nc.ChanSubscribe("foo", ch)
	defer sub.Unsubscribe()
	logrus.Error(err)

	for {
		msg := <-ch
		logrus.WithField("t", "chan").Info(string(msg.Data))
	}
}
