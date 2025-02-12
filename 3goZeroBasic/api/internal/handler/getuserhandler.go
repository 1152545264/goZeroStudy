package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goZeroStudy/3goZeroBasic/api/internal/logic"
	"goZeroStudy/3goZeroBasic/api/internal/svc"
	"goZeroStudy/3goZeroBasic/api/internal/types"
)

func getUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetUserLogic(r.Context(), svcCtx)
		resp, err := l.GetUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
