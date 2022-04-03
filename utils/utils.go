package utils

import (
	"fmt"
	"log"
)

/*
##########
# Print Message
##########
*/
func PrintMsg(msg string) {
	log.Println(fmt.Sprintf(": %v : msg='%v'", "info", msg))
}

/*
##########
# Print Error is any
##########
*/
func CheckErr(err error) {
	if err != nil {
		log.Println(fmt.Sprintf(": %v : msg='%v'", "error", err.Error()))
	}
}

/*
##########
# Panic if Error is found
##########
*/
func PanicIfErr(err error) {
	CheckErr(err)
	if err != nil {
		panic(err)
	}
}

// Function to get config variables
// should be using config file
func GetCfgVar(envVar string) string {
	localValueMap := map[string]string{
		"DB_IP":       "localhost",
		"DB_USER":     "raktim",
		"DB_PASSWORD": "12345678",
		"TEST_IP":     "localhost",
	}
	dockerValueMap := map[string]string{
		"DB_IP":       "postgres_app",
		"DB_USER":     "raktim",
		"DB_PASSWORD": "12345678",
		"TEST_IP":     "ipatser_app_deploy",
	}
	absValueMap := map[string]map[string]string{
		"local":  localValueMap,
		"docker": dockerValueMap,
	}
	valueMap := absValueMap["local"]
	envVarValue := ""
	if _, ok := valueMap[envVar]; ok {
		envVarValue = valueMap[envVar]
	}
	return envVarValue
}
