package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var plusCmd = &cobra.Command{
	Use:   "plus",
	Short: "Just a placeholder",
	Long:  `Just a placeholder`,

	Args: cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {
		num1, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid argument:", args[0])
			return
		}

		num2, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid argument:", args[1])
			return
		}

		fmt.Println(num1 + num2)
	},
}

func init() {
	RegisterSubCommand(plusCmd)
}
