package tasker

import (
	"log"
	"time"

	"github.com/gunawanpras/go-tasker/config"
)

type TaskScheduler struct {
	tasks []Task
}

type Task struct {
	Name     string
	Schedule time.Duration
	Job      func()
}

func NewScheduler(config config.Config) *TaskScheduler {
	return &TaskScheduler{
		tasks: []Task{},
	}
}

func (ts *TaskScheduler) AddTask(name string, schedule time.Duration, job func()) {
	ts.tasks = append(ts.tasks, Task{Name: name, Schedule: schedule, Job: job})
}

func (ts *TaskScheduler) Run() {
	for _, task := range ts.tasks {
		go func(t Task) {
			for {
				log.Println("Running task: [" + t.Name + "]")
				t.Job()
				println()
				time.Sleep(t.Schedule)
			}
		}(task)
	}
}
