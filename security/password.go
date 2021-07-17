package security

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePwd(pwd string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd)) == nil
}

func HashPwd(pwd string) (ret string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return ret, err
	}
	return string(hash), nil
}
