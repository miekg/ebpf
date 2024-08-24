//go:build ignore

// This command will generate the Go types and functions from 'bpftool btf dump -j file /sys/kernel/btf/vmlinux '
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// Some values are so huge that they don't fit in a int, currently we skip them as they are mostly(?) bitmasks.
type Val int

type BTF struct {
	ID         int
	TypeID     int `json:"type_id,omitempty"`
	RetTypeID  int `json:"ret_type_id"` // return type for functions
	Kind       string
	Name       string
	Size       int    `json:",omitempty"`
	BitsOffset int    `json:"bits_offset"`
	NrBits     int    `json:"nr_bits"`
	NrElems    int    `json:"nr_elems"`
	Encoding   string `json:",omitempty"`
	Vlen       int    `json:",omitempty"` // Number of Members or Params
	Val        Val    `json:",omitempty"`
	Members    []BTF  `json:",omitempty"`
	Params     []BTF  `json:",omitempty"`
	Values     []BTF  `json:",omitempty"`
}

type Kind string

const (
	EnumKind   = "ENUM"
	StructKind = "STRUCT"
)

func (v *Val) UnmarshalJSON(b []byte) error {
	var i int
	if err := json.Unmarshal(b, &i); err != nil {
		// print ?
		*v = 0
		return nil
	}
	*v = Val(i)
	return nil
}

func main() {
	data, err := os.ReadFile("vmlinux.json")
	if err != nil {
		log.Fatal(err)
	}
	linux := map[string][]BTF{}
	if err := json.Unmarshal(data, &linux); err != nil {
		log.Fatal(err)
	}

	for _, k := range linux["types"] {
		if k.Kind == EnumKind && k.Name == "bpf_map_type" {
			fmt.Printf("const (\n")
			for _, v := range k.Values {
				fmt.Printf("  %s = %d\n", makeGoName(v.Name), v.Val)
			}
			fmt.Printf(")\n")
		}
	}

}

func makeGoName(c string) string {
	fields := strings.SplitAfter(strings.ToLower(c), "_")
	start := 0
	if fields[0] == "bpf_" {
		start = 1
	}

	s := strings.Builder{}
	for _, f := range fields[start:] {
		f = strings.Trim(f, "_")
		s.WriteString(strings.Title(f))
	}
	return s.String()
}
