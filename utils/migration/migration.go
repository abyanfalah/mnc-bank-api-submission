package migration

import (
	"fmt"
	"mnc-bank-api/utils"
	"os"
)

func Migrate() {
	utils.Line()
	fmt.Println("Migrating (creating required files only)")
	defer fmt.Println("migration finished")

	tables := []string{
		"customer",
		"transaction",
		"activity_log",
	}

	for _, table := range tables {
		if !fileExists(filePath(table)) {
			file, err := os.Create(filePath(table))
			if err != nil {
				panic(err)
			}

			file.Close()
		}
	}
}

func filePath(tableName string) string {
	return "database/" + tableName + ".json"
}

func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}
