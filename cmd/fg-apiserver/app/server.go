package app

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/onexstack/fastgo/cmd/fg-apiserver/app/options"
	"github.com/onexstack/fastgo/pkg/version"
)

var configFile string // 配置文件路径

func NewFastGOCommand() *cobra.Command {
	opts := options.NewServerOptions()

	cmd := &cobra.Command{
		Use:          "fg-apiserver",
		Short:        "A very lightweight full go project",
		Long:         `A very lightweight full go project, designed to help beginners quickly learn Go project development.`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(opts)

		},
		Args: cobra.NoArgs,
	}

	cobra.OnInitialize(onInitialize)

	// cobra 支持持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	// 推荐使用配置文件来配置应用，便于管理配置项
	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", filePath(), "Path to the fg-apiserver configuration file.")

	// 添加 --version 标志
	version.AddFlags(cmd.PersistentFlags())

	return cmd
}

func run(opts *options.ServerOptions) error {
	// 如果传入 --version，则打印版本信息并退出
	version.PrintAndExitIfRequested()

	// 将 viper 中的配置解析到选项 opts 变量中.
	if err := viper.Unmarshal(opts); err != nil {
		return err
	}

	// 验证选项 opts 变量.
	if err := opts.Validate(); err != nil {
		return err
	}

	// 获取应用配置.
	// 将命令行选项和应用配置分开，可以更加灵活的处理 2 种不同类型的配置.
	cfg, err := opts.Config()
	if err != nil {
		return err
	}

	// 创建服务器实例.
	server, err := cfg.NewServer()
	if err != nil {
		return err
	}

	// 启动服务器
	return server.Run()

}
