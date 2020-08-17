package command

import (
	"fmt"
	"io/ioutil"

	schema "github.com/chathula/workspacer/app"
	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
)

var ws schema.WorkspacerSchema

// Install all
func Install(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("loading failed file: <%s>", color.RedString(file))
	}

	var workspacerSchema schema.WorkspacerSchema
	err = yaml.Unmarshal([]byte(data), &workspacerSchema)
	if err != nil {
		return fmt.Errorf(color.RedString("Failed parsing yaml file"))
	}

	schema := schema.WorkspacerSchema{
		Version: workspacerSchema.Version,
		Home:    workspacerSchema.Home,
		Git:     workspacerSchema.Git,
		Envs:    workspacerSchema.Envs,
	}

	fmt.Println(schema.Envs[0].Key)

	// user, _ := user.Current()
	// fmt.Printf("install all file %s\n", file)
	// fmt.Printf("home %s\n", user.HomeDir)
	// cmd := exec.Command("/bin/sh", "-c", "sudo touch /etc/testgofile")
	// fmt.Printf("cmd %s\n", cmd)
	return nil
}
