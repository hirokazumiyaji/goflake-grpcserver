package main

import (
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func TestGetID(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := NewGoflakeClient(conn)
	r, err := c.GetID(context.Background(), &GetIDRequest{})
	if err != nil {
		t.Fatalf("could not get id: %v", err)
	}
	if r.ID == "" {
		t.Fatalf("id is empty.")
	}
}
