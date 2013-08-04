package main

import (
	"crypto/md5"
	"errors"
	"io"
	"strconv"
	"strings"
)

func getInfo(uid string) (string, int64) {
	// TODO query host and timeout from the database
	host := ""
	timeout := int64(0)
	return host, timeout
}

func setInfo(token string, host string, timeout string) error {
	timeoutNum, err := strconv.Atoi(timeout)
	if err != nil {
		return err
	}
	if timeoutNum <= 0 {
		return errors.New("negative or zero value of timeout")
	}
	if len(strings.Split(host, ":")) != 2 {
		return errors.New("host format error")
	}
	uid := getUID(token)
	if uid == "" {
		return errors.New("token error")
	}
	// TODO store the mapping from uid to host and timeout to the database
	return nil
}

// get user
func getUID(token string) string {
	// TODO query the database to get uid of a token. If the token doesn't
	// exist, return ""
	return ""
}

// TODO we use md5 to generate token for new registered user. This has to be
// revised for security reasons.
func generateToken(uid string) string {
	h := md5.New()
	io.WriteString(h, uid)
	token := string(h.Sum(nil))
	// TODO write uid and token to data base
	return token
}
