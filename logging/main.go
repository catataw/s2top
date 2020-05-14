package logging

import (
	"fmt"
	"os"
	"time"

	"github.com/op/go-logging"
)

const (
	size = 1024
)

var (
	Log    *S2TOPLogger
	exited bool
	level  = logging.INFO // default level
	format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
)

type statusMsg struct {
	Text    string
	IsError bool
}

type S2TOPLogger struct {
	*logging.Logger
	backend *logging.MemoryBackend
	sLog    []statusMsg
}

func (c *S2TOPLogger) FlushStatus() chan statusMsg {
	ch := make(chan statusMsg)
	go func() {
		for _, sm := range c.sLog {
			ch <- sm
		}
		close(ch)
		c.sLog = []statusMsg{}
	}()
	return ch
}

func (c *S2TOPLogger) StatusQueued() bool     { return len(c.sLog) > 0 }
func (c *S2TOPLogger) Status(s string)        { c.addStatus(statusMsg{s, false}) }
func (c *S2TOPLogger) StatusErr(err error)    { c.addStatus(statusMsg{err.Error(), true}) }
func (c *S2TOPLogger) addStatus(sm statusMsg) { c.sLog = append(c.sLog, sm) }

func (c *S2TOPLogger) Statusf(s string, a ...interface{}) { c.Status(fmt.Sprintf(s, a...)) }

func Init() *S2TOPLogger {
	if Log == nil {
		logging.SetFormatter(format) // setup default formatter

		Log = &S2TOPLogger{
			logging.MustGetLogger("s2top"),
			logging.NewMemoryBackend(size),
			[]statusMsg{},
		}

		if debugMode() {
			level = logging.DEBUG
			StartServer()
		}

		backendLvl := logging.AddModuleLevel(Log.backend)
		backendLvl.SetLevel(level, "")

		logging.SetBackend(backendLvl)
		Log.Notice("logger initialized")
	}
	return Log
}

func (log *S2TOPLogger) tail() chan string {
	stream := make(chan string)

	node := log.backend.Head()
	go func() {
		for {
			stream <- node.Record.Formatted(0)
			for {
				nnode := node.Next()
				if nnode != nil {
					node = nnode
					break
				}
				if exited {
					close(stream)
					return
				}
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return stream
}

func (log *S2TOPLogger) Exit() {
	exited = true
	StopServer()
}

func debugMode() bool    { return os.Getenv("S2TOP_DEBUG") == "1" }
func debugModeTCP() bool { return os.Getenv("S2TOP_DEBUG_TCP") == "1" }
