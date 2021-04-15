package main

import (
	"fmt"
	"io"
	"os"
	"os/signal"
)

var (
	TestFile = "./test/unusual.txt"
)

func main() {
	go Handler(TestFile)
	Exit(TestFile)
}

type File struct {
	fd *os.File
}

var fileFd *File

func GetFileFd(filename string) (*File, error) {
	var tmpFd *File
	if fileFd != nil {
		return fileFd, nil
	}

	tmpFd, err := newFileFd(filename)
	if err != nil {
		return nil, err
	}

	return tmpFd, nil
}

func newFileFd(filename string) (*File, error) {
	fd, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return &File{fd: fd}, nil
}

func (dev *File) ReadCdev() ([]byte, error) {
	var data = make([]byte, 4096)
	_, err := dev.fd.Read(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (dev *File) WriteCdev(data []byte) error {
	_, err := dev.fd.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func Handler(filename string) error {
	var err error
	fileFd, err := GetFileFd(filename)
	if err != nil {
		return err
	}

	defer fileFd.fd.Close()
	for {
		data, err := fileFd.ReadCdev()
		if err == io.EOF || err != nil || len(data) == 0 {
			continue
		}
		fmt.Printf("键入新消息:%s", string(data))
	}
}

func Exit(filename string) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	s := <-c
	fmt.Println("截取到中断信号:", s)
	fmt.Println("...........\n清空文件")
	Empty(filename)
	fmt.Println("...........\n清空完成")
	os.Exit(0)
}

func Empty(filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
}
