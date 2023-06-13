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

	grpctrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

	helloworld "github.com/katainaka0503/grpc-pr-env-test-backend/helloworld"
	pb "github.com/katainaka0503/grpc-pr-env-test-frontend/executeGreeting"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "katainaka"
)

var (
	port = flag.Int("port", 50052, "The server port")
	addr = flag.String("addr", "backend-gateway:80", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

type server struct {
	connection helloworld.GreeterClient
	pb.UnimplementedExecuteGreetingServer
}

var _ pb.ExecuteGreetingServer = &server{}

func (s *server) ExecuteGreeting(ctx context.Context, in *pb.ExecuteGreetingRequest) (*pb.ExecuteGreetingReply, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	span, ok := tracer.SpanFromContext(ctx)
	if !ok {
		log.Printf("failed to fetch span")
	}

	span.Context().ForeachBaggageItem(func(k string, v string) bool {
		log.Printf("MetaData: %v: %v\n", k, v)
		return true
	})

	r, err := s.connection.SayHello(ctx, &helloworld.HelloRequest{Name: *name})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	log.Printf("Greeting: %s\n", r.GetMessage())

	return &pb.ExecuteGreetingReply{Message: r.GetMessage()}, nil
}

func main() {
	flag.Parse()

	tracer.Start()
	defer tracer.Stop()

	// InterceptorでOpenTelemetryを仕込む
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(grpctrace.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(grpctrace.StreamClientInterceptor()))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(grpctrace.UnaryServerInterceptor()),
		grpc.StreamInterceptor(grpctrace.StreamServerInterceptor()))

	pb.RegisterExecuteGreetingServer(s, &server{connection: c})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
