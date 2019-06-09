package admin

import (
	"dc-kanban/controller"
	"dc-kanban/dataStruct"
	"dc-kanban/lib/curl"
	"dc-kanban/lib/logger"
	"dc-kanban/model"
	"dc-kanban/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strings"
	"sync"
)

type HomeController struct{ controller.Base }

func (t *HomeController) Index(c *gin.Context) {

	moduleList := new(service.ModuleService).GetAll()
	var moduleIds = make([]uint, 0)
	for _, v := range moduleList {
		moduleIds = append(moduleIds, v.ID)
	}

	var itemList []model.ChartItem
	cate := new(service.CategoryService).GetListByModuleIds(moduleIds)
	if len(cate) > 0 {
		var cateIds = make([]uint, 0)
		for _, v := range cate {
			cateIds = append(cateIds, v.ID)
		}

		itemList = new(service.CategoryItemService).GetListByCateIds(cateIds)
	}

	t.Data(c, gin.H{
		"module":   moduleList,
		"category": cate,
		"item":     itemList,
	})
}

func (t *HomeController) Note(c *gin.Context) {

	t.Data(c, gin.H{
		"moduleList": new(service.ModuleService).GetAll(),
		"noteList":   new(service.NoteService).GetAll(),
	})
}

func (t *HomeController) IndustryData(c *gin.Context) {

	// 5分钟

	var vals = [3]string{
		`search=search index="index_jira" 问题类型 IN ("任务" "子任务") 
| stats max(_time) latest(*) AS * by 标识
| search 经办人 IN ("未分配","") 状态!="完成" 
| append 
    [ search index="index_jira" 问题类型 IN ("任务" "子任务") 
    | stats max(_time) latest(*) AS * by 标识
    | search 状态="待办"]
| dedup 标识
| rename 更新 AS _update_time 
| rename 优先级 as priority 
| eval update_time=strptime( _update_time,"%Y/%m/%d %H:%M") 
| eval fifteendaysago=relative_time(now(),"-15d@d")
| where (update_time >= fifteendaysago) OR ((priority = "高优先级" OR priority = "排除万难") AND update_time<fifteendaysago)
| search 项目="*"
| rename 子任务 as subtask
| eval subtask_tag=if(subtask!="",0,1)
| search subtask_tag=1
| fields - subtask,subtask_tag
| search sprint_state="ACTIVE"
| fields - sprint_state
|stats count&output_mode=json`, // 待办任务

		`search=search index="index_jira" 问题类型="缺陷"
	| stats max(_time) latest(*) AS * by 标识
	| search 项目="*"
	| stats count&output_mode=json`, // 任务缺陷

		`search=search index="index_app"  source="/opt/dcei/etc/apps/jira/bin/collecterapp.py"
	| stats max(_time) latest(*) AS * by 标识
	| search 问题类型="应用详情"
	| fillnull value="无" 应用对接优先级
	| stats count by 应用对接优先级&output_mode=json`, // 应用优先级
	}

	uri := `https://192.168.5.85:8089/services/search/jobs/export?`

	var data = make([]interface{}, 3)

	var wg sync.WaitGroup
	for k, v := range vals {
		wg.Add(1)
		go func(u string, k int) {
			b, err := curl.Get(u)
			if err != nil {
				t.Err(c, err.Error())
				return
			}
			arr := strings.Split(string(b), "\n")
			var origin = make([]map[string]interface{}, 0)
			for _, vv := range arr {
				var source map[string]interface{}
				if vv != "" {
					if err := json.Unmarshal([]byte(vv), &source); err != nil {
						logger.Info("Json Error:", err.Error())
						return
					}
					origin = append(origin, source)
				}
			}
			data[k] = origin
			wg.Done()
		}(uri+v, k)
	}

	wg.Wait()

	t.Data(c, data)
}

func (t *HomeController) OperateData(c *gin.Context) {

	// 15s
	ipList := map[string][]string{
		"ATS":     {"192.168.176.111", "192.168.176.112", "192.168.176.113"},
		"Captain": {"10.12.10.1"},
		"IOG":     {"192.168.111.2"},
		"NAT":     {"172.16.10.1", "172.16.10.2"},
		"NDC":     {"10.6.10.1", "10.6.10.2", "10.6.10.3", "10.6.10.4", "10.6.10.5"},
		"Public":  {"192.168.1.101"},
		"TIANFU":  {"10.12.10.1", "10.15.10.1"},
	}
	uri := `http://192.168.170.18:9090/api/v1/query?query=esxi_cpu_usage`

	var cpuData = make(map[string]interface{})
	for k, v := range ipList {
		for _, vv := range v {
			params := `{cluster="` + k + `",id="` + vv + `",exported_job="cpu_usage",instance="pushgateway",job="pushgateway"}`
			b, err := curl.Get(uri + params)
			var res dataStruct.CPUData
			if err == nil {
				if err := json.Unmarshal(b, &res); err == nil && res.Status == "success" && len(res.Data.Result) == 1 {
					cpuData[vv] = res.Data.Result[0].Value
					continue
				}
			}
			cpuData[vv] = []int{0, 0}
		}
	}

	uriI := `http://192.168.170.18:9090/api/v1/query?query=ifHCInOctets{ifAlias="eth2",ifDescr="WAN2",ifIndex="16777218",ifName="WAN2",instance="192.168.1.1",job="snmp",tag="DaoCloud_Route"}`
	uriO := `http://192.168.170.18:9090/api/v1/query?query=ifHCOutOctets{ifAlias="eth2",ifDescr="WAN2",ifIndex="16777218",ifName="WAN2",instance="192.168.1.1",job="snmp",tag="DaoCloud_Route"}`

	var ioData = make([]interface{}, 2)
	ioData[0] = []int{0, 0}
	ioData[1] = []int{0, 0}

	var vi dataStruct.IOData
	bi, err := curl.Get(uriI)
	if err == nil {
		if err := json.Unmarshal(bi, &vi); err == nil && vi.Status == "success" && len(vi.Data.Result) == 1 {
			ioData[0] = vi.Data.Result[0].Value
		}
	}

	var vo dataStruct.IOData
	bo, err := curl.Get(uriO)
	if err == nil {
		if err := json.Unmarshal(bo, &vo); err == nil && vo.Status == "success" && len(vo.Data.Result) == 1 {
			ioData[1] = vo.Data.Result[0].Value
		}
	}

	t.Data(c, gin.H{
		"cpu": cpuData,
		"io":  ioData,
	})
}

func (t *HomeController) IOData(c *gin.Context) {

	uriI := `http://192.168.170.18:9090/api/v1/query?query=ifHCInOctets{ifAlias="eth2",ifDescr="WAN2",ifIndex="16777218",ifName="WAN2",instance="192.168.1.1",job="snmp",tag="DaoCloud_Route"}[2h]&step=30s`
	uriO := `http://192.168.170.18:9090/api/v1/query?query=ifHCOutOctets{ifAlias="eth2",ifDescr="WAN2",ifIndex="16777218",ifName="WAN2",instance="192.168.1.1",job="snmp",tag="DaoCloud_Route"}[2h]&step=30s`

	bi, err := curl.Get(uriI)
	if err != nil {
		t.Err(c, "input获取错误")
		return
	}

	var vi dataStruct.IODataSet
	if err := json.Unmarshal(bi, &vi); err != nil {
		t.Err(c, "input解析错误")
		return
	}

	bo, err := curl.Get(uriO)
	if err != nil {
		t.Err(c, "output获取错误")
		return
	}

	var vo dataStruct.IODataSet
	if err := json.Unmarshal(bo, &vo); err != nil {
		t.Err(c, "output解析错误")
		return
	}

	var iData interface{}
	var oData interface{}
	if vi.Status == "success" && len(vi.Data.Result) == 1 {
		iData = vi.Data.Result[0].Values
	}
	if vo.Status == "success" && len(vo.Data.Result) == 1 {
		oData = vo.Data.Result[0].Values
	}

	t.Data(c, gin.H{
		"input":  iData,
		"output": oData,
	})
}
