package apputils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func Post(L, a, b string) string {
	url := "http://www.easyrgb.com/php/query_convert.php"
	lab := fmt.Sprintf("DAT1=%s&DAT2=%s&DAT3=%s&DAT4=", L, a, b)
	param := "SAV0=&SAV1=&SAV2=&SAV3=&SAV4=&SAV5=&SAV6=&SAV7=&SAV8=&SAV9=&SAVA=&SAVB=&SAVC=&SAVD=&SAVE=&SAVF=&RES0=&RES1=&RES2=&RES3=&RES4=&RES5=&RES6=&RES7=&RES8=&RES9=&RESA=&RESB=&RESC=&RESD=&RESE=&RESF=&RESX=&SPCE=LAB&ILLU=D65&OBSE=2&" + lab
	payload := strings.NewReader(param)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Cookie", "EasyRGBuser=lv1airr3vlnkk6jmf73glareo0")
	req.Header.Add("Origin", "http://www.easyrgb.com")
	req.Header.Add("Referer", "http://www.easyrgb.com/en/convert.php")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36")
	req.Header.Add("sec-gpc", "1")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(string(body))
	return string(body)
}

func GetFrom(str string) (string, string) {
	re := regexp.MustCompile(`Yxy\s*=\s*(\d+\.\d+)\s+(\d+\.\d+)\s+(\d+\.\d+)`)

	// Find the first match in the input string
	match := re.FindStringSubmatch(str)

	// Extract Yxy values from the match
	if len(match) > 0 {
		Y := match[1]
		x := match[2]
		y := match[3]
		fmt.Printf("Y=%s, x=%s, y=%s\n", Y, x, y)
		return x, y
	}
	return "", ""
}
