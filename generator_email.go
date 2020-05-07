package chinaid

import (
	"github.com/mritd/chinaid/metadata"
)

// Email 返回随机邮箱，邮箱目前只支持常见的域名后缀
func Email() string {
	return randStr(8) + "@" + randStr(5) + metadata.DomainSuffix[randInt(0, len(metadata.DomainSuffix))]
}
