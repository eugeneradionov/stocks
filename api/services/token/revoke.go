package token

func (srv service) Revoke(accessToken string) (*Claims, error) {
	token, err := srv.parseJWT(accessToken)
	if err != nil {
		return nil, err
	}

	claims := srv.parseClaims(token)
	if claims == nil {
		return nil, ErrTokenClaimsInvalid
	}

	err = srv.AddToBlacklist(claims.TokenID, claims.TTL())
	if err != nil {
		return nil, err
	}

	return claims, nil
}
