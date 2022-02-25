package activity

import "time"

type ActivityInfoDTO struct {
	ActivityName      string    `json:"activity_name"`
	ActivityStatus    uint      `json:"activity_status"`
	LimitBuy          uint      `json:"limit_buy"`
	ActivityStartTime time.Time `json:"activity_start_time"`
	ActivityEndTime   time.Time `json:"activity_end_time"`
}
