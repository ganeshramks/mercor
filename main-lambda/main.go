package main

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cast"

	"common-libs/model"
	"common-libs/utility"
)

func init() {
	fmt.Println("Initializing database connection...")
	if err := utility.ConnectToDatabase("mercor"); err != nil {
		fmt.Printf("Database connection failed: %v\n", err)
	}
}

func main() {
	job := createOrUpdateJob()
	if job == nil {
		return
	}

	timeLog := createOrUpdateTimeLog(job.Id, job.ContractorId)
	if timeLog == nil {
		return
	}

	createOrUpdatePaymentLineItem(job.Id, timeLog.Id, job.ContractorId)
}

func createOrUpdateJob() *model.Job {
	job := model.Job{
		Version:      1,
		JobId:        9,
		Rate:         0.29,
		Status:       "active",
		Title:        "Director Investment Associate",
		CompanyId:    4,
		ContractorId: 101909,
	}

	if err := model.CheckAndUpdateJob(job, []string{"status", "rate"}); err != nil {
		fmt.Printf("Failed to update job: %v\n", err)
		return nil
	}

	existingJob, err := model.GetJobByJobId(job.JobId)
	if err != nil && err != orm.ErrNoRows {
		fmt.Printf("Error retrieving job: %v\n", err)
		return nil
	}

	fmt.Println("Retrieved job:", existingJob)
	fmt.Println(model.GetLatestJobsByFieldAndStatus("companyId", existingJob.CompanyId, "active"))
	fmt.Println(model.GetLatestJobsByFieldAndStatus("contractorId", existingJob.ContractorId, "active"))
	return existingJob
}

func createOrUpdateTimeLog(jobId, contractorId int) *model.TimeLog {
	start := time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 5, 20, 23, 59, 59, 0, time.UTC)

	timeLog := model.TimeLog{
		Version:    1,
		TimeLogId:  23,
		Duration:   cast.ToInt(end.Sub(start).Minutes()),
		TimeStart:  start,
		TimeEnd:    end,
		Type:       "adjusted",
		Jobuid:     jobId,
	}

	if err := model.CheckAndUpdateTimeLog(timeLog, []string{"type", "duration"}); err != nil {
		fmt.Printf("Failed to update time log: %v\n", err)
		return nil
	}

	existingLog, err := model.GetTimeLogByTimeLogIdAndJobUid(timeLog.TimeLogId, jobId)
	if err != nil && err != orm.ErrNoRows {
		fmt.Printf("Error retrieving time log: %v\n", err)
		return nil
	}

	fmt.Println("Retrieved time log:", existingLog)
	fmt.Println(model.GetLatestTimeLogs(contractorId, start, end))
	return existingLog
}

func createOrUpdatePaymentLineItem(jobId, timeLogId, contractorId int) {
	item := model.PaymentLineItem{
		PaymentLineItemId: 101,
		JobUid:            jobId,
		TimeLogUid:        timeLogId,
		Amount:            708,
		Status:            "paid",
		Version:           2,
	}

	if err := model.CheckAndUpdatePaymentLineItem(item, []string{"amount", "status"}); err != nil {
		fmt.Printf("Failed to update payment line item: %v\n", err)
		return
	}

	start := time.Date(2020, 4, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 5, 20, 23, 59, 59, 0, time.UTC)
	fmt.Println(model.GetLatestPaymentLineItemsRaw(contractorId, start, end))
}
