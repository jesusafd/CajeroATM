package database

import (
	_ "embed"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
)

// configuracionBD es la estructura usitlizada para
// extraer los datos necesarios para la conexion con
// la BD, del archivo configuracion.yaml
type SettingsDB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

// FileSettings es la varible con la cual se extraera los
// datos del archivo .yaml

//go:embed settings.yaml
var FileSettings []byte

// Set es la funcion encargada de leer los datos
// del archivo .yaml y crear un objeto de tipo ConfiguracionBD el
// cula permita manipualar estos datos
func SetSettings() (SettingsDB, error) {
	var settingsDB SettingsDB

	err := yaml.Unmarshal(FileSettings, &settingsDB)
	if err != nil {
		log.Println("Error al extraer los datos de configuracion de la bd" + err.Error())
	}
	return settingsDB, err
}
