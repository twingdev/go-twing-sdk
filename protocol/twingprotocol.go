package protocol

import (
	"github.com/google/uuid"
	"go-twing-sdk/proto/pb"
)

const TWING_PROTOCOL_ID = "/twing/protocol/1.0"

func NewEncryptedMessage(peerID string, aad, msg []byte) *pb.EncryptedMessage {
	m := &pb.Message{}
	Id, _ := uuid.NewUUID()
	Ids := Id.String()
	m.Id = Ids
	return &pb.EncryptedMessage{MsgData: m, PeerID: peerID, Aad: string(aad), Data: string(msg)}
}
