package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	user "./user"
)

var rootCmd = &cobra.Command {
	Use: "userCRUD",
	Short: "add, read, update user",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("user management tool")
	},
}

func main() {
	rootCmd.AddCommand(user.AddCmd())
	log.Fatal(rootCmd.Execute())
}