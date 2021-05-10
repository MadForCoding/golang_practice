package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"golang_practice/AIRPAY-64062/model"
	"io"
	"os"
	"strings"
)

const (
	traceSplit = "traceId:"
	sep        = "|"
)

var (
	csvHeader = []string{"Trace Id", "Top Merchant Id", "Merchant Id", "Outlet Id", "Reason"}
)

func main() {
	reasonMap, err := ReadReasonMsg("/Users/weichen/Downloads/04/04_27/auth_func_04_26_vn.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	merchantInfoMap, err := ReadMerchantInfo("/Users/weichen/Downloads/04/04_27/auth_43_04_26_vn.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(len(reasonMap), len(merchantInfoMap))
	resultMap := combine(reasonMap, merchantInfoMap)
	fmt.Println("RESULTMAP", len(resultMap))

	// 去重
	res := handleDuplicate(resultMap)
	fmt.Println("AFTER", len(res))
	// 生成csv文件
	fp, err := os.OpenFile("/Users/weichen/Downloads/04/04_27/auth_04_26_vn.csv", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	csvWrite := csv.NewWriter(fp)
	csvWrite.Write(csvHeader)
	for _, value := range res {
		tmp := []string{value.TraceId, value.TopMerchantId, value.MerchantId, value.OutletId, value.Reason}
		fmt.Println(tmp)
		csvWrite.Write(tmp)
	}
	csvWrite.Flush()

}

func ReadReasonMsg(filePath string) (map[string]*model.ReasonMsg, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reasonMap := make(map[string]*model.ReasonMsg)
	br := bufio.NewReader(file)
	for {
		content, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		// 字符串处理
		arr := strings.Split(string(content), traceSplit)
		fmt.Println(arr)
		finalSlice := strings.Split(arr[1], sep)
		traceId := finalSlice[0]
		msg := finalSlice[len(finalSlice)-1]
		reason := &model.ReasonMsg{
			TraceId: strings.TrimSpace(traceId),
			Reason:  strings.TrimSpace(msg),
		}
		reasonMap[reason.TraceId] = reason
	}

	return reasonMap, nil
}

func ReadMerchantInfo(filePath string) (map[string]*model.MerchantInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	infoMap := make(map[string]*model.MerchantInfo)
	br := bufio.NewReader(file)
	for {
		content, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		//字符串处理
		arr := strings.Split(string(content), traceSplit)
		if len(arr) < 2 {
			continue
		}
		tmpSlice := strings.Split(arr[1], sep)
		traceId := tmpSlice[0]
		merchantSlice := strings.Split(tmpSlice[len(tmpSlice)-1], ",")
		merchantInfo := &model.MerchantInfo{
			TraceId:       strings.TrimSpace(traceId),
			TopMerchantId: strings.TrimSpace(merchantSlice[1]),
			MerchantId:    strings.TrimSpace(merchantSlice[2]),
			OutletId:      strings.TrimSpace(merchantSlice[3]),
		}
		infoMap[merchantInfo.TraceId] = merchantInfo
	}

	return infoMap, nil
}

func combine(reasonMap map[string]*model.ReasonMsg, merchantInfoMap map[string]*model.MerchantInfo) map[string]*model.Result {
	res := make(map[string]*model.Result)
	for reasonTrace, reason := range reasonMap {
		if info, ok := merchantInfoMap[reasonTrace]; ok {
			if _, ok := res[reasonTrace]; !ok {
				tmp := &model.Result{
					TraceId:       reasonTrace,
					TopMerchantId: info.TopMerchantId,
					MerchantId:    info.MerchantId,
					OutletId:      info.OutletId,
					Reason:        reason.Reason,
				}
				res[tmp.TraceId] = tmp
			}
		}
	}
	return res
}

func handleDuplicate(midMap map[string]*model.Result) map[string]*model.Result {
	mark := make(map[string]int)
	res := make(map[string]*model.Result)
	for _, value := range midMap {
		tSlice := strings.Split(value.TopMerchantId, ":")
		mSlice := strings.Split(value.MerchantId, ":")
		oSlice := strings.Split(value.OutletId, ":")
		msg := strings.TrimSpace(value.Reason)
		tmp := fmt.Sprintf("%s%s%s%s", strings.TrimSpace(tSlice[1]), strings.TrimSpace(mSlice[1]), strings.TrimSpace(oSlice[1]), msg)
		fmt.Println(tmp)
		i := mark[tmp]
		if i == 0 {
			t := &model.Result{
				TraceId:       value.TraceId,
				TopMerchantId: tSlice[1],
				MerchantId:    mSlice[1],
				OutletId:      oSlice[1],
				Reason:        value.Reason,
			}
			res[t.TraceId] = t
			mark[tmp] = 1
		}
	}
	return res
}
