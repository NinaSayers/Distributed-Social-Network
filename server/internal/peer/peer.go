package peer

import kademlia "github.com/jackverneda/godemlia/pkg"

type Peer interface {
	*kademlia.Node
}
