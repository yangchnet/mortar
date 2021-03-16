package main

import (
	"context"
	"fmt"
	"time"

	"github.com/wazsmwazsm/mortar"
)

func main() {
	pool, err := mortar.NewPool(10)
	if err != nil {
		panic(err)
	}

	pool.PanicHandler = func(r interface{}) {
		fmt.Printf("Warning!!! %s", r)
	}

	pool.Put(&mortar.Task{
		Handler: func(ctx context.Context, v ...interface{}) {
			panic("somthing wrong!")
		},
	})

	time.Sleep(1e9)
}
