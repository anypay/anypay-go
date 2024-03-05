package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var invoiceCmd = &cobra.Command{
	Use:   "invoice",
	Short: "Manage invoices",
	Long:  `All invoice related actions can be accessed through this command.`,
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new invoice",
	Long:  `Create a new invoice with specified details.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Here you can handle the invoice creation logic
		fmt.Println("Invoice created with amount:", amount, "and currency:", currency)
	},
}

// cancel an invoice
var cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel an invoice",
	Long:  `Cancel an invoice with specified details.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Here you can handle the invoice cancellation logic
		fmt.Println("Invoice cancelled")
	},
}

// fetch an invoice
var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch an invoice",
	Long:  `Fetch an invoice with specified details.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Here you can handle the invoice fetching logic
		fmt.Println("Invoice fetched")
	},
}

var amount int
var currency string

func init() {
	// Add flags for the `create` command
	createCmd.Flags().IntVarP(&amount, "amount", "a", 0, "Amount for the invoice")
	createCmd.Flags().StringVarP(&currency, "currency", "c", "", "Currency for the invoice")
	createCmd.MarkFlagRequired("amount")
	createCmd.MarkFlagRequired("currency")

	// Add `create` as a sub-command of `invoice`
	invoiceCmd.AddCommand(createCmd)
}
