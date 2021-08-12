package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
	"github.com/tidwall/gjson"
)

// Defining mutation function
func mutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {
	data, err := ioutil.ReadFile("validation.json")
	rootNode := gjson.ParseBytes(data)

	if err != nil {
		log.Fatal("Error while loading file", err)
	}

	for _, model := range b.Models {
		var childNode gjson.Result
		if childNode = rootNode.Get(model.Name); childNode.Exists() {
			for _, field := range model.Fields {
				var node gjson.Result
				if node = childNode.Get(field.Name); node.Exists() {
					if node.IsObject() {
						field.Tag += ` govalid:"`
						node.ForEach(func(key, value gjson.Result) bool {
							if value.String() == "" {
								field.Tag += key.String() + "|"
							} else {
								field.Tag += key.String() + ":" + value.String() + "|"
							}

							return true
						})
						field.Tag += `"`
					} else {
						field.Tag += ` govalid:"` + node.String() + `"`
					}
				}
			}
		}

	}

	return b
}

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	// Attaching the mutation function onto modelgen plugin
	p := modelgen.Plugin{
		MutateHook: mutateHook,
	}

	err = api.Generate(cfg,
		api.NoPlugins(),
		api.AddPlugin(&p),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
