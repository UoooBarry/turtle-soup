package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"uooobarry/soup/config"
	"uooobarry/soup/internal/model"
	"uooobarry/soup/internal/repository"
	"uooobarry/soup/internal/service"
	"uooobarry/soup/internal/soup"

	"github.com/spf13/cobra"
	"gorm.io/datatypes"
)

func parseID(arg string) (uint, error) {
	id, err := strconv.ParseUint(arg, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid ID format: %v", err)
	}
	return uint(id), nil
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch soup data and store it in the database",
	Long:  `Fetches soup data from the DeepSeek API and stores it in the SQLite database.`,
	Run: func(cmd *cobra.Command, args []string) {
		s := getSoupService()
		if err := soup.FetchSoupFromDS(s); err != nil {
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

		service := getSoupService()
		if err := service.Create(&soupEntry); err != nil {
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
		service := getSoupService()
		id, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			log.Printf("Invalid ID format: %v", err)
			return
		}
		soupEntry, err := service.GetByID(uint(id))
		if err != nil {
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
		s := getSoupService()
		id, err := parseID(args[0])
		if err != nil {
			log.Printf("Failed to parse ID: %v", err)
			return
		}
		soupEntry, err := s.GetByID(id)
		if err != nil {
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

		if err := s.Update(soupEntry); err != nil {
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
		s := getSoupService()
		id, err := parseID(args[0])
		if err != nil {
			log.Printf("Failed to parse ID: %v", err)
			return
		}
		if err := s.Delete(id); err != nil {
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
	rootCmd.AddCommand(registerCmd)
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

func getSoupService() *service.SoupService {
	repo := repository.NewSoupRepository(config.InitDB())
	return service.NewSoupService(repo)
}

func getAuthService() *service.AuthService {
	repo := repository.NewAuthRepository(config.InitDB())
	return service.NewAuthService(repo)
}

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a new user",
	Long:  `Registers a new user with the provided username, password, and email.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			username string
			password string
			email    string
		)

		fmt.Print("Enter Username: ")
		fmt.Scanln(&username)
		fmt.Print("Enter Password: ")
		fmt.Scanln(&password)
		fmt.Print("Enter Email: ")
		fmt.Scanln(&email)

		authService := getAuthService()
		user, err := authService.Register(username, password, email)
		if err != nil {
			log.Printf("Failed to register user: %v", err)
			return
		}

		fmt.Printf("Successfully registered user with ID: %d\n", user.ID)
	},
}
