package biztag

import (
	"net/http"

	"github.com/colinrs/goleaf/internal/logic/biztag"
	"github.com/colinrs/goleaf/internal/svc"
	"github.com/colinrs/goleaf/internal/types"
	"github.com/colinrs/goleaf/pkg/httpy"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateBizTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateBizTagRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpy.ResultCtx(r, w, nil, err)
			return
		}

		l := biztag.NewCreateBizTagLogic(r.Context(), svcCtx)
		resp, err := l.CreateBizTag(&req)
		httpy.ResultCtx(r, w, resp, err)
	}
}
