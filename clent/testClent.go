package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/testdata"
	"flag"
	"github.com/astaxie/beego"
	"grpc_test/testpb"
	"time"
	"golang.org/x/net/context"
	"google.golang.org/grpc/credentials"
)

const (
	testport = "localhost:10025"
	tlsbool  = false
)

var (
	testcaFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
	testserverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

func main() {
	var opts []grpc.DialOption

	if *testcaFile == "" {
		*testcaFile = testdata.Path("ca.pem")
	}
	//构建证书
	if !tlsbool {
		creds, err := credentials.NewClientTLSFromFile(*testcaFile, *testserverHostOverride)
		if err != nil {
			beego.Error("构建证书失败")
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	conn, err := grpc.Dial(testport, opts...)
	if err != nil {
		beego.Error("建立连接失败", err)
	}
	defer conn.Close()
	client := testpb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 12*time.Second)
	defer cancel()
	startTime := time.Now()
	for {
		value, err := client.GetValue(ctx, &testpb.InsertKey{"you dail!", "hahahaha "})
		time.Sleep(1 * time.Second)
		if err != nil {
			endTime := time.Now()
			beego.Debug(endTime.Sub(startTime))
			break
		} else {
			beego.Debug("得到的结果", value.Value)
		}
	}
}
