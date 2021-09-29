go 覆盖测试

localhost :: package/utiltest/cover ‹master*› » go test --cover                                                                                   1 ↵
PASS
coverage: 60.0% of statements
ok  	go-notes/package/utiltest/cover	0.435s

输出覆盖详情
localhost :: package/utiltest/cover ‹master*› » go test -coverprofile=size_coverage.out
PASS
coverage: 60.0% of statements
ok  	go-notes/package/utiltest/cover	0.196s

按函数覆盖率统计
localhost :: package/utiltest/cover ‹master*› » go tool cover -func=size_coverage.out
go-notes/package/utiltest/cover/cover.go:3:	Size		42.9%
go-notes/package/utiltest/cover/cover.go:20:	Sum		100.0%
total:						(statements)	60.0%

运行此命令时，浏览器将弹出窗口，已覆盖（绿色），未覆盖（红色）和 未埋点（灰色）
go tool cover -html=size_coverage.out

go test 命令接受 -covermode 标志将覆盖模式设置为三种设置之一：
set: 每个语句是否执行？是默认设置
count: 每个语句执行了几次？
atomic: 类似于 count, 但表示的是并行程序中的精确计数

go test  -covermode=count -coverprofile=size_count.out
go tool cover -html=size_count.out

对具体的报进行测试覆盖率
go test -covermode=count -coverprofile=fmt_count.out fmt



https://blog.csdn.net/u012855229/article/details/51930709/

https://studygolang.com/articles/33818

