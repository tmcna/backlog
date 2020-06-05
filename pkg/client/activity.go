package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// Activity is ...
type Activity struct {
	space   string
	apiKey  string
	query   url.Values
	next    bool
	lastID  int
	restnum int
	order   int
}

// NewActivity is constructor.
func NewActivity(space string, apiKey string, dispnum int, order int) *Activity {
	act := new(Activity)
	act.space = space
	act.apiKey = apiKey
	act.query = url.Values{}
	act.next = true
	act.lastID = 0
	act.restnum = dispnum
	act.order = order

	return act
}

// Query - Query parameter
func (act *Activity) Query(key string, value string) {
	act.query.Add(key, value)
}

// List function returns recent updates in your space.
func (act *Activity) List() ([]ActivityResponse, error) {
	api := "api/v2/space/activities"

	var count int
	if act.restnum == 0 || act.restnum > ActivityResponseCount {
		count = ActivityResponseCount //1回のAPIで取り出す件数
	} else {
		count = act.restnum //NewActivity()で指定された数
	}
	act.query.Set("count", strconv.Itoa(count)) //取得数を追加
	if act.order == DisplayOrderDesc {
		act.query.Set("order", "desc") //並び順を設定
	} else {
		act.query.Set("order", "asc") //並び順を設定
	}

	if act.lastID != 0 {
		if act.order == DisplayOrderDesc {
			act.query.Set("maxId", strconv.Itoa(act.lastID)) //降順の場合は最大IDを設定(2回目以降のみ)
		} else {
			act.query.Set("minId", strconv.Itoa(act.lastID)) //昇順の場合は最小IDを設定(2回目以降のみ)
		}
	}
	cli := NewClient(act.space, act.apiKey)
	body, err := cli.Get(api, act.query)
	if err != nil {
		return nil, err
	}

	//JSON形式で出力（デバッグ用）
	//fmt.Println(string(body))

	var r []ActivityResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	// 取得できた課題数が、指定した数より小さければ全取得と判断し、nextフラグをfalseにする。
	// 取得数=指定数の場合一度空振りするが頻度が少ないため許容する。
	if len(r) < count {
		act.next = false
	}
	if len(r) != 0 {
		// 空振りでない場合のみインデックスをずらすようにする。
		index := len(r)
		act.lastID = r[index-1].ID
	}
	act.restnum -= len(r)
	if act.restnum <= 0 {
		act.next = false
	}

	return r, nil
}

// HasNext is ...
func (act *Activity) HasNext() bool {
	return act.next
}

// PrintCSV prints list of activities in CSV format.
func (act *Activity) PrintCSV(r []ActivityResponse) {
	for _, n := range r  {
		fmt.Printf("%d,%s,%s,%s,%d,%d\n",
			n.ID,
			n.Project.ProjectKey,
			n.Project.Name,
			n.User.Name,
			n.Type,
			n.Reason)
	}
}
