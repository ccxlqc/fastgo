package apiserver

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	ms "github.com/onexstack/fastgo/internal/pkg/middleware"
	genericoptions "github.com/onexstack/fastgo/pkg/options"
)

type Config struct {
	MySQLOptions *genericoptions.MySQLOptions
	Addr         string
}

type Server struct {
	cfg *Config
	srv *http.Server
}

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		slog.Info("Request received", "method", c.Request.Method, "path", c.Request.URL.Path)
		c.Next()
	}
}

func (cfg *Config) NewServer() (*Server, error) {
	engine := gin.New()

	// gin.Recovery() 中间件，用来捕获任何 panic，并恢复
	mws := []gin.HandlerFunc{
		gin.Recovery(),
		ms.NoCache,
		ms.Cors,
		ms.RequestID(),
	}
	engine.Use(mws...)

	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PageNotFound", "message": "Page not found."})
	})

	engine.Handle("GET", "/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	srv := &http.Server{
		Addr:    cfg.Addr,
		Handler: engine,
	}

	return &Server{
		cfg: cfg,
		srv: srv,
	}, nil
}

func (s *Server) Run() error {
	slog.Info("Read MySQL host from config", "mysql.addr", s.cfg.MySQLOptions.Addr)

	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
