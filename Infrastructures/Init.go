package Infrastructures

import (
	"context"
	"fmt"
	"log"
	"mygofarm/Infrastructures/Redis"
	"mygofarm/Infrastructures/Rom"
	"mygofarm/Pkg/Logger"
	"mygofarm/Routes"
	"mygofarm/Settings"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func Init() (err error) {
	/////////////////////////////////////////////////////////////
	////////////1.加载配置////////////////////////////////////////
	////////////////////////////////////////////////////////////
	if err := Settings.Init(); err != nil {
		fmt.Printf("init setting failed, err:%v\n", err)
	}
	zap.L().Info("loading init success...")
	/////////////////////////////////////////////////////////////
	////////////2.初始化日志//////////////////////////////////////
	////////////////////////////////////////////////////////////
	if err := Logger.Init(Settings.Conf.LogConfig, Settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
	}
	zap.L().Sync() //延迟注册
	zap.L().Info("loading logger success...")
	/////////////////////////////////////////////////////////////
	////////////3.初始化MYSQL连接/////////////////////////////////
	////////////////////////////////////////////////////////////
	if err := Rom.Init(Settings.Conf.MysqlConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
	}
	zap.L().Info("loading mysqlconfig success...")
	/////////////////////////////////////////////////////////////
	////////////4.初始化Redis连接/////////////////////////////////
	////////////////////////////////////////////////////////////
	if err := Redis.Init(Settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
	}
	zap.L().Info("loading redisconfig success...")
	defer Redis.Close() //释放redis连接

	/////////////////////////////////////////////////////////////
	////////////7.初始化kafka生产者连接/////////////////////////////////
	////////////////////////////////////////////////////////////
	//if err := kafka.InitKafkaProducer(settings.Conf.KafkaConfig); err != nil {
	//	fmt.Printf("init kafka failed, err:%v\n", err)
	//}
	//zap.L().Info("loading kafka Producer success...")
	/////////////////////////////////////////////////////////////
	////////////7.初始化kafka消费者连接/////////////////////////////////
	////////////////////////////////////////////////////////////
	//if err := kafka.InitKafkaConsumer(settings.Conf.KafkaConfig); err != nil {
	//	fmt.Printf("init kafka failed, err:%v\n", err)
	//}
	//zap.L().Info("loading kafka Consumer success...")
	/////////////////////////////////////////////////////////////
	////////////8.注册路由////////////////////////////////////////
	/////////////////////////////////////////////////////////////
	r := Routes.Setup()
	/////////////////////////////////////////////////////////////
	////////////8.启动服务（优雅关机）///////////////////////////////
	/////////////////////////////////////////////////////////////
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", Settings.Conf.Port),
		Handler: r,
	}
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("Server Shutdown: ", zap.Error(err))
	}
	zap.L().Info("Server exiting ...")
	return

}
