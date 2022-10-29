package try

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mnc-bank-api/model"
	"mnc-bank-api/utils"
	"os"
	"strconv"
	"strings"
)

func Run() {
	var err error

	// for i := 0; i < 20; i++ {
	// 	err = JsonWriteData("customer", newCustomer(i))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	all, err := JsonReadData("customer")
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("all")
	fmt.Println(all[0])

}

func newCustomer(num int) model.Customer {
	return model.Customer{
		Id:       utils.GenerateId(),
		Name:     "cust" + strconv.Itoa(num),
		Username: "uname" + strconv.Itoa(num),
		Password: "12345",
		Balance:  num * 1450,
	}
}

func JsonEncode(model interface{}) string {
	jsonData, _ := json.Marshal(model)
	return string(jsonData)
}

func JsonReadData(tableName string) ([]interface{}, error) {
	var list []interface{}

	file, err := ioutil.ReadFile("database/" + tableName + ".json")
	if err != nil {
		return nil, errors.New("unable to read json file from table " + tableName + " : " + err.Error())
	}

	json.Unmarshal(file, &list)
	return list, nil
}

func JsonWriteData(tableName string, model interface{}) error {
	list, _ := JsonReadData(tableName)
	list = append(list, model)
	jsonByte, _ := json.Marshal(list)

	err := ioutil.WriteFile("database/"+tableName+".json", jsonByte, os.ModePerm)
	if err != nil {
		return errors.New("cannot write json data into table " + tableName + " : " + err.Error())
	}

	return nil
}
