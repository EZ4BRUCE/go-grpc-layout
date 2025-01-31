package service

import (
	"net/http"

	v1 "github.com/EZ4BRUCE/go-grpc-layout/api/http/v1"
	"github.com/EZ4BRUCE/go-grpc-layout/internal/ecode"
	"github.com/EZ4BRUCE/go-grpc-layout/pkg/utils/request"
	"github.com/EZ4BRUCE/go-grpc-layout/pkg/utils/response"
	"github.com/gin-gonic/gin"
)

func (s *HttpService) DebugPerf(c *gin.Context) {
	req := &v1.DebugPerfRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, ecode.ParamsFailed, err)
		return
	}
	res, err := s.uc.DebugPerf(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	response.Success(c, res)
}
