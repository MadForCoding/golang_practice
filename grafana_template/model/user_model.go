package model

type UserCustomPanel struct {
	GlobalDataSource string           // M 指定source, 可以是面板变量${datasource}, 也可以是写死的。 全局
	PanelList        []*UserEachPanel // M
}

type UserEachPanel struct {
	PanTitle       string
	DataSource     string // 指定source, 可以是面板变量${datasource}, 也可以是写死的。 panel级别 高于全局
	UserExpr       *UserExpr
	UserAlert      *UserAlert
	UserThresholds *UserThresholds
}

type UserExpr struct {
	Expr   string // C panel表达式
	Legend string // O 展示出来的名字， 也可以用${}取变量 可选

	ExprFormat       string   // C panel表达式, 变量替换裸格式
	ExprFormatVar    [][]any  // C 变量替换内容
	ExprLegendFormat []string // O 命别名， 要和ExprFormatVar的一维保持一致，

}

type UserAlert struct {
	Conditions            []Conditions      // M 评估语句
	ExecutionErrorState   AlarmExecErrState // O
	For                   string            // O 持续了for的时间还是pengding， 就会将状态转为alarm, 并且告警
	Frequency             string            // O 每分钟评估一次
	Message               string            // M 告警出来的消息内容
	Name                  string            // M 标题名字
	NoDataState           AlarmNoDataState  // O
	NotificationChannelID string            // O
}

// UserThresholds - 控制面板上的线
//
// 默认不需要传, 会有逻辑从alert里面读取你第一个
type UserThresholds struct {
	OP    string  // M
	Value float64 // M
}
