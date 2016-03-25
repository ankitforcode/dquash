package main

import (
	"flag"
	"fmt"
)

func main() {
        var tag string
	flag.StringVar(&tag, "t", "", "Repository name and tag for new image")
	flag.Usage = func() {
		fmt.Printf("\nUsage: dquash [options]\n\n")
		fmt.Printf("Dquashes the layers of a tar archive on STDIN and streams it to STDOUT\n\n")
		fmt.Printf("Options:\n")
                flag.PrintDefaults()
	}
        flag.Parse()
}
