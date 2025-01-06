package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/haierkeys/obsidian-image-api-gateway/pkg/path"

	"github.com/radovskyb/watcher"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type runFlags struct {
	dir     string // 项目根目录
	port    string // 启动端口
	runMode string // 启动模式
	config  string // 指定要使用的配置文件路径
}

func init() {
	runEnv := new(runFlags)

	var runCommand = &cobra.Command{
		Use:   "run [-c config_file] [-d working_dir] [-p port]",
		Short: "Run service",
		Run: func(cmd *cobra.Command, args []string) {

			if len(runEnv.dir) > 0 {
				err := os.Chdir(runEnv.dir)
				if err != nil {
					log.Println("failed to change the current working directory, ", err)
				}
				log.Println("working directory changed", zap.String("path", runEnv.dir).String)
			}

			if len(runEnv.config) <= 0 {
				if path.Exists("config.yaml") {
					runEnv.config = "config.yaml"
				} else if path.Exists("config/config-dev.yaml") {
					runEnv.config = "config/config-dev.yaml"
				} else {
					runEnv.config = "config/config.yaml"
				}
			}

			s, err := NewServer(runEnv)
			if err != nil {
				log.Println("failed to change the current working directory, ", err)
			}

			go func() {

				w := watcher.New()

				// 将 SetMaxEvents 设置为 1，以便在每个监听周期中至多接收 1 个事件
				// 如果没有设置 SetMaxEvents，默认情况下会发送所有事件。
				w.SetMaxEvents(1)

				// 只通知重命名和移动事件。
				w.FilterOps(watcher.Write)

				go func() {
					for {
						select {
						case event := <-w.Event:

							s.logger.Info("config watcher change", zap.String("event", event.Op.String()), zap.String("file", event.Path))
							s.sc.SendCloseSignal(nil)

							// 重新初始化 server
							s, err = NewServer(runEnv)
							if err != nil {
								s.logger.Error("service start err", zap.Error(err))
							}

						case err := <-w.Error:
							s.logger.Error("config watcher error", zap.Error(err))
						case <-w.Closed:
							log.Println("config watcher closed")
						}
					}
				}()

				// 监听 config.yaml 文件
				if err := w.Add(runEnv.config); err != nil {
					s.logger.Error("config watcher file error", zap.Error(err))
				}

				// 启动监听
				if err := w.Start(time.Second * 5); err != nil {
					s.logger.Error("config watcher start error", zap.Error(err))
				}
			}()

			quit1 := make(chan os.Signal)
			signal.Notify(quit1, syscall.SIGINT, syscall.SIGTERM)
			<-quit1
			s.sc.SendCloseSignal(nil)
			log.Println("service has been shut down.")

		},
	}

	rootCmd.AddCommand(runCommand)
	fs := runCommand.Flags()
	fs.StringVarP(&runEnv.dir, "dir", "d", "", "run dir")
	fs.StringVarP(&runEnv.port, "port", "p", "", "run port")
	fs.StringVarP(&runEnv.runMode, "mode", "m", "", "run mode")
	fs.StringVarP(&runEnv.config, "config", "c", "", "config file")

}
