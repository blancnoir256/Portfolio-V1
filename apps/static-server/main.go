package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lpernett/godotenv"
)

var (
	appDir    = ""
	secretDir = ""
	isDev     = false
)

func main() {
	commandLineArguments()
	envLoadAndSet()

	e := echo.New()
	setupMiddleware(e)
	setupRoutes(e)
	startServer(e)
}

// ミドルウェアを設定する関数
func setupMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

// ルーティングを設定する関数
// pingと静的ファイルのルーティングとSPAのフォールバックを設定
func setupRoutes(e *echo.Echo) {
	setupAPIRoutes(e)
	setupStaticRoutes(e)

	e.GET("/*", spaFallbackHandler)
}

// Echoサーバーを起動する関数
// 開発環境ではHTTPSを使用し、その他の環境ではHTTPを使用
// 本番環境ではk8sのIngressを使用してHTTPSを提供することが想定されている
func startServer(e *echo.Echo) {
	if isDev {
		e.Logger.Fatal(e.StartTLS(":443", secretDir+"/tls.crt", secretDir+"/tls.key"))
	} else {
		e.Logger.Fatal(e.Start(":8080"))
	}
}

// コマンドライン引数を処理する関数
// `-isDev` フラグを使用して開発環境かどうかを判定
func commandLineArguments() {
	var (
		isDev_ = flag.Bool("isDev", false, "Is this dev environment?")
	)
	flag.Parse()
	isDev = *isDev_
}

// 環境変数を読み込み、設定する関数
// どの環境変数を読むのかは`getEnv`関数で定義されている
func envLoadAndSet() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	getEnv()
}

// 環境変数を設定する関数
func getEnv() {
	appDir = os.Getenv("APP_DIR")
	secretDir = os.Getenv("SECRET_DIR")
}

// ping確認 後でいい感じに実装
func setupAPIRoutes(e *echo.Echo) {
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong!")
	})
}

// 静的ファイルをに対するルーティングを設定
func setupStaticRoutes(e *echo.Echo) {
	e.Static("/assets", filepath.Join(appDir, "assets"))

	publicFiles := map[string]string{
		"/BN256_icon_mini.png": "BN256_icon_mini.png",
		"/BN256_icon.png":      "BN256_icon.png",
		"/BN256.png":           "BN256.png",
		"/favicon.ico":         "favicon.ico",
	}

	for route, filename := range publicFiles {
		e.File(route, filepath.Join(appDir, filename))
	}
}

// SPAルーティングのフォールバック
func spaFallbackHandler(c echo.Context) error {
	return c.File(filepath.Join(appDir, "index.html"))
}
