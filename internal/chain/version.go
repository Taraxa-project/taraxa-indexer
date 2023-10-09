package chain

import (
	"strings"

	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	log "github.com/sirupsen/logrus"
)

const MinimumProtocolVersion = "1.5.0"

func versionStringToInts(version string) (int64, int64, int64) {
	v := strings.Split(version, ".")
	if len(v) != 3 {
		log.WithFields(log.Fields{"version": version}).Error("Invalid version string")
	}
	return common.ParseInt(v[0]), common.ParseInt(v[1]), common.ParseInt(v[2])
}

func CheckProtocolVersion(version string) bool {
	supportedMajor, supportedMinor, supportedPatch := versionStringToInts(MinimumProtocolVersion)
	major, minor, patch := versionStringToInts(version)
	if major < supportedMajor {
		return false
	} else if (major == supportedMajor) && minor < supportedMinor {
		return false
	} else if (major == supportedMajor) && (minor == supportedMinor) && patch < supportedPatch {
		return false
	}
	return true
}
