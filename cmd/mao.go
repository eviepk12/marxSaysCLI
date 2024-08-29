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

// maoCmd represents the mao command
var maoCmd = &cobra.Command{
	Use:   "mao",
	Short: "Outputs random quotes from Mao Zedong",
	Long: `
Outputs random Mao Quotes

Usage :
marxsays mao
	`,

	Run: func(cmd *cobra.Command, args []string) {
		maoQuotesJSON, err := os.ReadFile("./bin/mao.json")
		if err != nil {
			log.Fatal("Error when reading json file: ", err)
		}

		maoQuotesString := string(maoQuotesJSON[:])

		valueMaoQuote := gjson.Get(maoQuotesString, "mao.#.quote")
		valueMaoJSON := valueMaoQuote.Array()

		valueMaoQuoteSource := gjson.Get(maoQuotesString, "mao.#.source")
		valueMaoSourceJSON := valueMaoQuoteSource.Array()

		randomQuoteIndex := rand.Intn(147)

		fmt.Printf("_____")
		fmt.Printf("\n Cow Mao once said... \n \n %q",valueMaoJSON[147])
		fmt.Print("\n \n",valueMaoSourceJSON[randomQuoteIndex])
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
	rootCmd.AddCommand(maoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// maoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// maoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
