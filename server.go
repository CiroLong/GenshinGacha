package main

import (
	"Genshin/gacha"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := gin.Default()

	r.GET("/gacha", gachaOnceRes)
	r.GET("/gachaten", gachaTenRes)

	r.StaticFile("/test", "./source/test.html")

	{
		r.StaticFile("/favicon.ico", "./img/5star/character/优菈.jpg")
		r.StaticFile("/img/5star1", "./img/5star/character/刻晴.jpg")
		r.StaticFile("/img/5star2", "./img/5star/character/迪卢克.jpg")
		r.StaticFile("/img/5star3", "./img/5star/character/琴.jpg")
		r.StaticFile("/img/5star4", "./img/5star/character/莫娜.jpg")
		r.StaticFile("/img/5star5", "./img/5star/character/七七.jpg")
		r.StaticFile("/img/5star6", "./img/5star/character/温迪.jpg")
		r.StaticFile("/img/5star7", "./img/5star/character/可莉.jpg")
		r.StaticFile("/img/5star8", "./img/5star/character/阿贝多.jpg")
		r.StaticFile("/img/5star9", "./img/5star/character/钟离.jpg")
		r.StaticFile("/img/5star10", "./img/5star/character/甘雨.jpg")
		r.StaticFile("/img/5star11", "./img/5star/character/魈.jpg")
		r.StaticFile("/img/5star12", "./img/5star/character/胡桃.jpg")
		r.StaticFile("/img/5star13", "./img/5star/character/优菈.jpg")

		r.StaticFile("/img/5stararms1", "./img/5star/arms/阿莫斯之弓.jpg")
		r.StaticFile("/img/5stararms2", "./img/5star/arms/尘世之锁.jpg")
		r.StaticFile("/img/5stararms3", "./img/5star/arms/风鹰剑.jpg")
		r.StaticFile("/img/5stararms4", "./img/5star/arms/贯虹之槊.jpg")
		r.StaticFile("/img/5stararms5", "./img/5star/arms/和璞鸢.jpg")
		r.StaticFile("/img/5stararms6", "./img/5star/arms/护摩之杖.jpg")
		r.StaticFile("/img/5stararms7", "./img/5star/arms/狼的末路.jpg")
		r.StaticFile("/img/5stararms8", "./img/5star/arms/磐岩结绿.jpg")
		r.StaticFile("/img/5stararms9", "./img/5star/arms/四风原典.jpg")
		r.StaticFile("/img/5stararms10", "./img/5star/arms/松籁响起之时.jpg")
		r.StaticFile("/img/5stararms11", "./img/5star/arms/天空之傲.jpg")
		r.StaticFile("/img/5stararms12", "./img/5star/arms/天空之脊.jpg")
		r.StaticFile("/img/5stararms13", "./img/5star/arms/天空之卷.jpg")
		r.StaticFile("/img/5stararms14", "./img/5star/arms/天空之刃.jpg")
		r.StaticFile("/img/5stararms15", "./img/5star/arms/天空之翼.jpg")
		r.StaticFile("/img/5stararms16", "./img/5star/arms/无工之剑.jpg")
		r.StaticFile("/img/5stararms17", "./img/5star/arms/终末嗟叹之诗.jpg")
		r.StaticFile("/img/5stararms18", "./img/5star/arms/斫峰之刃.jpg")
	} //5star
	{
		r.StaticFile("/img/3star1", "./img/3star/翠玉法球.jpg")
		r.StaticFile("/img/3star2", "./img/3star/弹弓.jpg")
		r.StaticFile("/img/3star3", "./img/3star/飞天御剑.jpg")
		r.StaticFile("/img/3star4", "./img/3star/黑缨枪.jpg")
		r.StaticFile("/img/3star5", "./img/3star/冷刃.jpg")
		r.StaticFile("/img/3star6", "./img/3star/黎明神剑.jpg")
		r.StaticFile("/img/3star7", "./img/3star/魔导绪论.jpg")
		r.StaticFile("/img/3star8", "./img/3star/沐浴龙血的剑.jpg")
		r.StaticFile("/img/3star9", "./img/3star/神射手之誓.jpg")
		r.StaticFile("/img/3star10", "./img/3star/讨龙英杰谭.jpg")
		r.StaticFile("/img/3star11", "./img/3star/铁影阔剑.jpg")
		r.StaticFile("/img/3star12", "./img/3star/鸦羽弓.jpg")
		r.StaticFile("/img/3star13", "./img/3star/以理服人.jpg")
	} //3star

	{
		r.StaticFile("/img/4star1", "./img/4star/丽莎.jpg")
		r.StaticFile("/img/4star2", "./img/4star/凝光.jpg")
		r.StaticFile("/img/4star3", "./img/4star/凯亚.jpg")
		r.StaticFile("/img/4star4", "./img/4star/北斗.jpg")
		r.StaticFile("/img/4star5", "./img/4star/匣里灭辰.jpg")
		r.StaticFile("/img/4star6", "./img/4star/匣里龙吟.jpg")
		r.StaticFile("/img/4star7", "./img/4star/弓藏.jpg")
		r.StaticFile("/img/4star8", "./img/4star/昭心.jpg")
		r.StaticFile("/img/4star9", "./img/4star/流浪乐章.jpg")
		r.StaticFile("/img/4star10", "./img/4star/烟绯.jpg")
		r.StaticFile("/img/4star11", "./img/4star/班尼特.jpg")
		r.StaticFile("/img/4star12", "./img/4star/砂糖.jpg")
		r.StaticFile("/img/4star13", "./img/4star/祭礼大剑.jpg")
		r.StaticFile("/img/4star14", "./img/4star/祭礼弓.jpg")
		r.StaticFile("/img/4star15", "./img/4star/祭礼残章.jpg")
		r.StaticFile("/img/4star16", "./img/4star/笛剑.jpg")
		r.StaticFile("/img/4star17", "./img/4star/罗莎莉亚.jpg")
		r.StaticFile("/img/4star18", "./img/4star/芭芭拉.jpg")
		r.StaticFile("/img/4star19", "./img/4star/菲谢尔.jpg")
		r.StaticFile("/img/4star20", "./img/4star/行秋.jpg")
		r.StaticFile("/img/4star21", "./img/4star/西风剑.jpg")
		r.StaticFile("/img/4star22", "./img/4star/西风大剑.jpg")
		r.StaticFile("/img/4star23", "./img/4star/西风猎弓.jpg")
		r.StaticFile("/img/4star24", "./img/4star/诺艾尔.jpg")
		r.StaticFile("/img/4star25", "./img/4star/辛焱.jpg")
		r.StaticFile("/img/4star26", "./img/4star/迪奥娜.jpg")
		r.StaticFile("/img/4star27", "./img/4star/重云.jpg")
		r.StaticFile("/img/4star28", "./img/4star/雨裁.jpg")
		r.StaticFile("/img/4star29", "./img/4star/雷泽.jpg")
		r.StaticFile("/img/4star30", "./img/4star/香菱.jpg")
	} //4star
	r.Run(":8000")
}

func gachaOnceRes(c *gin.Context) {
	username := c.Query("user_name")
	if username == "" {
		c.JSON(200, gin.H{
			"error": "query error",
		})
		c.Abort()
		return
	}

	db, err := gorm.Open("mysql", "root:071212@(127.0.0.1:3306)/genshin_impact?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var curUser gacha.User
	var rusultCharacter gacha.Character

	rslt := db.Where("user_name = ?", username).First(&curUser)
	if rslt.RowsAffected == 0 {
		curUser = gacha.User{
			UserName:     username,
			Cardpool:     gacha.Cpool1,
			Gacha_record: make([]gacha.Record, 1),
			Small_pool:   1,
			Big_pool:     1,
			EorA:         false,
		}
		db.Create(&curUser)
		rusultCharacter = curUser.Gacha_once()
		db.Save(&curUser)
		c.JSON(200, gin.H{
			"setuser": "ok",
			"result":  rusultCharacter,
		})
	} else {
		curUser.Cardpool = gacha.Cpool1
		rusultCharacter = curUser.Gacha_once()
		db.Save(&curUser)
		c.JSON(200, gin.H{
			"result": rusultCharacter,
		})
	}
}

func gachaTenRes(c *gin.Context) {
	username := c.Query("user_name")
	if username == "" {
		c.JSON(200, gin.H{
			"error": "query error",
		})
		c.Abort()
		return
	}

	db, err := gorm.Open("mysql", "root:071212@(127.0.0.1:3306)/genshin_impact?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var curUser gacha.User
	var rusultCharacter []gacha.Character

	rslt := db.Where("user_name = ?", username).First(&curUser)
	if rslt.RowsAffected == 0 {
		curUser = gacha.User{
			UserName:     username,
			Cardpool:     gacha.Cpool1,
			Gacha_record: make([]gacha.Record, 1),
			Small_pool:   1,
			Big_pool:     1,
			EorA:         false,
		}
		db.Create(&curUser)
		rusultCharacter = curUser.Gacha_ten()
		db.Save(&curUser)
		c.JSON(200, gin.H{
			"setuser": "ok",
			"result":  rusultCharacter,
		})
	} else {
		curUser.Cardpool = gacha.Cpool1
		rusultCharacter = curUser.Gacha_ten()
		db.Save(&curUser)
		c.JSON(200, gin.H{
			"result": rusultCharacter,
		})
	}
}
