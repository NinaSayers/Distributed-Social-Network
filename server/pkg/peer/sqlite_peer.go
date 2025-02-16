package core

import (
	"crypto/sha1"

	"github.com/NinaSayers/Distributed-Social-Network/server/pkg/persistence"
	kademlia "github.com/jackverneda/godemlia/pkg"
	base58 "github.com/jbenet/go-base58"
)

type SqlitePeer struct {
	kademlia.Node
}

func NewSqlitePeer(ip string, port, bootstrapPort int, dbPath string, isBootstrapNode bool) *SqlitePeer {
	db := persistence.NewSqliteDb(dbPath) // Initialize SQLite DB
	newPeer := kademlia.NewNode(ip, port, bootstrapPort, db, isBootstrapNode)

	return &SqlitePeer{*newPeer}
}

func (p *SqlitePeer) Store(data *[]byte) (string, error) {
	hash := sha1.Sum(*data)
	key := base58.Encode(hash[:])

	_, err := p.StoreValue(key, data) // Store in Kademlia
	if err != nil {
		return "", err
	}

	return key, nil
}
