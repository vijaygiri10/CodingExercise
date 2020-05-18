package lib

import (
	util "CodingExercise/shared/helpers"
	"CodingExercise/shared/log"
	"fmt"
	"os"
)

func ParseConfiguration(configPath, logPath string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Execption ParseConfiguration err: ", err)
		}
	}()

	servicefilePath := configPath + "/service.yaml"
	var service ServiceConfiguration
	util.ParseConfiguration(servicefilePath, &service)

	postgresfilePath := configPath + "/postgres.yaml"
	var postgres PostgresConfiguration
	util.ParseConfiguration(postgresfilePath, &postgres)

	env := os.Getenv("ENV")
	ServiceConfig = Configuration{
		Service:  service.Env[env],
		Postgres: postgres.Env[env],
	}

	log.LogName = ServiceConfig.Service.LogName
	fmt.Println(env, " service: ", service.Env[env])
	fmt.Println(env, " postgres: ", postgres.Env[env])

	log.InitializeLogger(logPath)

}
