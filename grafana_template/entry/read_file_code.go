package entry

//
//import (
//	"bufio"
//	"encoding/json"
//	"fmt"
//	"golang_practice/grafana_template/model"
//	"io"
//	"io/ioutil"
//	"os"
//	"strings"
//)
//
//func readFile() {
//	dashBoardByte, err := ioutil.ReadFile("./grafana_template/entry/dashBoard_template.json")
//	if err != nil {
//		panic(err)
//	}
//	dashBoard := model.InitDashBoardByJson(dashBoardByte)
//	fmt.Println(dashBoard)
//
//	panelList := getPanelList()
//	if len(panelList) == 0 {
//		panic("not thing to add")
//	}
//	for _, p := range panelList {
//		dashBoard.PanelAppend(p, true)
//	}
//	dashBoard.PreCheckPanelsIDWithFillUp()
//	str := &strings.Builder{}
//	encoder := json.NewEncoder(str)
//	// avoid special character parse
//	encoder.SetEscapeHTML(false)
//	encoder.Encode(dashBoard)
//	fmt.Println("===========")
//	fmt.Println(str.String())
//}
//
//func getPanel(strRaw []string) *model.Panels {
//	title := ""
//	dataSourceVar := "Airpay TH"
//	targetStr := ""
//	alarmstr := ""
//	thresholdStr := ""
//	for i, v := range strRaw {
//		if i == 0 {
//			title = v
//		} else if i == 1 {
//			targetStr = v
//		} else if i == 2 {
//			alarmstr = v
//		} else if i == 3 {
//			thresholdStr = v
//		}
//	}
//	if len(title) == 0 {
//		panic("not allow empty title")
//	}
//	if len(targetStr) == 0 {
//		panic("not allow empty expression")
//	}
//	var target []model.Targets
//	var alarm *model.Alert
//	var threshold []model.Thresholds
//	var hasAlarm = false
//	if err := json.Unmarshal([]byte(targetStr), &target); err != nil {
//		panic(err)
//	}
//	if len(alarmstr) != 0 {
//		alarm = &model.Alert{}
//		if err := json.Unmarshal([]byte(alarmstr), alarm); err != nil {
//			panic(err)
//		}
//		hasAlarm = true
//	}
//	if len(thresholdStr) != 0 {
//		if err := json.Unmarshal([]byte(thresholdStr), &threshold); err != nil {
//			panic(err)
//		}
//	}
//
//	if hasAlarm {
//		return model.InitPanelWithAlarm(dataSourceVar, title, alarm, target, threshold)
//	} else {
//		return model.InitPanelBasic(dataSourceVar, title, target)
//	}
//}
//
//func getPanelList() []*model.Panels {
//	panelByteFile, err := os.Open("./grafana_template/entry/add_panels.txt")
//	if err != nil {
//		panic(err)
//	}
//	defer panelByteFile.Close()
//	var panelList []*model.Panels
//	var skipHeader = true
//	panelReader := bufio.NewReader(panelByteFile)
//	for {
//		line, err := panelReader.ReadString('\n')
//		if err == io.EOF {
//			break
//		} else if err != nil {
//			panic(err)
//		}
//		if skipHeader {
//			skipHeader = false
//			continue
//		}
//
//		subStr := strings.Split(line, "@@")
//		if len(subStr) == 0 {
//			continue
//		}
//		panel := getPanel(subStr)
//		panelList = append(panelList, panel)
//	}
//	return panelList
//}
