package config

const (
	CODE_SUCCESS = 200
	CODE_FAIL    = 400
)

// 配置文件路径
const PATH_ENV = "env"

// admin身份认证基础加密信息
const JWT_SECRET_ADMIN = "admin"

// admin身份认证过期时间(小时)
const JWT_EXP_ADMIN = 2

// 分页,单页默认条数
const PAGINATE_PAGESIZE = 15