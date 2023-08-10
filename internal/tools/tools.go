/**
* PACKAGE tickets
* Name tools
* Description TODO
* Author lifezqy@126.com
* Date 2023/8/10
 */
package main

import (
	"fmt"

	"github.com/lifezq/tickets/internal/constant"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type TicketsQuerier interface {
	// where("`id`=@id")
	FindByID(id int32) (*gen.T, error)
}

type MyDataQuerier interface {
	// where("`id`=@id")
	FindByID(id int32) (*gen.T, error)
}

func main() {

	g := gen.NewGenerator(gen.Config{
		OutPath:           "./internal/gormgen",
		ModelPkgPath:      "./model",
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		FieldCoverable:    false,
		FieldSignable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})

	gormdb, _ := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		constant.MysqlUser, constant.MysqlPassword, constant.MysqlHost,
		constant.MysqlPort, constant.MysqlDatabase)), &gorm.Config{})
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	g.ApplyInterface(func(TicketsQuerier) {}, g.GenerateModel("tickets"))
	g.ApplyInterface(func(MyDataQuerier) {}, g.GenerateModel("my_data"))

	// Generate the code
	g.Execute()
}
