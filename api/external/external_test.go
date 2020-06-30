package external

import (
	"testing"

	"example.org/luksam/kiwi-server/config"
	"example.org/luksam/kiwi-server/db"
	"example.org/luksam/kiwi-server/test"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"

	log "github.com/sirupsen/logrus"
)

var testConfig config.Database

// DatabaseTestSuiteBase provides the setup and teardown of the database for every test-run.
type DatabaseTestSuiteBase struct {
}

// SetupSuite is called once before starting the test-suite.
func (b *DatabaseTestSuiteBase) SetupSuite() {
	testConfig = test.GetTestConfiguration().Database

	// check if connecting is possible
	err := db.TryConnect(testConfig)
	if err != nil {

		panic(errors.Wrap(err, "could not connect to the test database, make sure the configured PostgreSQL instance is running ('sudo service postgresql start') and accessible and the required table is created (search for 'unit test howto: create required database' in code)"))

		// -----------------------------------------
		// unit test howto: create required database
		// -----------------------------------------
		//
		// The following steps are required for the unit test to work.
		//
		// start postgres console:
		//   sudo -u postgres psql
		//
		// create the kiwi_server_test_user user:
		//   CREATE ROLE kiwi_server_test_user WITH LOGIN PASSWORD 'dbpassword';
		//
		// create the kiwi_server_test database
		//   CREATE DATABASE kiwi_server_test WITH OWNER kiwi_server_test_user;
		//
		// connect to the kiwi_server_test database
		//   \connect kiwi_server_test
		//
		// create required extension
		//   CREATE EXTENSION hstore;
		//
		// quit
		//   \quit
	}

	// check if required table exists
	tableExists, err := db.CheckTableExists("device_up", testConfig)
	if err != nil {
		panic(err)
	}
	if !*tableExists {
		// create table since it does currently not exist
		log.Info("table 'device_up' did not exist yet, trying to create it")
		sqlStatementCreateTestTable := `
CREATE TABLE device_up (
	id UUID PRIMARY KEY,
	received_at TIMESTAMP WITH TIME ZONE NOT NULL,
	dev_eui BYTEA NOT NULL,
	device_name VARCHAR(100) NOT NULL,
	application_id BIGINT NOT NULL,
	application_name VARCHAR(100) NOT NULL,
	frequency BIGINT NOT NULL,
	dr SMALLINT NOT NULL,	
	adr BOOLEAN NOT NULL,
	f_cnt BIGINT NOT NULL,
	f_port SMALLINT NOT NULL,
	tags HSTORE NOT NULL,
	data BYTEA NOT NULL,
	rx_info JSONB NOT NULL,
	object JSONB NOT NULL
);

CREATE INDEX idx_device_up_received_at ON device_up(received_at);
CREATE INDEX idx_device_up_dev_eui ON device_up(dev_eui);
CREATE INDEX idx_device_up_application_id ON device_up(application_id);
CREATE INDEX idx_device_up_frequency ON device_up(frequency);
CREATE INDEX idx_device_up_dr ON device_up(dr);
CREATE INDEX idx_device_up_tags ON device_up(tags);
		`
		err := db.RunSQLStatementForUnitTest(sqlStatementCreateTestTable, testConfig)
		if err != nil {
			panic(err)
		}
	}

}

// SetupTest is called before every test.
func (b *DatabaseTestSuiteBase) SetupTest() {
	// remove old table content
	err := db.RunSQLStatementForUnitTest(`TRUNCATE device_up;`, testConfig)
	if err != nil {
		panic(err)
	}
}

// TearDownTest is called after every test.
func (b *DatabaseTestSuiteBase) TearDownTest() {
	// remove old table content
	err := db.RunSQLStatementForUnitTest(`TRUNCATE device_up;`, testConfig)
	if err != nil {
		panic(err)
	}
}

type APITestSuite struct {
	suite.Suite
	DatabaseTestSuiteBase
}

func TestAPI(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}
