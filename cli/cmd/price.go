/*
Copyright © 2025 Jahred Danker jahrede@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// priceCmd represents the price command
var priceCmd = &cobra.Command{
	Use:   "price",
	Short: "Returns the current price of the stock",
	Long: `Returns the current price of the ticker symbol provided. 
	Price is last known price during standard market hours`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("price called")
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
