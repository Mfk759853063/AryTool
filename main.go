package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"github.com/atotto/clipboard"
)

func main() {
	digital := flag.Int64("d", 0, "输入一个10进制数")
	hexString := flag.String("h", "", "输入一个16进制数")
	needCopy := flag.Bool("c",false,"将结果复制到剪切板")
	flag.Parse()
	if *digital != 0 {
		var d int64 = *digital
		hexString := strconv.FormatInt(d, 16)
		fmt.Printf("0x%s\n", hexString)
		if *needCopy {
			hexString = fmt.Sprintf("0x%s", hexString)
			clipboard.WriteAll(hexString)
		}
	} else if *hexString != "" {
		if strings.HasPrefix(*hexString, "0x") {
			*hexString = strings.Replace(*hexString, "0x", "", 1)
		}
		if parseNumber, err := strconv.ParseInt(*hexString, 16, 0); err != nil {
			fmt.Printf("error : %@", err)
		} else {
			fmt.Printf("%d\n", parseNumber)
			if *needCopy {
				parseStr := fmt.Sprintf("%d", parseNumber)
				clipboard.WriteAll(parseStr)
			}
		}

	}
}
