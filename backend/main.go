package main

import (
	"fmt"
	"log"
	"os"

	_ "fs-beego/routers"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

func init() {
	env := os.Getenv("ENV")

	if env == "" {
		env = "dev"
	}

	if err := orm.RegisterDriver("postgres", orm.DRPostgres); err != nil {
		log.Fatal(err.Error())
	}

	host, _ := config.String(env + "::db_host")
	port, _ := config.String(env + "::db_port")
	user, _ := config.String(env + "::db_user")
	pass, _ := config.String(env + "::db_pass")
	name, _ := config.String(env + "::db_name")

	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, name)

	if err := orm.RegisterDataBase("default", "postgres", dataSource); err != nil {
		log.Fatal(err.Error())
	}

	if err := orm.RunSyncdb("default", false, true); err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	beego.Run()
}
