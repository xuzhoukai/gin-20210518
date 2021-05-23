package tailf

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

var TailfInstance *tail.Tail

func Init() {
	var err error
	fileName := "./my.log"
	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,
	}
	TailfInstance, err = tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail file failed, err:", err)
		return
	}

}

func Run() {
	var (
		line *tail.Line
		ok   bool
	)
	for {
		line, ok = <-TailfInstance.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", TailfInstance.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line:", line.Text)
	}
}
