package main
import (
	"net"
	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
)
func main() {


	for {
		opcion := 0
		fmt.Println("1. Add User")
		fmt.Println("2. Search User")
		fmt.Println("3. Delete User")
		fmt.Print("Opcion: ")
		reader := bufio.NewReader(os.Stdin)
		optxt,_:= reader.ReadString('\n')
		opcion,_ = strconv.Atoi(strings.TrimSpace(string(optxt)))
		switch opcion {
		case 1:
			conn, _ := net.Dial("tcp", "127.0.0.1:399")
			conn.Write([]byte(optxt))
			message:=""
			for message!="Success\n"{
				message, _ = bufio.NewReader(conn).ReadString('\n')
				if message!="Success\n" {
					fmt.Print(message)
					text, _ := reader.ReadString('\n')
					conn.Write([]byte(text))
				}else if message=="Success\n"{
					opcion = -1
					conn.Close()
				}
			}
		case 2:
			fmt.Print("not this case")

		}
	}
}
