package ebpf

// naming:
// remove bpf_ prefix
// remove get_
// camelcase and remove _

// Printk ...
func Printk(format string, v ...interface{}) int {
	return 0
}

// CurrentPidTgid is the bpf_get_current_pid_tgid helper function.
func CurrentPidTgid() int64 {
	return 0
}
