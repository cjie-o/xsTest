package main

import (
	"testing"
)

func TestDeal(t *testing.T) {
	json := `{ "code": 0, "data": { "job_post_list": [{ "id": "7019185011899238695", "title": "产品设计师", "sub_title": null, "description": "", "job_category": { "id": "6791702736615344397", "name": "设计", "en_name": "Design", "i18n_name": "设计", "depth": 2, "parent": { "id": "6791698585114724616", "name": "互联网 / 电子 / 网游", "en_name": "Internet / Electronics / Games", "i18n_name": "互联网 / 电子 / 网游", "depth": 1, "parent": null, "children": null }, "children": null }, "city_info": { "code": "CT_11", "name": "北京", "en_name": "Beijing", "location_type": null, "i18n_name": "北京", "py_name": null }, "recruit_type": { "id": "202", "name": "实习", "en_name": "Intern", "i18n_name": "实习", "depth": 2, "parent": { "id": "2", "name": "校招", "en_name": "Campus", "i18n_name": "校招", "depth": 1, "parent": null, "children": null, "active_status": 1 }, "children": null, "active_status": 1 }, "publish_time": 1634281519814, "job_hot_flag": null, "job_subject": { "id": "7002102973326362919", "name": { "zh_cn": "2022年校招项目", "en_us": null, "i18n": "2022年校招项目" }, "limit_count": null, "active_status": 0 }, "code": "9552", "department_id": null } ], "count": 21, "extra": "{\"fe_tracking\":{\"log_id\":\"202111031820230102121810322F006E5C\",\"query_length\":0,\"total\":21}}" }, "message": "ok", "error": null }`
	// got := make(map[string]string)
	// got["name"] = "ghj"

	got := deal([]byte(json))
	if got[0]["title"] != "产品设计师" {
		t.Error("deal([]byte(json))[0][title] get ", got[0]["title"], "want 产品设计师")
	}
}
