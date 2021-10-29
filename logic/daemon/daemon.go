// Package daemon http://vitogo.tpddns.cn:9000/liushihao/daemon
package daemon

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func init() {
	if runtime.GOOS == "windows" {
		fmt.Println("Warning! on the windows platform, ignore daemon.")
		return
	}
	daemon := flag.Bool("daemon", false, "to run it as a full daemon. only support for linux and mac os")
	defaultDlogPath := filepath.Base(os.Args[0]) + ".log"
	dlog := flag.String("dlog", defaultDlogPath, `specify the daemon log path`)
	var argNoDaemon []string
	var dlogParsed bool
	for n, a := range os.Args {
		if strings.Contains(a, "-daemon") {
			*daemon = true
			continue
		}
		if !dlogParsed && strings.Contains(a, "-dlog") {
			// 兼容go自带的参数解析方式
			if strings.Contains(a, "-dlog=") {
				// dlog 直接带=
				ss := strings.Split(a, "=")
				if len(ss) == 0 {
					fmt.Println("dlog parse error. please input the right format. e.g. -dlog=hello.log, --dlog=/home/me/hello.log")
					os.Exit(1)
				}
				*dlog = ss[1]
				dlogParsed = true
				continue
			}
			// dlog 不带等号
			if n+1 >= len(os.Args) {
				fmt.Println("dlog parse error. dlog needs an argument. please input the right format. e.g. -dlog=hello.log, --dlog=/home/me/hello.log")
				os.Exit(1)
			}
			if strings.HasPrefix(os.Args[n+1], "-") {
				fmt.Println("dlog parse error. dlog needs an argument, and dlog can not begin with -.\n\tplease input the right format. e.g. -dlog=hello.log, --dlog=/home/me/hello.log")
				os.Exit(1)
			}
			*dlog = os.Args[n+1]
			dlogParsed = true
		}
		argNoDaemon = append(argNoDaemon, a)
	}
	if *daemon {
		f, err := os.Create(*dlog)
		if err != nil {
			panic(err)
		}
		cmd := exec.Command(argNoDaemon[0], argNoDaemon[1:]...)
		cmd.Stdout = f
		cmd.Stderr = f
		if err = cmd.Start(); err != nil {
			fmt.Printf("start %s failed, error: %v\n", os.Args[0], err)
			err = f.Close()
			if err != nil {
				panic(err)
			}
			os.Exit(1)
		}
		fmt.Printf("%+v PID: %d running...\n", argNoDaemon, cmd.Process.Pid)
		_, err = f.WriteString(fmt.Sprintf("%s %+v PID: %d\n",
			time.Now().Format("2006/01/02 15:04:05.000"), argNoDaemon, cmd.Process.Pid))
		if err != nil {
			panic(err)
		}
		err = f.Close()
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	}
}
