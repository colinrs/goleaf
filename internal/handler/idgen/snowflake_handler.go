package idgen

import (
	"net/http"

	"github.com/colinrs/goleaf/internal/logic/idgen"
	"github.com/colinrs/goleaf/internal/svc"
	"github.com/colinrs/goleaf/internal/types"
	"github.com/colinrs/goleaf/pkg/httpy"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SnowflakeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SnowflakeRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpy.ResultCtx(r, w, nil, err)
		}

		l := idgen.NewSnowflakeLogic(r.Context(), svcCtx)
		resp, err := l.Snowflake(&req)
		httpy.ResultCtx(r, w, resp, err)
	}
}
