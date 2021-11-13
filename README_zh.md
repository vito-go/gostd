# daemon

以后台进程方式运行项目程序（仅仅支持linux以及mac os平台）.

## 用法

`import _ "github.com/liushihao1993/daemon"`
运行时候添加命令行参数  `-daemon`
示例:
hello_world.go

```go
package main

import (
	"os"
	"strconv"
	"time"

	_ "github.com/liushihao1993/daemon"
)

func main() {
	f, err := os.Create("hello.log")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for i := 0; i < 10; i++ {
		f.WriteString(strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
```

```shell
$ go build ./hello_world.go
./hello_world -daemon
# 现在你可以退出shell会话， hello_world程序会继续执行下去 并看到hello.log会被创建，并会从1-10写入这个10个数字到该文件中(每秒中写一个).
```