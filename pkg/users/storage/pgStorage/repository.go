package pgStorage

import (
	"fmt"
	"github.ru/GeorgVartanov/myApp/pkg/database/pgDB"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type UserPostgresStorage struct {
	*pgDB.PostgresDB
}

func NewUserPostgresStorage(pg *pgDB.PostgresDB) *UserPostgresStorage {
	return &UserPostgresStorage{PostgresDB: pg}
}

// CreateUserTableORDrop if True Table will be create, false will drop table
func (u *UserPostgresStorage) CreateUserTableORDrop(status bool) error {
	baseDir, err := os.Getwd()
	fmt.Println(baseDir)
	if err != nil {
		log.Fatal(err)
	}
	sqlPath := filepath.Join(baseDir, "pkg", "users", "storage", "pgStorage", "sql")
	if status {
		createUserPath := filepath.Join(sqlPath, "createUserTable.sql")
		c, ioErr := ioutil.ReadFile(createUserPath)
		if ioErr != nil {
			return nil
		}
		sql := string(c)
		fmt.Println(sql)
		_, err := u.Exec(sql)
		if err != nil {
			return err
		}
	} else {
		dropUserPath := filepath.Join(sqlPath, "dropUserTable.sql")
		fmt.Println(dropUserPath)
		c, ioErr := ioutil.ReadFile(dropUserPath)
		if ioErr != nil {
			return nil
		}
		sql := string(c)
		fmt.Println(sql)
		_, err := u.Exec(sql)
		if err != nil {
			return err
		}

	}

	return nil
}
