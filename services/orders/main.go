package main

func main() {
	httpServer := NewHttpServer(":8000")
	go httpServer.Run() // since the http server is blocking, we need to run it in a goroutine

	gRPCServer := NewGRPCServer(":9000")
	gRPCServer.Run()
}
