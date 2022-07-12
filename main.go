/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	helloworld "github.com/katainaka0503/grpc-pr-env-test-backend/helloworld"
	pb "github.com/katainaka0503/grpc-pr-env-test-frontend/executeGreeting"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	port = flag.Int("port", 50052, "The server port")
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

type server struct {
	connection helloworld.GreeterClient
	pb.UnimplementedExecuteGreetingServer
}

func (s *server) ExecuteGreeting(ctx context.Context, in *pb.ExecuteGreetingRequest) (*pb.ExecuteGreetingReply, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	r, err := s.connection.SayHello(ctx, &helloworld.HelloRequest{Name: *name})
	if err != nil {
		return nil, err
	}
	log.Printf("Greeting: %s", r.GetMessage())

	return &pb.ExecuteGreetingReply{Message: r.GetMessage()}, nil
}

var _ pb.ExecuteGreetingServer = &server{}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()))

	pb.RegisterExecuteGreetingServer(s, &server{connection: c})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
