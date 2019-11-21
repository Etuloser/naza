// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package bitrate_test

import (
	"testing"
	"time"

	"github.com/q191201771/naza/pkg/assert"
	"github.com/q191201771/naza/pkg/bitrate"
)

func TestBitrate(t *testing.T) {
	var b *bitrate.Bitrate
	b = bitrate.NewBitrate(10)
	b.Add(1000)
	r := b.Rate()
	assert.Equal(t, 800, r)
	time.Sleep(100 * time.Millisecond)
	b.Rate()
}
