package entity

import (
	"github.com/catataw/s2top/logging"
	"github.com/catataw/s2top/connector/collector"
	"github.com/catataw/s2top/models"
	"github.com/catataw/s2top/cwidgets"
)

var (
	log = logging.Init()
)

type Entity interface {
	SetState(s string)
	Logs() collector.LogCollector
	GetMetaEntity() Meta
	SetUpdater(updater cwidgets.WidgetUpdater)
	GetMeta(v string) string
	GetId() string
	GetMetrics() models.Metrics
}

