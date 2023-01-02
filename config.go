package goconfig

import (
	"fmt"
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

func ParseYml[T interface{}](filepath string) *T {
	config := new(T)
	file, err := os.Open(filepath)
	if err != nil {
		err = fmt.Errorf("goconfig: %w", err)
		log.Fatalln(err)
	}
	if err := yaml.NewDecoder(file).Decode(config); err != nil {
		err = fmt.Errorf("goconfig: %w", err)
		log.Fatalln(err)
	}
	return config
}
