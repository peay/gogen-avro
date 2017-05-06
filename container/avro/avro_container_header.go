/*
 * CODE GENERATED AUTOMATICALLY WITH github.com/peay/gogen-avro
 * THIS FILE SHOULD NOT BE EDITED BY HAND
 *
 * SOURCES:
 *     block.avsc
 *     header.avsc
 */

package avro

import (
	"io"
)

type AvroContainerHeader struct {
	Magic Magic
	Meta  map[string][]byte
	Sync  Sync
}

func DeserializeAvroContainerHeader(r io.Reader) (*AvroContainerHeader, error) {
	return readAvroContainerHeader(r)
}

func (r *AvroContainerHeader) Schema() string {
	return "{\"fields\":[{\"name\":\"magic\",\"type\":{\"name\":\"Magic\",\"size\":4,\"type\":\"fixed\"}},{\"name\":\"meta\",\"type\":{\"type\":\"map\",\"values\":\"bytes\"}},{\"name\":\"sync\",\"type\":{\"name\":\"Sync\",\"size\":16,\"type\":\"fixed\"}}],\"name\":\"AvroContainerHeader\",\"type\":\"record\"}"
}

func (r *AvroContainerHeader) Serialize(w io.Writer) error {
	return writeAvroContainerHeader(r, w)
}
