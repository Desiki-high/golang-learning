package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	App App `yaml:"app" json:"app"`
	Log Log `yaml:"log" json:"log"`
}

type App struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

type Log struct {
	Suffix  string `yaml:"suffix" json:"suffix"`
	MaxSize int    `yaml:"maxSize" json:"maxsize"`
}

func main() {
	yamlFile, err := os.ReadFile("struct-yaml/config.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}

	cfg := Config{}
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- cfg: %v\n\n", cfg)

	d, err := yaml.Marshal(&cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- cfg dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal(yamlFile, &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))
}
