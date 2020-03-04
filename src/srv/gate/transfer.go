package main

import "github.com/jeckbjy/gsk/arpc"

func Transfer(next arpc.Handler) arpc.Handler {
	return func(ctx arpc.Context) error {
		msg := ctx.Message()
		if msg.IsAck() {
			// 应答消息,发给客户端
		} else {
			// 客户端消息,发给服务器
		}
		return next(ctx)
	}
}
