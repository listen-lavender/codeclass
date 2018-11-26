package main

import (
	"flag"
	"fmt"
	"go-concurrentMap-master"

	"code.google.com/p/go.net/websocket"
)

// 全局的网络信息结构
var G_PlayerData map[string]*NetDataConn
var G_PlayerNet map[string]int // 心跳结构信息存储的结构
var G_PlayerNetSys map[string]int
var G_Net_Count map[string]int
var M *concurrent.ConcurrentMap // 并发安全的
var addr = flag.String("addr", "127.0.0.1:8888", "http service address")
var WS *websocket.Conn

// 游戏服务器的初始化
func init() {
	// 初始化
	G_PlayerData = make(map[string]*NetDataConn)
	G_PlayerNet = make(map[string]int)
	G_PlayerNetSys = make(map[string]int)
	G_Net_Count = make(map[string]int)
	// 并发安全的初始化
	M = concurrent.NewConcurrentMap()
	//go G_timer()
	//go G_timeout_kick_Player()
	return
}

func Go_func() {
	fmt.Println("Golang语言社区")
	return
}

// 服务器的初始化
func GameServerINIT() {
	// 1 主动链接网关服  -- 获取IP+端口
	// 2 建立与网关服的心跳-- 确保server正常 -- kill
	url := "ws://" + *addr + "/GolangLtd"
	WS, err := websocket.Dial(url, "", "test://golang/")
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	go GS2GW_Timer(WS)
}
