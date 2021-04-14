package main

import "example/go-practise/dbus/method/pkg/serve"

func main() {
	srv := serve.GetService()
	err := srv.Init()
	if err != nil {
		panic(err)
	}
	srv.Loop()
}
