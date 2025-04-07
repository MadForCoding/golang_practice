package model

import (
	"fmt"
	"golang_practice/grafana_template/tool"
)

func InitPanelBasic(dataSource string, panelTitle string, targets []Targets) *Panels {
	// 这里假定每个panel只有一个prometheus语句
	// 告警是只能在一个prometheus语句上的原因
	panel := &Panels{
		AliasColors: AliasColors{},
		Bars:        false,
		DashLength:  10,
		Dashes:      false,
		Datasource:  dataSource,
		FieldConfig: FieldConfig{
			Defaults:  Defaults{},
			Overrides: []any{},
		},
		Fill:         1,
		FillGradient: 0,
		GridPos: GridPos{
			H: 8,
			W: 12, // total 24 for each line
			X: 0,  // pos x
			Y: 0,  // pos y
		},
		HiddenSeries: false,
		ID:           0, // 会在最终model 那边赋值了， 要防止重复
		Legend: Legend{
			AlignAsTable: true,
			Avg:          true,
			Current:      true,
			Max:          true,
			Min:          false,
			RightSide:    false,
			Show:         true,
			Total:        true,
			Values:       true,
		},
		Lines:         true,
		Linewidth:     1,
		NullPointMode: "null",
		Options: Options{
			AlertThreshold: true,
		},
		Percentage:      false,
		PluginVersion:   "7.5.17",
		Pointradius:     2,
		Points:          false,
		Renderer:        "flot",
		SeriesOverrides: []any{},
		SpaceLength:     10,
		Stack:           false,
		SteppedLine:     false,
		Targets:         targets,
		Thresholds:      []Thresholds{}, // 这个配置告警的话， 在alarm时候再加上吧， 面板上的那条红线
		TimeFrom:        nil,
		TimeRegions:     []any{},
		TimeShift:       nil,
		Title:           panelTitle,
		Tooltip: Tooltip{
			Shared:    true,
			Sort:      0,
			ValueType: "individual",
		},
		Type: "graph",
		Xaxis: Xaxis{
			Buckets: nil,
			Mode:    "time",
			Name:    nil,
			Show:    false,
			Values:  []any{},
		},
		Yaxes: []Yaxes{
			{
				Format:  "short",
				Label:   nil,
				LogBase: 1,
				Max:     nil,
				Min:     0,
				Show:    true,
			},
			{
				Format:  "short",
				Label:   nil,
				LogBase: 1,
				Max:     nil,
				Min:     nil,
				Show:    true,
			},
		},
		Yaxis: Yaxis{
			Align:      false,
			AlignLevel: nil,
		},
	}
	return panel
}

func InitPanelWithAlarm(dataSource string, panelTitle string, alarm *Alert, targets []Targets, thresholds []Thresholds) *Panels {
	panelWithAlarm := InitPanelBasic(dataSource, panelTitle, targets)
	panelWithAlarm.Alert = alarm
	panelWithAlarm.Targets = targets
	panelWithAlarm.Thresholds = thresholds
	return panelWithAlarm
}

func InitTargetByCustom(userExpr *UserExpr) []Targets {
	if userExpr == nil {
		panic("empty userExpr")
	}
	if len(userExpr.Expr) == 0 && len(userExpr.ExprFormat) == 0 {
		panic("both expr and exprFormat empty")
	}
	if len(userExpr.Expr) != 0 && len(userExpr.ExprFormat) != 0 {
		panic("both expr and exprFormat has value")
	}
	f := func(e string, l string) Targets {
		return Targets{
			Exemplar:     true,
			Expr:         e,
			Interval:     "",
			LegendFormat: l,
			RefID:        tool.GenerateChar.GetNextCharacter(),
		}
	}
	checkFunc := func(formatVarLen int, formatLegendLen int) {
		if formatLegendLen == 0 {
			return
		}
		if formatVarLen != formatLegendLen {
			panic(fmt.Sprintf("formatVarLen(%d) not equal to formatLegendLen(%d)", formatVarLen, formatLegendLen))
		}
	}
	if len(userExpr.Expr) != 0 {
		target := f(userExpr.Expr, userExpr.Legend)
		tool.GenerateChar.Reset()
		return []Targets{target}
	}
	if len(userExpr.ExprFormat) != 0 {
		var tList []Targets
		checkFunc(len(userExpr.ExprFormatVar), len(userExpr.ExprLegendFormat))
		for i, valueList := range userExpr.ExprFormatVar {
			exprStr := fmt.Sprintf(userExpr.ExprFormat, valueList...)
			target := f(exprStr, userExpr.ExprLegendFormat[i])
			tList = append(tList, target)
		}
		tool.GenerateChar.Reset()
		return tList
	}
	panic("weird path")
}

func InitThresholdsByCustom(thresholds *UserThresholds) []Thresholds {
	if thresholds == nil {
		panic("nil thresholds")
	}
	if len(thresholds.OP) == 0 {
		panic("userThresholds M field not fulfilled")
	}
	data := Thresholds{
		ColorMode: "critical",
		Fill:      true,
		Line:      true,
		Op:        thresholds.OP,
		Value:     thresholds.Value,
		Visible:   true,
	}
	return []Thresholds{data}
}

func InitAlarmByCustom(userAlarm *UserAlert) *Alert {
	if userAlarm == nil {
		panic("nil userAlarm")
	}
	if len(userAlarm.Conditions) == 0 || len(userAlarm.Message) == 0 || len(userAlarm.Name) == 0 {
		panic("userAlarm all M field not fulfilled")
	}
	type UserAlert struct {
		Conditions            []Conditions      // M 评估语句
		ExecutionErrorState   AlarmExecErrState // O
		For                   string            // O 持续了for的时间还是pengding， 就会告警
		Frequency             string            // O 每分钟评估一次
		Message               string            // M 告警出来的消息内容
		Name                  string            // M  标题名字
		NoDataState           AlarmNoDataState  // O
		NotificationChannelID string            // O
	}

	alert := initDefaultAlert()
	// M field
	alert.Conditions = userAlarm.Conditions
	alert.Message = userAlarm.Message
	alert.Name = userAlarm.Name

	// O field
	if len(userAlarm.ExecutionErrorState) != 0 {
		alert.ExecutionErrorState = userAlarm.ExecutionErrorState.Raw()
	}
	if len(userAlarm.For) != 0 {
		alert.For = userAlarm.For
	}
	if len(userAlarm.Frequency) != 0 {
		alert.Frequency = userAlarm.Frequency
	}
	if len(userAlarm.NoDataState) != 0 {
		alert.NoDataState = userAlarm.NoDataState.Raw()
	}
	if len(userAlarm.NotificationChannelID) != 0 {
		// TODO: 这里可以优化, 现在假设都是一个告警组
		alert.Notifications = []Notifications{{UID: userAlarm.NotificationChannelID}}
	}
	return alert
}

func initDefaultAlert() *Alert {
	return &Alert{
		AlertRuleTags:       AlertRuleTags{},
		Conditions:          nil,
		ExecutionErrorState: AlarmExecErrState_KeepState.Raw(),
		For:                 "5m",
		Frequency:           "1m",
		Handler:             1,
		Message:             "",
		Name:                "",
		NoDataState:         AlarmNoDataState_Alarming.Raw(),
		Notifications:       []Notifications{{UID: AlarmChannel}},
	}
}
