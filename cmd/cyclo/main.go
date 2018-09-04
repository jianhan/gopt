package main

import (
	"fmt"
	"log"
	"os/exec"
)

// func main() {
// 	cmd := exec.Command("/home/jian/go/bin/gocyclo", "~/Documents/test.go")
// 	stdout, err := cmd.StdoutPipe()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := cmd.Start(); err != nil {
// 		log.Fatal(err)
// 	}
// 	var output []byte
// 	n, err := stdout.Read(output)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(output), "test", n)
// }

func main() {
	cmd := exec.Command("/home/jian/go/bin/gocyclo", "/home/jian/Documents/test.go")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}
