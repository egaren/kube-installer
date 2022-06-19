package k3s

const (
	systemDirectory = "/etc/systemd/system"
	binDir          = "/usr/local/bin"
	binaryName      = "k3s"
	shaSumName      = "sha256sum"
	downloadPath    = "/tmp"
	configDir       = "/etc/rancher/k3s"
	configYamlDir   = "/etc/rancher/k3s/config.yaml.d"
	configFile      = "/etc/rancher/k3s/config.yaml"
	sourceUrl       = "https://github.com/k3s-io/k3s/releases/download"
	executableMode  = 0755
	criCtl          = "crictl"
	ctr             = "ctr"
	kubectl         = "kubectl"
)

const svc = `[Unit]
Description=Lightweight Kubernetes
Documentation=https://k3s.io
Wants=network-online.target
After=network-online.target

[Service]
Type=notify
ExecStartPre=-/sbin/modprobe br_netfilter
ExecStartPre=-/sbin/modprobe overlay
ExecStart=/usr/local/bin/k3s server --config /etc/rancher/k3s/config.yaml --advertise-address {{ .NodeAddr }} --node-ip {{ .NodeAddr }} --node-external-ip {{ .NodeAddr }} --disable local-storage
KillMode=process
Delegate=yes
LimitNOFILE=1048576
LimitNPROC=infinity
LimitCORE=infinity
TasksMax=infinity
TimeoutStartSec=0
Restart=always
RestartSec=5s

[Install]
WantedBy=multi-user.target
`
