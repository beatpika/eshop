package service

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alicebob/miniredis/v2"
	tokenRedis "github.com/beatpika/eshop/app/token/biz/dal/redis"
	"github.com/beatpika/eshop/app/token/conf"
	redisdb "github.com/redis/go-redis/v9"
)

func setupTestRedis(t *testing.T) (*miniredis.Miniredis, func()) {
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatal(err)
	}

	// 使用miniredis创建一个redis客户端
	client := redisdb.NewClient(&redisdb.Options{
		Addr: mr.Addr(),
	})

	// 初始化Redis管理器
	tokenRedis.TokenManager = tokenRedis.NewTokenStore(client)

	return mr, func() {
		mr.Close()
	}
}

func setupTestConfig(t *testing.T) func() {
	// 创建临时配置目录
	tmpDir := filepath.Join(os.TempDir(), "eshop_test")
	err := os.MkdirAll(tmpDir, 0o755)
	if err != nil {
		t.Fatal(err)
	}

	// 创建测试配置文件
	confContent := []byte(`kitex:
  service: "token"
  address: ":8882"
  log_level: info
  log_file_name: "log/kitex.log"
  log_max_size: 10
  log_max_age: 3
  log_max_backups: 50

registry:
  registry_address:
    - 127.0.0.1:8500
  username: ""
  password: ""

mysql:
  dsn: "root:root@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"

redis:
  address: "127.0.0.1:6379"
  username: ""
  password: ""
  db: 0

jwt:
  secret_key: "test-secret-key"`)

	confDir := filepath.Join(tmpDir, "test")
	err = os.MkdirAll(confDir, 0o755)
	if err != nil {
		t.Fatal(err)
	}

	confFile := filepath.Join(confDir, "conf.yaml")
	err = os.WriteFile(confFile, confContent, 0o644)
	if err != nil {
		t.Fatal(err)
	}

	// 设置环境变量
	oldGoEnv := os.Getenv("GO_ENV")
	os.Setenv("GO_ENV", "test")

	// 设置配置文件路径并重新初始化配置
	oldConfigPath := conf.GetConfigPath()
	conf.SetConfigPath(tmpDir)

	// 强制重新初始化配置
	conf.ResetConfig()
	conf.GetConf()

	// 返回清理函数
	return func() {
		os.RemoveAll(tmpDir)
		os.Setenv("GO_ENV", oldGoEnv)
		conf.SetConfigPath(oldConfigPath)
		conf.ResetConfig()
	}
}
