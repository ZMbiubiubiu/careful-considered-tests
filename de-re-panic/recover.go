package main

import (
	"context"
	"fmt"
	"runtime/debug"
)

func AsyncPanicHandler(ctx context.Context) {
	if err := recover(); err != nil {
		fmt.Printf("msg=async_panic_handler_recovered||err=%+v||stack=%s\n", err, string(debug.Stack()))
	}
}

func serviceTaskAutoStopOne(ctx context.Context, task string) (err error) {
	// 能够兜住panic
	defer AsyncPanicHandler(ctx)

	if task == "panic" {
		panic("panic")
	}
	return nil
}

func main() {
	serviceTaskAutoStopOne(context.TODO(), "panic")
}
