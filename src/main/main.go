package main

import (
	"community-api/config"
	"community-api/db"
	"community-api/grpc_server"
	"community-api/logging"
	"community-api/redis"
	"community-api/web"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	//加载初始化.toml默认配置
	if err := config.LoadConfigAndSetDefault(); err != nil {
		panic(err.Error())
	}

	//初始化日志配置
	if err := logging.InitZap(&config.GetConf().LogConf); err != nil {
		panic("InitLogger:" + err.Error())
	}
	//初始化mysql库连接
	if err := db.InitDbConn(config.GetConf()); err != nil {
		panic("init db " + err.Error())
	}

	//初始化redis连接
	redis.InitRedis(config.GetConf().RedisConf)

	//初始化lua脚本
	redis.StatusLua = redis.FromResource(config.GetConf().RedisConf.LuaPath+string(os.PathSeparator)+"status.lua", true)

	//启动grpc服务
	go grpc_server.StartComApiService(config.GetConf())
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	//启动web服务
	go web.Run()

	//测试函数
	//test.Test()

	//监听退出指令
	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			fmt.Println("退出", s)
			//服务退出
			web.Shutdown()
			grpc_server.StopComApiService()
			redis.Close()
			time.Sleep(time.Second * 3) //等待三秒
			os.Exit(0)
		default:
			fmt.Println("other", s)
		}
	}

}
