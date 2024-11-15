package customerror

import "github.com/raymondwongso/gogox/errorx"

func IsErrNotFound(err error) bool {
	goxErr, ok := errorx.Parse(err)
	if !ok {
		return false
	}
	return goxErr.Code == errorx.CodeNotFound
}
