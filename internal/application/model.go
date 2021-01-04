package application

import (
	"log"

	"github.com/hamba/avro"
)

type Message struct {
	Content  string `avro:"content"`
	Priority int    `avro:"priority"`
	BundleID string `avro:"bundle_id"`
}

var (
	// MessageSchema ...
	MessageSchema *avro.RecordSchema
)

func Schema() *avro.RecordSchema {
	if MessageSchema == nil {
		schema, err := avro.Parse(`{
			"type": "record",
			"name": "message",
			"fields" : [
				{"name": "content", "type": "string"},
				{"name": "priority", "type": "int"},
				{"name": "bundle_id", "type": "string"}
			]
		}`)
		if err != nil {
			log.Fatal(err)
		}
		MessageSchema = schema.(*avro.RecordSchema)
	}
	return MessageSchema
}
