package main

import (
	"antikode-test/config"
	"antikode-test/util"
	"fmt"
)

func main() {
	config := config.GetConfig()
	db := util.MysqlDatabaseConnection(config)
	fmt.Println(db)
}
