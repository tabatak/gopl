// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// my answer
		// if !strings.Contains(url, "http://") {
		// 	url = "http://" + url
		// }
		// answer by shibata
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		fmt.Fprintf(os.Stdout, "status code: %d", resp.StatusCode)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		// 	os.Exit(1)
		// }
		// fmt.Printf("%s", b)

		io.Copy(os.Stdout, resp.Body)
	}
}
