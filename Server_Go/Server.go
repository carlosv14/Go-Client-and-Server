package main
import (
	"fmt"
	"net"
	"bufio"
	"strings"
	"strconv"
	"log"
	"regexp"
	"encoding/json"
)

const _VALID_EMAIL = `\w[-._\w]*\w@\w[-._\w]*\w\.\w{2,3}`
const _VALID_ID = `(\d{4}\-\d{4}\-\d{5})`
const _VALID_DATE = `(?:0?[1-9]|[1-2]\d|3[01])\/(?:0?[1-9]|1[0-2])\/\d{4}`
func EvalRegex(re,value string) bool{
	regex, err := regexp.Compile(re)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if !regex.MatchString(value) {
		return false
	}
	return true
}

type User struct  {
	username,name,email,id,f_nac, foto string
}


func display(u User) {

	fmt.Println("Username: " + u.username)
	fmt.Println("Name: " + u.name)
	fmt.Println("Email: " + u.email)
	fmt.Println("ID: " + u.id)
	fmt.Println("Birthdate: " + u.f_nac)
	fmt.Println("Profile Picture: " + u.foto)
}

func Unique(users []User,param string, paramtype int ) bool{
	for _, element := range users {
		if paramtype==0 {
			if element.username == param {
				return false
			}
		}else if paramtype==1 {
			if element.email == param {
				return false
			}
		}else if paramtype==2 {
			if element.id == param{
				return false
			}
		}
	}
	return true
}


func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":399")

	var users []User
	cont:=0
	for {
		  conn, _ := ln.Accept()
		  message, _ := bufio.NewReader(conn).ReadString('\n')
		  fmt.Print("Message Received:", string(message))
		  opcion,err:= strconv.Atoi(strings.TrimSpace(string(message)))
		if err!= nil {
			log.Fatal(err)
		}else{
			switch opcion {
			case 1:

				username :=""
				email:=""
				name:=""
				id:=""
				bday:=""
				img:=""
				for {
					conn.Write([]byte("Enter Username: " + "\n"))
					username, _ = bufio.NewReader(conn).ReadString('\n')
					if Unique(users,username,0) && username!="" {
						break
					}
				}
				for {
					conn.Write([]byte("Enter Name : " + "\n"))
					name, _ = bufio.NewReader(conn).ReadString('\n')
					if name!="" {
						break
					}
				}
				for {
				conn.Write([]byte("Enter E-mail : " + "\n"))
				email,_= bufio.NewReader(conn).ReadString('\n')
					if EvalRegex(_VALID_EMAIL,email)  && Unique(users,email,1) && email!=""{
						break
					}
				}
				for {
					conn.Write([]byte("Enter ID : " + "\n"))
					id, _ = bufio.NewReader(conn).ReadString('\n')
					if id!="" && EvalRegex(_VALID_ID,id) && Unique(users,id,2)   {
						break
					}
				}
				for {
					conn.Write([]byte("Enter Birth Date : " + "\n"))
					bday, _ = bufio.NewReader(conn).ReadString('\n')
					if EvalRegex(_VALID_DATE,bday) && bday!="" {
						break
					}
				}
				for {
					conn.Write([]byte("Enter Profile Image : " + "\n"))
					img, _ = bufio.NewReader(conn).ReadString('\n')
					if(img!=""){
						break
					}
				}
				if(img!="" && name!="" && email!="" && id!="" && bday!="" && username!=""){
				user:= User{username,name,email,id,bday,img}
				users = append(users,user)
				cont = cont+1
				fmt.Print(cont)
					if cont==2 {
						b,_ := json.Marshal(users)
						fmt.Print(b)

					}
				}
				conn.Write([]byte("Success\n"))
			}
		}

		 }
}

