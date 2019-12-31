package persistence

import (
	"log"
	"time"

	"github.com/kmdkuk/my-blog-go/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kelseyhightower/envconfig"
)

var (
	DB  *gorm.DB
	env Env
)

type Env struct {
	User     string `envconfig:"MYSQL_USER" default:"default"`
	Password string `envconfig:"MYSQL_PASSWORD" default:"default"`
	Port     string `envconfig:"DB_PORT" default:"3306"`
	URL      string `envconfig:"DB_URL" default:"localhost"`
}

func init() {
	envconfig.Process("", &env)
	// Connection String
	DBMS := "mysql"
	USER := env.User
	PASS := env.Password
	PROTOCOL := "tcp(" + env.URL + ":" + env.Port + ")"
	DBNAME := "blog"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	// Create Connection
	var err error
	for range time.Tick(1 * time.Second) {
		DB, err = gorm.Open(DBMS, CONNECT)
		if err != nil {
			log.Println(err)
			continue
		}
		break
	}
	DB.LogMode(true)
	DB.Set("gorm:table_options", "ENGINE=InnoDB")
	DB.AutoMigrate(&model.Article{})
}
