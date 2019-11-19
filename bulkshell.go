package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/mattn/go-shellwords"
)

func main() {
	wg := &sync.WaitGroup{}
	i := 0
	shFilename := "shell.txt"
	dirFilename := "dir.txt"
	sq := make(chan string, 1)
	dq := make(chan string, 1)
	go func() {
		var dirs []string
		for target := range dq {
			dirs = append(dirs, target)
		}
		wg.Done()
		for command := range sq {
			for _, target := range dirs {
				out, err := runCmd(command, target)
				if err != nil {
					fmt.Printf(target+" error:%v\n", err)
					continue
				}
				splits := strings.Split(target, "/")
				name := splits[len(splits)-1]
				write("shell_"+fmt.Sprint(i), name, out)
			}
			i++
		}
		wg.Done()
	}()
	read(wg, dirFilename, dq)
	close(dq)
	read(wg, shFilename, sq)
	close(sq)
	wg.Wait()
}

func read(wg *sync.WaitGroup, filename string, q chan string) {
	wg.Add(1)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf(filename+" error:%v\n", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		q <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("scan error:%v", err)
	}
}

func write(dir string, name string, out []byte) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0777)
	}
	file, err := os.OpenFile(dir+"/"+name+".txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("write error:%v", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	bw := bufio.NewWriter(writer)
	bw.Write(out)
	bw.Flush()
}

func runCmd(cmdstr string, target string) ([]byte, error) {
	c, err := shellwords.Parse(cmdstr)
	if err != nil {
		return nil, fmt.Errorf("Parse %v:", err)
	}
	var out []byte
	switch len(c) {
	case 0:
		return nil, nil
	case 1:
		cmd := exec.Command(c[0])
		cmd.Dir = target
		out, err = cmd.Output()
	default:
		cmd := exec.Command(c[0], c[1:]...)
		cmd.Dir = target
		out, err = cmd.Output()
	}
	if err != nil {
		return nil, fmt.Errorf("Output %v:", err)
	}
	return out, nil
}
