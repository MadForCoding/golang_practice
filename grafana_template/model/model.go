package model

import (
	"encoding/json"
	"math/rand"
)

type AutoGenerated struct {
	Annotations   Annotations         `json:"annotations"`
	Editable      bool                `json:"editable"`
	GnetID        any                 `json:"gnetId"`
	GraphTooltip  int                 `json:"graphTooltip"`
	ID            int                 `json:"id"`
	Links         []any               `json:"links"`
	Panels        []*Panels           `json:"panels"`
	Refresh       string              `json:"refresh,omitempty"`
	SchemaVersion int                 `json:"schemaVersion"`
	Style         string              `json:"style"`
	Tags          []any               `json:"tags"`
	Templating    Templating          `json:"templating"`
	Time          Time                `json:"time"`
	Timepicker    Timepicker          `json:"timepicker"`
	Timezone      string              `json:"timezone"`
	Title         string              `json:"title"`
	UID           string              `json:"uid"`
	Version       int                 `json:"version"`
	panelMap      map[string]struct{} `json:"-"`
}

func InitDashBoardByJson(dashBoardJson []byte) *AutoGenerated {
	if len(dashBoardJson) == 0 {
		panic("empty dashBoard json string")
	}
	var dashBoard = &AutoGenerated{}
	if err := json.Unmarshal(dashBoardJson, dashBoard); err != nil {
		panic(err)
	}

	dashBoard.panelMap = make(map[string]struct{})
	for _, p := range dashBoard.Panels {
		dashBoard.panelMap[p.Title] = struct{}{}
	}
	if len(dashBoard.Panels) == 0 {
		dashBoard.Panels = make([]*Panels, 0, 10)
	}

	return dashBoard
}

func (r *AutoGenerated) PreCheckPanelsIDWithFillUp() {
	if r == nil {
		panic("empty object")
	}
	if len(r.Panels) == 0 {
		return
	}
	dupMap := make(map[int]struct{})
	for _, v := range r.Panels {
		if v.ID != 0 {
			_, has := dupMap[v.ID]
			if !has {
				dupMap[v.ID] = struct{}{}
				continue
			}
		}

		// 重复的，还有没有id的， 都会被重新生成
		v.ID = r.generateID()
		for {
			_, has := dupMap[v.ID]
			if !has {
				dupMap[v.ID] = struct{}{}
				break
			}
		}

	}
}

func (r *AutoGenerated) panelExist(panelTitle string) bool {
	_, has := r.panelMap[panelTitle]
	return has
}

func (r *AutoGenerated) generateID() int {
	casualNumber := rand.Intn(100)
	return 100 + casualNumber
}

func (r *AutoGenerated) PanelAppend(panel *Panels, existForUpdate bool) {
	if r == nil {
		panic("dashboard not exist as cannot adding panel")
	}
	if !r.panelExist(panel.Title) {
		r.Panels = append(r.Panels, panel)
		r.panelMap[r.Title] = struct{}{}
		return
	}

	if existForUpdate {
		for i, p := range r.Panels {
			if p.Title == panel.Title {
				r.Panels[i] = panel
				return
			}
		}
	}
	return

}

type Annotations struct {
	List []OutSideList `json:"list"`
}

type OutSideList struct {
	BuiltIn    int    `json:"builtIn"`
	Datasource string `json:"datasource"`
	Enable     bool   `json:"enable"`
	Hide       bool   `json:"hide"`
	IconColor  string `json:"iconColor"`
	Name       string `json:"name"`
	Type       string `json:"type"`
}

type Templating struct {
	List []any `json:"list"`
}

type Time struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Timepicker struct {
}
