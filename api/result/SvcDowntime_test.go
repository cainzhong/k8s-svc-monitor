package result

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	filename := "./k8s_svc_downtime.json"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	svcDowntime := &SvcDowntime{}
	json.Unmarshal(b, svcDowntime)
	fmt.Println(fmt.Sprintf("%+v",svcDowntime))
}
