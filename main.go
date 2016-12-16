package main

import (
	// "bufio"
	"fmt"
	// "os"
	"net/http"
	"os/exec"
	"strings"
	// "time"
	"runtime"
)

func main() {
	b, err := exec.Command(getUser()).Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Welcome, %s!\n", strings.TrimSpace(string(b)))
	// fmt.Print("Enter password: ")
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	input := scanner.Text()
	// 	if strings.TrimSpace(input) == ".exit" {
	// 		os.Exit(0)
	// 	}
	// 	fmt.Println(input)
	// 	fmt.Println("> ")
	// 	fmt.Println("\033[8m")

	// }

	http.HandleFunc("/", root)

	go open("http://localhost:8080/")
	panic(http.ListenAndServe(":8080", nil))

}

func root(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("wtf?"))
}

func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func getUser() string {
	switch runtime.GOOS {
	case "windows":
		return "whoami"
	case "darwin":
		//
	default: // "linux", "freebsd", "openbsd", "netbsd"
		//
	}

}
