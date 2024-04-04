package cmd

import (

	"github.com/spf13/cobra"
)

var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "tense a url",
	Long:  `tense a url`,
	Run:   tenseUrl,
}

func tenseUrl(cmd *cobra.Command, args []string) {
	url, err := cmd.Flags().GetString("url")

	if err != nil {
		cmd.Help()
		return
	}

	if url == "" {
		cmd.Help()
		return
	}

	n, err := cmd.Flags().GetInt64("number")

	if err != nil {
		cmd.Help()
		return
	}

	if n <= 0 {
		cmd.Help()
		return
	}

	err = tense(url, n)
	if err != nil {
		cmd.Help()
		return
	}
}

func init() {
	rootCmd.AddCommand(urlCmd)

	urlCmd.Flags().StringP("url", "u", "", "url to tense")
	urlCmd.Flags().Int64P("number", "n", 1_000_000_000, "number of times to tense the url")

}

