package core

import (
	"github.com/thetatoken/ukulele/common"
	"github.com/thetatoken/ukulele/crypto"
)

// ConsensusEngine is the interface of a consensus engine.
type ConsensusEngine interface {
	ID() string
	PrivateKey() *crypto.PrivateKey
	GetTip() *ExtendedBlock
	GetEpoch() uint64
	GetLedger() Ledger
	AddMessage(msg interface{})
	FinalizedBlocks() chan *Block
	GetLastFinalizedBlock() *ExtendedBlock
}

// ValidatorManager is the component for managing validator related logic for consensus engine.
type ValidatorManager interface {
	SetConsensusEngine(consensus ConsensusEngine)
	GetProposer(blockHash common.Hash, epoch uint64) Validator
	GetValidatorSet(blockHash common.Hash) *ValidatorSet
}
