package types

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

type ESSnapshotStatus struct {
	Snapshots []struct {
		Snapshot           string `json:"snapshot"`
		Repository         string `json:"repository"`
		UUID               string `json:"uuid"`
		State              string `json:"state"`
		IncludeGlobalState bool   `json:"include_global_state"`
		ShardsStats        struct {
			Initializing int64 `json:"initializing"`
			Started      int64 `json:"started"`
			Finalizing   int64 `json:"finalizing"`
			Done         int64 `json:"done"`
			Failed       int64 `json:"failed"`
			Total        int64 `json:"total"`
		} `json:"shards_stats"`
		Stats struct {
			Incremental struct {
				FileCount   int64 `json:"file_count"`
				SizeInBytes int64 `json:"size_in_bytes"`
			} `json:"incremental"`
			Processed struct {
				FileCount   int64 `json:"file_count"`
				SizeInBytes int64 `json:"size_in_bytes"`
			} `json:"processed"`
			Total struct {
				FileCount   int64 `json:"file_count"`
				SizeInBytes int64 `json:"size_in_bytes"`
			} `json:"total"`
			// TimeInMillis and StartTimeInMillis are not reliable
			StartTimeInMillis    int64 `json:"start_time_in_millis"`
			TimeInMillis         int64 `json:"time_in_millis"`
			NumberOfFiles        int64 `json:"number_of_files"`
			ProcessedFiles       int64 `json:"processed_files"`
			TotalSizeInBytes     int64 `json:"total_size_in_bytes"`
			ProcessedSizeInBytes int64 `json:"processed_size_in_bytes"`
		} `json:"stats"`
		Indices map[string]struct {
			ShardsStats struct {
				Initializing int64 `json:"initializing"`
				Started      int64 `json:"started"`
				Finalizing   int64 `json:"finalizing"`
				Done         int64 `json:"done"`
				Failed       int64 `json:"failed"`
				Total        int64 `json:"total"`
			} `json:"shards_stats"`
			Stats struct {
				Incremental struct {
					FileCount   int64 `json:"file_count"`
					SizeInBytes int64 `json:"size_in_bytes"`
				} `json:"incremental"`
				Total struct {
					FileCount   int64 `json:"file_count"`
					SizeInBytes int64 `json:"size_in_bytes"`
				} `json:"total"`
				// TimeInMillis and StartTimeInMillis are not reliable
				StartTimeInMillis    int64 `json:"start_time_in_millis"`
				TimeInMillis         int64 `json:"time_in_millis"`
				NumberOfFiles        int64 `json:"number_of_files"`
				ProcessedFiles       int64 `json:"processed_files"`
				TotalSizeInBytes     int64 `json:"total_size_in_bytes"`
				ProcessedSizeInBytes int64 `json:"processed_size_in_bytes"`
			} `json:"stats"`
			Shards map[string]struct {
				Stage string `json:"stage"`
				Stats struct {
					Incremental struct {
						FileCount   int64 `json:"file_count"`
						SizeInBytes int64 `json:"size_in_bytes"`
					} `json:"incremental"`
					Total struct {
						FileCount   int64 `json:"file_count"`
						SizeInBytes int64 `json:"size_in_bytes"`
					} `json:"total"`
					StartTimeInMillis    int64 `json:"start_time_in_millis"`
					TimeInMillis         int64 `json:"time_in_millis"`
					NumberOfFiles        int64 `json:"number_of_files"`
					ProcessedFiles       int64 `json:"processed_files"`
					TotalSizeInBytes     int64 `json:"total_size_in_bytes"`
					ProcessedSizeInBytes int64 `json:"processed_size_in_bytes"`
				} `json:"stats"`
				Node string `json:"node"`
			} `json:"shards"`
		} `json:"indices"`
	} `json:"snapshots"`
}

func (snapshotStatus *ESSnapshotStatus) Unmarshal(jsonbytes []byte) error {
	return json.Unmarshal(jsonbytes, snapshotStatus)
}

// Displays ESSnapshotStatus data at REPO level of detail
func (snapshotStatus *ESSnapshotStatus) PrintRepoTable() {
	table := tablewriter.NewWriter(os.Stdout)
	goodCell := tablewriter.Colors{tablewriter.Normal, tablewriter.FgHiGreenColor}
	goodRow := []tablewriter.Colors{goodCell, goodCell, goodCell, goodCell}
	table.SetHeader([]string{"Snapshot", "Repo", "State", "Percent"})
	for _, snapshot := range snapshotStatus.Snapshots {
		var percent int64
		if snapshot.Stats.TotalSizeInBytes != 0 {
			percent = 100 * snapshot.Stats.ProcessedSizeInBytes / snapshot.Stats.TotalSizeInBytes
		}
		table.Rich([]string{snapshot.Snapshot, snapshot.Repository, snapshot.State, fmt.Sprint(percent)}, goodRow)
	}
	table.Render()
}

// Displays ESSnapshotStatus data at INDEX level of detail
func (snapshotStatus *ESSnapshotStatus) PrintIndexTable() {
	table := tablewriter.NewWriter(os.Stdout)
	goodCell := tablewriter.Colors{tablewriter.Normal, tablewriter.FgHiGreenColor}
	// badCell := tablewriter.Colors{tablewriter.Normal, tablewriter.FgHiRedColor}
	goodRow := []tablewriter.Colors{goodCell, goodCell, goodCell, goodCell, goodCell, goodCell, goodCell}
	// badRow := []tablewriter.Colors{badCell, badCell, badCell, badCell, badCell, badCell, badCell}
	table.SetHeader([]string{"Snapshot", "Repo", "Index", "Percent"})
	for _, snapshot := range snapshotStatus.Snapshots {
		for indexname, index := range snapshot.Indices {
			var percent int64 = 0
			if index.Stats.TotalSizeInBytes != 0 {
				percent = 100 * index.Stats.ProcessedSizeInBytes / index.Stats.TotalSizeInBytes
			}
			table.Rich([]string{snapshot.Snapshot, snapshot.Repository, indexname, fmt.Sprint(percent)}, goodRow)
		}
	}
	table.Render()
}

// Displays ESSnapshotStatus data at SHARD level of detail
func (snapshotStatus *ESSnapshotStatus) PrintShardTable() {
	table := tablewriter.NewWriter(os.Stdout)
	goodCell := tablewriter.Colors{tablewriter.Normal, tablewriter.FgHiGreenColor}
	// badCell := tablewriter.Colors{tablewriter.Normal, tablewriter.FgHiRedColor}
	goodRow := []tablewriter.Colors{goodCell, goodCell, goodCell, goodCell, goodCell, goodCell, goodCell}
	// badRow := []tablewriter.Colors{badCell, badCell, badCell, badCell, badCell, badCell, badCell}
	table.SetHeader([]string{"Snapshot", "Repo", "Index", "Shard", "Node", "Stage", "Percent"})
	for _, snapshot := range snapshotStatus.Snapshots {
		for indexname, index := range snapshot.Indices {
			for shardname, shard := range index.Shards {
				var percent int64 = 0
				if shard.Stats.TotalSizeInBytes != 0 {
					percent = 100 * shard.Stats.ProcessedSizeInBytes / shard.Stats.TotalSizeInBytes
				}
				table.Rich([]string{snapshot.Snapshot, snapshot.Repository, indexname, shardname, shard.Node, shard.Stage, fmt.Sprint(percent)}, goodRow)
			}
		}
	}
	table.Render()
}
