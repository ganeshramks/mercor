package model

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"reflect"
	"github.com/spf13/cast"
	"errors"
	"strings"


	"github.com/fatih/structs"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/text/cases"
    "golang.org/x/text/language"
)


func UpdateWithTrace(newObj, currentObj interface{}, recordExists bool, checkColumns []string) (err error) {
	o := orm.NewOrm()
	err = UpdateWithTraceORM(newObj, currentObj, o, recordExists, checkColumns)
	return err
}

func UpdateWithTraceORM(newObj, currentObj interface{}, o orm.Ormer, recordExists bool, checkColumns []string) error {
	fmt.Println("In UpdateWithTraceORM")

	if newObj == nil {
		return errors.New("newObj cannot be nil")
	}

	// Convert new object to map
	newObjDataMap := structs.Map(newObj)
	fmt.Println("newObjDataMap after struct.Map =", newObjDataMap)

	var currentObjDataMap map[string]interface{}
	updateRequired := false

	if recordExists {
		fmt.Println("recordExists")
		currentObjDataMap = structs.Map(currentObj)
		fmt.Println("currentObjDataMap =", currentObjDataMap)

		for _, column := range checkColumns {
			field := cases.Title(language.English, cases.Compact).String(column)
			newVal, newOk := newObjDataMap[field]
			oldVal, oldOk := currentObjDataMap[field]

			fmt.Printf("Comparing %s: new=%v, old=%v\n", field, newVal, oldVal)

			if newOk && oldOk && newVal != oldVal {
				updateRequired = true
				break
			}
		}
	}

	typeName := cast.ToString(reflect.TypeOf(newObj))

	// Normalize type name
	if strings.HasPrefix(typeName, "*") {
		typeName = typeName[1:]
	}

	switch typeName {
	case "model.Job", "Job":
		return handleVersionedInsert[Job](newObjDataMap, currentObjDataMap, recordExists, updateRequired, o, AddJobOrm)
	case "model.TimeLog", "TimeLog":
		return handleVersionedInsert[TimeLog](newObjDataMap, currentObjDataMap, recordExists, updateRequired, o, AddTimeLogOrm)
	case "model.PaymentLineItem", "PaymentLineItem":
		return handleVersionedInsert[PaymentLineItem](newObjDataMap, currentObjDataMap, recordExists, updateRequired, o, AddPaymentLineItemOrm)
	default:
		return errors.New("unsupported model type")
	}
}

func handleVersionedInsert[T any](
	newMap map[string]interface{},
	currentMap map[string]interface{},
	recordExists bool,
	updateRequired bool,
	o orm.Ormer,
	insertFunc func(*T, orm.Ormer) (int64, error),
) error {
	var newData T
	if err := mapstructure.Decode(newMap, &newData); err != nil {
		return fmt.Errorf("error decoding new struct: %w", err)
	}

	if !recordExists {
		setVersion(&newData, 1)
	} else if updateRequired {
		var currentData T
		if err := mapstructure.Decode(currentMap, &currentData); err == nil {
			setVersion(&newData, getVersion(currentData)+1)
		}
	}

	if !recordExists || updateRequired {
		fmt.Printf("Inserting new version: %+v\n", newData)
		_, err := insertFunc(&newData, o)
		return err
	}

	return nil
}

func setVersion(obj interface{}, version int) {
	v := reflect.ValueOf(obj).Elem()
	if field := v.FieldByName("Version"); field.IsValid() && field.CanSet() {
		field.SetInt(int64(version))
	}
}

func getVersion(obj interface{}) int {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if field := v.FieldByName("Version"); field.IsValid() {
		return int(field.Int())
	}
	return 0
}

