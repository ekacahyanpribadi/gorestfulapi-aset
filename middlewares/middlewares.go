package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"restfulapi/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"

	log "github.com/sirupsen/logrus"
)

const charset = "012345678901234567890123456789012345678901234567890123456789"
const charset2 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomNumber(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func randomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset2[rand.Intn(len(charset2))])
	}
	return sb.String()
}

// rubah nama tabel sesuai database start
type Tabler interface {
	TableName() string
}

func (modelLoghit) TableName() string {
	return "log_hit"
}

//rubah nama tabel sesuai database end

// inisiasi model untuk tabel log hit start
type modelLoghit struct {
	Id          string    `json:"id" gorm:"primary_key"`
	Datetime    time.Time `json:"datetime"`
	Method      string    `json:"method"`
	Desc        string    `json:"desc"`
	Json_req    string    `json:"json_req"`
	Status_resp int       `json:"status_resp"`
	Json_resp   string    `json:"json_resp"`
}

//inisiasi model untuk tabel log hit end

type ginBodyLogger struct {
	// get all the methods implementation from the original one
	// override only the Write() method
	gin.ResponseWriter
	body bytes.Buffer
}

func (g *ginBodyLogger) Write(b []byte) (int, error) {
	g.body.Write(b)
	return g.ResponseWriter.Write(b)
}

func RequestLoggingMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ginBodyLogger := &ginBodyLogger{
			body:           bytes.Buffer{},
			ResponseWriter: ctx.Writer,
		}
		ctx.Writer = ginBodyLogger
		var req interface{}
		if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		data, err := json.Marshal(req)
		if err != nil {
			panic(fmt.Errorf("err while marshaling req msg: %v", err))
		}
		ctx.Next()

		statusCode := ctx.Writer.Status()

		logger.WithFields(log.Fields{
			"status":       ctx.Writer.Status(),
			"method":       ctx.Request.Method,
			"path":         ctx.Request.URL.Path,
			"query_params": ctx.Request.URL.Query(),
			"req_body":     string(data),
			"res_body":     ginBodyLogger.body.String(),
		}).Info("request details")

		currentTime := time.Now().UTC().Add(7 * time.Hour)
		const (
			df = "20060102"
		)
		currentTimex := currentTime.Format(df)
		genId := currentTimex + randomNumber(6) + randomString(6)

		//create log to db start
		postLoghit := modelLoghit{
			Id:          genId,
			Datetime:    currentTime,
			Method:      ctx.Request.Method,
			Desc:        ctx.Request.URL.Path,
			Json_req:    string(data),
			Status_resp: statusCode,
			Json_resp:   ginBodyLogger.body.String(),
		}

		models.DB.Create(&postLoghit)
		//create log to db end
	}
}
