package service

import (
	"github.com/stretchr/testify/assert"
	log "github.com/xushikuan/microlog"
	"testing"
)

func TestGetToken(t *testing.T) {
	resp, err := GetToken()
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("AccessToken : %v; ExpiresIn : %v", resp.AccessToken, resp.ExpiresIn)
	assert.Equal(t, resp.ExpiresIn, 7200)
}
