package libs

import (
	"encoding/json"
	"fmt"
	"github.com/xiaoJack/beegoAceAdmin/common"
)

type Ztree struct {

}

func GetProjectZtreeByProjectId(projectId int, projectName string)(ztreeJson string, err error)  {

	//树的数据结构
	//	[
	//	{ id:1, pId:0, name:"APP", open:true},
	//	{ id:3, pId:1, name:"个人中心", open:true},
	//	{ id:11, pId:3, name:"叶子333"},
	//	{ id:4, pId:1, name:"简历", open:true},
	//	{ id:21, pId:4, name:"叶子节点 2-1"},
	//	{ id:22, pId:4, name:"叶子节点 2-2"},
	//	{ id:23, pId:4, name:"叶子节点 2-3"},
	//	{ id:5, pId:1, name:"首页", open:true},
	//	{ id:31, pId:5, name:"叶子节点 3-1"},
	//	{ id:32, pId:5, name:"叶子节点 3-2"},
	//	{ id:33, pId:5, name:"叶子节点 3-3"},
	//];


	var project_label common.Project_label
	var project_api common.Project_api

	label,err := project_label.GetListByProjectId(projectId)
	if err != nil {
		return "", err
	}

	apiList,err := project_api.GetListByProjectId(projectId)
	if err != nil {
		return "", err
	}

	list := make([]map[string]interface{}, len(label)+len(apiList)+1)

	//固定项目名称
	top := make(map[string]interface{})
	top["id"] = projectId
	top["pId"] = 0
	top["name"] = projectName
	top["open"] = true
	list[0] = top



	//项目标签
	for k, v := range label {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["pId"] = v.Project_id
		row["name"] = v.Label_name
		row["open"] = true

		list[k+1] = row
	}

	//项目接口列表
	for k,v := range apiList {
		row := make(map[string]interface{})
		//row["id"] = "a_" + string(v.Id)
		row["id"] = fmt.Sprintf("%d_%d", v.Project_label_id, v.Id)
		row["pId"] = v.Project_label_id
		row["name"] = v.Api_name
		row["open"] = true

		list[k+len(label)+1] = row
	}



	jsonStr, err := json.Marshal(list)
	if err != nil {
		return "", err
	}

	return string(jsonStr), nil
}