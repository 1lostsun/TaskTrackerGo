package scheduler

import (
	"context"
	"log"
	"taskTrackerGo/internal/service"
	"time"
)

func StartEscalationScheduler(ctx context.Context, ts service.TaskService) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Print("Scheduler stopped")
				return
			default:
				err := ts.EscalateOverdueTasks(ctx)
				if err != nil {
					log.Fatalf("Escalating Overdue Tasks: %v", err)
				}

				log.Print("Escalating Overdue Tasks was finished")
				time.Sleep(10 * time.Minute)
			}
		}
	}()
}
