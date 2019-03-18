package ginproxy

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

type ProxyOptions struct {
	Target      string
	PathRewrite string
}

func HandleProxy(path string, proxyOpthion ProxyOptions) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if strings.Index(ctx.Request.RequestURI, path) == 0 {
			client := &http.Client{}
			requestUrl := strings.Replace(ctx.Request.RequestURI, proxyOpthion.PathRewrite, "", -1)
			url := proxyOpthion.Target+requestUrl
			req, err := http.NewRequest(ctx.Request.Method, url, ctx.Request.Body)
			if err != nil {
				println(err)
				return
			}
			req.Header = ctx.Request.Header
			resp, err := client.Do(req)
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			for key, value := range resp.Header {
				if len(value) == 1 {
					ctx.Writer.Header().Add(key, value[0])
				}
			}
			ctx.Status(resp.StatusCode)
			ctx.Writer.Write(body)

		} else {
			ctx.Next()
		}

	}
}