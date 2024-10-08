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
	Short: "Prints out a random quote from Karl Marx!.",
	Long: 
	`
Prints out random quotations from marx

Usage :
marxsays random
	`,
	Run: func(cmd *cobra.Command, args []string) {

		marxQuotesJSON, err := os.ReadFile("./bin/marxAndEngels.json")
		if err != nil {
			log.Fatal("Error when reading json file: ", err)
		}

		marxQuotesString := string(marxQuotesJSON[:])

		valueMarxQuote := gjson.Get(marxQuotesString, "marx.#.quote")
		valueMarxJSON := valueMarxQuote.Array()

		valueMarxQuoteSource := gjson.Get(marxQuotesString, "marx.#.source")
		valueMarxSourceJSON := valueMarxQuoteSource.Array()

		randomQuoteIndex := rand.Intn(143)

		fmt.Printf("_____")
		fmt.Printf("\n Cow Marx once said... \n \n %q",valueMarxJSON[randomQuoteIndex])
		fmt.Print("\n \n",valueMarxSourceJSON[randomQuoteIndex])
		fmt.Println("\n _____")
	
		cowAscii := 
		`
		   //
		  // -
		 //
		//
	 (    )
	  (oo)
)\.-----/(O O)
# ;       / u
   (  .   |} )
    |/ '.;|/;
    "     " "
		`

		fmt.Println(cowAscii)
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
