package telemetry

//InstanceInfo is a deploed server data dto
type InstanceInfo struct {
	ID string `json:"id,omitempty"`

	Commit      string `json:"commit,omitempty"`
	Tag         string `json:"tag,omitempty"`
	BuiltAt     string `json:"built_at,omitempty"`
	ServiceName string `json:"service,omitempty"`
	RunID       string `json:"run_id,omitempty"`
}

//Usage is a usage accounting dto
type Usage struct {
	DockerHubID string `json:"docker_hub_id,omitempty"`
	ServerStart int    `json:"server_start,omitempty"`
	ServerStop  int    `json:"server_stop,omitempty"`

	Events    uint64 `json:"events,omitempty"`
	Errors    uint64 `json:"errors,omitempty"`
	EventsSrc string `json:"events_src,omitempty"`

	Source     string `json:"hashed_source_id,omitempty"`
	SourceType string `json:"source_type,omitempty"`

	Destination        string `json:"hashed_destination_id,omitempty"`
	DestinationType    string `json:"destination_type,omitempty"`
	DestinationMode    string `json:"destination_mode,omitempty"`
	DestinationMapping string `json:"destination_mappings,omitempty"`
	DestinationPkKeys  bool   `json:"destination_primary_keys,omitempty"`
	UsersRecognition   bool   `json:"users_recognition,omitempty"`

	Coordination string `json:"coordination,omitempty"`

	CLICommand     string `json:"cli_command,omitempty"`
	CLIStart       int    `json:"cli_start,omitempty"`
	CLIDateFilters bool   `json:"cli_date_filters,omitempty"`
	CLIState       bool   `json:"cli_state,omitempty"`
	CLIChunkSize   int64  `json:"cli_chunk_size,omitempty"`
}

//Errors is a error accounting dto
type Errors struct {
	ID       int64 `json:"id,omitempty"`
	Quantity int64 `json:"quantity,omitempty"`
}

//UserData is a registered user data dto
type UserData struct {
	Email       string `json:"email,omitempty"`
	Name        string `json:"name,omitempty"`
	Company     string `json:"company,omitempty"`
	EmailOptout bool   `json:"email_optout"`
	UsageOptout bool   `json:"telemetry_usage_optout"`
}

//Request is a telemetry request dto
type Request struct {
	Timestamp    string        `json:"timestamp,omitempty"`
	InstanceInfo *InstanceInfo `json:"instance_info,omitempty"`
	MetricType   string        `json:"metric_type,omitempty"`
	Usage        *Usage        `json:"usage,omitempty"`
	Errors       *Errors       `json:"errors,omitempty"`
	User         *UserData     `json:"user,omitempty"`
}
