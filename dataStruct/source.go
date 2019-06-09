package dataStruct

type CPUData struct {
	Data struct {
		Result []struct {
			Metric struct {
				Name        string `json:"__name__"`
				Cluster     string `json:"cluster"`
				ExportedJob string `json:"exported_job"`
				ID          string `json:"id"`
				Instance    string `json:"instance"`
				Job         string `json:"job"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
		ResultType string `json:"resultType"`
	} `json:"data"`
	Status string `json:"status"`
}

type IOData struct {
	Data struct {
		Result []struct {
			Metric struct {
				Name     string `json:"__name__"`
				IfAlias  string `json:"ifAlias"`
				IfDescr  string `json:"ifDescr"`
				IfIndex  string `json:"ifIndex"`
				IfName   string `json:"ifName"`
				Instance string `json:"instance"`
				Job      string `json:"job"`
				Tag      string `json:"tag"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
		ResultType string `json:"resultType"`
	} `json:"data"`
	Status string `json:"status"`
}

type IODataSet struct {
	Data struct {
		Result []struct {
			Metric struct {
				Name     string `json:"__name__"`
				IfAlias  string `json:"ifAlias"`
				IfDescr  string `json:"ifDescr"`
				IfIndex  string `json:"ifIndex"`
				IfName   string `json:"ifName"`
				Instance string `json:"instance"`
				Job      string `json:"job"`
				Tag      string `json:"tag"`
			} `json:"metric"`
			Values [][]interface{} `json:"values"`
		} `json:"result"`
		ResultType string `json:"resultType"`
	} `json:"data"`
	Status string `json:"status"`
}