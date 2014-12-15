package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
)

func download(url string) {
	tmp := strings.Split(url, "/")
	fileName := tmp[len(tmp)-1]
	fmt.Println("'" + fileName + "' download...")

	response, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error: %s", err.Error())
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error: %s", err.Error())
		return
	}

	ioutil.WriteFile(fileName, body, 0644)
	fmt.Println("done.")
}

func main() {
	dwgo := cli.NewApp()
	dwgo.Name = "dwgo05"
	dwgo.Usage = "dwgo05 [<URI>]"
	dwgo.Action = func(c *cli.Context) {
		wg := new(sync.WaitGroup)
		runtime.GOMAXPROCS(runtime.NumCPU())
		for _, url := range os.Args[1:] {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				download(url)
			}(url)
			wg.Wait()
		}
	}
	dwgo.Run(os.Args)
}
