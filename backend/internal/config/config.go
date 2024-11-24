package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() error {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("internal/config")

	var err = viper.ReadInConfig()
	if err != nil {
		log.Printf("Error al cargar la configuración: %v", err)
		return err
	}

	viper.AutomaticEnv()
	log.Println("Archivo .env cargado correctamente")
	return nil
}