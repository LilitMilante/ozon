package app

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port int `yml:"Port"`

	Database DBConfig
}

type DBConfig struct {
	Host     string `yml:"DB_Host"`
	Port     int    `yml:"DB_Port"`
	Name     string `yml:"DB_Name"`
	User     string `yml:"DB_User"`
	Password string `yml:"DB_Password"`
}

func NewConfig() (Config, error) {
	data, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	var conn Config

	err = yaml.Unmarshal(data, &conn)
	if err != nil {
		log.Fatalf("error unmarshaling YAML: %v", err)
	}

	return conn, nil

}
