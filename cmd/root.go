package cmd

import (
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/fanaticscripter/EggOrganizer/config"
)

var (
	_verbose bool
	_debug   bool
	_config  *config.Config
	_rootCmd = &cobra.Command{
		Use: "EggOrganizer",
	}
)

func init() {
	cobra.OnInitialize(cobraHouseKeeping, configureLoggingLevel)

	_rootCmd.PersistentFlags().BoolVarP(&_verbose, "verbose", "v", false, "enable verbose logging")
	_rootCmd.PersistentFlags().BoolVar(&_debug, "debug", false, "enable debug logging")
}

func Execute() {
	if err := _rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func cobraHouseKeeping() {
	// Do not repeat error and print usage when a subcommand's RunE returns an
	// error. See https://github.com/spf13/cobra/issues/340.
	_rootCmd.SilenceUsage = true
	_rootCmd.SilenceErrors = true
}

func configureLoggingLevel() {
	log.SetLevel(log.WarnLevel)
	if _verbose {
		log.SetLevel(log.InfoLevel)
	}
	if _debug {
		log.SetLevel(log.DebugLevel)
	}
}

func initConfig() error {
	viper.SetConfigFile("config.toml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	_config = &config.Config{}
	if err := viper.Unmarshal(_config); err != nil {
		return err
	}
	if err := _config.Validate(); err != nil {
		return errors.Wrap(err, "invalid config")
	}
	return nil
}

func subcommandPreRunE(cmd *cobra.Command, args []string) error {
	if err := initConfig(); err != nil {
		return err
	}
	return nil
}
