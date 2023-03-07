package utils

import (
	"time"

	log "github.com/techidea8/restgo/pkg/log"
)

// @brief：耗时统计函数
// defer timeCost()()
func TimeCost(tags string, params ...interface{}) func() {
	start := time.Now()
	return func() {
		tc := time.Since(start)
		if len(params) > 0 {
			log.Debugf("%s,%dms,%dns", tags, tc/(1000*1000), tc, params)
		} else {
			log.Debugf("%s,%dms,%dns", tags, tc/(1000*1000), tc)
		}
	}
}
