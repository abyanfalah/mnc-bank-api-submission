package migration

import (
	"fmt"
	"mnc-bank-api/model"
	"mnc-bank-api/utils"
	"mnc-bank-api/utils/jsonrw"
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

			if table == "customer" {
				addDummyCustomer()
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

func addDummyCustomer() {

	customer1 := model.Customer{
		Id:       utils.GenerateId(),
		Name:     "Andi",
		Username: "andi",
		Password: "password",
		Balance:  500000,
	}

	customer2 := model.Customer{
		Id:       utils.GenerateId(),
		Name:     "Budi",
		Username: "budi",
		Password: "password",
		Balance:  500000,
	}

	jsonrw.JsonWriteData("customer", customer1)
	jsonrw.JsonWriteData("customer", customer2)
}
