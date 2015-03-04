package main

import (
	"fmt"

	"github.com/fortytw2/barrage/config"
	"github.com/fortytw2/barrage/importer"
	"github.com/fortytw2/barrage/models"
	"github.com/spf13/cobra"
)

//go:generate bash -c "lessc assets/less/barrage.less | cleancss > static/css/barrage.min.css"
//go:generate bash -c "browserify -t mithrilify assets/js/router.js > static/js/barrage.min.js"

func main() {
	var webCmd = &cobra.Command{
		Use:   "web",
		Short: "runs a webserver",
		Run: func(cmd *cobra.Command, args []string) {
			runWeb()
		},
	}

	var scrapeCmd = &cobra.Command{
		Use:   "scrape",
		Short: "attempts to generate toml from a name",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("no scraping yet")
		},
	}

	var importCmd = &cobra.Command{
		Use:   "import",
		Short: "manual reload of database",
		Run: func(cmd *cobra.Command, args []string) {
			db := models.OpenDB()
			importer.Import(config.VideoFolder, db)
		},
	}

	var rootCmd = &cobra.Command{Use: "barrage"}
	rootCmd.AddCommand(webCmd)
	rootCmd.AddCommand(scrapeCmd)
	rootCmd.AddCommand(importCmd)
	rootCmd.Execute()
}
