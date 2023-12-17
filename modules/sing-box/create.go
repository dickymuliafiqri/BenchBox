package singbox

import (
	"context"

	B "github.com/sagernet/sing-box"
	"github.com/sagernet/sing-box/option"
)

func Create(options option.Options) (*B.Box, error) {
	box, err := B.New(B.Options{
		Context: context.Background(),
		Options: options,
	})
	if err != nil {
		return nil, err
	}

	go box.Start()
	return box, nil
}
