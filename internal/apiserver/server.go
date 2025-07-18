package apiserver

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/onexstack/fastgo/internal/apiserver/biz"
	"github.com/onexstack/fastgo/internal/apiserver/handler"
	"github.com/onexstack/fastgo/internal/apiserver/pkg/conversion/validation"
	"github.com/onexstack/fastgo/internal/apiserver/store"
	"github.com/onexstack/fastgo/internal/pkg/core"
	"github.com/onexstack/fastgo/internal/pkg/errorsx"
	mw "github.com/onexstack/fastgo/internal/pkg/middleware"
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
		mw.NoCache,
		mw.Cors,
		mw.RequestID(),
	}
	engine.Use(mws...)

	// 初始化数据库连接
	db, err := cfg.MySQLOptions.NewDB()
	if err != nil {
		return nil, err
	}
	store := store.NewStore(db)
	cfg.InstallRESTAPI(engine, store)

	srv := &http.Server{
		Addr:    cfg.Addr,
		Handler: engine,
	}

	return &Server{
		cfg: cfg,
		srv: srv,
	}, nil
}

// 注册 API 路由。路由的路径和 HTTP 方法，严格遵循 REST 规范.
func (cfg *Config) InstallRESTAPI(engine *gin.Engine, store store.IStore) {
	// 注册 404 Handler.
	engine.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, nil, errorsx.ErrNotFound.WithMessage("Page not found"))
	})

	// 注册 /healthz handler.
	engine.GET("/healthz", func(c *gin.Context) {
		core.WriteResponse(c, map[string]string{"status": "ok"}, nil)
	})

	// 创建核心业务处理器
	handler := handler.NewHandler(biz.NewBiz(store), validation.NewValidator(store))
	authMiddlewares := []gin.HandlerFunc{mw.Authn()}

	// 注册用户登录和令牌刷新接口。这2个接口比较简单，所以没有 API 版本
	engine.POST("/login", handler.Login)
	engine.POST("/refresh-token", mw.Authn(), handler.RefreshToken)

	// 注册 v1 版本 API 路由分组
	v1 := engine.Group("/v1")
	{
		// 用户相关路由
		userv1 := v1.Group("/users")
		{
			// 创建用户。这里要注意：创建用户是不用进行认证和授权的
			userv1.POST("", handler.CreateUser)
			userv1.Use(authMiddlewares...)

			userv1.PUT(":userID", handler.UpdateUser)    // 更新用户信息
			userv1.DELETE(":userID", handler.DeleteUser) // 删除用户
			userv1.GET(":userID", handler.GetUser)       // 查询用户详情
			userv1.GET("", handler.ListUser)             // 查询用户列表.
			userv1.PUT(":userID/change-password", handler.ChangePassword)
		}

		// 博客相关路由
		postv1 := v1.Group("/posts", authMiddlewares...)
		{
			postv1.POST("", handler.CreatePost)       // 创建博客
			postv1.PUT(":postID", handler.UpdatePost) // 更新博客
			postv1.DELETE("", handler.DeletePost)     // 删除博客
			postv1.GET(":postID", handler.GetPost)    // 查询博客详情
			postv1.GET("", handler.ListPost)          // 查询博客列表
		}
	}
}

func (s *Server) Run() error {
	slog.Info("Read MySQL host from config", "mysql.addr", s.cfg.MySQLOptions.Addr)

	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Failed to start server", "error", err)
		}
	}()

	// 创建一个 os.Signal 类型的 channel，用于接收系统信号
	quit := make(chan os.Signal, 1)

	// 监听系统信号，如 SIGINT 和 SIGTERM
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	slog.Info("Shutting down server...")

	// 优雅关闭服务
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", "error", err)
		return err
	}

	slog.Info("Server exited")

	return nil
}
