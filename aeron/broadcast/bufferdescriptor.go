/*
Copyright 2016 Stanislav Liberman

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package broadcast

import (
	"github.com/lirm/aeron-go/aeron/util"
	"log"
)

var BufferDescriptor = struct {
	TAIL_INTENT_COUNTER_OFFSET int32
	TAIL_COUNTER_OFFSET        int32
	LATEST_COUNTER_OFFSET      int32
	TRAILER_LENGTH             int32
}{
	0,
	util.SIZEOF_INT64,
	util.SIZEOF_INT64 * 2,
	util.CACHE_LINE_LENGTH * 2,
}

func CheckCapacity(capacity int32) {
	if !util.IsPowerOfTwo(capacity) {
		log.Fatalf("Capacity must be a positive power of 2 + TRAILER_LENGTH: capacity=%d", capacity)
	}
}
