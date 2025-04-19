package model

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"reflect"
	"github.com/spf13/cast"
	"errors"


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

func UpdateWithTraceORM(newObj, currentObj interface{}, o orm.Ormer, recordExists bool, checkColumns []string) (err error) {
	fmt.Println("In UpdateWithTraceORM")

	if newObj == nil {
		err = errors.New("New Obj cannot be nil")
		return
	}

	currentObjDataMap := make(map[string]interface{})
	//extract interface value by converting it to map 
	newObjDataMap := structs.Map(newObj)
	fmt.Println("newObjDataMap after struct.Map = ", newObjDataMap)
	
	updateRequired := false

	// Note
	// currentObj will be nil if recordExists is false so have the below check to avoid panic and check if newObj differs from currentObj (if it exists)
	// for the range of columns specified in checkColumns
	if recordExists {
		fmt.Println("recordExists")
		currentObjDataMap = structs.Map(currentObj)
		fmt.Println("currentObjDataMap=", currentObjDataMap)
		for _, columnValue := range checkColumns {
			capitalizedColumnValue := cases.Title(language.English, cases.Compact).String(columnValue)
			newValue, ok := newObjDataMap[capitalizedColumnValue];
			currentValue, ok2 := currentObjDataMap[capitalizedColumnValue];
			fmt.Println("newValue=", newValue, "currentValue=", currentValue, newValue==currentValue)
			if ok && ok2 && newValue != currentValue  {
				fmt.Println("update is Required")
				updateRequired = true
				break
			}
		}
	}

	switch cast.ToString(reflect.TypeOf(newObj)) {
		case "model.Job", "*model.Job", "Job":
			var newObjData Job


			//convert map into custom struct (Job)
			mapstructure.Decode(newObjDataMap, &newObjData)

			// if no existing record found then insert the new object with Version as 1
			if !recordExists {
				newObjData.Version = 1
			}
			fmt.Println("newObjData.Version = ", newObjData.Version)
			// increment version by 1 to track the change in the existing record
			fmt.Println("vale of updateRequired = ", updateRequired)
			if updateRequired {
				var currentObjData Job
				mapstructure.Decode(currentObjDataMap, &currentObjData)
				newObjData.Version = currentObjData.Version + 1
				fmt.Println("newObjData.Version = ", newObjData.Version)
			}

			if !recordExists || updateRequired {
				fmt.Println("newObjData before insertion = ", newObjData)
				_, err = AddJobOrm(&newObjData, o)
				if err != nil {
					fmt.Println("UpdateWithTraceORM.AddJobOrm: Err in insert: ", err)
				}
				break;
			}
		case "model.TimeLog", "*model.TimeLog", "TimeLog":
			var newObjData TimeLog


			//convert map into custom struct (Job)
			mapstructure.Decode(newObjDataMap, &newObjData)

			// if no existing record found then insert the new object with Version as 1
			if !recordExists {
				newObjData.Version = 1
			}
			fmt.Println("newObjData.Version = ", newObjData.Version)
			// increment version by 1 to track the change in the existing record
			fmt.Println("vale of updateRequired = ", updateRequired)
			if updateRequired {
				var currentObjData Job
				mapstructure.Decode(currentObjDataMap, &currentObjData)
				newObjData.Version = currentObjData.Version + 1
				fmt.Println("newObjData.Version = ", newObjData.Version)
			}

			if !recordExists || updateRequired {
				fmt.Println("newObjData before insertion = ", newObjData)
				_, err = AddTimeLogOrm(&newObjData, o)
				if err != nil {
					fmt.Println("UpdateWithTraceORM.AddTimeLogOrm: Err in insert: ", err)
				}
				break;
			}
		case "model.PaymentLineItem", "*model.PaymentLineItem", "PaymentLineItem":
			var newObjData PaymentLineItem


			//convert map into custom struct (Job)
			mapstructure.Decode(newObjDataMap, &newObjData)

			// if no existing record found then insert the new object with Version as 1
			if !recordExists {
				newObjData.Version = 1
			}
			fmt.Println("newObjData.Version = ", newObjData.Version)
			// increment version by 1 to track the change in the existing record
			fmt.Println("vale of updateRequired = ", updateRequired)
			if updateRequired {
				var currentObjData Job
				mapstructure.Decode(currentObjDataMap, &currentObjData)
				newObjData.Version = currentObjData.Version + 1
				fmt.Println("newObjData.Version = ", newObjData.Version)
			}

			if !recordExists || updateRequired {
				fmt.Println("newObjData before insertion = ", newObjData)
				_, err = AddPaymentLineItemOrm(&newObjData, o)
				if err != nil {
					fmt.Println("UpdateWithTraceORM.AddPaymentLineItemOrm: Err in insert: ", err)
				}
				break;
			}
		default:
				err = errors.New("Unsupported model type")
	}

	return

}
