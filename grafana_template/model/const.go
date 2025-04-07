package model

type AlarmNoDataState string

const (
	AlarmNoDataState_Ok        AlarmNoDataState = "ok"
	AlarmNoDataState_KeepState AlarmNoDataState = "keep_state"
	AlarmNoDataState_Alarming  AlarmNoDataState = "alerting"
	AlarmNoDataState_NoData    AlarmNoDataState = "no_data"
)

func (r AlarmNoDataState) Raw() string {
	return string(r)
}

type AlarmExecErrState string

const (
	AlarmExecErrState_KeepState AlarmExecErrState = "keep_state"
	AlarmExecErrState_Alarming  AlarmExecErrState = "alerting"
)

func (r AlarmExecErrState) Raw() string {
	return string(r)
}
