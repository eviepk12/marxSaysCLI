/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Prints out a random quote from marx, engels, or both!",
	Long: 
	`
Prints out random quotations from marx, engels, or both.

Usage :
marxsays random
	`,
	Run: func(cmd *cobra.Command, args []string) {

		marxQuotesJSON, err := os.ReadFile("./bin/marxAndEngels.json")
		if err != nil {
			log.Fatal("Error when reading json file: ", err)
		}

		marxQutoesString := string(marxQuotesJSON[:])

		value := gjson.Get(marxQutoesString, "marx.#.quote")
		valueJson := value.Array()

		randomQuoteIndex := rand.Intn(100)

		fmt.Printf("Marx once said... \n %q", valueJson[randomQuoteIndex])
		fmt.Println("")
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
