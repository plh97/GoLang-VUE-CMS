package handler

import (
	v1 "go-nunu/api/v1"
	"go-nunu/internal/service"
	"go-nunu/pkg/aws" // 假设 aws 包被正确导入
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommonHandler struct {
	*Handler
	CommonService service.CommonService
	// R2Client      *aws.CloudflareR2 // 1. 结构体新增 R2 客户端字段
}

func NewCommonHandler(handler *Handler, commonService service.CommonService, r2Client *aws.CloudflareR2) *CommonHandler {
	return &CommonHandler{
		Handler:       handler,
		CommonService: commonService,
		// R2Client:      r2Client, // 2. 构造函数接收并赋值 R2 客户端
	}
}

func (h *CommonHandler) UploadPresignedUrl(ctx *gin.Context) {
	var req v1.UploadPresignedUrlRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, v1.ErrBadRequest, nil)
		return
	}

	preSignedUrl, endpointUrl, err := h.CommonService.UploadPresignedUrl(req.FileExt, req.UploadScene)
	// preSignedUrl, endpointUrl, err := h.R2Client.UploadPresignedUrl(req.FileExt, req.UploadScene)

	if err != nil {
		v1.HandleError(ctx, http.StatusOK, v1.ErrBadRequest, nil)
		return
	}

	v1.HandleSuccess(ctx, v1.UploadPresignedUrlResponseData{
		PreSignedUrl: preSignedUrl,
		EndpointUrl:  endpointUrl,
	})
}
