// To tun it : go run main.go -src img/ -dst output -task channel -poolsize 5

package main

import (
	"flag"
	"fmt"
	"time"
	"udemy/imgproc/task"

	"udemy/imgproc/filter"
)

func main() {
	srcDir := flag.String("src", "", "Input directory")
	dstDir := flag.String("dst", "", "Output directory")
	filterType := flag.String("filter", "grayscale", "grayscale/blur")
	taskType := flag.String("task", "channel", "channel/group")
	poolSize := flag.Int("poolsize", 4, "Workers pool size for channel task")
	flag.Parse()

	var f filter.Filter
	switch *filterType {
	case "grayscale":
		f = filter.GrayScale{}
	case "blur":
		f = filter.Blur{}
	}

	var t task.Tasker
	switch *taskType {
	case "channel":
		t = task.NewChanTask(*srcDir, *dstDir, f, *poolSize)
	case "group":
		t = task.NewWaitGrpTask(*srcDir, *dstDir, f)
	}

	start := time.Now()
	t.Process()
	elapsed := time.Since(start)
	fmt.Printf("Time elapsed: %s\n", elapsed)
}
