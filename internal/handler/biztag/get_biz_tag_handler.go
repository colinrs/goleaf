package biztag

import (
	"github.com/colinrs/goleaf/pkg/httpy"
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
			httpy.ResultCtx(r, w, nil, err)
			return
		}

		l := biztag.NewGetBizTagLogic(r.Context(), svcCtx)
		resp, err := l.GetBizTag(&req)
		httpy.ResultCtx(r, w, resp, err)
	}
}
