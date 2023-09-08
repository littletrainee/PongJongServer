package ServerHolder

import (
	"github.com/littletrainee/PongJongServer/GameData"
	"github.com/littletrainee/PongJongServer/PongJong/GameState"
	"github.com/littletrainee/PongJongServer/PongJong/Player"
)

func getdata(p1, p2 *Player.Player, gs *GameState.GameState) GameData.GameData {
	var GD GameData.GameData
	GD.GameOn = gs.GameOn.Get()
	GD.ThisPlayerHand = p1.Hand.Get()
	GD.ThisPlayerMeld = p1.Meld.Get()
	GD.ThisPlayerRiver = p1.River.Get()
	GD.ThisPlayerCodeNumber = p1.CodeNumber.Get()
	GD.ThisPlayerIsRiiChi = p1.IsRiiChi.Get()
	GD.OppositePlayerName = p2.Name.Get()
	GD.OppositePlayerMeld = p2.Meld.Get()
	GD.OppositePlayerRiver = p2.River.Get()
	GD.WhoToDiscard = gs.GameTurn.Get() + 1
	return GD
}
