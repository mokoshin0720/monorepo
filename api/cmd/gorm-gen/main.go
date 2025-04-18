package main

import (
	"fmt"

	"github.com/mokoshin0720/monorepo/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./app/infra/entity",
		ModelPkgPath:  "entity",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	dns := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=true",
		config.MySQL.User,
		config.MySQL.Password,
		config.MySQL.Host,
		config.MySQL.Port,
		config.MySQL.Database,
	)

	mysqlDB, err := gorm.Open(mysql.Open(dns))
	if err != nil {
		panic(err)
	}

	g.UseDB(mysqlDB)
	g.GenerateAllTable()

	genOrganization(g)

	g.Execute()
}
