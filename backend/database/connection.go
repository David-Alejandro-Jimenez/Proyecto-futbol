package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var DataBase *sql.DB

func InitDB() error {
	var err error
	//err = godotenv.Load("../internal/config/.env")
	//if err != nil {
		//log.Fatal("Error al cargar el archivo .env:", err)
	//}

	var user = viper.GetString("DB_USER")
	var password = viper.GetString("DB_PASSWORD")
	var host = viper.GetString("DB_HOST")
	var port = viper.GetString("DB_PORT")
	var database = viper.GetString("DB_NAME")

	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true", user, password, host, port, database)
	
	DataBase, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("no se pudo conectar a la base de datos: %w", err)
	}

	err = DataBase.Ping() 
	if err != nil {
        return fmt.Errorf("no se pudo verificar la conexión: %w", err)
    }

	log.Println("Conexión a la base de datos exitosa")
	return nil
}