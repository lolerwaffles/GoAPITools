package apitools

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Header Map Example
// var headersmap = map[string]string{
// 	"HeaderType": "HeaderData",
// }

func CallAPIReturnString(URL string, HeadersMap map[string]string) string {

	client := &http.Client{}
	request, err := http.NewRequest("GET", URL, nil)

	if err != nil {
		log.Fatalln(err)
	}
	if len(HeadersMap) > 0 {
		for key, value := range HeadersMap {
			request.Header.Set(key, value)
		}
	}
	resp, err := client.Do(request)

	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}
