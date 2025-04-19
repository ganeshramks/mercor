package utility

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func CPrint(){
	fmt.Println("Hello from utility")
}

func ConnectToDatabase(dbName string) (err error) {
	fmt.Println("In connectToDatabase")

	// orm.RegisterModel(new(User))
	dbCreds := fetchDBCreds(dbName)
	dsn:= dbCreds["userName"]+":"+dbCreds["pwd"]+"@/"+dbName+"?charset=utf8"

	_, err = orm.GetDB(dbName)
	if err != nil {
		err = orm.RegisterDataBase("default", "mysql", dsn)
	}
	fmt.Println("error:", err)
	err = testDBConn()
	if err != nil {
		fmt.Println("connection to "+dbName+" db failed with error:",err)
		return
	}
	fmt.Println("connection to "+dbName+" db success")
	return
}

func testDBConn() error {
	o := orm.NewOrm()
	// o.Using(dbName)
	_, err := o.Raw("SELECT 1").Exec()
	return err
}

func fetchDBCreds(dbName string) (map[string]string){
	fmt.Println("Inside fetchDBCreds")
	
	// Have a map storing userName and passwords in a place like AWS secrets or lambda env variables

	dbSecrets := make(map[string]map[string]string)

	mercorDBCredsMap := map[string]string{
		"userName": "mercor",
		"pwd":      "mercor123",
	}

	dbSecrets["mercor"] = mercorDBCredsMap

	val, ok := dbSecrets[dbName]
	if ok {
	    return val
	}

	return nil
}