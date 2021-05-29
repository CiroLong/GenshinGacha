package main

import (
	"Genshin/gacha"
	"fmt"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	db, err := gorm.Open("mysql", "root:071212@(127.0.0.1:3306)/genshin_impact?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var wg sync.WaitGroup
	wg.Add(1)

	for i := 1; i <= 1; i++ {
		go func() {
			username := "Long"

			var user1 gacha.User

			//db.CreateTable(&user1)
			db.Where("user_name = ?", username).First(&user1)
			fmt.Println(user1)
			user1.Cardpool = gacha.Cpool1
			//
			var res []gacha.Result
			for i := 1; i <= 1; i++ {
				res = append(res, user1.Gacha_ten())
				db.Save(&user1)
			}

			for i, j := range res {
				for k, s := range j {
					fmt.Println(10*(i)+k+1, "  ", s)
				}

			}

			//fmt.Println(100 - user1.Big_pool)

			//for i, value := range user1.Gacha_record {
			//	fmt.Println(i+1, value)
			//}
			wg.Done()
		}()
		time.Sleep(time.Millisecond * 10)
	}

	wg.Wait()
}
