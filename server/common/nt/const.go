package nt

const Token = "X-Auth-Token"

type Key string

const (
	DB Key = "db"

	SSH    = "ssh"
	RDP    = "rdp"
	VNC    = "vnc"
	Telnet = "telnet"
	K8s    = "kubernetes"

	AccessRuleAllow  = "allow"  // 允许访问
	AccessRuleReject = "reject" // 拒绝访问

	Custom     = "custom"      // 密码
	PrivateKey = "private-key" // 密钥

	JobStatusRunning    = "running"     // 计划任务运行状态
	JobStatusNotRunning = "not-running" // 计划任务未运行状态
	FuncShellJob        = "shell-job"   // 执行Shell脚本
	JobModeSelf         = "self"        // 本机

	SshMode      = "ssh-mode"      // ssh模式
	MailHost     = "mail-host"     // 邮件服务器地址
	MailPort     = "mail-port"     // 邮件服务器端口
	MailUsername = "mail-username" // 邮件服务账号
	MailPassword = "mail-password" // 邮件服务密码

	TypeUser  = "user"  // 普通用户
	TypeAdmin = "admin" // 管理员

	StatusEnabled  = "enabled"
	StatusDisabled = "disabled"

	SocksProxyEnable   = "socks-proxy-enable"
	SocksProxyHost     = "socks-proxy-host"
	SocksProxyPort     = "socks-proxy-port"
	SocksProxyUsername = "socks-proxy-username"
	SocksProxyPassword = "socks-proxy-password"

	LoginToken   = "login-token"
	AccessToken  = "access-token"
	ShareSession = "share-session"

	Anonymous = "anonymous"

	StorageLogActionRm       = "rm"       // 删除
	StorageLogActionUpload   = "upload"   // 上传
	StorageLogActionDownload = "download" // 下载
	StorageLogActionMkdir    = "mkdir"    // 创建文件夹
	StorageLogActionRename   = "rename"   // 重命名
)
