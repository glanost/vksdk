package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api"
)

func TestVK_StatusGet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StatusGet(api.Params{})
	noError(t, err)
}

func TestVK_StatusSet(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.StatusSet(api.Params{
		"text": "Hello world",
	})
	noError(t, err)
}
