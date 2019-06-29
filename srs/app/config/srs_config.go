package config

import (
	"encoding/json"
	"io/ioutil"
)

type SrsConfig struct {
	ListenPort     uint32                `json:"listen_port"`
	Pid            string                `json:"pid"`
	ChunkSize      uint32                `json:"chunk_size"`
	MaxConnections uint32                `json:"max_connection"`
	WorkDir        string                `json:"work_dir"`
	VHosts         map[string]*VHostConf `json:"vhosts"`
}

func (this *SrsConfig) amendDefault() {
	if this.ListenPort == 0 {
		this.ListenPort = 1935
	}

	if this.Pid == "" {
		this.Pid = "./srs.pid"
	}

	if this.ChunkSize == 0 {
		this.ChunkSize = 60000
	}

	if this.MaxConnections == 0 {
		this.MaxConnections = 1000
	}

	if this.WorkDir == "" {
		this.WorkDir = "./"
	}

	for _, v := range this.VHosts {
		v.amendDefault()
	}
}

type HeartBeatConf struct {
	Enabled   string  `json:"enabled"`
	Interval  float64 `json:"interval"`
	Url       string  `json:"url"`
	DeviceId  string  `json:"device_id"`
	Summeries string  `json:"summeries"`
}

func (this *HeartBeatConf) amendDefault() {
	if this.Enabled != "on" {
		this.Enabled = "off"
	}

	if this.Interval <= 0 {
		this.Interval = 9.3
	}

	if this.Url == "" {
		this.Url = "http://127.0.0.1:8085/api/v1/servers"
	}

	if this.Summeries == "" {
		this.Summeries = "on"
	}
}

type StatsConf struct {
	Enabled string   `json:"enabled"`
	Disk    []string `json:"disk"`
}

func (this *StatsConf) amendDefault() {
	if this.Enabled == "" {
		this.Enabled = "off"
	}
}

type HttpApiConf struct {
	Enabled     string `json:"enabled"`
	Listen      uint32 `json:"listen"`
	Crossdomain string `json:"crossdomain"`
}

func (this *HttpApiConf) amendDefault() {
	if this.Enabled == "" {
		this.Enabled = "off"
	}

	if this.Listen == 0 {
		this.Listen = 1985
	}

	if this.Crossdomain == "" {
		this.Crossdomain = "on"
	}
}

type HttpServerConf struct {
	Enabled string `json:"enabled"`
	Listen  uint32 `json:"listen"`
	Dir     string `json:"dir"`
}

func (this *HttpServerConf) amendDefault() {
	if this.Enabled == "" {
		this.Enabled = "off"
	}

	if this.Listen == 0 {
		this.Listen = 8080
	}

	if this.Dir == "" {
		this.Dir = "./html"
	}
}

type SecurityConf struct {
	Enabled string `json:"enabled"`
}

func (this *SecurityConf) amendDefault() {
	if this.Enabled == "" {
		this.Enabled = "off"
	}
}

type DvrConf struct {
	Enabled         string `json:"enabled"`
	DvrPlan         string `json:"dvr_plan"`
	DvrPath         string `json:"dvr_path"`
	DvrDuration     uint32 `json:"dvr_duration"`
	DvrWaitKeyFrame string `json:"dvr_wait_keyframe"`
	TimerJitter     string `json:"timer_jitter"` //full, zero, off
}

func (this *DvrConf) amendDefault() {
	if this.Enabled == "" {
		this.Enabled = "off"
	}

	if this.DvrPlan == "" {
		this.DvrPlan = "session"
	}

	if this.DvrPath == "" {
		this.DvrPath = "html/[app]/[stream].[timestamp].flv"
	}

	if this.DvrDuration == 0 {
		this.DvrDuration = 30
	}

	if this.DvrWaitKeyFrame == "" {
		this.DvrWaitKeyFrame = "on"
	}

	if this.TimerJitter == "" {
		this.TimerJitter = "full"
	}
}

type HttpStaticConf struct {
	Enabled string `json:"enabled"`
	Mount   string `json:"mount"`
	Dir     string `json:"dir"`
}

func (this *HttpStaticConf) amendDefault() {
	if this.Enabled == "" {
		this.Enabled = "off"
	}

	if this.Mount == "" {
		this.Mount = "[vhost]/hls"
	}

	if this.Dir == "" {
		this.Dir = "html/hls"
	}
}

type HttpRemuxConf struct {
	Enabled   string `json:"enabled"`
	FastCache uint32 `json:"fast_cache"`
	Mount     string `json:"mount"`
	HStrs     string `json:"hstrs"`
}

func (this *HttpRemuxConf) amendDefault() {
	if this.Enabled == "" {
		this.Enabled = "off"
	}

	if this.Mount == "" {
		this.Mount = "[vhost]/[app]/[stream].flv"
	}

	if this.HStrs == "" {
		this.HStrs = "on"
	}
}

type HlsConf struct {
	Enabled         string  `json:"enabled"`
	HlsFragment     uint32  `json:"hls_fragment"`      //the hls fragment in seconds, the duration of a piece of ts.
	HlsTdRatio      float64 `json:"hls_td_ratio"`      //the hls m3u8 target duration ratio
	HlsAofRatio     float64 `json:"hls_aof_ration"`    //the audio overflow ratio.
	HlsWindow       uint32  `json:"hls_window"`        //the hls window in seconds, the number of ts in m3u8.
	HlsOnError      string  `json:"hls_on_error"`      //the error strategy
	HlsPath         string  `json:"hls_path"`          //the hls output path.
	HlsM3u8File     string  `json:"hls_m3u8_file"`     //the hls m3u8 file name.
	HlsTsFile       string  `json:"hls_ts_file"`       //the hls ts file name.
	HlsTsFloor      string  `json:"hls_ts_floor"`      //whether use floor for the hls_ts_file path generation.
	HlsEntryPrefix  string  `json:"hls_entry_prefix"`  //the hls entry prefix, which is base url of ts url.
	HlsAcodec       string  `json:"hls_acodec"`        //the default audio codec of hls.
	HlsVcodec       string  `json:"hls_vcodec"`        //the default video codec of hls.
	HlsCleanup      string  `json:"hls_cleanup"`       //whether cleanup the old expired ts files.
	HlsDispose      uint32  `json:"hls_dispose"`       //the timeout in seconds to dispose the hls,dispose is to remove all hls files, m3u8 and ts files.
	HlsNbNotify     uint32  `json:"hls_nb_notify"`     //the max size to notify hls,to read max bytes from ts of specified cdn network,
	HlsWaitKeyframe string  `json:"hls_wait_keyframe"` //whether wait keyframe to reap segment,
}

func (this *HlsConf) amendDefault() {
	if this.Enabled == "" {
		this.Enabled = "off"
	}

	if this.HlsFragment == 0 {
		this.HlsFragment = 10
	}

	if this.HlsTdRatio <= 0 {
		this.HlsTdRatio = 1.5
	}

	if this.HlsAofRatio <= 0 {
		this.HlsAofRatio = 2.0
	}

	if this.HlsWindow <= 0 {
		this.HlsWindow = 60
	}

	if this.HlsOnError == "" {
		this.HlsOnError = "continue"
	}

	if this.HlsPath == "" {
		this.HlsPath = "html"
	}

	if this.HlsM3u8File == "" {
		this.HlsM3u8File = "[app]/[stream].m3u8"
	}

	if this.HlsTsFile == ""  {
		this.HlsTsFile = "[app]/[stream]-[seq].ts"
	}

	if this.HlsTsFloor == "" {
		this.HlsTsFloor = "off"
	}

	if this.HlsAcodec == "" {
		this.HlsAcodec = "aac"
	}

	if this.HlsVcodec == "" {
		this.HlsVcodec = "avc"
	}

	if this.HlsCleanup == "" {
		this.HlsCleanup = "on"
	}

	if this.HlsNbNotify == 0 {
		this.HlsNbNotify = 64
	}

	if this.HlsWaitKeyframe == "" {
		this.HlsWaitKeyframe = "on"
	}
}

type HttpHooksConf struct {
	Enabled     string `json:"enabled"`
	OnConnect   string `json:"on_connect"`
	OnClose     string `json:"on_close"`
	OnPublish   string `json:"on_publish"`
	OnUnpublish string `json:"on_unpublish"`
	OnPlay      string `json:"on_play"`
	OnStop      string `json:"on_stop"`
	OnDvr       string `json:"on_dvr"`
	OnHls       string `json:"on_hls"`
	OnHlsNotify string `json:"on_hls_notify"`
}

func (this *HttpHooksConf) amendDefault() {
	if this.Enabled == "" {
		this.Enabled = "off"
	}
}

type PublishConf struct {
	ParseSps string `json:"parse_sps"`
}

func (this *PublishConf) amendDefault() {
	if this.ParseSps == "" {
		this.ParseSps = "on"
	}
}

type VHostConf struct {
	Enabled              string          `json:"enabled"`
	MinLatency           string          `json:"min_latency"`
	GopCache             string          `json:"gop_cache"`
	QueueLength          uint32          `json:"queue_length"`
	SendMinInterval      uint32          `json:"send_min_interval"`
	ReduceSequenceHeader string          `json:"reduce_sequence_header"`
	Publish1stPktTimeout uint32          `json:"publish_1stpkt_timeout"`
	PublishNormalTimeout uint32          `json:"publish_normal_timeout"`
	Forward              []string        `json:"forward"`
	ChunkSize            uint32          `json:"chunk_size"`
	TimerJitter          string          `json:"time_jitter"`
	MixCorrect           string          `json:"mix_correct"`
	Atc                  string          `json:"atc"`
	AtcAuto              string          `json:"act_auto"`
	HeartBeat            *HeartBeatConf  `json:"heartbeat"`
	Stats                *StatsConf      `json:"stats"`
	HttpApi              *HttpApiConf    `json:"http_api"`
	HttpServer           *HttpServerConf `json:"http_server"`
	Security             *SecurityConf   `json:"security"`
	Dvr                  *DvrConf        `json:"dvr"`
	HttpStatic           *HttpStaticConf `json:"http_static"`
	HttpRemux            *HttpRemuxConf  `json:"http_remux"`
	Hls                  *HlsConf        `json:"hls"`
	HttpHooks            *HttpHooksConf  `json:"http_hooks"`
	Publish              *PublishConf    `json:"publish"`
}

func (this *VHostConf) amendDefault() {
	this.MinLatency = "on"
	if this.GopCache == "" {
		this.GopCache = "on"
	}

	if this.QueueLength == 0 {
		this.QueueLength = 10
	}

	if this.ReduceSequenceHeader == "" {
		this.ReduceSequenceHeader = "off"
	}

	if this.Publish1stPktTimeout == 0 {
		this.Publish1stPktTimeout = 20000
	}

	if this.PublishNormalTimeout == 0 {
		this.PublishNormalTimeout = 7000
	}

	if this.ChunkSize == 0 {
		this.ChunkSize = 65000
	}

	if this.TimerJitter == "" {
		this.TimerJitter = "full"
	}

	if this.MixCorrect == "" {
		this.MixCorrect = "off"
	}

	if this.Atc == "" {
		this.Atc = "off"
	}

	if this.AtcAuto == "" {
		this.AtcAuto = "on"
	}

	this.HeartBeat.amendDefault()
	this.Stats.amendDefault()
	this.HttpApi.amendDefault()
	this.HttpServer.amendDefault()
	this.Security.amendDefault()
	this.Dvr.amendDefault()
	this.HttpStatic.amendDefault()
	this.HttpRemux.amendDefault()
	this.Hls.amendDefault()
	this.HttpHooks.amendDefault()
	this.Publish.amendDefault()
}

var config *SrsConfig

func GetInstance() *SrsConfig {
	if config == nil {
		config = &SrsConfig{}
	}
	return config
}

func (this *SrsConfig) Init(file string) error {
	this.ListenPort = 1935
	this.Pid = "./srs.pid"
	this.ChunkSize = 60000
	this.MaxConnections = 1000
	this.WorkDir = "./"
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, this)
	if err != nil {
		return err
	}

	this.amendDefault()
	return nil
}

func init() {
	config = &SrsConfig{}
}
