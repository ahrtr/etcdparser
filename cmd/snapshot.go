package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"go.etcd.io/etcd/server/v3/wal"
	"go.etcd.io/etcd/server/v3/datadir"
	"go.etcd.io/etcd/server/v3/etcdserver"
	"go.etcd.io/etcd/server/v3/etcdserver/api/snap"
	"go.etcd.io/etcd/server/v3/etcdserver/api/v2store"
)

func parseSnapshot() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args[] string) error {
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

		if len(walSnaps) == 0 {
			fmt.Println("No any snapshots")
			return nil
		}

		// Load the newest snapshot
		ss := snap.New(nil, datadir.ToSnapDir(dataDir))
		snapshot, err := ss.LoadNewestAvailable(walSnaps)
		if err != nil {
			return err
		}

		if nil == snapshot {
			// It should run into this branch
			return errors.New("failed to load the newest snapshot")
		}

		if rawFormat {
			fmt.Printf("Snapshot Metadata: \n%+v\n", snapshot.Metadata)
			if showDetails {
				printSeparator()
				fmt.Printf("Snapshot Data:\n%+v\n", string(snapshot.Data))
			}
		} else {
			if snapMetadata, err := formatStructInJSON(snapshot.Metadata, rawFormat); err != nil {
				return fmt.Errorf("failed to marshal snapshot metadata, rawFormat: %t, error: %v", rawFormat, err)
			} else {
				fmt.Printf("Snapshot Metadata: \n%s\n", snapMetadata)
			}

			if showDetails {
				printSeparator()

				// restore store from snapshot
				st := v2store.New(etcdserver.StoreClusterPrefix, etcdserver.StoreKeysPrefix)
				if err = st.Recovery(snapshot.Data); err != nil {
					return fmt.Errorf("failed to recover from snapshot, error: %v", err)
				}

				if snapData, err := formatStructInJSON(st, rawFormat); err != nil {
					fmt.Errorf("failed to marshal snapshot data, rawFormat: %t, error: %v", rawFormat, err)
				} else {
					fmt.Printf("Snapshot Data: \n%s\n", snapData)
				}
			}
		}

		return nil
	}
}

func createSnapCommand() *cobra.Command{
	var snapCmd  = &cobra.Command {
		Use:   "snap",
		Short: "Parse snap files",
		Long: "Parse snap files",
		RunE: silenceUsage(parseSnapshot()),
	}

	return snapCmd
}
