package helpers

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//ParseConfiguration ...
func ParseConfiguration(filePath string, ServiceConfig interface{}) {
	//ctx := context.Background()

	// dir, _ := os.Getwd()
	// filePath := dir + path

	yamlDataByte, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(filePath, " Error! Failed to read yaml[ParseConfiguration]:", err)
		//log.Error(ctx, filePath, "\n Error! Failed to read yaml[ParseConfiguration]:", err)
		return
	}
	//fmt.Println(filePath, " ", string(yamlFile))
	err = yaml.Unmarshal(yamlDataByte, ServiceConfig)
	if err != nil {
		fmt.Println(filePath, " Error! json unmarshal[ParseConfiguration]", err)
		//log.Error(ctx, filePath, "\n Error! json unmarshal[ParseConfiguration]", err)
		return
	}
	//fmt.Println(filePath, " Configuration for Service::", ServiceConfig)
}
