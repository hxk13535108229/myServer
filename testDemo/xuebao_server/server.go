/*
 * @Date: 2022-10-10 18:24:54
 * @LastEditors: hxk
 * @LastEditTime: 2022-10-10 18:28:39
 * @FilePath: \myServer\testDemo\xuebao_server\server.go
 */
package main

import "xuebao/knet"

func main() {

	//1 创建一个server 句柄 s
	s := knet.NewServer("[zinx V0.1]")

	//2 开启服务
	s.Serve()
}
