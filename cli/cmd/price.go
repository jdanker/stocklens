/*
Copyright © 2025 Jahred Danker jahrede@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type priceResponse struct {
	Ticker string  `json:"ticker"`
	Price  float64 `json:"price"`
}

// priceCmd represents the price command
var priceCmd = &cobra.Command{
	Use:   "price [ticker]",
	Short: "Returns the current price of the stock",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ticker := args[0]
		var resp priceResponse
		if err := fetchJSON(priceURL(ticker), &resp); err != nil {
			return err
		}
		fmt.Fprintf(cmd.OutOrStdout(), "%s price: %.2f\n", resp.Ticker, resp.Price)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(priceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// priceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// priceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
