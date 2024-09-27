package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

var Viper *viper.Viper

func InitViper() {

	env := "development"

	// Initialize Viper
	v := viper.New()

	if err := checkFileExists(env); err != nil {
		log.Println(err)
		log.Printf("Loaded From Enivronment Variables. APP_ENV: %s \n", env)
		v.AutomaticEnv()
	} else {
		log.Printf("File .env.%s exists\n", env)
		log.Printf("Loaded .env.%s file\n", env)

		// Set the configuration file based on the environment
		v.SetConfigFile(fmt.Sprintf(".env.%s", env))
		v.SetConfigType("dotenv")
		v.AddConfigPath("./")

		_ = v.ReadInConfig()
	}

	Viper = v
}

func checkFileExists(env string) error {
	fileName := fmt.Sprintf(".env.%s", env)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return fmt.Errorf("file %s does not exist", fileName)
	}
	return nil
}
