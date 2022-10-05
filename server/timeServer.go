package main

import(
	gRPC "go_time_service/proto"
	//"google.golang.org/grpc"
)

type TimeServer struct {
    // an interface that the server needs to have
    gRPC.UnimplementedTemplateServer
    
    // here you can impliment other fields that you want
}