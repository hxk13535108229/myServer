/*
 * @Date: 2022-10-10 18:27:12
 * @LastEditors: hxk
 * @LastEditTime: 2022-10-10 18:27:16
 * @FilePath: \myServer\kiface\iserver.go
 */
package kiface

//定义服务器接口
type IServer interface{
    //启动服务器方法
    Start()
    //停止服务器方法
    Stop()
    //开启业务服务方法
    Serve()
}