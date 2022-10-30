package migration

import (
	"fmt"
	"mnc-bank-api/model"
	"mnc-bank-api/utils"
	"mnc-bank-api/utils/jsonrw"
	"os"
	"time"
)

func Migrate() {
	utils.Line()
	fmt.Println("Migrating (creating required files only)")
	defer fmt.Println("migration finished")

	if !fileExists("database") {
		err := os.Mkdir("./database", os.FileMode(0777))
		if err != nil {
			panic(err)
		}
	}

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

		if table == "customer" {
			os.Truncate(filePath("customer"), 0)
			AddDummyCustomer()
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

func AddDummyCustomer() {

	customer1 := model.Customer{
		Id:       utils.GenerateId(),
		Name:     "Andi",
		Username: "andi",
		Password: "5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8",
		Balance:  500000,
	}

	customer2 := model.Customer{
		Id:       utils.GenerateId(),
		Name:     "Budi",
		Username: "budi",
		Password: "5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8",
		Balance:  500000,
	}

	jsonrw.JsonWriteData("customer", customer1)
	jsonrw.JsonWriteData("customer", customer2)
}

func AddDummyTransaction() {

	tx1 := model.Transaction{
		Id:         "test",
		SenderId:   "test",
		ReceiverId: "test",
		Amount:     5000,
		Created_at: time.Now(),
	}

	tx2 := model.Transaction{
		Id:         "test",
		SenderId:   "test",
		ReceiverId: "test",
		Amount:     85858,
		Created_at: time.Now(),
	}

	jsonrw.JsonWriteData("transaction", tx1)
	jsonrw.JsonWriteData("transaction", tx2)
}
