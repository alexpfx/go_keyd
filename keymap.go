package keymap

import (
	"fmt"
	"regexp"
	"strconv"

	"os/exec"
	"strings"
)

var hexa = regexp.MustCompile(`(?:0x)([A-Fa-f0-9]+)`)

func Load() KeyMapper {
	cmdMm := exec.Command("xmodmap", "-pm")
	mm, err := cmdMm.CombinedOutput()
	check(err)

	modKeysHex := hexa.FindAllStringSubmatch(string(mm), -1)
	modKeys := make([]uint16, 0)
	for _, hexKey := range modKeysHex {
		p, _ := strconv.ParseUint(hexKey[1], 16, 16)
		modKeys = append(modKeys, uint16(p))
	}

	keys := make(map[uint16][]string)
	cmd := exec.Command("xmodmap", "-pke")

	output, err := cmd.CombinedOutput()
	check(err)

	lines := strings.Split(string(output), "\n")
	for _, s := range lines {
		if s == "" {
			continue
		}
		key, vals := splitKeyVals(s)

		ikey, err := strconv.Atoi(key)
		if err != nil {
			continue
		}
		ix := uint16(ikey)

		keys[ix] = vals
	}

	return KeyMapper{keys: keys, modKeys: modKeys}

}

func splitKeyVals(s string) (string, []string) {
	spt := strings.Split(s, "=")
	k := spt[0]
	v := spt[1]
	return strings.Fields(k)[1], strings.Fields(v)
}

type KeyMapper struct {
	keys    map[uint16][]string
	modKeys []uint16
}

func (k KeyMapper) Get(keyCode uint16, nth int) string {
	if val, ok := k.keys[keyCode]; ok && len(val) >= nth {
		return val[nth]
	}
	return ""
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
