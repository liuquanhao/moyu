package page_data_controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gofiber/websocket/v2"
	"github.com/liuquanhao/moyu/model"
	"github.com/liuquanhao/moyu/model/remote_url_model"
)

// define: page/backend/service/page_data.go
type PageData struct {
	Uptime        uint64  `json:"uptime"`
	CPUPercent    float64 `json:"cpu_percent"`
	MemoryPercent float64 `json:"memory_percent"`
	DiskPercent   float64 `json:"disk_percent"`
	NetSendTotal  uint64  `json:"net_send"`
	NetRecvTotal  uint64  `json:"net_recv"`
	Timestamp     int64   `json:"timestamp"`
}

type PageInfo struct {
	PageId   uint32    `json:"page_id"`
	PageData *PageData `json:"page_data"`
}

func getPageData(myClient *http.Client, url string, target interface{}) error {
	resp, err := http.Get(url + "api/page_data")
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)
}

func loopPullPageData(pageId uint32, url string, ch chan *PageInfo) {
	var myClient = &http.Client{Timeout: 5 * time.Second}
	pageData := new(PageData)
	for {
		start := time.Now()
		err := getPageData(myClient, url, pageData)
		if err != nil {
			time.Sleep(3 * time.Second)
		} else {
			ch <- &PageInfo{
				PageId:   pageId,
				PageData: pageData,
			}
		}
		end := time.Now()
		useTime := end.Sub(start)
		if useTime < 1*time.Second {
			time.Sleep(1000*time.Millisecond - useTime)
		}
	}
}

func PushPageData(c *websocket.Conn) {
	db := model.DBConn
	var pageUrls []remote_url_model.PageUrl
	db.Find(&pageUrls)
	ch := make(chan *PageInfo, len(pageUrls))
	defer close(ch)
	for _, pageUrl := range pageUrls {
		go loopPullPageData(pageUrl.Id, pageUrl.PageUrl, ch)
	}
	for pageInfo := range ch {
		c.WriteJSON(pageInfo)
	}
}
