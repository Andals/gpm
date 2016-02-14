/**
* @file misc.go
* @brief misc supermarket
* @author ligang
* @date 2015-12-11
 */

package misc

import ()

const (
	TIME_FMT_STR_YEAR   = "2006"
	TIME_FMT_STR_MONTH  = "01"
	TIME_FMT_STR_DAY    = "02"
	TIME_FMT_STR_HOUR   = "15"
	TIME_FMT_STR_MINUTE = "04"
	TIME_FMT_STR_SECOND = "05"
)

func TimeGeneralLayout() string {
	layout := TIME_FMT_STR_YEAR + "-" + TIME_FMT_STR_MONTH + "-" + TIME_FMT_STR_DAY + " "
	layout += TIME_FMT_STR_HOUR + ":" + TIME_FMT_STR_MINUTE + ":" + TIME_FMT_STR_SECOND

	return layout
}