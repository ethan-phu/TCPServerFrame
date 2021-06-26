package main

import "zinx/znet"

func main() {
	s := znet.NewServer("[zinx 0.1]")
	s.Serve()
}
