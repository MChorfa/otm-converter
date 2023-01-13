package converter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/MChorfa/otm-converter/pkg/converter"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"goa.design/model/mdl"
)

var (
	version = "v0.0.1-alpha"
	name    = "otm"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Launch is the entrypoint that sets up and runs the Fleet commands.
func Execute() {
	initLogging(name, version)
	rootCmd := createRootCmd()
	rootCmd.AddCommand(createConvertCmd())
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Sorry. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

// initFatal prints an error message and exits with a non-zero status.
func initFatal(err error, message string) {
	fmt.Printf("Failed to start: %s: %v\n", message, err)
	os.Exit(1)
}

func createRootCmd() *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	rootCmd := &cobra.Command{
		Version: version,
		Short:   "Convert to Open Threat Model",
	}

	// rootCmd.PersistentFlags().StringP("config", "c", "", "Path to a configuration file")

	return rootCmd
}

func createConvertCmd() *cobra.Command {
	inputPath := ""
	ouputPath := ""
	convertCmd := &cobra.Command{
		Use:   "convert",
		Short: "Launch the Converter",
		Long:  "Launch the Converter to convert dsl file to otm file",
		Run: func(cmd *cobra.Command, args []string) {
			err := handleConvertCmd(inputPath, ouputPath)
			if err != nil {
				fmt.Printf("%v", err)
			}
			out := cmd.OutOrStdout()
			fmt.Fprintln(out)
		},
	}
	convertCmd.Flags().StringVarP(&inputPath, "input-path", "i", "design.json", "What is the input path?")
	convertCmd.Flags().StringVarP(&ouputPath, "output-path", "o", "otm.yaml", "What is the output path?")
	return convertCmd
}

func handleConvertCmd(inputPath string, ouputPath string) error {
	var design mdl.Design
	input := path.Join(inputPath)
	if fileExists(input) {
		b, err := ioutil.ReadFile(input)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(b, &design); err != nil {
			return fmt.Errorf("failed to load design: %s", err.Error())
		}
	}
	converter := converter.NewConverter(&design)
	converter.Convert(ouputPath)
	return nil
}

func initLogging(name, version string) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.TimestampFieldName = "ts"
	zerolog.LevelFieldName = "level"
	zerolog.MessageFieldName = "msg"
	zerolog.ErrorFieldName = "err"
	zerolog.CallerFieldName = "caller"
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
