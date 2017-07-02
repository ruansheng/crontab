package model

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Parse(str_con string) bool {
	fields := validate(str_con)
	if 0 == len(fields) {
		return false
	}

	times := formatTime()

	var arr_cron [][]int

	arr_cron = append(arr_cron, parseCronPart(fields[0], 0, 59))
	arr_cron = append(arr_cron, parseCronPart(fields[1], 0, 23))
	arr_cron = append(arr_cron, parseCronPart(fields[2], 1, 31))
	arr_cron = append(arr_cron, parseCronPart(fields[3], 1, 12))
	arr_cron = append(arr_cron, parseCronPart(fields[4], 0, 6))

	return check(times, arr_cron)
}

func validate(str_con string) []string {
	var fields []string
	str_cons := strings.Fields(str_con)
	if 5 != len(str_cons) {
		return fields
	}

	//正则校验
	var regxs [5]string
	regxs[0] = "(\\*|([0-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|5[0-9])|\\*\\/([0-9]|1[0-9]|2[0-9]|3[0-9]|4[0-9]|5[0-9]))"
	regxs[1] = "(\\*|([0-9]|1[0-9]|2[0-3])|\\*\\/([0-9]|1[0-9]|2[0-3]))"
	regxs[2] = "(\\*|([1-9]|1[0-9]|2[0-9]|3[0-1])|\\*\\/([1-9]|1[0-9]|2[0-9]|3[0-1]))"
	regxs[3] = "(\\*|([1-9]|1[0-2])|\\*\\/([1-9]|1[0-2]))"
	regxs[4] = "(\\*|([0-6])|\\*\\/([0-6]))"

	pass := true
	for index, field := range str_cons {
		match, _ := regexp.MatchString(regxs[index], field)
		if true != match {
			pass = false
			break
		}
	}
	if false == pass {
		return fields
	}
	return str_cons
}

func formatTime() []int {
	times := make([]int, 0, 5)

	minute := time.Now().Minute()
	hour := time.Now().Hour()
	day := time.Now().Day()
	month := time.Now().Month()
	weekday := time.Now().Weekday()

	times = append(times, minute, hour, day, int(month), int(weekday))

	return times
}

func parseCronPart(field string, f_min int, f_max int) []int {
	var cron_arr []int

	//处理"," -- 列表
	if strings.Contains(field, ",") {
		ifields := strings.Split(field, ",")
		for _, ifield := range ifields {
			t := parseCronPart(ifield, f_min, f_max)
			cron_arr = mergeArray(cron_arr, t)
		}
		return cron_arr
	}

	//处理"/" -- 间隔
	tmp := strings.Split(field, "/")
	field = tmp[0]
	step := 1
	if 2 == len(tmp) {
		step, _ = strconv.Atoi(tmp[1])
	}

	//处理"-" -- 范围
	var min int
	var max int
	if strings.Contains(field, "-") {
		min_max := strings.Split(field, "-")
		min, _ = strconv.Atoi(min_max[0])
		max, _ = strconv.Atoi(min_max[1])
	} else if "*" == field {
		min = f_min
		max = f_max
	} else {
		min, _ = strconv.Atoi(field)
		max, _ = strconv.Atoi(field)
	}

	//如果是*，则填充
	if min == f_min && max == f_max && step == 1 {
		return rangeArray(min, max, step)
	}

	//越界判断 数值越界  应该：分0-59，时0-59，日1-31，月1-12，周0-6
	if min < f_min || max > f_max {
		return cron_arr
	}

	if (max - min) > step {
		cron_arr = rangeArray(min, max, step)
	} else {
		cron_arr = append(cron_arr, min)
	}

	return cron_arr
}

func check(times []int, arr_cron [][]int) bool {
	if (5 != len(times)) || (5 != len(arr_cron)) {
		return false
	}

	minute_flag := (0 != len(arr_cron[0])) && inArray(times[0], arr_cron[0])
	hour_flag := (0 != len(arr_cron[1])) && inArray(times[1], arr_cron[1])
	day_flag := (0 != len(arr_cron[2])) && inArray(times[2], arr_cron[2])
	month_flag := (0 != len(arr_cron[3])) && inArray(times[3], arr_cron[3])
	weekday_flag := (0 != len(arr_cron[4])) && inArray(times[4], arr_cron[4])
	if minute_flag && hour_flag && day_flag && month_flag && weekday_flag {
		return true
	}
	return false
}

func mergeArray(a []int, b []int) []int {
	for _, value := range b {
		a = append(a, value)
	}
	return a
}

func rangeArray(min int, max int, step int) []int {
	var range_arr []int
	i := 0
	for i = min; i <= max; i += step {
		range_arr = append(range_arr, i)
	}
	return range_arr
}

func inArray(a int, b []int) bool {
	flag := false
	for _, value := range b {
		if value == a {
			flag = true
		}
	}
	return flag
}
