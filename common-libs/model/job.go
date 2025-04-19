package model

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	// "reflect"
)

type Job struct {
	Id                     int        `orm:"column(id);auto" description:"Id"`
	Version                int        `orm:"column(version);"`
	JobId              	   int        `orm:"column(jobId);"`
	Rate        		   float64    `orm:"column(rate);"`
	Status				   string     `orm:"column(status);"`
	Title                  string     `orm:"column(title);"`
	CompanyId              int        `orm:"column(companyId);"`
	ContractorId		   int        `orm:"column(contractorId);"`
}

func (t *Job) TableName() string {
	return "job"
}

func init() {
	orm.RegisterModel(new(Job))
}

// AddJob inserts a new job into job table and returns last inserted Id on success.
func AddJob(m *Job) (id int64, err error) {
	o := orm.NewOrm()
	return AddJobOrm(m, o)
}

func AddJobOrm(m *Job, o orm.Ormer) (id int64, err error) {
	id, err = o.Insert(m)
	if err != nil {
		fmt.Println("Err in AddJobOrm: ", err)
	}
	return
}

// GetJobById retrieves Job by jobId with latest version. Returns error if not found
func GetJobByJobId(jobId int) (v *Job, err error) {
	o := orm.NewOrm()
	return GetJobByJobIdWithOrm(jobId, o)
}

func GetJobByJobIdWithOrm(jobId int, o orm.Ormer) (v *Job, err error) {
	v = &Job{}
	if err = o.QueryTable(new(Job)).Filter("jobId", jobId).OrderBy("-version").One(v); err == nil {
		return v, nil
	}
	return nil, err
}


func GetLatestJobsByFieldAndStatus(field string, fieldValue int, status string) ([]Job, error) {
    o := orm.NewOrm()
    var jobs []Job

    rawSQL := fmt.Sprintf(`
        SELECT j.*
        FROM job j
        JOIN (
            SELECT jobId, MAX(version) AS latest_version
            FROM job
            WHERE %s = ?
            GROUP BY jobId
        ) latest_jobs ON j.jobId = latest_jobs.jobId AND j.version = latest_jobs.latest_version
        WHERE j.%s = ? AND j.status = ?;
    `, field, field)

    _, err := o.Raw(rawSQL, fieldValue, fieldValue, status).QueryRows(&jobs)
    return jobs, err
}

func CheckAndUpdateRecord(newJob Job, checkColumns []string) (err error) {
	// check and get if any record exists with the jobId received in newJob object input
	recordExists := true
	currentRecord, err := GetJobByJobId(newJob.JobId)
	if err != nil && err != orm.ErrNoRows {
		fmt.Println("Err in GetJobByJobId: ", err)
		return 
	}
	if err == orm.ErrNoRows {
		recordExists = false 
	}

	err = UpdateWithTrace(newJob, currentRecord, recordExists, checkColumns)
	if err != nil {
		fmt.Println("UpdateWithTrace err: ", err) 
	}

	return
}