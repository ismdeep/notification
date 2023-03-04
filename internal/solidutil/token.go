package solidutil

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/btcsuite/btcd/btcutil/base58"

	"github.com/ismdeep/notification/pkg/core"
)

// DefaultSwarmKey default notification swarm key
const DefaultSwarmKey = "nsk_4k3ipmdSvQdRffwMXkV6gjJkuHYmgxHb"

var SwarmKey string

func init() {
	SwarmKey = os.Getenv("NOTIFICATION_SWARM_KEY")
	switch SwarmKey {
	case "":
		SwarmKey = DefaultSwarmKey
	}
}

// key headers
const (
	KeyHeaderTokenID           = "t_"  // token id
	KeyHeaderUserID            = "u_"  // user id
	KeyHeaderProjectID         = "p_"  // project id
	KeyHeaderAESKey            = "a_"  // aes key
	KeyHeaderNotificationToken = "nt_" // notification token
	KeyHeaderMsgID             = "m_"  // msg id
)

func newID(n int) string {
	const charset = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	bytes := make([]byte, n)
	for i := range bytes {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}
		bytes[i] = charset[randomIndex.Int64()]
	}
	return string(bytes)
}

func solid(s string) string {
	h := sha256.New()
	_, err1 := h.Write([]byte(SwarmKey))
	core.PanicIf(err1)
	_, err2 := h.Write([]byte(s))
	core.PanicIf(err2)
	return base58.Encode(h.Sum(nil))
}

func UserID(username string) string {
	return KeyHeaderUserID + solid(username)
}

func TokenID(userID string, tokenName string) string {
	return KeyHeaderTokenID + solid(fmt.Sprintf("%v|%v", userID, tokenName))
}

func MsgID(userID string, customerMsgID string, content string) string {
	return KeyHeaderMsgID + solid(fmt.Sprintf("%v|%v|%v", userID, customerMsgID, content))
}

func ProjectID(projectName string) string {
	return KeyHeaderProjectID + solid(projectName)
}

func RandToken() string {
	return KeyHeaderNotificationToken + newID(32)
}

func RandAESKey() string {
	return KeyHeaderAESKey + newID(32)
}
