package ueditor

import (
	"net/http"
	"strconv"
	"io/ioutil"
	"path"
	"sort"
	"encoding/json"
)

//排序
type ItemSortList []map[string]interface{}

func (self ItemSortList) Len() int  {
	return len(self)
}

func (self ItemSortList)  Less(i, j int) bool  {
	var iItem = self[i]
	var jItem = self[j]
	var iTime = iItem["mtime"].(int64)
	var jTime = jItem["mtime"].(int64)
	return iTime > jTime
}

func (self ItemSortList)  Swap(i, j int)   {
	self[i], self[j] = self[j], self[i]
}
//结束


func getLimit(r *http.Request) (startVal int,sizeVal int) {
	var start = r.FormValue("start")
	var size = r.FormValue("size")
	if start != "" {
		startVal,_ = strconv.Atoi(start)
	} else {
		startVal = 0
	}
	if size != "" {
		sizeVal,_ = strconv.Atoi(size)
	} else {
		sizeVal = 20
	}
	return startVal, sizeVal
}


func getAllList( pathString string )  []map[string]interface{} {
	var picPathList = make([]map[string]interface{},0)
	var fileList, e = ioutil.ReadDir(pathString)
	if e != nil {
		return picPathList
	}
	for _, v := range fileList {
		if v.IsDir() {
			continue
		}
		var picPath  = path.Join("/", pathString,v.Name())
		var mtime = v.ModTime().Unix()
		var item = make(map[string]interface{})
		item["url"] = picPath
		item["mtime"] = mtime
		picPathList = append(picPathList, item)
	}
	sort.Sort(ItemSortList(picPathList))
	return picPathList
}

func getList( r *http.Request, pathString string ) (  reData []map[string]interface{}, reStart int, reTotal int ) {
	var picPathList = getAllList(pathString)
	var total = len(picPathList)
	var start, size = getLimit(r)
	var end =  start + size
	if end > total{
		end = total
	}
	if start >= end {
		return nil,0,0
	}
	picPathList = picPathList[start:end]
	return picPathList, start, total
}

func listResult(picPathList []map[string]interface{},start int, total int , w http.ResponseWriter)  {
	var result = make(map[string]interface{})
	result["state"] = "SUCCESS"
	result["list"] = picPathList
	result["start"] = start
	result["total"] = total

	var jsonBytes,_ = json.Marshal(result)
	w.Write([]byte(jsonBytes))
}

