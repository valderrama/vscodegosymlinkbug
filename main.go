package main

import (
	"github.com/hyperledger/fabric/membersrvc/ca"
	"google.golang.org/grpc"
)

// Client
type Client struct {
	connection *grpc.ClientConn
}

func main() {

	client := &Client{}

	connection, _ := ca.GetClientConn("127.0.0.1:7054", "THIS_VALUE_IS_NOT_USED")

	client.connection = connection
}
