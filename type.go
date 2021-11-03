package main

type responseD struct {
	Code int `json:"code"`
	Data struct {
		JobPostList []struct {
			ID          string      `json:"id"`
			Title       string      `json:"title"`
			SubTitle    interface{} `json:"sub_title"`
			Description string      `json:"description"`
			Requirement string      `json:"requirement"`
			JobCategory struct {
				ID       string `json:"id"`
				Name     string `json:"name"`
				EnName   string `json:"en_name"`
				I18NName string `json:"i18n_name"`
				Depth    int    `json:"depth"`
				Parent   struct {
					ID       string `json:"id"`
					Name     string `json:"name"`
					EnName   string `json:"en_name"`
					I18NName string `json:"i18n_name"`
					Depth    int    `json:"depth"`
					Parent   struct {
						ID       string      `json:"id"`
						Name     string      `json:"name"`
						EnName   string      `json:"en_name"`
						I18NName string      `json:"i18n_name"`
						Depth    int         `json:"depth"`
						Parent   interface{} `json:"parent"`
						Children interface{} `json:"children"`
					} `json:"parent"`
					Children interface{} `json:"children"`
				} `json:"parent"`
				Children interface{} `json:"children"`
			} `json:"job_category"`
			CityInfo struct {
				Code         string      `json:"code"`
				Name         string      `json:"name"`
				EnName       string      `json:"en_name"`
				LocationType interface{} `json:"location_type"`
				I18NName     string      `json:"i18n_name"`
				PyName       interface{} `json:"py_name"`
			} `json:"city_info"`
			RecruitType struct {
				ID       string `json:"id"`
				Name     string `json:"name"`
				EnName   string `json:"en_name"`
				I18NName string `json:"i18n_name"`
				Depth    int    `json:"depth"`
				Parent   struct {
					ID           string      `json:"id"`
					Name         string      `json:"name"`
					EnName       string      `json:"en_name"`
					I18NName     string      `json:"i18n_name"`
					Depth        int         `json:"depth"`
					Parent       interface{} `json:"parent"`
					Children     interface{} `json:"children"`
					ActiveStatus int         `json:"active_status"`
				} `json:"parent"`
				Children     interface{} `json:"children"`
				ActiveStatus int         `json:"active_status"`
			} `json:"recruit_type"`
			PublishTime int64       `json:"publish_time"`
			JobHotFlag  interface{} `json:"job_hot_flag"`
			JobSubject  struct {
				ID   string `json:"id"`
				Name struct {
					ZhCn string      `json:"zh_cn"`
					EnUs interface{} `json:"en_us"`
					I18N string      `json:"i18n"`
				} `json:"name"`
				LimitCount   interface{} `json:"limit_count"`
				ActiveStatus int         `json:"active_status"`
			} `json:"job_subject"`
			Code         string      `json:"code"`
			DepartmentID interface{} `json:"department_id"`
		} `json:"job_post_list"`
		Count int    `json:"count"`
		Extra string `json:"extra"`
	} `json:"data"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}
