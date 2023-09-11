package zero

import (
	"os"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/windows"
)

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
	if err != nil {
		log.SetFormatter(&ColorNotFormatter{})
		log.Errorln("设置有色输出失败,默认输出无色")
	} else {
		log.SetFormatter(&ColorFormatter{})
	}
}
