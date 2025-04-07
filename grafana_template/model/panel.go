package model

type Panels struct {
	Alert           *Alert       `json:"alert,omitempty"`
	AliasColors     AliasColors  `json:"aliasColors"`
	Bars            bool         `json:"bars"`
	DashLength      int          `json:"dashLength"`
	Dashes          bool         `json:"dashes"`
	Datasource      string       `json:"datasource"`
	FieldConfig     FieldConfig  `json:"fieldConfig"`
	Fill            int          `json:"fill"`
	FillGradient    int          `json:"fillGradient"`
	GridPos         GridPos      `json:"gridPos"`
	HiddenSeries    bool         `json:"hiddenSeries"`
	ID              int          `json:"id"`
	Legend          Legend       `json:"legend,omitempty"`
	Lines           bool         `json:"lines"`
	Linewidth       int          `json:"linewidth"`
	NullPointMode   string       `json:"nullPointMode"`
	Options         Options      `json:"options"`
	Percentage      bool         `json:"percentage"`
	PluginVersion   string       `json:"pluginVersion"`
	Pointradius     int          `json:"pointradius"`
	Points          bool         `json:"points"`
	Renderer        string       `json:"renderer"`
	SeriesOverrides []any        `json:"seriesOverrides"`
	SpaceLength     int          `json:"spaceLength"`
	Stack           bool         `json:"stack"`
	SteppedLine     bool         `json:"steppedLine"`
	Targets         []Targets    `json:"targets"`
	Thresholds      []Thresholds `json:"thresholds"`
	TimeFrom        any          `json:"timeFrom"`
	TimeRegions     []any        `json:"timeRegions"`
	TimeShift       any          `json:"timeShift"`
	Title           string       `json:"title"`
	Tooltip         Tooltip      `json:"tooltip"`
	Type            string       `json:"type"`
	Xaxis           Xaxis        `json:"xaxis"`
	Yaxes           []Yaxes      `json:"yaxes"`
	Yaxis           Yaxis        `json:"yaxis"`
}

type AliasColors struct {
}
type Defaults struct {
}

type GridPos struct {
	H int `json:"h"`
	W int `json:"w"`
	X int `json:"x"`
	Y int `json:"y"`
}

type FieldConfig struct {
	Defaults  Defaults `json:"defaults"`
	Overrides []any    `json:"overrides"`
}

type Alert struct {
	AlertRuleTags       AlertRuleTags   `json:"alertRuleTags"`
	Conditions          []Conditions    `json:"conditions"`
	ExecutionErrorState string          `json:"executionErrorState"`
	For                 string          `json:"for"`
	Frequency           string          `json:"frequency"`
	Handler             int             `json:"handler"`
	Message             string          `json:"message"`
	Name                string          `json:"name"`
	NoDataState         string          `json:"noDataState"`
	Notifications       []Notifications `json:"notifications"`
}

type Notifications struct {
	UID string `json:"uid"`
}

type AlertRuleTags struct {
}

type Evaluator struct {
	Params []float64 `json:"params"`
	Type   string    `json:"type"`
}
type Operator struct {
	Type string `json:"type"`
}
type Query struct {
	Params []string `json:"params"`
}
type Reducer struct {
	Params []any  `json:"params"`
	Type   string `json:"type"`
}

type Conditions struct {
	Evaluator Evaluator `json:"evaluator"`
	Operator  Operator  `json:"operator"`
	Query     Query     `json:"query"`
	Reducer   Reducer   `json:"reducer"`
	Type      string    `json:"type"`
}

type Options struct {
	AlertThreshold bool `json:"alertThreshold"`
}
type Targets struct {
	Exemplar     bool   `json:"exemplar"`
	Expr         string `json:"expr"`
	Interval     string `json:"interval"`
	LegendFormat string `json:"legendFormat"`
	QueryType    string `json:"queryType,omitempty"`
	RefID        string `json:"refId"`
}
type Thresholds struct {
	ColorMode string  `json:"colorMode"`
	Fill      bool    `json:"fill"`
	Line      bool    `json:"line"`
	Op        string  `json:"op"`
	Value     float64 `json:"value"`
	Visible   bool    `json:"visible"`
}
type Tooltip struct {
	Shared    bool   `json:"shared"`
	Sort      int    `json:"sort"`
	ValueType string `json:"value_type"`
}
type Xaxis struct {
	Buckets any    `json:"buckets"`
	Mode    string `json:"mode"`
	Name    any    `json:"name"`
	Show    bool   `json:"show"`
	Values  []any  `json:"values"`
}
type Yaxes struct {
	//HashKey string `json:"$$hashKey"` // ngAugular properites, can disable, usually is only for js cache use it.
	Format  string `json:"format"`
	Label   any    `json:"label"`
	LogBase int    `json:"logBase"`
	Max     any    `json:"max"`
	Min     any    `json:"min"`
	Show    bool   `json:"show"`
}
type Yaxis struct {
	Align      bool `json:"align"`
	AlignLevel any  `json:"alignLevel"`
}

type Legend struct {
	AlignAsTable bool `json:"alignAsTable"`
	Avg          bool `json:"avg"`
	Current      bool `json:"current"`
	Max          bool `json:"max"`
	Min          bool `json:"min"`
	RightSide    bool `json:"rightSide"`
	Show         bool `json:"show,omitempty"`
	Total        bool `json:"total"`
	Values       bool `json:"values"`
}
