package cron

import (
	"seanime/internal/core"
	"time"
)

type JobCtx struct {
	App *core.App
}

func RunJobs(app *core.App) {

	// Run the jobs only if the server is online
	if !app.IsOffline() {

		ctx := &JobCtx{
			App: app,
		}

		refreshAnilistTicker := time.NewTicker(10 * time.Minute)
		refreshLocalDataTicker := time.NewTicker(31 * time.Minute)
		refetchReleaseTicker := time.NewTicker(1 * time.Hour)

		go func() {
			for {
				select {
				case <-refreshAnilistTicker.C:
					RefreshAnilistDataJob(ctx)
				case <-refreshLocalDataTicker.C:
					SyncLocalDataJob(ctx)
				case <-refetchReleaseTicker.C:
					app.Updater.ShouldRefetchReleases()
				}
			}
		}()

	}
}
