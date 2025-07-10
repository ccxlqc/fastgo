package options

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func (o *MySQLOptions) NewDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(o.Addr), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(o.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(o.MaxOpenConnections)
	sqlDB.SetConnMaxLifetime(o.MaxConnectionLifeTime)

	return db, nil
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

func (o *MySQLOptions) Validate() error {
	// 验证MySQL地址格式
	if o.Addr == "" {
		return fmt.Errorf("mysql server address cannot be empty")
	}

	// 检查地址格式是否为host:port
	host, portStr, err := net.SplitHostPort(o.Addr)
	if err != nil {
		return fmt.Errorf("invalid MySQL address format '%s': %w", o.Addr, err)
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
	if o.Username == "" {
		return fmt.Errorf("mysql username cannot be empty")
	}

	if o.Password == "" {
		return fmt.Errorf("mysql password cannot be empty")
	}

	if o.Database == "" {
		return fmt.Errorf("mysql database name cannot be empty")
	}

	// 验证连接池参数
	if o.MaxIdleConnections <= 0 {
		return fmt.Errorf("mysql max idle connections must be greater than 0")
	}

	if o.MaxOpenConnections <= 0 {
		return fmt.Errorf("mysql max open connections must be greater than 0")
	}

	if o.MaxIdleConnections > o.MaxOpenConnections {
		return fmt.Errorf("mysql max idle connections cannot be greater than max open connections")
	}

	if o.MaxConnectionLifeTime <= 0 {
		return fmt.Errorf("mysql max connection lifetime must be greater than 0")
	}

	return nil
}
