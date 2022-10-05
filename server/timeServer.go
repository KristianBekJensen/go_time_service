package main

import(
	"context"
	"flag"
	"fmt"
	//"io"
	"log"
	"net"
	//"os"
	"sync"
	"time"
	
	gRPC "github.com/kbekj/go_time_service/proto"
	"google.golang.org/grpc"
)

type TimeServer struct {
    // an interface that the server needs to have
    gRPC.UnimplementedTemplateServer
	time string

	mutex sync.Mutex
    
    // here you can impliment other fields that you want
}

var serverName = flag.String("name", "default", "Senders name") // set with "-name <name>" in terminal
var port = flag.String("port", "5400", "Server port")      

func main() {

	// setLog() //uncomment this line to log to a log.txt file instead of the console

	// This parses the flags and sets the correct/given corresponding values.
	flag.Parse()
	fmt.Println(".:server is starting:.")

	// starts a goroutine executing the launchServer method.
	go launchServer()

	// This makes sure that the main method is "kept alive"/keeps running
	for {
		time.Sleep(time.Second * 5)
	}
}

func (s *TimeServer) SendMyTime(ctx context.Context, Time *gPRC.Time) (*gRPC.Hi, error) {
    s.mutex.Lock()
	defer s.mutex.Unlock()

	return &gRPC.Hi{NewValue: s.time}, nil
}

func launchServer() {
	log.Printf("Server %s: Attempts to create listener on port %s\n", *serverName, *port)

	// Create listener tcp on given port or default port 5400
	list, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", *port))
	if err != nil {
		log.Printf("Server %s: Failed to listen on port %s: %v", *serverName, *port, err) //If it fails to listen on the port, run launchServer method again with the next value/port in ports array
		return
	}

	// makes gRPC server using the options
	// you can add options here if you want or remove the options part entirely
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	// makes a new server instance using the name and port from the flags.
	server := &TimeServer{
		time: time.Now().String(), // gives default value, but not sure if it is necessary
	}

	gRPC.RegisterTemplateServer(grpcServer, server) //Registers the server to the gRPC server.

	log.Printf("Server %s: Listening on port %s\n", *serverName, *port)

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
	// code here is unreachable because grpcServer.Serve occupies the current thread.
}