package main

import (
	appconfig "Test_Example/conf"
	"Test_Example/models"
	"flag"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/astaxie/beego"
	"os"
	"strings"
)
func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}
func main() {
	flag.Usage = Usage
	protocol := flag.String("P", "binary", "Specify the protocol (binary, compact, json, simplejson)")
	framed := flag.Bool("framed", false, "Use framed transport")
	buffered := flag.Bool("buffered", false, "Use buffered transport")
	addr := flag.String("addr", "localhost:9090", "Address to listen to")
	flag.Parse()

	beego.LoadAppConfig("ini", "/home/hoangviet/go/src/Test_Example/conf/app.conf")
	config := &appconfig.AppConfig{}
	etcdEndpoints := beego.AppConfig.String("etcd::etcdEndpoints")
	config.EtcdServerEndpoints = strings.Split(etcdEndpoints, ",")
	config.SourceMappingNewSID = beego.AppConfig.String("source_mapping_new::sid")
	config.SourceMappingNewHost = beego.AppConfig.String("source_mapping_new::host")
	config.SourceMappingNewPort = beego.AppConfig.String("source_mapping_new::port")
	config.SourceMappingNewProtocol = beego.AppConfig.String("source_mapping_new::protocol")
	appconfig.Config = config
	models.InitBigSetIf()
	var protocolFactory thrift.TProtocolFactory
	switch *protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	default:
		fmt.Fprint(os.Stderr, "Invalid protocol specified", protocol, "\n")
		Usage()
		os.Exit(1)
	}

	var transportFactory thrift.TTransportFactory
	if *buffered {
		transportFactory = thrift.NewTBufferedTransportFactory(8192)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}

	if *framed {
		transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	}
	// always run server here
	if err := runServer(transportFactory, protocolFactory, *addr); err != nil {
		fmt.Println("error running server:", err)
	}
}

