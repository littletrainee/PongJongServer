package ServerHolder

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/littletrainee/PongJongServer/ClientData"
	Enum "github.com/littletrainee/PongJongServer/EnumHolder"
)

func (sh *ServerHolder) MakePong(writer http.ResponseWriter, reader *http.Request) {
	var CD ClientData.ClientData
	if reader.Method == "POST" {
		if err := json.NewDecoder(reader.Body).Decode(&CD); err != nil {
			log.Println(err)
		} else {
			for i := range sh.GameHandlerList {
				if sh.GameHandlerList[i].Room.Get()[0] == CD.RoomUUID {
					if CD.GameData.Action == Enum.IsPong {
						sh.GameHandlerList[i].Action.Set(CD.GameData.Action)
					} else if CD.GameData.Action == Enum.NotPong {
						sh.GameHandlerList[i].Action.Set(Enum.DrawCard)
					}
					break
				}
			}
			if err := json.NewEncoder(writer).Encode(&CD); err != nil {
				log.Println(err)
			}
		}
	}
}
