package app

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type T struct {
	Version   string   `yaml:"version"`
	Path      *string  `yaml:"path,omitempty"`
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

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	fmt.Println(t.Path)

	if t.Path != nil {
		if err := os.Chdir(*t.Path); err != nil {
			log.Fatal(err)
		}
	}

	for _, f := range t.Structure {
		path := fmt.Sprintf("%s/%s", t.Output, f)
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			log.Fatal(err)
		}

		if err := os.Mkdir(path, os.ModeDir); err != nil {
			log.Fatal(err)
		}

		log.Println("> Path:", path)
	}
}
