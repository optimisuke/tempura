package jobcommand

import (
	"context"

	acalldi "github.com/acall-inc/acall-di"
	definition "github.com/hoge"
)

type (
	server struct {
	}
)

// NewServer はコンストラクタです
func NewServer() definition.hogeServiceServer {
	return &server{}
}

func (t *server) rpcName(context context.Context, req *definition.rpcNameRequest) (*definition.rpcNameResponse, error) {
	container := context.Value(shared.ContainerContextKey).(acalldi.Container)
	res := definition.rpcNameResponse{}

	if err := container.Invoke(func(usecase Usecase) error {
		ures, err := usecase.funcName(&funcNameRequest{
            // TODO
		})
        if err != nil {
			return err
		}
        // TODO: res.hoge = ures.hoge
        return nil
	}); err != nil {
		return nil, err
	}
	return &res, nil
}
