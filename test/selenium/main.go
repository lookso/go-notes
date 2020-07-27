//package main
//
//import (
//	"fmt"
//	"github.com/tebeka/selenium"
//	"github.com/tebeka/selenium/chrome"
//	"net"
//	"os"
//	"time"
//)
//
//func pickUnusedPort() (int, error) {
//
//	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
//	if err != nil {
//		return 0, err
//	}
//
//	l, err := net.ListenTCP("tcp", addr)
//	if err != nil {
//		return 0, err
//	}
//	port := l.Addr().(*net.TCPAddr).Port
//	if err := l.Close(); err != nil {
//		return 0, err
//	}
//	return port, nil
//}
//
//func main() {
//	port, err := pickUnusedPort()
//	fmt.Println("port", port)
//	opts := []selenium.ServiceOption{
//		selenium.Output(os.Stderr),
//	}
//	selenium.SetDebug(false)
//
//	service, err := selenium.NewChromeDriverService("/home/lukun/data/netcode/chromedriver", port, opts...)
//	if err != nil {
//		panic(err)
//	}
//	defer service.Stop()
//
//	fmt.Println("here 1")
//
//	// 起新线程在新标签页打开窗口
//	go func() {
//		caps := selenium.Capabilities{"browserName": "chrome"}
//		chromeCaps := chrome.Capabilities{
//			Path: "",
//			Args: []string{
//				"--Content-Type=application/json",
//				"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7",
//			},
//		}
//		caps.AddChrome(chromeCaps)
//		wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d", port))
//		if err != nil {
//			panic(err)
//		}
//		defer wd.Quit()
//
//		if err := wd.Get("https://www.baidu.com"); err != nil {
//			panic(err)
//		}
//
//		time.Sleep(time.Second * 10)
//
//		// 在窗口调用js 脚本
//		wd.ExecuteScript(`window.open("https://www.qq.com", "_blank");`, nil)
//
//		time.Sleep(time.Second * 20)
//	}()
//
//	// 打开一个在一个标签页里打开一个窗口
//	caps := selenium.Capabilities{"browserName": "internet explorer"}
//	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d", port))
//	if err != nil {
//		panic(err)
//	}
//	defer wd.Quit()
//
//	// Navigate to the simple playground interface.
//	if err := wd.Get("http://play.golang.org/?simple=1"); err != nil {
//		panic(err)
//	}
//
//	// Get a reference to the text box containing code.
//	elem, err := wd.FindElement(selenium.ByCSSSelector, "#code")
//	if err != nil {
//		panic(err)
//	}
//	// Remove the boilerplate code already in the text box.
//	if err := elem.Clear(); err != nil {
//		panic(err)
//	}
//
//	// Enter some new code in text box.
//	err = elem.SendKeys(`
//		package main
//		import "fmt"
//		func main() {
//			fmt.Println("Hello WebDriver!\n")
//		}
//	`)
//	if err != nil {
//		panic(err)
//	}
//
//	// Click the run button.
//	btn, err := wd.FindElement(selenium.ByCSSSelector, "#run")
//	if err != nil {
//		panic(err)
//	}
//	if err := btn.Click(); err != nil {
//		panic(err)
//	}
//
//	// Wait for the program to finish running and get the output.
//	outputDiv, err := wd.FindElement(selenium.ByCSSSelector, "#output")
//	if err != nil {
//		panic(err)
//	}
//
//	var output string
//	for {
//		output, err = outputDiv.Text()
//		if err != nil {
//			panic(err)
//		}
//		if output != "Waiting for remote server..." {
//			break
//		}
//		time.Sleep(time.Second * 1)
//	}
//
//	fmt.Println("waiting....")
//	time.Sleep(time.Second * 80)
//
//	// Example Output:
//	// Hello WebDriver!
//	//
//	// Program exited.
//}

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

const (
	port = 8080
)

func main() {
	opts := []selenium.ServiceOption{
		// Enable fake XWindow session.
		// selenium.StartFrameBuffer(),
		selenium.Output(os.Stderr), // Output debug information to STDERR
	}

	// Enable debug info.
	// Enable debug info.
	// selenium.SetDebug(true)
	service, err := selenium.NewChromeDriverService("/home/lukun/data/netcode/chromedriver", port, opts...)
	if err != nil {
		panic(err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://127.0.0.1:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	//wd.Get("https://y.qq.com/n/yqq/toplist/27.html#stat=y_new.toplist.menu.27")
	wd.Get("http://learn.itech8.com/")
	fmt.Println(wd.Title())
	elem, err := wd.FindElement(selenium.ByClassName, "home blog wp-custom-logo hfeed")
	if err != nil {
		panic(err)
	}
	fmt.Println(elem.Text())

	time.Sleep(5 * time.Second)
}
