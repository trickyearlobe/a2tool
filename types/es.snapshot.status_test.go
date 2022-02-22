package types

import (
	"os"
	"testing"
)

var esSnapshotStatusTestData []byte
var esSnapshotStatus ESSnapshotStatus

// Load our sample data for later tests
func TestEsSnapshotStatusLoadTestData(t *testing.T) {
	data, err := os.ReadFile("../test_data/es/_snapshot/_status.json")
	if err != nil {
		t.Fatalf("Failed to read sample data: %v", err)
	}
	esSnapshotStatusTestData = data
}

// Make sure we can unmarshal our sample data for later tests
func TestEsSnapshotStatusUnmarshal(t *testing.T) {
	if esSnapshotStatusTestData == nil {
		t.SkipNow()
	}
	err := esSnapshotStatus.Unmarshal(esSnapshotStatusTestData)
	if err != nil {
		t.Fatalf("Failed to unmarshal data: %v", err)
	}
}

// Make sure we have a slice of snapshots containing 1 snapshot
func TestEsSnapshotStatusSnapshots(t *testing.T) {
	if len(esSnapshotStatus.Snapshots) == 0 {
		t.SkipNow()
	}
	if len(esSnapshotStatus.Snapshots) != 1 {
		t.Errorf("Expected 1 snapshot but found %v", len(esSnapshotStatus.Snapshots))
	}
}

// Check we can read the snapshot IDs
func TestEsSnapshotStatusSnapshotsID(t *testing.T) {
	if len(esSnapshotStatus.Snapshots) == 0 {
		t.SkipNow()
	}
	if (esSnapshotStatus.Snapshots[0].Snapshot) != "20220221003105" {
		t.Errorf("Expected snapshot name to be 20220221003105 but got %v", esSnapshotStatus.Snapshots[0].Snapshot)
	}
}

// Check we can read the data about shards for a named index
func TestEsSnapshotStatusSnapshotsIndexShardsInitializing(t *testing.T) {
	if len(esSnapshotStatus.Snapshots) == 0 {
		t.SkipNow()
	}
	numShards := esSnapshotStatus.Snapshots[0].Indices["comp-7-r-2022.02.18"].ShardsStats.Initializing
	if numShards != 5 {
		t.Errorf("Expected named index shards initializing to be 5 but got %v", numShards)
	}
}

func TestEsSnapshotStatusSnapshotsPrintRepo(t *testing.T) {
	if len(esSnapshotStatus.Snapshots) == 0 {
		t.SkipNow()
	}
	esSnapshotStatus.PrintRepoTable()
}

func TestEsSnapshotStatusSnapshotsPrintIndex(t *testing.T) {
	if len(esSnapshotStatus.Snapshots) == 0 {
		t.SkipNow()
	}
	esSnapshotStatus.PrintIndexTable()
}

func TestEsSnapshotStatusSnapshotsPrintShard(t *testing.T) {
	if len(esSnapshotStatus.Snapshots) == 0 {
		t.SkipNow()
	}
	esSnapshotStatus.PrintShardTable()
}
