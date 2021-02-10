package cmd

import (
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/cheggaaa/pb/v3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/fanaticscripter/EggOrganizer/api"
	"github.com/fanaticscripter/EggOrganizer/config"
	"github.com/fanaticscripter/EggOrganizer/ei"
	"github.com/fanaticscripter/EggOrganizer/util"
)

type contractStatus int

const (
	Unknown contractStatus = iota
	NeverAttempted
	Attempted
	Done
)

var (
	_statusHideDone bool
	_statusNoId     bool
)

var _statusCommand = &cobra.Command{
	Use:     "status <contract-id>",
	Short:   "Bulk-retrieve members' status of a contract",
	Args:    cobra.ExactArgs(1),
	PreRunE: subcommandPreRunE,
	RunE: func(cmd *cobra.Command, args []string) error {
		contractId := strings.ToLower(args[0])
		players := _config.Players

		type result struct {
			jobId  int
			player config.Player
			status contractStatus
		}
		resultCh := make(chan result, len(players))
		done := make(chan bool)
		var results []result

		bar := pb.Simple.Start(len(players))
		go func() {
			for res := range resultCh {
				results = append(results, res)
				bar.Increment()
			}
			bar.Finish()
			done <- true
		}()

		var wg sync.WaitGroup
		for i, player := range players {
			wg.Add(1)
			go func(jobId int, contractId string, player config.Player) {
				resultCh <- result{
					jobId:  jobId,
					player: player,
					status: retrieveContractStatus(contractId, player.UserId),
				}
				wg.Done()
			}(i, contractId, player)
			// Do not hammer server too hard.
			time.Sleep(250 * time.Millisecond)
		}
		wg.Wait()
		close(resultCh)
		<-done

		table := [][]string{
			{"User ID", "Nickname", "Status"},
			{"-------", "--------", "------"},
		}
		sort.Slice(results, func(i, j int) bool {
			return results[i].jobId < results[j].jobId
		})
		for _, res := range results {
			if !_statusHideDone || res.status != Done {
				table = append(table, []string{
					res.player.UserId,
					res.player.Nickname,
					res.status.String(),
				})
			}
		}
		if _statusNoId {
			for i, row := range table {
				table[i] = row[1:]
			}
		}
		util.PrintTable(table)

		return nil
	},
}

func init() {
	_rootCmd.AddCommand(_statusCommand)
	_statusCommand.Flags().BoolVarP(&_statusHideDone, "hide-done", "H", false,
		"hide players who have completed this contract")
	_statusCommand.Flags().BoolVarP(&_statusNoId, "no-id", "n", false,
		"do not print the User ID column")
}

func retrieveContractStatus(contractId, userId string) contractStatus {
	resp, err := api.RequestFirstContact(
		&ei.EggIncFirstContactRequest{
			EiUserId: &userId,
		},
	)
	if err != nil {
		log.Errorf("failed to retrieve backup for user %s: %s", userId, err)
		return Unknown
	}
	if resp.Backup == nil || resp.Backup.Contracts == nil {
		encoded, _ := protojson.Marshal(resp)
		log.Errorf("invalid backup data for user %s: %s", userId, string(encoded))
		return Unknown
	}
	status := NeverAttempted
	for _, contract := range resp.Backup.Contracts.Archive {
		if *contract.Contract.Identifier == contractId {
			if int(*contract.NumGoalsAchieved) < len(contract.Contract.Goals) {
				status = Attempted
			} else {
				status = Done
			}
			break
		}
	}
	return status
}

func (s contractStatus) String() string {
	switch s {
	case NeverAttempted:
		return "NEVER ATTEMPTED"
	case Attempted:
		return "ATTEMPTED"
	case Done:
		return "DONE"
	default:
		return "UNKNOWN"
	}
}
