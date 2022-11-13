package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// type Greeting func(name string) string

// func say(g Greeting, n string) { fmt.Println(g(n)) }

// func french(name string) string { return "Bonjour, " + name }

// func main() {
//         english := func(name string) string { return "Hello, " + name }

//         say(english, "ANisus")
//         say(french, "ANisus")
// }

func main() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req.Header.Add("Accept", `application/json`)
	req.Header.Add("Authorization", fmt.Sprintf("token %s", os.Getenv("GITHUB_TOKEN")))
	resp, err := c.Do(req)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	fmt.Printf("Body : %s \n ", body)
	fmt.Printf("Response status : %s \n", resp.Status)
}
