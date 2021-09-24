# daemon

runs the computer program as a background process.

## Usage

Just `import _ "gitea.com/liushihao/daemon"` in the program and run with command argument `-daemon`   
For example:  
hello_world.go

```go
package main

import (
	_ "gitea.com/liushihao/daemon"
	"os"
	"strconv"
	"time"
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
# then you can exit the shell, and hello_world will go on.
```