package biztag

import (
	"github.com/colinrs/goleaf/pkg/httpy"
	"net/http"

	"github.com/colinrs/goleaf/internal/logic/biztag"
	"github.com/colinrs/goleaf/internal/svc"
	"github.com/colinrs/goleaf/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateBizTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateBizTageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpy.ResultCtx(r, w, nil, err)
		}

		l := biztag.NewUpdateBizTagLogic(r.Context(), svcCtx)
		resp, err := l.UpdateBizTag(&req)
		httpy.ResultCtx(r, w, resp, err)
	}
}
