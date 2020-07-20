package timer

import "time"

func GetNowTime() time.Time {
	// 考虑时区问题
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

// 进行时间的推算
func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTimer.Add(duration), nil
}

//
