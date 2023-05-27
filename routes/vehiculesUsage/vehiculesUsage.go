package vehiculesUsage

import "time"

type VehiculesUsageDb struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	VehiculesId int       `json:"vehicules_id"`
	Usage       string    `json:"usage"`
	Time        time.Time `json:"time"`
}
