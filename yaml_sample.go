package main

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// structたち
type Data struct {
	Users []User `yaml:"users"`
}

type User struct {
	Name             string           `yaml:"common"`
	FullName         fullName         `yaml:"full_name"`
	Sex              string           `yaml:"sex"`
	SelfIntroduction selfIntroduction `yaml:"self_introduction"`
	ImageURLs        []string         `yaml:"image_urls"`
	Shemale          bool             `yaml:"shemale"`
}

type fullName struct {
	FirstName string `yaml:"first_name"`
	LastName  string `yaml:"last_name"`
}

type selfIntroduction struct {
	Long  string `yaml:"long"`
	Short string `yaml:"short"`
}

func main() {
	// yamlを読み込む
	buf, err := ioutil.ReadFile("./sample.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("buf: %+v", string(buf))
	fmt.Printf("\n\n")

	// structにUnmasrshal
	var d Data
	err = yaml.Unmarshal(buf, &d)
	if err != nil {
		panic(err)
	}
	fmt.Printf("d: %+v", d)
	fmt.Printf("\n\n")

}
