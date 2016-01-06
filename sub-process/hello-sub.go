package main

import (
    "fmt"
	"log"
    "os"
	"os/exec"
    "io"
)

type prepender struct {
    prefix string
    writer io.Writer
}

func (pre prepender) Write(p []byte) (n int, err error){
    return fmt.Fprintf(pre.writer, "%s: %s", pre.prefix, p)
}

func main(){

    redisLogger := prepender{"redis", os.Stdout}
    redisLogger2 := prepender{"redis2", os.Stdout}
    
    cmd := exec.Command("redis-server", "--port", "6380")
    cmd2 := exec.Command("redis-server", "--port", "6381")

    cmd.Stdout = redisLogger
    cmd.Stderr = redisLogger
    cmd2.Stdout = redisLogger2
    cmd2.Stderr = redisLogger2

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err2 := cmd2.Start()
	if err2 != nil {
		log.Fatal(err2)
	}

	cmd.Wait()
}

