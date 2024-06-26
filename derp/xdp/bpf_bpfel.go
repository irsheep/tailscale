// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64 || arm || arm64 || loong64 || mips64le || mipsle || ppc64le || riscv64

package xdp

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpfConfig struct{ DstPort uint16 }

type bpfCounterKeyAf uint32

const (
	bpfCounterKeyAfCOUNTER_KEY_AF_UNKNOWN bpfCounterKeyAf = 0
	bpfCounterKeyAfCOUNTER_KEY_AF_IPV4    bpfCounterKeyAf = 1
	bpfCounterKeyAfCOUNTER_KEY_AF_IPV6    bpfCounterKeyAf = 2
	bpfCounterKeyAfCOUNTER_KEY_AF_LEN     bpfCounterKeyAf = 3
)

type bpfCounterKeyPacketsBytesAction uint32

const (
	bpfCounterKeyPacketsBytesActionCOUNTER_KEY_PACKETS_PASS_TOTAL       bpfCounterKeyPacketsBytesAction = 0
	bpfCounterKeyPacketsBytesActionCOUNTER_KEY_BYTES_PASS_TOTAL         bpfCounterKeyPacketsBytesAction = 1
	bpfCounterKeyPacketsBytesActionCOUNTER_KEY_PACKETS_ABORTED_TOTAL    bpfCounterKeyPacketsBytesAction = 2
	bpfCounterKeyPacketsBytesActionCOUNTER_KEY_BYTES_ABORTED_TOTAL      bpfCounterKeyPacketsBytesAction = 3
	bpfCounterKeyPacketsBytesActionCOUNTER_KEY_PACKETS_TX_TOTAL         bpfCounterKeyPacketsBytesAction = 4
	bpfCounterKeyPacketsBytesActionCOUNTER_KEY_BYTES_TX_TOTAL           bpfCounterKeyPacketsBytesAction = 5
	bpfCounterKeyPacketsBytesActionCOUNTER_KEY_PACKETS_DROP_TOTAL       bpfCounterKeyPacketsBytesAction = 6
	bpfCounterKeyPacketsBytesActionCOUNTER_KEY_BYTES_DROP_TOTAL         bpfCounterKeyPacketsBytesAction = 7
	bpfCounterKeyPacketsBytesActionCOUNTER_KEY_PACKETS_BYTES_ACTION_LEN bpfCounterKeyPacketsBytesAction = 8
)

type bpfCounterKeyProgEnd uint32

const (
	bpfCounterKeyProgEndCOUNTER_KEY_END_UNSPECIFIED                bpfCounterKeyProgEnd = 0
	bpfCounterKeyProgEndCOUNTER_KEY_END_UNEXPECTED_FIRST_STUN_ATTR bpfCounterKeyProgEnd = 1
	bpfCounterKeyProgEndCOUNTER_KEY_END_INVALID_UDP_CSUM           bpfCounterKeyProgEnd = 2
	bpfCounterKeyProgEndCOUNTER_KEY_END_INVALID_IP_CSUM            bpfCounterKeyProgEnd = 3
	bpfCounterKeyProgEndCOUNTER_KEY_END_NOT_STUN_PORT              bpfCounterKeyProgEnd = 4
	bpfCounterKeyProgEndCOUNTER_KEY_END_INVALID_SW_ATTR_VAL        bpfCounterKeyProgEnd = 5
	bpfCounterKeyProgEndCOUNTER_KEY_END_LEN                        bpfCounterKeyProgEnd = 6
)

type bpfCountersKey struct {
	Unused  uint8
	Af      uint8
	Pba     uint8
	ProgEnd uint8
}

// loadBpf returns the embedded CollectionSpec for bpf.
func loadBpf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf: %w", err)
	}

	return spec, err
}

// loadBpfObjects loads bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpfObjects
//	*bpfPrograms
//	*bpfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfSpecs struct {
	bpfProgramSpecs
	bpfMapSpecs
}

// bpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfProgramSpecs struct {
	XdpProgFunc *ebpf.ProgramSpec `ebpf:"xdp_prog_func"`
}

// bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfMapSpecs struct {
	ConfigMap   *ebpf.MapSpec `ebpf:"config_map"`
	CountersMap *ebpf.MapSpec `ebpf:"counters_map"`
}

// bpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfObjects struct {
	bpfPrograms
	bpfMaps
}

func (o *bpfObjects) Close() error {
	return _BpfClose(
		&o.bpfPrograms,
		&o.bpfMaps,
	)
}

// bpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfMaps struct {
	ConfigMap   *ebpf.Map `ebpf:"config_map"`
	CountersMap *ebpf.Map `ebpf:"counters_map"`
}

func (m *bpfMaps) Close() error {
	return _BpfClose(
		m.ConfigMap,
		m.CountersMap,
	)
}

// bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfPrograms struct {
	XdpProgFunc *ebpf.Program `ebpf:"xdp_prog_func"`
}

func (p *bpfPrograms) Close() error {
	return _BpfClose(
		p.XdpProgFunc,
	)
}

func _BpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_bpfel.o
var _BpfBytes []byte
