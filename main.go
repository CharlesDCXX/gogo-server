package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"time"

	"github.com/patrickmn/go-cache"
)

// ProductFeatCache 产品特性缓存
var productFeatCache *cache.Cache

func init() {
	//todo_dual_camera 修改过期时间
	productFeatCache = cache.New(24*time.Hour, 24*time.Hour)
}
func main() {
	// 假设有一个 map[string]int 变量
	m := map[string]int{
		"1": 1,
		"2": 2,
	}
	a := 1
	result, ok := m[string(a)]
	fmt.Println(result, ok)
}

type dog struct {
	weight int
	name   string
}

func goCache() {
	productFeatCache := cache.New(24*time.Hour, 24*time.Hour)
	productFeatCache.SetDefault("1", dog{})
	doginter, ok := productFeatCache.Get("1")
	fmt.Println(ok)
	fmt.Println(doginter.(dog).name + "1")
}

func GetChanNosAndConsistency(context context.Context, ProductFeature string) (result map[string]int, err error) {
	result = make(map[string]int)
	if ProductFeature == "" {
		return
	}
	// 将查到的数据转为json
	js, err := simplejson.NewJson([]byte(ProductFeature))
	if err != nil {
		fmt.Println(err, "New")
	}

	cloudInfoChanMap, err := js.Get("cloud_info_chan_map").Map()
	if err != nil {
		fmt.Println(err)
	}
	for k, v := range cloudInfoChanMap {
		value, ok := v.(json.Number)
		if ok {
			intValue, err := value.Int64()
			if err == nil {
				result[k] = int(intValue)
			} else {
				return nil, err
			}
		} else {
			return
		}
	}
	return
}
