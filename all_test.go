package ipinfo

import (
	"log"
	"testing"
)

func TestIpInfo(t *testing.T) {
	info, err := GetIpInfo("")
	log.Println(info, err)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(info)
	}
}
