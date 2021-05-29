package gacha

import (
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
)

type Result []Character

type Record struct {
	Time  time.Time `json:"time" time-format:"2006-05-25"`
	Chact Character `json:"chact"`
}

type User struct {
	gorm.Model `json:"-"`
	UserName   string `json:"user_name"` //  用户名

	Cardpool     `json:"cardpool" gorm:"-"` //卡池
	Gacha_record []Record                   `json:"gacha_record" gorm:"-"` //历史抽卡记录

	Small_pool int  `json:"small_pool"` //小井
	Big_pool   int  `json:"big_pool"`   //大井
	EorA       bool `json:"eora"`       //是否大保底
}

var Users []User

func (c *User) Gacha_once() Character {

	time.Sleep(time.Microsecond)
	if c.Big_pool == BIG_POOL {
		//大保底
		if c.EorA {
			result := c.Special[0]
			//修改
			c.Big_pool = 1
			c.Small_pool = 1
			c.EorA = false
			c.Gacha_record = append(c.Gacha_record, Record{time.Now(), result})
			return result
		}
		//随机
		rand.Seed(time.Now().UnixNano())
		rand_number := rand.Intn(len(c.Content.Star_5))
		result := c.Content.Star_5[rand_number]
		if result.ID != c.Special[0].ID { //是否为up
			c.EorA = true //不为up则下次必出
		}
		c.Big_pool = 1
		c.Small_pool = 1
		c.Gacha_record = append(c.Gacha_record, Record{time.Now(), result})
		return result
	}
	if c.Small_pool == SMALL_POOL {
		//必出4星
		rand.Seed(time.Now().UnixNano())
		var result Character
		rand1 := rand.Intn(len(c.Content.Star_4))
		result = c.Content.Star_4[rand1]
		c.Big_pool += 1
		c.Small_pool = 1
		c.Gacha_record = append(c.Gacha_record, Record{time.Now(), result})
		return result
	}
	rand.Seed(time.Now().UnixNano())
	rand_number := rand.Float32() //随机数函数
	var result Character
	if rand_number <= rate_5star { //出5星
		rand1 := rand.Intn(len(c.Content.Star_5))
		result = c.Content.Star_5[rand1]
		if result.ID != c.Special[0].ID { //是否为up
			c.EorA = true //不为up则下次必出
		}
		c.Big_pool = 1
		c.Small_pool = 1
		c.Gacha_record = append(c.Gacha_record, Record{time.Now(), result})
	} else if rand_number <= rate_4star { //出四星
		rand.Seed(time.Now().UnixNano())
		rand1 := rand.Intn(len(c.Content.Star_4))
		result = c.Content.Star_4[rand1]
		c.Big_pool += 1
		c.Small_pool = 1
		c.Gacha_record = append(c.Gacha_record, Record{time.Now(), result})
	} else { //三星
		rand.Seed(time.Now().UnixNano())
		rand1 := rand.Intn(len(c.Content.Star_3))
		result = c.Content.Star_3[rand1]
		c.Big_pool += 1
		c.Small_pool += 1
		c.Gacha_record = append(c.Gacha_record, Record{time.Now(), result})
	}
	return result
}

func (c *User) Gacha_ten() Result {
	var results Result
	for i := 0; i < 10; i++ {
		results = append(results, c.Gacha_once())
	}
	return results
}
