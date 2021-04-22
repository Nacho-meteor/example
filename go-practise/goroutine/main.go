package main

import (
	"flag"
	"os/exec"
	"sync"
)

var (
	num = flag.Int("num", 100, "请求次数")
)

//Execute command line
func Task(cmdline string, wg *sync.WaitGroup) {
	cmd := exec.Command("/bin/bash", "-c", cmdline)
	cmd.CombinedOutput()
	wg.Done()
}

func main() {
	flag.Parse()

	wg := sync.WaitGroup{}
	wg.Add(*num)
	for i := 0; i < *num; i++ {
		go Task("wget https://www.baidu.com/", &wg)
	}
	wg.Wait()
}
