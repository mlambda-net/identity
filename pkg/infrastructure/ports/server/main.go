package main

import "github.com/mlambda-net/identity/pkg/infrastructure/endpoint/server"

func main() {
	s := server.NewServer()
	s.Start()
	s.Wait()
}
