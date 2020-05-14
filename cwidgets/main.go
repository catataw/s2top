package cwidgets

import (
	"github.com/catataw/s2top/logging"
	"github.com/catataw/s2top/models"
)

var log = logging.Init()

type WidgetUpdater interface {
	SetMeta(string, string)
	SetMetrics(models.Metrics)
}
