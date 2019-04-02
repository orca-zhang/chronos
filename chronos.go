package chronos

func toUnix(year, month, day, hour, min, sec int) int64 {
	if month < 1 || month > 12 {
		return 0
	}

	leap := 0
	if (year%4 == 0 && (year%100 != 0 || year%400 == 0)) && month >= 3 {
		leap = 1 // February 29
	}

	return int64((365*year-719528+day-1+(year+3)/4-(year+99)/100+(year+399)/400+int([]int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}[month-1])+leap)*86400 + hour*3600 + min*60 + sec)
}
