package examples

import (
	"context"
	"fmt"
	"log"
	"time"
)

type ctxKey struct{}

var myCtxKey = &ctxKey{}

func passValueWithCtx(v interface{}) {
	ctx := context.WithValue(context.Background(), myCtxKey, v)
	go nextFn(ctx)
	time.Sleep(1 * time.Second)
}

func nextFn(ctx context.Context) {
	v := ctx.Value(myCtxKey)
	fmt.Printf("%T, %v\n", v, v)
}

func timeoutWithCtx() {
	log.Printf("now: %v\n", time.Now())

	ctx, fn := context.WithTimeout(context.Background(), 2*time.Second)
	defer fn()
	go timeoutFn(ctx)

	time.Sleep(3 * time.Second)

	select {
	case <-ctx.Done():
		log.Println("timeoutWithCtx timeout")
		return
	default:
	}
}

func timeoutFn(ctx context.Context) {
	go timeoutFn1(ctx)

	for {
		select {
		case <-ctx.Done():
			log.Printf("ctx: %v\n", ctx)
			log.Println("timeoutFn---", ctx.Err())
			return
		default:
		}
	}
}

func timeoutFn1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("ctx: %v\n", ctx)
			log.Println("timeoutFn1---", ctx.Err())
			return
		default:
		}
	}
}

func cancelCtx() {
	newCtx, Fn := context.WithCancel(context.Background())
	go cancelFn(newCtx)

	// time.Sleep(2 * time.Second)
	Fn()
	select {
	case <-newCtx.Done():
		log.Println("ctx done")
		return
	default:
	}
}

func cancelFn(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Printf("cancelFn ctx: %v\n", ctx)
			log.Println("cancelFn ctx Done")
			return
		default:
		}
	}
}

func deadlineCtx() {
	ctx, fn := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer fn()

	log.Printf("now: %v\n", time.Now())

	go someTimeConsumingFn(ctx)

	for {
		select {
		case <-time.After(2 * time.Second):
			log.Println("after 2 seconds we reach deadline")
			return
		case <-ctx.Done():
			log.Println("deadlineCtx ---", ctx.Err())
			return
		default:
		}
	}
}

func someTimeConsumingFn(ctx context.Context) {
	v, ok := ctx.Deadline()
	log.Println("deadline: ", v, ok)

	for {
		select {
		case <-ctx.Done():
			log.Printf("someTimeConsumingFn: %v, %v\n", time.Now(), ctx.Err())
			return
		default:
		}
	}
}
