package bootstrap

import (
	"github.com/goravel/framework/contracts/queue"

	"order-service/app/jobs"
)

func Jobs() []queue.Job {
	return []queue.Job{
		&jobs.ProcessOrder{},
	}
}
