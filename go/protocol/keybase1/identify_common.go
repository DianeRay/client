// Auto-generated by avdl-compiler v1.3.13 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/identify_common.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
)

type TrackToken string
type TrackDiffType int

const (
	TrackDiffType_NONE               TrackDiffType = 0
	TrackDiffType_ERROR              TrackDiffType = 1
	TrackDiffType_CLASH              TrackDiffType = 2
	TrackDiffType_REVOKED            TrackDiffType = 3
	TrackDiffType_UPGRADED           TrackDiffType = 4
	TrackDiffType_NEW                TrackDiffType = 5
	TrackDiffType_REMOTE_FAIL        TrackDiffType = 6
	TrackDiffType_REMOTE_WORKING     TrackDiffType = 7
	TrackDiffType_REMOTE_CHANGED     TrackDiffType = 8
	TrackDiffType_NEW_ELDEST         TrackDiffType = 9
	TrackDiffType_NONE_VIA_TEMPORARY TrackDiffType = 10
)

var TrackDiffTypeMap = map[string]TrackDiffType{
	"NONE":               0,
	"ERROR":              1,
	"CLASH":              2,
	"REVOKED":            3,
	"UPGRADED":           4,
	"NEW":                5,
	"REMOTE_FAIL":        6,
	"REMOTE_WORKING":     7,
	"REMOTE_CHANGED":     8,
	"NEW_ELDEST":         9,
	"NONE_VIA_TEMPORARY": 10,
}

var TrackDiffTypeRevMap = map[TrackDiffType]string{
	0:  "NONE",
	1:  "ERROR",
	2:  "CLASH",
	3:  "REVOKED",
	4:  "UPGRADED",
	5:  "NEW",
	6:  "REMOTE_FAIL",
	7:  "REMOTE_WORKING",
	8:  "REMOTE_CHANGED",
	9:  "NEW_ELDEST",
	10: "NONE_VIA_TEMPORARY",
}

func (e TrackDiffType) String() string {
	if v, ok := TrackDiffTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type TrackDiff struct {
	Type          TrackDiffType `codec:"type" json:"type"`
	DisplayMarkup string        `codec:"displayMarkup" json:"displayMarkup"`
}

type TrackSummary struct {
	Username string `codec:"username" json:"username"`
	Time     Time   `codec:"time" json:"time"`
	IsRemote bool   `codec:"isRemote" json:"isRemote"`
}

// TrackStatus is a summary of this track before the track is approved by the
// user.
// NEW_*: New tracks
// UPDATE_*: Update to an existing track
// NEW_OK: Everything ok
// NEW_ZERO_PROOFS: User being tracked has no proofs
// NEW_FAIL_PROOFS: User being tracked has some failed proofs
// UPDATE_BROKEN: Previous tracking statement broken, this one will fix it.
// UPDATE_NEW_PROOFS: Previous tracking statement ok, but there are new proofs since previous tracking statement generated
// UPDATE_OK: No changes to previous tracking statement
type TrackStatus int

const (
	TrackStatus_NEW_OK                      TrackStatus = 1
	TrackStatus_NEW_ZERO_PROOFS             TrackStatus = 2
	TrackStatus_NEW_FAIL_PROOFS             TrackStatus = 3
	TrackStatus_UPDATE_BROKEN_FAILED_PROOFS TrackStatus = 4
	TrackStatus_UPDATE_NEW_PROOFS           TrackStatus = 5
	TrackStatus_UPDATE_OK                   TrackStatus = 6
	TrackStatus_UPDATE_BROKEN_REVOKED       TrackStatus = 7
)

var TrackStatusMap = map[string]TrackStatus{
	"NEW_OK":                      1,
	"NEW_ZERO_PROOFS":             2,
	"NEW_FAIL_PROOFS":             3,
	"UPDATE_BROKEN_FAILED_PROOFS": 4,
	"UPDATE_NEW_PROOFS":           5,
	"UPDATE_OK":                   6,
	"UPDATE_BROKEN_REVOKED":       7,
}

var TrackStatusRevMap = map[TrackStatus]string{
	1: "NEW_OK",
	2: "NEW_ZERO_PROOFS",
	3: "NEW_FAIL_PROOFS",
	4: "UPDATE_BROKEN_FAILED_PROOFS",
	5: "UPDATE_NEW_PROOFS",
	6: "UPDATE_OK",
	7: "UPDATE_BROKEN_REVOKED",
}

func (e TrackStatus) String() string {
	if v, ok := TrackStatusRevMap[e]; ok {
		return v
	}
	return ""
}

type TrackOptions struct {
	LocalOnly     bool `codec:"localOnly" json:"localOnly"`
	BypassConfirm bool `codec:"bypassConfirm" json:"bypassConfirm"`
	ForceRetrack  bool `codec:"forceRetrack" json:"forceRetrack"`
	ExpiringLocal bool `codec:"expiringLocal" json:"expiringLocal"`
	ForPGPPull    bool `codec:"forPGPPull" json:"forPGPPull"`
}

type IdentifyReasonType int

const (
	IdentifyReasonType_NONE       IdentifyReasonType = 0
	IdentifyReasonType_ID         IdentifyReasonType = 1
	IdentifyReasonType_TRACK      IdentifyReasonType = 2
	IdentifyReasonType_ENCRYPT    IdentifyReasonType = 3
	IdentifyReasonType_DECRYPT    IdentifyReasonType = 4
	IdentifyReasonType_VERIFY     IdentifyReasonType = 5
	IdentifyReasonType_RESOURCE   IdentifyReasonType = 6
	IdentifyReasonType_BACKGROUND IdentifyReasonType = 7
)

var IdentifyReasonTypeMap = map[string]IdentifyReasonType{
	"NONE":       0,
	"ID":         1,
	"TRACK":      2,
	"ENCRYPT":    3,
	"DECRYPT":    4,
	"VERIFY":     5,
	"RESOURCE":   6,
	"BACKGROUND": 7,
}

var IdentifyReasonTypeRevMap = map[IdentifyReasonType]string{
	0: "NONE",
	1: "ID",
	2: "TRACK",
	3: "ENCRYPT",
	4: "DECRYPT",
	5: "VERIFY",
	6: "RESOURCE",
	7: "BACKGROUND",
}

func (e IdentifyReasonType) String() string {
	if v, ok := IdentifyReasonTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type IdentifyReason struct {
	Type     IdentifyReasonType `codec:"type" json:"type"`
	Reason   string             `codec:"reason" json:"reason"`
	Resource string             `codec:"resource" json:"resource"`
}

type IdentifyOutcome struct {
	Username          string         `codec:"username" json:"username"`
	Status            *Status        `codec:"status,omitempty" json:"status,omitempty"`
	Warnings          []string       `codec:"warnings" json:"warnings"`
	TrackUsed         *TrackSummary  `codec:"trackUsed,omitempty" json:"trackUsed,omitempty"`
	TrackStatus       TrackStatus    `codec:"trackStatus" json:"trackStatus"`
	NumTrackFailures  int            `codec:"numTrackFailures" json:"numTrackFailures"`
	NumTrackChanges   int            `codec:"numTrackChanges" json:"numTrackChanges"`
	NumProofFailures  int            `codec:"numProofFailures" json:"numProofFailures"`
	NumRevoked        int            `codec:"numRevoked" json:"numRevoked"`
	NumProofSuccesses int            `codec:"numProofSuccesses" json:"numProofSuccesses"`
	Revoked           []TrackDiff    `codec:"revoked" json:"revoked"`
	TrackOptions      TrackOptions   `codec:"trackOptions" json:"trackOptions"`
	ForPGPPull        bool           `codec:"forPGPPull" json:"forPGPPull"`
	Reason            IdentifyReason `codec:"reason" json:"reason"`
}

type IdentifyRes struct {
	User       *User           `codec:"user,omitempty" json:"user,omitempty"`
	PublicKeys []PublicKey     `codec:"publicKeys" json:"publicKeys"`
	Outcome    IdentifyOutcome `codec:"outcome" json:"outcome"`
	TrackToken TrackToken      `codec:"trackToken" json:"trackToken"`
}

type RemoteProof struct {
	ProofType     ProofType `codec:"proofType" json:"proofType"`
	Key           string    `codec:"key" json:"key"`
	Value         string    `codec:"value" json:"value"`
	DisplayMarkup string    `codec:"displayMarkup" json:"displayMarkup"`
	SigID         SigID     `codec:"sigID" json:"sigID"`
	MTime         Time      `codec:"mTime" json:"mTime"`
}

type IdentifyCommonInterface interface {
}

func IdentifyCommonProtocol(i IdentifyCommonInterface) rpc.Protocol {
	return rpc.Protocol{
		Name:    "keybase.1.identifyCommon",
		Methods: map[string]rpc.ServeHandlerDescription{},
	}
}

type IdentifyCommonClient struct {
	Cli rpc.GenericClient
}
