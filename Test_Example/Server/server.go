package main

import (
	"Test_Example/conf/gen-go/tutorial"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TServerTransport
	var err error
	transport, err = thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}
	handler := NewServiceHandler()
	processor := tutorial.NewUserServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	fmt.Println("Running server..."+addr)
	return server.Serve()
}
