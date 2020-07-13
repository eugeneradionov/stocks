package token

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const blackListPrefix = "blacklist_token_id"

func (srv service) RemoveFromBlacklist(tokenID uuid.UUID) error {
	return srv.rds.Del(srv.TokenBlackListKey(tokenID))
}

func (srv service) AddToBlacklist(tokenID uuid.UUID, ttl int64) error {
	if ttl <= 0 {
		ttl = int64(srv.expirationTimeSec)
	}

	return srv.rds.Set(srv.TokenBlackListKey(tokenID), []byte{}, time.Duration(ttl)*time.Second)
}

func (srv service) TokenBlackListKey(tokenID uuid.UUID) string { // nolint:interfacer
	return fmt.Sprintf("%s:%s", blackListPrefix, tokenID.String())
}
