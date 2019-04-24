package result

type SvcDowntime struct {
	StartTimestamp int64  `json:"start_timestamp"`
	List           []Svc `json:list`
}

type Svc struct {
	Name   string   `json:name`
	Details []Detail `json:"detail"`
}

type Detail struct {
	Timestamp   int64 `json:timestamp`
	NumEndpoint int8  `json:num_endpoint`
}
