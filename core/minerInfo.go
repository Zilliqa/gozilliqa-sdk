package core

import "container/list"

type MinerInfoDSComm struct {
	// inner type is string
	DSNodes        *list.List
	DSNodesEjected []string
}
