package main

import (
	"fmt"
	"os"

	"hpgo.io/s3indexsync/internal/sync"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage %s <directory> <s3uri>\n", os.Args[0])
		os.Exit(1)
	}

	dir := os.Args[1]
	s3uri := os.Args[2]

	err := sync.Do(dir, s3uri)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

}
