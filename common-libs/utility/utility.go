package utility

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

// CPrint is a sample utility print function.
func CPrint() {
	fmt.Println("Hello from utility")
}

// ConnectToDatabase registers and tests a connection to the specified database.
func ConnectToDatabase(dbName string) error {
	fmt.Println("Connecting to database:", dbName)

	creds, err := fetchDBCreds(dbName)
	if err != nil {
		return fmt.Errorf("failed to fetch DB credentials: %w", err)
	}

	dsn := fmt.Sprintf("%s:%s@/%s?charset=utf8", creds["userName"], creds["pwd"], dbName)

	if _, err := orm.GetDB(dbName); err != nil {
		if regErr := orm.RegisterDataBase("default", "mysql", dsn); regErr != nil {
			return fmt.Errorf("failed to register DB: %w", regErr)
		}
	}

	if err := testDBConn(); err != nil {
		return fmt.Errorf("DB connection test failed: %w", err)
	}

	fmt.Println("Successfully connected to database:", dbName)
	return nil
}

// testDBConn pings the database with a simple query.
func testDBConn() error {
	o := orm.NewOrm()
	_, err := o.Raw("SELECT 1").Exec()
	return err
}

// fetchDBCreds retrieves database credentials.
func fetchDBCreds(dbName string) (map[string]string, error) {
	// NOTE: In production, I will use environment variables or secret manager (eg: AWS Secrets Manager).
	dbSecrets := map[string]map[string]string{
		"mercor": {
			"userName": "mercor",
			"pwd":      "mercor123",
		},
	}

	if creds, ok := dbSecrets[dbName]; ok {
		return creds, nil
	}

	return nil, errors.New("database credentials not found")
}
