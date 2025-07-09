package app

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/onexstack/fastgo/cmd/fg-apiserver/app/options"
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
			// 将 viper 中的配置解析到选项 opts 变量中.
			if err := viper.Unmarshal(opts); err != nil {
				return err
			}

			// 验证选项 opts 变量.
			if err := opts.Validate(); err != nil {
				return err
			}

			fmt.Printf("Read MySQL host from Viper: %s\n\n", viper.GetString("mysql.host"))
			jsonData, _ := json.MarshalIndent(opts, "", "  ")
			fmt.Println(string(jsonData))

			return nil
		},
		Args: cobra.NoArgs,
	}

	cobra.OnInitialize(onInitialize)

	// cobra 支持持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	// 推荐使用配置文件来配置应用，便于管理配置项
	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", filePath(), "Path to the fg-apiserver configuration file.")

	return cmd
}
