package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os/exec"
	"regexp"

	"google.golang.org/grpc/codes"
)

var re = regexp.MustCompile(`Get the answer: (\d+ [\+\-\*\/\%]{1,2} \d+) = \?`)

func main() {
	str := codes.OK.String()
	fmt.Println(str)
	return

	conn, err := net.Dial("tcp", "ppc-ctf.o3.ru:7090")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	r := bufio.NewReader(conn)

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		sm := re.FindStringSubmatch(line)
		if len(sm) > 1 {
			expr := sm[1]

			byteRes, err := exec.Command("python3", "-c", "print("+expr+")").Output()
			if err != nil {
				log.Fatal(err)
			}

			res := string(byteRes)

			fmt.Println(res)

			_, err = fmt.Fprintf(conn, res)
			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Println(line)
	}
}
