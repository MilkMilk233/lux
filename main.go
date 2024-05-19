package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"

	"github.com/iawia002/lux/app"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: lux [flag] [urls]")
		return
	}

	flag := os.Args[1]
	switch flag {
	case "-1":
		// 执行与"-1"相关的操作
		urls := os.Args[1:]
		fmt.Println(urls) // 添加这行来打印urls
		if err := app.New().Run(urls); err != nil {
			fmt.Fprintf(
				color.Output,
				"Run %s failed: %s\n",
				color.CyanString("%s", app.Name), color.RedString("%v", err),
			)
			os.Exit(1)
		}
	case "-2":
		// 执行与"-2"相关的操作
		urls := os.Args[1:]
		fmt.Println(urls) // 添加这行来打印urls
		if err := app.New().Run(urls); err != nil {
			fmt.Fprintf(
				color.Output,
				"Run %s failed: %s\n",
				color.CyanString("%s", app.Name), color.RedString("%v", err),
			)
			os.Exit(1)
		}
	case "-3":
		// 执行与"-3"相关的操作 优先级下载
		// num, err := strconv.Atoi(os.Args[2])
		// if err != nil {
		// 	fmt.Println("Invalid num:", os.Args[2])
		// 	return
		// }
		// urls := os.Args[3 : 3+num]
		// priorities := os.Args[3+num:]
	default:
		fmt.Println("Invalid flag:", flag)
	}
}
