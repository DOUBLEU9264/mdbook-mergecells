package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// 处理命令行参数
	args := os.Args[1:]

	// 没有命令行参数则处理Stdin内容
	if len(args) < 2 {
		// 从标准输入读取 JSON
		input, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalf("读取输入失败: %v", err)
		}

		processBookContent(input)
		os.Exit(0)
	}

	if args[0] != "supports" || args[1] != "html" {
		os.Exit(1)
	}
	os.Exit(0)
}

func processBookContent(input []byte) {
	// 解析 JSON
	bookData, err := NewBookDataFromJson(input)
	if err != nil {
		log.Fatalf("解析书籍JSON失败: %v", err)
	}

	// 处理书籍内容
	bookData.ProcessBook()

	// 序列化处理后的数据
	output, err := bookData.ToJsonIndent()
	if err != nil {
		log.Fatalf("将书籍重新序列化JSON时失败: %v", err)
	}

	// 输出到标准输出
	fmt.Print(string(output))
}
