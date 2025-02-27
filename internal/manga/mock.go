package manga

import (
	"path/filepath"
	"seanime/internal/database/db"
	"seanime/internal/events"
	"seanime/internal/test_utils"
	"seanime/internal/util"
	"seanime/internal/util/filecache"
	"testing"
)

func GetMockRepository(t *testing.T, db *db.Database) *Repository {
	logger := util.NewLogger()
	fileCacher, err := filecache.NewCacher(filepath.Join(test_utils.ConfigData.Path.DataDir, "cache"))
	if err != nil {
		t.Fatal(err)
	}

	repository := NewRepository(&NewRepositoryOptions{
		Logger:         logger,
		FileCacher:     fileCacher,
		ServerURI:      "",
		WsEventManager: events.NewMockWSEventManager(logger),
		DownloadDir:    filepath.Join(test_utils.ConfigData.Path.DataDir, "manga"),
		Database:       db,
	})

	return repository
}
