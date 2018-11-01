package Config

// 配置文件结构信息
type Config struct {
	ID   string // ID 序号信息
	Name string // 名字
	IP   string // 服务器IP
	Port string // 服务器端口
	Dsp  string // 备注信息
}

// 获取配置文件
func init() {
	// 初始化数据操作
}

// 读取配置文件-- 加载配置文件数据
func ReadDataCSV(strFileName string, strFilePath string) {

	return
}
