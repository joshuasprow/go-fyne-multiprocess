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
	"os"
	"time"

	pb "github.com/joshuasprow/go-fyne-multiprocess/api"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func newClient() (
	client pb.GreeterClient,
	done func() error,
	err error,
) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return client, func() error { return nil }, errors.Wrap(err, "grpc.Dial")
	}

	c := pb.NewGreeterClient(conn)

	return c, conn.Close, nil
}

func sayHello() (string, error) {
	c, done, err := newClient()
	if err != nil {
		return "", errors.Wrap(err, "newClient")
	}
	defer done()

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		return "", errors.Wrap(err, "c.SayHello")
	}

	return r.GetMessage(), nil
}

func sayGoodbye() (string, error) {
	c, done, err := newClient()
	if err != nil {
		return "", errors.Wrap(err, "newClient")
	}
	defer done()

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayGoodbye(ctx, &pb.GoodbyeRequest{Name: name})
	if err != nil {
		return "", errors.Wrap(err, "c.SayGoodbye")
	}

	return r.GetMessage(), nil
}
