package config

import (
	"encoding/json"

	"os"
	"path"
	"strings"

	"github.com/Jarover/crf/pkg/utils"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type CopyPattern struct {
	Pattern     string `yaml:"pattern" json:"pattern"`
	To          string `yaml:"to" json:"to"`
	Type        string `yaml:"type" json:"type"`
	Move        bool   `yaml:"move" json:"move"`
	Achive      bool   `yaml:"archive" json:"archive"`
	ArchivePath string `yaml:"archive_path" json:"archive_path"`
}

// Config - структура для считывания конфигурационного файла
type Config struct {
	//Db_url       string `yaml:"db_url" json:"db_url"`
	//Port         int    `yaml:"port" json:"port" `
	//Noobject_url string `yaml:"noobject_url" json:"noobject_url"`
	LogLevel      string `yaml:"logLevel" json:"logLevel"`
	Env           bool   `yaml:"env" json:"env" `
	Save          bool   `yaml:"save" json:"save" `
	CryptPath     string `yaml:"crypt_path" json:"crypt_path"`
	UnzipPath     string `yaml:"unzip_path" json:"unzip_path"`
	ModifyTimeout int    `yaml:"modify_timeout" json:"modify_timeout" `
	MaxWork       int    `yaml:"max_work" json:"max_work"`
	/*
		CountMessages int    `yaml:"count_messages" json:"count_messages" `
		EmailPassword string `yaml:"email_password" json:"email_password"`
		EmailHost     string `yaml:"email_host" json:"email_host"`
		EmailUser     string `yaml:"email_user" json:"email_user"`
		EmailPort     string `yaml:"email_port" json:"email_port"`
	*/
	DbName string `yaml:"db_name" json:"db_name"`
	DbPort string `yaml:"db_port" json:"db_port" `
	DbHost string `yaml:"db_host" json:"db_host"`
	DbUser string `yaml:"db_user" json:"db_user"`
	DbPass string `yaml:"db_pass" json:"db_pass"`

	CopyRules []struct {
		Achive      bool   `yaml:"archive" json:"archive"`
		ArchivePath string `yaml:"archive_path" json:"archive_path"`
		Paths       []struct {
			Path     string `yaml:"path" json:"path"`
			Patterns []struct {
				CopyPattern
			} `yaml:"patterns" json:"patterns"`
		}
	} `yaml:"copyrules" json:"copyrules"`

	Rules []struct {
		From  string `yaml:"from" json:"from"`
		Paths []struct {
			InPath  string `yaml:"inPath" json:"inPath"`
			OutPath string `yaml:"outPath" json:"outPath"`
		} `yaml:"paths" json:"paths"`
	} `yaml:"rules" json:"rules"`
}

var conf *Config

func ReadConfig(ConfigName string) (err error) {
	var file []byte
	if file, err = os.ReadFile(ConfigName); err != nil {
		return err
	}

	switch strings.ToLower(path.Ext(ConfigName)) {

	case ".yaml", ".yml":
		err = yaml.Unmarshal(file, &conf)
	case ".json":
		err = json.Unmarshal(file, &conf)

	}

	if err != nil {
		return err
	}

	if conf.Env {
		dir, err := utils.GetDir()
		if err != nil {
			return err
		}
		e := godotenv.Load(dir + "/.env") //Загрузить файл .env
		if e != nil {
			return e
		}
		/*
				if os.Getenv("EMAIL_HOST") != "" {
					conf.EmailHost = os.Getenv("EMAIL_HOST")
				}

				if os.Getenv("EMAIL_PASSWORD") != "" {
					conf.EmailPassword = os.Getenv("EMAIL_PASSWORD")
				}

				if os.Getenv("EMAIL_USER") != "" {
					conf.EmailUser = os.Getenv("EMAIL_USER")
				}

				if os.Getenv("EMAIL_PORT") != "" {
					conf.EmailPort = os.Getenv("EMAIL_PORT")
				}

			if os.Getenv("DB_PORT") != "" {
				conf.DbPort = os.Getenv("DB_PORT")
			}
			if os.Getenv("DB_HOST") != "" {
				conf.DbHost = os.Getenv("DB_HOST")
			}
			if os.Getenv("DB_NAME") != "" {
				conf.DbName = os.Getenv("DB_NAME")
			}
			if os.Getenv("DB_USER") != "" {
				conf.DbUser = os.Getenv("DB_USER")
			}
			if os.Getenv("DB_PASS") != "" {
				conf.DbPass = os.Getenv("DB_PASS")
			}
		*/

	}

	//username := os.Getenv("DB_USER")

	return nil
}

// возвращает дескриптор объекта DB
func GetConfig() *Config {
	return conf
}
