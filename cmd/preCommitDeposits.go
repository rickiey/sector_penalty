package cmd

import (
	"bytes"
	"context"
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/builtin/v9/miner"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/rickiey/loggo"
	"github.com/rickiey/sector_penalty/pkg"
	"github.com/spf13/cobra"
)

// precommitDepositsCmd
var preCommitDepositsCmd = &cobra.Command{
	Use:   "preCommitDeposits miner",
	Short: "miner preCommit deposits",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			loggo.Error("no params miner address")
			return
		}

		maddr, err := address.NewFromString(args[0])
		if err != nil {
			loggo.Error("miner address is invalid")
			return
		}

		head, err := lapi.ChainHead(ctx)
		if err != nil {
			loggo.Error(err.Error())
			return
		}

		tsk := types.NewTipSetKey(head.Cids()...)

		//minfo, err := lapi.StateMinerInfo(ctx, maddr, tsk)
		//if err != nil {
		//	loggo.Error(err.Error())
		//	return
		//}

		ma, err := lapi.StateGetActor(context.Background(), maddr, tsk)
		if err != nil {
			loggo.Error(err)
			return
		}

		//actorHead, err := service.RequestChainReadObj(ma.Head.String())
		actorHead, err := lapi.ChainReadObj(ctx, ma.Head)
		if err != nil {
			loggo.Error(err)
			return
		}

		var minerState miner.State

		if err := minerState.UnmarshalCBOR(bytes.NewReader(actorHead)); err != nil {
			loggo.Error(err)
			return
		}

		fmt.Printf("%v attoFIL = %v FIL \n", minerState.PreCommitDeposits, pkg.ToFloat64(minerState.PreCommitDeposits))
	},
}

func init() {
	rootCmd.AddCommand(preCommitDepositsCmd)
}
