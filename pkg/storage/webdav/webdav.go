// webdav.go

package webdav

import (
	"fmt"
	"log"
	"net/http"

	"github.com/studio-b12/gowebdav"
)

// Config 结构体用于存储 WebDAV 连接信息。
type Config struct {
	URL      string `yaml:"url"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// WebDAV 结构体表示 WebDAV 客户端。
type WebDAV struct {
	Client *gowebdav.Client
	Config *Config
}

// NewClient 创建一个新的 WebDAV 客户端实例。
func NewClient(configMap map[string]any) (*WebDAV, error) {
	var (
		url      string
		user     string
		password string
	)

	// 从 map 中提取配置信息
	urlI, ok := configMap["URL"]
	if !ok {
		return nil, fmt.Errorf("URL 配置缺失")
	}
	url, ok = urlI.(string)
	if !ok {
		return nil, fmt.Errorf("URL 配置类型错误")
	}

	userI, ok := configMap["User"]
	if !ok {
		return nil, fmt.Errorf("User 配置缺失")
	}
	user, ok = userI.(string)
	if !ok {
		return nil, fmt.Errorf("User 配置类型错误")
	}

	passwordI, ok := configMap["Password"]
	if !ok {
		return nil, fmt.Errorf("Password 配置缺失")
	}
	password, ok = passwordI.(string)
	if !ok {
		return nil, fmt.Errorf("Password 配置类型错误")
	}

	config := &Config{
		URL:      url,
		User:     user,
		Password: password,
	}

	client := gowebdav.NewClient(config.URL, config.User, config.Password)
	// 设置传输方式为 HTTP
	client.SetTransport(&http.Transport{})

	return &WebDAV{
		Client: client,
		Config: config,
	}, nil
}

// Start 启动 WebDAV 客户端（用于测试和调试）。
func (w *WebDAV) Start() {
	log.Printf("WebDAV 客户端已初始化，准备进行操作")
}
