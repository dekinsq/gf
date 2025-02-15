package main

import (
	"fmt"
	"github.com/dekinsq/gf/database/gdb"
	"github.com/dekinsq/gf/frame/g"
)

func main() {
	//db := g.DB()

	gdb.AddDefaultConfigNode(gdb.ConfigNode{
		LinkInfo: "root:12345678@tcp(127.0.0.1:3306)/test?parseTime=true&loc=Local",
		Type:     "mysql",
		Charset:  "utf8",
	})
	db, _ := gdb.New()

	db.SetDebug(true)

	r, e := db.Table("user").Data(g.Map{
		"create_at": "now()",
	}).Unscoped().Insert()
	if e != nil {
		panic(e)
	}
	if r != nil {
		fmt.Println(r.LastInsertId())
	}
}
