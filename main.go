/**
* PACKAGE main
* Name main
* Description TODO
* Author lifezqy@126.com
* Date 2023/8/10
 */
package main

import (
	"github.com/lifezq/tickets/internal/driver"
	"github.com/lifezq/tickets/internal/service"
)

func main() {
	driver.Connect()
	//service.DataSync()
	//service.Data2Sync()

	service.Analysis()
}
