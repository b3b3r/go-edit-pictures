package task

import (
	"fmt"
	"path"
	"path/filepath"
	"udemy/imgproc/filter"
)

type ChanTask struct {
	dirCtx
	Filter   filter.Filter
	PoolSize int
}

func NewChanTask(srcDir, dstDir string, filter filter.Filter, poolSize int) Tasker {
	return &ChanTask{
		Filter: filter,
		dirCtx: dirCtx{
			SrcDir: srcDir,
			DstDir: dstDir,
			files:  buildFilesList(srcDir),
		},
		PoolSize: poolSize,
	}
}

type jobReq struct {
	src string
	dst string
}

func (c *ChanTask) Process() error {
	size := len(c.files)
	jobs := make(chan jobReq, size)
	results := make(chan string, size)

	//init workers
	for i := 1; i < c.PoolSize; i++ {
		go worker(i, c, jobs, results)
	}

	//start workers
	for _, f := range c.files {
		filename := filepath.Base(f)
		dst := path.Join(c.DstDir, filename)
		jobs <- jobReq{
			src: f,
			dst: dst,
		}
	}
	close(jobs)

	for range c.files {
		fmt.Println(<-results)
	}
	return nil
}

func worker(id int, chanTask *ChanTask, jobs <-chan jobReq, results chan<- string) {
	for j := range jobs {
		fmt.Printf("Worker %d, started job %v\n", id, j)
		chanTask.Filter.Process(j.src, j.dst)
		fmt.Printf("Worker %d, finished job %v\n", id, j)
		results <- j.dst
	}
}
