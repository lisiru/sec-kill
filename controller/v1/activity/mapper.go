package activity

import "sec-kill/model"

func convertActivityInfoDTO(activity *model.Activity) *ActivityInfoDTO {
	return &ActivityInfoDTO{
		ActivityName:      activity.ActivityName,
		ActivityStartTime: activity.StartTime,
		ActivityEndTime:   activity.EndTime,
		ActivityStatus:    activity.Status,
		LimitBuy:          activity.LimitBuy,
	}

}
