package mq

import (
	"context"
	"time"

	"github.com/nsqio/go-nsq"
)

type NsqConsumer struct {
	lookupds            []string        // nsq lookupds 地址
	lookupdPollInterval time.Duration   // nsq consumer config
	maxInFlight         int             // nsq consumer config
	consumers           []*nsq.Consumer // 注册的 consumers 对象
	middlewares         []middleware    // 中间件数组
}

// 消息 handler 函数定义
// 参考了 grpc， echo 的 handler 函数。引入两个参数一个是 context 上下文，一个是 nsq 的消息体
type MqHandlerFunc func(ctx context.Context, msg *nsq.Message) error

// 初始化 Nsq consumer 对象
func NewNsqConsumer(lookupds []string, pollInterval time.Duration, maxInFlight int) *NsqConsumer {
	return &NsqConsumer{
		lookupds:            lookupds,
		lookupdPollInterval: pollInterval,
		maxInFlight:         maxInFlight,
	}
}

// 将 我们定义的 MqHandlerFunc 转换成 nsq.Consumer 内接受的 nsq.HandlerFunc
func (n *NsqConsumer) toNsqHandler(handlerFunc MqHandlerFunc) nsq.HandlerFunc {
	return func(msg *nsq.Message) error {
		ctx := context.TODO()
		return handlerFunc(ctx, msg)
	}
}

// 中间函数格式
type middleware func(handler MqHandlerFunc) MqHandlerFunc

// 将中间存入数组
func (n *NsqConsumer) Use(mids ...middleware) {
	for _, mid := range mids {
		n.middlewares = append(n.middlewares, mid)
	}
}

// 将中间件串联起来，
func (n *NsqConsumer) wrap(h MqHandlerFunc) MqHandlerFunc {
	// 这边为什么采用数组倒序？
	// 一般情况下，我们期望最先加入中间件数组的函数时最先被调用的（位于洋葱模型的最外层）
	// wrap 函数中传入 handler 函数是最原始的 handler 函数，没有被中间件函数包裹的
	// 按照 code2 中的  newHandler = middlewareFunc1(middlewareFunc2(middlewareFunc3(handler))) 来看
	// 接受最原始的 handler 函数的中间件函数肯定是最后加入中间数组的入参，依次类推，因此 中间件数组迭代的时候是倒序的
	for i := len(n.middlewares) - 1; i >= 0; i-- {
		mid := n.middlewares[i]
		h = mid(h)
	}
	return h
}

// 注册 topic 和 handler Func, 类似 echo 的路由配置
func (n *NsqConsumer) RegisterHandler(topic, channel string, handler MqHandlerFunc) error {
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = n.lookupdPollInterval
	c, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		return err
	}
	c.ChangeMaxInFlight(n.maxInFlight)
	c.AddHandler(n.toNsqHandler(handler))
	n.consumers = append(n.consumers, c)
	return nil
}

// 模拟 grpc 和 echo 的开启服务的逻辑
// 因为 consumer 启动的时候本身就才采用了 goroutine, 所以在 nsq Start 的时候就不用像 echo 和 grpc 一样用 goroutine
func (n *NsqConsumer) Start() error {
	for _, h := range n.consumers {
		if err := h.ConnectToNSQLookupds(n.lookupds); err != nil {
			return err
		}
	}
	return nil
}

// gracefully close consumer
func (n *NsqConsumer) Close() {
	for _, h := range n.consumers {
		h.Stop()
	}
}
