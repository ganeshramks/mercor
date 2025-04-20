package model

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"time"
	// "reflect"
)

type TimeLog struct {
	Id                     int        `orm:"column(id);auto" description:"Id"`
	Version                int        `orm:"column(version);"`
	TimeLogId              int        `orm:"column(timeLogId);"`
	Duration        	   int    	  `orm:"column(duration);"`
	TimeStart 			   time.Time  `orm:"column(timeStart)"`
    TimeEnd   			   time.Time  `orm:"column(timeEnd)"`
	// Status				   string     `orm:"column(status);"`
	Type                   string     `orm:"column(type);"`
	Jobuid                 int        `orm:"column(jobUid);"`
	// ContractorId		   int        `orm:"column(contractorId);"`
}

func (t *TimeLog) TableName() string {
	return "timeLog"
}

func init() {
	orm.RegisterModel(new(TimeLog))
}

// AddTimeLog inserts a new timelog into timelog table and returns last inserted Id on success.
func AddTimeLog(m *TimeLog) (id int64, err error) {
	o := orm.NewOrm()
	return AddTimeLogOrm(m, o)
}

func AddTimeLogOrm(m *TimeLog, o orm.Ormer) (id int64, err error) {
	id, err = o.Insert(m)
	if err != nil {
		fmt.Println("Err in AddTimeLogOrm: ", err)
	}
	return
}

// GetTimeLogByTimeLogId retrieves timelog by timeLogId and jobUid with latest version. Returns error if not found
func GetTimeLogByTimeLogIdAndJobUid(timeLogId, jobUid int) (v *TimeLog, err error) {
	o := orm.NewOrm()
	return GetTimeLogByTimeLogIdAndJobUidWithORM(timeLogId, jobUid, o)
}

func GetTimeLogByTimeLogIdAndJobUidWithORM(timeLogId, jobUid int, o orm.Ormer) (v *TimeLog, err error) {
	v = &TimeLog{}
	if err = o.QueryTable(new(TimeLog)).Filter("timeLogId", timeLogId).Filter("jobUid", jobUid).OrderBy("-version").One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetLatestTimeLogs(contractorId int, startTime, endTime time.Time) ([]TimeLog, error) {
    o := orm.NewOrm()
    var logs []TimeLog

    rawSQL := `
        SELECT tl.*
        FROM timeLog tl
        JOIN (
            SELECT timeLogId, jobUid, MAX(version) AS latest_version
            FROM timeLog
            GROUP BY timeLogId, jobUid
        ) latest 
            ON tl.timeLogId = latest.timeLogId 
            AND tl.jobUid = latest.jobUid 
            AND tl.version = latest.latest_version
        JOIN job j 
            ON tl.jobUid = j.id
        WHERE j.contractorId = ?
          AND tl.timeStart >= ?
          AND tl.timeEnd <= ?;
    `

    _, err := o.Raw(rawSQL, contractorId, startTime, endTime).QueryRows(&logs)
    return logs, err
}


func CheckAndUpdateTimeLogRecord(newTimeLog TimeLog, checkColumns []string) error {
	return checkAndUpdateRecordGeneric(
		newTimeLog,
		checkColumns,
		func() (interface{}, error) {
			return GetTimeLogByTimeLogIdAndJobUid(newTimeLog.TimeLogId, newTimeLog.Jobuid)
		},
	)
}