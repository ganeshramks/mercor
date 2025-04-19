package main

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"github.com/spf13/cast"


	"common-libs/utility"
	"common-libs/model"
)

func init() {
	fmt.Println("init")

	dbName:="mercor"
	err := utility.ConnectToDatabase(dbName)
	if err != nil {
		fmt.Println("ConnectToDatabase failed with err:", err)
		return
	}

}


func main() {

	newJob := model.Job {
		Version:1,
		JobId:9,
		Rate:0.29,
		Status:"active",
		Title:"Director Investment Associate",
		CompanyId:4,
		ContractorId:101909,
	}

	err := model.CheckAndUpdateRecord(newJob, []string{"status", "rate"})
	if err != nil {
		fmt.Println("err in CheckAndUpdateRecord: ", err) 
	}

	job, err := model.GetJobByJobId(newJob.JobId)
	if err != nil && err != orm.ErrNoRows {
		fmt.Println("err in GetJobById: ", err)
	}
	fmt.Println("New job: ", job, "________________________________")

	// Get all Jobs for a company with latest versions as active/inactive etc
	fmt.Println(model.GetLatestJobsByFieldAndStatus("companyId",newJob.CompanyId, "active"))

	// Get all Jobs for a contractor with latest versions as active/inactive etc
	fmt.Println(model.GetLatestJobsByFieldAndStatus("contractorId", newJob.ContractorId, "active"))


	start := time.Date(2023, 4, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2025, 4, 20, 23, 59, 59, 0, time.UTC)

	fmt.Println(model.GetLatestTimeLogs(newJob.ContractorId, start, end))


	
	fmt.Println("Time Log starts--------===========================")

	newTimeLog := model.TimeLog {
		Version: 1,
		TimeLogId: 23,
		Duration: cast.ToInt(end.Sub(start).Minutes()), //duration is in minutes
		TimeStart: time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
		TimeEnd:time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC),	
		Type: "adjusted",
		Jobuid: job.Id,
	}


	err = model.CheckAndUpdateTimeLogRecord(newTimeLog, []string{"type", "duration"})
	if err != nil {
		fmt.Println("err in CheckAndUpdateTimeLogRecord: ", err) 
	}

}