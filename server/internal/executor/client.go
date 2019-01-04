package executor

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc/keepalive"

	"github.com/illush123/hanamizuki/proto"
	"github.com/illush123/hanamizuki/server/config"
	"google.golang.org/grpc"
)

type ExecutorClient struct {
	conn *grpc.ClientConn
}

var client *ExecutorClient

func Init() {
	conn, err := grpc.Dial(
		config.Config.Executor.Host+":"+config.Config.Executor.Port,
		grpc.WithInsecure(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                30 * time.Millisecond,
			Timeout:             20 * time.Millisecond,
			PermitWithoutStream: true,
		}))
	if err != nil {
		log.Fatal(err)
	}
	client = &ExecutorClient{conn: conn}
}

func StoreFiles(in *proto.VmissRequest) (*proto.Empty, error) {
	c := proto.NewVmissRPCClient(client.conn)
	return c.StoreFiles(context.Background(), in)
}
