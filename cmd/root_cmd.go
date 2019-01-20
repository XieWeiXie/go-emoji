package cmd

import (
	"fmt"
	"go-emoji/cmd/emoji"
	"go-emoji/models"
	"go-emoji/pkg/database"
	"log"
	"os"

	"github.com/jinzhu/gorm"

	"github.com/spf13/cobra"
)

const (
	AppName = "emoji"
)

var rootCMD = &cobra.Command{
	Use: AppName,
	Run: fetch,
}

func fetch(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		helper()
		return
	}
	database.DataBaeInit()
	defer database.POSTGRESDIALECT.Close()
	switch args[0] {
	case "import":
		if err := importEmoji(); err != nil {
			panic("import emoji data error")
		}
	case "export":
		if err := exportEmoji(); err != nil {
			panic("export emoji data error")
		}
	case "db":
		log.Println(args)
		switch args[1] {
		case "migrate":
			migrateTable(database.POSTGRESDIALECT)
		}
	}

}

func helper() string {
	return "You should add one argument at least"
}

func importEmoji() error {
	fetcher := emoji_fetch.NewFetcher("https://emojipedia.org/emoji/")
	fetcher.Run()
	return nil
}

func exportEmoji() error {
	path := "./src/emoji_map.go"
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		os.Create(path)
	}

	return nil
}

func migrateTable(dialect *gorm.DB) {
	for _, i := range models.Collection() {
		fmt.Println("---Migrate:")
		dialect.AutoMigrate(i)
	}
	var version models.Version
	if dbError := dialect.Where("version_name = ?", "12.0").First(&version).Error; dbError != nil {
		version.VersionName = "12.0"
		dialect.Save(&version)
	}
}

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
