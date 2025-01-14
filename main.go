package main

import (
	comment "github.com/ClubWeGo/commentmicro/kitex_gen/comment/commentservice"
	"log"
	"net"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/ClubWeGo/commentmicro/dal"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
)

func main() {
	dsn := "tk:123456@tcp(127.0.0.1:3306)/simpletk?charset=utf8&parseTime=True&loc=Local"
	dal.InitDB(dsn)

	r, err := etcd.NewEtcdRegistry([]string{"0.0.0.0:2379"})
	if err != nil {
		log.Fatal(err)
	}

	addr, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:10010")
	svr := comment.NewServer(new(CommentServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "commentservice"}),
		server.WithRegistry(r),
		server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
