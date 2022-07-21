package handlers

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConfigDBConnection(connString string, schemaName string) (db *gorm.DB, err error) {

	db, err = gorm.Open(postgres.Open(connString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   fmt.Sprintf("%s.", schemaName),
			SingularTable: true,
		},
	})

	return db, err
}
