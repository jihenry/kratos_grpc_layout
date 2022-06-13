package cache

import (
	"strconv"
	"strings"
	"time"
)

//多字段
func KeyOfDemoMultiKey(uniq string, sid uint64) string {
	return strings.Join([]string{"user:si", strconv.FormatUint(sid, 10), uniq}, ":")
}

//单字段
func KeyOfDemo(id int32) string {
	return "qrcode:param:" + strconv.FormatInt(int64(id), 10)
}

//常量形式的redis key，或者单参数前缀

//缓存过期时间常量
const (
	DemoConst = 24 * time.Hour //用户模块通用的过期时间
)
