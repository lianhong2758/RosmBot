package zero

import (
	"bytes"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/windows"
)

type myfoFormatter struct {
	HasColor bool
}

func init() {
	stdin := windows.Handle(os.Stdin.Fd())

	var mode uint32
	err := windows.GetConsoleMode(stdin, &mode)
	if err != nil {
		panic(err)
	}

	mode &^= windows.ENABLE_QUICK_EDIT_MODE // 禁用快速编辑模式
	mode |= windows.ENABLE_EXTENDED_FLAGS   // 启用扩展标志

	mode &^= windows.ENABLE_MOUSE_INPUT    // 禁用鼠标输入
	mode |= windows.ENABLE_PROCESSED_INPUT // 启用控制输入

	mode &^= windows.ENABLE_INSERT_MODE                           // 禁用插入模式
	mode |= windows.ENABLE_ECHO_INPUT | windows.ENABLE_LINE_INPUT // 启用输入回显&逐行输入

	mode &^= windows.ENABLE_WINDOW_INPUT           // 禁用窗口输入
	mode &^= windows.ENABLE_VIRTUAL_TERMINAL_INPUT // 禁用虚拟终端输入

	err = windows.SetConsoleMode(stdin, mode)
	if err != nil {
		panic(err)
	}

	stdout := windows.Handle(os.Stdout.Fd())
	err = windows.GetConsoleMode(stdout, &mode)
	if err != nil {
		panic(err)
	}

	mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING // 启用虚拟终端处理
	mode |= windows.ENABLE_PROCESSED_OUTPUT            // 启用处理后的输出

	err = windows.SetConsoleMode(stdout, mode)
	log.SetFormatter(&myfoFormatter{HasColor: err == nil})
	if err != nil {
		log.Errorln("设置有色输出失败,默认输出无色")
	}
}

func (m myfoFormatter) Format(entry *log.Entry) ([]byte, error) {
	var color int
	switch entry.Level {
	case log.ErrorLevel:
		color = 1 //red
	case log.WarnLevel:
		color = 3 //yellow
	case log.InfoLevel:
		color = 2 //green
	case log.DebugLevel:
		color = 5
	default:
		color = 7 //白
	}
	var buff *bytes.Buffer
	if entry.Buffer == nil {
		buff = new(bytes.Buffer)
	} else {
		buff = entry.Buffer
	}
	//时间
	formatTime := entry.Time.Format("15:04:06")
	//设置格式
	if m.HasColor {
		fmt.Fprintf(buff, "\033[3%dm[%s]\033[0m %s %s\n", color, entry.Level, formatTime, entry.Message)
	} else {
		fmt.Fprintf(buff, "[%s] %s %s\n", entry.Level, formatTime, entry.Message)
	}
	return buff.Bytes(), nil
}
