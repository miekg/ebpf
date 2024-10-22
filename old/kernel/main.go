package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
)

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
	Members    []BTF  `json:",omitempty"`
	Params     []BTF  `json:",omitempty"`
}

type Kind string

const (
	StructKind = "STRUCT"
)

type Lookup map[int]BTF

func main() {
	stdin := bufio.NewReader(os.Stdin)
	data, err := io.ReadAll(stdin)
	if err != nil {
		log.Fatal(err)
	}

	j := map[string][]BTF{}
	if err := json.Unmarshal(data, &j); err != nil {
		log.Fatal(err)
	}

	l := Lookup{}
	for _, k := range j["types"] {
		l[k.ID] = k
	}
	// 90 = "name": "callback_head", (on my machine, now)
	resolve(l, 90)
}

// What to return? Something we can convert into Go code... So reflect stuff? Or a fuller BTF struct?

// resolve resolves the element with typeIP from l. It returns a fleshed out BTF where the type ID is followed
// and a set to 0 and replaced with its target BTF.
func resolve(l Lookup, typeID int) *BTF {
	b := l[typeID]

}
