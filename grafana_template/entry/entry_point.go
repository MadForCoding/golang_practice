package entry

import (
	"golang_practice/grafana_template/model"
	"golang_practice/grafana_template/tool"
	"io/ioutil"
	"math/rand"
	"time"
)

// DashBoardWithPanelsByTemplateGeneration - main entry point
//
// panelExistForUpdate - false: 根据title 如果存在在读取的json文件里面，不会被创建
// true: 如果有相同title的面板存在, 会被覆盖
func DashBoardWithPanelsByTemplateGeneration(userCustom *model.UserCustomPanel, templateRelativePath string, panelExistForUpdate bool) string {
	rand.Seed(time.Now().Unix())
	dashBoard := initDashBoardFromFile(templateRelativePath)

	panelsAddList := convertUserDefinePanels(userCustom)
	for _, addPanel := range panelsAddList {
		dashBoard.PanelAppend(addPanel, panelExistForUpdate)
	}
	dashBoard.PreCheckPanelsIDWithFillUp()

	jsonStr := tool.JsonTool.JsonEncoding(dashBoard, false, true)
	return jsonStr
}

func initDashBoardFromFile(relativePath string) *model.AutoGenerated {
	dashBoardByte, err := ioutil.ReadFile(relativePath)
	if err != nil {
		panic(err)
	}
	dashBoard := model.InitDashBoardByJson(dashBoardByte)
	return dashBoard
}

func convertUserDefinePanels(userCustom *model.UserCustomPanel) []*model.Panels {
	var globalDataSource = userCustom.GlobalDataSource
	var panelList []*model.Panels
	for _, customPanel := range userCustom.PanelList {
		var (
			currentDataSource = globalDataSource
			advancePanel      = false
		)

		if len(customPanel.DataSource) != 0 {
			currentDataSource = customPanel.DataSource
		}
		if customPanel.UserAlert != nil {
			advancePanel = true
		}
		targets := model.InitTargetByCustom(customPanel.UserExpr)
		if advancePanel {
			alarm := model.InitAlarmByCustom(customPanel.UserAlert)
			threshold := model.InitThresholdsByCustom(customPanel.UserThresholds)
			p := model.InitPanelWithAlarm(currentDataSource, customPanel.PanTitle, alarm, targets, threshold)
			panelList = append(panelList, p)
		} else {
			p := model.InitPanelBasic(currentDataSource, customPanel.PanTitle, targets)
			panelList = append(panelList, p)
		}
	}
	return panelList
}
