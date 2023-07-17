package util

import "strconv"

type LocaleStrcut struct {
	Th string `json:"th"`
	En string `json:"en"`
}

func MapOrderStatusEmail(stateMachineStatus string) LocaleStrcut {

	// watingToAllotStatus := []string{"SETTLED", "APPROVED", "WAITING"}
	// pendingStatus := []string{"CREATED", "CONFIRMED", "SUBMITTED"}
	// succeedStatus := []string{"ALLOTTED"}
	// failedStatus := []string{"CANCELLED", "ERROR", "REJECTED", "SUSPENDING"}
	mapStatus := map[string]string{
		"SETTLED": "watingToAllotStatus", "APPROVED": "watingToAllotStatus", "WAITING": "watingToAllotStatus",
		"CREATED": "pendingStatus", "CONFIRMED": "pendingStatus", "SUBMITTED": "pendingStatus",
		"ALLOTTED":  "succeedStatus",
		"CANCELLED": "failedStatus", "ERROR": "failedStatus", "REJECTED": "failedStatus", "SUSPENDING": "failedStatus",
	}
	locateStatus := map[string]LocaleStrcut{
		"watingToAllotStatus": {
			Th: "รอจัดสรร",
			En: "Waiting for allocation",
		},
		"pendingStatus": {
			Th: "รอดำเนินการ",
			En: "Waiting to be processed",
		},
		"succeedStatus": {
			Th: "สำเร็จ",
			En: "Completed",
		},
		"failedStatus": {
			Th: "ส่งคำสั่งไม่สำเร็จ",
			En: "Unsuccessful order",
		},
	}

	return locateStatus[mapStatus[stateMachineStatus]]
}

type MapSchedulerDcaEmailInput struct {
	DcaType    string `json:"dcaType"`
	DayOfMonth int    `json:"dayOfMonth"`
	DayOfWeek  string `json:"dayOfWeek"`
}

func MapSchedulerDcaEmail(input MapSchedulerDcaEmailInput) LocaleStrcut {
	dayEn := ""
	dayTh := ""
	switch input.DcaType {
	case "M":
		if input.DayOfMonth == 99 {
			dayEn = "end of month"
			dayTh = "สิ้นเดือน"
		} else {
			dayEn = OrdinalSuffixOfDay(input.DayOfMonth)
			dayTh = "วันที่ " + strconv.Itoa(input.DayOfMonth)
		}
	case "W":
		fullDateName := ShortDayMapToFullDay[input.DayOfWeek]
		dayEn = fullDateName.En
		dayTh = fullDateName.Th
	}

	schedulerStatus := map[string]LocaleStrcut{
		"M": {
			Th: "รายเดือน (ทุก" + dayTh + ")",
			En: "Monthly (Every " + dayEn + ")",
		},
		"W": {
			Th: "รายสัปดาห์ (ทุก " + dayTh + ")",
			En: "Weekly (Every " + dayEn + ")",
		},
	}

	return schedulerStatus[input.DcaType]
}
