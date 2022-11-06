package main

import (
	"context"
	"fmt"
	"log"

	"github.com/quartz-technology/quicknode-sdks/go-sdk/client"
)

func main() {
	c := client.New()

	collection, err := c.GetCollectionDetails(context.TODO(), "0x60e4d786628fea6478f785a6d7e704777c86a7c6")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", collection)
}
