/*
 * @Date: 2022-10-10 18:31:06
 * @LastEditors: hxk
 * @LastEditTime: 2022-10-10 18:31:09
 * @FilePath: \myServer\kiface\irequest.go
 */
package kiface

type IRequest interface{
	GetConnection() IConnection	//获取请求连接信息
	GetData() []byte			//获取请求消息的数据
}