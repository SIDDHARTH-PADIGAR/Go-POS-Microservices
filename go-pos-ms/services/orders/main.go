package main

func main() {
	httpServer := NewHttpServer(":8081")
	go httpServer.Run() //by adding go keyword, the http server will run in a goroutine, and convert it to a non-blocking call

	grpcServer := NewGRPCServer(":8080")
	grpcServer.Run()
}
