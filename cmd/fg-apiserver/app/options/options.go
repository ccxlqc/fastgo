package options

import (
	"fmt"
	"net"
	"strconv"

	"github.com/onexstack/fastgo/internal/apiserver"
	genericoptions "github.com/onexstack/fastgo/pkg/options"
)

type ServerOptions struct {
	MySQLOptions *genericoptions.MySQLOptions `json:"mysql" mapstructure:"mysql"`
	Addr         string                       `json:"addr" mapstructure:"addr"`
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		MySQLOptions: genericoptions.NewMySQLOptions(),
		Addr:         "0.0.0.0:6666",
	}
}

func (o *ServerOptions) Validate() error {
	if err := o.MySQLOptions.Validate(); err != nil {
		return err
	}

	// 验证服务器地址
	if o.Addr == "" {
		return fmt.Errorf("server address cannot be empty")
	}

	_, portStr, err := net.SplitHostPort(o.Addr)
	if err != nil {
		return fmt.Errorf("invalid server address: %s", o.Addr)
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return fmt.Errorf("invalid server port: %s", portStr)
	}

	if port <= 0 || port > 65535 {
		return fmt.Errorf("invalid server port: %d", port)
	}

	return nil
}

func (o *ServerOptions) Config() (*apiserver.Config, error) {
	return &apiserver.Config{
		MySQLOptions: o.MySQLOptions,
		Addr:         o.Addr,
	}, nil
}
