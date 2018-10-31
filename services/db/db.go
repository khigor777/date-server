package db

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Db struct {
	fileName string
}

func (db *Db) Set(timeF float64) error {
	return ioutil.WriteFile(db.fileName, []byte(fmt.Sprintf("%09.6f", timeF)), os.ModePerm)
}

func (db *Db) Get() (float64, error) {
	b, err := ioutil.ReadFile(db.fileName)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(string(b), 64)
}

func (db *Db) Remove() error {
	return os.Remove(db.fileName)
}
