package cmd

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/fanaticscripter/EggOrganizer/config"
)

var Version string

const _cfgFileEnvVar = "EGGORGANIZER_CONFIG_FILE"

var (
	_verbose bool
	_debug   bool
	_cfgFile string
	_config  *config.Config
	_rootCmd = &cobra.Command{
		Use:     filepath.Base(os.Args[0]),
		Version: getVersion(),
	}
)

func init() {
	cobra.OnInitialize(cobraHouseKeeping, configureLoggingLevel)

	_rootCmd.PersistentFlags().StringVar(&_cfgFile, "config", "",
		"config file, could also be set through env var "+_cfgFileEnvVar+" (default is config.toml in pwd or the program directory)")
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
	if _cfgFile != "" {
		viper.SetConfigFile(_cfgFile)
	} else if os.Getenv(_cfgFileEnvVar) != "" {
		viper.SetConfigFile(os.Getenv(_cfgFileEnvVar))
	} else {
		viper.AddConfigPath(".")
		if prog, err := os.Executable(); err == nil {
			viper.AddConfigPath(filepath.Dir(prog))
		}
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
	}
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

func getVersion() string {
	if Version != "" {
		return Version
	}
	return "dev"
}
