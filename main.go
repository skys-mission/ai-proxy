package main

import (
	"context"
	"fmt"

	"github.com/skys-mission/pkg/go/iplocation"
)

func main() {
	client := iplocation.New()
	address, err := client.GetIPAddress(context.Background())
	if err != nil {
		return
	}
	fmt.Print("%+v\n", address)
}
