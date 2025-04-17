package cmd

import (
	"fmt"
	"log"
	"os"
	"uooobarry/soup/internal/model"
	"uooobarry/soup/internal/soup"

	"github.com/spf13/cobra"
	"gorm.io/datatypes"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch soup data and store it in the database",
	Long:  `Fetches soup data from the DeepSeek API and stores it in the SQLite database.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := soup.FetchSoupFromDS(); err != nil {
			log.Println("Failed to fetch soup: %v", err)
		}
		fmt.Println("Successfully fetched and stored soup data.")
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new soup entry",
	Long:  `Creates a new soup entry with the provided question, answer, and tags.`,
	Run: func(cmd *cobra.Command, args []string) {
		var soupEntry model.Soup
		fmt.Print("Enter Soup Question: ")
		fmt.Scanln(&soupEntry.SoupQuestion)
		fmt.Print("Enter Soup Answer: ")
		fmt.Scanln(&soupEntry.SoupAnswer)
		fmt.Print("Enter Soup Tags (JSON): ")
		fmt.Scanln(&soupEntry.SoupTag)

		if err := model.CreateSoup(&soupEntry); err != nil {
			log.Printf("Failed to create soup: %v", err)
			return
		}
		fmt.Printf("Successfully created soup with ID: %d\n", soupEntry.ID)
	},
}

var readCmd = &cobra.Command{
	Use:   "read [id]",
	Short: "Read a soup entry by ID",
	Long:  `Fetches and displays a soup entry by its ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var soupEntry model.Soup
		if err := model.GetSoupByID(args[0], &soupEntry); err != nil {
			log.Printf("Failed to read soup: %v", err)
			return
		}
		fmt.Printf("Soup Question: %s\nSoup Answer: %s\nTags: %s\n", soupEntry.SoupQuestion, soupEntry.SoupAnswer, soupEntry.SoupTag)
	},
}

var updateCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Update a soup entry by ID",
	Long:  `Updates a soup entry with the provided question, answer, and tags.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var soupEntry model.Soup
		if err := model.GetSoupByID(args[0], &soupEntry); err != nil {
			log.Printf("Failed to fetch soup for update: %v", err)
			return
		}

		fmt.Printf("Current Question: %s\nNew Question (leave blank to keep current): ", soupEntry.SoupQuestion)
		var newQuestion string
		fmt.Scanln(&newQuestion)
		if newQuestion != "" {
			soupEntry.SoupQuestion = newQuestion
		}

		fmt.Printf("Current Answer: %s\nNew Answer (leave blank to keep current): ", soupEntry.SoupAnswer)
		var newAnswer string
		fmt.Scanln(&newAnswer)
		if newAnswer != "" {
			soupEntry.SoupAnswer = newAnswer
		}

		fmt.Printf("Current Tags: %s\nNew Tags (JSON, leave blank to keep current): ", soupEntry.SoupTag)
		var newTags string
		fmt.Scanln(&newTags)
		if newTags != "" {
			soupEntry.SoupTag = datatypes.JSON(newTags)
		}

		if err := model.UpdateSoup(&soupEntry); err != nil {
			log.Printf("Failed to update soup: %v", err)
			return
		}
		fmt.Println("Successfully updated soup.")
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a soup entry by ID",
	Long:  `Deletes a soup entry by its ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := model.DeleteSoup(args[0]); err != nil {
			log.Printf("Failed to delete soup: %v", err)
			return
		}
		fmt.Println("Successfully deleted soup.")
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(readCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(deleteCmd)
}

var rootCmd = &cobra.Command{
	Use:   "soup-cli",
	Short: "A CLI tool for managing soup data",
	Long:  `soup-cli is a command-line tool to interact with the soup database and APIs.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
