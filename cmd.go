package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version   string
	buildDate string
	verbose   string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go-typevis",
	Short: "Welcome to go-typevis",
	Long:  `go-typevis is go type visualize.`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of go-typevis",
	Long:  `Print the version number of go-typevis`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("go-typevis Version: %s(%s)\n", version, buildDate)
	},
}

func Execute(_version, _buildDate string) {
	version = _version
	buildDate = _buildDate

	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.Flags().StringVarP(
		&verbose,
		"v", "v", "5", "verbose log",
	)
	RootCmd.AddCommand(versionCmd)

	typesCmd.Flags().StringVarP(
		&pkgPath,
		"package", "p", "", "package path",
	)
	RootCmd.AddCommand(typesCmd)
}

func initConfig() {
}

var (
	pkgPath string
)

var typesCmd = &cobra.Command{
	Use:   "types",
	Short: "",
	Long: `
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return types()
	},
}

func types() error {
	pkg := analysis(typeOption{pkgPath: pkgPath})
	render(pkg)
	return nil
}
