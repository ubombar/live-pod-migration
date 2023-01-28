package clients

import "time"

const (
	ClientPodman = "podman"
)

type Client interface {
	// Get runtime eg 'docker'
	Runtime() string

	// Get version of the runtime
	Version() string

	// Checkpointing functions (demo specific)
	InspectContainer(containerId string) (*ContainerInspectResult, error)
	CheckpointContainer(containerId string, checkpointPath string) error
	RestoreContainer(checkpointArchive string, randomizeName bool) (*ContainerInspectResult, error)
	ClearContainer(containerId string)
}

type ContainerInspectResult struct {
	ID      string    `json:"Id"`
	Created time.Time `json:"Created"`
	Path    string    `json:"Path"`
	Args    []string  `json:"Args"`
	State   struct {
		OciVersion string    `json:"OciVersion"`
		Status     string    `json:"Status"`
		Running    bool      `json:"Running"`
		Paused     bool      `json:"Paused"`
		Restarting bool      `json:"Restarting"`
		OOMKilled  bool      `json:"OOMKilled"`
		Dead       bool      `json:"Dead"`
		Pid        int       `json:"Pid"`
		ConmonPid  int       `json:"ConmonPid"`
		ExitCode   int       `json:"ExitCode"`
		Error      string    `json:"Error"`
		StartedAt  time.Time `json:"StartedAt"`
		FinishedAt time.Time `json:"FinishedAt"`
		Health     struct {
			Status        string      `json:"Status"`
			FailingStreak int         `json:"FailingStreak"`
			Log           interface{} `json:"Log"`
		} `json:"Health"`
		CgroupPath     string    `json:"CgroupPath"`
		CheckpointedAt time.Time `json:"CheckpointedAt"`
		RestoredAt     time.Time `json:"RestoredAt"`
	} `json:"State"`
	Image           string        `json:"Image"`
	ImageDigest     string        `json:"ImageDigest"`
	ImageName       string        `json:"ImageName"`
	Rootfs          string        `json:"Rootfs"`
	Pod             string        `json:"Pod"`
	ResolvConfPath  string        `json:"ResolvConfPath"`
	HostnamePath    string        `json:"HostnamePath"`
	HostsPath       string        `json:"HostsPath"`
	StaticDir       string        `json:"StaticDir"`
	OCIConfigPath   string        `json:"OCIConfigPath"`
	OCIRuntime      string        `json:"OCIRuntime"`
	ConmonPidFile   string        `json:"ConmonPidFile"`
	PidFile         string        `json:"PidFile"`
	Name            string        `json:"Name"`
	RestartCount    int           `json:"RestartCount"`
	Driver          string        `json:"Driver"`
	MountLabel      string        `json:"MountLabel"`
	ProcessLabel    string        `json:"ProcessLabel"`
	AppArmorProfile string        `json:"AppArmorProfile"`
	EffectiveCaps   []string      `json:"EffectiveCaps"`
	BoundingCaps    []string      `json:"BoundingCaps"`
	ExecIDs         []interface{} `json:"ExecIDs"`
	GraphDriver     struct {
		Name string `json:"Name"`
		Data struct {
			LowerDir  string `json:"LowerDir"`
			MergedDir string `json:"MergedDir"`
			UpperDir  string `json:"UpperDir"`
			WorkDir   string `json:"WorkDir"`
		} `json:"Data"`
	} `json:"GraphDriver"`
	Mounts          []interface{} `json:"Mounts"`
	Dependencies    []interface{} `json:"Dependencies"`
	NetworkSettings struct {
		EndpointID             string `json:"EndpointID"`
		Gateway                string `json:"Gateway"`
		IPAddress              string `json:"IPAddress"`
		IPPrefixLen            int    `json:"IPPrefixLen"`
		IPv6Gateway            string `json:"IPv6Gateway"`
		GlobalIPv6Address      string `json:"GlobalIPv6Address"`
		GlobalIPv6PrefixLen    int    `json:"GlobalIPv6PrefixLen"`
		MacAddress             string `json:"MacAddress"`
		Bridge                 string `json:"Bridge"`
		SandboxID              string `json:"SandboxID"`
		HairpinMode            bool   `json:"HairpinMode"`
		LinkLocalIPv6Address   string `json:"LinkLocalIPv6Address"`
		LinkLocalIPv6PrefixLen int    `json:"LinkLocalIPv6PrefixLen"`
		Ports                  struct {
		} `json:"Ports"`
		SandboxKey string `json:"SandboxKey"`
	} `json:"NetworkSettings"`
	Namespace string `json:"Namespace"`
	IsInfra   bool   `json:"IsInfra"`
	IsService bool   `json:"IsService"`
	Config    struct {
		Hostname     string      `json:"Hostname"`
		Domainname   string      `json:"Domainname"`
		User         string      `json:"User"`
		AttachStdin  bool        `json:"AttachStdin"`
		AttachStdout bool        `json:"AttachStdout"`
		AttachStderr bool        `json:"AttachStderr"`
		Tty          bool        `json:"Tty"`
		OpenStdin    bool        `json:"OpenStdin"`
		StdinOnce    bool        `json:"StdinOnce"`
		Env          []string    `json:"Env"`
		Cmd          []string    `json:"Cmd"`
		Image        string      `json:"Image"`
		Volumes      interface{} `json:"Volumes"`
		WorkingDir   string      `json:"WorkingDir"`
		Entrypoint   string      `json:"Entrypoint"`
		OnBuild      interface{} `json:"OnBuild"`
		Labels       interface{} `json:"Labels"`
		Annotations  struct {
			IoContainerManager               string    `json:"io.container.manager"`
			IoKubernetesCriOCreated          time.Time `json:"io.kubernetes.cri-o.Created"`
			IoKubernetesCriOTTY              string    `json:"io.kubernetes.cri-o.TTY"`
			IoPodmanAnnotationsAutoremove    string    `json:"io.podman.annotations.autoremove"`
			IoPodmanAnnotationsInit          string    `json:"io.podman.annotations.init"`
			IoPodmanAnnotationsPrivileged    string    `json:"io.podman.annotations.privileged"`
			IoPodmanAnnotationsPublishAll    string    `json:"io.podman.annotations.publish-all"`
			OrgOpencontainersImageStopSignal string    `json:"org.opencontainers.image.stopSignal"`
		} `json:"Annotations"`
		StopSignal                 int      `json:"StopSignal"`
		HealthcheckOnFailureAction string   `json:"HealthcheckOnFailureAction"`
		CreateCommand              []string `json:"CreateCommand"`
		Umask                      string   `json:"Umask"`
		Timeout                    int      `json:"Timeout"`
		StopTimeout                int      `json:"StopTimeout"`
		Passwd                     bool     `json:"Passwd"`
		SdNotifyMode               string   `json:"sdNotifyMode"`
	} `json:"Config"`
	HostConfig struct {
		Binds           []interface{} `json:"Binds"`
		CgroupManager   string        `json:"CgroupManager"`
		CgroupMode      string        `json:"CgroupMode"`
		ContainerIDFile string        `json:"ContainerIDFile"`
		LogConfig       struct {
			Type   string      `json:"Type"`
			Config interface{} `json:"Config"`
			Path   string      `json:"Path"`
			Tag    string      `json:"Tag"`
			Size   string      `json:"Size"`
		} `json:"LogConfig"`
		NetworkMode  string `json:"NetworkMode"`
		PortBindings struct {
		} `json:"PortBindings"`
		RestartPolicy struct {
			Name              string `json:"Name"`
			MaximumRetryCount int    `json:"MaximumRetryCount"`
		} `json:"RestartPolicy"`
		AutoRemove      bool          `json:"AutoRemove"`
		VolumeDriver    string        `json:"VolumeDriver"`
		VolumesFrom     interface{}   `json:"VolumesFrom"`
		CapAdd          []interface{} `json:"CapAdd"`
		CapDrop         []string      `json:"CapDrop"`
		DNS             []interface{} `json:"Dns"`
		DNSOptions      []interface{} `json:"DnsOptions"`
		DNSSearch       []interface{} `json:"DnsSearch"`
		ExtraHosts      []interface{} `json:"ExtraHosts"`
		GroupAdd        []interface{} `json:"GroupAdd"`
		IpcMode         string        `json:"IpcMode"`
		Cgroup          string        `json:"Cgroup"`
		Cgroups         string        `json:"Cgroups"`
		Links           interface{}   `json:"Links"`
		OomScoreAdj     int           `json:"OomScoreAdj"`
		PidMode         string        `json:"PidMode"`
		Privileged      bool          `json:"Privileged"`
		PublishAllPorts bool          `json:"PublishAllPorts"`
		ReadonlyRootfs  bool          `json:"ReadonlyRootfs"`
		SecurityOpt     []interface{} `json:"SecurityOpt"`
		Tmpfs           struct {
		} `json:"Tmpfs"`
		UTSMode              string        `json:"UTSMode"`
		UsernsMode           string        `json:"UsernsMode"`
		ShmSize              int           `json:"ShmSize"`
		Runtime              string        `json:"Runtime"`
		ConsoleSize          []int         `json:"ConsoleSize"`
		Isolation            string        `json:"Isolation"`
		CPUShares            int           `json:"CpuShares"`
		Memory               int           `json:"Memory"`
		NanoCpus             int           `json:"NanoCpus"`
		CgroupParent         string        `json:"CgroupParent"`
		BlkioWeight          int           `json:"BlkioWeight"`
		BlkioWeightDevice    interface{}   `json:"BlkioWeightDevice"`
		BlkioDeviceReadBps   interface{}   `json:"BlkioDeviceReadBps"`
		BlkioDeviceWriteBps  interface{}   `json:"BlkioDeviceWriteBps"`
		BlkioDeviceReadIOps  interface{}   `json:"BlkioDeviceReadIOps"`
		BlkioDeviceWriteIOps interface{}   `json:"BlkioDeviceWriteIOps"`
		CPUPeriod            int           `json:"CpuPeriod"`
		CPUQuota             int           `json:"CpuQuota"`
		CPURealtimePeriod    int           `json:"CpuRealtimePeriod"`
		CPURealtimeRuntime   int           `json:"CpuRealtimeRuntime"`
		CpusetCpus           string        `json:"CpusetCpus"`
		CpusetMems           string        `json:"CpusetMems"`
		Devices              []interface{} `json:"Devices"`
		DiskQuota            int           `json:"DiskQuota"`
		KernelMemory         int           `json:"KernelMemory"`
		MemoryReservation    int           `json:"MemoryReservation"`
		MemorySwap           int           `json:"MemorySwap"`
		MemorySwappiness     int           `json:"MemorySwappiness"`
		OomKillDisable       bool          `json:"OomKillDisable"`
		PidsLimit            int           `json:"PidsLimit"`
		Ulimits              []interface{} `json:"Ulimits"`
		CPUCount             int           `json:"CpuCount"`
		CPUPercent           int           `json:"CpuPercent"`
		IOMaximumIOps        int           `json:"IOMaximumIOps"`
		IOMaximumBandwidth   int           `json:"IOMaximumBandwidth"`
		CgroupConf           interface{}   `json:"CgroupConf"`
	} `json:"HostConfig"`
}
