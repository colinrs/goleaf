package biztag

import (
	"net/http"

	"github.com/colinrs/goleaf/internal/logic/biztag"
	"github.com/colinrs/goleaf/internal/svc"
	"github.com/colinrs/goleaf/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetBizTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetBizTageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := biztag.NewGetBizTagLogic(r.Context(), svcCtx)
		resp, err := l.GetBizTag(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}