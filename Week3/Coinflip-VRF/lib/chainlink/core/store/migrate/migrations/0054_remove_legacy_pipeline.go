package migrations

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/pressly/goose/v3"
)

const up54 = `
ALTER TABLE log_broadcasts DROP COLUMN job_id;
DROP TABLE service_agreements;
DROP TABLE eth_task_run_txes;
DROP TABLE task_runs;
DROP TABLE task_specs;
DROP TABLE flux_monitor_round_stats;
DROP TABLE job_runs;
DROP TABLE job_spec_errors;
DROP TABLE initiators;
DROP TABLE job_specs;

DROP TABLE run_results;
DROP TABLE run_requests;
DROP TABLE sync_events;

ALTER TABLE log_broadcasts RENAME COLUMN job_id_v2 TO job_id;
ALTER TABLE job_spec_errors_v2 RENAME TO job_spec_errors;
`

type queryer interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

func Up54(ctx context.Context, tx *sql.Tx) error {
	if err := CheckNoLegacyJobs(ctx, tx); err != nil {
		return err
	}
	if _, err := tx.ExecContext(ctx, up54); err != nil {
		return err
	}
	return nil
}

func Down54(ctx context.Context, tx *sql.Tx) error {
	return errors.New("irreversible migration")
}

// CheckNoLegacyJobs ensures that there are no legacy job specs
func CheckNoLegacyJobs(ctx context.Context, ds queryer) error {
	var count int
	if err := ds.QueryRowContext(ctx, `SELECT COUNT(*) FROM job_specs WHERE deleted_at IS NULL`).Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		return errors.Errorf("cannot migrate; this release removes support for legacy job specs but there are still %d in the database. Please migrate these jobs specs to the V2 pipeline (further details found here: https://docs.chain.link/docs/jobs/migration-v1-v2/) and make sure job_specs table is empty (run sql command: `TRUNCATE job_specs CASCADE;`), then run the migration again. These operations are NOT REVERSIBLE, so it is STRONGLY RECOMMENDED that you take a database backup before continuing", count)
	}
	return nil
}

var Migration54 = goose.NewGoMigration(54, &goose.GoFunc{RunTx: Up54}, &goose.GoFunc{RunTx: Down54})
