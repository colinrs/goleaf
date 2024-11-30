package idgen

import (
	"net/http"

	"github.com/colinrs/goleaf/internal/logic/idgen"
	"github.com/colinrs/goleaf/internal/svc"
	"github.com/colinrs/goleaf/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DecodeSnowflakeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DecodeSnowflakeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := idgen.NewDecodeSnowflakeLogic(r.Context(), svcCtx)
		resp, err := l.DecodeSnowflake(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
