package controller

import (
	"dm/apilogic"
	"dm/service/resp"
	"github.com/zeromicro/go-zero/rest/httpx"
	"golang.org/x/net/context"
	"net/http"
)

func TestSubmit(svcCtx *context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, _, errF := r.FormFile("file")
		if errF != nil {
			httpx.ErrorCtx(r.Context(), w, resp.NewError(resp.Err_ParamsError, "请上传文件"))
			return
		}
		defer file.Close()

		l := apilogic.NewUploadLogic(r.Context(), r)
		res, errU := l.UploadImage()
		if errU != nil {
			httpx.Error(w, errU)
		} else {
			httpx.OkJsonCtx(r.Context(), w, res)
		}
	}
}