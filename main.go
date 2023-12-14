package main

import (
	"context"
	"fmt"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/patrickmn/go-cache"
)

func main() {
	d := dog{}
	GetChanNosAndConsistency(context.TODO(), d.name)
}

type dog struct {
	weight int
	name string
}

func goCache() {
	productFeatCache := cache.New(24*time.Hour, 24*time.Hour)
	productFeatCache.SetDefault("1", dog{})
	doginter, ok := productFeatCache.Get("1")
	fmt.Println(ok)
	fmt.Println(doginter.(dog).weight)
}

func GetChanNosAndConsistency(context context.Context, ProductFeature string) (chanNos string, consistency int, err error) {
	if ProductFeature == "" {
		return "1", 0, err
	}
	// 将查到的数据转为json
	js, err := simplejson.NewJson([]byte(ProductFeature))
	if err != nil {
		fmt.Println(err, "New")
	}

	chanNos, err = js.Get("chan_nos").String()
	if err != nil {
		fmt.Println(err)
		chanNos = "0"
	}

	consistency, err = js.Get("chan_consistency").Int()
	if err != nil {
		fmt.Println(err)
	}
	return
}
