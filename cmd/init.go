package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize zeed in your repository",
	Long: `Initialize zeed in your repository.

If no repository provided, this command will init zeed in the current directory:
	1. create .zeed directory inside your repository
	2. create .zeed.yaml config file inside .zeed
All files related to zeed will be inside .zeed`,
	RunE:          initRun,
	SilenceErrors: true, // errors are handled by cmd.Execute()
}

func initRun(_ *cobra.Command, _ []string) error {
	if viper.ConfigFileUsed() != "" {
		return errors.New(fmt.Sprintf("zeed is already initialized in `%s`", repository))
	}
	err := os.MkdirAll(cfgDir(), os.ModePerm)
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to create `%s` directory", cfgDir()))
	}
	err = ioutil.WriteFile(cfgFile(), []byte(""), 0644)
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to create `%s`", cfgFile()))
	}
	initConfig()
	fmt.Println(fmt.Sprintf("Successfully initialized zeed in `%s`", repository))
	fmt.Println(fmt.Sprintf("A zeed config file was created (`%s`)", cfgFile()))
	fmt.Println("Edit it according to your needs.")

	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)
}
