/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	// "fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		marxQuotesJSON, err := os.ReadFile("./bin/marxAndEngels.json")
		if err != nil {
			log.Fatal("Error when reading json file: ", err)
		}

		marxQutoesString := string(marxQuotesJSON[:])

		value := gjson.Get(marxQutoesString, "marx.#.quote")
		valueJson := value.Array()

		log.Fatal(valueJson[2])

		// log.Fatal(marxQutoesString)

		// const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

		// value := gjson.Get(json, "name.last")
		// println(value.String())

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
