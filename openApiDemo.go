package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	url := "http://api.letvcloud.com/open.php"
	secretkey := "9d43ac5d51a2360edd0a9b744d624a3e"
	param := make(map[string]string)
	param["api"] = "video.get"
	param["video_id"] = "8243607"
	param["ver"] = "2.0"
	param["user_unique"] = "c191a085a2"
	param["format"] = "json"
	param["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	sign := getSign(param, secretkey)
	requestUrl := createUrl(param, url, sign)
	resp, err := http.Get(requestUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("%s", b)
}

func getSign(param map[string]string, secretkey string) string {

	sorted_keys := make([]string, 0)
	for k, _ := range param {
		sorted_keys = append(sorted_keys, k)
	}

	// sort 'string' key in increasing order
	sort.Strings(sorted_keys)
	str := ""
	for _, k := range sorted_keys {
		str += k + param[k]
	}

	h := md5.New()
	h.Write([]byte(str + secretkey))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func createUrl(param map[string]string, url string, sign string) string {
	newUrl := ""
	for k, v := range param {
		newUrl += k + "=" + v + "&"
	}
	newUrl = url + "?" + newUrl + "sign=" + sign
	return newUrl
}
