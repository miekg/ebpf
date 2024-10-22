//go:build ignore

// This command will generate the Go types and functions from 'bpftool btf dump -j file /sys/kernel/btf/vmlinux '
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"log"
	"os"
	"strings"
)

// BTF is the JSON structure that we parse via bpftool.
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
	Encoding   string `json:",omitempty"` // FIXME: Not used yet.
	Vlen       int    `json:",omitempty"` // Number of Members, Params or Values.
	Val        Val    `json:",omitempty"`
	Members    []BTF  `json:",omitempty"`
	Params     []BTF  `json:",omitempty"`
	Values     []BTF  `json:",omitempty"`
}

// Package is the Go code we output. This is mostly variable definitions.
type Package struct {
	Package string // Package controls the package name.
	Sources []Source
}

// Source is the actual (printable) Go code.
type Source struct {
	Kind string        // How to place the code in the package.
	Text *bytes.Buffer // Source is the directly printable Go code
}

const (
	EnumKind   = "ENUM"
	StructKind = "STRUCT"
)

// KindToType maps the C Kinds to Go language constructs.
var KindToType = map[string]string{
	EnumKind:   "const",
	StructKind: "struct",
}

// Some values are so huge that they don't fit in a int, currently we skip them as they are mostly(?) bitmasks.
type Val int

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

	pBpf := Package{Package: "bpf"}

	for _, k := range linux["types"] {
		if k.Kind == EnumKind && k.Name == "bpf_map_type" {
			source := Source{Kind: KindToType[EnumKind], Text: new(bytes.Buffer)}
			for _, v := range k.Values {
				fmt.Fprintf(source.Text, "%s = %d\n", makeGoName(v.Name), v.Val)
			}
			pBpf.Sources = append(pBpf.Sources, source)
		}
	}

	final := new(bytes.Buffer)
	for i, source := range pBpf.Sources {
		if i == 0 {
			fmt.Printf("package %s\n\n", pBpf.Package)

		}
		switch source.Kind {
		case "const":
			fmt.Fprintf(final, "const (\n")
			fmt.Fprintf(final, "%s", source.Text.Bytes())
			fmt.Fprintf(final, ")\n")

		case "struct":
		}
	}

	res, err := format.Source(final.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", res)
}

// makeGoName takes a string like XXX_YYYY and converts it to Go's camelcase XxxYyyy. BPF_ is removed from the beginning
// of the string, as those indentifiers are put in bpf package.
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
