package gacha

const (
	rate_5star float32 = 0.006
	rate_4star float32 = 0.051
	SMALL_POOL         = 10
	BIG_POOL           = 90
)

type Content struct {
	Star_3 []Character `json:"star_3"`
	Star_4 []Character `json:"star_4"`
	Star_5 []Character `json:"star_5"`
} //卡池内容

type Cardpool struct {
	Content `json:"content"` //匿名成员
	Special []Character      `json:"special"` //指定up
}

var Cpool1 Cardpool = Cardpool{
	Content: Content{
		Star_3: star3s,
		Star_4: star4s,
		Star_5: star5s,
	},
	Special: []Character{
		{
			Name:    "温迪",
			Star:    5,
			ImgPath: "/img/5star6",
		},
	},
}
