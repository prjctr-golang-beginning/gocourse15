package providers

import (
	"github.com/samber/do"
	"gocourse15/app"
	"gocourse15/pkg/myhttp"
	"gocourse15/pkg/scheduler"
	"gocourse15/pkg/users"
)

func Provide(ing *do.Injector) {
	do.Provide(ing, myhttp.RouterProvider)
	do.Provide(ing, myhttp.ServerProvider)
	do.Provide(ing, scheduler.ProvideScheduler)
	do.Provide(ing, users.ProvideUsersRepository)
	do.Provide(ing, app.ProvideService)
}
