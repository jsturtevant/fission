/*
Copyright 2016 The Fission Authors.

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

package messageQueue

import (
	// "bytes"
	// "errors"
	// "fmt"
	// "io/ioutil"
	// "net/http"
	// "strings"

	// log "github.com/sirupsen/logrus"

	//"github.com/fission/fission"
	"github.com/fission/fission/tpr"
)

const (
	asqQueue = "queuename"
)

type (
	Asq struct {
		conn      string
		routerUrl string
	}
)

func makeAsqMessageQueue(routerUrl string, mqCfg MessageQueueConfig) (MessageQueue, error) {
	asq := Asq{
		conn:      "connection go here",
		routerUrl: routerUrl,
	}
	return asq, nil
}

func (asq Asq) subscribe(trigger *tpr.Messagequeuetrigger) (messageQueueSubscription, error) {
	return nil, nil
}

func (asq Asq) unsubscribe(subscription messageQueueSubscription) error {
	return nil
}
