This library extends [JWT-GO](https://github.com/dgrijalva/jwt-go) by adding support of Atlassian's custom QSH (query string hash) claim.
It may be useful for development Atlassian Cloud plugins with Go. Here is example of REST API call with jwt authentication
```Go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	atlsnjwt "github.com/AndreyLevchenko/atlassian-jwt"
)

const (
	appKey       = "org.evoja.jira.todoable" //plugin key
	sharedSecret = "this is secret" //substitute value obtained while plugin installation
	url          = "https://xxx.atlassian.net/rest/api/latest/project"
)

func main() {
	jwt, err := atlsnjwt.Encode("GET", "/rest/api/latest/project", appKey, sharedSecret, 0)
	if err != nil {
		fmt.Printf("encoding error %s\n", err)
	}

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Set("Authorization", fmt.Sprintf("jwt %s", jwt))

	resp, _ := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("Error making request: %s", err)
	}

	b, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}

	fmt.Println(string(b[:]))

}
```
