package config

import (
	"fmt"
	"os"

	env "github.com/caarlos0/env/v10"
	"gopkg.in/yaml.v3"
)

// ReadConfig reads configuration from file or from env variables if configFile set to "none"
// use envPrefix,yaml tags like `envPrefix:"DB_" yaml:"storage"`.
func ReadConfig(configFile string, cfg any) error {
	if configFile != "" && configFile != "none" {
		err := parseYaml(configFile, cfg)
		if err != nil {
			return err
		}
	} else {
		err := parseEnv(cfg)
		if err != nil {
			return err
		}
	}
	return nil
}

func parseEnv(cfg any) error {
	err := env.Parse(cfg)
	if err != nil {
		return err
	}

	return nil
}

//nolint:forbidigo
func PrintEnv(cfg any) error {
	params, err := env.GetFieldParams(cfg)
	if err != nil {
		return err
	}

	space := "    "
	fmt.Println("list of env")
	for _, p := range params {
		fmt.Println(space + p.Key)
	}

	fmt.Println()
	fmt.Println("env for cluster deploy format")
	for i, p := range params {
		fmt.Println(space + p.Key + ":" + fmt.Sprintf(" val%d", i))
	}

	fmt.Println()
	fmt.Println("env for docker compose deploy format")
	for i, p := range params {
		fmt.Println(space + p.Key + "=" + fmt.Sprintf("val%d", i))
	}

	return nil
}

func parseYaml(file string, cfg any) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	yamlDecoder := yaml.NewDecoder(f)
	err = yamlDecoder.Decode(cfg)
	if err != nil {
		return err
	}

	return nil
}

func CreateYaml(outputFileName string, cfg any) error {
	f, err := os.Create(outputFileName)
	if err != nil {
		return err
	}

	defer f.Close()

	yamlDecoder := yaml.NewEncoder(f)
	err = yamlDecoder.Encode(cfg)
	if err != nil {
		return err
	}
	return nil
}
