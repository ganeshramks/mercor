package model

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type PaymentLineItem struct {
    Id                 int     `orm:"column(id);pk"`
    PaymentLineItemId  int     `orm:"column(paymentLineItemId)"`
    JobUid             int     `orm:"column(jobUid)"`
    TimeLogUid         int     `orm:"column(timeLogUid)"`
    Amount             float64 `orm:"column(amount)"`
    Status             string 	`orm:"column(status)"`
    Version            int 		`orm:"column(version)"`
}

func (t *PaymentLineItem) TableName() string {
	return "paymentLineItems"
}

func init() {
	orm.RegisterModel(new(PaymentLineItem))
}

// AddPaymentLineItem inserts a new PaymentLineItem into PaymentLineItem table and returns last inserted Id on success.
func AddPaymentLineItem(m *PaymentLineItem) (id int64, err error) {
	o := orm.NewOrm()
	return AddPaymentLineItemOrm(m, o)
}

func AddPaymentLineItemOrm(m *PaymentLineItem, o orm.Ormer) (id int64, err error) {
	id, err = o.Insert(m)
	if err != nil {
		fmt.Println("Err in AddPaymentLineItemOrm: ", err)
	}
	return
}

// GetTimeLogByTimeLogId retrieves timelog by timeLogId and jobUid with latest version. Returns error if not found
func GetPaymentLineItemByPaymentLineIdAndJobUidAndTimeLogId(paymentLineItemId, timeLogId, jobUid int) (v *PaymentLineItem, err error) {
	o := orm.NewOrm()
	return GetPaymentLineItemByPaymentLineIdAndJobUidAndTimeLogIdWithORM(paymentLineItemId, timeLogId, jobUid, o)
}

func GetPaymentLineItemByPaymentLineIdAndJobUidAndTimeLogIdWithORM(paymentLineItemId, timeLogId, jobUid int, o orm.Ormer) (v *PaymentLineItem, err error) {
	v = &PaymentLineItem{}
	if err = o.QueryTable(new(PaymentLineItem)).Filter("paymentLineItemId", paymentLineItemId).Filter("timeLogUid", timeLogId).Filter("jobUid", jobUid).OrderBy("-version").One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetLatestPaymentLineItemsRaw(contractorId int, startTime, endTime time.Time) ([]PaymentLineItem, error) {
    o := orm.NewOrm()
    var items []PaymentLineItem

    rawSQL := `
        SELECT pli.*
        FROM paymentLineItems pli
        JOIN (
            SELECT paymentLineItemId, jobUid, timeLogUid, MAX(version) AS latest_version
            FROM paymentLineItems
            GROUP BY paymentLineItemId, jobUid, timeLogUid
        ) latest
            ON pli.paymentLineItemId = latest.paymentLineItemId
            AND pli.jobUid = latest.jobUid
            AND pli.timeLogUid = latest.timeLogUid
            AND pli.version = latest.latest_version
        JOIN job j
            ON pli.jobUid = j.id
        JOIN timeLog tl
            ON pli.timeLogUid = tl.id
        WHERE j.contractorId = ?
          AND tl.timeStart >= ?
          AND tl.timeEnd <= ?;
    `

    _, err := o.Raw(rawSQL, contractorId, startTime, endTime).QueryRows(&items)
    return items, err
}

func CheckAndUpdatePaymentLineItem(newItem PaymentLineItem, checkColumns []string) error {
	return CheckAndUpdateRecordGeneric[PaymentLineItem](
		newItem,
		checkColumns,
		func() (*PaymentLineItem, error) {
			return GetPaymentLineItemByPaymentLineIdAndJobUidAndTimeLogId(
				newItem.PaymentLineItemId,
				newItem.TimeLogUid,
				newItem.JobUid,
			)
		},
		func(newData PaymentLineItem, oldData *PaymentLineItem, recordExists bool, columns []string) error {
			return UpdateWithTrace(newData, oldData, recordExists, columns)
		},
	)
}
