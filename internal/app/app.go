package app

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type T struct {
	Version   string   `yaml:"version"`
	Path      string   `yaml:"path"`
	Name      string   `yaml:"name"`
	Output    string   `yaml:"output"`
	Structure []string `yaml:"structure"`
}

func Parse(data []byte) {
	t := &T{}

	err := yaml.Unmarshal(data, &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//fmt.Printf("--- t:\n%v\n\n", t)

	_, err = yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//fmt.Printf("--- t dump:\n%s\n\n", string(d))

	for _, f := range t.Structure {
		path := fmt.Sprintf("%s/%s", t.Output, f)
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			log.Fatal(err)
		}
		log.Println("> Path:", path)
	}
}
