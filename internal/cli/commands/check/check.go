package check

import (
	"fmt"

	"github.com/Wa4h1h/websites-status-checker/internal/statuschecker"
	"github.com/spf13/cobra"
)

var (
	src         string
	concurrency int32
)

const DefaultNumOfConcurrentRequests = 10

func New(checker statuschecker.StatusChecker) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check [urls] [--src]",
		Short: "Check the HTTP status of one or more URLs",
		Long:  "Check the HTTP status of one or more URLs from a list of urls and/or --src",
		Run: func(cmd *cobra.Command, args []string) {
			results, done, err := checker.Check(int(concurrency), src, args...)
			if err != nil {
				return
			}

			for {
				select {
				case res := <-results:
					fmt.Println(res.Url, "---", res.Up)
				case <-done:
					return
				}
			}
		},
	}

	cmd.Flags().StringVarP(&src, "src", "s", "", "Read URLs from file")
	cmd.Flags().Int32VarP(&concurrency, "concurrency", "c", DefaultNumOfConcurrentRequests, "Number of concurrent requests")

	return cmd
}
