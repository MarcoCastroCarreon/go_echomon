package mongo

import (
	"github.com/fatih/color"
	"github.com/kamva/mgm/v3"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitConnection(uri string, db string) (conneted bool, conectionError error) {
	color.Green("Connecting to Database...")
	err := mgm.SetDefaultConfig(nil, db, options.Client().ApplyURI(uri))

	connected := err == nil
	color.Cyan("Connected -> %v", connected)
	return connected, err
}
