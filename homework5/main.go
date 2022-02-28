package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"sync"
)


func RandomFatRate() float64 {
	base := 3 + rand.Intn(50)
	fatRate := float64(base) +  rand.Float64()
	return fatRate
}

func UserOrder(order int) string{
	if order >= 1 && order <= 9 {
		return "000" + strconv.Itoa(order)
	} else if order >= 10 && order <= 99 {
		return "00" + strconv.Itoa(order)
	} else if order >= 100 && order <= 999 {
		return "0" + strconv.Itoa(order)
	} else {
		return strconv.Itoa(order)
	}
}

func UserRankingByChannel(userId string, userOrder string, channelMessage chan string, userSlice []string) {
	fatRate := RandomFatRate()

	message := strconv.FormatFloat(fatRate, 'f', 3, 64) + "," + userOrder + "," + userId
	fmt.Println(message)
	channelMessage <- message

	UserGetRanking(userOrder, channelMessage, userSlice)
}

func UserGetRanking(userOrder string, channelMessage chan string, userSlice []string) {
	message := <- channelMessage
	userSlice = append(userSlice, message)
	sort.Strings(userSlice)
	for size := 1; size <= 1000; size++ {
		readMessage := userSlice[size]
		if strings.Contains(readMessage, "," + userOrder + ",") {
			fmt.Println(userOrder, " rank No is", size)
		}
	}
}



func UseChannel() {
	fmt.Println("使用 Channel 实现排行榜")
	channelMessage := make(chan string, 1000)
	userSlice := make([]string, 1000)

	for size := 1; size <=1000; size++ {
		userId := "user" + strconv.Itoa(size)
		userOrder := UserOrder(size)
		go UserRankingByChannel(userId, userOrder, channelMessage, userSlice)
	}



}



func UseLock() {
	fmt.Println("使用 锁 实现排行榜")

	// userSlice := make([]string, 1000)
	// userFatRateMap := make(map[int]string)
	userFatRateMap := sync.Map{}
	messageList := list.New()
	for size := 1; size <=1000; size++ {
		userId := "user" + strconv.Itoa(size)
		userOrder := UserOrder(size)
		fatRate := RandomFatRate()
		message := strconv.FormatFloat(fatRate, 'f', 3, 64) + "," + userOrder + "," + userId + "," + strconv.Itoa(size)
		//userSlice = append(userSlice, message)
		// userFatRateMap[size] = message
		userFatRateMap.Store(size, message)
		messageList.PushBack(message)



		fmt.Println(userFatRateMap.Load(size))


	}


}


func main() {
	fmt.Println("体脂率排行榜!")

	// UseChannel()
	UseLock()

}
