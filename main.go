package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	client *http.Client
)

func init() {
	client = &http.Client{}
}
func main() {
	t := getToken()
	fmt.Println(t)
	d := getData(t)
	fmt.Println(string(d))
	deal(d)
}

func getData(token string) []byte {
	data, header := make(map[string]int), make(map[string]string)

	data["limit"] = 10
	data["offset"] = 0
	data["portal_type"] = 6
	data["portal_entrance"] = 1
	Pdata, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	header["Host"] = "xskydata.jobs.feishu.cn"
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:93.0) Gecko/20100101 Firefox/93.0"
	header["Accept"] = "application/json, text/plain, */*"
	header["Content-Type"] = "application/json"
	header["website-path"] = "school"
	header["x-csrf-token"] = token[0:43] + "="
	header["Origin"] = "https://xskydata.jobs.feishu.cn"
	header["Connection"] = "keep-alive"
	header["Referer"] = "https://xskydata.jobs.feishu.cn/school/"
	header["Cookie"] = "device-id=; channel=saas-career; platform=pc; atsx-csrf-token=" + token + "; SLARDAR_WEB_ID=9c873871-2f6e-4f8a-b1dc-15c0d210101a; s_v_web_id=verify_kvjh4p79_xh1FZjvN_22lc_4Ssy_AJYl_cqSdxgGA07XJ"

	request, _ := http.NewRequest("POST", "https://xskydata.jobs.feishu.cn/api/v1/search/job/posts", bytes.NewBuffer(Pdata))
	for k, v := range header {
		request.Header.Set(k, v)
	}

	res, _ := client.Do(request)
	resb, _ := io.ReadAll(res.Body)
	return resb
}

func getToken() string {
	data, header := make(map[string]int), make(map[string]string)

	data["portal_entrance"] = 1

	Pdata, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	header["Host"] = "xskydata.jobs.feishu.cn"
	header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:93.0) Gecko/20100101 Firefox/93.0"
	header["Accept"] = "application/json, text/plain, */*"
	header["Content-Type"] = "application/json"
	header["website-path"] = "school"
	header["x-csrf-token"] = "undefined"
	header["Origin"] = "https://xskydata.jobs.feishu.cn"
	header["Connection"] = "keep-alive"
	header["Referer"] = "https://xskydata.jobs.feishu.cn/school/"

	request, _ := http.NewRequest("POST", "https://xskydata.jobs.feishu.cn/api/v1/csrf/token", bytes.NewBuffer(Pdata))
	for k, v := range header {
		request.Header.Set(k, v)
	}

	res, _ := client.Do(request)
	return string(res.Cookies()[0].Value)
}

func deal(data []byte) {
	a := &responseD{}
	json.Unmarshal(data, a)
	saveD := make([]map[string]string, len(a.Data.JobPostList))
	for k, v := range a.Data.JobPostList {
		// saveD1:=make(map[]type)
		saveD[k] = map[string]string{}
		saveD[k]["id"] = v.ID
		saveD[k]["title"] = v.Title
		// saveD[k]["sub_title"] = v.SubTitle
		saveD[k]["description"] = v.Description
		saveD[k]["requirement"] = v.Requirement
		saveD[k]["job_category"] = v.JobCategory.Name
		saveD[k]["city_info"] = v.CityInfo.Name
		saveD[k]["recruit_type"] = v.RecruitType.Name
		saveD[k]["publish_time"] = time.Unix(v.PublishTime, 0).Local().Format("2006-01-02 15:04:05")
		// saveD[k]["job_hot_flag"] = v.JobHotFlag
		saveD[k]["job_subject"] = v.JobSubject.Name.ZhCn
		saveD[k]["code"] = v.Code
		// saveD[k]["department_id"] = v.DepartmentID.
	}
	fmt.Println(saveD)
}
