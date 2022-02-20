package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

// Users -
type Users struct {
	ID        int         `orm:"column(id);auto"`
	Name      string      `orm:"column(name)"`
	Posts     []*Posts    `orm:"reverse(many)"`
	Comments  []*Comments `orm:"reverse(many)"`
	CreatedAt time.Time   `orm:"column(created_at);auto_now_add;type(datetime)"`
	UpdatedAt time.Time   `orm:"column(updated_at);auto_now;type(datetime)"`
}

type Posts struct {
	ID        int       `orm:"column(id);auto"`
	Title     string    `orm:"column(title)"`
	Content   string    `orm:"column(content);type(text)"`
	User      *Users    `orm:"rel(fk);column(user_id)"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)"`
}

type Comments struct {
	ID        int       `orm:"column(id);auto"`
	Content   string    `orm:"column(content);type(text)"`
	Post      *Posts    `orm:"rel(fk);column(post_id)"`
	User      *Users    `orm:"rel(fk);column(user_id)"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)"`
}

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
	orm.RegisterModel(new(Users), new(Posts), new(Comments))

	// create table
	// orm.RunSyncdb("default", false, true)
}

func main() {
	// o := orm.NewOrm()

	// user := Users{Name: "Michael"}

	// insert
	// id, err := o.Insert(&user)
	// fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	// user.Name = "Jerry"
	// num, err := o.Update(&user)
	// fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// // read one
	// u := Users{ID: user.ID}
	// err = o.Read(&u)
	// fmt.Printf("ERR: %v\n", err)

	// // delete
	// num, err = o.Delete(&u)
	// fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
