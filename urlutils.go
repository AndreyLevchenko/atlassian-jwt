package atlsnjwt

import (
	"crypto/sha256"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

func hashUrl(httpMethod string, url string) (string, error) {
	req, err := canonicalizeRequest(httpMethod, url)
	if err != nil {
		return "", err
	}
	sha256encoded := sha256.Sum256([]byte(req))
	return fmt.Sprintf("%x", sha256encoded), nil
}
func canonicalizeQS(query string) (string, error) {
	/*
		Example of conversion:
		zee_last=param&repeated=parameter 1&first=param&repeated=parameter 2

		first=param&repeated=parameter%201,parameter%202&zee_last=param
	*/

	mm, err := url.ParseQuery(query)
	if err != nil {
		return "", err
	}
	encodedParams := make(map[string]string)
	keys := make([]string, 0, len(mm))
	for name, values := range mm {
		if name == "jwt" {
			continue
		}
		encodedName := url.QueryEscape(name)
		//encode values and sort
		for i, v := range values {
			values[i] = url.QueryEscape(v)
		}
		sort.Strings(values)

		keys = append(keys, encodedName)

		encodedParams[encodedName] = strings.Join(values, ",")
	}
	sort.Strings(keys)

	var result string

	for i, key := range keys {
		var delimiter string

		if i == 0 {
			delimiter = ""
		} else {
			delimiter = "&"
		}

		result += delimiter + key + "=" + encodedParams[key]
	}

	return result, nil
}
func canonicalizeRequest(httpMethod string, requestUrl string) (string, error) {
	requestUrl = strings.TrimLeft(requestUrl, "/")
	u, err := url.Parse(requestUrl)
	if err != nil {
		return "", err
	}

	canonicalURI := "/" + strings.ReplaceAll(strings.Trim(u.Path, "/"), "&", url.QueryEscape("&"))
	canonicalQS, err := canonicalizeQS(u.RawQuery)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s&%s&%s", strings.ToUpper(httpMethod), canonicalURI, canonicalQS), nil
}
