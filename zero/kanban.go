package zero

import (
	"fmt"
	"time"
)

var banner string = "* Villa + RosmBot + Golang\n" +
	"* " + time.Now().Format("2006-01-02 :15:04:06") + "+0800 CST\n" +
	"* Project: https://github.com/lianhong2758/RosmBot"
var line string = "-----------------------------------------------------"

func Kanban() {
	fmt.Println(line)
	fmt.Println(banner)
	fmt.Println(line)
}
