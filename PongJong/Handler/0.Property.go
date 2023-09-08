package Handler

import (
	"sync"

	"github.com/littletrainee/GetSet"
	Enum "github.com/littletrainee/PongJongServer/EnumHolder"

	"github.com/littletrainee/PongJongServer/PongJong/GameState"
	"github.com/littletrainee/PongJongServer/PongJong/Player"
	PS "github.com/littletrainee/PongJongServer/PongJong/PrintWinnerAndScore"
	"github.com/littletrainee/PongJongServer/PongJong/Wall"
)

type Handler struct {
	Room                     GetSet.Type[[]string]
	Wall                     Wall.Wall
	Player1, Player2, RongBy Player.Player
	GameState                GameState.GameState
	wg                       *sync.WaitGroup
	PlayerDecision           GetSet.Type[string]
	Winner                   GetSet.Type[string]
	Action                   GetSet.Type[Enum.Action]
	PrintWin                 PS.PrintWinnerAndScore
}
