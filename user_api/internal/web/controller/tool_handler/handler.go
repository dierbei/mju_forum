package tool_handler

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	HashIdsView() core.HandlerFunc
	LogsView() core.HandlerFunc
	CacheView() core.HandlerFunc
	DataView() core.HandlerFunc
	WebsocketView() core.HandlerFunc
}

type handler struct {
	logger *zap.Logger
	cache  cache.Repo
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo) Handler {
	return &handler{
		logger: logger,
		cache:  cache,
	}
}

func (h *handler) i() {}
