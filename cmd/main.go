package main

import (
	"fmt"

	"github.com/luopengift/log"
	"github.com/luopengift/tree"
)

func main() {

	fmt.Println("init")
	tr := tree.InitTree("/", "/")
	tr.Put("/test/config/luopeng")
	tr.Put("/test/config/luopeng1")
	tr.Put("/login")
	tr.Put("/logout")

	//用户信息
	tr.Put("/user")

	//上传文件
	tr.Put("/aws/s3")

	tr.Put("/api/v1/page")

	tr.Put("/service")
	tr.Put("/api/v1/service")
	tr.Put("/api/v1/service/version")
	tr.Put("/api/v1/service/healthcheck")
	tr.Put("/api/v1/service/runtime")
	tr.Put("/api/v1/ServiceIpaddrEnv")

	//发布系统
	tr.Put("/api/v1/publish")
	tr.Put("/api/v1/publish/task")
	tr.Put("/api/v1/publish/log")
	tr.Put("/api/v1/publish/state")
	tr.Put("/api/v1/script")
	tr.Put("/api/v1/script.sh")
	tr.Put("/api/v1/pkgs/version")
	tr.Put("/api/v1/env")

	tr.Put("/api/v1/aws/ri")
	tr.Put("/api/v1/aws/si")
	tr.Put("/api/v1/aws/ec2")
	tr.Put("/api/v1/aws/ec2/type")
	tr.Put("/api/v1/aws/ec2/snapshot")
	tr.Put("/api/v1/aws/ec2/tags")
	tr.Put("/api/v1/aws/sqs")

	tr.Put("/upload")

	tr.Put("/api/v1/sync/cdn")

	//Prometheus alertManager
	tr.Put("/alert")
	tr.Put("/prometheus/config")
	tr.Put("/prometheus/rule")
	tr.Put("/prometheus/silence")
	tr.Put("/prometheus/sync_rule")

	tr.Put("/wechat")
	tr.Put("/wechat/sendwx")
	tr.Put("/api/v1/mail")

	tr.Put("/web")
	tr.Put("/")
	log.Display("$$", tr.Node)
	tr.Dump("txt")
}
