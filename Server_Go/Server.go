package main
import (
	"fmt"
	"net"
	"bufio"
	"strings"
	"strconv"
	"regexp"
	"encoding/json"
	"io/ioutil"
	"os"
)

const _VALID_EMAIL = `\w[-._\w]*\w@\w[-._\w]*\w\.\w{2,3}`
const _VALID_ID = `(\d{4}\-\d{4}\-\d{5})`
const _VALID_DATE = `(?:0?[1-9]|[1-2]\d|3[01])\/(?:0?[1-9]|1[0-2])\/\d{4}`

func check(e error) int{
	if e != nil {
		fmt.Print(os.Stderr,e)
		if strings.Contains(e.Error(),"An existing connection was forcibly closed") {
			return 2
		}
		return 1
	}
	return 0
}

func WriteFile(u []User) bool{
	usersFile, _ := json.Marshal(u)
	err := ioutil.WriteFile("users.txt",usersFile,0644)
	if err!=nil {
		return false
	}
	return true
}

func ReadFile() []User{
	var users []User
	usersF,err := ioutil.ReadFile("users.txt")
	if check(err)==0 {
		err = json.Unmarshal(usersF,&users)
		return users
	}else{
		fmt.Print("Error While Loading Users..")
		return nil
	}


}
func EvalRegex(re,value string) bool{
	regex, err := regexp.Compile(re)
	check(err)
	if !regex.MatchString(value) {
		return false
	}
	return true
}

type User struct  {
	Username,Name,Email,Id,F_nac, Foto string
}


func display(u User)string {
	return "Username: " + u.Username  + " Name: " + u.Name  + " Email: " + u.Email + " ID: " + u.Id + " Birthdate: " + u.Foto + " Profile Picture: " + u.Foto +"\n"
}

func Unique(users []User,param string, paramtype int ) bool{
	for _, element := range users {
		if paramtype==0 {
			if element.Username == param {
				return false
			}
		}else if paramtype==1 {
			if element.Email == param {
				return false
			}
		}else if paramtype==2 {
			if element.Id == param{
				return false
			}
		}
	}
	return true
}

func SearchUser(users[]User,username string) int{
	for i,element:= range users{
		if element.Username == strings.Split(username,"\r\n")[0]{
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":399")

	var users []User
	res:=ReadFile()
	if res!=nil {
		users = res
	}

	for {
		  conn, _ := ln.Accept()
		  message, _ := bufio.NewReader(conn).ReadString('\n')
		  fmt.Print("Message Received:", string(message))
		  opcion,err:= strconv.Atoi(strings.TrimSpace(string(message)))
		if err!= nil {
			fmt.Print(os.Stderr,err)
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
					username, err = bufio.NewReader(conn).ReadString('\n')
					if check(err) ==2{
						WriteFile(users)
						os.Exit(0)
					}
					username = strings.Split(username,"\r\n")[0]
					if Unique(users,username,0) && username!="" {
						break
					}
				}
				for {
					conn.Write([]byte("Enter Name : " + "\n"))
					name, err = bufio.NewReader(conn).ReadString('\n')
					if check(err) ==2{
						WriteFile(users)
						os.Exit(0)
					}
					name = strings.Split(name,"\r\n")[0]
					if name!="" {
						break
					}
				}
				for {
					conn.Write([]byte("Enter E-mail : " + "\n"))
					email,err= bufio.NewReader(conn).ReadString('\n')
					if check(err) ==2{
						WriteFile(users)
						os.Exit(0)
					}
					email = strings.Split(email,"\r\n")[0]
					if EvalRegex(_VALID_EMAIL,email)  && Unique(users,email,1) && email!=""{
						break
					}
				}
				for {
					conn.Write([]byte("Enter ID : " + "\n"))
					id, err = bufio.NewReader(conn).ReadString('\n')
					if check(err) ==2{
						WriteFile(users)
						os.Exit(0)
					}
					id = strings.Split(id,"\r\n")[0]
					if id!="" && EvalRegex(_VALID_ID,id) && Unique(users,id,2)   {
						break
					}
				}
				for {
					conn.Write([]byte("Enter Birth Date : " + "\n"))
					bday, err = bufio.NewReader(conn).ReadString('\n')
					if check(err) ==2{
						WriteFile(users)
						os.Exit(0)
					}
					bday = strings.Split(bday,"\r\n")[0]
					if EvalRegex(_VALID_DATE,bday) && bday!="" {
						break
					}
				}
				for {
					conn.Write([]byte("Enter Profile Image : " + "\n"))
					img, err = bufio.NewReader(conn).ReadString('\n')
					if check(err) ==2{
						WriteFile(users)
						os.Exit(0)
					}
					img = strings.Split(img,"\r\n")[0]
					if(img!=""){
						break
					}
				}
				if(img!="" && name!="" && email!="" && id!="" && bday!="" && username!=""){
				user:= User{username,name,email,id,bday,img}
				users = append(users,user)
				}
				conn.Write([]byte("Success\n"))

			case 2:
				conn.Write([]byte("Enter User Name: " + "\n"))
				search, err := bufio.NewReader(conn).ReadString('\n')
				if check(err) ==2{
					WriteFile(users)
					os.Exit(0)
				}else {
					pos := SearchUser(users, search)
					if pos >= 0 {
						conn.Write([]byte(display(users[pos])))
					}else{
						conn.Write([]byte("not Found!" + "\n"))
					}

			    }
			case 3:
				conn.Write([]byte("Enter User Name: " + "\n"))
				search, err := bufio.NewReader(conn).ReadString('\n')
				if check(err) ==2{
					WriteFile(users)
					os.Exit(0)
				}else {
					pos := SearchUser(users, search)
					if pos >= 0 {
						users = append(users[:pos],users[pos+1:]...)
						conn.Write([]byte("Success \n"))
						WriteFile(users)
					}else{
						conn.Write([]byte("not Found!" + "\n"))
					}

				}

			case 5:
				b,err:= json.Marshal(users)
				check(err)
				fmt.Print(string(b))
				if WriteFile(users) {
						fmt.Print("Succesfully Saved")
				}
			}
		}

	}
}

