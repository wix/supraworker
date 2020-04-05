package worker

import (
	"context"
	"fmt"
	// "os/exec"
	"sync"
	"testing"
	"time"
	// "github.com/weldpua2008/supraworker/job"
	"github.com/sirupsen/logrus"
	model "github.com/weldpua2008/supraworker/model"
	"github.com/weldpua2008/supraworker/model/cmdtest"
)

func TestHelperProcess(t *testing.T) {
	cmdtest.TestHelperProcess(t)
}
func init() {
	logrus.SetLevel(logrus.WarnLevel)
}

func TestExecuteJobSuccess(t *testing.T) {
	var wg sync.WaitGroup
	jobs := make(chan *model.Job, 1)

	wg.Add(1)
	go StartWorker(0, jobs, &wg)
	jobOne := model.NewTestJob(fmt.Sprintf("job-%v", cmdtest.GetFunctionName(t.Name)), cmdtest.CMDForTest("echo  &&exit 0"))
	jobOne.TTR = 10000000

	jobs <- jobOne
	close(jobs)
	wg.Wait()
	if jobOne.Status != model.JOB_STATUS_SUCCESS {
		t.Errorf("Expected %s, got %s\n", model.JOB_STATUS_SUCCESS, jobOne.Status)
	}
}

func TestExecuteJobFail(t *testing.T) {
	var wg sync.WaitGroup
	jobs := make(chan *model.Job, 1)

	wg.Add(1)
	go StartWorker(0, jobs, &wg)

	jobOne := model.NewTestJob(fmt.Sprintf("job-%v", cmdtest.GetFunctionName(t.Name)), cmdtest.CMDForTest("echo  &&exit 1"))
	jobOne.TTR = 10000000

	jobs <- jobOne
	close(jobs)
	wg.Wait()
	if jobOne.Status != model.JOB_STATUS_ERROR {
		t.Errorf("Expected %s, got %s\n", model.JOB_STATUS_ERROR, jobOne.Status)
	}
}

func TestExecuteJobContextCancel(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel() // cancel when we are getting the kill signal or exit
	time.Sleep(1 * time.Millisecond)
	var wg sync.WaitGroup
	jobs := make(chan *model.Job, 1)

	wg.Add(1)
	go StartWorker(0, jobs, &wg)

	jobOne := model.NewTestJob(fmt.Sprintf("job-%v", cmdtest.GetFunctionName(t.Name)), cmdtest.CMDForTest("echo  && sleep 100 &&exit 0"))
	jobOne.SetContext(ctx)
	jobOne.TTR = 10000000

	jobs <- jobOne
	close(jobs)
	wg.Wait()
	if jobOne.Status != model.JOB_STATUS_ERROR {
		t.Errorf("Expected %s, got %s\n", model.JOB_STATUS_ERROR, jobOne.Status)
	}
}

func TestExecuteJobTTRCanceled(t *testing.T) {
	var wg sync.WaitGroup
	jobs := make(chan *model.Job, 1)

	wg.Add(1)
	go StartWorker(0, jobs, &wg)

	jobOne := model.NewTestJob(fmt.Sprintf("job-%v", cmdtest.GetFunctionName(t.Name)), cmdtest.CMDForTest("echo  && sleep 100 &&exit 0"))
	jobOne.TTR = 1

	jobs <- jobOne
	close(jobs)
	wg.Wait()
	// time.Sleep(10 * time.Millisecond)
	if jobOne.Status != model.JOB_STATUS_ERROR {
		t.Errorf("Expected %s, got %s\n", model.JOB_STATUS_ERROR, jobOne.Status)
	}
}