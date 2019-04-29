package main


import "fmt"
import "math/rand"

const (

	win            = 100 //在一场Pig游戏中获胜的分数
	gamesPerSeries = 10  //每次连续模拟游戏的数量
)

//总分 score包括每个玩家前几轮的得分以及本轮中当前玩家的得分
type score struct {
	player,opponent,thisTurn int
}

// action 将一个动作随机转换为一个分数
type action func(current score) (result score, turnIsOver bool)

