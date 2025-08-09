package scheduler

import (
	"context"
	"log"
	"taskTrackerGo/internal/service"
	"time"
)

func StartEscalationScheduler(ctx context.Context, ts service.TaskService) {
	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		defer ticker.Stop()

		if err := ts.EscalateOverdueTasks(ctx); err != nil {
			log.Fatalf("Escalating Overdue Tasks: %v", err)
		}
		log.Print("Escalating Overdue Tasks was finished")

		for {
			select {
			case <-ctx.Done():
				log.Print("Scheduler stopped")
				return
			case <-ticker.C:
				if err := ts.EscalateOverdueTasks(ctx); err != nil {
					log.Fatalf("Escalating Overdue Tasks: %v", err)
				}
				log.Print("Escalating Overdue Tasks was finished")
			}
		}
	}()
}
