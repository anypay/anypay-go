package cmd

import (
	"fmt"

	"github.com/anypay/anypay-go/prices"
	"github.com/spf13/cobra"
)

// pricesCmd represents the base command when called without any subcommands
var pricesCmd = &cobra.Command{
	Use:   "prices",
	Short: "Manage prices",
	Long:  `All price related actions can be accessed through this command.`,
}

// listCmd represents the command to list all prices
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all prices",
	Long:  `List all available prices.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Here you can call the prices module to list prices
		fmt.Println("Listing all prices:")
		pricesList := prices.ListAll() // Assuming ListAll() is a function in your prices module
		for _, price := range pricesList {
			fmt.Printf("ID: %d, Amount: %f, Currency: %s\n", price.ID, price.Amount, price.Currency)
		}
	},
}

// showCmd represents the command to show a specific price by ID
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show a specific price",
	Long:  `Show details of a specific price by ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Here you can call the prices module to show a specific price
		if len(args) < 1 {
			fmt.Println("Please provide an ID for the price you want to show.")
			return
		}
		id := args[0]
		fmt.Printf("Showing details for price ID: %s\n", id)
		price, err := prices.FindByID(id) // Assuming FindByID() is a function in your prices module
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("ID: %d, Amount: %f, Currency: %s\n", price.ID, price.Amount, price.Currency)
	},
}

func init() {
	rootCmd.AddCommand(pricesCmd)
	pricesCmd.AddCommand(listCmd)
	pricesCmd.AddCommand(showCmd)
}
