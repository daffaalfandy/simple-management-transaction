package config

import (
	"os"
	"strconv"
)

type Authenticator struct {
	Secret string
	Expire int
}

func initAuthenticator() Authenticator {
	expireTime, err := strconv.Atoi(os.Getenv("JWT_EXPIRED_TOKEN"))
	if err != nil {
		panic(err)
	}

	return Authenticator{
		Secret: os.Getenv("SECRET_KEY"),
		Expire: expireTime,
	}
}

func (r *Authenticator) GetSecret() string {
	return r.Secret
}

func (r *Authenticator) GetExpire() int {
	return r.Expire
}
