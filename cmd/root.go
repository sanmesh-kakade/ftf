package cmd

import (
	"os"

	"github.com/sanmesh-kakade/ftf/cmd/module"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logLevel string
)

var rootCmd = &cobra.Command{
	Use:   "ftf",
	Short: "FTF CLI is a command-line interface (CLI) tool that aids in the development of Facets.Cloud modules via innersourcing flow",
	Long:  `FTF CLI is a command-line interface (CLI) tool that aids in the development of Facets.Cloud modules via innersourcing flow. It packs a set of commands that help in generating, validating, testing and registering Facets.Cloud modules.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubCommandPalettes() {
	rootCmd.AddCommand(module.ModuleCmd)
}

func init() {

	cobra.OnInitialize(initLogger)
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "log level to set (debug, info, warn, error, fatal, panic)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubCommandPalettes()
}
func initLogger() {
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "fatal":
		level = zap.FatalLevel
	case "panic":
		level = zap.PanicLevel
	default:
		level = zap.InfoLevel
	}

	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
		Development: false,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "message",
			LevelKey:       "level",
			TimeKey:        "time",
			NameKey:        "logger",
			CallerKey:      "caller",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	// Replace the global logger
	zap.ReplaceGlobals(logger)
}
