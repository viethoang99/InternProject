package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"Test_Example/conf/gen-go/tutorial"
)

func handleClient(client *tutorial.UserServiceClient) (err error) {
	client.Ping(context.Background())
	fmt.Println("ping......")
	////khai bao user
	user:=tutorial.NewUser()
	user.UserID="a_1234"
	//user.Username="Hoang...................."
	//user.DepartmentID="3"
	//////them user
	////s,err := client.AddUser(context.Background(),user)
	////fmt.Println(s)
	////sua user
	//s1,err:=client.Put(context.Background(),user)
	//fmt.Println(s1)
	//Delete
	s2,err:=client.Delete(context.Background(),user.UserID)
	fmt.Println(s2)
	//Lay tat ca user
	users,err := client.GetAll(context.Background())
	for i:=0;i<len(users);i++{
		fmt.Printf("Người thứ %d Tên: %s,%s\n", i, users[i].UserID,users[i].Username)
	}
	return err
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TTransport
	var err error
	transport, err = thrift.NewTSocket(addr)
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return err
	}
	if transport == nil {
		return fmt.Errorf("Error opening socket, got nil transport. Is server available?")
	}
	transport,_ = transportFactory.GetTransport(transport)
	if transport == nil {
		return fmt.Errorf("Error from transportFactory.GetTransport(), got nil transport. Is server available?")
	}

	err = transport.Open()
	if err != nil {
		return err
	}
	defer transport.Close()
	return handleClient(tutorial.NewUserServiceClientFactory(transport, protocolFactory))
}
