package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		host string
		port int
		help bool
	)
	flag.StringVar(&host, "H", "127.0.0.1", "主机")
	flag.IntVar(&port, "P", 22, "端口号")
	//flag.BoolVar(&help,"s",false,"帮助")
	fmt.Println(help)
	//flag.Usage = func() {
	//	fmt.Println("使用说明")
	//	flag.PrintDefaults()
	//
	//}
	flag.Parse()

	//if help {
	//	flag.Usage()
	//	return
	//}
	//fmt.Println(host,port,help)

}
