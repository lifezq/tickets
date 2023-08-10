/**
* PACKAGE driver
* Name mysql
* Description TODO
* Author lifezqy@126.com
* Date 2023/8/10
 */
package driver

import (
	"fmt"

	"github.com/lifezq/tickets/internal/constant"
	"github.com/lifezq/tickets/internal/gormgen"
	"github.com/lifezq/tickets/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		constant.MysqlUser, constant.MysqlPassword, constant.MysqlHost,
		constant.MysqlPort, constant.MysqlDatabase)), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}

	// Migrate the schema
	err = db.AutoMigrate(&model.Ticket{}, &model.MyDatum{})
	if err != nil {
		panic("failed to auto migrate, err:" + err.Error())
	}

	gormgen.SetDefault(db)
}
