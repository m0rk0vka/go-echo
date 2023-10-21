package service

import (
	"encoding/json"
	"io/ioutil"
)

type Data struct {
	UserId int
	id     int
	Title  string
	Body   string
}

type Payload struct {
	Data []Data
}

func raw() ([]Data, error) {
	b, err := ioutil.ReadFile("data.json")
	if err != nil {
		return nil, err
	}
	var payload Payload
	if err := json.Unmarshal(b, &payload.Data); err != nil {
		return nil, err
	}
	return payload.Data, nil
}

func GetAll() ([]Data, error) {
	data, err := raw()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetById(id int) (any, error) {
	data, err := raw()
	if err != nil {
		return Data{}, err
	}
	if id < 0 || id > len(data) {
		return []string{}, err
	}
	return data[id], nil
}
