// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package mqtt

import (
	"testing"

	"github.com/dapr/components-contrib/bindings"
	"github.com/stretchr/testify/assert"
)

func TestParseMetadata(t *testing.T) {
	m := bindings.Metadata{}
	m.Properties = map[string]string{"URL": "a", "Topic": "a"}
	mq := MQTT{}
	mm, err := mq.getMQTTMetadata(m)
	assert.Nil(t, err)
	assert.Equal(t, "a", mm.URL)
	assert.Equal(t, "a", mm.Topic)
}
