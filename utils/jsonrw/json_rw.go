package jsonrw

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

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

func JsonUpdateList(tableName string, newList interface{}) error {
	if newList == nil {
		return errors.New("aborted, no new list received")
	}

	tablePath := "database/" + tableName + ".json"

	err := os.Truncate(tablePath, 0)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(tablePath, []byte(JsonEncode(newList)), os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
