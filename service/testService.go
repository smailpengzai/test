package main

import (
	"context"
	"grpc_test/testpb"
	"math/rand"
	"net"
	"fmt"
	"github.com/astaxie/beego"
	"google.golang.org/grpc"
	"flag"
	"google.golang.org/grpc/testdata"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

type testServer struct {
}

const testport = ":10025"

var (
	testcertFile = flag.String("cert_file", "", "The TLS cert file")
	testkeyFile  = flag.String("key_file", "", "The TLS key file")
)

func (this *testServer) GetValue(ctx context.Context, key *testpb.InsertKey) ( *testpb.GetoutValue,error) {

	var (
		testmap map[string]string = make(map[string]string)
		value   testpb.GetoutValue
	)
	testmap[key.Key]="信不信打死你！"+ key.Name
	value.Value = testmap[key.Key]
	beego.Debug(value.Value)
	value.Index = int64(rand.Intn(100))
	return  &value ,nil
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost%s", testport))
	if err != nil {
		beego.Error("开启监听失败！", err)
	}
	var (
		opts        []grpc.ServerOption
		testServer	testServer

	)
	if *testcertFile == "" {
		*testcertFile = testdata.Path("server1.pem")
	}
	if *testkeyFile == "" {
		*testkeyFile = testdata.Path("server1.key")
	}
	testcreds, err := credentials.NewServerTLSFromFile(*testcertFile, *testkeyFile)
	//证书
	opts = []grpc.ServerOption{grpc.Creds(testcreds)}

	grpcServer := grpc.NewServer(opts...)
	testpb.RegisterGreeterServer(grpcServer, &testServer)
	grpcServer.Serve(listener)

	reflection.Register(grpcServer)
}
