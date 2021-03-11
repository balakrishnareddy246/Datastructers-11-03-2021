package main

import(
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
    "time"
)

func main() {

    var name string
    var userRating string

//Front end
reader := bufio.NewReader(os.Stdin)
fmt.Println("please enter your full name")
name, _ = reader.ReadString('\n')
// take rating dorm user convert it to float 
reader = bufio.NewReaderWriter(os.Stdin)
fmt.Println("Please rate our kloudone comapny betweeen 1 and 5 : ")
userRating, _ = reader.ReadString('\n')
mynumRating, _ := strconv.ParseFlaot(string.TrimSpace(userRating),64)



//back end
fmt.Printf("Hello %v,\n Thank for rating our kloudone technologyes %v star rating recorded in our system at %v",
name, myRating, time.Now().Format)

}
