package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	pb "go.etcd.io/etcd/api/v3/etcdserverpb"
	"go.etcd.io/etcd/pkg/v3/pbutil"
	"go.etcd.io/etcd/server/v3/wal"
	"go.etcd.io/etcd/server/v3/wal/walpb"
	"go.etcd.io/etcd/server/v3/datadir"
	"go.etcd.io/etcd/server/v3/etcdserver/api/snap"
	"go.etcd.io/etcd/raft/v3/raftpb"
	"io"
)

func parseWAL() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args[] string) error {
		var (
			walsnap walpb.Snapshot
			w *wal.WAL
			wmetadata []byte
			st raftpb.HardState
			ents []raftpb.Entry
			err error
		)

		// check whether the data directory exist or not
		if err := checkDataDir(); err != nil {
			return err
		}

		// Load all valid snapshot entries
		walDir := datadir.ToWalDir(dataDir)
		walSnaps, err := wal.ValidSnapshotEntries(nil, walDir)
		if err != nil {
			return err
		}

		// Load the newest snapshot
		ss := snap.New(nil, datadir.ToSnapDir(dataDir))
		snapshot, err := ss.LoadNewestAvailable(walSnaps)
		if err != nil {
			return err
		}

		// Load WAL data after the newest snapshot
		if snapshot != nil {
			walsnap.Index, walsnap.Term = snapshot.Metadata.Index, snapshot.Metadata.Term
		}

		repaired := false
		for {
			if w, err = wal.Open(nil, walDir, walsnap); err != nil {
				return fmt.Errorf("failed to open WAL, error: %v", err)
			}

			if wmetadata, st, ents, err = w.ReadAll(); err != nil {
				w.Close()
				// we can only repair ErrUnexpectedEOF and we never repair twice.
				if repaired || err != io.ErrUnexpectedEOF {
					return fmt.Errorf("failed to read WAL, cannot be repaired, error: %v", err)
				}
				if !wal.Repair(nil, walDir) {
					return fmt.Errorf("failed to repair WAL, error: %v", err)
				} else {
					fmt.Printf("repaired WAL, error: %v\n", err)
					repaired = true
				}
				continue
			}
			break
		}

		// Print results
		// 1. print newest snapshot's metadata
		if snapshot != nil {
			if s, err := formatStructInJSON(snapshot.Metadata, rawFormat); err != nil {
				return fmt.Errorf("failed to marshal snapshot metadata, rawFormat: %t, error: %v", rawFormat, err)
			} else {
				printJsonObject("Snapshot Metadata", s)
			}
		} else {
			printJsonObject("Snapshot metadata", "No any snapshots")
		}
		printSeparator()

		// 2. print metadata
		var metadata pb.Metadata
		pbutil.MustUnmarshal(&metadata, wmetadata)
		if s, err := formatStructInJSON(metadata, rawFormat); err != nil {
			return fmt.Errorf("failed to marshal metadata, rawFormat: %t, error: %v", rawFormat, err)
		} else {
			printJsonObject("Cluster Metadata", s)
		}
		printSeparator()

		// 3. print HardState
		if s, err := formatStructInJSON(st, rawFormat); err != nil {
			return fmt.Errorf("failed to marshal HardState, rawFormat: %t, error: %v", rawFormat, err)
		} else {
			printJsonObject("HardState", s)
		}
		printSeparator()

		// 4. print entries
		printJsonObject("Entry", fmt.Sprintf("Entry number: %d\n", len(ents)))
		if showDetails {
			for _, entry := range ents {
				if s, err := formatStructInJSON(entry, rawFormat); err != nil {
					return fmt.Errorf("failed to marshal Enntry, rawFormat: %t, error: %v", rawFormat, err)
				} else {
					printJsonObject("", s)
				}
			}
		}

		return nil
	}
}

func createWALCommand() *cobra.Command{
	var walCmd  = &cobra.Command {
		Use:   "wal",
		Short: "Parse wal files",
		Long: "Parse wal files",
		RunE: silenceUsage(parseWAL()),
	}

	return walCmd
}
