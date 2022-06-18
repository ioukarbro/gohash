package datetime

import "github.com/golang-module/carbon/v2"

const timezone = carbon.Shanghai

// NowDate 2006-01-02
func NowDate() string {
	return now().ToDateString()
}

func NowCarbonDate() carbon.Date {
	return carbon.Date{Carbon: carbon.Now()}
}
func NowCarbonDatetime() carbon.DateTime {
	return carbon.DateTime{Carbon: carbon.Now()}
}

func now() carbon.Carbon {
	return carbon.Now(timezone)
}

// NowDatetime 2006-01-02 13:02:03
func NowDatetime() string {
	return now().ToDateTimeString()
}

func NowDateMilliString() string {
	return now().ToDateMilliString()
}

func NowTimestamp() int64 {
	return now().Timestamp()
}
