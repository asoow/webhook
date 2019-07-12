package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "传参需要header：%v和get请求http://localhost:10086/?program=%v", "X-Gitlab-Token", "项目名")

	cfg, _ := goconfig.LoadConfigFile("conf.ini")
	Token, _ := cfg.GetValue("config", "token")

	tokens, yes := r.Header["X-Gitlab-Token"]
	if !yes || len(tokens) < 1 {
		log.Println("Header param 'X-Gitlab-Token' is missing")
		return
	}

	token := string(tokens[0])
	if !strings.EqualFold(token, Token) {
		log.Println("Header param 'X-Gitlab-Token' is error")
		return
	}

	programs, ok := r.URL.Query()["program"]
	if !ok || len(programs) < 1 {
		log.Println("Url Param 'program' is missing")
		return
	}
	program := string(programs[0])

	execShell(program)

}

func execShell(program string) {

	command := fmt.Sprintf("shell/%v.sh", program)

	if _, err := os.Stat(command); os.IsNotExist(err) {
		fmt.Printf("Shell file %v is not exists", command)
		return
	}
	cmd := exec.Command("/bin/bash", "-c", command)

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))

}

func main() {

	http.HandleFunc("/", index)

	err := http.ListenAndServe(":10086", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
