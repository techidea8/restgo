package wxapp

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
)

type UserConfigOption func(*miniProgram.UserConfig)

func UseRedisCache(addr, password string) UserConfigOption {
	return func(cfg *miniProgram.UserConfig) {
		cfg.Cache = kernel.NewRedisClient(&kernel.RedisOptions{
			Addr:     addr,
			Password: password,
			DB:       0,
		})
	}
}

var conf WxConf

func InitializeMiniapp(wxconf WxConf) {
	conf = wxconf
}

type AppInfo interface {
	GetAppId() string
	GetSecret() string
}

func TakeMiniProgramApp(appInfo AppInfo) (ret *miniProgram.MiniProgram, err error) {
	appId := appInfo.GetAppId()
	secret := appInfo.GetSecret()
	if app, ok := mapMiniProgramApp.Load(appId); ok {
		return app.(*miniProgram.MiniProgram), nil
	}
	app, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:     appId,  // 小程序appid
		Secret:    secret, // 小程序app secret
		HttpDebug: conf.HttpDebug,
		Log: miniProgram.Log{
			Level: conf.Log.Level,
			File:  conf.Log.File,
		},
		Cache: kernel.NewRedisClient(&kernel.RedisOptions{
			Addr:     conf.CacheConf.Addr,
			Password: conf.CacheConf.Password,
			DB:       conf.CacheConf.DbIndex,
		}),
	})
	mapMiniProgramApp.Store(appId, app)
	return app, err
}
