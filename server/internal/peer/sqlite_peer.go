package peer

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"

	"github.com/NinaSayers/Distributed-Social-Network/server/internal/persistence"
	kademlia "github.com/jackverneda/godemlia/pkg"
	base58 "github.com/jbenet/go-base58"
)

type SqlitePeer struct {
	kademlia.Node
}

func InitSqlitePeer(ip string, port, bootPort int, dbPath string, script string, bootstrap bool) *SqlitePeer {
	peer := NewSqlitePeer(ip, 8080, 32140, dbPath, script, bootstrap)
	addr := fmt.Sprintf("%s:%d", ip, port)
	go peer.CreateGRPCServer(addr)
	return peer
}

func NewSqlitePeer(ip string, port, bootstrapPort int, dbPath string, script string, isBootstrapNode bool) *SqlitePeer {
	db := persistence.NewSqliteDb(dbPath, script) // Initialize SQLite DB
	newPeer := kademlia.NewNode(ip, port, bootstrapPort, db, isBootstrapNode)

	return &SqlitePeer{*newPeer}
}

func (p *SqlitePeer) Store(entity string, data *[]byte) (string, error) {
	payload := map[string]interface{}{}
	err := json.Unmarshal(*data, &payload)
	if err != nil {
		return "", err
	}

	hash := sha1.Sum(*data)
	id := base58.Encode(hash[:])

	switch entity {
	case "user":
		email, ok := payload["email"].(string)
		if !ok {
			return "", fmt.Errorf("email is not a string")
		}
		hash = sha1.Sum(base58.Decode(email))
		id = base58.Encode(hash[:])

	default:
	}

	_, err = p.StoreValue(entity, id, data) // Store in Kademlia
	if err != nil {
		return "", err
	}

	return id, nil
}
