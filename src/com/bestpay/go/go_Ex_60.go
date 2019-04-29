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

//roll 返回模拟一次投掷骰产生的（result,turnIsOver）
//若投掷骰的结果 outcome 为 1，那么这一轮 thisTurn的分数就会被抛弃，
//然后玩家的角色互换。否则，此次掷骰的值就会计入这一轮 thisTurn 中。
func roll(s score)(score,bool){

	outcome := rand.Intn(6) + 1

	if outcome == 1{

		return score{s.opponent,s.player,0},true
	}

	return score{s.player,s.opponent,outcome + s.thisTurn},false
}

//stay 返回停留时产生的结果（result,turnIsOver）。
//这一轮 thisTurn 的分数会计入该玩家的总分中，然后玩家的角色互换。
func stay(s score)(score,bool)  {

	return  score{s.opponent,s.player + s.thisTurn,0},true

}

// strategy 为任何给定的分数 score 返回一个动作 action
type strategy func(score) action

//strategy 返回一个策略，该策略继续投掷骰直到这一轮 thisTurn 至少为 k,然后停留
func stayAtK(k int) strategy  {

	return func(s score) action {
		if s.thisTurn >= k {
			return stay
		}

		return roll
	}
}

func play(strategy0,strategy1 strategy) int {

	strategies := []strategy{strategy0,strategy1}
	var s score
	var turnIsOver bool
	currentPlayer := rand.Intn(2)

	for s.player + s.thisTurn < win{
		action := strategies[currentPlayer](s)
		s,turnIsOver = action(s)
		if turnIsOver {
			currentPlayer = (currentPlayer + 1) % 2
		}
	}

	return currentPlayer

}

