package finality

// This should be updated to match the gRPC interface of the Babylon finality gadget client
// https://github.com/babylonlabs-io/finality-gadget

import "github.com/babylonlabs-io/finality-gadget/types"

type IFinalityGadgetClient interface {
	/// QueryIsBlockBabylonFinalized checks if the given L2 block is finalized by the Babylon finality gadget
	QueryIsBlockBabylonFinalized(block *types.Block) (bool, error)

	/// QueryBlockRangeBabylonFinalized searches for a row of consecutive finalized blocks in the block range, and returns
	QueryBlockRangeBabylonFinalized(blocks []*types.Block) (*uint64, error)

	/// QueryBtcStakingActivatedTimestamp returns the timestamp when the BTC staking is activated
	QueryBtcStakingActivatedTimestamp() (uint64, error)

	/// QueryIsBlockFinalizedByHeight returns the btc finalization status of a block at given height by querying the local db
	QueryIsBlockFinalizedByHeight(height uint64) (bool, error)

	/// QueryIsBlockFinalizedByHash returns the btc finalization status of a block at given hash by querying the local db
	QueryIsBlockFinalizedByHash(hash string) (bool, error)

	/// QueryLatestFinalizedBlock returns the latest finalized block by querying the local db
	QueryLatestFinalizedBlock() (*types.Block, error)

	// Closes the client
	Close() error
}
