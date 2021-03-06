package KCP_Net

import (
	"glog-master"

	kcp "github.com/xtaci/kcp-go"
)

// 初始化操作
func init() {
	Connect()
	return
}

// 网络模型的初始化操作
func Connect() {

	// 初始化数据
	lis, err := kcp.ListenWithOptions(":10000", nil, 10, 3)
	if err != nil {
		glog.Error(err)
		return
	}

	_ = lis

	return
}

//Client: full demo
// 例子管理操作
//kcpconn, err := kcp.DialWithOptions("192.168.0.1:10000", nil, 10, 3)
//Server: full demo
// 数据整理操作 服务器
//lis, err := kcp.ListenWithOptions(":10000", nil, 10, 3)
