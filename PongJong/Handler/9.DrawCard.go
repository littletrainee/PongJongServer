package Handler

import Enum "github.com/littletrainee/PongJongServer/EnumHolder"

func (h *Handler) DrawCard() {
	switch h.GameState.GameTurn.Get() {
	case 0:
		h.Player1.DrawCard(&h.Wall)
	case 1:
		h.Player2.DrawCard(&h.Wall)
	}
	h.Action.Set(Enum.CheckTsumo)
}
