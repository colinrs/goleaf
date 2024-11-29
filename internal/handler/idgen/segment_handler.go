package idgen

import (
	"net/http"

	"github.com/colinrs/goleaf/internal/logic/idgen"
	"github.com/colinrs/goleaf/internal/svc"
	"github.com/colinrs/goleaf/internal/types"
	"github.com/colinrs/goleaf/pkg/httpy"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SegmentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SegmentRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpy.ResultCtx(r, w, nil, err)
			return
		}
		l := idgen.NewSegmentLogic(r.Context(), svcCtx)
		resp, err := l.Segment(&req)
		httpy.ResultCtx(r, w, resp, err)
	}
}
