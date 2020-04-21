package inputs

import (
	"sync"

	"github.com/Duet3D/DSF-APIs/godsfapi/v2/types"
)

// Inputs holds all available code channels
type Inputs []InputChannel

var channelMap map[types.CodeChannel]*InputChannel
var initChannelMap sync.Once

func (ch *Inputs) init() {
	channelMap = make(map[types.CodeChannel]*InputChannel)
	for _, i := range *ch {
		channelMap[types.CodeChannel(i.Name)] = &i
	}
}

// Get will return the Channel to the given types.CodeChannel.
// It will return SPI for unknown types.
func (ch *Inputs) Get(cc types.CodeChannel) *InputChannel {
	initChannelMap.Do(ch.init)

	return channelMap[cc]
}
