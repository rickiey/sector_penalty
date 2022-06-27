package cmd

import (
	"bytes"
	"fmt"

	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/builtin/v8/miner"
	"github.com/filecoin-project/go-state-types/builtin/v8/power"
	"github.com/filecoin-project/go-state-types/builtin/v8/reward"
	"github.com/filecoin-project/go-state-types/builtin/v8/util/smoothing"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/rickiey/loggo"
	"github.com/spf13/cobra"

	"github.com/rickiey/sector_penalty/pkg"
)

const (
	ss2KiB   = 2 << 10
	ss8MiB   = 8 << 20
	ss512MiB = 512 << 20
	ss32GiB  = 32 << 30
	ss64GiB  = 64 << 30
)

// pledgeCmd represents the pledge command
var pledgeCmd = &cobra.Command{
	Use:   "init-pledge",
	Short: "sector current InitialPledge calculate",

	Run: func(cmd *cobra.Command, args []string) {
		head, err := lapi.ChainHead(ctx)
		if err != nil {
			loggo.Error(err.Error())
			return
		}

		tsk := types.NewTipSetKey(head.Cids()...)

		act, err := lapi.StateGetActor(ctx, builtin.RewardActorAddr, tsk)
		if err != nil {
			loggo.Error(err)
			return
		}

		actorHead, err := lapi.ChainReadObj(ctx, act.Head)
		if err != nil {
			loggo.Error(err)
			return
		}

		var rewardActorState reward.State

		if err := rewardActorState.UnmarshalCBOR(bytes.NewReader(actorHead)); err != nil {
			loggo.Error(err)
			return
		}

		actst, err := lapi.StateGetActor(ctx, builtin.StoragePowerActorAddr, tsk)
		if err != nil {
			loggo.Error(err)
			return
		}

		stactorHead, err := lapi.ChainReadObj(ctx, actst.Head)
		if err != nil {
			loggo.Error(err)
			return
		}

		var powerActorState power.State

		if err := powerActorState.UnmarshalCBOR(bytes.NewReader(stactorHead)); err != nil {
			loggo.Error(err)
			return
		}

		circSupply, err := lapi.StateVMCirculatingSupplyInternal(ctx, tsk)
		if err != nil {
			return
		}

		weight := miner.QAPowerForWeight(ss32GiB, 1555200, big.NewInt(0), big.NewInt(0))
		InitialPledge := miner.InitialPledgeForPower(weight, powerActorState.ThisEpochQualityAdjPower, rewardActorState.
			ThisEpochRewardSmoothed, smoothing.FilterEstimate(powerActorState.ThisEpochQAPowerSmoothed), circSupply.FilCirculating)

		fmt.Println("InitialPledge for 32 GiB (FIL): ", pkg.ToFloat64(InitialPledge))
		fmt.Println("InitialPledge for 1  TiB (FIL): ", pkg.ToFloat64(big.Mul(InitialPledge, big.NewInt(32))))
	},
}

func init() {
	rootCmd.AddCommand(pledgeCmd)
}
