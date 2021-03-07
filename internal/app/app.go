package app

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
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
	bad := color.New(color.FgRed).PrintlnFunc()
	info := color.New(color.FgCyan).PrintlnFunc()
	log := color.New(color.FgHiCyan).PrintlnFunc()
	ok := color.New(color.FgHiGreen).PrintlnFunc()

	t := T{}

	err := yaml.Unmarshal(data, &t)
	if err != nil {
		bad("error: %v", err)
		return
	}
	//fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		bad(err.Error())
		log("error: %v", err)
		return
	}

	info(string(d))

	if t.Path != nil { // Path is defined
		// let check if the path exists
		if _, err := os.Stat(*t.Path); os.IsNotExist(err) {
			info(fmt.Printf("Path does not exists: %s", *t.Path))
			if err := os.MkdirAll(*t.Path, 0777); err != nil {
				bad(fmt.Printf("\tError creating path: %s", *t.Path))
				bad(err.Error())
				os.Exit(1)
			}
			info(fmt.Printf("Path created: %s", *t.Path))
		}

		if err := os.Chdir(*t.Path); err != nil {
			bad(fmt.Printf("Cannot change to path: %s", *t.Path))
			bad(err.Error())
			os.Exit(1)
		}
	} else {
		// A path is not defined, assuming current directory
		// we can remove this else later on...
		current, _ := os.Getwd()
		t.Path = &current
		info(fmt.Printf("Path: %s", current))
	}

	for _, path := range t.Structure {
		log("Trying path:", path)
		if err := os.MkdirAll(path, 0777); err != nil {
			bad(err.Error())
		}
		ok("\tCreated:", path)
	}

	c, _ := os.Getwd()
	log("Final structure", c)
	cmd := exec.Command("tree")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		bad(err.Error())
	}
}
