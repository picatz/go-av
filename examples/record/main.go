package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/picatz/go-av"
)

func main() {
	fh, err := os.CreateTemp(os.TempDir(), "go-av-example-*.m4a")
	if err != nil {
		panic(err)
	}
	fh.Close()
	fmt.Println("Recording to", fh.Name())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	av.RecordAudioToFile(ctx, fh.Name())

}
