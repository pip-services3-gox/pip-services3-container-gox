package main

import (
	"context"
	"github.com/pip-services3-gox/pip-services3-container-gox/examples"
	"os"
)

func main() {
	process := examples.NewDummyProcess()
	process.Run(context.Background(), os.Args)
}
