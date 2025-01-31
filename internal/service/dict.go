package service

import (
	"github.com/gin-gonic/gin"

	"github.com/EZ4BRUCE/go-grpc-layout/internal/consts"
	"github.com/EZ4BRUCE/go-grpc-layout/internal/ecode"
	"github.com/EZ4BRUCE/go-grpc-layout/pkg/consul"
	"github.com/EZ4BRUCE/go-grpc-layout/pkg/utils/response"
)

func (s *HttpService) GetDictList(c *gin.Context) {
	cli, err := consul.NewConsulClient(consts.Conf.Consul.Host, consts.Conf.Consul.Token)
	if err != nil {
		response.Fail(c, ecode.GetDictListFailed, err)
		return
	}

	dict := make(map[string]interface{}, 0)
	_, err = consul.GetConsulKV(cli, consts.Conf.Consul.Kv.DictPath, &dict)
	if err != nil {
		response.Fail(c, ecode.GetDictListFailed, err)
		return
	}
	res := map[string]interface{}{
		"filename": "dict.json",
		"dictInfo": dict,
	}
	response.Success(c, res)
}
