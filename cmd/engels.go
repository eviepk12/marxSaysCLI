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

// engelsCmd represents the engels command
var engelsCmd = &cobra.Command{
	Use:   "engels",
	Short: "Outputs random quotes from Friedrich Engels!",
	Long: `
Outputs random quotes from Friedrich Engels!

Usage :
marxsays engels
	`,
	Run: func(cmd *cobra.Command, args []string) {
		engelsQuotesJSON, err := os.ReadFile("./bin/marxAndEngels.json")
		if err != nil {
			log.Fatal("Error when reading json file: ", err)
		}

		engelsQutoesString := string(engelsQuotesJSON[:])

		valueEngelsQuote := gjson.Get(engelsQutoesString, "engels.#.quote")
		valueEngelsJSON := valueEngelsQuote.Array()

		valueMarxQuoteSource := gjson.Get(engelsQutoesString, "engels.#.source")
		valueMarxSourceJSON := valueMarxQuoteSource.Array()

		randomQuoteIndex := rand.Intn(42)

		fmt.Printf("_____")
		fmt.Printf("\n Cow Engels once said... \n \n %q",valueEngelsJSON[randomQuoteIndex])
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
	rootCmd.AddCommand(engelsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// engelsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// engelsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
