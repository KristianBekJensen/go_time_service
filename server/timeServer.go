package main

import(
	gRPC "https://github.com/kbekj/go_time_service/proto"
	//"google.golang.org/grpc"
)

type TimeServer struct {
    // an interface that the server needs to have
    gRPC.UnimplementedTemplateServer
    
    // here you can impliment other fields that you want
}

func (s *Server) <endpoint name>(ctx context.Context, Time *gPRC.Time) (*gRPC.Hi, error) {
    //some code here
    ...
    ack :=  // make an instance of your return type
    return (ack, nil)
}

grpcServer := grpc.NewServer(opts...)
server := &Server{
    // set fields here 
}
gRPC.RegisterTemplateServer(grpcServer, server)
grpcServer.Serve(list)