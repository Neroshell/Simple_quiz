package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "fasttrack-quiz",
    Short: "A simple quiz application",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Welcome to the Fast Track Quiz!")
        // Call the StartQuiz function from quiz.go
        StartQuiz()
    },
}

// Execute adds all child commands to the root command and sets flags appropriately
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
