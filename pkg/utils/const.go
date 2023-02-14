package utils

const (
	ResourceName = "yoyoworks.com/xpu-shares"
	CountName    = "yoyoworks.com/xpu-counts"

	EnvNVGPU              = "NVIDIA_VISIBLE_DEVICES"
	EnvResourceIndex      = "YOYOWORKS_XPU_SHARES_INDEX"
	EnvResourceByPod      = "YOYOWORKS_XPU_SHARES_POD"
	EnvResourceByDev      = "YOYOWORKS_XPU_SHARES_TOTAL"
	EnvAssignedFlag       = "YOYOWORKS_XPU_SHARES_ALLOCATED"
	EnvResourceAssumeTime = "YOYOWORKS_XPU_SHARES_FILTER_STAMP"
)

