package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/win5do/go-scaffold/pkg/logi"
	"github.com/win5do/go-scaffold/pkg/scaffold"
	"go.uber.org/zap"
)

var (
	log *zap.SugaredLogger

	rootCmd = &cobra.Command{
		Use: "scaffold",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func newCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "new",
		Short: "Generate scaffold project layout for Go.",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Debugf("args: %+v", args)
			if len(args) < 1 {
				return fmt.Errorf("please provide project name")
			}

			err := scaffold.New().Generate(os.Args[0], args[0])
			if err == nil {
				log.Info("creat success, exec `make run` to start service.")
			}

			return nil
		},
	}
}

func init() {
	// logi.SetLogger(logi.Logger(false))
	log = logi.Log.Sugar()
	rootCmd.AddCommand(newCmd())
}

func main() {
	_ = Execute()
}
