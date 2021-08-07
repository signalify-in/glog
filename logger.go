package glog

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Logger struct {
	Bot       *telegrambot
	BotLevels []Level

	DirPath   string
	DirLevels []Level
	Lot       bool // Lot - Log on Terminal
	LogPrefix string
}

func New(level uint8, otherArgs ...string) *Logger {
	dirLevels := []Level{
		Panic,
		Fatal,
		Error,
		Warning,
		Info,
		Debug,
		Trace,
	}
	/* botLevels := []Level{
		Panic,
		Fatal,
		Error,
	} */
	prefix := ""
	if len(otherArgs) == 1 {
		prefix = otherArgs[0]
	}
	if level >= uint8(len(dirLevels)) {
		return &Logger{DirLevels: dirLevels, LogPrefix: prefix}
	}
	return &Logger{DirLevels: dirLevels[:level], LogPrefix: prefix}
}

func (l *Logger) NewBot(token string, chatId int, levels []Level) error {
	bot, err := new(token, chatId)
	if err != nil {
		l.Error(err)
		return errors.New("Error to connect the bot")
	}
	bot.ChatID = int64(chatId)
	l.Bot = bot
	l.BotLevels = levels
	return nil
}

func (l *Logger) NewDir(path string, levels []Level) error {
	result, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			l.Error(err)
			return errors.New("Error to create directory")
		}

		l.DirPath = path
		l.DirLevels = levels

		return nil
	}
	if result.IsDir() {
		l.DirPath = path
		l.DirLevels = levels
		return nil
	}
	l.Error(err)
	return errors.New("Error to create directory, a file with the same name already exists")
}

// Loggers
func (l *Logger) LogToFile(path string, level Level, args ...interface{}) {
	write := l.getLogStr(level, args...)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		f, err := os.Create(path)
		if err != nil {
			l.Error(errors.New("Error to create file"), err)
			return
		}

		defer f.Close()
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		l.Error(errors.New("Error to open file"), err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(write)
	if err != nil {
		l.Error(errors.New("Error to write file"), err)
		return
	}
}

func (l *Logger) Log(path string, level Level, args ...interface{}) {
	result := l.checkToArray(level, l.DirLevels)
	if result {
		l.LogToFile(path, level, args...)
	}

	result = l.checkToArray(level, l.BotLevels)
	if result {
		msg := l.getLogStr(level, args...)
		l.Bot.Send(msg)
	}
	if l.Lot {
		write := l.getLogStr(level, args...)
		fmt.Println(write)
	}
}

func (l *Logger) Trace(args ...interface{}) {
	l.Log("./log/trace.Log", Trace, args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.Log("./log/debug.Log", Debug, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.Log("./log/info.Log", Info, args...)
}

func (l *Logger) Warning(args ...interface{}) {
	l.Log("./log/warn.Log", Warning, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.Log("./log/error.Log", Error, args...)
}

func (l *Logger) Notify(args ...interface{}) {
	l.Log("./log/notify.Log", Notify, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.Log("./log/fatal.Log", Fatal, args...)
	os.Exit(1)
}

func (l *Logger) Panic(args ...interface{}) {
	l.Log("./log/panic.Log", Panic, args...)
	panic(args)
}

// Helpers
func (l *Logger) checkToArray(level Level, array []Level) bool {
	for _, value := range array {
		if value == level {
			return true
		}
	}
	return false
}

func (l *Logger) getLogStr(level Level, args ...interface{}) string {
	now := time.Now().Format("2006.01.02 15:04:05")
	head := fmt.Sprint(args[0])
	for _, v := range args[1:] {
		head = strings.Replace(head, "%v", fmt.Sprint(v), 1)
	}
	// due to go's fmt directive reports for
	// extra %v error we have to eliminate extra %v
	str := head //fmt.Sprintf(fmtDirective, args...)
	if l.LogPrefix == "" {
		return fmt.Sprintf("[%s] %s %s \n", now, level, str)
	} else {
		return fmt.Sprintf("[%s] [%s] %s %s \n", now, l.LogPrefix, level, str)
	}
}
