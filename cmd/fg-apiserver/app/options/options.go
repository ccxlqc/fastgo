package options

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

type MySQLOptions struct {
	Addr                  string        `json:"addr,omitempty" mapstructure:"addr"`
	Username              string        `json:"username,omitempty" mapstructure:"username"`
	Password              string        `json:"-" mapstructure:"password"`
	Database              string        `json:"database" mapstructure:"database"`
	MaxIdleConnections    int           `json:"max-idle-connections,omitempty" mapstructure:"max-idle-connections,omitempty"`
	MaxOpenConnections    int           `json:"max-open-connections,omitempty" mapstructure:"max-open-connections"`
	MaxConnectionLifeTime time.Duration `json:"max-connection-life-time,omitempty" mapstructure:"max-connection-life-time"`
}

func NewMySQLOptions() *MySQLOptions {
	return &MySQLOptions{
		Addr:                  "127.0.0.1:3306",
		Username:              "onex",
		Password:              "onex(#)666",
		Database:              "onex",
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: time.Duration(10) * time.Second,
	}
}

type ServerOptions struct {
	MySQLOptions *MySQLOptions `json:"mysql" mapstructure:"mysql"`
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		MySQLOptions: NewMySQLOptions(),
	}
}

func (o *ServerOptions) Validate() error {
	// 验证MySQL地址格式
	if o.MySQLOptions.Addr == "" {
		return fmt.Errorf("mysql server address cannot be empty")
	}

	// 检查地址格式是否为host:port
	host, portStr, err := net.SplitHostPort(o.MySQLOptions.Addr)
	if err != nil {
		return fmt.Errorf("invalid MySQL address format '%s': %w", o.MySQLOptions.Addr, err)
	}

	// 验证端口是否为数字
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("invalid MySQL port: %s", portStr)
	}

	// 验证主机名是否为空
	if host == "" {
		return fmt.Errorf("mysql hostname cannot be empty")
	}

	// 验证凭据和数据库名
	if o.MySQLOptions.Username == "" {
		return fmt.Errorf("mysql username cannot be empty")
	}

	if o.MySQLOptions.Password == "" {
		return fmt.Errorf("mysql password cannot be empty")
	}

	if o.MySQLOptions.Database == "" {
		return fmt.Errorf("mysql database name cannot be empty")
	}

	// 验证连接池参数
	if o.MySQLOptions.MaxIdleConnections <= 0 {
		return fmt.Errorf("mysql max idle connections must be greater than 0")
	}

	if o.MySQLOptions.MaxOpenConnections <= 0 {
		return fmt.Errorf("mysql max open connections must be greater than 0")
	}

	if o.MySQLOptions.MaxIdleConnections > o.MySQLOptions.MaxOpenConnections {
		return fmt.Errorf("mysql max idle connections cannot be greater than max open connections")
	}

	if o.MySQLOptions.MaxConnectionLifeTime <= 0 {
		return fmt.Errorf("mysql max connection lifetime must be greater than 0")
	}

	return nil
}
