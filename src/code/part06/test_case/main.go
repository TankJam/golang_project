package main

/*
单元测试
	- 测试函数:
		- 格式: 函数名前缀为 Test
		- 作用: 测试程序的一些逻辑行为是否正确
	- 基准函数:
		- 格式: 函数前缀为Benchmark
		- 作用: 测试函数的性能
	- 示例函数:
		- 格式: 函数前缀为Example
		- 作用: 为文档提供示例文档

go test  执行测试函数
	- 会遍历所有的Test开头的函数去测试

go test -v  在t.Run() 可以查看详情

go test -cover  查看测试覆盖率，百分比，测试基本上 60% 就很不错了

go test -cover -coverprofile=c.out  将覆盖率信息保存到文c.out文件中

go tool cover -html=c.out   通过html来打开c.out文件，绿色的代码是被覆盖的

go test -bench=Split -benchmem  // 执行性能测试
	BenchmarkSplit-8：操作系统的核数
	25067053：执行了2000多万次
	46.8 ns/op：每次操作耗费了46.8 纳秒
	16 B/op：每次操作消耗 16个字节
	1 allocs/op：每次做三次的内存申请，申请内存会非常耗时

	- 根据返回的参数对代码进行优化
*/


