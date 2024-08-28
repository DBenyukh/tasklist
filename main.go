package main

import (
	"context"
	"sync"

	"tasklist/env"
	"tasklist/server"
)

func main() {
	env.Load()

	ctx := context.Background()
	wg := &sync.WaitGroup{}

	server.Start(ctx, wg)

	wg.Wait()

}
