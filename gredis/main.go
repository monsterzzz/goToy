package main

import (
	"fmt"
	redis2 "github.com/go-redis/redis/v7"
	redis "goToy/gredis/cache"
	"goToy/gredis/model"
	"golang.org/x/exp/rand"
	"reflect"
	"strconv"
	"sync"
	"time"
)

const exampleKey = "hello"
const hashExample = "myHash"
const listExample1 = "myList1"
const listExample2 = "myList2"
const setExample1 = "mySet1"
const setExample2 = "mySet2"
const sortSetExample1 = "mySortSet"

func main() {

}

func redisLock() {
	client := redis.Cache
	defer client.Close()
	var wg *sync.WaitGroup
	wg = &sync.WaitGroup{}
	//
	lock := sync.Mutex{}
	client.Set("go", 0, 1*time.Hour)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			lock.Lock()
			fmt.Println(fmt.Sprintf("thread-%d get LOCK", i))
			v1, _ := client.Get("go").Int()
			if v1 == 0 {
				client.Incr("go")
			}
			v2, _ := client.Get("go").Int()
			fmt.Println(v2)
			lock.Unlock()
			fmt.Println(fmt.Sprintf("thread-%d release LOCK", i))
			wg.Done()
		}(i)
	}
	wg.Wait()
	v1, _ := client.Get("go").Int()
	fmt.Println(v1)
}

func rank() {
	client := redis.Cache
	defer client.Close()
	// 排行榜
	rand.Seed(uint64(time.Now().Unix()))
	keys := make([]string, 0)
	for i := 0; i <= 100; i++ {
		keys = append(keys, makeRandomString(10))
	}
	// 模拟点击
	for i := 0; i <= 100000; i++ {
		client.ZIncrBy("rank", 1, keys[rand.Intn(100)])
	}
	v1, e := client.ZRange("rank", 0, 100).Result()
	fmt.Println(v1, e)
	v2, e := client.ZRevRangeWithScores("rank", 0, 9).Result()
	fmt.Println(v2, e, reflect.TypeOf(v2))
}

func SortSet() {
	client := redis.Cache
	defer client.Close()
	client.ZRemRangeByScore(sortSetExample1, "0", "20")
	rand.Seed(uint64(time.Now().Unix()))
	for i := 0; i < 10; i++ {
		tmp := &redis2.Z{
			Score:  float64(rand.Intn(20)),
			Member: makeRandomString(5),
		}
		client.ZAdd(sortSetExample1, tmp)
	}

	v1, _ := client.ZRange(sortSetExample1, 0, 20).Result()
	fmt.Println(v1)
	v2, _ := client.ZCard(sortSetExample1).Result()
	fmt.Println(v2)
	v3, _ := client.ZCount(sortSetExample1, "0", "10").Result()
	fmt.Println(v3)
	v4, _ := client.ZRevRange(sortSetExample1, 0, 10).Result()
	fmt.Println(v4)
	v5, _ := client.ZIncrBy(sortSetExample1, 1.0, "jcbik").Result()
	fmt.Println(v5)
	v6, _ := client.ZScore(sortSetExample1, "jcbik").Result()
	fmt.Println(v6)
	v7, _ := client.ZRangeByScoreWithScores(sortSetExample1, &redis2.ZRangeBy{
		Min:    "0",
		Max:    "20",
		Offset: 0,
		Count:  0,
	}).Result()
	fmt.Println(v7)
}

func Set() {
	client := redis.Cache
	defer client.Close()
	client.Del(setExample1)
	client.Del(setExample2)
	for i := 0; i < 10; i++ {
		client.SAdd(setExample1, "monster"+strconv.Itoa(i))
	}
	for i := 5; i < 15; i++ {
		client.SAdd(setExample2, "monster"+strconv.Itoa(i))
	}

	v1, _ := client.SCard(setExample1).Result()
	fmt.Println(v1)                                          // 10
	v2, _ := client.SDiff(setExample2, setExample1).Result() // 取差集
	fmt.Println(v2)                                          // [monster10 monster12 monster14 monster11 monster13]
	v3, _ := client.SDiff(setExample1, setExample2).Result() // 取差集
	fmt.Println(v3)                                          // [monster4 monster3 monster2 monster1 monster0]
	v4, _ := client.SDiffStore("save", setExample1, setExample2).Result()
	fmt.Println(v4)
	v5, err := client.SMembers("save").Result()
	fmt.Println(v5, err)
	v6, _ := client.SIsMember(setExample1, "monster1").Result()
	fmt.Println(v6) // true
	v7, _ := client.SIsMember(setExample1, "monster").Result()
	fmt.Println(v7) // false
	v8, _ := client.SMove(setExample1, setExample2, "monster1").Result()
	fmt.Println(v8)
	v9, _ := client.SMembers(setExample1).Result()
	fmt.Println(v9)

}

func List() {
	client := redis.Cache
	defer client.Close()

	client.Del(listExample1)
	for i := 0; i <= 10; i++ {
		client.LPush(listExample1, "monster"+strconv.Itoa(i))
	}
	v1, _ := client.LRange(listExample1, 0, 20).Result()
	fmt.Println(v1)
	v11, _ := client.LPop(listExample1).Result() // GET HEAD
	fmt.Println(v11)
	v12, _ := client.RPop(listExample1).Result() // GET LAST
	fmt.Println(v12)
	v13, _ := client.BLPop(5*time.Second, listExample1).Result() // wait for 5s
	fmt.Println(v13)

	client.Del(listExample2)
	for i := 0; i <= 10; i++ {
		client.RPush(listExample2, "monster"+strconv.Itoa(i))
	}
	v2, _ := client.LRange(listExample2, 0, 20).Result()
	fmt.Println(v2)
}

func HmAPI() {
	client := redis.Cache
	defer client.Close()
	client.HSet(hashExample, "money", 22.5)
	v, _ := client.HIncrBy(hashExample, "age", 1).Result()
	fmt.Println(v)

	v1, _ := client.HIncrByFloat(hashExample, "money", 0.654).Result()
	fmt.Println(v1)
	v2, _ := client.HKeys(hashExample).Result()
	fmt.Println(v2)
	v3, _ := client.HVals(hashExample).Result()
	fmt.Println(v3)
	v4, _ := client.HLen(hashExample).Result()
	fmt.Println(v4)
}

func HGetAll() {
	client := redis.Cache
	defer client.Close()
	// 获得所有
	v, _ := client.HGetAll(hashExample).Result()
	fmt.Println(v)
}

func HExists() {
	// hash表 key中是否存在这个字段
	client := redis.Cache
	defer client.Close()
	v, _ := client.HExists(hashExample, "name").Result()
	fmt.Println(v, reflect.TypeOf(v))
}

func HmSet() {
	client := redis.Cache
	defer client.Close()

	m := make(map[string]interface{})
	m["name"] = "monster"
	m["age"] = 16
	m["desc"] = "hello world"
	client.HMSet(hashExample, m)
	v, _ := client.HMGet(hashExample, "age").Result()
	fmt.Println(v)
}

func Mget() {
	// 获取多个
	client := redis.Cache
	defer client.Close()
	client.Set("hello1", "ss", 1*time.Hour)
	v, _ := client.MGet(exampleKey, "hello1").Result()
	fmt.Println(v) // [f ss]
}

func setBit() {
	// 偏移位数
	client := redis.Cache
	defer client.Close()
	fmt.Println([]rune("abc")) // 97 98 99
	client.Set(exampleKey, "a", 1*time.Hour)
	// 97 二进制 1100001
	// 98 二进制 1100010
	// 99 二进制 1100011
	client.SetBit(exampleKey, 6, 1)
	v, _ := client.Get(exampleKey).Result()
	fmt.Println(v)
	client.SetBit(exampleKey, 7, 0)
	v, _ = client.Get(exampleKey).Result()
	fmt.Println(v)
	client.SetBit(exampleKey, 5, 1)
	v, _ = client.Get(exampleKey).Result()
	fmt.Println(v)
}

func getSet() {
	// 设定新值并返回旧值
	client := redis.Cache
	client.Set(exampleKey, "world", 1*time.Hour)
	v, _ := client.GetSet(exampleKey, "a").Result()
	fmt.Println(v)
}

func getSubStr() {
	// 获取截取字符
	client := redis.Cache
	client.Set(exampleKey, "world", 1*time.Hour)
	v, _ := client.GetRange(exampleKey, 0, 10).Result()
	fmt.Println(v, len(v))
}

func set() {
	// set
	// key value 过期时间
	client := redis.Cache
	client.Set(exampleKey, "world", 1*time.Second)
	v, _ := client.Get(exampleKey).Result()
	fmt.Println(v)
	time.Sleep(3 * time.Second)
	v, e := client.Get(exampleKey).Result()
	fmt.Println(v, e) // error : redis:nil 键不存在
}

func makeRandomString(length int) string {
	tmpStr := make([]rune, 0)
	for i := 0; i < length; i++ {
		tmpStr = append(tmpStr, rune(rand.Intn(122-97)+97))
	}
	return string(tmpStr)
}

func InsertRandomMessage() {
	for i := 0; i < 20; i++ {
		model.DB.Create(&model.Student{
			Name:        makeRandomString(15),
			Age:         rand.Intn(100),
			Sex:         rand.Intn(2),
			Description: makeRandomString(30),
		})
	}

	var students []model.Student
	model.DB.Find(&students)
	fmt.Println(students)
}
