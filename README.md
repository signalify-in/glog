# Glog 

(This repo is a hard fork of Tele [https://github.com/4FR4KO-POVELECKO/glog]).

Glog is a small console,file and Telegram Bot logger for Go.

Logging in:
- Telegram bot (in-progress)
- File ```.log```
- Terminal

## Installation

```bash
go get -u github.com/signalify.in/glog
```

## Examples

First, create bot in [BotFather](https://telegram.me/BotFather).

Start:
```go
package main

import (
	"github.com/signalify.in/glog"
)

func main() {
	// Create new logger
	logger := glog.New()

	// Log to file 
	dirLevels := []glog.Level{ // choose levels
		glog.Panic,
		glog.Fatal,
		glog.Error,
		glog.Warning,
		glog.Info,
		glog.Debug,
		glog.Trace,
	}

	err := logger.NewDir("./log", dirLevels)
	if err != nil {
		logger.Error(err)
	}

	// Log to telegram bot
	botLevels := []glog.Level{ // choose levels
		glog.Panic,
		glog.Fatal,
		glog.Error,
		glog.Warning,
		glog.Info,
		glog.Debug,
		glog.Trace,
	}

	err = logger.NewBot("BOT_TOKEN", botLevels)
	if err != nil {
		logger.Error(err)
	}
}

```

Usage:
```go
logger.Log("./log/trace.Log", glog.Info, "text")
logger.Trace("text")
logger.Debug("text")
logger.Info("text")
logger.Error("text")
logger.Warning("text")
logger.Fatal("text")
logger.Panic("text")
```

Levels:
```go
package glog

type Level string

const (
	Panic   Level = "PANIC:"
	Fatal   Level = "FATAL:"
	Error   Level = "ERROR:"
	Warning Level = "WARNING:"
	Info    Level = "INFO:"
	Debug   Level = "DEBUG:"
	Trace   Level = "TRACE:"
)
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)