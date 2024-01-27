package providers

import (
	"github.com/samber/do"
	"gocourse15/pkg/myhttp"
	"gocourse15/pkg/scheduler"
)

func Provide(ing *do.Injector) {
	do.Provide(ing, myhttp.RouterProvider)
	do.Provide(ing, myhttp.ServerProvider)
	do.Provide(ing, scheduler.ProvideScheduler)
}
