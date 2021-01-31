package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	log *zap.SugaredLogger
)

func init() {
	initLogger()

	go func() {
		sign := <-getExitSignalChanel()
		log.Infof("Got the signal %+v", sign)
		os.Exit(0)
	}()

}

func getExitSignalChanel() chan os.Signal {
	channel := make(chan os.Signal, 1)

	signal.Notify(channel,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGHUP,
	)

	return channel
}

func initLogger() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	log = logger.Sugar()
	log.Info("Logger for development mode is initialized")
}

var rootCmd = &cobra.Command{
	Use: "fact finder test microservice",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}

// Execute func
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Infof("Terminate program %+v", err)
		os.Exit(1)
	}
}
