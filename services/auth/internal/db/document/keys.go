package db

import (
	"github.com/gstones/platform/services/common/nosql/document"
)

func userAuthenticationKey(appId, userId string) (document.Key, error) {
	return document.NewKeyFromParts(appId, "users", userId)
}

func userPlatformAuthenticationKey(appId, platformType, platformUid string) (document.Key, error) {
	return document.NewKeyFromParts(appId, platformType, "users", platformUid)
}
