/*
 * @Date: 2022-10-10 18:32:23
 * @LastEditors: hxk
 * @LastEditTime: 2022-10-10 18:38:14
 * @FilePath: \myServer\knet\request.go
 */
package knet

import "xuebao/kiface"

type Request struct {
	conn kiface.IConnection //已经和客户端建立好的 链接
	data []byte             //客户端请求的数据
}

//获取请求连接信息
func (r *Request) GetConnection() kiface.IConnection {
	return r.conn
}

//获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.data
}
