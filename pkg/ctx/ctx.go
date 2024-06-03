package ctx

import (
	"alarm_collector/internal/cache"
	"alarm_collector/internal/ck"
	"alarm_collector/internal/repo"
	"context"
)

type Context struct {
	DB    repo.InterEntryRepo
	Redis cache.InterEntryCache
	CK    ck.InterEntryRepo
	Ctx   context.Context
}

var (
	DB    repo.InterEntryRepo
	Redis cache.InterEntryCache
	CK    ck.InterEntryRepo
	Ctx   context.Context
)

func NewContext(ctx context.Context, db repo.InterEntryRepo, redis cache.InterEntryCache, clickHouse ck.InterEntryRepo) *Context {
	DB = db
	Redis = redis
	CK = clickHouse
	Ctx = ctx
	return &Context{
		DB:    db,
		Redis: redis,
		CK:    clickHouse,
		Ctx:   ctx,
	}
}

func DO() *Context {
	return &Context{
		DB:    DB,
		Redis: Redis,
		CK:    CK,
		Ctx:   Ctx,
	}
}
