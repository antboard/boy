package main

import "fmt"

var Commit string
var BuildDate string
var GitTag string
var GitBranch string
var GoVersion string

func main() {
	fmt.Println("Commit:", Commit)
	fmt.Println("BuildDate:", BuildDate)
	fmt.Println("GitTag:", GitTag)
	fmt.Println("GitBranch:", GitBranch)
	fmt.Println("GoVersion:", GoVersion)
	// 启动api服务 或者启动 grpc服务
}
