package main

import (
	_ "blog/routers"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// set default database
	orm.RegisterDriver("postgres", orm.DRPostgres)

	// set default RegisterDataBasedatabase
	pg_user, _ := web.AppConfig.String("PgUser")
	pg_pwd, _ := web.AppConfig.String("PgPassword")
	pg_host, _ := web.AppConfig.String("PgHost")
	pg_db, _ := web.AppConfig.String("PgDatabase")
	db_url := "postgres://" + pg_user + ":" + pg_pwd + "@" + pg_host + "/" + pg_db + "?sslmode=disable"
	orm.RegisterDataBase("default", "postgres", db_url)

	// register model
	// orm.RegisterModel(new(models.User), new(models.Post), new(models.Comment))

	// create table
	// orm.RunSyncdb("default", false, true)
}
func main() {
	beego.Run()
}
