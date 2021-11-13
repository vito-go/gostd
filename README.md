# daemon

runs the computer program as a background process.

## Usage

just `import _ "github.com/liushihao1993/daemon"` in the program and run with command argument `-daemon`
For example:  
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
$ ./hello_world -daemon
# then you can exit the shell, and hello_world will go on.
```