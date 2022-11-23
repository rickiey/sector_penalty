package cmd

import (
	"bytes"
	"fmt"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/lotus/chain/types"

	"github.com/filecoin-project/go-address"
	"github.com/rickiey/sector_penalty/pkg"

	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/builtin/v9/power"
	"github.com/filecoin-project/go-state-types/builtin/v9/reward"
	"github.com/rickiey/loggo"
	"github.com/spf13/cobra"
)

// windowPostCmd represents the windowPost command
var windowPostCmd = &cobra.Command{
	Use:   "windowPost [miner]",
	Short: "windowPoSt failed penalty for one sector",
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

		penalty32G := ExpectedRewardForPower(rewardActorState.ThisEpochRewardSmoothed, powerActorState.ThisEpochQAPowerSmoothed,
			big.NewInt(32<<30), 2880*351/100)

		fmt.Printf("Current network penalty for windowPoSt failure of 32GB sector: %v FIL = %v attoFIL \n", pkg.ToFloat64(penalty32G), penalty32G)

		if len(args) != 1 {
			return
		}

		mineraddr, err := address.NewFromString(args[0])
		if err != nil {
			loggo.Error(err)
			return
		}

		fs, err := lapi.StateMinerFaults(ctx, mineraddr, tsk)
		if err != nil {
			loggo.Error(err)
			return
		}

		rs, err := lapi.StateMinerRecoveries(ctx, mineraddr, tsk)
		if err != nil {
			loggo.Error(err)
			return
		}

		fsc, err := fs.Count()
		if err != nil {
			loggo.Error(err)
			return
		}

		rsc, err := rs.Count()
		if err != nil {
			loggo.Error(err)
			return
		}
		minfo, err := lapi.StateMinerInfo(ctx, mineraddr, tsk)
		if err != nil {
			loggo.Error(err)
			return
		}

		faultsectorPower := big.NewInt(int64((fsc + rsc) * uint64(minfo.SectorSize/(32<<30))))

		wpostPenalty24h := big.Mul(faultsectorPower, penalty32G)
		fmt.Printf("miner %v fault sectors %v , recoveries sectors %v \n", mineraddr.String(), fsc, rsc)
		fmt.Printf("miner %v Estimated 24-hour windowPoSt penalty for failed sectors: %v FIL = %v attoFIL \n",
			mineraddr.String(), pkg.ToFloat64(wpostPenalty24h), wpostPenalty24h)
	},
}

func init() {
	rootCmd.AddCommand(windowPostCmd)
}
