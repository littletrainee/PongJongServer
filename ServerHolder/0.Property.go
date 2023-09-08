package ServerHolder

import (
	"net/http"
	"sync"

	"github.com/littletrainee/GetSet"
	Enum "github.com/littletrainee/PongJongServer/EnumHolder"
	"github.com/littletrainee/PongJongServer/PongJong/Handler"
)

type ServerHolder struct {
	Address         GetSet.Type[string]
	Port            GetSet.Type[[]string]
	PatternMap      map[Enum.Pattern]func(http.ResponseWriter, *http.Request)
	ConnectList     []string
	WaitForGameList []string
	GameHandlerList []Handler.Handler
	wg              *sync.WaitGroup
}
