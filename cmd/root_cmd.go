package cmd

import (
	"fmt"
	"go-emoji/pkg/database"
	"os"

	"github.com/spf13/cobra"
)

const (
	AppName = "go-emoji"
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
		switch args[1] {
		case "migrate":
			migrateTable(args[2])
		}
	}

}

func helper() string {
	return "You should add one argument at least"
}

func importEmoji() error {
	return nil
}

func exportEmoji() error {
	return nil
}

func migrateTable(dialect string) {

	for _, i := range database.ModelsCollection() {
		if dialect == "sqlite" {
			database.SQLITE3DIALECT.AutoMigrate(i)
		} else {
			database.POSTGRESDIALECT.AutoMigrate(i)
		}
	}
}

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
