// Code generated by "stringer -type=flapsAction"; DO NOT EDIT.

package flaps

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[none-0]
	_ = x[appCreate-1]
	_ = x[machineLaunch-2]
	_ = x[machineUpdate-3]
	_ = x[machineStart-4]
	_ = x[machineWait-5]
	_ = x[machineStop-6]
	_ = x[machineRestart-7]
	_ = x[machineGet-8]
	_ = x[machineList-9]
	_ = x[machineDestroy-10]
	_ = x[machineKill-11]
	_ = x[machineFindLease-12]
	_ = x[machineAcquireLease-13]
	_ = x[machineRefreshLease-14]
	_ = x[machineReleaseLease-15]
	_ = x[machineExec-16]
	_ = x[machinePs-17]
	_ = x[machineCordon-18]
	_ = x[machineUncordon-19]
	_ = x[volumeList-20]
	_ = x[volumeCreate-21]
	_ = x[volumetUpdate-22]
	_ = x[volumeGet-23]
	_ = x[volumeSnapshotCreate-24]
	_ = x[volumeSnapshotList-25]
	_ = x[volumeExtend-26]
	_ = x[volumeDelete-27]
}

const _flapsAction_name = "noneappCreatemachineLaunchmachineUpdatemachineStartmachineWaitmachineStopmachineRestartmachineGetmachineListmachineDestroymachineKillmachineFindLeasemachineAcquireLeasemachineRefreshLeasemachineReleaseLeasemachineExecmachinePsmachineCordonmachineUncordonvolumeListvolumeCreatevolumetUpdatevolumeGetvolumeSnapshotCreatevolumeSnapshotListvolumeExtendvolumeDelete"

var _flapsAction_index = [...]uint16{0, 4, 13, 26, 39, 51, 62, 73, 87, 97, 108, 122, 133, 149, 168, 187, 206, 217, 226, 239, 254, 264, 276, 289, 298, 318, 336, 348, 360}

func (i flapsAction) String() string {
	if i < 0 || i >= flapsAction(len(_flapsAction_index)-1) {
		return "flapsAction(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _flapsAction_name[_flapsAction_index[i]:_flapsAction_index[i+1]]
}