package env

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

var configs = "ports : \n" +
	"  grpc : " + getENV("grpc_port") + "\n" +
	"database : \n" +
	"  pgdb : \n" +
	"   db_url : " + fmt.Sprintf("%q", getENV("db_url")) + "\n" +
	"   db_type : " + getENV("db_type") + "\n" +
	"   db_user : " + getENV("db_user") + "\n" +
	"   db_pass: " + getENV("db_pass") + "\n" +
	"   db_host: " + getENV("db_host") + "\n" +
	"   db_port: " + getENV("db_port") + "\n" +
	"   db_name: " + getENV("db_name") + "\n" +
	"   db_additional : \n" +
	"      postgres_ssl_mode: " + getENV("db_additional_ssl_mode") + "\n" +
	"services : \n" +
	"   base_url_auth: " + getENV("base_url_auth") + "\n" +
	"   base_url_out: " + getENV("base_url_out") + "\n" +
	"   base_url_customer: " + getENV("base_url_customer") + "\n" +
	"   base_url_monitoring: " + getENV("base_url_monitoring") + "\n" +
	"   base_url_coupon_reward: " + getENV("base_url_coupon_reward") + "\n" +
	"kafka :  " + getENV("kafka") + "\n" +
	"kafka_producer :  " + getENV("kafka_producer") + "\n" +
	"kafka_consumer :  " + getENV("kafka_consumer") + "\n" +
	"kafka_greward :  " + getENV("kafka_greward") + "\n" +
	"event_topics : \n" +
	"   point_topic_kafka: " + getENV("point_topic_kafka") + "\n" +
	"   point_topic_monitoring: " + getENV("point_topic_monitoring") + "\n" +
	"   kafka_redeem_coupon_topic: " + getENV("kafka_redeem_coupon_topic") + "\n" +
	"event_group_id :  " + getENV("event_group_id") + "\n" +
	"additional_prefix_repush :  " + getENV("additional_prefix_repush") + "\n" +
	"toggle_send_sms :  " + getENV("toggle_send_sms") + "\n" +
	"percent_disc : " + getENV("percent_disc") + "\n"

func getENV(envName string) string {
	return os.Getenv(envName)
}

type Dependency struct {
	Err    error
	DB     *sql.DB
	Params EnvironmentParameters
}

type (
	EnvironmentParameters struct {
		Ports struct {
			Gin string `yaml:"gin"`
		} `yaml:"ports"`
		Database struct {
			PGDB PostGresDB `yaml:"pgdb"`
		} `yaml:"database"`
	}

	PostGresDB struct {
		DBURL        string `yaml:"db_url"`
		DBType       string `yaml:"db_type"`
		DBUser       string `yaml:"db_user"`
		DBPass       string `yaml:"db_pass"`
		DBHost       string `yaml:"db_host"`
		DBPort       string `yaml:"db_port"`
		DBName       string `yaml:"db_name"`
		DBAdditional struct {
			PostgresSSLMode string `yaml:"postgres_ssl_mode"`
		} `yaml:"db_additional"`
		DBConfig string `yaml:"db_config"`
	}
)

// NewENV : reading through provided config file
func NewENV(configPath string) (*Dependency, error) {
	var settings Dependency
	config, err := os.Open(configPath)
	if err != nil {
		//if error use from configmap
		return NewENVFromMap()
	}

	defer func(config *os.File) {
		err = config.Close()
	}(config)
	d := yaml.NewDecoder(config)
	if err = d.Decode(&settings.Params); err != nil {
		return nil, err
	}

	return &settings, err
}

// NewENVFromMap : reading through provided configmap
func NewENVFromMap() (*Dependency, error) {
	var settings Dependency
	log.Println(configs)
	d := yaml.NewDecoder(strings.NewReader(configs))
	err := d.Decode(&settings.Params)
	if err != nil {
		return nil, err
	}
	return &settings, err
}

// SetupPostGresDBConnection : parse database parameters
func (eP *EnvironmentParameters) SetupPostGresDBConnection() *EnvironmentParameters {
	eP.Database.PGDB.DBConfig = fmt.Sprintf(eP.Database.PGDB.DBURL,
		eP.Database.PGDB.DBType,
		eP.Database.PGDB.DBUser,
		eP.Database.PGDB.DBPass,
		eP.Database.PGDB.DBHost,
		eP.Database.PGDB.DBPort,
		eP.Database.PGDB.DBName,
		eP.Database.PGDB.DBAdditional.PostgresSSLMode,
	)
	return eP
}
