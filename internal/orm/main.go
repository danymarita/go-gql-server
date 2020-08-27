package orm

import (
	"errors"

	"github.com/danymarita/go-gql-server/internal/gql/resolvers/transformations"

	"github.com/markbates/goth"

	"github.com/danymarita/go-gql-server/internal/logger"
	"github.com/danymarita/go-gql-server/internal/orm/models"

	"github.com/danymarita/go-gql-server/internal/orm/migration"

	"github.com/danymarita/go-gql-server/pkg/utils"
	//Imports the database dialect of choice
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
)

var autoMigrate, logMode, seedDB bool
var dsn, dialect string

// ORM struct to holds the gorm pointer to db
type ORM struct {
	DB *gorm.DB
}

func init() {
	dialect = utils.MustGet("GORM_DIALECT")
	dsn = utils.MustGet("GORM_CONNECTION_DSN")
	seedDB = utils.MustGetBool("GORM_SEED_DB")
	logMode = utils.MustGetBool("GORM_LOGMODE")
	autoMigrate = utils.MustGetBool("GORM_AUTOMIGRATE")
}

// Factory creates a db connection with the selected dialect and connection string
func Factory() (*ORM, error) {
	db, err := gorm.Open(dialect, dsn)
	if err != nil {
		log.Panic("[ORM] err: ", err)
	}
	orm := &ORM{
		DB: db,
	}

	// Log every SQL command on dev, @prod: this should be disabled?
	db.LogMode(logMode)
	// Automigrate tables
	if autoMigrate {
		err = migration.ServiceAutoMigration(orm.DB)
	}
	log.Info("[ORM] Database connecion initialized")
	return orm, err
}

//FindUserByAPIKey finds the user that is related to the API key
func (o *ORM) FindUserByAPIKey(apiKey string) (*models.User, error) {
	db := o.DB.New()
	uak := &models.UserAPIKey{}
	if apiKey == "" {
		return nil, errors.New("API key is empty")
	}
	if err := db.Preload("User").Where("api_key = ?", apiKey).Find(uak).Error; err != nil {
		return nil, err
	}
	return &uak.User, nil
}

// FindUserByJWT finds the user that is related to the APIKey token
func (o *ORM) FindUserByJWT(email string, provider string, userID string) (*models.User, error) {
	db := o.DB.New()
	up := &models.UserProfile{}
	
}