package disassemble

import (
	"encoding/binary"
	"reflect"
	"testing"
)

func Test_decompose_single_instr(t *testing.T) {
	var results [1024]byte
	type args struct {
		instructionValue uint32
		address          uint64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ummla	v1.4s, v16.16b, v31.16b",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x01, 0xa6, 0x9f, 0x6e}),
				address:          0,
			},
			want:    "ummla\tv1.4s, v16.16b, v31.16b",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decompose(tt.args.address, tt.args.instructionValue, &results)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decompose() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got.Disassembly, tt.want) {
				t.Errorf("Disassembly = %v, want %v", got.Disassembly, tt.want)
			}
		})
	}
}

func Test_decompose_v8_1a(t *testing.T) {
	var results [1024]byte
	type args struct {
		instructionValue uint32
		address          uint64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		//
		// llvm/test/MC/AArch64/armv8.1a-atomic.s
		//
		{
			name: "ldadda	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x00, 0xa0, 0xf8}),
				address:          0,
			},
			want:    "ldadda\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldclrl	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x10, 0x60, 0xf8}),
				address:          0,
			},
			want:    "ldclrl\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldeoral	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x20, 0xe0, 0xf8}),
				address:          0,
			},
			want:    "ldeoral\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldset	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x30, 0x20, 0xf8}),
				address:          0,
			},
			want:    "ldset\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsmaxa	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x40, 0xa0, 0xb8}),
				address:          0,
			},
			want:    "ldsmaxa\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsminlb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x50, 0x60, 0x38}),
				address:          0,
			},
			want:    "ldsminlb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldumaxalh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x60, 0xe0, 0x78}),
				address:          0,
			},
			want:    "ldumaxalh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldumin	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x70, 0x20, 0xb8}),
				address:          0,
			},
			want:    "ldumin\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsminb	w2, w3, [x5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0x50, 0x22, 0x38}),
				address:          0,
			},
			want:    "ldsminb\tw2, w3, [x5]",
			wantErr: false,
		},
		{
			name: "staddlb	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x00, 0x60, 0x38}),
				address:          0,
			},
			want:    "staddlb\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stclrlh	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x10, 0x60, 0x78}),
				address:          0,
			},
			want:    "stclrlh\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "steorl	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x20, 0x60, 0xb8}),
				address:          0,
			},
			want:    "steorl\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stsetl	x0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x30, 0x60, 0xf8}),
				address:          0,
			},
			want:    "stsetl\tx0, [x2]",
			wantErr: false,
		},
		{
			name: "stsmaxb	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x40, 0x20, 0x38}),
				address:          0,
			},
			want:    "stsmaxb\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stsminh	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x50, 0x20, 0x78}),
				address:          0,
			},
			want:    "stsminh\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stumax	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x60, 0x20, 0xb8}),
				address:          0,
			},
			want:    "stumax\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stumin	x0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x70, 0x20, 0xf8}),
				address:          0,
			},
			want:    "stumin\tx0, [x2]",
			wantErr: false,
		},
		{
			name: "stsminl	x29, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x53, 0x7d, 0xf8}),
				address:          0,
			},
			want:    "stsminl\tfp, [sp]",
			wantErr: false,
		},
		{
			name: "swp	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x80, 0x20, 0xf8}),
				address:          0,
			},
			want:    "swp\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "swpb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x80, 0x20, 0x38}),
				address:          0,
			},
			want:    "swpb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "swplh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x80, 0x60, 0x78}),
				address:          0,
			},
			want:    "swplh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "swpal	x0, x1, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe1, 0x83, 0xe0, 0xf8}),
				address:          0,
			},
			want:    "swpal\tx0, x1, [sp]",
			wantErr: false,
		},
		{
			name: "casp	x0, x1, x2, x3, [x4]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x82, 0x7c, 0x20, 0x48}),
				address:          0,
			},
			want:    "casp\tx0, x1, x2, x3, [x4]",
			wantErr: false,
		},
		{
			name: "casp	w0, w1, w2, w3, [x4]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x82, 0x7c, 0x20, 0x08}),
				address:          0,
			},
			want:    "casp\tw0, w1, w2, w3, [x4]",
			wantErr: false,
		},
		//
		// llvm/test/MC/AArch64/armv8.1a-lor.s
		//
		{
			name: "ldlarb	w0, [x1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x7c, 0xdf, 0x08}),
				address:          0,
			},
			want:    "ldlarb\tw0, [x1]",
			wantErr: false,
		},
		{
			name: "ldlarh	w0, [x1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x7c, 0xdf, 0x48}),
				address:          0,
			},
			want:    "ldlarh\tw0, [x1]",
			wantErr: false,
		},
		{
			name: "ldlar	w0, [x1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x7c, 0xdf, 0x88}),
				address:          0,
			},
			want:    "ldlar\tw0, [x1]",
			wantErr: false,
		},
		{
			name: "ldlar	x0, [x1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x7c, 0xdf, 0xc8}),
				address:          0,
			},
			want:    "ldlar\tx0, [x1]",
			wantErr: false,
		},
		{
			name: "stllrb	w0, [x1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x7c, 0x9f, 0x08}),
				address:          0,
			},
			want:    "stllrb\tw0, [x1]",
			wantErr: false,
		},
		{
			name: "stllrh	w0, [x1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x7c, 0x9f, 0x48}),
				address:          0,
			},
			want:    "stllrh\tw0, [x1]",
			wantErr: false,
		},
		{
			name: "stllr	w0, [x1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x7c, 0x9f, 0x88}),
				address:          0,
			},
			want:    "stllr\tw0, [x1]",
			wantErr: false,
		},
		{
			name: "stllr	x0, [x1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x7c, 0x9f, 0xc8}),
				address:          0,
			},
			want:    "stllr\tx0, [x1]",
			wantErr: false,
		},
		{
			name: "msr	lorsa_el1, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0xa4, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tlorsa_el1, x0",
			wantErr: false,
		},
		{
			name: "msr	lorea_el1, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xa4, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tlorea_el1, x0",
			wantErr: false,
		},
		{
			name: "msr	lorn_el1, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0xa4, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tlorn_el1, x0",
			wantErr: false,
		},
		{
			name: "msr	lorc_el1, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x60, 0xa4, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tlorc_el1, x0",
			wantErr: false,
		},
		{
			name: "mrs	x0, lorid_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0xa4, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, lorid_el1",
			wantErr: false,
		},
		//
		// llvm/test/MC/AArch64/armv8.1a-pan.s
		//
		{
			name: "msr	pan, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x40, 0x00, 0xd5}),
				address:          0,
			},
			want:    "msr\tpan, #0",
			wantErr: false,
		},
		{
			name: "msr	pan, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x41, 0x00, 0xd5}),
				address:          0,
			},
			want:    "msr\tpan, #0x1",
			wantErr: false,
		},
		{
			name: "msr	pan, x5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x65, 0x42, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tpan, x5",
			wantErr: false,
		},
		{
			name: "mrs	x13, pan",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6d, 0x42, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx13, pan",
			wantErr: false,
		},
		//
		// llvm/test/MC/AArch64/armv8.1a-rdma.s
		//
		{
			name: "sqrdmlah	v0.4h, v1.4h, v2.4h",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x84, 0x42, 0x2e}),
				address:          0,
			},
			want:    "sqrdmlah\tv0.4h, v1.4h, v2.4h",
			wantErr: false,
		},
		{
			name: "sqrdmlsh	v0.4h, v1.4h, v2.4h",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x8c, 0x42, 0x2e}),
				address:          0,
			},
			want:    "sqrdmlsh\tv0.4h, v1.4h, v2.4h",
			wantErr: false,
		},
		{
			name: "sqrdmlah	v0.2s, v1.2s, v2.2s",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x84, 0x82, 0x2e}),
				address:          0,
			},
			want:    "sqrdmlah\tv0.2s, v1.2s, v2.2s",
			wantErr: false,
		},
		{
			name: "sqrdmlsh	v0.2s, v1.2s, v2.2s",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x8c, 0x82, 0x2e}),
				address:          0,
			},
			want:    "sqrdmlsh\tv0.2s, v1.2s, v2.2s",
			wantErr: false,
		},
		{
			name: "sqrdmlah	v0.4s, v1.4s, v2.4s",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x84, 0x82, 0x6e}),
				address:          0,
			},
			want:    "sqrdmlah\tv0.4s, v1.4s, v2.4s",
			wantErr: false,
		},
		{
			name: "sqrdmlsh	v0.4s, v1.4s, v2.4s",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x8c, 0x82, 0x6e}),
				address:          0,
			},
			want:    "sqrdmlsh\tv0.4s, v1.4s, v2.4s",
			wantErr: false,
		},
		{
			name: "sqrdmlah	v0.8h, v1.8h, v2.8h",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x84, 0x42, 0x6e}),
				address:          0,
			},
			want:    "sqrdmlah\tv0.8h, v1.8h, v2.8h",
			wantErr: false,
		},
		{
			name: "sqrdmlsh	v0.8h, v1.8h, v2.8h",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x8c, 0x42, 0x6e}),
				address:          0,
			},
			want:    "sqrdmlsh\tv0.8h, v1.8h, v2.8h",
			wantErr: false,
		},
		{
			name: "sqrdmlah	h0, h1, h2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x84, 0x42, 0x7e}),
				address:          0,
			},
			want:    "sqrdmlah\th0, h1, h2",
			wantErr: false,
		},
		{
			name: "sqrdmlsh	h0, h1, h2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x8c, 0x42, 0x7e}),
				address:          0,
			},
			want:    "sqrdmlsh\th0, h1, h2",
			wantErr: false,
		},
		{
			name: "sqrdmlah	s0, s1, s2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x84, 0x82, 0x7e}),
				address:          0,
			},
			want:    "sqrdmlah\ts0, s1, s2",
			wantErr: false,
		},
		{
			name: "sqrdmlsh	s0, s1, s2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x8c, 0x82, 0x7e}),
				address:          0,
			},
			want:    "sqrdmlsh\ts0, s1, s2",
			wantErr: false,
		},
		{
			name: "sqrdmlah	v0.4h, v1.4h, v2.h[3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xd0, 0x72, 0x2f}),
				address:          0,
			},
			want:    "sqrdmlah\tv0.4h, v1.4h, v2.h[3]",
			wantErr: false,
		},
		{
			name: "sqrdmlsh	v0.4h, v1.4h, v2.h[3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xf0, 0x72, 0x2f}),
				address:          0,
			},
			want:    "sqrdmlsh\tv0.4h, v1.4h, v2.h[3]",
			wantErr: false,
		},
		{
			name: "sqrdmlah	v0.2s, v1.2s, v2.s[1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xd0, 0xa2, 0x2f}),
				address:          0,
			},
			want:    "sqrdmlah\tv0.2s, v1.2s, v2.s[1]",
			wantErr: false,
		},
		{
			name: "sqrdmlsh	v0.2s, v1.2s, v2.s[1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xf0, 0xa2, 0x2f}),
				address:          0,
			},
			want:    "sqrdmlsh\tv0.2s, v1.2s, v2.s[1]",
			wantErr: false,
		},
		{
			name: "sqrdmlah	v0.8h, v1.8h, v2.h[3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xd0, 0x72, 0x6f}),
				address:          0,
			},
			want:    "sqrdmlah\tv0.8h, v1.8h, v2.h[3]",
			wantErr: false,
		},
		{
			name: "sqrdmlsh	v0.8h, v1.8h, v2.h[3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xf0, 0x72, 0x6f}),
				address:          0,
			},
			want:    "sqrdmlsh\tv0.8h, v1.8h, v2.h[3]",
			wantErr: false,
		},
		{
			name: "sqrdmlah	v0.4s, v1.4s, v2.s[3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xd8, 0xa2, 0x6f}),
				address:          0,
			},
			want:    "sqrdmlah\tv0.4s, v1.4s, v2.s[3]",
			wantErr: false,
		},
		{
			name: "sqrdmlsh	v0.4s, v1.4s, v2.s[3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xf8, 0xa2, 0x6f}),
				address:          0,
			},
			want:    "sqrdmlsh\tv0.4s, v1.4s, v2.s[3]",
			wantErr: false,
		},
		{
			name: "sqrdmlah	h0, h1, v2.h[3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xd0, 0x72, 0x7f}),
				address:          0,
			},
			want:    "sqrdmlah\th0, h1, v2.h[3]",
			wantErr: false,
		},
		{
			name: "sqrdmlsh	h0, h1, v2.h[3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xf0, 0x72, 0x7f}),
				address:          0,
			},
			want:    "sqrdmlsh\th0, h1, v2.h[3]",
			wantErr: false,
		},
		{
			name: "sqrdmlah	s0, s1, v2.s[3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xd8, 0xa2, 0x7f}),
				address:          0,
			},
			want:    "sqrdmlah\ts0, s1, v2.s[3]",
			wantErr: false,
		},
		{
			name: "sqrdmlsh	s0, s1, v2.s[3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xf8, 0xa2, 0x7f}),
				address:          0,
			},
			want:    "sqrdmlsh\ts0, s1, v2.s[3]",
			wantErr: false,
		},
		//
		// llvm/test/MC/AArch64/armv8.1a-vhe.s
		//
		{
			name: "msr	ttbr1_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x20, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tttbr1_el2, x0",
			wantErr: false,
		},
		{
			name: "msr	contextidr_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xd0, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tcontextidr_el2, x0",
			wantErr: false,
		},
		{
			name: "msr	cnthv_tval_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0xe3, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tcnthv_tval_el2, x0",
			wantErr: false,
		},
		{
			name: "msr	cnthv_cval_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0xe3, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tcnthv_cval_el2, x0",
			wantErr: false,
		},
		{
			name: "msr	cnthv_ctl_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xe3, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tcnthv_ctl_el2, x0",
			wantErr: false,
		},
		{
			name: "msr	sctlr_el12, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x10, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\tsctlr_el12, x0",
			wantErr: false,
		},
		{
			name: "msr	cpacr_el12, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0x10, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\tcpacr_el12, x0",
			wantErr: false,
		},
		{
			name: "msr	ttbr0_el12, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x20, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\tttbr0_el12, x0",
			wantErr: false,
		},
		{
			name: "msr	ttbr1_el12, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x20, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\tttbr1_el12, x0",
			wantErr: false,
		},
		{
			name: "msr	tcr_el12, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0x20, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\ttcr_el12, x0",
			wantErr: false,
		},
		{
			name: "msr	afsr0_el12, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x51, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\tafsr0_el12, x0",
			wantErr: false,
		},
		{
			name: "msr	afsr1_el12, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x51, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\tafsr1_el12, x0",
			wantErr: false,
		},
		{
			name: "msr	esr_el12, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x52, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\tesr_el12, x0",
			wantErr: false,
		},
		{
			name: "msr	far_el12, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x60, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\tfar_el12, x0",
			wantErr: false,
		},
		{
			name: "msr	mair_el12, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0xa2, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\tmair_el12, x0",
			wantErr: false,
		},
		{
			name: "msr	amair_el12, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0xa3, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\tamair_el12, x0",
			wantErr: false,
		},
		{
			name: "msr	vbar_el12, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0xc0, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\tvbar_el12, x0",
			wantErr: false,
		},
		{
			name: "msr	contextidr_el12, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xd0, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\tcontextidr_el12, x0",
			wantErr: false,
		},
		{
			name: "msr	cntkctl_el12, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0xe1, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\tcntkctl_el12, x0",
			wantErr: false,
		},
		{
			name: "msr	cntp_tval_el02, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0xe2, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\tcntp_tval_el02, x0",
			wantErr: false,
		},
		{
			name: "msr	cntp_ctl_el02, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xe2, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\tcntp_ctl_el02, x0",
			wantErr: false,
		},
		{
			name: "msr	cntp_cval_el02, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0xe2, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\tcntp_cval_el02, x0",
			wantErr: false,
		},
		{
			name: "msr	cntv_tval_el02, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0xe3, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\tcntv_tval_el02, x0",
			wantErr: false,
		},
		{
			name: "msr	cntv_ctl_el02, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xe3, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\tcntv_ctl_el02, x0",
			wantErr: false,
		},
		{
			name: "msr	cntv_cval_el02, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0xe3, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\tcntv_cval_el02, x0",
			wantErr: false,
		},
		{
			name: "msr	spsr_el12, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x40, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\tspsr_el12, x0",
			wantErr: false,
		},
		{
			name: "msr	elr_el12, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x40, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\telr_el12, x0",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decompose(tt.args.address, tt.args.instructionValue, &results)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decompose() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got.Disassembly, tt.want) {
				t.Errorf("Disassembly = %v, want %v", got.Disassembly, tt.want)
			}
		})
	}
}

func Test_decompose_v8_1a_LSE(t *testing.T) {
	var results [1024]byte
	type args struct {
		instructionValue uint32
		address          uint64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		//
		// llvm/test/MC/AArch64/armv8.1a-lse.s
		//
		{
			name: "cas	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x7c, 0xa0, 0x88}),
				address:          0,
			},
			want:    "cas\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "cas	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x7f, 0xa2, 0x88}),
				address:          0,
			},
			want:    "cas\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "casa	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x7c, 0xe0, 0x88}),
				address:          0,
			},
			want:    "casa\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "casa	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x7f, 0xe2, 0x88}),
				address:          0,
			},
			want:    "casa\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "casl	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0xfc, 0xa0, 0x88}),
				address:          0,
			},
			want:    "casl\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "casl	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0xff, 0xa2, 0x88}),
				address:          0,
			},
			want:    "casl\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "casal	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0xfc, 0xe0, 0x88}),
				address:          0,
			},
			want:    "casal\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "casal	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0xff, 0xe2, 0x88}),
				address:          0,
			},
			want:    "casal\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "casb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x7c, 0xa0, 0x08}),
				address:          0,
			},
			want:    "casb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "casb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x7f, 0xa2, 0x08}),
				address:          0,
			},
			want:    "casb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "cash	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x7c, 0xa0, 0x48}),
				address:          0,
			},
			want:    "cash\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "cash	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x7f, 0xa2, 0x48}),
				address:          0,
			},
			want:    "cash\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "casab	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x7c, 0xe0, 0x08}),
				address:          0,
			},
			want:    "casab\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "casab	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x7f, 0xe2, 0x08}),
				address:          0,
			},
			want:    "casab\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "caslb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0xfc, 0xa0, 0x08}),
				address:          0,
			},
			want:    "caslb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "caslb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0xff, 0xa2, 0x08}),
				address:          0,
			},
			want:    "caslb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "casalb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0xfc, 0xe0, 0x08}),
				address:          0,
			},
			want:    "casalb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "casalb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0xff, 0xe2, 0x08}),
				address:          0,
			},
			want:    "casalb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "casah	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x7c, 0xe0, 0x48}),
				address:          0,
			},
			want:    "casah\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "casah	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x7f, 0xe2, 0x48}),
				address:          0,
			},
			want:    "casah\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "caslh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0xfc, 0xa0, 0x48}),
				address:          0,
			},
			want:    "caslh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "caslh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0xff, 0xa2, 0x48}),
				address:          0,
			},
			want:    "caslh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "casalh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0xfc, 0xe0, 0x48}),
				address:          0,
			},
			want:    "casalh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "casalh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0xff, 0xe2, 0x48}),
				address:          0,
			},
			want:    "casalh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "cas	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x7c, 0xa0, 0xc8}),
				address:          0,
			},
			want:    "cas\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "cas	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x7f, 0xa2, 0xc8}),
				address:          0,
			},
			want:    "cas\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "casa	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x7c, 0xe0, 0xc8}),
				address:          0,
			},
			want:    "casa\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "casa	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x7f, 0xe2, 0xc8}),
				address:          0,
			},
			want:    "casa\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "casl	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0xfc, 0xa0, 0xc8}),
				address:          0,
			},
			want:    "casl\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "casl	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0xff, 0xa2, 0xc8}),
				address:          0,
			},
			want:    "casl\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "casal	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0xfc, 0xe0, 0xc8}),
				address:          0,
			},
			want:    "casal\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "casal	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0xff, 0xe2, 0xc8}),
				address:          0,
			},
			want:    "casal\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "swp	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x80, 0x20, 0xb8}),
				address:          0,
			},
			want:    "swp\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "swp	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x83, 0x22, 0xb8}),
				address:          0,
			},
			want:    "swp\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "swpa	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x80, 0xa0, 0xb8}),
				address:          0,
			},
			want:    "swpa\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "swpa	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x83, 0xa2, 0xb8}),
				address:          0,
			},
			want:    "swpa\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "swpl	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x80, 0x60, 0xb8}),
				address:          0,
			},
			want:    "swpl\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "swpl	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x83, 0x62, 0xb8}),
				address:          0,
			},
			want:    "swpl\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "swpal	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x80, 0xe0, 0xb8}),
				address:          0,
			},
			want:    "swpal\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "swpal	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x83, 0xe2, 0xb8}),
				address:          0,
			},
			want:    "swpal\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "swpb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x80, 0x20, 0x38}),
				address:          0,
			},
			want:    "swpb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "swpb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x83, 0x22, 0x38}),
				address:          0,
			},
			want:    "swpb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "swph	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x80, 0x20, 0x78}),
				address:          0,
			},
			want:    "swph\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "swph	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x83, 0x22, 0x78}),
				address:          0,
			},
			want:    "swph\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "swpab	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x80, 0xa0, 0x38}),
				address:          0,
			},
			want:    "swpab\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "swpab	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x83, 0xa2, 0x38}),
				address:          0,
			},
			want:    "swpab\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "swplb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x80, 0x60, 0x38}),
				address:          0,
			},
			want:    "swplb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "swplb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x83, 0x62, 0x38}),
				address:          0,
			},
			want:    "swplb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "swpalb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x80, 0xe0, 0x38}),
				address:          0,
			},
			want:    "swpalb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "swpalb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x83, 0xe2, 0x38}),
				address:          0,
			},
			want:    "swpalb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "swpah	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x80, 0xa0, 0x78}),
				address:          0,
			},
			want:    "swpah\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "swpah	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x83, 0xa2, 0x78}),
				address:          0,
			},
			want:    "swpah\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "swplh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x80, 0x60, 0x78}),
				address:          0,
			},
			want:    "swplh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "swplh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x83, 0x62, 0x78}),
				address:          0,
			},
			want:    "swplh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "swpalh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x80, 0xe0, 0x78}),
				address:          0,
			},
			want:    "swpalh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "swpalh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x83, 0xe2, 0x78}),
				address:          0,
			},
			want:    "swpalh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "swp	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x80, 0x20, 0xf8}),
				address:          0,
			},
			want:    "swp\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "swp	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x83, 0x22, 0xf8}),
				address:          0,
			},
			want:    "swp\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "swpa	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x80, 0xa0, 0xf8}),
				address:          0,
			},
			want:    "swpa\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "swpa	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x83, 0xa2, 0xf8}),
				address:          0,
			},
			want:    "swpa\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "swpl	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x80, 0x60, 0xf8}),
				address:          0,
			},
			want:    "swpl\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "swpl	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x83, 0x62, 0xf8}),
				address:          0,
			},
			want:    "swpl\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "swpal	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x80, 0xe0, 0xf8}),
				address:          0,
			},
			want:    "swpal\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "swpal	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x83, 0xe2, 0xf8}),
				address:          0,
			},
			want:    "swpal\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "casp	w0, w1, w2, w3, [x5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa2, 0x7c, 0x20, 0x08}),
				address:          0,
			},
			want:    "casp\tw0, w1, w2, w3, [x5]",
			wantErr: false,
		},
		{
			name: "casp	w4, w5, w6, w7, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe6, 0x7f, 0x24, 0x08}),
				address:          0,
			},
			want:    "casp\tw4, w5, w6, w7, [sp]",
			wantErr: false,
		},
		{
			name: "casp	x0, x1, x2, x3, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x42, 0x7c, 0x20, 0x48}),
				address:          0,
			},
			want:    "casp\tx0, x1, x2, x3, [x2]",
			wantErr: false,
		},
		{
			name: "casp	x4, x5, x6, x7, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe6, 0x7f, 0x24, 0x48}),
				address:          0,
			},
			want:    "casp\tx4, x5, x6, x7, [sp]",
			wantErr: false,
		},
		{
			name: "caspa	w0, w1, w2, w3, [x5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa2, 0x7c, 0x60, 0x08}),
				address:          0,
			},
			want:    "caspa\tw0, w1, w2, w3, [x5]",
			wantErr: false,
		},
		{
			name: "caspa	w4, w5, w6, w7, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe6, 0x7f, 0x64, 0x08}),
				address:          0,
			},
			want:    "caspa\tw4, w5, w6, w7, [sp]",
			wantErr: false,
		},
		{
			name: "caspa	x0, x1, x2, x3, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x42, 0x7c, 0x60, 0x48}),
				address:          0,
			},
			want:    "caspa\tx0, x1, x2, x3, [x2]",
			wantErr: false,
		},
		{
			name: "caspa	x4, x5, x6, x7, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe6, 0x7f, 0x64, 0x48}),
				address:          0,
			},
			want:    "caspa\tx4, x5, x6, x7, [sp]",
			wantErr: false,
		},
		{
			name: "caspl	w0, w1, w2, w3, [x5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa2, 0xfc, 0x20, 0x08}),
				address:          0,
			},
			want:    "caspl\tw0, w1, w2, w3, [x5]",
			wantErr: false,
		},
		{
			name: "caspl	w4, w5, w6, w7, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe6, 0xff, 0x24, 0x08}),
				address:          0,
			},
			want:    "caspl\tw4, w5, w6, w7, [sp]",
			wantErr: false,
		},
		{
			name: "caspl	x0, x1, x2, x3, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x42, 0xfc, 0x20, 0x48}),
				address:          0,
			},
			want:    "caspl\tx0, x1, x2, x3, [x2]",
			wantErr: false,
		},
		{
			name: "caspl	x4, x5, x6, x7, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe6, 0xff, 0x24, 0x48}),
				address:          0,
			},
			want:    "caspl\tx4, x5, x6, x7, [sp]",
			wantErr: false,
		},
		{
			name: "caspal	w0, w1, w2, w3, [x5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa2, 0xfc, 0x60, 0x08}),
				address:          0,
			},
			want:    "caspal\tw0, w1, w2, w3, [x5]",
			wantErr: false,
		},
		{
			name: "caspal	w4, w5, w6, w7, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe6, 0xff, 0x64, 0x08}),
				address:          0,
			},
			want:    "caspal\tw4, w5, w6, w7, [sp]",
			wantErr: false,
		},
		{
			name: "caspal	x0, x1, x2, x3, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x42, 0xfc, 0x60, 0x48}),
				address:          0,
			},
			want:    "caspal\tx0, x1, x2, x3, [x2]",
			wantErr: false,
		},
		{
			name: "caspal	x4, x5, x6, x7, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe6, 0xff, 0x64, 0x48}),
				address:          0,
			},
			want:    "caspal\tx4, x5, x6, x7, [sp]",
			wantErr: false,
		},
		{
			name: "ldadd	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x00, 0x20, 0xb8}),
				address:          0,
			},
			want:    "ldadd\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldadd	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0x22, 0xb8}),
				address:          0,
			},
			want:    "ldadd\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldadda	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x00, 0xa0, 0xb8}),
				address:          0,
			},
			want:    "ldadda\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldadda	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0xa2, 0xb8}),
				address:          0,
			},
			want:    "ldadda\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldaddl	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x00, 0x60, 0xb8}),
				address:          0,
			},
			want:    "ldaddl\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldaddl	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0x62, 0xb8}),
				address:          0,
			},
			want:    "ldaddl\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldaddal	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x00, 0xe0, 0xb8}),
				address:          0,
			},
			want:    "ldaddal\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldaddal	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0xe2, 0xb8}),
				address:          0,
			},
			want:    "ldaddal\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldaddb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x00, 0x20, 0x38}),
				address:          0,
			},
			want:    "ldaddb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldaddb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0x22, 0x38}),
				address:          0,
			},
			want:    "ldaddb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldaddh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x00, 0x20, 0x78}),
				address:          0,
			},
			want:    "ldaddh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldaddh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0x22, 0x78}),
				address:          0,
			},
			want:    "ldaddh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldaddab	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x00, 0xa0, 0x38}),
				address:          0,
			},
			want:    "ldaddab\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldaddab	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0xa2, 0x38}),
				address:          0,
			},
			want:    "ldaddab\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldaddlb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x00, 0x60, 0x38}),
				address:          0,
			},
			want:    "ldaddlb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldaddlb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0x62, 0x38}),
				address:          0,
			},
			want:    "ldaddlb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldaddalb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x00, 0xe0, 0x38}),
				address:          0,
			},
			want:    "ldaddalb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldaddalb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0xe2, 0x38}),
				address:          0,
			},
			want:    "ldaddalb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldaddah	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x00, 0xa0, 0x78}),
				address:          0,
			},
			want:    "ldaddah\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldaddah	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0xa2, 0x78}),
				address:          0,
			},
			want:    "ldaddah\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldaddlh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x00, 0x60, 0x78}),
				address:          0,
			},
			want:    "ldaddlh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldaddlh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0x62, 0x78}),
				address:          0,
			},
			want:    "ldaddlh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldaddalh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x00, 0xe0, 0x78}),
				address:          0,
			},
			want:    "ldaddalh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldaddalh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0xe2, 0x78}),
				address:          0,
			},
			want:    "ldaddalh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldadd	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x00, 0x20, 0xf8}),
				address:          0,
			},
			want:    "ldadd\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldadd	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0x22, 0xf8}),
				address:          0,
			},
			want:    "ldadd\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldadda	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x00, 0xa0, 0xf8}),
				address:          0,
			},
			want:    "ldadda\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldadda	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0xa2, 0xf8}),
				address:          0,
			},
			want:    "ldadda\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldaddl	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x00, 0x60, 0xf8}),
				address:          0,
			},
			want:    "ldaddl\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldaddl	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0x62, 0xf8}),
				address:          0,
			},
			want:    "ldaddl\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldaddal	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x00, 0xe0, 0xf8}),
				address:          0,
			},
			want:    "ldaddal\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldaddal	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0xe2, 0xf8}),
				address:          0,
			},
			want:    "ldaddal\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldclr	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x10, 0x20, 0xb8}),
				address:          0,
			},
			want:    "ldclr\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldclr	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x13, 0x22, 0xb8}),
				address:          0,
			},
			want:    "ldclr\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldclra	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x10, 0xa0, 0xb8}),
				address:          0,
			},
			want:    "ldclra\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldclra	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x13, 0xa2, 0xb8}),
				address:          0,
			},
			want:    "ldclra\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldclrl	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x10, 0x60, 0xb8}),
				address:          0,
			},
			want:    "ldclrl\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldclrl	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x13, 0x62, 0xb8}),
				address:          0,
			},
			want:    "ldclrl\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldclral	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x10, 0xe0, 0xb8}),
				address:          0,
			},
			want:    "ldclral\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldclral	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x13, 0xe2, 0xb8}),
				address:          0,
			},
			want:    "ldclral\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldclrb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x10, 0x20, 0x38}),
				address:          0,
			},
			want:    "ldclrb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldclrb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x13, 0x22, 0x38}),
				address:          0,
			},
			want:    "ldclrb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldclrh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x10, 0x20, 0x78}),
				address:          0,
			},
			want:    "ldclrh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldclrh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x13, 0x22, 0x78}),
				address:          0,
			},
			want:    "ldclrh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldclrab	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x10, 0xa0, 0x38}),
				address:          0,
			},
			want:    "ldclrab\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldclrab	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x13, 0xa2, 0x38}),
				address:          0,
			},
			want:    "ldclrab\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldclrlb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x10, 0x60, 0x38}),
				address:          0,
			},
			want:    "ldclrlb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldclrlb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x13, 0x62, 0x38}),
				address:          0,
			},
			want:    "ldclrlb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldclralb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x10, 0xe0, 0x38}),
				address:          0,
			},
			want:    "ldclralb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldclralb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x13, 0xe2, 0x38}),
				address:          0,
			},
			want:    "ldclralb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldclrah	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x10, 0xa0, 0x78}),
				address:          0,
			},
			want:    "ldclrah\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldclrah	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x13, 0xa2, 0x78}),
				address:          0,
			},
			want:    "ldclrah\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldclrlh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x10, 0x60, 0x78}),
				address:          0,
			},
			want:    "ldclrlh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldclrlh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x13, 0x62, 0x78}),
				address:          0,
			},
			want:    "ldclrlh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldclralh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x10, 0xe0, 0x78}),
				address:          0,
			},
			want:    "ldclralh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldclralh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x13, 0xe2, 0x78}),
				address:          0,
			},
			want:    "ldclralh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldclr	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x10, 0x20, 0xf8}),
				address:          0,
			},
			want:    "ldclr\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldclr	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x13, 0x22, 0xf8}),
				address:          0,
			},
			want:    "ldclr\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldclra	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x10, 0xa0, 0xf8}),
				address:          0,
			},
			want:    "ldclra\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldclra	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x13, 0xa2, 0xf8}),
				address:          0,
			},
			want:    "ldclra\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldclrl	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x10, 0x60, 0xf8}),
				address:          0,
			},
			want:    "ldclrl\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldclrl	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x13, 0x62, 0xf8}),
				address:          0,
			},
			want:    "ldclrl\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldclral	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x10, 0xe0, 0xf8}),
				address:          0,
			},
			want:    "ldclral\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldclral	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x13, 0xe2, 0xf8}),
				address:          0,
			},
			want:    "ldclral\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldeor	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x20, 0x20, 0xb8}),
				address:          0,
			},
			want:    "ldeor\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldeor	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x23, 0x22, 0xb8}),
				address:          0,
			},
			want:    "ldeor\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldeora	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x20, 0xa0, 0xb8}),
				address:          0,
			},
			want:    "ldeora\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldeora	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x23, 0xa2, 0xb8}),
				address:          0,
			},
			want:    "ldeora\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldeorl	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x20, 0x60, 0xb8}),
				address:          0,
			},
			want:    "ldeorl\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldeorl	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x23, 0x62, 0xb8}),
				address:          0,
			},
			want:    "ldeorl\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldeoral	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x20, 0xe0, 0xb8}),
				address:          0,
			},
			want:    "ldeoral\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldeoral	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x23, 0xe2, 0xb8}),
				address:          0,
			},
			want:    "ldeoral\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldeorb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x20, 0x20, 0x38}),
				address:          0,
			},
			want:    "ldeorb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldeorb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x23, 0x22, 0x38}),
				address:          0,
			},
			want:    "ldeorb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldeorh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x20, 0x20, 0x78}),
				address:          0,
			},
			want:    "ldeorh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldeorh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x23, 0x22, 0x78}),
				address:          0,
			},
			want:    "ldeorh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldeorab	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x20, 0xa0, 0x38}),
				address:          0,
			},
			want:    "ldeorab\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldeorab	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x23, 0xa2, 0x38}),
				address:          0,
			},
			want:    "ldeorab\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldeorlb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x20, 0x60, 0x38}),
				address:          0,
			},
			want:    "ldeorlb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldeorlb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x23, 0x62, 0x38}),
				address:          0,
			},
			want:    "ldeorlb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldeoralb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x20, 0xe0, 0x38}),
				address:          0,
			},
			want:    "ldeoralb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldeoralb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x23, 0xe2, 0x38}),
				address:          0,
			},
			want:    "ldeoralb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldeorah	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x20, 0xa0, 0x78}),
				address:          0,
			},
			want:    "ldeorah\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldeorah	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x23, 0xa2, 0x78}),
				address:          0,
			},
			want:    "ldeorah\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldeorlh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x20, 0x60, 0x78}),
				address:          0,
			},
			want:    "ldeorlh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldeorlh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x23, 0x62, 0x78}),
				address:          0,
			},
			want:    "ldeorlh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldeoralh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x20, 0xe0, 0x78}),
				address:          0,
			},
			want:    "ldeoralh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldeoralh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x23, 0xe2, 0x78}),
				address:          0,
			},
			want:    "ldeoralh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldeor	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x20, 0x20, 0xf8}),
				address:          0,
			},
			want:    "ldeor\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldeor	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x23, 0x22, 0xf8}),
				address:          0,
			},
			want:    "ldeor\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldeora	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x20, 0xa0, 0xf8}),
				address:          0,
			},
			want:    "ldeora\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldeora	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x23, 0xa2, 0xf8}),
				address:          0,
			},
			want:    "ldeora\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldeorl	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x20, 0x60, 0xf8}),
				address:          0,
			},
			want:    "ldeorl\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldeorl	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x23, 0x62, 0xf8}),
				address:          0,
			},
			want:    "ldeorl\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldeoral	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x20, 0xe0, 0xf8}),
				address:          0,
			},
			want:    "ldeoral\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldeoral	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x23, 0xe2, 0xf8}),
				address:          0,
			},
			want:    "ldeoral\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldset	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x30, 0x20, 0xb8}),
				address:          0,
			},
			want:    "ldset\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldset	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x33, 0x22, 0xb8}),
				address:          0,
			},
			want:    "ldset\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldseta	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x30, 0xa0, 0xb8}),
				address:          0,
			},
			want:    "ldseta\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldseta	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x33, 0xa2, 0xb8}),
				address:          0,
			},
			want:    "ldseta\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsetl	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x30, 0x60, 0xb8}),
				address:          0,
			},
			want:    "ldsetl\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsetl	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x33, 0x62, 0xb8}),
				address:          0,
			},
			want:    "ldsetl\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsetal	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x30, 0xe0, 0xb8}),
				address:          0,
			},
			want:    "ldsetal\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsetal	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x33, 0xe2, 0xb8}),
				address:          0,
			},
			want:    "ldsetal\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsetb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x30, 0x20, 0x38}),
				address:          0,
			},
			want:    "ldsetb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsetb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x33, 0x22, 0x38}),
				address:          0,
			},
			want:    "ldsetb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldseth	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x30, 0x20, 0x78}),
				address:          0,
			},
			want:    "ldseth\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldseth	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x33, 0x22, 0x78}),
				address:          0,
			},
			want:    "ldseth\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsetab	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x30, 0xa0, 0x38}),
				address:          0,
			},
			want:    "ldsetab\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsetab	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x33, 0xa2, 0x38}),
				address:          0,
			},
			want:    "ldsetab\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsetlb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x30, 0x60, 0x38}),
				address:          0,
			},
			want:    "ldsetlb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsetlb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x33, 0x62, 0x38}),
				address:          0,
			},
			want:    "ldsetlb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsetalb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x30, 0xe0, 0x38}),
				address:          0,
			},
			want:    "ldsetalb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsetalb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x33, 0xe2, 0x38}),
				address:          0,
			},
			want:    "ldsetalb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsetah	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x30, 0xa0, 0x78}),
				address:          0,
			},
			want:    "ldsetah\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsetah	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x33, 0xa2, 0x78}),
				address:          0,
			},
			want:    "ldsetah\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsetlh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x30, 0x60, 0x78}),
				address:          0,
			},
			want:    "ldsetlh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsetlh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x33, 0x62, 0x78}),
				address:          0,
			},
			want:    "ldsetlh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsetalh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x30, 0xe0, 0x78}),
				address:          0,
			},
			want:    "ldsetalh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsetalh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x33, 0xe2, 0x78}),
				address:          0,
			},
			want:    "ldsetalh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldset	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x30, 0x20, 0xf8}),
				address:          0,
			},
			want:    "ldset\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldset	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x33, 0x22, 0xf8}),
				address:          0,
			},
			want:    "ldset\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldseta	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x30, 0xa0, 0xf8}),
				address:          0,
			},
			want:    "ldseta\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldseta	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x33, 0xa2, 0xf8}),
				address:          0,
			},
			want:    "ldseta\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsetl	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x30, 0x60, 0xf8}),
				address:          0,
			},
			want:    "ldsetl\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsetl	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x33, 0x62, 0xf8}),
				address:          0,
			},
			want:    "ldsetl\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsetal	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x30, 0xe0, 0xf8}),
				address:          0,
			},
			want:    "ldsetal\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsetal	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x33, 0xe2, 0xf8}),
				address:          0,
			},
			want:    "ldsetal\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsmax	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x40, 0x20, 0xb8}),
				address:          0,
			},
			want:    "ldsmax\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsmax	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x43, 0x22, 0xb8}),
				address:          0,
			},
			want:    "ldsmax\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsmaxa	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x40, 0xa0, 0xb8}),
				address:          0,
			},
			want:    "ldsmaxa\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsmaxa	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x43, 0xa2, 0xb8}),
				address:          0,
			},
			want:    "ldsmaxa\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsmaxl	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x40, 0x60, 0xb8}),
				address:          0,
			},
			want:    "ldsmaxl\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsmaxl	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x43, 0x62, 0xb8}),
				address:          0,
			},
			want:    "ldsmaxl\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsmaxal	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x40, 0xe0, 0xb8}),
				address:          0,
			},
			want:    "ldsmaxal\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsmaxal	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x43, 0xe2, 0xb8}),
				address:          0,
			},
			want:    "ldsmaxal\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsmaxb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x40, 0x20, 0x38}),
				address:          0,
			},
			want:    "ldsmaxb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsmaxb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x43, 0x22, 0x38}),
				address:          0,
			},
			want:    "ldsmaxb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsmaxh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x40, 0x20, 0x78}),
				address:          0,
			},
			want:    "ldsmaxh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsmaxh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x43, 0x22, 0x78}),
				address:          0,
			},
			want:    "ldsmaxh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsmaxab	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x40, 0xa0, 0x38}),
				address:          0,
			},
			want:    "ldsmaxab\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsmaxab	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x43, 0xa2, 0x38}),
				address:          0,
			},
			want:    "ldsmaxab\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsmaxlb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x40, 0x60, 0x38}),
				address:          0,
			},
			want:    "ldsmaxlb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsmaxlb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x43, 0x62, 0x38}),
				address:          0,
			},
			want:    "ldsmaxlb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsmaxalb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x40, 0xe0, 0x38}),
				address:          0,
			},
			want:    "ldsmaxalb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsmaxalb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x43, 0xe2, 0x38}),
				address:          0,
			},
			want:    "ldsmaxalb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsmaxah	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x40, 0xa0, 0x78}),
				address:          0,
			},
			want:    "ldsmaxah\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsmaxah	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x43, 0xa2, 0x78}),
				address:          0,
			},
			want:    "ldsmaxah\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsmaxlh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x40, 0x60, 0x78}),
				address:          0,
			},
			want:    "ldsmaxlh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsmaxlh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x43, 0x62, 0x78}),
				address:          0,
			},
			want:    "ldsmaxlh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsmaxalh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x40, 0xe0, 0x78}),
				address:          0,
			},
			want:    "ldsmaxalh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsmaxalh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x43, 0xe2, 0x78}),
				address:          0,
			},
			want:    "ldsmaxalh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsmax	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x40, 0x20, 0xf8}),
				address:          0,
			},
			want:    "ldsmax\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsmax	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x43, 0x22, 0xf8}),
				address:          0,
			},
			want:    "ldsmax\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsmaxa	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x40, 0xa0, 0xf8}),
				address:          0,
			},
			want:    "ldsmaxa\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsmaxa	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x43, 0xa2, 0xf8}),
				address:          0,
			},
			want:    "ldsmaxa\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsmaxl	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x40, 0x60, 0xf8}),
				address:          0,
			},
			want:    "ldsmaxl\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsmaxl	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x43, 0x62, 0xf8}),
				address:          0,
			},
			want:    "ldsmaxl\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsmaxal	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x40, 0xe0, 0xf8}),
				address:          0,
			},
			want:    "ldsmaxal\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsmaxal	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x43, 0xe2, 0xf8}),
				address:          0,
			},
			want:    "ldsmaxal\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsmin	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x50, 0x20, 0xb8}),
				address:          0,
			},
			want:    "ldsmin\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsmin	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x53, 0x22, 0xb8}),
				address:          0,
			},
			want:    "ldsmin\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsmina	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x50, 0xa0, 0xb8}),
				address:          0,
			},
			want:    "ldsmina\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsmina	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x53, 0xa2, 0xb8}),
				address:          0,
			},
			want:    "ldsmina\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsminl	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x50, 0x60, 0xb8}),
				address:          0,
			},
			want:    "ldsminl\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsminl	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x53, 0x62, 0xb8}),
				address:          0,
			},
			want:    "ldsminl\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsminal	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x50, 0xe0, 0xb8}),
				address:          0,
			},
			want:    "ldsminal\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsminal	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x53, 0xe2, 0xb8}),
				address:          0,
			},
			want:    "ldsminal\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsminb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x50, 0x20, 0x38}),
				address:          0,
			},
			want:    "ldsminb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsminb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x53, 0x22, 0x38}),
				address:          0,
			},
			want:    "ldsminb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsminh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x50, 0x20, 0x78}),
				address:          0,
			},
			want:    "ldsminh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsminh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x53, 0x22, 0x78}),
				address:          0,
			},
			want:    "ldsminh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsminab	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x50, 0xa0, 0x38}),
				address:          0,
			},
			want:    "ldsminab\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsminab	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x53, 0xa2, 0x38}),
				address:          0,
			},
			want:    "ldsminab\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsminlb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x50, 0x60, 0x38}),
				address:          0,
			},
			want:    "ldsminlb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsminlb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x53, 0x62, 0x38}),
				address:          0,
			},
			want:    "ldsminlb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsminalb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x50, 0xe0, 0x38}),
				address:          0,
			},
			want:    "ldsminalb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsminalb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x53, 0xe2, 0x38}),
				address:          0,
			},
			want:    "ldsminalb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsminah	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x50, 0xa0, 0x78}),
				address:          0,
			},
			want:    "ldsminah\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsminah	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x53, 0xa2, 0x78}),
				address:          0,
			},
			want:    "ldsminah\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsminlh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x50, 0x60, 0x78}),
				address:          0,
			},
			want:    "ldsminlh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsminlh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x53, 0x62, 0x78}),
				address:          0,
			},
			want:    "ldsminlh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsminalh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x50, 0xe0, 0x78}),
				address:          0,
			},
			want:    "ldsminalh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsminalh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x53, 0xe2, 0x78}),
				address:          0,
			},
			want:    "ldsminalh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsmin	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x50, 0x20, 0xf8}),
				address:          0,
			},
			want:    "ldsmin\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsmin	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x53, 0x22, 0xf8}),
				address:          0,
			},
			want:    "ldsmin\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsmina	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x50, 0xa0, 0xf8}),
				address:          0,
			},
			want:    "ldsmina\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsmina	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x53, 0xa2, 0xf8}),
				address:          0,
			},
			want:    "ldsmina\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsminl	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x50, 0x60, 0xf8}),
				address:          0,
			},
			want:    "ldsminl\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsminl	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x53, 0x62, 0xf8}),
				address:          0,
			},
			want:    "ldsminl\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldsminal	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x50, 0xe0, 0xf8}),
				address:          0,
			},
			want:    "ldsminal\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldsminal	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x53, 0xe2, 0xf8}),
				address:          0,
			},
			want:    "ldsminal\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldumax	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x60, 0x20, 0xb8}),
				address:          0,
			},
			want:    "ldumax\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldumax	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x63, 0x22, 0xb8}),
				address:          0,
			},
			want:    "ldumax\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldumaxa	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x60, 0xa0, 0xb8}),
				address:          0,
			},
			want:    "ldumaxa\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldumaxa	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x63, 0xa2, 0xb8}),
				address:          0,
			},
			want:    "ldumaxa\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldumaxl	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x60, 0x60, 0xb8}),
				address:          0,
			},
			want:    "ldumaxl\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldumaxl	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x63, 0x62, 0xb8}),
				address:          0,
			},
			want:    "ldumaxl\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldumaxal	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x60, 0xe0, 0xb8}),
				address:          0,
			},
			want:    "ldumaxal\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldumaxal	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x63, 0xe2, 0xb8}),
				address:          0,
			},
			want:    "ldumaxal\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldumaxb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x60, 0x20, 0x38}),
				address:          0,
			},
			want:    "ldumaxb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldumaxb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x63, 0x22, 0x38}),
				address:          0,
			},
			want:    "ldumaxb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldumaxh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x60, 0x20, 0x78}),
				address:          0,
			},
			want:    "ldumaxh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldumaxh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x63, 0x22, 0x78}),
				address:          0,
			},
			want:    "ldumaxh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldumaxab	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x60, 0xa0, 0x38}),
				address:          0,
			},
			want:    "ldumaxab\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldumaxab	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x63, 0xa2, 0x38}),
				address:          0,
			},
			want:    "ldumaxab\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldumaxlb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x60, 0x60, 0x38}),
				address:          0,
			},
			want:    "ldumaxlb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldumaxlb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x63, 0x62, 0x38}),
				address:          0,
			},
			want:    "ldumaxlb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldumaxalb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x60, 0xe0, 0x38}),
				address:          0,
			},
			want:    "ldumaxalb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldumaxalb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x63, 0xe2, 0x38}),
				address:          0,
			},
			want:    "ldumaxalb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldumaxah	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x60, 0xa0, 0x78}),
				address:          0,
			},
			want:    "ldumaxah\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldumaxah	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x63, 0xa2, 0x78}),
				address:          0,
			},
			want:    "ldumaxah\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldumaxlh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x60, 0x60, 0x78}),
				address:          0,
			},
			want:    "ldumaxlh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldumaxlh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x63, 0x62, 0x78}),
				address:          0,
			},
			want:    "ldumaxlh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldumaxalh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x60, 0xe0, 0x78}),
				address:          0,
			},
			want:    "ldumaxalh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldumaxalh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x63, 0xe2, 0x78}),
				address:          0,
			},
			want:    "ldumaxalh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldumax	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x60, 0x20, 0xf8}),
				address:          0,
			},
			want:    "ldumax\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldumax	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x63, 0x22, 0xf8}),
				address:          0,
			},
			want:    "ldumax\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldumaxa	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x60, 0xa0, 0xf8}),
				address:          0,
			},
			want:    "ldumaxa\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldumaxa	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x63, 0xa2, 0xf8}),
				address:          0,
			},
			want:    "ldumaxa\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldumaxl	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x60, 0x60, 0xf8}),
				address:          0,
			},
			want:    "ldumaxl\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldumaxl	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x63, 0x62, 0xf8}),
				address:          0,
			},
			want:    "ldumaxl\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldumaxal	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x60, 0xe0, 0xf8}),
				address:          0,
			},
			want:    "ldumaxal\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldumaxal	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x63, 0xe2, 0xf8}),
				address:          0,
			},
			want:    "ldumaxal\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldumin	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x70, 0x20, 0xb8}),
				address:          0,
			},
			want:    "ldumin\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldumin	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x73, 0x22, 0xb8}),
				address:          0,
			},
			want:    "ldumin\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldumina	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x70, 0xa0, 0xb8}),
				address:          0,
			},
			want:    "ldumina\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "ldumina	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x73, 0xa2, 0xb8}),
				address:          0,
			},
			want:    "ldumina\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "lduminl	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x70, 0x60, 0xb8}),
				address:          0,
			},
			want:    "lduminl\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "lduminl	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x73, 0x62, 0xb8}),
				address:          0,
			},
			want:    "lduminl\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "lduminal	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x70, 0xe0, 0xb8}),
				address:          0,
			},
			want:    "lduminal\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "lduminal	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x73, 0xe2, 0xb8}),
				address:          0,
			},
			want:    "lduminal\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "lduminb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x70, 0x20, 0x38}),
				address:          0,
			},
			want:    "lduminb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "lduminb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x73, 0x22, 0x38}),
				address:          0,
			},
			want:    "lduminb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "lduminh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x70, 0x20, 0x78}),
				address:          0,
			},
			want:    "lduminh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "lduminh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x73, 0x22, 0x78}),
				address:          0,
			},
			want:    "lduminh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "lduminab	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x70, 0xa0, 0x38}),
				address:          0,
			},
			want:    "lduminab\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "lduminab	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x73, 0xa2, 0x38}),
				address:          0,
			},
			want:    "lduminab\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "lduminlb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x70, 0x60, 0x38}),
				address:          0,
			},
			want:    "lduminlb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "lduminlb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x73, 0x62, 0x38}),
				address:          0,
			},
			want:    "lduminlb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "lduminalb	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x70, 0xe0, 0x38}),
				address:          0,
			},
			want:    "lduminalb\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "lduminalb	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x73, 0xe2, 0x38}),
				address:          0,
			},
			want:    "lduminalb\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "lduminah	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x70, 0xa0, 0x78}),
				address:          0,
			},
			want:    "lduminah\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "lduminah	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x73, 0xa2, 0x78}),
				address:          0,
			},
			want:    "lduminah\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "lduminlh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x70, 0x60, 0x78}),
				address:          0,
			},
			want:    "lduminlh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "lduminlh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x73, 0x62, 0x78}),
				address:          0,
			},
			want:    "lduminlh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "lduminalh	w0, w1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x70, 0xe0, 0x78}),
				address:          0,
			},
			want:    "lduminalh\tw0, w1, [x2]",
			wantErr: false,
		},
		{
			name: "lduminalh	w2, w3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x73, 0xe2, 0x78}),
				address:          0,
			},
			want:    "lduminalh\tw2, w3, [sp]",
			wantErr: false,
		},
		{
			name: "ldumin	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x70, 0x20, 0xf8}),
				address:          0,
			},
			want:    "ldumin\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldumin	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x73, 0x22, 0xf8}),
				address:          0,
			},
			want:    "ldumin\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "ldumina	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x70, 0xa0, 0xf8}),
				address:          0,
			},
			want:    "ldumina\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "ldumina	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x73, 0xa2, 0xf8}),
				address:          0,
			},
			want:    "ldumina\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "lduminl	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x70, 0x60, 0xf8}),
				address:          0,
			},
			want:    "lduminl\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "lduminl	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x73, 0x62, 0xf8}),
				address:          0,
			},
			want:    "lduminl\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "lduminal	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x70, 0xe0, 0xf8}),
				address:          0,
			},
			want:    "lduminal\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "lduminal	x2, x3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x73, 0xe2, 0xf8}),
				address:          0,
			},
			want:    "lduminal\tx2, x3, [sp]",
			wantErr: false,
		},
		{
			name: "stadd	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x00, 0x20, 0xb8}),
				address:          0,
			},
			want:    "stadd\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stadd	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x22, 0xb8}),
				address:          0,
			},
			want:    "stadd\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "staddl	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x00, 0x60, 0xb8}),
				address:          0,
			},
			want:    "staddl\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "staddl	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x62, 0xb8}),
				address:          0,
			},
			want:    "staddl\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "staddb	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x00, 0x20, 0x38}),
				address:          0,
			},
			want:    "staddb\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "staddb	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x22, 0x38}),
				address:          0,
			},
			want:    "staddb\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "staddh	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x00, 0x20, 0x78}),
				address:          0,
			},
			want:    "staddh\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "staddh	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x22, 0x78}),
				address:          0,
			},
			want:    "staddh\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "staddlb	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x00, 0x60, 0x38}),
				address:          0,
			},
			want:    "staddlb\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "staddlb	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x62, 0x38}),
				address:          0,
			},
			want:    "staddlb\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "staddlh	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x00, 0x60, 0x78}),
				address:          0,
			},
			want:    "staddlh\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "staddlh	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x62, 0x78}),
				address:          0,
			},
			want:    "staddlh\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stadd	x0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x00, 0x20, 0xf8}),
				address:          0,
			},
			want:    "stadd\tx0, [x2]",
			wantErr: false,
		},
		{
			name: "stadd	x2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x22, 0xf8}),
				address:          0,
			},
			want:    "stadd\tx2, [sp]",
			wantErr: false,
		},
		{
			name: "staddl	x0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x00, 0x60, 0xf8}),
				address:          0,
			},
			want:    "staddl\tx0, [x2]",
			wantErr: false,
		},
		{
			name: "staddl	x2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x62, 0xf8}),
				address:          0,
			},
			want:    "staddl\tx2, [sp]",
			wantErr: false,
		},
		{
			name: "stclr	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x10, 0x20, 0xb8}),
				address:          0,
			},
			want:    "stclr\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stclr	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x13, 0x22, 0xb8}),
				address:          0,
			},
			want:    "stclr\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stclrl	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x10, 0x60, 0xb8}),
				address:          0,
			},
			want:    "stclrl\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stclrl	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x13, 0x62, 0xb8}),
				address:          0,
			},
			want:    "stclrl\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stclrb	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x10, 0x20, 0x38}),
				address:          0,
			},
			want:    "stclrb\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stclrb	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x13, 0x22, 0x38}),
				address:          0,
			},
			want:    "stclrb\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stclrh	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x10, 0x20, 0x78}),
				address:          0,
			},
			want:    "stclrh\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stclrh	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x13, 0x22, 0x78}),
				address:          0,
			},
			want:    "stclrh\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stclrlb	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x10, 0x60, 0x38}),
				address:          0,
			},
			want:    "stclrlb\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stclrlb	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x13, 0x62, 0x38}),
				address:          0,
			},
			want:    "stclrlb\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stclrlh	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x10, 0x60, 0x78}),
				address:          0,
			},
			want:    "stclrlh\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stclrlh	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x13, 0x62, 0x78}),
				address:          0,
			},
			want:    "stclrlh\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stclr	x0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x10, 0x20, 0xf8}),
				address:          0,
			},
			want:    "stclr\tx0, [x2]",
			wantErr: false,
		},
		{
			name: "stclr	x2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x13, 0x22, 0xf8}),
				address:          0,
			},
			want:    "stclr\tx2, [sp]",
			wantErr: false,
		},
		{
			name: "stclrl	x0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x10, 0x60, 0xf8}),
				address:          0,
			},
			want:    "stclrl\tx0, [x2]",
			wantErr: false,
		},
		{
			name: "stclrl	x2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x13, 0x62, 0xf8}),
				address:          0,
			},
			want:    "stclrl\tx2, [sp]",
			wantErr: false,
		},
		{
			name: "steor	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x20, 0x20, 0xb8}),
				address:          0,
			},
			want:    "steor\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "steor	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x23, 0x22, 0xb8}),
				address:          0,
			},
			want:    "steor\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "steorl	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x20, 0x60, 0xb8}),
				address:          0,
			},
			want:    "steorl\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "steorl	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x23, 0x62, 0xb8}),
				address:          0,
			},
			want:    "steorl\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "steorb	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x20, 0x20, 0x38}),
				address:          0,
			},
			want:    "steorb\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "steorb	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x23, 0x22, 0x38}),
				address:          0,
			},
			want:    "steorb\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "steorh	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x20, 0x20, 0x78}),
				address:          0,
			},
			want:    "steorh\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "steorh	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x23, 0x22, 0x78}),
				address:          0,
			},
			want:    "steorh\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "steorlb	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x20, 0x60, 0x38}),
				address:          0,
			},
			want:    "steorlb\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "steorlb	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x23, 0x62, 0x38}),
				address:          0,
			},
			want:    "steorlb\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "steorlh	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x20, 0x60, 0x78}),
				address:          0,
			},
			want:    "steorlh\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "steorlh	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x23, 0x62, 0x78}),
				address:          0,
			},
			want:    "steorlh\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "steor	x0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x20, 0x20, 0xf8}),
				address:          0,
			},
			want:    "steor\tx0, [x2]",
			wantErr: false,
		},
		{
			name: "steor	x2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x23, 0x22, 0xf8}),
				address:          0,
			},
			want:    "steor\tx2, [sp]",
			wantErr: false,
		},
		{
			name: "steorl	x0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x20, 0x60, 0xf8}),
				address:          0,
			},
			want:    "steorl\tx0, [x2]",
			wantErr: false,
		},
		{
			name: "steorl	x2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x23, 0x62, 0xf8}),
				address:          0,
			},
			want:    "steorl\tx2, [sp]",
			wantErr: false,
		},
		{
			name: "stset	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x30, 0x20, 0xb8}),
				address:          0,
			},
			want:    "stset\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stset	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x33, 0x22, 0xb8}),
				address:          0,
			},
			want:    "stset\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stsetl	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x30, 0x60, 0xb8}),
				address:          0,
			},
			want:    "stsetl\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stsetl	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x33, 0x62, 0xb8}),
				address:          0,
			},
			want:    "stsetl\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stsetb	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x30, 0x20, 0x38}),
				address:          0,
			},
			want:    "stsetb\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stsetb	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x33, 0x22, 0x38}),
				address:          0,
			},
			want:    "stsetb\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stseth	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x30, 0x20, 0x78}),
				address:          0,
			},
			want:    "stseth\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stseth	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x33, 0x22, 0x78}),
				address:          0,
			},
			want:    "stseth\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stsetlb	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x30, 0x60, 0x38}),
				address:          0,
			},
			want:    "stsetlb\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stsetlb	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x33, 0x62, 0x38}),
				address:          0,
			},
			want:    "stsetlb\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stsetlh	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x30, 0x60, 0x78}),
				address:          0,
			},
			want:    "stsetlh\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stsetlh	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x33, 0x62, 0x78}),
				address:          0,
			},
			want:    "stsetlh\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stset	x0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x30, 0x20, 0xf8}),
				address:          0,
			},
			want:    "stset\tx0, [x2]",
			wantErr: false,
		},
		{
			name: "stset	x2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x33, 0x22, 0xf8}),
				address:          0,
			},
			want:    "stset\tx2, [sp]",
			wantErr: false,
		},
		{
			name: "stsetl	x0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x30, 0x60, 0xf8}),
				address:          0,
			},
			want:    "stsetl\tx0, [x2]",
			wantErr: false,
		},
		{
			name: "stsetl	x2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x33, 0x62, 0xf8}),
				address:          0,
			},
			want:    "stsetl\tx2, [sp]",
			wantErr: false,
		},
		{
			name: "stsmax	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x40, 0x20, 0xb8}),
				address:          0,
			},
			want:    "stsmax\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stsmax	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x43, 0x22, 0xb8}),
				address:          0,
			},
			want:    "stsmax\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stsmaxl	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x40, 0x60, 0xb8}),
				address:          0,
			},
			want:    "stsmaxl\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stsmaxl	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x43, 0x62, 0xb8}),
				address:          0,
			},
			want:    "stsmaxl\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stsmaxb	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x40, 0x20, 0x38}),
				address:          0,
			},
			want:    "stsmaxb\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stsmaxb	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x43, 0x22, 0x38}),
				address:          0,
			},
			want:    "stsmaxb\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stsmaxh	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x40, 0x20, 0x78}),
				address:          0,
			},
			want:    "stsmaxh\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stsmaxh	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x43, 0x22, 0x78}),
				address:          0,
			},
			want:    "stsmaxh\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stsmaxlb	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x40, 0x60, 0x38}),
				address:          0,
			},
			want:    "stsmaxlb\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stsmaxlb	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x43, 0x62, 0x38}),
				address:          0,
			},
			want:    "stsmaxlb\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stsmaxlh	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x40, 0x60, 0x78}),
				address:          0,
			},
			want:    "stsmaxlh\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stsmaxlh	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x43, 0x62, 0x78}),
				address:          0,
			},
			want:    "stsmaxlh\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stsmax	x0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x40, 0x20, 0xf8}),
				address:          0,
			},
			want:    "stsmax\tx0, [x2]",
			wantErr: false,
		},
		{
			name: "stsmax	x2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x43, 0x22, 0xf8}),
				address:          0,
			},
			want:    "stsmax\tx2, [sp]",
			wantErr: false,
		},
		{
			name: "stsmaxl	x0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x40, 0x60, 0xf8}),
				address:          0,
			},
			want:    "stsmaxl\tx0, [x2]",
			wantErr: false,
		},
		{
			name: "stsmaxl	x2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x43, 0x62, 0xf8}),
				address:          0,
			},
			want:    "stsmaxl\tx2, [sp]",
			wantErr: false,
		},
		{
			name: "stsmin	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x50, 0x20, 0xb8}),
				address:          0,
			},
			want:    "stsmin\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stsmin	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x53, 0x22, 0xb8}),
				address:          0,
			},
			want:    "stsmin\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stsminl	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x50, 0x60, 0xb8}),
				address:          0,
			},
			want:    "stsminl\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stsminl	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x53, 0x62, 0xb8}),
				address:          0,
			},
			want:    "stsminl\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stsminb	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x50, 0x20, 0x38}),
				address:          0,
			},
			want:    "stsminb\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stsminb	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x53, 0x22, 0x38}),
				address:          0,
			},
			want:    "stsminb\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stsminh	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x50, 0x20, 0x78}),
				address:          0,
			},
			want:    "stsminh\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stsminh	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x53, 0x22, 0x78}),
				address:          0,
			},
			want:    "stsminh\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stsminlb	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x50, 0x60, 0x38}),
				address:          0,
			},
			want:    "stsminlb\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stsminlb	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x53, 0x62, 0x38}),
				address:          0,
			},
			want:    "stsminlb\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stsminlh	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x50, 0x60, 0x78}),
				address:          0,
			},
			want:    "stsminlh\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stsminlh	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x53, 0x62, 0x78}),
				address:          0,
			},
			want:    "stsminlh\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stsmin	x0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x50, 0x20, 0xf8}),
				address:          0,
			},
			want:    "stsmin\tx0, [x2]",
			wantErr: false,
		},
		{
			name: "stsmin	x2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x53, 0x22, 0xf8}),
				address:          0,
			},
			want:    "stsmin\tx2, [sp]",
			wantErr: false,
		},
		{
			name: "stsminl	x0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x50, 0x60, 0xf8}),
				address:          0,
			},
			want:    "stsminl\tx0, [x2]",
			wantErr: false,
		},
		{
			name: "stsminl	x2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x53, 0x62, 0xf8}),
				address:          0,
			},
			want:    "stsminl\tx2, [sp]",
			wantErr: false,
		},
		{
			name: "stumax	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x60, 0x20, 0xb8}),
				address:          0,
			},
			want:    "stumax\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stumax	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x63, 0x22, 0xb8}),
				address:          0,
			},
			want:    "stumax\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stumaxl	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x60, 0x60, 0xb8}),
				address:          0,
			},
			want:    "stumaxl\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stumaxl	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x63, 0x62, 0xb8}),
				address:          0,
			},
			want:    "stumaxl\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stumaxb	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x60, 0x20, 0x38}),
				address:          0,
			},
			want:    "stumaxb\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stumaxb	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x63, 0x22, 0x38}),
				address:          0,
			},
			want:    "stumaxb\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stumaxh	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x60, 0x20, 0x78}),
				address:          0,
			},
			want:    "stumaxh\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stumaxh	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x63, 0x22, 0x78}),
				address:          0,
			},
			want:    "stumaxh\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stumaxlb	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x60, 0x60, 0x38}),
				address:          0,
			},
			want:    "stumaxlb\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stumaxlb	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x63, 0x62, 0x38}),
				address:          0,
			},
			want:    "stumaxlb\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stumaxlh	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x60, 0x60, 0x78}),
				address:          0,
			},
			want:    "stumaxlh\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stumaxlh	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x63, 0x62, 0x78}),
				address:          0,
			},
			want:    "stumaxlh\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stumax	x0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x60, 0x20, 0xf8}),
				address:          0,
			},
			want:    "stumax\tx0, [x2]",
			wantErr: false,
		},
		{
			name: "stumax	x2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x63, 0x22, 0xf8}),
				address:          0,
			},
			want:    "stumax\tx2, [sp]",
			wantErr: false,
		},
		{
			name: "stumaxl	x0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x60, 0x60, 0xf8}),
				address:          0,
			},
			want:    "stumaxl\tx0, [x2]",
			wantErr: false,
		},
		{
			name: "stumaxl	x2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x63, 0x62, 0xf8}),
				address:          0,
			},
			want:    "stumaxl\tx2, [sp]",
			wantErr: false,
		},
		{
			name: "stumin	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x70, 0x20, 0xb8}),
				address:          0,
			},
			want:    "stumin\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stumin	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x73, 0x22, 0xb8}),
				address:          0,
			},
			want:    "stumin\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stuminl	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x70, 0x60, 0xb8}),
				address:          0,
			},
			want:    "stuminl\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stuminl	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x73, 0x62, 0xb8}),
				address:          0,
			},
			want:    "stuminl\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stuminb	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x70, 0x20, 0x38}),
				address:          0,
			},
			want:    "stuminb\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stuminb	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x73, 0x22, 0x38}),
				address:          0,
			},
			want:    "stuminb\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "stuminh	w0, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x70, 0x20, 0x78}),
				address:          0,
			},
			want:    "stuminh\tw0, [x2]",
			wantErr: false,
		},
		{
			name: "stuminh	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x73, 0x22, 0x78}),
				address:          0,
			},
			want:    "stuminh\tw2, [sp]",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decompose(tt.args.address, tt.args.instructionValue, &results)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decompose() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got.Disassembly, tt.want) {
				t.Errorf("Disassembly = %v, want %v", got.Disassembly, tt.want)
			}
		})
	}
}

func Test_decompose_v8_2a(t *testing.T) {
	var results [1024]byte
	type args struct {
		instructionValue uint32
		address          uint64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// llvm/test/MC/AArch64/armv8.2a-at.s
		{
			name: "at	s1e1rp, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x01, 0x79, 0x08, 0xd5}),
				address:          0,
			},
			want:    "at\tS1E1RP, x1",
			wantErr: false,
		},
		{
			name: "at	s1e1wp, x2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x22, 0x79, 0x08, 0xd5}),
				address:          0,
			},
			want:    "at\tS1E1WP, x2",
			wantErr: false,
		},
		// llvm/test/MC/AArch64/armv8.2a-crypto-apple.s
		{
			name: "sha512h.2d	q0, q1, v2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x80, 0x62, 0xce}),
				address:          0,
			},
			want:    "sha512h\tq0, q1, v2.2d",
			wantErr: false,
		},
		{
			name: "sha512h2.2d	q0, q1, v2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x84, 0x62, 0xce}),
				address:          0,
			},
			want:    "sha512h2\tq0, q1, v2.2d",
			wantErr: false,
		},
		{
			name: "sha512su0.2d	v11, v12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8b, 0x81, 0xc0, 0xce}),
				address:          0,
			},
			want:    "sha512su0\tv11.2d, v12.2d",
			wantErr: false,
		},
		{
			name: "sha512su1.2d	v11, v13, v14",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xab, 0x89, 0x6e, 0xce}),
				address:          0,
			},
			want:    "sha512su1\tv11.2d, v13.2d, v14.2d",
			wantErr: false,
		},
		{
			name: "eor3.16b	v25, v12, v7, v2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x99, 0x09, 0x07, 0xce}),
				address:          0,
			},
			want:    "eor3\tv25.16b, v12.16b, v7.16b, v2.16b",
			wantErr: false,
		},
		{
			name: "rax1.2d	v30, v29, v26",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbe, 0x8f, 0x7a, 0xce}),
				address:          0,
			},
			want:    "rax1\tv30.2d, v29.2d, v26.2d",
			wantErr: false,
		},
		{
			name: "xar.2d	v26, v21, v27, #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xba, 0xfe, 0x9b, 0xce}),
				address:          0,
			},
			want:    "xar\tv26.2d, v21.2d, v27.2d, #0x3f",
			wantErr: false,
		},
		{
			name: "bcax.16b	v31, v26, v2, v1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x07, 0x22, 0xce}),
				address:          0,
			},
			want:    "bcax\tv31.16b, v26.16b, v2.16b, v1.16b",
			wantErr: false,
		},
		{
			name: "sm3ss1.4s	v20, v23, v21, v22",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x5a, 0x55, 0xce}),
				address:          0,
			},
			want:    "sm3ss1\tv20.4s, v23.4s, v21.4s, v22.4s",
			wantErr: false,
		},
		{
			name: "sm3tt1a.4s	v20, v23, v21[3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0xb2, 0x55, 0xce}),
				address:          0,
			},
			want:    "sm3tt1a\tv20.4s, v23.4s, v21.s[3]",
			wantErr: false,
		},
		{
			name: "sm3tt1b.4s	v20, v23, v21[3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0xb6, 0x55, 0xce}),
				address:          0,
			},
			want:    "sm3tt1b\tv20.4s, v23.4s, v21.s[3]",
			wantErr: false,
		},
		{
			name: "sm3tt2a.4s	v20, v23, v21[3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0xba, 0x55, 0xce}),
				address:          0,
			},
			want:    "sm3tt2a\tv20.4s, v23.4s, v21.s[3]",
			wantErr: false,
		},
		{
			name: "sm3tt2b.4s	v20, v23, v21[3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0xbe, 0x55, 0xce}),
				address:          0,
			},
			want:    "sm3tt2b\tv20.4s, v23.4s, v21.s[3]",
			wantErr: false,
		},
		{
			name: "sm3partw1.4s	v30, v29, v26",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbe, 0xc3, 0x7a, 0xce}),
				address:          0,
			},
			want:    "sm3partw1\tv30.4s, v29.4s, v26.4s",
			wantErr: false,
		},
		{
			name: "sm3partw2.4s	v30, v29, v26",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbe, 0xc7, 0x7a, 0xce}),
				address:          0,
			},
			want:    "sm3partw2\tv30.4s, v29.4s, v26.4s",
			wantErr: false,
		},
		{
			name: "sm4ekey.4s	v11, v11, v19",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6b, 0xc9, 0x73, 0xce}),
				address:          0,
			},
			want:    "sm4ekey\tv11.4s, v11.4s, v19.4s",
			wantErr: false,
		},
		{
			name: "sm4e.4s	v2, v15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe2, 0x85, 0xc0, 0xce}),
				address:          0,
			},
			want:    "sm4e\tv2.4s, v15.4s",
			wantErr: false,
		},
		// llvm/test/MC/AArch64/armv8.2a-crypto.s
		{
			name: "sha512h	q0, q1, v2.2d",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x80, 0x62, 0xce}),
				address:          0,
			},
			want:    "sha512h\tq0, q1, v2.2d",
			wantErr: false,
		},
		{
			name: "sha512h2	q0, q1, v2.2d",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x84, 0x62, 0xce}),
				address:          0,
			},
			want:    "sha512h2\tq0, q1, v2.2d",
			wantErr: false,
		},
		{
			name: "sha512su0	v11.2d, v12.2d",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8b, 0x81, 0xc0, 0xce}),
				address:          0,
			},
			want:    "sha512su0\tv11.2d, v12.2d",
			wantErr: false,
		},
		{
			name: "sha512su1	v11.2d, v13.2d, v14.2d",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xab, 0x89, 0x6e, 0xce}),
				address:          0,
			},
			want:    "sha512su1\tv11.2d, v13.2d, v14.2d",
			wantErr: false,
		},
		{
			name: "eor3	v25.16b, v12.16b, v7.16b, v2.16b",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x99, 0x09, 0x07, 0xce}),
				address:          0,
			},
			want:    "eor3\tv25.16b, v12.16b, v7.16b, v2.16b",
			wantErr: false,
		},
		{
			name: "rax1	v30.2d, v29.2d, v26.2d",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbe, 0x8f, 0x7a, 0xce}),
				address:          0,
			},
			want:    "rax1\tv30.2d, v29.2d, v26.2d",
			wantErr: false,
		},
		{
			name: "xar	v26.2d, v21.2d, v27.2d, #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xba, 0xfe, 0x9b, 0xce}),
				address:          0,
			},
			want:    "xar\tv26.2d, v21.2d, v27.2d, #0x3f",
			wantErr: false,
		},
		{
			name: "bcax	v31.16b, v26.16b, v2.16b, v1.16b",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x07, 0x22, 0xce}),
				address:          0,
			},
			want:    "bcax\tv31.16b, v26.16b, v2.16b, v1.16b",
			wantErr: false,
		},
		{
			name: "sm3ss1	v20.4s, v23.4s, v21.4s, v22.4s",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x5a, 0x55, 0xce}),
				address:          0,
			},
			want:    "sm3ss1\tv20.4s, v23.4s, v21.4s, v22.4s",
			wantErr: false,
		},
		{
			name: "sm3tt1a	v20.4s, v23.4s, v21.s[3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0xb2, 0x55, 0xce}),
				address:          0,
			},
			want:    "sm3tt1a\tv20.4s, v23.4s, v21.s[3]",
			wantErr: false,
		},
		{
			name: "sm3tt1b	v20.4s, v23.4s, v21.s[3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0xb6, 0x55, 0xce}),
				address:          0,
			},
			want:    "sm3tt1b\tv20.4s, v23.4s, v21.s[3]",
			wantErr: false,
		},
		{
			name: "sm3tt2a	v20.4s, v23.4s, v21.s[3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0xba, 0x55, 0xce}),
				address:          0,
			},
			want:    "sm3tt2a\tv20.4s, v23.4s, v21.s[3]",
			wantErr: false,
		},
		{
			name: "sm3tt2b	v20.4s, v23.4s, v21.s[3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0xbe, 0x55, 0xce}),
				address:          0,
			},
			want:    "sm3tt2b\tv20.4s, v23.4s, v21.s[3]",
			wantErr: false,
		},
		{
			name: "sm3partw1	v30.4s, v29.4s, v26.4s",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbe, 0xc3, 0x7a, 0xce}),
				address:          0,
			},
			want:    "sm3partw1\tv30.4s, v29.4s, v26.4s",
			wantErr: false,
		},
		{
			name: "sm3partw2	v30.4s, v29.4s, v26.4s",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbe, 0xc7, 0x7a, 0xce}),
				address:          0,
			},
			want:    "sm3partw2\tv30.4s, v29.4s, v26.4s",
			wantErr: false,
		},
		{
			name: "sm4ekey	v11.4s, v11.4s, v19.4s",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6b, 0xc9, 0x73, 0xce}),
				address:          0,
			},
			want:    "sm4ekey\tv11.4s, v11.4s, v19.4s",
			wantErr: false,
		},
		{
			name: "sm4e	v2.4s, v15.4s",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe2, 0x85, 0xc0, 0xce}),
				address:          0,
			},
			want:    "sm4e\tv2.4s, v15.4s",
			wantErr: false,
		},
		// llvm/test/MC/AArch64/armv8.2a-dotprod.s
		{
			name: "udot	v0.2s, v1.8b, v2.8b",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x94, 0x82, 0x2e}),
				address:          0,
			},
			want:    "udot\tv0.2s, v1.8b, v2.8b",
			wantErr: false,
		},
		{
			name: "sdot	v0.2s, v1.8b, v2.8b",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x94, 0x82, 0x0e}),
				address:          0,
			},
			want:    "sdot\tv0.2s, v1.8b, v2.8b",
			wantErr: false,
		},
		{
			name: "udot	v0.4s, v1.16b, v2.16b",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x94, 0x82, 0x6e}),
				address:          0,
			},
			want:    "udot\tv0.4s, v1.16b, v2.16b",
			wantErr: false,
		},
		{
			name: "sdot	v0.4s, v1.16b, v2.16b",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x94, 0x82, 0x4e}),
				address:          0,
			},
			want:    "sdot\tv0.4s, v1.16b, v2.16b",
			wantErr: false,
		},
		{
			name: "udot	v0.2s, v1.8b, v2.4b[0]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xe0, 0x82, 0x2f}),
				address:          0,
			},
			want:    "udot\tv0.2s, v1.8b, v2.4b[0]",
			wantErr: false,
		},
		{
			name: "sdot	v0.2s, v1.8b, v2.4b[1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xe0, 0xa2, 0x0f}),
				address:          0,
			},
			want:    "sdot\tv0.2s, v1.8b, v2.4b[1]",
			wantErr: false,
		},
		{
			name: "udot	v0.4s, v1.16b, v2.4b[2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xe8, 0x82, 0x6f}),
				address:          0,
			},
			want:    "udot\tv0.4s, v1.16b, v2.4b[2]",
			wantErr: false,
		},
		{
			name: "sdot	v0.4s, v1.16b, v2.4b[3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xe8, 0xa2, 0x4f}),
				address:          0,
			},
			want:    "sdot\tv0.4s, v1.16b, v2.4b[3]",
			wantErr: false,
		},
		{
			name: "udot	v0.2s, v1.8b, v2.4b[0]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xe0, 0x82, 0x2f}),
				address:          0,
			},
			want:    "udot\tv0.2s, v1.8b, v2.4b[0]",
			wantErr: false,
		},
		{
			name: "udot	v0.4s, v1.16b, v2.4b[2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xe8, 0x82, 0x6f}),
				address:          0,
			},
			want:    "udot\tv0.4s, v1.16b, v2.4b[2]",
			wantErr: false,
		},
		// llvm/test/MC/AArch64/armv8.2a-persistent-memory.s
		{
			name: "dc	cvap, x7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x27, 0x7c, 0x0b, 0xd5}),
				address:          0,
			},
			want:    "dc\tCVAP, x7",
			wantErr: false,
		},
		// llvm/test/MC/AArch64/armv8.2a-statistical-profiling.s
		{
			name: "psb	csync",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0x22, 0x03, 0xd5}),
				address:          0,
			},
			want:    "psb\tcsync",
			wantErr: false,
		},
		{
			name: "msr	pmblimitr_el1, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x9a, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmblimitr_el1, x0",
			wantErr: false,
		},
		{
			name: "msr	pmbptr_el1, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x9a, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmbptr_el1, x0",
			wantErr: false,
		},
		{
			name: "msr	pmbsr_el1, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x60, 0x9a, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmbsr_el1, x0",
			wantErr: false,
		},
		{
			name: "msr	pmscr_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x99, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmscr_el2, x0",
			wantErr: false,
		},
		{
			name: "msr	pmscr_el12, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x99, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmscr_el12, x0",
			wantErr: false,
		},
		{
			name: "msr	pmscr_el1, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x99, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmscr_el1, x0",
			wantErr: false,
		},
		{
			name: "msr	pmsicr_el1, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0x99, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmsicr_el1, x0",
			wantErr: false,
		},
		{
			name: "msr	pmsirr_el1, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x60, 0x99, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmsirr_el1, x0",
			wantErr: false,
		},
		{
			name: "msr	pmsfcr_el1, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x80, 0x99, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmsfcr_el1, x0",
			wantErr: false,
		},
		{
			name: "msr	pmsevfr_el1, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa0, 0x99, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmsevfr_el1, x0",
			wantErr: false,
		},
		{
			name: "msr	pmslatfr_el1, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc0, 0x99, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmslatfr_el1, x0",
			wantErr: false,
		},
		{
			name: "mrs	x0, pmblimitr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x9a, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, pmblimitr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x0, pmbptr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x9a, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, pmbptr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x0, pmbsr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x60, 0x9a, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, pmbsr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x0, pmbidr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x9a, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, pmbidr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x0, pmscr_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x99, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, pmscr_el2",
			wantErr: false,
		},
		{
			name: "mrs	x0, pmscr_el12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x99, 0x3d, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, pmscr_el12",
			wantErr: false,
		},
		{
			name: "mrs	x0, pmscr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x99, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, pmscr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x0, pmsicr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0x99, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, pmsicr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x0, pmsirr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x60, 0x99, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, pmsirr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x0, pmsfcr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x80, 0x99, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, pmsfcr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x0, pmsevfr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa0, 0x99, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, pmsevfr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x0, pmslatfr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc0, 0x99, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, pmslatfr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x0, pmsidr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x99, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, pmsidr_el1",
			wantErr: false,
		},
		// llvm/test/MC/AArch64/armv8.2a-uao.s
		{
			name: "msr	uao, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x40, 0x00, 0xd5}),
				address:          0,
			},
			want:    "msr\tuao, #0",
			wantErr: false,
		},
		{
			name: "msr	uao, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x41, 0x00, 0xd5}),
				address:          0,
			},
			want:    "msr\tuao, #0x1",
			wantErr: false,
		},
		{
			name: "msr	uao, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x81, 0x42, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tuao, x1",
			wantErr: false,
		},
		{
			name: "mrs	x2, uao",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x82, 0x42, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx2, uao",
			wantErr: false,
		},
		//
		// llvm/test/MC/AArch64/armv8a-fpmul.s
		//
		{
			name: "fmlal	v0.2s, v1.2h, v2.2h",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xec, 0x22, 0x0e}),
				address:          0,
			},
			want:    "fmlal\tv0.2s, v1.2h, v2.2h",
			wantErr: false,
		},
		{
			name: "fmlsl	v0.2s, v1.2h, v2.2h",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xec, 0xa2, 0x0e}),
				address:          0,
			},
			want:    "fmlsl\tv0.2s, v1.2h, v2.2h",
			wantErr: false,
		},
		{
			name: "fmlal	v0.4s, v1.4h, v2.4h",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xec, 0x22, 0x4e}),
				address:          0,
			},
			want:    "fmlal\tv0.4s, v1.4h, v2.4h",
			wantErr: false,
		},
		{
			name: "fmlsl	v0.4s, v1.4h, v2.4h",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xec, 0xa2, 0x4e}),
				address:          0,
			},
			want:    "fmlsl\tv0.4s, v1.4h, v2.4h",
			wantErr: false,
		},
		{
			name: "fmlal2	v0.2s, v1.2h, v2.2h",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xcc, 0x22, 0x2e}),
				address:          0,
			},
			want:    "fmlal2\tv0.2s, v1.2h, v2.2h",
			wantErr: false,
		},
		{
			name: "fmlsl2	v0.2s, v1.2h, v2.2h",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xcc, 0xa2, 0x2e}),
				address:          0,
			},
			want:    "fmlsl2\tv0.2s, v1.2h, v2.2h",
			wantErr: false,
		},
		{
			name: "fmlal2	v0.4s, v1.4h, v2.4h",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xcc, 0x22, 0x6e}),
				address:          0,
			},
			want:    "fmlal2\tv0.4s, v1.4h, v2.4h",
			wantErr: false,
		},
		{
			name: "fmlsl2	v0.4s, v1.4h, v2.4h",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xcc, 0xa2, 0x6e}),
				address:          0,
			},
			want:    "fmlsl2\tv0.4s, v1.4h, v2.4h",
			wantErr: false,
		},
		{
			name: "fmlal	v0.2s, v1.2h, v2.h[7]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x08, 0xb2, 0x0f}),
				address:          0,
			},
			want:    "fmlal\tv0.2s, v1.2h, v2.h[7]",
			wantErr: false,
		},
		{
			name: "fmlsl	v0.2s, v1.2h, v2.h[7]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x48, 0xb2, 0x0f}),
				address:          0,
			},
			want:    "fmlsl\tv0.2s, v1.2h, v2.h[7]",
			wantErr: false,
		},
		{
			name: "fmlal	v0.4s, v1.4h, v2.h[7]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x08, 0xb2, 0x4f}),
				address:          0,
			},
			want:    "fmlal\tv0.4s, v1.4h, v2.h[7]",
			wantErr: false,
		},
		{
			name: "fmlsl	v0.4s, v1.4h, v2.h[7]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x48, 0xb2, 0x4f}),
				address:          0,
			},
			want:    "fmlsl\tv0.4s, v1.4h, v2.h[7]",
			wantErr: false,
		},
		{
			name: "fmlal2	v0.2s, v1.2h, v2.h[7]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x88, 0xb2, 0x2f}),
				address:          0,
			},
			want:    "fmlal2\tv0.2s, v1.2h, v2.h[7]",
			wantErr: false,
		},
		{
			name: "fmlsl2	v0.2s, v1.2h, v2.h[7]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xc8, 0xb2, 0x2f}),
				address:          0,
			},
			want:    "fmlsl2\tv0.2s, v1.2h, v2.h[7]",
			wantErr: false,
		},
		{
			name: "fmlal2	v0.4s, v1.4h, v2.h[7]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x88, 0xb2, 0x6f}),
				address:          0,
			},
			want:    "fmlal2\tv0.4s, v1.4h, v2.h[7]",
			wantErr: false,
		},
		{
			name: "fmlsl2	v0.4s, v1.4h, v2.h[7]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xc8, 0xb2, 0x6f}),
				address:          0,
			},
			want:    "fmlsl2\tv0.4s, v1.4h, v2.h[7]",
			wantErr: false,
		},
		{
			name: "fmlal	v0.2s, v1.2h, v2.h[5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x08, 0x92, 0x0f}),
				address:          0,
			},
			want:    "fmlal\tv0.2s, v1.2h, v2.h[5]",
			wantErr: false,
		},
		{
			name: "fmlsl	v0.2s, v1.2h, v2.h[5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x48, 0x92, 0x0f}),
				address:          0,
			},
			want:    "fmlsl\tv0.2s, v1.2h, v2.h[5]",
			wantErr: false,
		},
		{
			name: "fmlal	v0.4s, v1.4h, v2.h[5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x08, 0x92, 0x4f}),
				address:          0,
			},
			want:    "fmlal\tv0.4s, v1.4h, v2.h[5]",
			wantErr: false,
		},
		{
			name: "fmlsl	v0.4s, v1.4h, v2.h[5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x48, 0x92, 0x4f}),
				address:          0,
			},
			want:    "fmlsl\tv0.4s, v1.4h, v2.h[5]",
			wantErr: false,
		},
		{
			name: "fmlal2	v0.2s, v1.2h, v2.h[5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x88, 0x92, 0x2f}),
				address:          0,
			},
			want:    "fmlal2\tv0.2s, v1.2h, v2.h[5]",
			wantErr: false,
		},
		{
			name: "fmlsl2	v0.2s, v1.2h, v2.h[5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xc8, 0x92, 0x2f}),
				address:          0,
			},
			want:    "fmlsl2\tv0.2s, v1.2h, v2.h[5]",
			wantErr: false,
		},
		{
			name: "fmlal2	v0.4s, v1.4h, v2.h[5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x88, 0x92, 0x6f}),
				address:          0,
			},
			want:    "fmlal2\tv0.4s, v1.4h, v2.h[5]",
			wantErr: false,
		},
		{
			name: "fmlsl2	v0.4s, v1.4h, v2.h[5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xc8, 0x92, 0x6f}),
				address:          0,
			},
			want:    "fmlsl2\tv0.4s, v1.4h, v2.h[5]",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decompose(tt.args.address, tt.args.instructionValue, &results)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decompose() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got.Disassembly, tt.want) {
				t.Errorf("Disassembly = %v, want %v", got.Disassembly, tt.want)
			}
		})
	}
}

func Test_decompose_v8_3a(t *testing.T) {
	var results [1024]byte
	type args struct {
		instructionValue uint32
		address          uint64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		//
		// llvm/test/MC/AArch64/armv8.3a-complex_nofp16.s
		//
		{
			name: "fcmla	v0.2s, v1.2s, v2.2s, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xc4, 0x82, 0x2e}),
				address:          0,
			},
			want:    "fcmla\tv0.2s, v1.2s, v2.2s, #0",
			wantErr: false,
		},
		{
			name: "fcmla	v0.4s, v1.4s, v2.4s, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xc4, 0x82, 0x6e}),
				address:          0,
			},
			want:    "fcmla\tv0.4s, v1.4s, v2.4s, #0",
			wantErr: false,
		},
		{
			name: "fcmla	v0.2d, v1.2d, v2.2d, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xc4, 0xc2, 0x6e}),
				address:          0,
			},
			want:    "fcmla\tv0.2d, v1.2d, v2.2d, #0",
			wantErr: false,
		},
		{
			name: "fcmla	v0.2s, v1.2s, v2.2s, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xc4, 0x82, 0x2e}),
				address:          0,
			},
			want:    "fcmla\tv0.2s, v1.2s, v2.2s, #0",
			wantErr: false,
		},
		{
			name: "fcmla	v0.2s, v1.2s, v2.2s, #90",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xcc, 0x82, 0x2e}),
				address:          0,
			},
			want:    "fcmla\tv0.2s, v1.2s, v2.2s, #0x5a",
			wantErr: false,
		},
		{
			name: "fcmla	v0.2s, v1.2s, v2.2s, #180",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xd4, 0x82, 0x2e}),
				address:          0,
			},
			want:    "fcmla\tv0.2s, v1.2s, v2.2s, #0xb4",
			wantErr: false,
		},
		{
			name: "fcmla	v0.2s, v1.2s, v2.2s, #270",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xdc, 0x82, 0x2e}),
				address:          0,
			},
			want:    "fcmla\tv0.2s, v1.2s, v2.2s, #0x10e",
			wantErr: false,
		},
		{
			name: "fcadd	v0.2s, v1.2s, v2.2s, #90",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xe4, 0x82, 0x2e}),
				address:          0,
			},
			want:    "fcadd\tv0.2s, v1.2s, v2.2s, #0x5a",
			wantErr: false,
		},
		{
			name: "fcadd	v0.4s, v1.4s, v2.4s, #90",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xe4, 0x82, 0x6e}),
				address:          0,
			},
			want:    "fcadd\tv0.4s, v1.4s, v2.4s, #0x5a",
			wantErr: false,
		},
		{
			name: "fcadd	v0.2d, v1.2d, v2.2d, #90",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xe4, 0xc2, 0x6e}),
				address:          0,
			},
			want:    "fcadd\tv0.2d, v1.2d, v2.2d, #0x5a",
			wantErr: false,
		},
		{
			name: "fcadd	v0.2s, v1.2s, v2.2s, #90",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xe4, 0x82, 0x2e}),
				address:          0,
			},
			want:    "fcadd\tv0.2s, v1.2s, v2.2s, #0x5a",
			wantErr: false,
		},
		{
			name: "fcadd	v0.2s, v1.2s, v2.2s, #270",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xf4, 0x82, 0x2e}),
				address:          0,
			},
			want:    "fcadd\tv0.2s, v1.2s, v2.2s, #0x10e",
			wantErr: false,
		},
		{
			name: "fcmla	v0.4s, v1.4s, v2.s[0], #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x10, 0x82, 0x6f}),
				address:          0,
			},
			want:    "fcmla\tv0.4s, v1.4s, v2.s[0], #0",
			wantErr: false,
		},
		{
			name: "fcmla	v0.4s, v1.4s, v2.s[0], #90",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x30, 0x82, 0x6f}),
				address:          0,
			},
			want:    "fcmla\tv0.4s, v1.4s, v2.s[0], #0x5a",
			wantErr: false,
		},
		{
			name: "fcmla	v0.4s, v1.4s, v2.s[0], #180",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x50, 0x82, 0x6f}),
				address:          0,
			},
			want:    "fcmla\tv0.4s, v1.4s, v2.s[0], #0xb4",
			wantErr: false,
		},
		{
			name: "fcmla	v0.4s, v1.4s, v2.s[0], #270",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x70, 0x82, 0x6f}),
				address:          0,
			},
			want:    "fcmla\tv0.4s, v1.4s, v2.s[0], #0x10e",
			wantErr: false,
		},
		{
			name: "fcmla	v0.4s, v1.4s, v2.s[1], #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x18, 0x82, 0x6f}),
				address:          0,
			},
			want:    "fcmla\tv0.4s, v1.4s, v2.s[1], #0",
			wantErr: false,
		},

		//
		// llvm/test/MC/AArch64/armv8.3a-complex.s
		//
		{
			name: "fcmla	v0.4h, v1.4h, v2.4h, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xc4, 0x42, 0x2e}),
				address:          0,
			},
			want:    "fcmla\tv0.4h, v1.4h, v2.4h, #0",
			wantErr: false,
		},
		{
			name: "fcmla	v0.8h, v1.8h, v2.8h, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xc4, 0x42, 0x6e}),
				address:          0,
			},
			want:    "fcmla\tv0.8h, v1.8h, v2.8h, #0",
			wantErr: false,
		},
		{
			name: "fcmla	v0.2s, v1.2s, v2.2s, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xc4, 0x82, 0x2e}),
				address:          0,
			},
			want:    "fcmla\tv0.2s, v1.2s, v2.2s, #0",
			wantErr: false,
		},
		{
			name: "fcmla	v0.4s, v1.4s, v2.4s, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xc4, 0x82, 0x6e}),
				address:          0,
			},
			want:    "fcmla\tv0.4s, v1.4s, v2.4s, #0",
			wantErr: false,
		},
		{
			name: "fcmla	v0.2d, v1.2d, v2.2d, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xc4, 0xc2, 0x6e}),
				address:          0,
			},
			want:    "fcmla\tv0.2d, v1.2d, v2.2d, #0",
			wantErr: false,
		},
		{
			name: "fcmla	v0.2s, v1.2s, v2.2s, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xc4, 0x82, 0x2e}),
				address:          0,
			},
			want:    "fcmla\tv0.2s, v1.2s, v2.2s, #0",
			wantErr: false,
		},
		{
			name: "fcmla	v0.2s, v1.2s, v2.2s, #90",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xcc, 0x82, 0x2e}),
				address:          0,
			},
			want:    "fcmla\tv0.2s, v1.2s, v2.2s, #0x5a",
			wantErr: false,
		},
		{
			name: "fcmla	v0.2s, v1.2s, v2.2s, #180",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xd4, 0x82, 0x2e}),
				address:          0,
			},
			want:    "fcmla\tv0.2s, v1.2s, v2.2s, #0xb4",
			wantErr: false,
		},
		{
			name: "fcmla	v0.2s, v1.2s, v2.2s, #270",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xdc, 0x82, 0x2e}),
				address:          0,
			},
			want:    "fcmla\tv0.2s, v1.2s, v2.2s, #0x10e",
			wantErr: false,
		},
		{
			name: "fcadd	v0.4h, v1.4h, v2.4h, #90",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xe4, 0x42, 0x2e}),
				address:          0,
			},
			want:    "fcadd\tv0.4h, v1.4h, v2.4h, #0x5a",
			wantErr: false,
		},
		{
			name: "fcadd	v0.8h, v1.8h, v2.8h, #90",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xe4, 0x42, 0x6e}),
				address:          0,
			},
			want:    "fcadd\tv0.8h, v1.8h, v2.8h, #0x5a",
			wantErr: false,
		},
		{
			name: "fcadd	v0.2s, v1.2s, v2.2s, #90",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xe4, 0x82, 0x2e}),
				address:          0,
			},
			want:    "fcadd\tv0.2s, v1.2s, v2.2s, #0x5a",
			wantErr: false,
		},
		{
			name: "fcadd	v0.4s, v1.4s, v2.4s, #90",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xe4, 0x82, 0x6e}),
				address:          0,
			},
			want:    "fcadd\tv0.4s, v1.4s, v2.4s, #0x5a",
			wantErr: false,
		},
		{
			name: "fcadd	v0.2d, v1.2d, v2.2d, #90",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xe4, 0xc2, 0x6e}),
				address:          0,
			},
			want:    "fcadd\tv0.2d, v1.2d, v2.2d, #0x5a",
			wantErr: false,
		},
		{
			name: "fcadd	v0.2s, v1.2s, v2.2s, #90",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xe4, 0x82, 0x2e}),
				address:          0,
			},
			want:    "fcadd\tv0.2s, v1.2s, v2.2s, #0x5a",
			wantErr: false,
		},
		{
			name: "fcadd	v0.2s, v1.2s, v2.2s, #270",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xf4, 0x82, 0x2e}),
				address:          0,
			},
			want:    "fcadd\tv0.2s, v1.2s, v2.2s, #0x10e",
			wantErr: false,
		},
		{
			name: "fcmla	v0.4h, v1.4h, v2.h[0], #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x10, 0x42, 0x2f}),
				address:          0,
			},
			want:    "fcmla\tv0.4h, v1.4h, v2.h[0], #0",
			wantErr: false,
		},
		{
			name: "fcmla	v0.8h, v1.8h, v2.h[0], #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x10, 0x42, 0x6f}),
				address:          0,
			},
			want:    "fcmla\tv0.8h, v1.8h, v2.h[0], #0",
			wantErr: false,
		},
		{
			name: "fcmla	v0.4s, v1.4s, v2.s[0], #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x10, 0x82, 0x6f}),
				address:          0,
			},
			want:    "fcmla\tv0.4s, v1.4s, v2.s[0], #0",
			wantErr: false,
		},
		{
			name: "fcmla	v0.4s, v1.4s, v2.s[0], #90",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x30, 0x82, 0x6f}),
				address:          0,
			},
			want:    "fcmla\tv0.4s, v1.4s, v2.s[0], #0x5a",
			wantErr: false,
		},
		{
			name: "fcmla	v0.4s, v1.4s, v2.s[0], #180",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x50, 0x82, 0x6f}),
				address:          0,
			},
			want:    "fcmla\tv0.4s, v1.4s, v2.s[0], #0xb4",
			wantErr: false,
		},
		{
			name: "fcmla	v0.4s, v1.4s, v2.s[0], #270",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x70, 0x82, 0x6f}),
				address:          0,
			},
			want:    "fcmla\tv0.4s, v1.4s, v2.s[0], #0x10e",
			wantErr: false,
		},
		{
			name: "fcmla	v0.4h, v1.4h, v2.h[1], #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x10, 0x62, 0x2f}),
				address:          0,
			},
			want:    "fcmla\tv0.4h, v1.4h, v2.h[1], #0",
			wantErr: false,
		},
		{
			name: "fcmla	v0.8h, v1.8h, v2.h[3], #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x18, 0x62, 0x6f}),
				address:          0,
			},
			want:    "fcmla\tv0.8h, v1.8h, v2.h[3], #0",
			wantErr: false,
		},
		{
			name: "fcmla	v0.4s, v1.4s, v2.s[1], #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x18, 0x82, 0x6f}),
				address:          0,
			},
			want:    "fcmla\tv0.4s, v1.4s, v2.s[1], #0",
			wantErr: false,
		},

		//
		// llvm/test/MC/AArch64/armv8.3a-ID_ISAR6_EL1.s
		//
		{
			name: "mrs	x0, id_isar6_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x02, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, id_isar6_el1",
			wantErr: false,
		},

		//
		// llvm/test/MC/AArch64/armv8.3a-js.s
		//
		{
			name: "fjcvtzs	w0, d0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x00, 0x7e, 0x1e}),
				address:          0,
			},
			want:    "fjcvtzs\tw0, d0",
			wantErr: false,
		},

		// llvm/test/MC/AArch64/armv8.3a-rcpc.s
		{
			name: "ldaprb	w0, [x0]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0xc0, 0xbf, 0x38}),
				address:          0,
			},
			want:    "ldaprb\tw0, [x0]",
			wantErr: false,
		},
		{
			name: "ldaprh	w0, [x17]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xc2, 0xbf, 0x78}),
				address:          0,
			},
			want:    "ldaprh\tw0, [x17]",
			wantErr: false,
		},
		{
			name: "ldapr	w0, [x1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xc0, 0xbf, 0xb8}),
				address:          0,
			},
			want:    "ldapr\tw0, [x1]",
			wantErr: false,
		},
		{
			name: "ldapr	x0, [x0]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0xc0, 0xbf, 0xf8}),
				address:          0,
			},
			want:    "ldapr\tx0, [x0]",
			wantErr: false,
		},
		{
			name: "ldapr	w18, [x0]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x12, 0xc0, 0xbf, 0xb8}),
				address:          0,
			},
			want:    "ldapr\tw18, [x0]",
			wantErr: false,
		},
		{
			name: "ldapr	x15, [x0]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0f, 0xc0, 0xbf, 0xf8}),
				address:          0,
			},
			want:    "ldapr\tx15, [x0]",
			wantErr: false,
		},

		// llvm/test/MC/AArch64/armv8.3a-signed-pointer.s
		{
			name: "mrs	x0, apiakeylo_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x21, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, apiakeylo_el1",
			wantErr: false,
		},
		{
			name: "mrs	x0, apiakeyhi_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x21, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, apiakeyhi_el1",
			wantErr: false,
		},
		{
			name: "mrs	x0, apibkeylo_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0x21, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, apibkeylo_el1",
			wantErr: false,
		},
		{
			name: "mrs	x0, apibkeyhi_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x60, 0x21, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, apibkeyhi_el1",
			wantErr: false,
		},
		{
			name: "mrs	x0, apdakeylo_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x22, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, apdakeylo_el1",
			wantErr: false,
		},
		{
			name: "mrs	x0, apdakeyhi_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x22, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, apdakeyhi_el1",
			wantErr: false,
		},
		{
			name: "mrs	x0, apdbkeylo_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0x22, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, apdbkeylo_el1",
			wantErr: false,
		},
		{
			name: "mrs	x0, apdbkeyhi_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x60, 0x22, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, apdbkeyhi_el1",
			wantErr: false,
		},
		{
			name: "mrs	x0, apgakeylo_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x23, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, apgakeylo_el1",
			wantErr: false,
		},
		{
			name: "mrs	x0, apgakeyhi_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x23, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, apgakeyhi_el1",
			wantErr: false,
		},
		{
			name: "msr	apiakeylo_el1, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x21, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tapiakeylo_el1, x0",
			wantErr: false,
		},
		{
			name: "msr	apiakeyhi_el1, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x21, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tapiakeyhi_el1, x0",
			wantErr: false,
		},
		{
			name: "msr	apibkeylo_el1, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0x21, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tapibkeylo_el1, x0",
			wantErr: false,
		},
		{
			name: "msr	apibkeyhi_el1, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x60, 0x21, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tapibkeyhi_el1, x0",
			wantErr: false,
		},
		{
			name: "msr	apdakeylo_el1, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x22, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tapdakeylo_el1, x0",
			wantErr: false,
		},
		{
			name: "msr	apdakeyhi_el1, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x22, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tapdakeyhi_el1, x0",
			wantErr: false,
		},
		{
			name: "msr	apdbkeylo_el1, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0x22, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tapdbkeylo_el1, x0",
			wantErr: false,
		},
		{
			name: "msr	apdbkeyhi_el1, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x60, 0x22, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tapdbkeyhi_el1, x0",
			wantErr: false,
		},
		{
			name: "msr	apgakeylo_el1, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x23, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tapgakeylo_el1, x0",
			wantErr: false,
		},
		{
			name: "msr	apgakeyhi_el1, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x23, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tapgakeyhi_el1, x0",
			wantErr: false,
		},
		{
			name: "paciasp",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0x23, 0x03, 0xd5}),
				address:          0,
			},
			want:    "paciasp",
			wantErr: false,
		},
		{
			name: "autiasp",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x23, 0x03, 0xd5}),
				address:          0,
			},
			want:    "autiasp",
			wantErr: false,
		},
		{
			name: "paciaz",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x23, 0x03, 0xd5}),
				address:          0,
			},
			want:    "paciaz",
			wantErr: false,
		},
		{
			name: "autiaz",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x23, 0x03, 0xd5}),
				address:          0,
			},
			want:    "autiaz",
			wantErr: false,
		},
		{
			name: "pacia1716",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x21, 0x03, 0xd5}),
				address:          0,
			},
			want:    "pacia1716",
			wantErr: false,
		},
		{
			name: "autia1716",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x21, 0x03, 0xd5}),
				address:          0,
			},
			want:    "autia1716",
			wantErr: false,
		},
		{
			name: "pacibsp",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x23, 0x03, 0xd5}),
				address:          0,
			},
			want:    "pacibsp",
			wantErr: false,
		},
		{
			name: "autibsp",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x23, 0x03, 0xd5}),
				address:          0,
			},
			want:    "autibsp",
			wantErr: false,
		},
		{
			name: "pacibz",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x23, 0x03, 0xd5}),
				address:          0,
			},
			want:    "pacibz",
			wantErr: false,
		},
		{
			name: "autibz",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0x23, 0x03, 0xd5}),
				address:          0,
			},
			want:    "autibz",
			wantErr: false,
		},
		{
			name: "pacib1716",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x21, 0x03, 0xd5}),
				address:          0,
			},
			want:    "pacib1716",
			wantErr: false,
		},
		{
			name: "autib1716",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0x21, 0x03, 0xd5}),
				address:          0,
			},
			want:    "autib1716",
			wantErr: false,
		},
		{
			name: "xpaclri",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x20, 0x03, 0xd5}),
				address:          0,
			},
			want:    "xpaclri",
			wantErr: false,
		},
		{
			name: "pacia	x0, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x00, 0xc1, 0xda}),
				address:          0,
			},
			want:    "pacia\tx0, x1",
			wantErr: false,
		},
		{
			name: "autia	x0, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x10, 0xc1, 0xda}),
				address:          0,
			},
			want:    "autia\tx0, x1",
			wantErr: false,
		},
		{
			name: "pacda	x0, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x08, 0xc1, 0xda}),
				address:          0,
			},
			want:    "pacda\tx0, x1",
			wantErr: false,
		},
		{
			name: "autda	x0, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x18, 0xc1, 0xda}),
				address:          0,
			},
			want:    "autda\tx0, x1",
			wantErr: false,
		},
		{
			name: "pacib	x0, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x04, 0xc1, 0xda}),
				address:          0,
			},
			want:    "pacib\tx0, x1",
			wantErr: false,
		},
		{
			name: "autib	x0, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x14, 0xc1, 0xda}),
				address:          0,
			},
			want:    "autib\tx0, x1",
			wantErr: false,
		},
		{
			name: "pacdb	x0, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x0c, 0xc1, 0xda}),
				address:          0,
			},
			want:    "pacdb\tx0, x1",
			wantErr: false,
		},
		{
			name: "autdb	x0, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x1c, 0xc1, 0xda}),
				address:          0,
			},
			want:    "autdb\tx0, x1",
			wantErr: false,
		},
		{
			name: "pacga	x0, x1, x2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x30, 0xc2, 0x9a}),
				address:          0,
			},
			want:    "pacga\tx0, x1, x2",
			wantErr: false,
		},
		{
			name: "paciza	x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x23, 0xc1, 0xda}),
				address:          0,
			},
			want:    "paciza\tx0",
			wantErr: false,
		},
		{
			name: "autiza	x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x33, 0xc1, 0xda}),
				address:          0,
			},
			want:    "autiza\tx0",
			wantErr: false,
		},
		{
			name: "pacdza	x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x2b, 0xc1, 0xda}),
				address:          0,
			},
			want:    "pacdza\tx0",
			wantErr: false,
		},
		{
			name: "autdza	x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x3b, 0xc1, 0xda}),
				address:          0,
			},
			want:    "autdza\tx0",
			wantErr: false,
		},
		{
			name: "pacizb	x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x27, 0xc1, 0xda}),
				address:          0,
			},
			want:    "pacizb\tx0",
			wantErr: false,
		},
		{
			name: "autizb	x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x37, 0xc1, 0xda}),
				address:          0,
			},
			want:    "autizb\tx0",
			wantErr: false,
		},
		{
			name: "pacdzb	x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x2f, 0xc1, 0xda}),
				address:          0,
			},
			want:    "pacdzb\tx0",
			wantErr: false,
		},
		{
			name: "autdzb	x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x3f, 0xc1, 0xda}),
				address:          0,
			},
			want:    "autdzb\tx0",
			wantErr: false,
		},
		{
			name: "xpaci	x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x43, 0xc1, 0xda}),
				address:          0,
			},
			want:    "xpaci\tx0",
			wantErr: false,
		},
		{
			name: "xpacd	x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x47, 0xc1, 0xda}),
				address:          0,
			},
			want:    "xpacd\tx0",
			wantErr: false,
		},
		{
			name: "braa	x0, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x01, 0x08, 0x1f, 0xd7}),
				address:          0,
			},
			want:    "braa\tx0, x1",
			wantErr: false,
		},
		{
			name: "brab	x0, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x01, 0x0c, 0x1f, 0xd7}),
				address:          0,
			},
			want:    "brab\tx0, x1",
			wantErr: false,
		},
		{
			name: "blraa	x0, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x01, 0x08, 0x3f, 0xd7}),
				address:          0,
			},
			want:    "blraa\tx0, x1",
			wantErr: false,
		},
		{
			name: "blrab	x0, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x01, 0x0c, 0x3f, 0xd7}),
				address:          0,
			},
			want:    "blrab\tx0, x1",
			wantErr: false,
		},
		{
			name: "braaz	x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x08, 0x1f, 0xd6}),
				address:          0,
			},
			want:    "braaz\tx0",
			wantErr: false,
		},
		{
			name: "brabz	x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x0c, 0x1f, 0xd6}),
				address:          0,
			},
			want:    "brabz\tx0",
			wantErr: false,
		},
		{
			name: "blraaz	x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x08, 0x3f, 0xd6}),
				address:          0,
			},
			want:    "blraaz\tx0",
			wantErr: false,
		},
		{
			name: "blrabz	x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x0c, 0x3f, 0xd6}),
				address:          0,
			},
			want:    "blrabz\tx0",
			wantErr: false,
		},
		{
			name: "retaa",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x0b, 0x5f, 0xd6}),
				address:          0,
			},
			want:    "retaa",
			wantErr: false,
		},
		{
			name: "retab",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x0f, 0x5f, 0xd6}),
				address:          0,
			},
			want:    "retab",
			wantErr: false,
		},
		{
			name: "eretaa",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x0b, 0x9f, 0xd6}),
				address:          0,
			},
			want:    "eretaa",
			wantErr: false,
		},
		{
			name: "eretab",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x0f, 0x9f, 0xd6}),
				address:          0,
			},
			want:    "eretab",
			wantErr: false,
		},
		{
			name: "ldraa	x0, [x1, #4088]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xf4, 0x3f, 0xf8}),
				address:          0,
			},
			want:    "ldraa\tx0, [x1, #0xff8]",
			wantErr: false,
		},
		{
			name: "ldraa	x0, [x1, #-4096]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x04, 0x60, 0xf8}),
				address:          0,
			},
			want:    "ldraa\tx0, [x1, #-0x1000]",
			wantErr: false,
		},
		{
			name: "ldrab	x0, [x1, #4088]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xf4, 0xbf, 0xf8}),
				address:          0,
			},
			want:    "ldrab\tx0, [x1, #0xff8]",
			wantErr: false,
		},
		{
			name: "ldrab	x0, [x1, #-4096]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x04, 0xe0, 0xf8}),
				address:          0,
			},
			want:    "ldrab\tx0, [x1, #-0x1000]",
			wantErr: false,
		},
		{
			name: "ldraa	x0, [x1, #4088]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xfc, 0x3f, 0xf8}),
				address:          0,
			},
			want:    "ldraa\tx0, [x1, #0xff8]!",
			wantErr: false,
		},
		{
			name: "ldraa	x0, [x1, #-4096]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x0c, 0x60, 0xf8}),
				address:          0,
			},
			want:    "ldraa\tx0, [x1, #-0x1000]!",
			wantErr: false,
		},
		{
			name: "ldrab	x0, [x1, #4088]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xfc, 0xbf, 0xf8}),
				address:          0,
			},
			want:    "ldrab\tx0, [x1, #0xff8]!",
			wantErr: false,
		},
		{
			name: "ldrab	x0, [x1, #-4096]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x0c, 0xe0, 0xf8}),
				address:          0,
			},
			want:    "ldrab\tx0, [x1, #-0x1000]!",
			wantErr: false,
		},
		{
			name: "ldraa	x0, [x1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x04, 0x20, 0xf8}),
				address:          0,
			},
			want:    "ldraa\tx0, [x1]",
			wantErr: false,
		},
		{
			name: "ldrab	x0, [x1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x04, 0xa0, 0xf8}),
				address:          0,
			},
			want:    "ldrab\tx0, [x1]",
			wantErr: false,
		},
		{
			name: "ldraa	x0, [x1, #0]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x0c, 0x20, 0xf8}),
				address:          0,
			},
			want:    "ldraa\tx0, [x1, #0]!",
			wantErr: false,
		},
		{
			name: "ldrab	x0, [x1, #0]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x0c, 0xa0, 0xf8}),
				address:          0,
			},
			want:    "ldrab\tx0, [x1, #0]!",
			wantErr: false,
		},
		{
			name: "ldraa	xzr, [sp, #-4096]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x0f, 0x60, 0xf8}),
				address:          0,
			},
			want:    "ldraa\txzr, [sp, #-0x1000]!",
			wantErr: false,
		},
		{
			name: "ldrab	xzr, [sp, #-4096]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x0f, 0xe0, 0xf8}),
				address:          0,
			},
			want:    "ldrab\txzr, [sp, #-0x1000]!",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decompose(tt.args.address, tt.args.instructionValue, &results)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decompose() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got.Disassembly, tt.want) {
				t.Errorf("Disassembly = %v, want %v", got.Disassembly, tt.want)
			}
		})
	}
}

func Test_decompose_v8_4a(t *testing.T) {
	var results [1024]byte
	type args struct {
		instructionValue uint32
		address          uint64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// llvm/test/MC/AArch64/armv8.4a-ldst.s
		{
			name: "stlurb	w1, [x10]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x01, 0x00, 0x19}),
				address:          0,
			},
			want:    "stlurb\tw1, [x10]",
			wantErr: false,
		},
		{
			name: "stlurb	w1, [x10, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x01, 0x10, 0x19}),
				address:          0,
			},
			want:    "stlurb\tw1, [x10, #-0x100]",
			wantErr: false,
		},
		{
			name: "stlurb	w2, [x11, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xf1, 0x0f, 0x19}),
				address:          0,
			},
			want:    "stlurb\tw2, [x11, #0xff]",
			wantErr: false,
		},
		{
			name: "stlurb	w3, [sp, #-3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0xd3, 0x1f, 0x19}),
				address:          0,
			},
			want:    "stlurb\tw3, [sp, #-0x3]",
			wantErr: false,
		},
		{
			name: "ldapurb	wzr, [x12]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x01, 0x40, 0x19}),
				address:          0,
			},
			want:    "ldapurb\twzr, [x12]",
			wantErr: false,
		},
		{
			name: "ldapurb	w4, [x12]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x84, 0x01, 0x40, 0x19}),
				address:          0,
			},
			want:    "ldapurb\tw4, [x12]",
			wantErr: false,
		},
		{
			name: "ldapurb	w4, [x12, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x84, 0x01, 0x50, 0x19}),
				address:          0,
			},
			want:    "ldapurb\tw4, [x12, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldapurb	w5, [x13, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa5, 0xf1, 0x4f, 0x19}),
				address:          0,
			},
			want:    "ldapurb\tw5, [x13, #0xff]",
			wantErr: false,
		},
		{
			name: "ldapurb	w6, [sp, #-2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe6, 0xe3, 0x5f, 0x19}),
				address:          0,
			},
			want:    "ldapurb\tw6, [sp, #-0x2]",
			wantErr: false,
		},
		{
			name: "ldapursb	w7, [x14]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc7, 0x01, 0xc0, 0x19}),
				address:          0,
			},
			want:    "ldapursb\tw7, [x14]",
			wantErr: false,
		},
		{
			name: "ldapursb	w7, [x14, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc7, 0x01, 0xd0, 0x19}),
				address:          0,
			},
			want:    "ldapursb\tw7, [x14, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldapursb	w8, [x15, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe8, 0xf1, 0xcf, 0x19}),
				address:          0,
			},
			want:    "ldapursb\tw8, [x15, #0xff]",
			wantErr: false,
		},
		{
			name: "ldapursb	w9, [sp, #-1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xf3, 0xdf, 0x19}),
				address:          0,
			},
			want:    "ldapursb\tw9, [sp, #-0x1]",
			wantErr: false,
		},
		{
			name: "ldapursb	x0, [x16]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x02, 0x80, 0x19}),
				address:          0,
			},
			want:    "ldapursb\tx0, [x16]",
			wantErr: false,
		},
		{
			name: "ldapursb	x0, [x16, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x02, 0x90, 0x19}),
				address:          0,
			},
			want:    "ldapursb\tx0, [x16, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldapursb	x1, [x17, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x21, 0xf2, 0x8f, 0x19}),
				address:          0,
			},
			want:    "ldapursb\tx1, [x17, #0xff]",
			wantErr: false,
		},
		{
			name: "ldapursb	x2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe2, 0x03, 0x80, 0x19}),
				address:          0,
			},
			want:    "ldapursb\tx2, [sp]",
			wantErr: false,
		},
		{
			name: "ldapursb	x2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe2, 0x03, 0x80, 0x19}),
				address:          0,
			},
			want:    "ldapursb\tx2, [sp]",
			wantErr: false,
		},
		{
			name: "stlurh	w10, [x18]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4a, 0x02, 0x00, 0x59}),
				address:          0,
			},
			want:    "stlurh\tw10, [x18]",
			wantErr: false,
		},
		{
			name: "stlurh	w10, [x18, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4a, 0x02, 0x10, 0x59}),
				address:          0,
			},
			want:    "stlurh\tw10, [x18, #-0x100]",
			wantErr: false,
		},
		{
			name: "stlurh	w11, [x19, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6b, 0xf2, 0x0f, 0x59}),
				address:          0,
			},
			want:    "stlurh\tw11, [x19, #0xff]",
			wantErr: false,
		},
		{
			name: "stlurh	w12, [sp, #1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x13, 0x00, 0x59}),
				address:          0,
			},
			want:    "stlurh\tw12, [sp, #0x1]",
			wantErr: false,
		},
		{
			name: "ldapurh	w13, [x20]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8d, 0x02, 0x40, 0x59}),
				address:          0,
			},
			want:    "ldapurh\tw13, [x20]",
			wantErr: false,
		},
		{
			name: "ldapurh	w13, [x20, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8d, 0x02, 0x50, 0x59}),
				address:          0,
			},
			want:    "ldapurh\tw13, [x20, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldapurh	w14, [x21, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xae, 0xf2, 0x4f, 0x59}),
				address:          0,
			},
			want:    "ldapurh\tw14, [x21, #0xff]",
			wantErr: false,
		},
		{
			name: "ldapurh	w15, [sp, #2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xef, 0x23, 0x40, 0x59}),
				address:          0,
			},
			want:    "ldapurh\tw15, [sp, #0x2]",
			wantErr: false,
		},
		{
			name: "ldapursh	w16, [x22]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd0, 0x02, 0xc0, 0x59}),
				address:          0,
			},
			want:    "ldapursh\tw16, [x22]",
			wantErr: false,
		},
		{
			name: "ldapursh	w16, [x22, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd0, 0x02, 0xd0, 0x59}),
				address:          0,
			},
			want:    "ldapursh\tw16, [x22, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldapursh	w17, [x23, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf1, 0xf2, 0xcf, 0x59}),
				address:          0,
			},
			want:    "ldapursh\tw17, [x23, #0xff]",
			wantErr: false,
		},
		{
			name: "ldapursh	w18, [sp, #3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf2, 0x33, 0xc0, 0x59}),
				address:          0,
			},
			want:    "ldapursh\tw18, [sp, #0x3]",
			wantErr: false,
		},
		{
			name: "ldapursh	x3, [x24]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x03, 0x03, 0x80, 0x59}),
				address:          0,
			},
			want:    "ldapursh\tx3, [x24]",
			wantErr: false,
		},
		{
			name: "ldapursh	x3, [x24, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x03, 0x03, 0x90, 0x59}),
				address:          0,
			},
			want:    "ldapursh\tx3, [x24, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldapursh	x4, [x25, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x24, 0xf3, 0x8f, 0x59}),
				address:          0,
			},
			want:    "ldapursh\tx4, [x25, #0xff]",
			wantErr: false,
		},
		{
			name: "ldapursh	x5, [sp, #4]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe5, 0x43, 0x80, 0x59}),
				address:          0,
			},
			want:    "ldapursh\tx5, [sp, #0x4]",
			wantErr: false,
		},
		{
			name: "stlur	w19, [x26]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x53, 0x03, 0x00, 0x99}),
				address:          0,
			},
			want:    "stlur\tw19, [x26]",
			wantErr: false,
		},
		{
			name: "stlur	w19, [x26, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x53, 0x03, 0x10, 0x99}),
				address:          0,
			},
			want:    "stlur\tw19, [x26, #-0x100]",
			wantErr: false,
		},
		{
			name: "stlur	w20, [x27, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x74, 0xf3, 0x0f, 0x99}),
				address:          0,
			},
			want:    "stlur\tw20, [x27, #0xff]",
			wantErr: false,
		},
		{
			name: "stlur	w21, [sp, #5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf5, 0x53, 0x00, 0x99}),
				address:          0,
			},
			want:    "stlur\tw21, [sp, #0x5]",
			wantErr: false,
		},
		{
			name: "ldapur	w22, [x28]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x96, 0x03, 0x40, 0x99}),
				address:          0,
			},
			want:    "ldapur\tw22, [x28]",
			wantErr: false,
		},
		{
			name: "ldapur	w22, [x28, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x96, 0x03, 0x50, 0x99}),
				address:          0,
			},
			want:    "ldapur\tw22, [x28, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldapur	w23, [x29, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb7, 0xf3, 0x4f, 0x99}),
				address:          0,
			},
			want:    "ldapur\tw23, [fp, #0xff]",
			wantErr: false,
		},
		{
			name: "ldapur	w24, [sp, #6]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf8, 0x63, 0x40, 0x99}),
				address:          0,
			},
			want:    "ldapur\tw24, [sp, #0x6]",
			wantErr: false,
		},
		{
			name: "ldapursw	x6, [x30]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc6, 0x03, 0x80, 0x99}),
				address:          0,
			},
			want:    "ldapursw\tx6, [lr]",
			wantErr: false,
		},
		{
			name: "ldapursw	x6, [x30, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc6, 0x03, 0x90, 0x99}),
				address:          0,
			},
			want:    "ldapursw\tx6, [lr, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldapursw	x7, [x0, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x07, 0xf0, 0x8f, 0x99}),
				address:          0,
			},
			want:    "ldapursw\tx7, [x0, #0xff]",
			wantErr: false,
		},
		{
			name: "ldapursw	x8, [sp, #7]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe8, 0x73, 0x80, 0x99}),
				address:          0,
			},
			want:    "ldapursw\tx8, [sp, #0x7]",
			wantErr: false,
		},
		{
			name: "stlur	x9, [x1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x00, 0x00, 0xd9}),
				address:          0,
			},
			want:    "stlur\tx9, [x1]",
			wantErr: false,
		},
		{
			name: "stlur	x9, [x1, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x00, 0x10, 0xd9}),
				address:          0,
			},
			want:    "stlur\tx9, [x1, #-0x100]",
			wantErr: false,
		},
		{
			name: "stlur	x10, [x2, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4a, 0xf0, 0x0f, 0xd9}),
				address:          0,
			},
			want:    "stlur\tx10, [x2, #0xff]",
			wantErr: false,
		},
		{
			name: "stlur	x11, [sp, #8]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xeb, 0x83, 0x00, 0xd9}),
				address:          0,
			},
			want:    "stlur\tx11, [sp, #0x8]",
			wantErr: false,
		},
		{
			name: "ldapur	x12, [x3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6c, 0x00, 0x40, 0xd9}),
				address:          0,
			},
			want:    "ldapur\tx12, [x3]",
			wantErr: false,
		},
		{
			name: "ldapur	x12, [x3, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6c, 0x00, 0x50, 0xd9}),
				address:          0,
			},
			want:    "ldapur\tx12, [x3, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldapur	x13, [x4, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8d, 0xf0, 0x4f, 0xd9}),
				address:          0,
			},
			want:    "ldapur\tx13, [x4, #0xff]",
			wantErr: false,
		},
		{
			name: "ldapur	x14, [sp, #9]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xee, 0x93, 0x40, 0xd9}),
				address:          0,
			},
			want:    "ldapur\tx14, [sp, #0x9]",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decompose(tt.args.address, tt.args.instructionValue, &results)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decompose() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got.Disassembly, tt.want) {
				t.Errorf("Disassembly = %v, want %v", got.Disassembly, tt.want)
			}
		})
	}
}

func Test_decompose_v8_5a(t *testing.T) {
	var results [1024]byte
	type args struct {
		instructionValue uint32
		address          uint64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// // llvm/test/MC/AArch64/armv8.5a-altnzcv.s
		// {
		// 	name: "xaflag",
		// 	args: args{
		// 		instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0x40, 0x00, 0xd5}),
		// 		address:          0,
		// 	},
		// 	want:    "xaflag",
		// 	wantErr: false,
		// },
		// {
		// 	name: "axflag",
		// 	args: args{
		// 		instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x40, 0x00, 0xd5}),
		// 		address:          0,
		// 	},
		// 	want:    "axflag",
		// 	wantErr: false,
		// },
		// llvm/test/MC/AArch64/armv8.5a-bti.s
		{
			name: "bti",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x24, 0x03, 0xd5}),
				address:          0,
			},
			want:    "bti",
			wantErr: false,
		},
		{
			name: "bti	c",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x24, 0x03, 0xd5}),
				address:          0,
			},
			want:    "bti\tc",
			wantErr: false,
		},
		{
			name: "bti	j",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x24, 0x03, 0xd5}),
				address:          0,
			},
			want:    "bti\tj",
			wantErr: false,
		},
		{
			name: "bti	jc",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0x24, 0x03, 0xd5}),
				address:          0,
			},
			want:    "bti\tjc",
			wantErr: false,
		},
		// llvm/test/MC/AArch64/armv8.5a-frint.s
		{
			name: "frint32z	s0, s1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x40, 0x28, 0x1e}),
				address:          0,
			},
			want:    "frint32z\ts0, s1",
			wantErr: false,
		},
		{
			name: "frint32z	d0, d1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x40, 0x68, 0x1e}),
				address:          0,
			},
			want:    "frint32z\td0, d1",
			wantErr: false,
		},
		{
			name: "frint64z	s2, s3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x40, 0x29, 0x1e}),
				address:          0,
			},
			want:    "frint64z\ts2, s3",
			wantErr: false,
		},
		{
			name: "frint64z	d2, d3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x40, 0x69, 0x1e}),
				address:          0,
			},
			want:    "frint64z\td2, d3",
			wantErr: false,
		},
		{
			name: "frint32x	s4, s5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa4, 0xc0, 0x28, 0x1e}),
				address:          0,
			},
			want:    "frint32x\ts4, s5",
			wantErr: false,
		},
		{
			name: "frint32x	d4, d5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa4, 0xc0, 0x68, 0x1e}),
				address:          0,
			},
			want:    "frint32x\td4, d5",
			wantErr: false,
		},
		{
			name: "frint64x	s6, s7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe6, 0xc0, 0x29, 0x1e}),
				address:          0,
			},
			want:    "frint64x\ts6, s7",
			wantErr: false,
		},
		{
			name: "frint64x	d6, d7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe6, 0xc0, 0x69, 0x1e}),
				address:          0,
			},
			want:    "frint64x\td6, d7",
			wantErr: false,
		},
		{
			name: "frint32z	v0.2s, v1.2s",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xe8, 0x21, 0x0e}),
				address:          0,
			},
			want:    "frint32z\tv0.2s, v1.2s",
			wantErr: false,
		},
		{
			name: "frint32z	v0.2d, v1.2d",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xe8, 0x61, 0x4e}),
				address:          0,
			},
			want:    "frint32z\tv0.2d, v1.2d",
			wantErr: false,
		},
		{
			name: "frint32z	v0.4s, v1.4s",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xe8, 0x21, 0x4e}),
				address:          0,
			},
			want:    "frint32z\tv0.4s, v1.4s",
			wantErr: false,
		},
		{
			name: "frint64z	v2.2s, v3.2s",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xf8, 0x21, 0x0e}),
				address:          0,
			},
			want:    "frint64z\tv2.2s, v3.2s",
			wantErr: false,
		},
		{
			name: "frint64z	v2.2d, v3.2d",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xf8, 0x61, 0x4e}),
				address:          0,
			},
			want:    "frint64z\tv2.2d, v3.2d",
			wantErr: false,
		},
		{
			name: "frint64z	v2.4s, v3.4s",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xf8, 0x21, 0x4e}),
				address:          0,
			},
			want:    "frint64z\tv2.4s, v3.4s",
			wantErr: false,
		},
		{
			name: "frint32x	v4.2s, v5.2s",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa4, 0xe8, 0x21, 0x2e}),
				address:          0,
			},
			want:    "frint32x\tv4.2s, v5.2s",
			wantErr: false,
		},
		{
			name: "frint32x	v4.2d, v5.2d",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa4, 0xe8, 0x61, 0x6e}),
				address:          0,
			},
			want:    "frint32x\tv4.2d, v5.2d",
			wantErr: false,
		},
		{
			name: "frint32x	v4.4s, v5.4s",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa4, 0xe8, 0x21, 0x6e}),
				address:          0,
			},
			want:    "frint32x\tv4.4s, v5.4s",
			wantErr: false,
		},
		{
			name: "frint64x	v6.2s, v7.2s",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe6, 0xf8, 0x21, 0x2e}),
				address:          0,
			},
			want:    "frint64x\tv6.2s, v7.2s",
			wantErr: false,
		},
		{
			name: "frint64x	v6.2d, v7.2d",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe6, 0xf8, 0x61, 0x6e}),
				address:          0,
			},
			want:    "frint64x\tv6.2d, v7.2d",
			wantErr: false,
		},
		{
			name: "frint64x	v6.4s, v7.4s",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe6, 0xf8, 0x21, 0x6e}),
				address:          0,
			},
			want:    "frint64x\tv6.4s, v7.4s",
			wantErr: false,
		},
		// llvm/test/MC/AArch64/armv8.5a-persistent-memory.s
		{
			name: "dc	cvadp, x7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x27, 0x7d, 0x0b, 0xd5}),
				address:          0,
			},
			want:    "dc\tCVADP, x7",
			wantErr: false,
		},
		// llvm/test/MC/AArch64/armv8.5a-predres.s
		{
			name: "cfp	rctx, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x80, 0x73, 0x0b, 0xd5}),
				address:          0,
			},
			want:    "cfp\trctx, x0",
			wantErr: false,
		},
		{
			name: "dvp	rctx, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa1, 0x73, 0x0b, 0xd5}),
				address:          0,
			},
			want:    "dvp\trctx, x1",
			wantErr: false,
		},
		{
			name: "cpp	rctx, x2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe2, 0x73, 0x0b, 0xd5}),
				address:          0,
			},
			want:    "cpp\trctx, x2",
			wantErr: false,
		},
		// llvm/test/MC/AArch64/armv8.5a-rand.s
		{
			name: "mrs	x0, s3_3_c2_c4_0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x24, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_3_c2_c4_0",
			wantErr: false,
		},
		{
			name: "mrs	x1, s3_3_c2_c4_1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x21, 0x24, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx1, s3_3_c2_c4_1",
			wantErr: false,
		},
		// llvm/test/MC/AArch64/armv8.5a-sb.s
		{
			name: "sb",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x30, 0x03, 0xd5}),
				address:          0,
			},
			want:    "sb",
			wantErr: false,
		},
		// llvm/test/MC/AArch64/armv8.5a-specrestrict.s
		{
			name: "mrs	x9, id_pfr2_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x03, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_pfr2_el1",
			wantErr: false,
		},
		{
			name: "mrs	x8, scxtnum_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe8, 0xd0, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx8, scxtnum_el0",
			wantErr: false,
		},
		{
			name: "mrs	x7, scxtnum_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe7, 0xd0, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx7, scxtnum_el1",
			wantErr: false,
		},
		{
			name: "mrs	x6, scxtnum_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe6, 0xd0, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx6, scxtnum_el2",
			wantErr: false,
		},
		{
			name: "mrs	x5, scxtnum_el3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe5, 0xd0, 0x3e, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx5, scxtnum_el3",
			wantErr: false,
		},
		{
			name: "mrs	x4, scxtnum_el12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe4, 0xd0, 0x3d, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx4, scxtnum_el12",
			wantErr: false,
		},
		{
			name: "msr	scxtnum_el0, x8",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe8, 0xd0, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tscxtnum_el0, x8",
			wantErr: false,
		},
		{
			name: "msr	scxtnum_el1, x7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe7, 0xd0, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tscxtnum_el1, x7",
			wantErr: false,
		},
		{
			name: "msr	scxtnum_el2, x6",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe6, 0xd0, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tscxtnum_el2, x6",
			wantErr: false,
		},
		{
			name: "msr	scxtnum_el3, x5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe5, 0xd0, 0x1e, 0xd5}),
				address:          0,
			},
			want:    "msr\tscxtnum_el3, x5",
			wantErr: false,
		},
		{
			name: "msr	scxtnum_el12, x4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe4, 0xd0, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\tscxtnum_el12, x4",
			wantErr: false,
		},
		// llvm/test/MC/AArch64/armv8.5a-ssbs.s
		{
			name: "mrs	x2, ssbs",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc2, 0x42, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx2, ssbs",
			wantErr: false,
		},
		{
			name: "msr	ssbs, x3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc3, 0x42, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tssbs, x3",
			wantErr: false,
		},
		{
			name: "msr	ssbs, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0x41, 0x03, 0xd5}),
				address:          0,
			},
			want:    "msr\tssbs, #0x1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decompose(tt.args.address, tt.args.instructionValue, &results)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decompose() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got.Disassembly, tt.want) {
				t.Errorf("Disassembly = %v, want %v", got.Disassembly, tt.want)
			}
		})
	}
}

func Test_decompose_MTE(t *testing.T) {
	var results [1024]byte
	type args struct {
		instructionValue uint32
		address          uint64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "irg	x0, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x10, 0xdf, 0x9a}),
				address:          0,
			},
			want:    "irg\tx0, x1",
			wantErr: false,
		},
		{
			name: "irg	sp, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0x10, 0xdf, 0x9a}),
				address:          0,
			},
			want:    "irg\tsp, x1",
			wantErr: false,
		},
		{
			name: "irg	x0, sp",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x13, 0xdf, 0x9a}),
				address:          0,
			},
			want:    "irg\tx0, sp",
			wantErr: false,
		},
		{
			name: "irg	x0, x1, x2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x10, 0xc2, 0x9a}),
				address:          0,
			},
			want:    "irg\tx0, x1, x2",
			wantErr: false,
		},
		{
			name: "irg	sp, x1, x2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0x10, 0xc2, 0x9a}),
				address:          0,
			},
			want:    "irg\tsp, x1, x2",
			wantErr: false,
		},
		{
			name: "addg	x0, x1, #0, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x04, 0x80, 0x91}),
				address:          0,
			},
			want:    "addg\tx0, x1, #0, #0x1",
			wantErr: false,
		},
		{
			name: "addg	sp, x2, #32, #3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x0c, 0x82, 0x91}),
				address:          0,
			},
			want:    "addg\tsp, x2, #0x20, #0x3",
			wantErr: false,
		},
		{
			name: "addg	x0, sp, #64, #5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x17, 0x84, 0x91}),
				address:          0,
			},
			want:    "addg\tx0, sp, #0x40, #0x5",
			wantErr: false,
		},
		{
			name: "addg	x3, x4, #1008, #6",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x83, 0x18, 0xbf, 0x91}),
				address:          0,
			},
			want:    "addg\tx3, x4, #0x3f0, #0x6",
			wantErr: false,
		},
		{
			name: "addg	x5, x6, #112, #15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc5, 0x3c, 0x87, 0x91}),
				address:          0,
			},
			want:    "addg\tx5, x6, #0x70, #0xf",
			wantErr: false,
		},
		{
			name: "subg	x0, x1, #0, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x04, 0x80, 0xd1}),
				address:          0,
			},
			want:    "subg\tx0, x1, #0, #0x1",
			wantErr: false,
		},
		{
			name: "subg	sp, x2, #32, #3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x0c, 0x82, 0xd1}),
				address:          0,
			},
			want:    "subg\tsp, x2, #0x20, #0x3",
			wantErr: false,
		},
		{
			name: "subg	x0, sp, #64, #5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x17, 0x84, 0xd1}),
				address:          0,
			},
			want:    "subg\tx0, sp, #0x40, #0x5",
			wantErr: false,
		},
		{
			name: "subg	x3, x4, #1008, #6",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x83, 0x18, 0xbf, 0xd1}),
				address:          0,
			},
			want:    "subg\tx3, x4, #0x3f0, #0x6",
			wantErr: false,
		},
		{
			name: "subg	x5, x6, #112, #15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc5, 0x3c, 0x87, 0xd1}),
				address:          0,
			},
			want:    "subg\tx5, x6, #0x70, #0xf",
			wantErr: false,
		},
		{
			name: "gmi	x0, x1, x2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x14, 0xc2, 0x9a}),
				address:          0,
			},
			want:    "gmi\tx0, x1, x2",
			wantErr: false,
		},
		{
			name: "gmi	x3, sp, x4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x17, 0xc4, 0x9a}),
				address:          0,
			},
			want:    "gmi\tx3, sp, x4",
			wantErr: false,
		},
		{
			name: "gmi	xzr, x0, x30",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x14, 0xde, 0x9a}),
				address:          0,
			},
			want:    "gmi\txzr, x0, lr",
			wantErr: false,
		},
		{
			name: "gmi	x30, x0, xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1e, 0x14, 0xdf, 0x9a}),
				address:          0,
			},
			want:    "gmi\tlr, x0, xzr",
			wantErr: false,
		},
		{
			name: "stg	x0, [x1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x08, 0x20, 0xd9}),
				address:          0,
			},
			want:    "stg\tx0, [x1]",
			wantErr: false,
		},
		{
			name: "stg	x1, [x1, #-4096]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x21, 0x08, 0x30, 0xd9}),
				address:          0,
			},
			want:    "stg\tx1, [x1, #-0x1000]",
			wantErr: false,
		},
		{
			name: "stg	x2, [x2, #4080]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x42, 0xf8, 0x2f, 0xd9}),
				address:          0,
			},
			want:    "stg\tx2, [x2, #0xff0]",
			wantErr: false,
		},
		{
			name: "stg	x3, [sp, #16]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x1b, 0x20, 0xd9}),
				address:          0,
			},
			want:    "stg\tx3, [sp, #0x10]",
			wantErr: false,
		},
		{
			name: "stg	sp, [sp, #16]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x1b, 0x20, 0xd9}),
				address:          0,
			},
			want:    "stg\tsp, [sp, #0x10]",
			wantErr: false,
		},
		{
			name: "stzg	x0, [x1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x08, 0x60, 0xd9}),
				address:          0,
			},
			want:    "stzg\tx0, [x1]",
			wantErr: false,
		},
		{
			name: "stzg	x1, [x1, #-4096]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x21, 0x08, 0x70, 0xd9}),
				address:          0,
			},
			want:    "stzg\tx1, [x1, #-0x1000]",
			wantErr: false,
		},
		{
			name: "stzg	x2, [x2, #4080]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x42, 0xf8, 0x6f, 0xd9}),
				address:          0,
			},
			want:    "stzg\tx2, [x2, #0xff0]",
			wantErr: false,
		},
		{
			name: "stzg	x3, [sp, #16]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x1b, 0x60, 0xd9}),
				address:          0,
			},
			want:    "stzg\tx3, [sp, #0x10]",
			wantErr: false,
		},
		{
			name: "stzg	sp, [sp, #16]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x1b, 0x60, 0xd9}),
				address:          0,
			},
			want:    "stzg\tsp, [sp, #0x10]",
			wantErr: false,
		},
		{
			name: "stg	x0, [x1, #-4096]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x0c, 0x30, 0xd9}),
				address:          0,
			},
			want:    "stg\tx0, [x1, #-0x1000]!",
			wantErr: false,
		},
		{
			name: "stg	x1, [x2, #4080]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0xfc, 0x2f, 0xd9}),
				address:          0,
			},
			want:    "stg\tx1, [x2, #0xff0]!",
			wantErr: false,
		},
		{
			name: "stg	x2, [sp, #16]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe2, 0x1f, 0x20, 0xd9}),
				address:          0,
			},
			want:    "stg\tx2, [sp, #0x10]!",
			wantErr: false,
		},
		{
			name: "stg	sp, [sp, #16]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x1f, 0x20, 0xd9}),
				address:          0,
			},
			want:    "stg\tsp, [sp, #0x10]!",
			wantErr: false,
		},
		{
			name: "stzg	x0, [x1, #-4096]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x0c, 0x70, 0xd9}),
				address:          0,
			},
			want:    "stzg\tx0, [x1, #-0x1000]!",
			wantErr: false,
		},
		{
			name: "stzg	x1, [x2, #4080]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0xfc, 0x6f, 0xd9}),
				address:          0,
			},
			want:    "stzg\tx1, [x2, #0xff0]!",
			wantErr: false,
		},
		{
			name: "stzg	x2, [sp, #16]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe2, 0x1f, 0x60, 0xd9}),
				address:          0,
			},
			want:    "stzg\tx2, [sp, #0x10]!",
			wantErr: false,
		},
		{
			name: "stzg	sp, [sp, #16]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x1f, 0x60, 0xd9}),
				address:          0,
			},
			want:    "stzg\tsp, [sp, #0x10]!",
			wantErr: false,
		},
		{
			name: "stg	x0, [x1], #-4096",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x04, 0x30, 0xd9}),
				address:          0,
			},
			want:    "stg\tx0, [x1], #-0x1000",
			wantErr: false,
		},
		{
			name: "stg	x1, [x2], #4080",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0xf4, 0x2f, 0xd9}),
				address:          0,
			},
			want:    "stg\tx1, [x2], #0xff0",
			wantErr: false,
		},
		{
			name: "stg	x2, [sp], #16",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe2, 0x17, 0x20, 0xd9}),
				address:          0,
			},
			want:    "stg\tx2, [sp], #0x10",
			wantErr: false,
		},
		{
			name: "stg	sp, [sp], #16",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x17, 0x20, 0xd9}),
				address:          0,
			},
			want:    "stg\tsp, [sp], #0x10",
			wantErr: false,
		},
		{
			name: "stzg	x0, [x1], #-4096",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x04, 0x70, 0xd9}),
				address:          0,
			},
			want:    "stzg\tx0, [x1], #-0x1000",
			wantErr: false,
		},
		{
			name: "stzg	x1, [x2], #4080",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0xf4, 0x6f, 0xd9}),
				address:          0,
			},
			want:    "stzg\tx1, [x2], #0xff0",
			wantErr: false,
		},
		{
			name: "stzg	x2, [sp], #16",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe2, 0x17, 0x60, 0xd9}),
				address:          0,
			},
			want:    "stzg\tx2, [sp], #0x10",
			wantErr: false,
		},
		{
			name: "stzg	sp, [sp], #16",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x17, 0x60, 0xd9}),
				address:          0,
			},
			want:    "stzg\tsp, [sp], #0x10",
			wantErr: false,
		},
		{
			name: "st2g	x0, [x1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x08, 0xa0, 0xd9}),
				address:          0,
			},
			want:    "st2g\tx0, [x1]",
			wantErr: false,
		},
		{
			name: "st2g	x1, [x1, #-4096]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x21, 0x08, 0xb0, 0xd9}),
				address:          0,
			},
			want:    "st2g\tx1, [x1, #-0x1000]",
			wantErr: false,
		},
		{
			name: "st2g	x2, [x2, #4080]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x42, 0xf8, 0xaf, 0xd9}),
				address:          0,
			},
			want:    "st2g\tx2, [x2, #0xff0]",
			wantErr: false,
		},
		{
			name: "st2g	x3, [sp, #16]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x1b, 0xa0, 0xd9}),
				address:          0,
			},
			want:    "st2g\tx3, [sp, #0x10]",
			wantErr: false,
		},
		{
			name: "st2g	sp, [sp, #16]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x1b, 0xa0, 0xd9}),
				address:          0,
			},
			want:    "st2g\tsp, [sp, #0x10]",
			wantErr: false,
		},
		{
			name: "stz2g	x0, [x1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x08, 0xe0, 0xd9}),
				address:          0,
			},
			want:    "stz2g\tx0, [x1]",
			wantErr: false,
		},
		{
			name: "stz2g	x1, [x1, #-4096]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x21, 0x08, 0xf0, 0xd9}),
				address:          0,
			},
			want:    "stz2g\tx1, [x1, #-0x1000]",
			wantErr: false,
		},
		{
			name: "stz2g	x2, [x2, #4080]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x42, 0xf8, 0xef, 0xd9}),
				address:          0,
			},
			want:    "stz2g\tx2, [x2, #0xff0]",
			wantErr: false,
		},
		{
			name: "stz2g	x3, [sp, #16]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x1b, 0xe0, 0xd9}),
				address:          0,
			},
			want:    "stz2g\tx3, [sp, #0x10]",
			wantErr: false,
		},
		{
			name: "stz2g	sp, [sp, #16]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x1b, 0xe0, 0xd9}),
				address:          0,
			},
			want:    "stz2g\tsp, [sp, #0x10]",
			wantErr: false,
		},
		{
			name: "st2g	x0, [x1, #-4096]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x0c, 0xb0, 0xd9}),
				address:          0,
			},
			want:    "st2g\tx0, [x1, #-0x1000]!",
			wantErr: false,
		},
		{
			name: "st2g	x1, [x2, #4080]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0xfc, 0xaf, 0xd9}),
				address:          0,
			},
			want:    "st2g\tx1, [x2, #0xff0]!",
			wantErr: false,
		},
		{
			name: "st2g	x2, [sp, #16]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe2, 0x1f, 0xa0, 0xd9}),
				address:          0,
			},
			want:    "st2g\tx2, [sp, #0x10]!",
			wantErr: false,
		},
		{
			name: "st2g	sp, [sp, #16]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x1f, 0xa0, 0xd9}),
				address:          0,
			},
			want:    "st2g\tsp, [sp, #0x10]!",
			wantErr: false,
		},
		{
			name: "stz2g	x0, [x1, #-4096]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x0c, 0xf0, 0xd9}),
				address:          0,
			},
			want:    "stz2g\tx0, [x1, #-0x1000]!",
			wantErr: false,
		},
		{
			name: "stz2g	x1, [x2, #4080]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0xfc, 0xef, 0xd9}),
				address:          0,
			},
			want:    "stz2g\tx1, [x2, #0xff0]!",
			wantErr: false,
		},
		{
			name: "stz2g	x2, [sp, #16]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe2, 0x1f, 0xe0, 0xd9}),
				address:          0,
			},
			want:    "stz2g\tx2, [sp, #0x10]!",
			wantErr: false,
		},
		{
			name: "stz2g	sp, [sp, #16]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x1f, 0xe0, 0xd9}),
				address:          0,
			},
			want:    "stz2g\tsp, [sp, #0x10]!",
			wantErr: false,
		},
		{
			name: "st2g	x0, [x1], #-4096",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x04, 0xb0, 0xd9}),
				address:          0,
			},
			want:    "st2g\tx0, [x1], #-0x1000",
			wantErr: false,
		},
		{
			name: "st2g	x1, [x2], #4080",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0xf4, 0xaf, 0xd9}),
				address:          0,
			},
			want:    "st2g\tx1, [x2], #0xff0",
			wantErr: false,
		},
		{
			name: "st2g	x2, [sp], #16",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe2, 0x17, 0xa0, 0xd9}),
				address:          0,
			},
			want:    "st2g\tx2, [sp], #0x10",
			wantErr: false,
		},
		{
			name: "st2g	sp, [sp], #16",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x17, 0xa0, 0xd9}),
				address:          0,
			},
			want:    "st2g\tsp, [sp], #0x10",
			wantErr: false,
		},
		{
			name: "stz2g	x0, [x1], #-4096",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x04, 0xf0, 0xd9}),
				address:          0,
			},
			want:    "stz2g\tx0, [x1], #-0x1000",
			wantErr: false,
		},
		{
			name: "stz2g	x1, [x2], #4080",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0xf4, 0xef, 0xd9}),
				address:          0,
			},
			want:    "stz2g\tx1, [x2], #0xff0",
			wantErr: false,
		},
		{
			name: "stz2g	x2, [sp], #16",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe2, 0x17, 0xe0, 0xd9}),
				address:          0,
			},
			want:    "stz2g\tx2, [sp], #0x10",
			wantErr: false,
		},
		{
			name: "stz2g	sp, [sp], #16",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x17, 0xe0, 0xd9}),
				address:          0,
			},
			want:    "stz2g\tsp, [sp], #0x10",
			wantErr: false,
		},
		{
			name: "stgp	x0, x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0x04, 0x00, 0x69}),
				address:          0,
			},
			want:    "stgp\tx0, x1, [x2]",
			wantErr: false,
		},
		{
			name: "stgp	x0, x1, [x2, #-1024]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0x04, 0x20, 0x69}),
				address:          0,
			},
			want:    "stgp\tx0, x1, [x2, #-0x400]",
			wantErr: false,
		},
		{
			name: "stgp	x0, x1, [x2, #1008]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0x84, 0x1f, 0x69}),
				address:          0,
			},
			want:    "stgp\tx0, x1, [x2, #0x3f0]",
			wantErr: false,
		},
		{
			name: "stgp	x0, x1, [sp, #16]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x87, 0x00, 0x69}),
				address:          0,
			},
			want:    "stgp\tx0, x1, [sp, #0x10]",
			wantErr: false,
		},
		{
			name: "stgp	xzr, x1, [x2, #16]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x84, 0x00, 0x69}),
				address:          0,
			},
			want:    "stgp\txzr, x1, [x2, #0x10]",
			wantErr: false,
		},
		{
			name: "stgp	x0, xzr, [x2, #16]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0xfc, 0x00, 0x69}),
				address:          0,
			},
			want:    "stgp\tx0, xzr, [x2, #0x10]",
			wantErr: false,
		},
		{
			name: "stgp	x0, x1, [x2, #-1024]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0x04, 0xa0, 0x69}),
				address:          0,
			},
			want:    "stgp\tx0, x1, [x2, #-0x400]!",
			wantErr: false,
		},
		{
			name: "stgp	x0, x1, [x2, #1008]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0x84, 0x9f, 0x69}),
				address:          0,
			},
			want:    "stgp\tx0, x1, [x2, #0x3f0]!",
			wantErr: false,
		},
		{
			name: "stgp	x0, x1, [sp, #16]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x87, 0x80, 0x69}),
				address:          0,
			},
			want:    "stgp\tx0, x1, [sp, #0x10]!",
			wantErr: false,
		},
		{
			name: "stgp	xzr, x1, [x2, #16]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x84, 0x80, 0x69}),
				address:          0,
			},
			want:    "stgp\txzr, x1, [x2, #0x10]!",
			wantErr: false,
		},
		{
			name: "stgp	x0, xzr, [x2, #16]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0xfc, 0x80, 0x69}),
				address:          0,
			},
			want:    "stgp\tx0, xzr, [x2, #0x10]!",
			wantErr: false,
		},
		{
			name: "stgp	x0, x1, [x2], #-1024",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0x04, 0xa0, 0x68}),
				address:          0,
			},
			want:    "stgp\tx0, x1, [x2], #-0x400",
			wantErr: false,
		},
		{
			name: "stgp	x0, x1, [x2], #1008",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0x84, 0x9f, 0x68}),
				address:          0,
			},
			want:    "stgp\tx0, x1, [x2], #0x3f0",
			wantErr: false,
		},
		{
			name: "stgp	x0, x1, [sp], #16",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x87, 0x80, 0x68}),
				address:          0,
			},
			want:    "stgp\tx0, x1, [sp], #0x10",
			wantErr: false,
		},
		{
			name: "stgp	xzr, x1, [x2], #16",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x84, 0x80, 0x68}),
				address:          0,
			},
			want:    "stgp\txzr, x1, [x2], #0x10",
			wantErr: false,
		},
		{
			name: "stgp	x0, xzr, [x2], #16",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0xfc, 0x80, 0x68}),
				address:          0,
			},
			want:    "stgp\tx0, xzr, [x2], #0x10",
			wantErr: false,
		},
		{
			name: "dc	igvac, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x60, 0x76, 0x08, 0xd5}),
				address:          0,
			},
			want:    "dc\tIGVAC, x0",
			wantErr: false,
		},
		{
			name: "dc	igsw, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x81, 0x76, 0x08, 0xd5}),
				address:          0,
			},
			want:    "dc\tIGSW, x1",
			wantErr: false,
		},
		{
			name: "dc	cgsw, x2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x82, 0x7a, 0x08, 0xd5}),
				address:          0,
			},
			want:    "dc\tCGSW, x2",
			wantErr: false,
		},
		{
			name: "dc	cigsw, x3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x83, 0x7e, 0x08, 0xd5}),
				address:          0,
			},
			want:    "dc\tCIGSW, x3",
			wantErr: false,
		},
		{
			name: "dc	cgvac, x4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x64, 0x7a, 0x0b, 0xd5}),
				address:          0,
			},
			want:    "dc\tCGVAC, x4",
			wantErr: false,
		},
		{
			name: "dc	cgvap, x5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x65, 0x7c, 0x0b, 0xd5}),
				address:          0,
			},
			want:    "dc\tCGVAP, x5",
			wantErr: false,
		},
		{
			name: "dc	cgvadp, x6",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x66, 0x7d, 0x0b, 0xd5}),
				address:          0,
			},
			want:    "dc\tCGVADP, x6",
			wantErr: false,
		},
		{
			name: "dc	cigvac, x7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x67, 0x7e, 0x0b, 0xd5}),
				address:          0,
			},
			want:    "dc\tCIGVAC, x7",
			wantErr: false,
		},
		{
			name: "dc	gva, x8",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x68, 0x74, 0x0b, 0xd5}),
				address:          0,
			},
			want:    "dc\tGVA, x8",
			wantErr: false,
		},
		{
			name: "dc	igdvac, x9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x76, 0x08, 0xd5}),
				address:          0,
			},
			want:    "dc\tIGDVAC, x9",
			wantErr: false,
		},
		{
			name: "dc	igdsw, x10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xca, 0x76, 0x08, 0xd5}),
				address:          0,
			},
			want:    "dc\tIGDSW, x10",
			wantErr: false,
		},
		{
			name: "dc	cgdsw, x11",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcb, 0x7a, 0x08, 0xd5}),
				address:          0,
			},
			want:    "dc\tCGDSW, x11",
			wantErr: false,
		},
		{
			name: "dc	cigdsw, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x7e, 0x08, 0xd5}),
				address:          0,
			},
			want:    "dc\tCIGDSW, x12",
			wantErr: false,
		},
		{
			name: "dc	cgdvac, x13",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xad, 0x7a, 0x0b, 0xd5}),
				address:          0,
			},
			want:    "dc\tCGDVAC, x13",
			wantErr: false,
		},
		{
			name: "dc	cgdvap, x14",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xae, 0x7c, 0x0b, 0xd5}),
				address:          0,
			},
			want:    "dc\tCGDVAP, x14",
			wantErr: false,
		},
		{
			name: "dc	cgdvadp, x15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xaf, 0x7d, 0x0b, 0xd5}),
				address:          0,
			},
			want:    "dc\tCGDVADP, x15",
			wantErr: false,
		},
		{
			name: "dc	cigdvac, x16",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb0, 0x7e, 0x0b, 0xd5}),
				address:          0,
			},
			want:    "dc\tCIGDVAC, x16",
			wantErr: false,
		},
		{
			name: "dc	gzva, x17",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x91, 0x74, 0x0b, 0xd5}),
				address:          0,
			},
			want:    "dc\tGZVA, x17",
			wantErr: false,
		},
		{
			name: "mrs	x0, TCO",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x42, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, tco",
			wantErr: false,
		},
		{
			name: "mrs	x1, GCR_EL1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc1, 0x10, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx1, gcr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x2, RGSR_EL1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa2, 0x10, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx2, rgsr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x3, TFSR_EL1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x03, 0x56, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx3, tfsr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x4, TFSR_EL2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x04, 0x56, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx4, tfsr_el2",
			wantErr: false,
		},
		{
			name: "mrs	x5, TFSR_EL3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x05, 0x56, 0x3e, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx5, tfsr_el3",
			wantErr: false,
		},
		{
			name: "mrs	x6, TFSR_EL12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x06, 0x56, 0x3d, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx6, tfsr_el12",
			wantErr: false,
		},
		{
			name: "mrs	x7, TFSRE0_EL1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x27, 0x56, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx7, tfsre0_el1",
			wantErr: false,
		},
		{
			name: "mrs	x7, GMID_EL1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x87, 0x00, 0x39, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx7, s3_1_c0_c0_4",
			wantErr: false,
		},
		{
			name: "msr	TCO, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x40, 0x03, 0xd5}),
				address:          0,
			},
			want:    "msr\ttco, #0",
			wantErr: false,
		},
		{
			name: "msr	TCO, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x42, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\ttco, x0",
			wantErr: false,
		},
		{
			name: "msr	GCR_EL1, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc1, 0x10, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tgcr_el1, x1",
			wantErr: false,
		},
		{
			name: "msr	RGSR_EL1, x2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa2, 0x10, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\trgsr_el1, x2",
			wantErr: false,
		},
		{
			name: "msr	TFSR_EL1, x3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x03, 0x56, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\ttfsr_el1, x3",
			wantErr: false,
		},
		{
			name: "msr	TFSR_EL2, x4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x04, 0x56, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ttfsr_el2, x4",
			wantErr: false,
		},
		{
			name: "msr	TFSR_EL3, x5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x05, 0x56, 0x1e, 0xd5}),
				address:          0,
			},
			want:    "msr\ttfsr_el3, x5",
			wantErr: false,
		},
		{
			name: "msr	TFSR_EL12, x6",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x06, 0x56, 0x1d, 0xd5}),
				address:          0,
			},
			want:    "msr\ttfsr_el12, x6",
			wantErr: false,
		},
		{
			name: "msr	TFSRE0_EL1, x7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x27, 0x56, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\ttfsre0_el1, x7",
			wantErr: false,
		},
		{
			name: "subp	x0, x1, x2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x00, 0xc2, 0x9a}),
				address:          0,
			},
			want:    "subp\tx0, x1, x2",
			wantErr: false,
		},
		{
			name: "subp	x0, sp, sp",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x03, 0xdf, 0x9a}),
				address:          0,
			},
			want:    "subp\tx0, sp, sp",
			wantErr: false,
		},
		{
			name: "subps	x0, x1, x2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x00, 0xc2, 0xba}),
				address:          0,
			},
			want:    "subps\tx0, x1, x2",
			wantErr: false,
		},
		{
			name: "subps	x0, sp, sp",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x03, 0xdf, 0xba}),
				address:          0,
			},
			want:    "subps\tx0, sp, sp",
			wantErr: false,
		},
		{
			name: "subps	xzr, x0, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x00, 0xc1, 0xba}),
				address:          0,
			},
			want:    "cmpp\tx0, x1",
			wantErr: false,
		},
		{
			name: "subps	xzr, sp, sp",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0xdf, 0xba}),
				address:          0,
			},
			want:    "cmpp\tsp, sp",
			wantErr: false,
		},
		{
			name: "ldg	x0, [x1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x00, 0x60, 0xd9}),
				address:          0,
			},
			want:    "ldg\tx0, [x1]",
			wantErr: false,
		},
		{
			name: "ldg	x2, [sp, #-4096]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe2, 0x03, 0x70, 0xd9}),
				address:          0,
			},
			want:    "ldg\tx2, [sp, #-0x1000]",
			wantErr: false,
		},
		{
			name: "ldg	x3, [x4, #4080]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x83, 0xf0, 0x6f, 0xd9}),
				address:          0,
			},
			want:    "ldg\tx3, [x4, #0xff0]",
			wantErr: false,
		},
		{
			name: "ldgm	x0, [x1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x00, 0xe0, 0xd9}),
				address:          0,
			},
			want:    "ldgm\tx0, [x1]",
			wantErr: false,
		},
		{
			name: "ldgm	x1, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe1, 0x03, 0xe0, 0xd9}),
				address:          0,
			},
			want:    "ldgm\tx1, [sp]",
			wantErr: false,
		},
		{
			name: "ldgm	xzr, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x00, 0xe0, 0xd9}),
				address:          0,
			},
			want:    "ldgm\txzr, [x2]",
			wantErr: false,
		},
		{
			name: "stgm	x0, [x1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x00, 0xa0, 0xd9}),
				address:          0,
			},
			want:    "stgm\tx0, [x1]",
			wantErr: false,
		},
		{
			name: "stgm	x1, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe1, 0x03, 0xa0, 0xd9}),
				address:          0,
			},
			want:    "stgm\tx1, [sp]",
			wantErr: false,
		},
		{
			name: "stgm	xzr, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x00, 0xa0, 0xd9}),
				address:          0,
			},
			want:    "stgm\txzr, [x2]",
			wantErr: false,
		},
		{
			name: "stzgm	x0, [x1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x00, 0x20, 0xd9}),
				address:          0,
			},
			want:    "stzgm\tx0, [x1]",
			wantErr: false,
		},
		{
			name: "stzgm	x1, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe1, 0x03, 0x20, 0xd9}),
				address:          0,
			},
			want:    "stzgm\tx1, [sp]",
			wantErr: false,
		},
		{
			name: "stzgm	xzr, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x00, 0x20, 0xd9}),
				address:          0,
			},
			want:    "stzgm\txzr, [x2]",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decompose(tt.args.address, tt.args.instructionValue, &results)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decompose() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got.Disassembly, tt.want) {
				t.Errorf("Disassembly = %v, want %v", got.Disassembly, tt.want)
			}
		})
	}
}

func Test_decompose_v8_6a(t *testing.T) {
	var results [1024]byte
	type args struct {
		instructionValue uint32
		address          uint64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// llvm/test/MC/AArch64/armv8.6a-amvs.s
		{
			name: "msr	amevcntvoff00_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0xd8, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c8_0, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff01_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xd8, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c8_1, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff02_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0xd8, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c8_2, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff03_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x60, 0xd8, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c8_3, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff04_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x80, 0xd8, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c8_4, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff05_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa0, 0xd8, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c8_5, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff06_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc0, 0xd8, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c8_6, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff07_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0xd8, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c8_7, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff08_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0xd9, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c9_0, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff09_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xd9, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c9_1, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff010_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0xd9, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c9_2, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff011_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x60, 0xd9, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c9_3, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff012_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x80, 0xd9, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c9_4, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff013_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa0, 0xd9, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c9_5, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff014_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc0, 0xd9, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c9_6, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff015_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0xd9, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c9_7, x0",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff00_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0xd8, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c8_0",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff01_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xd8, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c8_1",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff02_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0xd8, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c8_2",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff03_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x60, 0xd8, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c8_3",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff04_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x80, 0xd8, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c8_4",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff05_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa0, 0xd8, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c8_5",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff06_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc0, 0xd8, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c8_6",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff07_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0xd8, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c8_7",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff08_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0xd9, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c9_0",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff09_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xd9, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c9_1",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff010_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0xd9, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c9_2",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff011_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x60, 0xd9, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c9_3",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff012_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x80, 0xd9, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c9_4",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff013_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa0, 0xd9, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c9_5",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff014_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc0, 0xd9, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c9_6",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff015_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0xd9, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c9_7",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff10_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0xda, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c10_0, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff11_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xda, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c10_1, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff12_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0xda, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c10_2, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff13_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x60, 0xda, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c10_3, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff14_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x80, 0xda, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c10_4, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff15_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa0, 0xda, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c10_5, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff16_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc0, 0xda, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c10_6, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff17_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0xda, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c10_7, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff18_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0xdb, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c11_0, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff19_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xdb, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c11_1, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff110_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0xdb, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c11_2, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff111_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x60, 0xdb, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c11_3, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff112_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x80, 0xdb, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c11_4, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff113_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa0, 0xdb, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c11_5, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff114_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc0, 0xdb, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c11_6, x0",
			wantErr: false,
		},
		{
			name: "msr	amevcntvoff115_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0xdb, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c13_c11_7, x0",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff10_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0xda, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c10_0",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff11_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xda, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c10_1",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff12_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0xda, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c10_2",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff13_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x60, 0xda, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c10_3",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff14_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x80, 0xda, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c10_4",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff15_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa0, 0xda, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c10_5",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff16_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc0, 0xda, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c10_6",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff17_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0xda, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c10_7",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff18_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0xdb, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c11_0",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff19_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xdb, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c11_1",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff110_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0xdb, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c11_2",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff111_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x60, 0xdb, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c11_3",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff112_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x80, 0xdb, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c11_4",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff113_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa0, 0xdb, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c11_5",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff114_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc0, 0xdb, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c11_6",
			wantErr: false,
		},
		{
			name: "mrs	x0, amevcntvoff115_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0xdb, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c13_c11_7",
			wantErr: false,
		},

		// llvm/test/MC/AArch64/armv8.6a-bf16.s
		{
			name: "bfdot	v2.2s, v3.4h, v4.4h",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xfc, 0x44, 0x2e}),
				address:          0,
			},
			want:    "bfdot\tv2.2s, v3.2h, v4.2h",
			wantErr: false,
		},
		{
			name: "bfdot	v2.4s, v3.8h, v4.8h",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xfc, 0x44, 0x6e}),
				address:          0,
			},
			want:    "bfdot\tv2.4s, v3.4h, v4.4h",
			wantErr: false,
		},
		{
			name: "bfdot	v2.2s, v3.4h, v4.2h[0]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xf0, 0x44, 0x0f}),
				address:          0,
			},
			want:    "bfdot	v2.2s, v3.4h, v4.2h[0]",
			wantErr: true,
		},
		{
			name: "bfdot	v2.2s, v3.4h, v4.2h[1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xf0, 0x64, 0x0f}),
				address:          0,
			},
			want:    "bfdot	v2.2s, v3.4h, v4.2h[1]",
			wantErr: true,
		},
		{
			name: "bfdot	v2.2s, v3.4h, v4.2h[2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xf8, 0x44, 0x0f}),
				address:          0,
			},
			want:    "bfdot	v2.2s, v3.4h, v4.2h[2]",
			wantErr: true,
		},
		{
			name: "bfdot	v2.2s, v3.4h, v4.2h[3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xf8, 0x64, 0x0f}),
				address:          0,
			},
			want:    "bfdot	v2.2s, v3.4h, v4.2h[3]",
			wantErr: true,
		},
		{
			name: "bfdot	v2.4s, v3.8h, v4.2h[0]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xf0, 0x44, 0x4f}),
				address:          0,
			},
			want:    "bfdot	v2.4s, v3.8h, v4.2h[0]",
			wantErr: true,
		},
		{
			name: "bfdot	v2.4s, v3.8h, v4.2h[1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xf0, 0x64, 0x4f}),
				address:          0,
			},
			want:    "bfdot	v2.4s, v3.8h, v4.2h[1]",
			wantErr: true,
		},
		{
			name: "bfdot	v2.4s, v3.8h, v4.2h[2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xf8, 0x44, 0x4f}),
				address:          0,
			},
			want:    "bfdot	v2.4s, v3.8h, v4.2h[2]",
			wantErr: true,
		},
		{
			name: "bfdot	v2.4s, v3.8h, v4.2h[3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xf8, 0x64, 0x4f}),
				address:          0,
			},
			want:    "bfdot	v2.4s, v3.8h, v4.2h[3]",
			wantErr: true,
		},
		{
			name: "bfmmla	v2.4s, v3.8h, v4.8h",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xec, 0x44, 0x6e}),
				address:          0,
			},
			want:    "bfmmla	v2.4s, v3.8h, v4.8h",
			wantErr: true,
		},
		{
			name: "bfmmla	v3.4s, v4.8h, v5.8h",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x83, 0xec, 0x45, 0x6e}),
				address:          0,
			},
			want:    "bfmmla	v3.4s, v4.8h, v5.8h",
			wantErr: true,
		},
		{
			name: "bfcvtn	v5.4h, v5.4s",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa5, 0x68, 0xa1, 0x0e}),
				address:          0,
			},
			want:    "bfcvtn\tv5.2d, v5.4s",
			wantErr: false,
		},
		{
			name: "bfcvtn2	v5.8h, v5.4s",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa5, 0x68, 0xa1, 0x4e}),
				address:          0,
			},
			want:    "bfcvtn2\tv5.2d, v5.4s",
			wantErr: false,
		},
		{
			name: "bfcvt	h5, s3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x65, 0x40, 0x63, 0x1e}),
				address:          0,
			},
			want:    "bfcvt\th5, s3",
			wantErr: false,
		},
		{
			name: "bfmlalb	v10.4s, v21.8h, v14.8h",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xaa, 0xfe, 0xce, 0x2e}),
				address:          0,
			},
			want:    "bfmlalb	v10.4s, v21.8h, v14.8h",
			wantErr: true,
		},
		{
			name: "bfmlalt	v21.4s, v14.8h, v10.8h",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd5, 0xfd, 0xca, 0x6e}),
				address:          0,
			},
			want:    "bfmlalt	v21.4s, v14.8h, v10.8h",
			wantErr: true,
		},
		{
			name: "bfmlalb	v14.4s, v21.8h, v10.h[1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xae, 0xf2, 0xda, 0x0f}),
				address:          0,
			},
			want:    "bfmlalb	v14.4s, v21.8h, v10.h[1]",
			wantErr: true,
		},
		{
			name: "bfmlalb	v14.4s, v21.8h, v10.h[2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xae, 0xf2, 0xea, 0x0f}),
				address:          0,
			},
			want:    "bfmlalb	v14.4s, v21.8h, v10.h[2]",
			wantErr: true,
		},
		{
			name: "bfmlalb	v14.4s, v21.8h, v10.h[7]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xae, 0xfa, 0xfa, 0x0f}),
				address:          0,
			},
			want:    "bfmlalb	v14.4s, v21.8h, v10.h[7]",
			wantErr: true,
		},
		{
			name: "bfmlalt	v21.4s, v10.8h, v14.h[1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x55, 0xf1, 0xde, 0x4f}),
				address:          0,
			},
			want:    "bfmlalt	v21.4s, v10.8h, v14.h[1]",
			wantErr: true,
		},
		{
			name: "bfmlalt	v21.4s, v10.8h, v14.h[2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x55, 0xf1, 0xee, 0x4f}),
				address:          0,
			},
			want:    "bfmlalt	v21.4s, v10.8h, v14.h[2]",
			wantErr: true,
		},
		{
			name: "bfmlalt	v21.4s, v10.8h, v14.h[7]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x55, 0xf9, 0xfe, 0x4f}),
				address:          0,
			},
			want:    "bfmlalt	v21.4s, v10.8h, v14.h[7]",
			wantErr: true,
		},
		// llvm/test/MC/AArch64/armv8.6a-ecv.s
		{
			name: "msr	cntscale_el2, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x81, 0xe0, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c14_c0_4, x1",
			wantErr: false,
		},
		{
			name: "msr	cntiscale_el2, x11",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xab, 0xe0, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c14_c0_5, x11",
			wantErr: false,
		},
		{
			name: "msr	cntpoff_el2, x22",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd6, 0xe0, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c14_c0_6, x22",
			wantErr: false,
		},
		{
			name: "msr	cntvfrq_el2, x3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0xe0, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_4_c14_c0_7, x3",
			wantErr: false,
		},
		{
			name: "msr	cntpctss_el0, x13",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xad, 0xe0, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tcntpctss_el0, x13",
			wantErr: false,
		},
		{
			name: "msr	cntvctss_el0, x23",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd7, 0xe0, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tcntvctss_el0, x23",
			wantErr: false,
		},
		{
			name: "mrs	x0, cntscale_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x80, 0xe0, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx0, s3_4_c14_c0_4",
			wantErr: false,
		},
		{
			name: "mrs	x5, cntiscale_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa5, 0xe0, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx5, s3_4_c14_c0_5",
			wantErr: false,
		},
		{
			name: "mrs	x10, cntpoff_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xca, 0xe0, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx10, s3_4_c14_c0_6",
			wantErr: false,
		},
		{
			name: "mrs	x15, cntvfrq_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xef, 0xe0, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx15, s3_4_c14_c0_7",
			wantErr: false,
		},
		{
			name: "mrs	x20, cntpctss_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb4, 0xe0, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx20, cntpctss_el0",
			wantErr: false,
		},
		{
			name: "mrs	x30, cntvctss_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xde, 0xe0, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tlr, cntvctss_el0",
			wantErr: false,
		},

		// llvm/test/MC/AArch64/armv8.6a-fgt.s
		{
			name: "msr	hfgrtr_el2, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x80, 0x11, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\thfgrtr_el2, x0",
			wantErr: false,
		},
		{
			name: "msr	hfgwtr_el2, x5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa5, 0x11, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\thfgwtr_el2, x5",
			wantErr: false,
		},
		{
			name: "msr	hfgitr_el2, x10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xca, 0x11, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\thfgitr_el2, x10",
			wantErr: false,
		},
		{
			name: "msr	hdfgrtr_el2, x15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8f, 0x31, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\thdfgrtr_el2, x15",
			wantErr: false,
		},
		{
			name: "msr	hdfgwtr_el2, x20",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb4, 0x31, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\thdfgwtr_el2, x20",
			wantErr: false,
		},
		{
			name: "mrs	x30, hfgrtr_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9e, 0x11, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tlr, hfgrtr_el2",
			wantErr: false,
		},
		{
			name: "mrs	x25, hfgwtr_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb9, 0x11, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx25, hfgwtr_el2",
			wantErr: false,
		},
		{
			name: "mrs	x20, hfgitr_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd4, 0x11, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx20, hfgitr_el2",
			wantErr: false,
		},
		{
			name: "mrs	x15, hdfgrtr_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8f, 0x31, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx15, hdfgrtr_el2",
			wantErr: false,
		},
		{
			name: "mrs	x10, hdfgwtr_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xaa, 0x31, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx10, hdfgwtr_el2",
			wantErr: false,
		},

		// llvm/test/MC/AArch64/armv8.6a-simd-matmul.s
		//
		// The BF16 element/matmul forms below currently fail decode with
		// DECODE_STATUS_ERROR_OPERANDS. Keep those expectations explicit until
		// decoder support lands rather than silently treating them as covered.
		{
			name: "smmla	v1.4s, v16.16b, v31.16b",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x01, 0xa6, 0x9f, 0x4e}),
				address:          0,
			},
			want:    "smmla\tv1.4s, v16.16b, v31.16b",
			wantErr: false,
		},
		{
			name: "ummla	v1.4s, v16.16b, v31.16b",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x01, 0xa6, 0x9f, 0x6e}),
				address:          0,
			},
			want:    "ummla\tv1.4s, v16.16b, v31.16b",
			wantErr: false,
		},
		{
			name: "usmmla	v1.4s, v16.16b, v31.16b",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x01, 0xae, 0x9f, 0x4e}),
				address:          0,
			},
			want:    "usmmla\tv1.4s, v16.16b, v31.16b",
			wantErr: false,
		},
		{
			name: "usdot	v3.2s, v15.8b, v30.8b",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x9d, 0x9e, 0x0e}),
				address:          0,
			},
			want:    "usdot\tv3.2s, v15.8b, v30.8b",
			wantErr: false,
		},
		{
			name: "usdot	v3.4s, v15.16b, v30.16b",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x9d, 0x9e, 0x4e}),
				address:          0,
			},
			want:    "usdot\tv3.4s, v15.16b, v30.16b",
			wantErr: false,
		},
		{
			name: "usdot	v31.2s, v1.8b, v2.4b[0]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0xf8, 0xa2, 0x0f}),
				address:          0,
			},
			want:    "usdot\tv31.2s, v1.8b, v2.4b[0]",
			wantErr: false,
		},
		{
			name: "usdot	v31.4s, v1.16b, v2.4b[0]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0xf8, 0xa2, 0x4f}),
				address:          0,
			},
			want:    "usdot\tv31.4s, v1.16b, v2.4b[0]",
			wantErr: false,
		},
		{
			name: "sudot	v31.2s, v1.8b, v2.4b[0]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0xf8, 0x22, 0x0f}),
				address:          0,
			},
			want:    "sudot\tv31.2s, v1.8b, v2.4b[0]",
			wantErr: false,
		},
		{
			name: "sudot	v31.4s, v1.16b, v2.4b[0]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0xf8, 0x22, 0x4f}),
				address:          0,
			},
			want:    "sudot\tv31.4s, v1.16b, v2.4b[0]",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decompose(tt.args.address, tt.args.instructionValue, &results)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decompose() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got.Disassembly, tt.want) {
				t.Errorf("Disassembly = %v, want %v", got.Disassembly, tt.want)
			}
		})
	}
}

func Test_decompose_basic(t *testing.T) {
	var results [1024]byte
	type args struct {
		instructionValue uint32
		address          uint64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "add	x2, x4, w5, uxtb",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x82, 0x00, 0x25, 0x8b}),
				address:          0,
			},
			want:    "add\tx2, x4, w5, uxtb",
			wantErr: false,
		},
		{
			name: "add	x20, sp, w19, uxth",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x23, 0x33, 0x8b}),
				address:          0,
			},
			want:    "add\tx20, sp, w19, uxth",
			wantErr: false,
		},
		{
			name: "add	x12, x1, w20, uxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x40, 0x34, 0x8b}),
				address:          0,
			},
			want:    "add\tx12, x1, w20, uxtw",
			wantErr: false,
		},
		{
			name: "add	x20, x3, x13, uxtx",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x74, 0x60, 0x2d, 0x8b}),
				address:          0,
			},
			want:    "add\tx20, x3, x13, uxtx",
			wantErr: false,
		},
		{
			name: "add	x17, x25, w20, sxtb",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x31, 0x83, 0x34, 0x8b}),
				address:          0,
			},
			want:    "add\tx17, x25, w20, sxtb",
			wantErr: false,
		},
		{
			name: "add	x18, x13, w19, sxth",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb2, 0xa1, 0x33, 0x8b}),
				address:          0,
			},
			want:    "add\tx18, x13, w19, sxth",
			wantErr: false,
		},
		{
			name: "add	sp, x2, w3, sxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0xc0, 0x23, 0x8b}),
				address:          0,
			},
			want:    "add\tsp, x2, w3, sxtw",
			wantErr: false,
		},
		{
			name: "add	x3, x5, x9, sxtx",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xe0, 0x29, 0x8b}),
				address:          0,
			},
			want:    "add\tx3, x5, x9, sxtx",
			wantErr: false,
		},
		{
			name: "add	w2, w5, w7, uxtb",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa2, 0x00, 0x27, 0x0b}),
				address:          0,
			},
			want:    "add\tw2, w5, w7, uxtb",
			wantErr: false,
		},
		{
			name: "add	w21, w15, w17, uxth",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf5, 0x21, 0x31, 0x0b}),
				address:          0,
			},
			want:    "add\tw21, w15, w17, uxth",
			wantErr: false,
		},
		{
			name: "add	w30, w29, wzr, uxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbe, 0x43, 0x3f, 0x0b}),
				address:          0,
			},
			want:    "add\tw30, w29, wzr, uxtw",
			wantErr: false,
		},
		{
			name: "add	w19, w17, w1, uxtx",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x33, 0x62, 0x21, 0x0b}),
				address:          0,
			},
			want:    "add\tw19, w17, w1, uxtx",
			wantErr: false,
		},
		{
			name: "add	w2, w5, w1, sxtb",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa2, 0x80, 0x21, 0x0b}),
				address:          0,
			},
			want:    "add\tw2, w5, w1, sxtb",
			wantErr: false,
		},
		{
			name: "add	w26, w17, w19, sxth",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3a, 0xa2, 0x33, 0x0b}),
				address:          0,
			},
			want:    "add\tw26, w17, w19, sxth",
			wantErr: false,
		},
		{
			name: "add	w0, w2, w3, sxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0xc0, 0x23, 0x0b}),
				address:          0,
			},
			want:    "add\tw0, w2, w3, sxtw",
			wantErr: false,
		},
		{
			name: "add	w2, w3, w5, sxtx",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xe0, 0x25, 0x0b}),
				address:          0,
			},
			want:    "add\tw2, w3, w5, sxtx",
			wantErr: false,
		},
		{
			name: "add	x2, x3, w5, sxtb",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x80, 0x25, 0x8b}),
				address:          0,
			},
			want:    "add\tx2, x3, w5, sxtb",
			wantErr: false,
		},
		{
			name: "add	x7, x11, w13, uxth #4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x67, 0x31, 0x2d, 0x8b}),
				address:          0,
			},
			want:    "add\tx7, x11, w13, uxth #0x4",
			wantErr: false,
		},
		{
			name: "add	w17, w19, w23, uxtw #2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x71, 0x4a, 0x37, 0x0b}),
				address:          0,
			},
			want:    "add\tw17, w19, w23, uxtw #0x2",
			wantErr: false,
		},
		{
			name: "add	w29, w23, w17, uxtx #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfd, 0x66, 0x31, 0x0b}),
				address:          0,
			},
			want:    "add\tw29, w23, w17, uxtx #0x1",
			wantErr: false,
		},
		{
			name: "sub	x2, x4, w5, uxtb #2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x82, 0x08, 0x25, 0xcb}),
				address:          0,
			},
			want:    "sub\tx2, x4, w5, uxtb #0x2",
			wantErr: false,
		},
		{
			name: "sub	x20, sp, w19, uxth #4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x33, 0x33, 0xcb}),
				address:          0,
			},
			want:    "sub\tx20, sp, w19, uxth #0x4",
			wantErr: false,
		},
		{
			name: "sub	x12, x1, w20, uxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x40, 0x34, 0xcb}),
				address:          0,
			},
			want:    "sub\tx12, x1, w20, uxtw",
			wantErr: false,
		},
		{
			name: "sub	x20, x3, x13, uxtx",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x74, 0x60, 0x2d, 0xcb}),
				address:          0,
			},
			want:    "sub\tx20, x3, x13, uxtx",
			wantErr: false,
		},
		{
			name: "sub	x17, x25, w20, sxtb",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x31, 0x83, 0x34, 0xcb}),
				address:          0,
			},
			want:    "sub\tx17, x25, w20, sxtb",
			wantErr: false,
		},
		{
			name: "sub	x18, x13, w19, sxth",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb2, 0xa1, 0x33, 0xcb}),
				address:          0,
			},
			want:    "sub\tx18, x13, w19, sxth",
			wantErr: false,
		},
		{
			name: "sub	sp, x2, w3, sxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0xc0, 0x23, 0xcb}),
				address:          0,
			},
			want:    "sub\tsp, x2, w3, sxtw",
			wantErr: false,
		},
		{
			name: "sub	x3, x5, x9, sxtx",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xe0, 0x29, 0xcb}),
				address:          0,
			},
			want:    "sub\tx3, x5, x9, sxtx",
			wantErr: false,
		},
		{
			name: "sub	w2, w5, w7, uxtb",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa2, 0x00, 0x27, 0x4b}),
				address:          0,
			},
			want:    "sub\tw2, w5, w7, uxtb",
			wantErr: false,
		},
		{
			name: "sub	w21, w15, w17, uxth",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf5, 0x21, 0x31, 0x4b}),
				address:          0,
			},
			want:    "sub\tw21, w15, w17, uxth",
			wantErr: false,
		},
		{
			name: "sub	w30, w29, wzr, uxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbe, 0x43, 0x3f, 0x4b}),
				address:          0,
			},
			want:    "sub\tw30, w29, wzr, uxtw",
			wantErr: false,
		},
		{
			name: "sub	w19, w17, w1, uxtx",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x33, 0x62, 0x21, 0x4b}),
				address:          0,
			},
			want:    "sub\tw19, w17, w1, uxtx",
			wantErr: false,
		},
		{
			name: "sub	w2, w5, w1, sxtb",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa2, 0x80, 0x21, 0x4b}),
				address:          0,
			},
			want:    "sub\tw2, w5, w1, sxtb",
			wantErr: false,
		},
		{
			name: "sub	w26, wsp, w19, sxth",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfa, 0xa3, 0x33, 0x4b}),
				address:          0,
			},
			want:    "sub\tw26, wsp, w19, sxth",
			wantErr: false,
		},
		{
			name: "sub	wsp, w2, w3, sxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0xc0, 0x23, 0x4b}),
				address:          0,
			},
			want:    "sub\twsp, w2, w3, sxtw",
			wantErr: false,
		},
		{
			name: "sub	w2, w3, w5, sxtx",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xe0, 0x25, 0x4b}),
				address:          0,
			},
			want:    "sub\tw2, w3, w5, sxtx",
			wantErr: false,
		},
		{
			name: "adds	x2, x4, w5, uxtb #2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x82, 0x08, 0x25, 0xab}),
				address:          0,
			},
			want:    "adds\tx2, x4, w5, uxtb #0x2",
			wantErr: false,
		},
		{
			name: "adds	x20, sp, w19, uxth #4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x33, 0x33, 0xab}),
				address:          0,
			},
			want:    "adds\tx20, sp, w19, uxth #0x4",
			wantErr: false,
		},
		{
			name: "adds	x12, x1, w20, uxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x40, 0x34, 0xab}),
				address:          0,
			},
			want:    "adds\tx12, x1, w20, uxtw",
			wantErr: false,
		},
		{
			name: "adds	x20, x3, x13, uxtx",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x74, 0x60, 0x2d, 0xab}),
				address:          0,
			},
			want:    "adds\tx20, x3, x13, uxtx",
			wantErr: false,
		},
		{
			name: "cmn	x25, w20, sxtb #3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0x8f, 0x34, 0xab}),
				address:          0,
			},
			want:    "cmn\tx25, w20, sxtb #0x3",
			wantErr: false,
		},
		{
			name: "adds	x18, sp, w19, sxth",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf2, 0xa3, 0x33, 0xab}),
				address:          0,
			},
			want:    "adds\tx18, sp, w19, sxth",
			wantErr: false,
		},
		{
			name: "cmn	x2, w3, sxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0xc0, 0x23, 0xab}),
				address:          0,
			},
			want:    "cmn\tx2, w3, sxtw",
			wantErr: false,
		},
		{
			name: "adds	x3, x5, x9, sxtx #2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xe8, 0x29, 0xab}),
				address:          0,
			},
			want:    "adds\tx3, x5, x9, sxtx #0x2",
			wantErr: false,
		},
		{
			name: "adds	w2, w5, w7, uxtb",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa2, 0x00, 0x27, 0x2b}),
				address:          0,
			},
			want:    "adds\tw2, w5, w7, uxtb",
			wantErr: false,
		},
		{
			name: "adds	w21, w15, w17, uxth",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf5, 0x21, 0x31, 0x2b}),
				address:          0,
			},
			want:    "adds\tw21, w15, w17, uxth",
			wantErr: false,
		},
		{
			name: "adds	w30, w29, wzr, uxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbe, 0x43, 0x3f, 0x2b}),
				address:          0,
			},
			want:    "adds\tw30, w29, wzr, uxtw",
			wantErr: false,
		},
		{
			name: "adds	w19, w17, w1, uxtx",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x33, 0x62, 0x21, 0x2b}),
				address:          0,
			},
			want:    "adds\tw19, w17, w1, uxtx",
			wantErr: false,
		},
		{
			name: "adds	w2, w5, w1, sxtb #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa2, 0x84, 0x21, 0x2b}),
				address:          0,
			},
			want:    "adds\tw2, w5, w1, sxtb #0x1",
			wantErr: false,
		},
		{
			name: "adds	w26, wsp, w19, sxth",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfa, 0xa3, 0x33, 0x2b}),
				address:          0,
			},
			want:    "adds\tw26, wsp, w19, sxth",
			wantErr: false,
		},
		{
			name: "cmn	w2, w3, sxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0xc0, 0x23, 0x2b}),
				address:          0,
			},
			want:    "cmn\tw2, w3, sxtw",
			wantErr: false,
		},
		{
			name: "adds	w2, w3, w5, sxtx",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xe0, 0x25, 0x2b}),
				address:          0,
			},
			want:    "adds\tw2, w3, w5, sxtx",
			wantErr: false,
		},
		{
			name: "subs	x2, x4, w5, uxtb #2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x82, 0x08, 0x25, 0xeb}),
				address:          0,
			},
			want:    "subs\tx2, x4, w5, uxtb #0x2",
			wantErr: false,
		},
		{
			name: "subs	x20, sp, w19, uxth #4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x33, 0x33, 0xeb}),
				address:          0,
			},
			want:    "subs\tx20, sp, w19, uxth #0x4",
			wantErr: false,
		},
		{
			name: "subs	x12, x1, w20, uxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x40, 0x34, 0xeb}),
				address:          0,
			},
			want:    "subs\tx12, x1, w20, uxtw",
			wantErr: false,
		},
		{
			name: "subs	x20, x3, x13, uxtx",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x74, 0x60, 0x2d, 0xeb}),
				address:          0,
			},
			want:    "subs\tx20, x3, x13, uxtx",
			wantErr: false,
		},
		{
			name: "cmp	x25, w20, sxtb #3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0x8f, 0x34, 0xeb}),
				address:          0,
			},
			want:    "cmp\tx25, w20, sxtb #0x3",
			wantErr: false,
		},
		{
			name: "subs	x18, sp, w19, sxth",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf2, 0xa3, 0x33, 0xeb}),
				address:          0,
			},
			want:    "subs\tx18, sp, w19, sxth",
			wantErr: false,
		},
		{
			name: "cmp	x2, w3, sxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0xc0, 0x23, 0xeb}),
				address:          0,
			},
			want:    "cmp\tx2, w3, sxtw",
			wantErr: false,
		},
		{
			name: "subs	x3, x5, x9, sxtx #2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xe8, 0x29, 0xeb}),
				address:          0,
			},
			want:    "subs\tx3, x5, x9, sxtx #0x2",
			wantErr: false,
		},
		{
			name: "subs	w2, w5, w7, uxtb",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa2, 0x00, 0x27, 0x6b}),
				address:          0,
			},
			want:    "subs\tw2, w5, w7, uxtb",
			wantErr: false,
		},
		{
			name: "subs	w21, w15, w17, uxth",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf5, 0x21, 0x31, 0x6b}),
				address:          0,
			},
			want:    "subs\tw21, w15, w17, uxth",
			wantErr: false,
		},
		{
			name: "subs	w30, w29, wzr, uxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbe, 0x43, 0x3f, 0x6b}),
				address:          0,
			},
			want:    "subs\tw30, w29, wzr, uxtw",
			wantErr: false,
		},
		{
			name: "subs	w19, w17, w1, uxtx",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x33, 0x62, 0x21, 0x6b}),
				address:          0,
			},
			want:    "subs\tw19, w17, w1, uxtx",
			wantErr: false,
		},
		{
			name: "subs	w2, w5, w1, sxtb #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa2, 0x84, 0x21, 0x6b}),
				address:          0,
			},
			want:    "subs\tw2, w5, w1, sxtb #0x1",
			wantErr: false,
		},
		{
			name: "subs	w26, wsp, w19, sxth",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfa, 0xa3, 0x33, 0x6b}),
				address:          0,
			},
			want:    "subs\tw26, wsp, w19, sxth",
			wantErr: false,
		},
		{
			name: "cmp	w2, w3, sxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0xc0, 0x23, 0x6b}),
				address:          0,
			},
			want:    "cmp\tw2, w3, sxtw",
			wantErr: false,
		},
		{
			name: "subs	w2, w3, w5, sxtx",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xe0, 0x25, 0x6b}),
				address:          0,
			},
			want:    "subs\tw2, w3, w5, sxtx",
			wantErr: false,
		},
		{
			name: "cmp	x4, w5, uxtb #2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x08, 0x25, 0xeb}),
				address:          0,
			},
			want:    "cmp\tx4, w5, uxtb #0x2",
			wantErr: false,
		},
		{
			name: "cmp	sp, w19, uxth #4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x33, 0x33, 0xeb}),
				address:          0,
			},
			want:    "cmp\tsp, w19, uxth #0x4",
			wantErr: false,
		},
		{
			name: "cmp	x1, w20, uxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0x40, 0x34, 0xeb}),
				address:          0,
			},
			want:    "cmp\tx1, w20, uxtw",
			wantErr: false,
		},
		{
			name: "cmp	x3, x13, uxtx",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x60, 0x2d, 0xeb}),
				address:          0,
			},
			want:    "cmp\tx3, x13, uxtx",
			wantErr: false,
		},
		{
			name: "cmp	x25, w20, sxtb #3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0x8f, 0x34, 0xeb}),
				address:          0,
			},
			want:    "cmp\tx25, w20, sxtb #0x3",
			wantErr: false,
		},
		{
			name: "cmp	sp, w19, sxth",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xa3, 0x33, 0xeb}),
				address:          0,
			},
			want:    "cmp\tsp, w19, sxth",
			wantErr: false,
		},
		{
			name: "cmp	x2, w3, sxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0xc0, 0x23, 0xeb}),
				address:          0,
			},
			want:    "cmp\tx2, w3, sxtw",
			wantErr: false,
		},
		{
			name: "cmp	x5, x9, sxtx #2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0xe8, 0x29, 0xeb}),
				address:          0,
			},
			want:    "cmp\tx5, x9, sxtx #0x2",
			wantErr: false,
		},
		{
			name: "cmp	w5, w7, uxtb",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x00, 0x27, 0x6b}),
				address:          0,
			},
			want:    "cmp\tw5, w7, uxtb",
			wantErr: false,
		},
		{
			name: "cmp	w15, w17, uxth",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x21, 0x31, 0x6b}),
				address:          0,
			},
			want:    "cmp\tw15, w17, uxth",
			wantErr: false,
		},
		{
			name: "cmp	w29, wzr, uxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x43, 0x3f, 0x6b}),
				address:          0,
			},
			want:    "cmp\tw29, wzr, uxtw",
			wantErr: false,
		},
		{
			name: "cmp	w17, w1, uxtx",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0x62, 0x21, 0x6b}),
				address:          0,
			},
			want:    "cmp\tw17, w1, uxtx",
			wantErr: false,
		},
		{
			name: "cmp	w5, w1, sxtb #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x84, 0x21, 0x6b}),
				address:          0,
			},
			want:    "cmp\tw5, w1, sxtb #0x1",
			wantErr: false,
		},
		{
			name: "cmp	wsp, w19, sxth",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xa3, 0x33, 0x6b}),
				address:          0,
			},
			want:    "cmp\twsp, w19, sxth",
			wantErr: false,
		},
		{
			name: "cmp	w2, w3, sxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0xc0, 0x23, 0x6b}),
				address:          0,
			},
			want:    "cmp\tw2, w3, sxtw",
			wantErr: false,
		},
		{
			name: "cmp	w3, w5, sxtx",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0xe0, 0x25, 0x6b}),
				address:          0,
			},
			want:    "cmp\tw3, w5, sxtx",
			wantErr: false,
		},
		{
			name: "cmn	x4, w5, uxtb #2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x08, 0x25, 0xab}),
				address:          0,
			},
			want:    "cmn\tx4, w5, uxtb #0x2",
			wantErr: false,
		},
		{
			name: "cmn	sp, w19, uxth #4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x33, 0x33, 0xab}),
				address:          0,
			},
			want:    "cmn\tsp, w19, uxth #0x4",
			wantErr: false,
		},
		{
			name: "cmn	x1, w20, uxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0x40, 0x34, 0xab}),
				address:          0,
			},
			want:    "cmn\tx1, w20, uxtw",
			wantErr: false,
		},
		{
			name: "cmn	x3, x13, uxtx",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x60, 0x2d, 0xab}),
				address:          0,
			},
			want:    "cmn\tx3, x13, uxtx",
			wantErr: false,
		},
		{
			name: "cmn	x25, w20, sxtb #3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0x8f, 0x34, 0xab}),
				address:          0,
			},
			want:    "cmn\tx25, w20, sxtb #0x3",
			wantErr: false,
		},
		{
			name: "cmn	sp, w19, sxth",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xa3, 0x33, 0xab}),
				address:          0,
			},
			want:    "cmn\tsp, w19, sxth",
			wantErr: false,
		},
		{
			name: "cmn	x2, w3, sxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0xc0, 0x23, 0xab}),
				address:          0,
			},
			want:    "cmn\tx2, w3, sxtw",
			wantErr: false,
		},
		{
			name: "cmn	x5, x9, sxtx #2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0xe8, 0x29, 0xab}),
				address:          0,
			},
			want:    "cmn\tx5, x9, sxtx #0x2",
			wantErr: false,
		},
		{
			name: "cmn	w5, w7, uxtb",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x00, 0x27, 0x2b}),
				address:          0,
			},
			want:    "cmn\tw5, w7, uxtb",
			wantErr: false,
		},
		{
			name: "cmn	w15, w17, uxth",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x21, 0x31, 0x2b}),
				address:          0,
			},
			want:    "cmn\tw15, w17, uxth",
			wantErr: false,
		},
		{
			name: "cmn	w29, wzr, uxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x43, 0x3f, 0x2b}),
				address:          0,
			},
			want:    "cmn\tw29, wzr, uxtw",
			wantErr: false,
		},
		{
			name: "cmn	w17, w1, uxtx",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0x62, 0x21, 0x2b}),
				address:          0,
			},
			want:    "cmn\tw17, w1, uxtx",
			wantErr: false,
		},
		{
			name: "cmn	w5, w1, sxtb #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x84, 0x21, 0x2b}),
				address:          0,
			},
			want:    "cmn\tw5, w1, sxtb #0x1",
			wantErr: false,
		},
		{
			name: "cmn	wsp, w19, sxth",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xa3, 0x33, 0x2b}),
				address:          0,
			},
			want:    "cmn\twsp, w19, sxth",
			wantErr: false,
		},
		{
			name: "cmn	w2, w3, sxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0xc0, 0x23, 0x2b}),
				address:          0,
			},
			want:    "cmn\tw2, w3, sxtw",
			wantErr: false,
		},
		{
			name: "cmn	w3, w5, sxtx",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0xe0, 0x25, 0x2b}),
				address:          0,
			},
			want:    "cmn\tw3, w5, sxtx",
			wantErr: false,
		},
		{
			name: "cmp	x20, w29, uxtb #3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x0e, 0x3d, 0xeb}),
				address:          0,
			},
			want:    "cmp\tx20, w29, uxtb #0x3",
			wantErr: false,
		},
		{
			name: "cmp	x12, x13, uxtx #4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x71, 0x2d, 0xeb}),
				address:          0,
			},
			want:    "cmp\tx12, x13, uxtx #0x4",
			wantErr: false,
		},
		{
			name: "cmp	wsp, w1, uxtb",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x21, 0x6b}),
				address:          0,
			},
			want:    "cmp\twsp, w1, uxtb",
			wantErr: false,
		},
		{
			name: "cmn	wsp, wzr, sxtw",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xc3, 0x3f, 0x2b}),
				address:          0,
			},
			want:    "cmn\twsp, wzr, sxtw",
			wantErr: false,
		},
		{
			name: "sub	sp, x3, x7, lsl #4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x70, 0x27, 0xcb}),
				address:          0,
			},
			want:    "sub\tsp, x3, x7, lsl #0x4",
			wantErr: false,
		},
		{
			name: "add	w2, wsp, w3, lsl #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe2, 0x47, 0x23, 0x0b}),
				address:          0,
			},
			want:    "add\tw2, wsp, w3, lsl #0x1",
			wantErr: false,
		},
		{
			name: "cmp	wsp, w9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x43, 0x29, 0x6b}),
				address:          0,
			},
			want:    "cmp\twsp, w9, uxtw",
			wantErr: false,
		},
		{
			name: "cmn	wsp, w3, lsl #4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x53, 0x23, 0x2b}),
				address:          0,
			},
			want:    "cmn\twsp, w3, lsl #0x4",
			wantErr: false,
		},
		{
			name: "subs	x3, sp, x9, lsl #2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x6b, 0x29, 0xeb}),
				address:          0,
			},
			want:    "subs\tx3, sp, x9, lsl #0x2",
			wantErr: false,
		},
		{
			name: "add	w4, w5, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa4, 0x00, 0x00, 0x11}),
				address:          0,
			},
			want:    "add\tw4, w5, #0",
			wantErr: false,
		},
		{
			name: "add	w2, w3, #4095",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xfc, 0x3f, 0x11}),
				address:          0,
			},
			want:    "add\tw2, w3, #0xfff",
			wantErr: false,
		},
		{
			name: "add	w30, w29, #1, lsl #12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbe, 0x07, 0x40, 0x11}),
				address:          0,
			},
			want:    "add\tw30, w29, #0x1, lsl #0xc",
			wantErr: false,
		},
		{
			name: "add	w13, w5, #4095, lsl #12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xad, 0xfc, 0x7f, 0x11}),
				address:          0,
			},
			want:    "add\tw13, w5, #0xfff, lsl #0xc",
			wantErr: false,
		},
		{
			name: "add	x5, x7, #1638",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe5, 0x98, 0x19, 0x91}),
				address:          0,
			},
			want:    "add\tx5, x7, #0x666",
			wantErr: false,
		},
		{
			name: "add	w20, wsp, #801",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x87, 0x0c, 0x11}),
				address:          0,
			},
			want:    "add\tw20, wsp, #0x321",
			wantErr: false,
		},
		{
			name: "add	wsp, wsp, #1104",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x43, 0x11, 0x11}),
				address:          0,
			},
			want:    "add\twsp, wsp, #0x450",
			wantErr: false,
		},
		{
			name: "add	wsp, w30, #4084",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0xd3, 0x3f, 0x11}),
				address:          0,
			},
			want:    "add\twsp, w30, #0xff4",
			wantErr: false,
		},
		{
			name: "add	x0, x24, #291",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x8f, 0x04, 0x91}),
				address:          0,
			},
			want:    "add\tx0, x24, #0x123",
			wantErr: false,
		},
		{
			name: "add	x3, x24, #4095, lsl #12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x03, 0xff, 0x7f, 0x91}),
				address:          0,
			},
			want:    "add\tx3, x24, #0xfff, lsl #0xc",
			wantErr: false,
		},
		{
			name: "add	x8, sp, #1074",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe8, 0xcb, 0x10, 0x91}),
				address:          0,
			},
			want:    "add\tx8, sp, #0x432",
			wantErr: false,
		},
		{
			name: "add	sp, x29, #3816",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0xa3, 0x3b, 0x91}),
				address:          0,
			},
			want:    "add\tsp, fp, #0xee8",
			wantErr: false,
		},
		{
			name: "sub	w0, wsp, #4077",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0xb7, 0x3f, 0x51}),
				address:          0,
			},
			want:    "sub\tw0, wsp, #0xfed",
			wantErr: false,
		},
		{
			name: "sub	w4, w20, #546, lsl #12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x84, 0x8a, 0x48, 0x51}),
				address:          0,
			},
			want:    "sub\tw4, w20, #0x222, lsl #0xc",
			wantErr: false,
		},
		{
			name: "sub	sp, sp, #288",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x83, 0x04, 0xd1}),
				address:          0,
			},
			want:    "sub\tsp, sp, #0x120",
			wantErr: false,
		},
		{
			name: "sub	wsp, w19, #16",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x42, 0x00, 0x51}),
				address:          0,
			},
			want:    "sub\twsp, w19, #0x10",
			wantErr: false,
		},
		{
			name: "adds	w13, w23, #291, lsl #12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xed, 0x8e, 0x44, 0x31}),
				address:          0,
			},
			want:    "adds\tw13, w23, #0x123, lsl #0xc",
			wantErr: false,
		},
		{
			name: "cmn	w2, #4095",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0xfc, 0x3f, 0x31}),
				address:          0,
			},
			want:    "cmn\tw2, #0xfff",
			wantErr: false,
		},
		{
			name: "adds	w20, wsp, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x03, 0x00, 0x31}),
				address:          0,
			},
			want:    "adds\tw20, wsp, #0",
			wantErr: false,
		},
		{
			name: "cmn	x3, #1, lsl #12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x04, 0x40, 0xb1}),
				address:          0,
			},
			want:    "cmn\tx3, #0x1, lsl #0xc",
			wantErr: false,
		},
		{
			name: "cmp	sp, #20, lsl #12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x53, 0x40, 0xf1}),
				address:          0,
			},
			want:    "cmp\tsp, #0x14, lsl #0xc",
			wantErr: false,
		},
		{
			name: "cmp	x30, #4095",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0xff, 0x3f, 0xf1}),
				address:          0,
			},
			want:    "cmp\tlr, #0xfff",
			wantErr: false,
		},
		{
			name: "subs	x4, sp, #3822",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe4, 0xbb, 0x3b, 0xf1}),
				address:          0,
			},
			want:    "subs\tx4, sp, #0xeee",
			wantErr: false,
		},
		{
			name: "cmn	w3, #291, lsl #12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x8c, 0x44, 0x31}),
				address:          0,
			},
			want:    "cmn\tw3, #0x123, lsl #0xc",
			wantErr: false,
		},
		{
			name: "cmn	wsp, #1365",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x57, 0x15, 0x31}),
				address:          0,
			},
			want:    "cmn\twsp, #0x555",
			wantErr: false,
		},
		{
			name: "cmn	sp, #1092, lsl #12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x13, 0x51, 0xb1}),
				address:          0,
			},
			want:    "cmn\tsp, #0x444, lsl #0xc",
			wantErr: false,
		},
		{
			name: "cmp	x4, #300, lsl #12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0xb0, 0x44, 0xf1}),
				address:          0,
			},
			want:    "cmp\tx4, #0x12c, lsl #0xc",
			wantErr: false,
		},
		{
			name: "cmp	wsp, #500",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xd3, 0x07, 0x71}),
				address:          0,
			},
			want:    "cmp\twsp, #0x1f4",
			wantErr: false,
		},
		{
			name: "cmp	sp, #200",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x23, 0x03, 0xf1}),
				address:          0,
			},
			want:    "cmp\tsp, #0xc8",
			wantErr: false,
		},
		{
			name: "mov	sp, x30",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0x03, 0x00, 0x91}),
				address:          0,
			},
			want:    "mov\tsp, lr",
			wantErr: false,
		},
		{
			name: "mov	wsp, w20",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x02, 0x00, 0x11}),
				address:          0,
			},
			want:    "mov\twsp, w20",
			wantErr: false,
		},
		{
			name: "mov	x11, sp",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xeb, 0x03, 0x00, 0x91}),
				address:          0,
			},
			want:    "mov\tx11, sp",
			wantErr: false,
		},
		{
			name: "mov	w24, wsp",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf8, 0x03, 0x00, 0x11}),
				address:          0,
			},
			want:    "mov\tw24, wsp",
			wantErr: false,
		},
		{
			name: "add	w3, w5, w7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0x00, 0x07, 0x0b}),
				address:          0,
			},
			want:    "add\tw3, w5, w7",
			wantErr: false,
		},
		{
			name: "add	wzr, w3, w5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x00, 0x05, 0x0b}),
				address:          0,
			},
			want:    "add\twzr, w3, w5",
			wantErr: false,
		},
		{
			name: "add	w20, wzr, w4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x03, 0x04, 0x0b}),
				address:          0,
			},
			want:    "add\tw20, wzr, w4",
			wantErr: false,
		},
		{
			name: "add	w4, w6, wzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc4, 0x00, 0x1f, 0x0b}),
				address:          0,
			},
			want:    "add\tw4, w6, wzr",
			wantErr: false,
		},
		{
			name: "add	w11, w13, w15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xab, 0x01, 0x0f, 0x0b}),
				address:          0,
			},
			want:    "add\tw11, w13, w15",
			wantErr: false,
		},
		{
			name: "add	w9, w3, wzr, lsl #10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0x28, 0x1f, 0x0b}),
				address:          0,
			},
			want:    "add\tw9, w3, wzr, lsl #0xa",
			wantErr: false,
		},
		{
			name: "add	w17, w29, w20, lsl #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb1, 0x7f, 0x14, 0x0b}),
				address:          0,
			},
			want:    "add\tw17, w29, w20, lsl #0x1f",
			wantErr: false,
		},
		{
			name: "add	w17, w29, w20, lsl #29",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb1, 0x77, 0x14, 0x0b}),
				address:          0,
			},
			want:    "add\tw17, w29, w20, lsl #0x1d",
			wantErr: false,
		},
		{
			name: "add	w21, w22, w23, lsr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd5, 0x02, 0x57, 0x0b}),
				address:          0,
			},
			want:    "add\tw21, w22, w23, lsr #0",
			wantErr: false,
		},
		{
			name: "add	w24, w25, w26, lsr #18",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x38, 0x4b, 0x5a, 0x0b}),
				address:          0,
			},
			want:    "add\tw24, w25, w26, lsr #0x12",
			wantErr: false,
		},
		{
			name: "add	w27, w28, w29, lsr #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9b, 0x7f, 0x5d, 0x0b}),
				address:          0,
			},
			want:    "add\tw27, w28, w29, lsr #0x1f",
			wantErr: false,
		},
		{
			name: "add	w27, w28, w29, lsr #29",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9b, 0x77, 0x5d, 0x0b}),
				address:          0,
			},
			want:    "add\tw27, w28, w29, lsr #0x1d",
			wantErr: false,
		},
		{
			name: "add	w2, w3, w4, asr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x00, 0x84, 0x0b}),
				address:          0,
			},
			want:    "add\tw2, w3, w4, asr #0",
			wantErr: false,
		},
		{
			name: "add	w5, w6, w7, asr #21",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc5, 0x54, 0x87, 0x0b}),
				address:          0,
			},
			want:    "add\tw5, w6, w7, asr #0x15",
			wantErr: false,
		},
		{
			name: "add	w8, w9, w10, asr #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x28, 0x7d, 0x8a, 0x0b}),
				address:          0,
			},
			want:    "add\tw8, w9, w10, asr #0x1f",
			wantErr: false,
		},
		{
			name: "add	w8, w9, w10, asr #29",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x28, 0x75, 0x8a, 0x0b}),
				address:          0,
			},
			want:    "add\tw8, w9, w10, asr #0x1d",
			wantErr: false,
		},
		{
			name: "add	x3, x5, x7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0x00, 0x07, 0x8b}),
				address:          0,
			},
			want:    "add\tx3, x5, x7",
			wantErr: false,
		},
		{
			name: "add	xzr, x3, x5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x00, 0x05, 0x8b}),
				address:          0,
			},
			want:    "add\txzr, x3, x5",
			wantErr: false,
		},
		{
			name: "add	x20, xzr, x4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x03, 0x04, 0x8b}),
				address:          0,
			},
			want:    "add\tx20, xzr, x4",
			wantErr: false,
		},
		{
			name: "add	x4, x6, xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc4, 0x00, 0x1f, 0x8b}),
				address:          0,
			},
			want:    "add\tx4, x6, xzr",
			wantErr: false,
		},
		{
			name: "add	x11, x13, x15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xab, 0x01, 0x0f, 0x8b}),
				address:          0,
			},
			want:    "add\tx11, x13, x15",
			wantErr: false,
		},
		{
			name: "add	x9, x3, xzr, lsl #10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0x28, 0x1f, 0x8b}),
				address:          0,
			},
			want:    "add\tx9, x3, xzr, lsl #0xa",
			wantErr: false,
		},
		{
			name: "add	x17, x29, x20, lsl #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb1, 0xff, 0x14, 0x8b}),
				address:          0,
			},
			want:    "add\tx17, fp, x20, lsl #0x3f",
			wantErr: false,
		},
		{
			name: "add	x17, x29, x20, lsl #58",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb1, 0xeb, 0x14, 0x8b}),
				address:          0,
			},
			want:    "add\tx17, fp, x20, lsl #0x3a",
			wantErr: false,
		},
		{
			name: "add	x21, x22, x23, lsr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd5, 0x02, 0x57, 0x8b}),
				address:          0,
			},
			want:    "add\tx21, x22, x23, lsr #0",
			wantErr: false,
		},
		{
			name: "add	x24, x25, x26, lsr #18",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x38, 0x4b, 0x5a, 0x8b}),
				address:          0,
			},
			want:    "add\tx24, x25, x26, lsr #0x12",
			wantErr: false,
		},
		{
			name: "add	x27, x28, x29, lsr #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9b, 0xff, 0x5d, 0x8b}),
				address:          0,
			},
			want:    "add\tx27, x28, fp, lsr #0x3f",
			wantErr: false,
		},
		{
			name: "add	x17, x29, x20, lsr #58",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb1, 0xeb, 0x54, 0x8b}),
				address:          0,
			},
			want:    "add\tx17, fp, x20, lsr #0x3a",
			wantErr: false,
		},
		{
			name: "add	x2, x3, x4, asr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x00, 0x84, 0x8b}),
				address:          0,
			},
			want:    "add\tx2, x3, x4, asr #0",
			wantErr: false,
		},
		{
			name: "add	x5, x6, x7, asr #21",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc5, 0x54, 0x87, 0x8b}),
				address:          0,
			},
			want:    "add\tx5, x6, x7, asr #0x15",
			wantErr: false,
		},
		{
			name: "add	x8, x9, x10, asr #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x28, 0xfd, 0x8a, 0x8b}),
				address:          0,
			},
			want:    "add\tx8, x9, x10, asr #0x3f",
			wantErr: false,
		},
		{
			name: "add	x17, x29, x20, asr #58",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb1, 0xeb, 0x94, 0x8b}),
				address:          0,
			},
			want:    "add\tx17, fp, x20, asr #0x3a",
			wantErr: false,
		},
		{
			name: "adds	w3, w5, w7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0x00, 0x07, 0x2b}),
				address:          0,
			},
			want:    "adds\tw3, w5, w7",
			wantErr: false,
		},
		{
			name: "cmn	w3, w5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x00, 0x05, 0x2b}),
				address:          0,
			},
			want:    "cmn\tw3, w5",
			wantErr: false,
		},
		{
			name: "adds	w20, wzr, w4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x03, 0x04, 0x2b}),
				address:          0,
			},
			want:    "adds\tw20, wzr, w4",
			wantErr: false,
		},
		{
			name: "adds	w4, w6, wzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc4, 0x00, 0x1f, 0x2b}),
				address:          0,
			},
			want:    "adds\tw4, w6, wzr",
			wantErr: false,
		},
		{
			name: "adds	w11, w13, w15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xab, 0x01, 0x0f, 0x2b}),
				address:          0,
			},
			want:    "adds\tw11, w13, w15",
			wantErr: false,
		},
		{
			name: "adds	w9, w3, wzr, lsl #10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0x28, 0x1f, 0x2b}),
				address:          0,
			},
			want:    "adds\tw9, w3, wzr, lsl #0xa",
			wantErr: false,
		},
		{
			name: "adds	w17, w29, w20, lsl #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb1, 0x7f, 0x14, 0x2b}),
				address:          0,
			},
			want:    "adds\tw17, w29, w20, lsl #0x1f",
			wantErr: false,
		},
		{
			name: "adds	w21, w22, w23, lsr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd5, 0x02, 0x57, 0x2b}),
				address:          0,
			},
			want:    "adds\tw21, w22, w23, lsr #0",
			wantErr: false,
		},
		{
			name: "adds	w24, w25, w26, lsr #18",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x38, 0x4b, 0x5a, 0x2b}),
				address:          0,
			},
			want:    "adds\tw24, w25, w26, lsr #0x12",
			wantErr: false,
		},
		{
			name: "adds	w27, w28, w29, lsr #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9b, 0x7f, 0x5d, 0x2b}),
				address:          0,
			},
			want:    "adds\tw27, w28, w29, lsr #0x1f",
			wantErr: false,
		},
		{
			name: "adds	w2, w3, w4, asr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x00, 0x84, 0x2b}),
				address:          0,
			},
			want:    "adds\tw2, w3, w4, asr #0",
			wantErr: false,
		},
		{
			name: "adds	w5, w6, w7, asr #21",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc5, 0x54, 0x87, 0x2b}),
				address:          0,
			},
			want:    "adds\tw5, w6, w7, asr #0x15",
			wantErr: false,
		},
		{
			name: "adds	w8, w9, w10, asr #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x28, 0x7d, 0x8a, 0x2b}),
				address:          0,
			},
			want:    "adds\tw8, w9, w10, asr #0x1f",
			wantErr: false,
		},
		{
			name: "adds	x3, x5, x7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0x00, 0x07, 0xab}),
				address:          0,
			},
			want:    "adds\tx3, x5, x7",
			wantErr: false,
		},
		{
			name: "cmn	x3, x5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x00, 0x05, 0xab}),
				address:          0,
			},
			want:    "cmn\tx3, x5",
			wantErr: false,
		},
		{
			name: "adds	x20, xzr, x4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x03, 0x04, 0xab}),
				address:          0,
			},
			want:    "adds\tx20, xzr, x4",
			wantErr: false,
		},
		{
			name: "adds	x4, x6, xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc4, 0x00, 0x1f, 0xab}),
				address:          0,
			},
			want:    "adds\tx4, x6, xzr",
			wantErr: false,
		},
		{
			name: "adds	x11, x13, x15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xab, 0x01, 0x0f, 0xab}),
				address:          0,
			},
			want:    "adds\tx11, x13, x15",
			wantErr: false,
		},
		{
			name: "adds	x9, x3, xzr, lsl #10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0x28, 0x1f, 0xab}),
				address:          0,
			},
			want:    "adds\tx9, x3, xzr, lsl #0xa",
			wantErr: false,
		},
		{
			name: "adds	x17, x29, x20, lsl #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb1, 0xff, 0x14, 0xab}),
				address:          0,
			},
			want:    "adds\tx17, fp, x20, lsl #0x3f",
			wantErr: false,
		},
		{
			name: "adds	x21, x22, x23, lsr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd5, 0x02, 0x57, 0xab}),
				address:          0,
			},
			want:    "adds\tx21, x22, x23, lsr #0",
			wantErr: false,
		},
		{
			name: "adds	x24, x25, x26, lsr #18",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x38, 0x4b, 0x5a, 0xab}),
				address:          0,
			},
			want:    "adds\tx24, x25, x26, lsr #0x12",
			wantErr: false,
		},
		{
			name: "adds	x27, x28, x29, lsr #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9b, 0xff, 0x5d, 0xab}),
				address:          0,
			},
			want:    "adds\tx27, x28, fp, lsr #0x3f",
			wantErr: false,
		},
		{
			name: "adds	x2, x3, x4, asr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x00, 0x84, 0xab}),
				address:          0,
			},
			want:    "adds\tx2, x3, x4, asr #0",
			wantErr: false,
		},
		{
			name: "adds	x5, x6, x7, asr #21",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc5, 0x54, 0x87, 0xab}),
				address:          0,
			},
			want:    "adds\tx5, x6, x7, asr #0x15",
			wantErr: false,
		},
		{
			name: "adds	x8, x9, x10, asr #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x28, 0xfd, 0x8a, 0xab}),
				address:          0,
			},
			want:    "adds\tx8, x9, x10, asr #0x3f",
			wantErr: false,
		},
		{
			name: "sub	w3, w5, w7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0x00, 0x07, 0x4b}),
				address:          0,
			},
			want:    "sub\tw3, w5, w7",
			wantErr: false,
		},
		{
			name: "sub	wzr, w3, w5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x00, 0x05, 0x4b}),
				address:          0,
			},
			want:    "sub\twzr, w3, w5",
			wantErr: false,
		},
		{
			name: "neg	w20, w4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x03, 0x04, 0x4b}),
				address:          0,
			},
			want:    "neg\tw20, w4",
			wantErr: false,
		},
		{
			name: "sub	w4, w6, wzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc4, 0x00, 0x1f, 0x4b}),
				address:          0,
			},
			want:    "sub\tw4, w6, wzr",
			wantErr: false,
		},
		{
			name: "sub	w11, w13, w15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xab, 0x01, 0x0f, 0x4b}),
				address:          0,
			},
			want:    "sub\tw11, w13, w15",
			wantErr: false,
		},
		{
			name: "sub	w9, w3, wzr, lsl #10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0x28, 0x1f, 0x4b}),
				address:          0,
			},
			want:    "sub\tw9, w3, wzr, lsl #0xa",
			wantErr: false,
		},
		{
			name: "sub	w17, w29, w20, lsl #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb1, 0x7f, 0x14, 0x4b}),
				address:          0,
			},
			want:    "sub\tw17, w29, w20, lsl #0x1f",
			wantErr: false,
		},
		{
			name: "sub	w21, w22, w23, lsr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd5, 0x02, 0x57, 0x4b}),
				address:          0,
			},
			want:    "sub\tw21, w22, w23, lsr #0",
			wantErr: false,
		},
		{
			name: "sub	w24, w25, w26, lsr #18",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x38, 0x4b, 0x5a, 0x4b}),
				address:          0,
			},
			want:    "sub\tw24, w25, w26, lsr #0x12",
			wantErr: false,
		},
		{
			name: "sub	w27, w28, w29, lsr #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9b, 0x7f, 0x5d, 0x4b}),
				address:          0,
			},
			want:    "sub\tw27, w28, w29, lsr #0x1f",
			wantErr: false,
		},
		{
			name: "sub	w2, w3, w4, asr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x00, 0x84, 0x4b}),
				address:          0,
			},
			want:    "sub\tw2, w3, w4, asr #0",
			wantErr: false,
		},
		{
			name: "sub	w5, w6, w7, asr #21",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc5, 0x54, 0x87, 0x4b}),
				address:          0,
			},
			want:    "sub\tw5, w6, w7, asr #0x15",
			wantErr: false,
		},
		{
			name: "sub	w8, w9, w10, asr #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x28, 0x7d, 0x8a, 0x4b}),
				address:          0,
			},
			want:    "sub\tw8, w9, w10, asr #0x1f",
			wantErr: false,
		},
		{
			name: "sub	x3, x5, x7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0x00, 0x07, 0xcb}),
				address:          0,
			},
			want:    "sub\tx3, x5, x7",
			wantErr: false,
		},
		{
			name: "sub	xzr, x3, x5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x00, 0x05, 0xcb}),
				address:          0,
			},
			want:    "sub\txzr, x3, x5",
			wantErr: false,
		},
		{
			name: "neg	x20, x4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x03, 0x04, 0xcb}),
				address:          0,
			},
			want:    "neg\tx20, x4",
			wantErr: false,
		},
		{
			name: "sub	x4, x6, xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc4, 0x00, 0x1f, 0xcb}),
				address:          0,
			},
			want:    "sub\tx4, x6, xzr",
			wantErr: false,
		},
		{
			name: "sub	x11, x13, x15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xab, 0x01, 0x0f, 0xcb}),
				address:          0,
			},
			want:    "sub\tx11, x13, x15",
			wantErr: false,
		},
		{
			name: "sub	x9, x3, xzr, lsl #10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0x28, 0x1f, 0xcb}),
				address:          0,
			},
			want:    "sub\tx9, x3, xzr, lsl #0xa",
			wantErr: false,
		},
		{
			name: "sub	x17, x29, x20, lsl #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb1, 0xff, 0x14, 0xcb}),
				address:          0,
			},
			want:    "sub\tx17, fp, x20, lsl #0x3f",
			wantErr: false,
		},
		{
			name: "sub	x21, x22, x23, lsr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd5, 0x02, 0x57, 0xcb}),
				address:          0,
			},
			want:    "sub\tx21, x22, x23, lsr #0",
			wantErr: false,
		},
		{
			name: "sub	x24, x25, x26, lsr #18",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x38, 0x4b, 0x5a, 0xcb}),
				address:          0,
			},
			want:    "sub\tx24, x25, x26, lsr #0x12",
			wantErr: false,
		},
		{
			name: "sub	x27, x28, x29, lsr #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9b, 0xff, 0x5d, 0xcb}),
				address:          0,
			},
			want:    "sub\tx27, x28, fp, lsr #0x3f",
			wantErr: false,
		},
		{
			name: "sub	x2, x3, x4, asr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x00, 0x84, 0xcb}),
				address:          0,
			},
			want:    "sub\tx2, x3, x4, asr #0",
			wantErr: false,
		},
		{
			name: "sub	x5, x6, x7, asr #21",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc5, 0x54, 0x87, 0xcb}),
				address:          0,
			},
			want:    "sub\tx5, x6, x7, asr #0x15",
			wantErr: false,
		},
		{
			name: "sub	x8, x9, x10, asr #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x28, 0xfd, 0x8a, 0xcb}),
				address:          0,
			},
			want:    "sub\tx8, x9, x10, asr #0x3f",
			wantErr: false,
		},
		{
			name: "subs	w3, w5, w7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0x00, 0x07, 0x6b}),
				address:          0,
			},
			want:    "subs\tw3, w5, w7",
			wantErr: false,
		},
		{
			name: "cmp	w3, w5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x00, 0x05, 0x6b}),
				address:          0,
			},
			want:    "cmp\tw3, w5",
			wantErr: false,
		},
		{
			name: "negs	w20, w4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x03, 0x04, 0x6b}),
				address:          0,
			},
			want:    "negs\tw20, w4",
			wantErr: false,
		},
		{
			name: "subs	w4, w6, wzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc4, 0x00, 0x1f, 0x6b}),
				address:          0,
			},
			want:    "subs\tw4, w6, wzr",
			wantErr: false,
		},
		{
			name: "subs	w11, w13, w15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xab, 0x01, 0x0f, 0x6b}),
				address:          0,
			},
			want:    "subs\tw11, w13, w15",
			wantErr: false,
		},
		{
			name: "subs	w9, w3, wzr, lsl #10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0x28, 0x1f, 0x6b}),
				address:          0,
			},
			want:    "subs\tw9, w3, wzr, lsl #0xa",
			wantErr: false,
		},
		{
			name: "subs	w17, w29, w20, lsl #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb1, 0x7f, 0x14, 0x6b}),
				address:          0,
			},
			want:    "subs\tw17, w29, w20, lsl #0x1f",
			wantErr: false,
		},
		{
			name: "subs	w21, w22, w23, lsr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd5, 0x02, 0x57, 0x6b}),
				address:          0,
			},
			want:    "subs\tw21, w22, w23, lsr #0",
			wantErr: false,
		},
		{
			name: "subs	w24, w25, w26, lsr #18",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x38, 0x4b, 0x5a, 0x6b}),
				address:          0,
			},
			want:    "subs\tw24, w25, w26, lsr #0x12",
			wantErr: false,
		},
		{
			name: "subs	w27, w28, w29, lsr #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9b, 0x7f, 0x5d, 0x6b}),
				address:          0,
			},
			want:    "subs\tw27, w28, w29, lsr #0x1f",
			wantErr: false,
		},
		{
			name: "subs	w2, w3, w4, asr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x00, 0x84, 0x6b}),
				address:          0,
			},
			want:    "subs\tw2, w3, w4, asr #0",
			wantErr: false,
		},
		{
			name: "subs	w5, w6, w7, asr #21",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc5, 0x54, 0x87, 0x6b}),
				address:          0,
			},
			want:    "subs\tw5, w6, w7, asr #0x15",
			wantErr: false,
		},
		{
			name: "subs	w8, w9, w10, asr #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x28, 0x7d, 0x8a, 0x6b}),
				address:          0,
			},
			want:    "subs\tw8, w9, w10, asr #0x1f",
			wantErr: false,
		},
		{
			name: "subs	x3, x5, x7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0x00, 0x07, 0xeb}),
				address:          0,
			},
			want:    "subs\tx3, x5, x7",
			wantErr: false,
		},
		{
			name: "cmp	x3, x5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x00, 0x05, 0xeb}),
				address:          0,
			},
			want:    "cmp\tx3, x5",
			wantErr: false,
		},
		{
			name: "negs	x20, x4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x03, 0x04, 0xeb}),
				address:          0,
			},
			want:    "negs\tx20, x4",
			wantErr: false,
		},
		{
			name: "subs	x4, x6, xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc4, 0x00, 0x1f, 0xeb}),
				address:          0,
			},
			want:    "subs\tx4, x6, xzr",
			wantErr: false,
		},
		{
			name: "subs	x11, x13, x15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xab, 0x01, 0x0f, 0xeb}),
				address:          0,
			},
			want:    "subs\tx11, x13, x15",
			wantErr: false,
		},
		{
			name: "subs	x9, x3, xzr, lsl #10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0x28, 0x1f, 0xeb}),
				address:          0,
			},
			want:    "subs\tx9, x3, xzr, lsl #0xa",
			wantErr: false,
		},
		{
			name: "subs	x17, x29, x20, lsl #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb1, 0xff, 0x14, 0xeb}),
				address:          0,
			},
			want:    "subs\tx17, fp, x20, lsl #0x3f",
			wantErr: false,
		},
		{
			name: "subs	x21, x22, x23, lsr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd5, 0x02, 0x57, 0xeb}),
				address:          0,
			},
			want:    "subs\tx21, x22, x23, lsr #0",
			wantErr: false,
		},
		{
			name: "subs	x24, x25, x26, lsr #18",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x38, 0x4b, 0x5a, 0xeb}),
				address:          0,
			},
			want:    "subs\tx24, x25, x26, lsr #0x12",
			wantErr: false,
		},
		{
			name: "subs	x27, x28, x29, lsr #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9b, 0xff, 0x5d, 0xeb}),
				address:          0,
			},
			want:    "subs\tx27, x28, fp, lsr #0x3f",
			wantErr: false,
		},
		{
			name: "subs	x2, x3, x4, asr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x00, 0x84, 0xeb}),
				address:          0,
			},
			want:    "subs\tx2, x3, x4, asr #0",
			wantErr: false,
		},
		{
			name: "subs	x5, x6, x7, asr #21",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc5, 0x54, 0x87, 0xeb}),
				address:          0,
			},
			want:    "subs\tx5, x6, x7, asr #0x15",
			wantErr: false,
		},
		{
			name: "subs	x8, x9, x10, asr #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x28, 0xfd, 0x8a, 0xeb}),
				address:          0,
			},
			want:    "subs\tx8, x9, x10, asr #0x3f",
			wantErr: false,
		},
		{
			name: "cmn	w0, w3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x00, 0x03, 0x2b}),
				address:          0,
			},
			want:    "cmn\tw0, w3",
			wantErr: false,
		},
		{
			name: "cmn	wzr, w4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x04, 0x2b}),
				address:          0,
			},
			want:    "cmn\twzr, w4",
			wantErr: false,
		},
		{
			name: "cmn	w5, wzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x00, 0x1f, 0x2b}),
				address:          0,
			},
			want:    "cmn\tw5, wzr",
			wantErr: false,
		},
		{
			name: "cmn	wsp, w6",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x43, 0x26, 0x2b}),
				address:          0,
			},
			want:    "cmn\twsp, w6, uxtw",
			wantErr: false,
		},
		{
			name: "cmn	w6, w7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0x00, 0x07, 0x2b}),
				address:          0,
			},
			want:    "cmn\tw6, w7",
			wantErr: false,
		},
		{
			name: "cmn	w8, w9, lsl #15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x3d, 0x09, 0x2b}),
				address:          0,
			},
			want:    "cmn\tw8, w9, lsl #0xf",
			wantErr: false,
		},
		{
			name: "cmn	w10, w11, lsl #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x7d, 0x0b, 0x2b}),
				address:          0,
			},
			want:    "cmn\tw10, w11, lsl #0x1f",
			wantErr: false,
		},
		{
			name: "cmn	w12, w13, lsr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x01, 0x4d, 0x2b}),
				address:          0,
			},
			want:    "cmn\tw12, w13, lsr #0",
			wantErr: false,
		},
		{
			name: "cmn	w14, w15, lsr #21",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0x55, 0x4f, 0x2b}),
				address:          0,
			},
			want:    "cmn\tw14, w15, lsr #0x15",
			wantErr: false,
		},
		{
			name: "cmn	w16, w17, lsr #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x7e, 0x51, 0x2b}),
				address:          0,
			},
			want:    "cmn\tw16, w17, lsr #0x1f",
			wantErr: false,
		},
		{
			name: "cmn	w18, w19, asr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x02, 0x93, 0x2b}),
				address:          0,
			},
			want:    "cmn\tw18, w19, asr #0",
			wantErr: false,
		},
		{
			name: "cmn	w20, w21, asr #22",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x5a, 0x95, 0x2b}),
				address:          0,
			},
			want:    "cmn\tw20, w21, asr #0x16",
			wantErr: false,
		},
		{
			name: "cmn	w22, w23, asr #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0x7e, 0x97, 0x2b}),
				address:          0,
			},
			want:    "cmn\tw22, w23, asr #0x1f",
			wantErr: false,
		},
		{
			name: "cmn	x0, x3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x00, 0x03, 0xab}),
				address:          0,
			},
			want:    "cmn\tx0, x3",
			wantErr: false,
		},
		{
			name: "cmn	xzr, x4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x04, 0xab}),
				address:          0,
			},
			want:    "cmn\txzr, x4",
			wantErr: false,
		},
		{
			name: "cmn	x5, xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x00, 0x1f, 0xab}),
				address:          0,
			},
			want:    "cmn\tx5, xzr",
			wantErr: false,
		},
		{
			name: "cmn	sp, x6",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x63, 0x26, 0xab}),
				address:          0,
			},
			want:    "cmn\tsp, x6",
			wantErr: false,
		},
		{
			name: "cmn	x6, x7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0x00, 0x07, 0xab}),
				address:          0,
			},
			want:    "cmn\tx6, x7",
			wantErr: false,
		},
		{
			name: "cmn	x8, x9, lsl #15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x3d, 0x09, 0xab}),
				address:          0,
			},
			want:    "cmn\tx8, x9, lsl #0xf",
			wantErr: false,
		},
		{
			name: "cmn	x10, x11, lsl #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0xfd, 0x0b, 0xab}),
				address:          0,
			},
			want:    "cmn\tx10, x11, lsl #0x3f",
			wantErr: false,
		},
		{
			name: "cmn	x12, x13, lsr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x01, 0x4d, 0xab}),
				address:          0,
			},
			want:    "cmn\tx12, x13, lsr #0",
			wantErr: false,
		},
		{
			name: "cmn	x14, x15, lsr #41",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0xa5, 0x4f, 0xab}),
				address:          0,
			},
			want:    "cmn\tx14, x15, lsr #0x29",
			wantErr: false,
		},
		{
			name: "cmn	x16, x17, lsr #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0xfe, 0x51, 0xab}),
				address:          0,
			},
			want:    "cmn\tx16, x17, lsr #0x3f",
			wantErr: false,
		},
		{
			name: "cmn	x18, x19, asr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x02, 0x93, 0xab}),
				address:          0,
			},
			want:    "cmn\tx18, x19, asr #0",
			wantErr: false,
		},
		{
			name: "cmn	x20, x21, asr #55",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0xde, 0x95, 0xab}),
				address:          0,
			},
			want:    "cmn\tx20, x21, asr #0x37",
			wantErr: false,
		},
		{
			name: "cmn	x22, x23, asr #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0xfe, 0x97, 0xab}),
				address:          0,
			},
			want:    "cmn\tx22, x23, asr #0x3f",
			wantErr: false,
		},
		{
			name: "cmp	w0, w3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x00, 0x03, 0x6b}),
				address:          0,
			},
			want:    "cmp\tw0, w3",
			wantErr: false,
		},
		{
			name: "cmp	wzr, w4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x04, 0x6b}),
				address:          0,
			},
			want:    "cmp\twzr, w4",
			wantErr: false,
		},
		{
			name: "cmp	w5, wzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x00, 0x1f, 0x6b}),
				address:          0,
			},
			want:    "cmp\tw5, wzr",
			wantErr: false,
		},
		{
			name: "cmp	wsp, w6",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x43, 0x26, 0x6b}),
				address:          0,
			},
			want:    "cmp\twsp, w6, uxtw",
			wantErr: false,
		},
		{
			name: "cmp	w6, w7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0x00, 0x07, 0x6b}),
				address:          0,
			},
			want:    "cmp\tw6, w7",
			wantErr: false,
		},
		{
			name: "cmp	w8, w9, lsl #15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x3d, 0x09, 0x6b}),
				address:          0,
			},
			want:    "cmp\tw8, w9, lsl #0xf",
			wantErr: false,
		},
		{
			name: "cmp	w10, w11, lsl #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x7d, 0x0b, 0x6b}),
				address:          0,
			},
			want:    "cmp\tw10, w11, lsl #0x1f",
			wantErr: false,
		},
		{
			name: "cmp	w12, w13, lsr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x01, 0x4d, 0x6b}),
				address:          0,
			},
			want:    "cmp\tw12, w13, lsr #0",
			wantErr: false,
		},
		{
			name: "cmp	w14, w15, lsr #21",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0x55, 0x4f, 0x6b}),
				address:          0,
			},
			want:    "cmp\tw14, w15, lsr #0x15",
			wantErr: false,
		},
		{
			name: "cmp	w16, w17, lsr #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x7e, 0x51, 0x6b}),
				address:          0,
			},
			want:    "cmp\tw16, w17, lsr #0x1f",
			wantErr: false,
		},
		{
			name: "cmp	w18, w19, asr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x02, 0x93, 0x6b}),
				address:          0,
			},
			want:    "cmp\tw18, w19, asr #0",
			wantErr: false,
		},
		{
			name: "cmp	w20, w21, asr #22",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x5a, 0x95, 0x6b}),
				address:          0,
			},
			want:    "cmp\tw20, w21, asr #0x16",
			wantErr: false,
		},
		{
			name: "cmp	w22, w23, asr #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0x7e, 0x97, 0x6b}),
				address:          0,
			},
			want:    "cmp\tw22, w23, asr #0x1f",
			wantErr: false,
		},
		{
			name: "cmp	x0, x3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x00, 0x03, 0xeb}),
				address:          0,
			},
			want:    "cmp\tx0, x3",
			wantErr: false,
		},
		{
			name: "cmp	xzr, x4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x04, 0xeb}),
				address:          0,
			},
			want:    "cmp\txzr, x4",
			wantErr: false,
		},
		{
			name: "cmp	x5, xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x00, 0x1f, 0xeb}),
				address:          0,
			},
			want:    "cmp\tx5, xzr",
			wantErr: false,
		},
		{
			name: "cmp	sp, x6",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x63, 0x26, 0xeb}),
				address:          0,
			},
			want:    "cmp\tsp, x6",
			wantErr: false,
		},
		{
			name: "cmp	x6, x7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0x00, 0x07, 0xeb}),
				address:          0,
			},
			want:    "cmp\tx6, x7",
			wantErr: false,
		},
		{
			name: "cmp	x8, x9, lsl #15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x3d, 0x09, 0xeb}),
				address:          0,
			},
			want:    "cmp\tx8, x9, lsl #0xf",
			wantErr: false,
		},
		{
			name: "cmp	x10, x11, lsl #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0xfd, 0x0b, 0xeb}),
				address:          0,
			},
			want:    "cmp\tx10, x11, lsl #0x3f",
			wantErr: false,
		},
		{
			name: "cmp	x12, x13, lsr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x01, 0x4d, 0xeb}),
				address:          0,
			},
			want:    "cmp\tx12, x13, lsr #0",
			wantErr: false,
		},
		{
			name: "cmp	x14, x15, lsr #41",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0xa5, 0x4f, 0xeb}),
				address:          0,
			},
			want:    "cmp\tx14, x15, lsr #0x29",
			wantErr: false,
		},
		{
			name: "cmp	x16, x17, lsr #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0xfe, 0x51, 0xeb}),
				address:          0,
			},
			want:    "cmp\tx16, x17, lsr #0x3f",
			wantErr: false,
		},
		{
			name: "cmp	x18, x19, asr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x02, 0x93, 0xeb}),
				address:          0,
			},
			want:    "cmp\tx18, x19, asr #0",
			wantErr: false,
		},
		{
			name: "cmp	x20, x21, asr #55",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0xde, 0x95, 0xeb}),
				address:          0,
			},
			want:    "cmp\tx20, x21, asr #0x37",
			wantErr: false,
		},
		{
			name: "cmp	x22, x23, asr #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0xfe, 0x97, 0xeb}),
				address:          0,
			},
			want:    "cmp\tx22, x23, asr #0x3f",
			wantErr: false,
		},
		{
			name: "neg	w29, w30",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfd, 0x03, 0x1e, 0x4b}),
				address:          0,
			},
			want:    "neg\tw29, w30",
			wantErr: false,
		},
		{
			name: "neg	w30, wzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfe, 0x03, 0x1f, 0x4b}),
				address:          0,
			},
			want:    "neg\tw30, wzr",
			wantErr: false,
		},
		{
			name: "neg	wzr, w0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x00, 0x4b}),
				address:          0,
			},
			want:    "neg\twzr, w0",
			wantErr: false,
		},
		{
			name: "neg	w28, w27",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfc, 0x03, 0x1b, 0x4b}),
				address:          0,
			},
			want:    "neg\tw28, w27",
			wantErr: false,
		},
		{
			name: "neg	w26, w25, lsl #29",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfa, 0x77, 0x19, 0x4b}),
				address:          0,
			},
			want:    "neg\tw26, w25, lsl #0x1d",
			wantErr: false,
		},
		{
			name: "neg	w24, w23, lsl #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf8, 0x7f, 0x17, 0x4b}),
				address:          0,
			},
			want:    "neg\tw24, w23, lsl #0x1f",
			wantErr: false,
		},
		{
			name: "neg	w22, w21, lsr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf6, 0x03, 0x55, 0x4b}),
				address:          0,
			},
			want:    "neg\tw22, w21, lsr #0",
			wantErr: false,
		},
		{
			name: "neg	w20, w19, lsr #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x07, 0x53, 0x4b}),
				address:          0,
			},
			want:    "neg\tw20, w19, lsr #0x1",
			wantErr: false,
		},
		{
			name: "neg	w18, w17, lsr #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf2, 0x7f, 0x51, 0x4b}),
				address:          0,
			},
			want:    "neg\tw18, w17, lsr #0x1f",
			wantErr: false,
		},
		{
			name: "neg	w16, w15, asr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf0, 0x03, 0x8f, 0x4b}),
				address:          0,
			},
			want:    "neg\tw16, w15, asr #0",
			wantErr: false,
		},
		{
			name: "neg	w14, w13, asr #12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xee, 0x33, 0x8d, 0x4b}),
				address:          0,
			},
			want:    "neg\tw14, w13, asr #0xc",
			wantErr: false,
		},
		{
			name: "neg	w12, w11, asr #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x7f, 0x8b, 0x4b}),
				address:          0,
			},
			want:    "neg\tw12, w11, asr #0x1f",
			wantErr: false,
		},
		{
			name: "neg	x29, x30",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfd, 0x03, 0x1e, 0xcb}),
				address:          0,
			},
			want:    "neg\tfp, lr",
			wantErr: false,
		},
		{
			name: "neg	x30, xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfe, 0x03, 0x1f, 0xcb}),
				address:          0,
			},
			want:    "neg\tlr, xzr",
			wantErr: false,
		},
		{
			name: "neg	xzr, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x00, 0xcb}),
				address:          0,
			},
			want:    "neg\txzr, x0",
			wantErr: false,
		},
		{
			name: "neg	x28, x27",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfc, 0x03, 0x1b, 0xcb}),
				address:          0,
			},
			want:    "neg\tx28, x27",
			wantErr: false,
		},
		{
			name: "neg	x26, x25, lsl #29",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfa, 0x77, 0x19, 0xcb}),
				address:          0,
			},
			want:    "neg\tx26, x25, lsl #0x1d",
			wantErr: false,
		},
		{
			name: "neg	x24, x23, lsl #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf8, 0x7f, 0x17, 0xcb}),
				address:          0,
			},
			want:    "neg\tx24, x23, lsl #0x1f",
			wantErr: false,
		},
		{
			name: "neg	x22, x21, lsr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf6, 0x03, 0x55, 0xcb}),
				address:          0,
			},
			want:    "neg\tx22, x21, lsr #0",
			wantErr: false,
		},
		{
			name: "neg	x20, x19, lsr #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x07, 0x53, 0xcb}),
				address:          0,
			},
			want:    "neg\tx20, x19, lsr #0x1",
			wantErr: false,
		},
		{
			name: "neg	x18, x17, lsr #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf2, 0x7f, 0x51, 0xcb}),
				address:          0,
			},
			want:    "neg\tx18, x17, lsr #0x1f",
			wantErr: false,
		},
		{
			name: "neg	x16, x15, asr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf0, 0x03, 0x8f, 0xcb}),
				address:          0,
			},
			want:    "neg\tx16, x15, asr #0",
			wantErr: false,
		},
		{
			name: "neg	x14, x13, asr #12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xee, 0x33, 0x8d, 0xcb}),
				address:          0,
			},
			want:    "neg\tx14, x13, asr #0xc",
			wantErr: false,
		},
		{
			name: "neg	x12, x11, asr #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x7f, 0x8b, 0xcb}),
				address:          0,
			},
			want:    "neg\tx12, x11, asr #0x1f",
			wantErr: false,
		},
		{
			name: "negs	w29, w30",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfd, 0x03, 0x1e, 0x6b}),
				address:          0,
			},
			want:    "negs\tw29, w30",
			wantErr: false,
		},
		{
			name: "negs	w30, wzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfe, 0x03, 0x1f, 0x6b}),
				address:          0,
			},
			want:    "negs\tw30, wzr",
			wantErr: false,
		},
		{
			name: "cmp	wzr, w0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x00, 0x6b}),
				address:          0,
			},
			want:    "cmp\twzr, w0",
			wantErr: false,
		},
		{
			name: "negs	w28, w27",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfc, 0x03, 0x1b, 0x6b}),
				address:          0,
			},
			want:    "negs\tw28, w27",
			wantErr: false,
		},
		{
			name: "negs	w26, w25, lsl #29",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfa, 0x77, 0x19, 0x6b}),
				address:          0,
			},
			want:    "negs\tw26, w25, lsl #0x1d",
			wantErr: false,
		},
		{
			name: "negs	w24, w23, lsl #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf8, 0x7f, 0x17, 0x6b}),
				address:          0,
			},
			want:    "negs\tw24, w23, lsl #0x1f",
			wantErr: false,
		},
		{
			name: "negs	w22, w21, lsr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf6, 0x03, 0x55, 0x6b}),
				address:          0,
			},
			want:    "negs\tw22, w21, lsr #0",
			wantErr: false,
		},
		{
			name: "negs	w20, w19, lsr #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x07, 0x53, 0x6b}),
				address:          0,
			},
			want:    "negs\tw20, w19, lsr #0x1",
			wantErr: false,
		},
		{
			name: "negs	w18, w17, lsr #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf2, 0x7f, 0x51, 0x6b}),
				address:          0,
			},
			want:    "negs\tw18, w17, lsr #0x1f",
			wantErr: false,
		},
		{
			name: "negs	w16, w15, asr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf0, 0x03, 0x8f, 0x6b}),
				address:          0,
			},
			want:    "negs\tw16, w15, asr #0",
			wantErr: false,
		},
		{
			name: "negs	w14, w13, asr #12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xee, 0x33, 0x8d, 0x6b}),
				address:          0,
			},
			want:    "negs\tw14, w13, asr #0xc",
			wantErr: false,
		},
		{
			name: "negs	w12, w11, asr #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x7f, 0x8b, 0x6b}),
				address:          0,
			},
			want:    "negs\tw12, w11, asr #0x1f",
			wantErr: false,
		},
		{
			name: "negs	x29, x30",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfd, 0x03, 0x1e, 0xeb}),
				address:          0,
			},
			want:    "negs\tfp, lr",
			wantErr: false,
		},
		{
			name: "negs	x30, xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfe, 0x03, 0x1f, 0xeb}),
				address:          0,
			},
			want:    "negs\tlr, xzr",
			wantErr: false,
		},
		{
			name: "cmp	xzr, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x00, 0xeb}),
				address:          0,
			},
			want:    "cmp\txzr, x0",
			wantErr: false,
		},
		{
			name: "negs	x28, x27",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfc, 0x03, 0x1b, 0xeb}),
				address:          0,
			},
			want:    "negs\tx28, x27",
			wantErr: false,
		},
		{
			name: "negs	x26, x25, lsl #29",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfa, 0x77, 0x19, 0xeb}),
				address:          0,
			},
			want:    "negs\tx26, x25, lsl #0x1d",
			wantErr: false,
		},
		{
			name: "negs	x24, x23, lsl #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf8, 0x7f, 0x17, 0xeb}),
				address:          0,
			},
			want:    "negs\tx24, x23, lsl #0x1f",
			wantErr: false,
		},
		{
			name: "negs	x22, x21, lsr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf6, 0x03, 0x55, 0xeb}),
				address:          0,
			},
			want:    "negs\tx22, x21, lsr #0",
			wantErr: false,
		},
		{
			name: "negs	x20, x19, lsr #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x07, 0x53, 0xeb}),
				address:          0,
			},
			want:    "negs\tx20, x19, lsr #0x1",
			wantErr: false,
		},
		{
			name: "negs	x18, x17, lsr #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf2, 0x7f, 0x51, 0xeb}),
				address:          0,
			},
			want:    "negs\tx18, x17, lsr #0x1f",
			wantErr: false,
		},
		{
			name: "negs	x16, x15, asr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf0, 0x03, 0x8f, 0xeb}),
				address:          0,
			},
			want:    "negs\tx16, x15, asr #0",
			wantErr: false,
		},
		{
			name: "negs	x14, x13, asr #12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xee, 0x33, 0x8d, 0xeb}),
				address:          0,
			},
			want:    "negs\tx14, x13, asr #0xc",
			wantErr: false,
		},
		{
			name: "negs	x12, x11, asr #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x7f, 0x8b, 0xeb}),
				address:          0,
			},
			want:    "negs\tx12, x11, asr #0x1f",
			wantErr: false,
		},
		{
			name: "adc	w29, w27, w25",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7d, 0x03, 0x19, 0x1a}),
				address:          0,
			},
			want:    "adc\tw29, w27, w25",
			wantErr: false,
		},
		{
			name: "adc	wzr, w3, w4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x00, 0x04, 0x1a}),
				address:          0,
			},
			want:    "adc\twzr, w3, w4",
			wantErr: false,
		},
		{
			name: "adc	w9, wzr, w10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x03, 0x0a, 0x1a}),
				address:          0,
			},
			want:    "adc\tw9, wzr, w10",
			wantErr: false,
		},
		{
			name: "adc	w20, w0, wzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x14, 0x00, 0x1f, 0x1a}),
				address:          0,
			},
			want:    "adc\tw20, w0, wzr",
			wantErr: false,
		},
		{
			name: "adc	x29, x27, x25",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7d, 0x03, 0x19, 0x9a}),
				address:          0,
			},
			want:    "adc\tfp, x27, x25",
			wantErr: false,
		},
		{
			name: "adc	xzr, x3, x4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x00, 0x04, 0x9a}),
				address:          0,
			},
			want:    "adc\txzr, x3, x4",
			wantErr: false,
		},
		{
			name: "adc	x9, xzr, x10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x03, 0x0a, 0x9a}),
				address:          0,
			},
			want:    "adc\tx9, xzr, x10",
			wantErr: false,
		},
		{
			name: "adc	x20, x0, xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x14, 0x00, 0x1f, 0x9a}),
				address:          0,
			},
			want:    "adc\tx20, x0, xzr",
			wantErr: false,
		},
		{
			name: "adcs	w29, w27, w25",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7d, 0x03, 0x19, 0x3a}),
				address:          0,
			},
			want:    "adcs\tw29, w27, w25",
			wantErr: false,
		},
		{
			name: "adcs	wzr, w3, w4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x00, 0x04, 0x3a}),
				address:          0,
			},
			want:    "adcs\twzr, w3, w4",
			wantErr: false,
		},
		{
			name: "adcs	w9, wzr, w10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x03, 0x0a, 0x3a}),
				address:          0,
			},
			want:    "adcs\tw9, wzr, w10",
			wantErr: false,
		},
		{
			name: "adcs	w20, w0, wzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x14, 0x00, 0x1f, 0x3a}),
				address:          0,
			},
			want:    "adcs\tw20, w0, wzr",
			wantErr: false,
		},
		{
			name: "adcs	x29, x27, x25",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7d, 0x03, 0x19, 0xba}),
				address:          0,
			},
			want:    "adcs\tfp, x27, x25",
			wantErr: false,
		},
		{
			name: "adcs	xzr, x3, x4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x00, 0x04, 0xba}),
				address:          0,
			},
			want:    "adcs\txzr, x3, x4",
			wantErr: false,
		},
		{
			name: "adcs	x9, xzr, x10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x03, 0x0a, 0xba}),
				address:          0,
			},
			want:    "adcs\tx9, xzr, x10",
			wantErr: false,
		},
		{
			name: "adcs	x20, x0, xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x14, 0x00, 0x1f, 0xba}),
				address:          0,
			},
			want:    "adcs\tx20, x0, xzr",
			wantErr: false,
		},
		{
			name: "sbc	w29, w27, w25",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7d, 0x03, 0x19, 0x5a}),
				address:          0,
			},
			want:    "sbc\tw29, w27, w25",
			wantErr: false,
		},
		{
			name: "sbc	wzr, w3, w4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x00, 0x04, 0x5a}),
				address:          0,
			},
			want:    "sbc\twzr, w3, w4",
			wantErr: false,
		},
		{
			name: "ngc	w9, w10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x03, 0x0a, 0x5a}),
				address:          0,
			},
			want:    "ngc\tw9, w10",
			wantErr: false,
		},
		{
			name: "sbc	w20, w0, wzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x14, 0x00, 0x1f, 0x5a}),
				address:          0,
			},
			want:    "sbc\tw20, w0, wzr",
			wantErr: false,
		},
		{
			name: "sbc	x29, x27, x25",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7d, 0x03, 0x19, 0xda}),
				address:          0,
			},
			want:    "sbc\tfp, x27, x25",
			wantErr: false,
		},
		{
			name: "sbc	xzr, x3, x4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x00, 0x04, 0xda}),
				address:          0,
			},
			want:    "sbc\txzr, x3, x4",
			wantErr: false,
		},
		{
			name: "ngc	x9, x10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x03, 0x0a, 0xda}),
				address:          0,
			},
			want:    "ngc\tx9, x10",
			wantErr: false,
		},
		{
			name: "sbc	x20, x0, xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x14, 0x00, 0x1f, 0xda}),
				address:          0,
			},
			want:    "sbc\tx20, x0, xzr",
			wantErr: false,
		},
		{
			name: "sbcs	w29, w27, w25",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7d, 0x03, 0x19, 0x7a}),
				address:          0,
			},
			want:    "sbcs\tw29, w27, w25",
			wantErr: false,
		},
		{
			name: "sbcs	wzr, w3, w4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x00, 0x04, 0x7a}),
				address:          0,
			},
			want:    "sbcs\twzr, w3, w4",
			wantErr: false,
		},
		{
			name: "ngcs	w9, w10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x03, 0x0a, 0x7a}),
				address:          0,
			},
			want:    "ngcs\tw9, w10",
			wantErr: false,
		},
		{
			name: "sbcs	w20, w0, wzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x14, 0x00, 0x1f, 0x7a}),
				address:          0,
			},
			want:    "sbcs\tw20, w0, wzr",
			wantErr: false,
		},
		{
			name: "sbcs	x29, x27, x25",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7d, 0x03, 0x19, 0xfa}),
				address:          0,
			},
			want:    "sbcs\tfp, x27, x25",
			wantErr: false,
		},
		{
			name: "sbcs	xzr, x3, x4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x00, 0x04, 0xfa}),
				address:          0,
			},
			want:    "sbcs\txzr, x3, x4",
			wantErr: false,
		},
		{
			name: "ngcs	x9, x10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x03, 0x0a, 0xfa}),
				address:          0,
			},
			want:    "ngcs\tx9, x10",
			wantErr: false,
		},
		{
			name: "sbcs	x20, x0, xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x14, 0x00, 0x1f, 0xfa}),
				address:          0,
			},
			want:    "sbcs\tx20, x0, xzr",
			wantErr: false,
		},
		{
			name: "ngc	w3, w12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0x0c, 0x5a}),
				address:          0,
			},
			want:    "ngc\tw3, w12",
			wantErr: false,
		},
		{
			name: "ngc	wzr, w9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x09, 0x5a}),
				address:          0,
			},
			want:    "ngc\twzr, w9",
			wantErr: false,
		},
		{
			name: "ngc	w23, wzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf7, 0x03, 0x1f, 0x5a}),
				address:          0,
			},
			want:    "ngc\tw23, wzr",
			wantErr: false,
		},
		{
			name: "ngc	x29, x30",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfd, 0x03, 0x1e, 0xda}),
				address:          0,
			},
			want:    "ngc\tfp, lr",
			wantErr: false,
		},
		{
			name: "ngc	xzr, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x00, 0xda}),
				address:          0,
			},
			want:    "ngc\txzr, x0",
			wantErr: false,
		},
		{
			name: "ngc	x0, xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x03, 0x1f, 0xda}),
				address:          0,
			},
			want:    "ngc\tx0, xzr",
			wantErr: false,
		},
		{
			name: "ngcs	w3, w12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0x0c, 0x7a}),
				address:          0,
			},
			want:    "ngcs\tw3, w12",
			wantErr: false,
		},
		{
			name: "ngcs	wzr, w9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x09, 0x7a}),
				address:          0,
			},
			want:    "ngcs\twzr, w9",
			wantErr: false,
		},
		{
			name: "ngcs	w23, wzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf7, 0x03, 0x1f, 0x7a}),
				address:          0,
			},
			want:    "ngcs\tw23, wzr",
			wantErr: false,
		},
		{
			name: "ngcs	x29, x30",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfd, 0x03, 0x1e, 0xfa}),
				address:          0,
			},
			want:    "ngcs\tfp, lr",
			wantErr: false,
		},
		{
			name: "ngcs	xzr, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x00, 0xfa}),
				address:          0,
			},
			want:    "ngcs\txzr, x0",
			wantErr: false,
		},
		{
			name: "ngcs	x0, xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x03, 0x1f, 0xfa}),
				address:          0,
			},
			want:    "ngcs\tx0, xzr",
			wantErr: false,
		},
		{
			name: "sbfx	x1, x2, #3, #2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x10, 0x43, 0x93}),
				address:          0,
			},
			want:    "sbfx\tx1, x2, #0x3, #0x2",
			wantErr: false,
		},
		{
			name: "asr	x3, x4, #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x83, 0xfc, 0x7f, 0x93}),
				address:          0,
			},
			want:    "asr\tx3, x4, #0x3f",
			wantErr: false,
		},
		{
			name: "asr	wzr, wzr, #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x7f, 0x1f, 0x13}),
				address:          0,
			},
			want:    "asr\twzr, wzr, #0x1f",
			wantErr: false,
		},
		{
			name: "sbfx	w12, w9, #0, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x01, 0x00, 0x13}),
				address:          0,
			},
			want:    "sbfx\tw12, w9, #0, #0x1",
			wantErr: false,
		},
		{
			name: "ubfiz	x4, x5, #52, #11",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa4, 0x28, 0x4c, 0xd3}),
				address:          0,
			},
			want:    "ubfiz\tx4, x5, #0x34, #0xb",
			wantErr: false,
		},
		{
			name: "ubfx	xzr, x4, #0, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x00, 0x40, 0xd3}),
				address:          0,
			},
			want:    "ubfx\txzr, x4, #0, #0x1",
			wantErr: false,
		},
		{
			name: "ubfiz	x4, xzr, #1, #6",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe4, 0x17, 0x7f, 0xd3}),
				address:          0,
			},
			want:    "ubfiz\tx4, xzr, #0x1, #0x6",
			wantErr: false,
		},
		{
			name: "lsr	x5, x6, #12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc5, 0xfc, 0x4c, 0xd3}),
				address:          0,
			},
			want:    "lsr\tx5, x6, #0xc",
			wantErr: false,
		},
		{
			name: "bfi	x4, x5, #52, #11",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa4, 0x28, 0x4c, 0xb3}),
				address:          0,
			},
			want:    "bfi\tx4, x5, #0x34, #0xb",
			wantErr: false,
		},
		{
			name: "bfxil	xzr, x4, #0, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x00, 0x40, 0xb3}),
				address:          0,
			},
			want:    "bfxil\txzr, x4, #0, #0x1",
			wantErr: false,
		},
		{
			name: "bfc	x4, #1, #6",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe4, 0x17, 0x7f, 0xb3}),
				address:          0,
			},
			want:    "bfc\tx4, #0x1, #0x6",
			wantErr: false,
		},
		{
			name: "bfxil	x5, x6, #12, #52",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc5, 0xfc, 0x4c, 0xb3}),
				address:          0,
			},
			want:    "bfxil\tx5, x6, #0xc, #0x34",
			wantErr: false,
		},
		{
			name: "sxtb	w1, w2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x1c, 0x00, 0x13}),
				address:          0,
			},
			want:    "sxtb\tw1, w2",
			wantErr: false,
		},
		{
			name: "sxtb	xzr, w3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x1c, 0x40, 0x93}),
				address:          0,
			},
			want:    "sxtb\txzr, w3",
			wantErr: false,
		},
		{
			name: "sxth	w9, w10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x3d, 0x00, 0x13}),
				address:          0,
			},
			want:    "sxth\tw9, w10",
			wantErr: false,
		},
		{
			name: "sxth	x0, w1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x3c, 0x40, 0x93}),
				address:          0,
			},
			want:    "sxth\tx0, w1",
			wantErr: false,
		},
		{
			name: "sxtw	x3, w30",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc3, 0x7f, 0x40, 0x93}),
				address:          0,
			},
			want:    "sxtw\tx3, w30",
			wantErr: false,
		},
		{
			name: "uxtb	w1, w2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x1c, 0x00, 0x53}),
				address:          0,
			},
			want:    "uxtb\tw1, w2",
			wantErr: false,
		},
		{
			name: "uxtb	wzr, w3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x1c, 0x00, 0x53}),
				address:          0,
			},
			want:    "uxtb\twzr, w3",
			wantErr: false,
		},
		{
			name: "uxth	w9, w10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x3d, 0x00, 0x53}),
				address:          0,
			},
			want:    "uxth\tw9, w10",
			wantErr: false,
		},
		{
			name: "uxth	w0, w1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x3c, 0x00, 0x53}),
				address:          0,
			},
			want:    "uxth\tw0, w1",
			wantErr: false,
		},
		{
			name: "asr	w3, w2, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x43, 0x7c, 0x00, 0x13}),
				address:          0,
			},
			want:    "asr\tw3, w2, #0",
			wantErr: false,
		},
		{
			name: "asr	w9, w10, #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x7d, 0x1f, 0x13}),
				address:          0,
			},
			want:    "asr\tw9, w10, #0x1f",
			wantErr: false,
		},
		{
			name: "asr	x20, x21, #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb4, 0xfe, 0x7f, 0x93}),
				address:          0,
			},
			want:    "asr\tx20, x21, #0x3f",
			wantErr: false,
		},
		{
			name: "asr	w1, wzr, #3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe1, 0x7f, 0x03, 0x13}),
				address:          0,
			},
			want:    "asr\tw1, wzr, #0x3",
			wantErr: false,
		},
		{
			name: "lsr	w3, w2, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x43, 0x7c, 0x00, 0x53}),
				address:          0,
			},
			want:    "lsr\tw3, w2, #0",
			wantErr: false,
		},
		{
			name: "lsr	w9, w10, #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x7d, 0x1f, 0x53}),
				address:          0,
			},
			want:    "lsr\tw9, w10, #0x1f",
			wantErr: false,
		},
		{
			name: "lsr	x20, x21, #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb4, 0xfe, 0x7f, 0xd3}),
				address:          0,
			},
			want:    "lsr\tx20, x21, #0x3f",
			wantErr: false,
		},
		{
			name: "lsr	wzr, wzr, #3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x7f, 0x03, 0x53}),
				address:          0,
			},
			want:    "lsr\twzr, wzr, #0x3",
			wantErr: false,
		},
		{
			name: "lsr	w3, w2, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x43, 0x7c, 0x00, 0x53}),
				address:          0,
			},
			want:    "lsr\tw3, w2, #0",
			wantErr: false,
		},
		{
			name: "lsl	w9, w10, #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x01, 0x01, 0x53}),
				address:          0,
			},
			want:    "lsl\tw9, w10, #0x1f",
			wantErr: false,
		},
		{
			name: "lsl	x20, x21, #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb4, 0x02, 0x41, 0xd3}),
				address:          0,
			},
			want:    "lsl\tx20, x21, #0x3f",
			wantErr: false,
		},
		{
			name: "lsl	w1, wzr, #3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe1, 0x73, 0x1d, 0x53}),
				address:          0,
			},
			want:    "lsl\tw1, wzr, #0x3",
			wantErr: false,
		},
		{
			name: "sbfx	w9, w10, #0, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x01, 0x00, 0x13}),
				address:          0,
			},
			want:    "sbfx\tw9, w10, #0, #0x1",
			wantErr: false,
		},
		{
			name: "sbfiz	x2, x3, #63, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x00, 0x41, 0x93}),
				address:          0,
			},
			want:    "sbfiz\tx2, x3, #0x3f, #0x1",
			wantErr: false,
		},
		{
			name: "asr	x19, x20, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0xfe, 0x40, 0x93}),
				address:          0,
			},
			want:    "asr\tx19, x20, #0",
			wantErr: false,
		},
		{
			name: "sbfiz	x9, x10, #5, #59",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xe9, 0x7b, 0x93}),
				address:          0,
			},
			want:    "sbfiz\tx9, x10, #0x5, #0x3b",
			wantErr: false,
		},
		{
			name: "asr	w9, w10, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x7d, 0x00, 0x13}),
				address:          0,
			},
			want:    "asr\tw9, w10, #0",
			wantErr: false,
		},
		{
			name: "sbfiz	w11, w12, #31, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8b, 0x01, 0x01, 0x13}),
				address:          0,
			},
			want:    "sbfiz\tw11, w12, #0x1f, #0x1",
			wantErr: false,
		},
		{
			name: "sbfiz	w13, w14, #29, #3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcd, 0x09, 0x03, 0x13}),
				address:          0,
			},
			want:    "sbfiz\tw13, w14, #0x1d, #0x3",
			wantErr: false,
		},
		{
			name: "sbfiz	xzr, xzr, #10, #11",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x2b, 0x76, 0x93}),
				address:          0,
			},
			want:    "sbfiz\txzr, xzr, #0xa, #0xb",
			wantErr: false,
		},
		{
			name: "sbfx	w9, w10, #0, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x01, 0x00, 0x13}),
				address:          0,
			},
			want:    "sbfx\tw9, w10, #0, #0x1",
			wantErr: false,
		},
		{
			name: "asr	x2, x3, #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xfc, 0x7f, 0x93}),
				address:          0,
			},
			want:    "asr\tx2, x3, #0x3f",
			wantErr: false,
		},
		{
			name: "asr	x19, x20, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0xfe, 0x40, 0x93}),
				address:          0,
			},
			want:    "asr\tx19, x20, #0",
			wantErr: false,
		},
		{
			name: "asr	x9, x10, #5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xfd, 0x45, 0x93}),
				address:          0,
			},
			want:    "asr\tx9, x10, #0x5",
			wantErr: false,
		},
		{
			name: "asr	w9, w10, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x7d, 0x00, 0x13}),
				address:          0,
			},
			want:    "asr\tw9, w10, #0",
			wantErr: false,
		},
		{
			name: "asr	w11, w12, #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8b, 0x7d, 0x1f, 0x13}),
				address:          0,
			},
			want:    "asr\tw11, w12, #0x1f",
			wantErr: false,
		},
		{
			name: "asr	w13, w14, #29",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcd, 0x7d, 0x1d, 0x13}),
				address:          0,
			},
			want:    "asr\tw13, w14, #0x1d",
			wantErr: false,
		},
		{
			name: "sbfx	xzr, xzr, #10, #11",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x53, 0x4a, 0x93}),
				address:          0,
			},
			want:    "sbfx\txzr, xzr, #0xa, #0xb",
			wantErr: false,
		},
		{
			name: "bfxil	w9, w10, #0, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x01, 0x00, 0x33}),
				address:          0,
			},
			want:    "bfxil\tw9, w10, #0, #0x1",
			wantErr: false,
		},
		{
			name: "bfi	x2, x3, #63, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x00, 0x41, 0xb3}),
				address:          0,
			},
			want:    "bfi\tx2, x3, #0x3f, #0x1",
			wantErr: false,
		},
		{
			name: "bfxil	x19, x20, #0, #64",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0xfe, 0x40, 0xb3}),
				address:          0,
			},
			want:    "bfxil\tx19, x20, #0, #0x40",
			wantErr: false,
		},
		{
			name: "bfi	x9, x10, #5, #59",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xe9, 0x7b, 0xb3}),
				address:          0,
			},
			want:    "bfi\tx9, x10, #0x5, #0x3b",
			wantErr: false,
		},
		{
			name: "bfxil	w9, w10, #0, #32",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x7d, 0x00, 0x33}),
				address:          0,
			},
			want:    "bfxil\tw9, w10, #0, #0x20",
			wantErr: false,
		},
		{
			name: "bfi	w11, w12, #31, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8b, 0x01, 0x01, 0x33}),
				address:          0,
			},
			want:    "bfi\tw11, w12, #0x1f, #0x1",
			wantErr: false,
		},
		{
			name: "bfi	w13, w14, #29, #3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcd, 0x09, 0x03, 0x33}),
				address:          0,
			},
			want:    "bfi\tw13, w14, #0x1d, #0x3",
			wantErr: false,
		},
		{
			name: "bfc	xzr, #10, #11",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x2b, 0x76, 0xb3}),
				address:          0,
			},
			want:    "bfc\txzr, #0xa, #0xb",
			wantErr: false,
		},
		{
			name: "bfxil	w9, w10, #0, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x01, 0x00, 0x33}),
				address:          0,
			},
			want:    "bfxil\tw9, w10, #0, #0x1",
			wantErr: false,
		},
		{
			name: "bfxil	x2, x3, #63, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xfc, 0x7f, 0xb3}),
				address:          0,
			},
			want:    "bfxil\tx2, x3, #0x3f, #0x1",
			wantErr: false,
		},
		{
			name: "bfxil	x19, x20, #0, #64",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0xfe, 0x40, 0xb3}),
				address:          0,
			},
			want:    "bfxil\tx19, x20, #0, #0x40",
			wantErr: false,
		},
		{
			name: "bfxil	x9, x10, #5, #59",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xfd, 0x45, 0xb3}),
				address:          0,
			},
			want:    "bfxil\tx9, x10, #0x5, #0x3b",
			wantErr: false,
		},
		{
			name: "bfxil	w9, w10, #0, #32",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x7d, 0x00, 0x33}),
				address:          0,
			},
			want:    "bfxil\tw9, w10, #0, #0x20",
			wantErr: false,
		},
		{
			name: "bfxil	w11, w12, #31, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8b, 0x7d, 0x1f, 0x33}),
				address:          0,
			},
			want:    "bfxil\tw11, w12, #0x1f, #0x1",
			wantErr: false,
		},
		{
			name: "bfxil	w13, w14, #29, #3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcd, 0x7d, 0x1d, 0x33}),
				address:          0,
			},
			want:    "bfxil\tw13, w14, #0x1d, #0x3",
			wantErr: false,
		},
		{
			name: "bfxil	xzr, xzr, #10, #11",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x53, 0x4a, 0xb3}),
				address:          0,
			},
			want:    "bfxil\txzr, xzr, #0xa, #0xb",
			wantErr: false,
		},
		{
			name: "ubfx	w9, w10, #0, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x01, 0x00, 0x53}),
				address:          0,
			},
			want:    "ubfx\tw9, w10, #0, #0x1",
			wantErr: false,
		},
		{
			name: "lsl	x2, x3, #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x00, 0x41, 0xd3}),
				address:          0,
			},
			want:    "lsl\tx2, x3, #0x3f",
			wantErr: false,
		},
		{
			name: "lsr	x19, x20, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0xfe, 0x40, 0xd3}),
				address:          0,
			},
			want:    "lsr\tx19, x20, #0",
			wantErr: false,
		},
		{
			name: "lsl	x9, x10, #5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xe9, 0x7b, 0xd3}),
				address:          0,
			},
			want:    "lsl\tx9, x10, #0x5",
			wantErr: false,
		},
		{
			name: "lsr	w9, w10, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x7d, 0x00, 0x53}),
				address:          0,
			},
			want:    "lsr\tw9, w10, #0",
			wantErr: false,
		},
		{
			name: "lsl	w11, w12, #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8b, 0x01, 0x01, 0x53}),
				address:          0,
			},
			want:    "lsl\tw11, w12, #0x1f",
			wantErr: false,
		},
		{
			name: "lsl	w13, w14, #29",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcd, 0x09, 0x03, 0x53}),
				address:          0,
			},
			want:    "lsl\tw13, w14, #0x1d",
			wantErr: false,
		},
		{
			name: "ubfiz	xzr, xzr, #10, #11",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x2b, 0x76, 0xd3}),
				address:          0,
			},
			want:    "ubfiz\txzr, xzr, #0xa, #0xb",
			wantErr: false,
		},
		{
			name: "ubfx	w9, w10, #0, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x01, 0x00, 0x53}),
				address:          0,
			},
			want:    "ubfx\tw9, w10, #0, #0x1",
			wantErr: false,
		},
		{
			name: "lsr	x2, x3, #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xfc, 0x7f, 0xd3}),
				address:          0,
			},
			want:    "lsr\tx2, x3, #0x3f",
			wantErr: false,
		},
		{
			name: "lsr	x19, x20, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0xfe, 0x40, 0xd3}),
				address:          0,
			},
			want:    "lsr\tx19, x20, #0",
			wantErr: false,
		},
		{
			name: "lsr	x9, x10, #5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xfd, 0x45, 0xd3}),
				address:          0,
			},
			want:    "lsr\tx9, x10, #0x5",
			wantErr: false,
		},
		{
			name: "lsr	w9, w10, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x7d, 0x00, 0x53}),
				address:          0,
			},
			want:    "lsr\tw9, w10, #0",
			wantErr: false,
		},
		{
			name: "lsr	w11, w12, #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8b, 0x7d, 0x1f, 0x53}),
				address:          0,
			},
			want:    "lsr\tw11, w12, #0x1f",
			wantErr: false,
		},
		{
			name: "lsr	w13, w14, #29",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcd, 0x7d, 0x1d, 0x53}),
				address:          0,
			},
			want:    "lsr\tw13, w14, #0x1d",
			wantErr: false,
		},
		{
			name: "ubfx	xzr, xzr, #10, #11",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x53, 0x4a, 0xd3}),
				address:          0,
			},
			want:    "ubfx\txzr, xzr, #0xa, #0xb",
			wantErr: false,
		},
		{
			name: "bfc	w3, #0, #32",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x7f, 0x00, 0x33}),
				address:          0,
			},
			want:    "bfxil\tw3, wzr, #0, #0x20",
			wantErr: false,
		},
		{
			name: "bfc	wzr, #31, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x01, 0x33}),
				address:          0,
			},
			want:    "bfc\twzr, #0x1f, #0x1",
			wantErr: false,
		},
		{
			name: "bfc	x0, #5, #9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x23, 0x7b, 0xb3}),
				address:          0,
			},
			want:    "bfc\tx0, #0x5, #0x9",
			wantErr: false,
		},
		{
			name: "bfc	xzr, #63, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x41, 0xb3}),
				address:          0,
			},
			want:    "bfc\txzr, #0x3f, #0x1",
			wantErr: false,
		},
		{
			name: "cbz	w5, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x05, 0x00, 0x00, 0x34}),
				address:          0,
			},
			want:    "cbz\tw5, 0x0",
			wantErr: false,
		},
		{
			name: "cbnz	x3, #-4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0xff, 0xff, 0xb5}),
				address:          0,
			},
			want:    "cbnz\tx3, 0xfffffffffffffffc",
			wantErr: false,
		},
		{
			name: "cbz	w20, #1048572",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0xff, 0x7f, 0x34}),
				address:          0,
			},
			want:    "cbz\tw20, 0xffffc",
			wantErr: false,
		},
		{
			name: "cbnz	xzr, #-1048576",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x00, 0x80, 0xb5}),
				address:          0,
			},
			want:    "cbnz\txzr, 0xfffffffffff00000",
			wantErr: false,
		},
		{
			name: "b.eq	#0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x00, 0x00, 0x54}),
				address:          0,
			},
			want:    "b.eq\t0x0",
			wantErr: false,
		},
		{
			name: "b.lt	#-4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xeb, 0xff, 0xff, 0x54}),
				address:          0,
			},
			want:    "b.lt\t0xfffffffffffffffc",
			wantErr: false,
		},
		{
			name: "b.lo	#1048572",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0xff, 0x7f, 0x54}),
				address:          0,
			},
			want:    "b.lo\t0xffffc",
			wantErr: false,
		},
		{
			name: "ccmp	w1, #31, #0, eq",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x08, 0x5f, 0x7a}),
				address:          0,
			},
			want:    "ccmp\tw1, #0x1f, #0, eq",
			wantErr: false,
		},
		{
			name: "ccmp	w3, #0, #15, hs",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6f, 0x28, 0x40, 0x7a}),
				address:          0,
			},
			want:    "ccmp\tw3, #0, #0xf, cs",
			wantErr: false,
		},
		{
			name: "ccmp	wzr, #15, #13, hs",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xed, 0x2b, 0x4f, 0x7a}),
				address:          0,
			},
			want:    "ccmp\twzr, #0xf, #0xd, cs",
			wantErr: false,
		},
		{
			name: "ccmp	x9, #31, #0, le",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xd9, 0x5f, 0xfa}),
				address:          0,
			},
			want:    "ccmp\tx9, #0x1f, #0, le",
			wantErr: false,
		},
		{
			name: "ccmp	x3, #0, #15, gt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6f, 0xc8, 0x40, 0xfa}),
				address:          0,
			},
			want:    "ccmp\tx3, #0, #0xf, gt",
			wantErr: false,
		},
		{
			name: "ccmp	xzr, #5, #7, ne",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe7, 0x1b, 0x45, 0xfa}),
				address:          0,
			},
			want:    "ccmp\txzr, #0x5, #0x7, ne",
			wantErr: false,
		},
		{
			name: "ccmn	w1, #31, #0, eq",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x08, 0x5f, 0x3a}),
				address:          0,
			},
			want:    "ccmn\tw1, #0x1f, #0, eq",
			wantErr: false,
		},
		{
			name: "ccmn	w3, #0, #15, hs",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6f, 0x28, 0x40, 0x3a}),
				address:          0,
			},
			want:    "ccmn\tw3, #0, #0xf, cs",
			wantErr: false,
		},
		{
			name: "ccmn	wzr, #15, #13, hs",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xed, 0x2b, 0x4f, 0x3a}),
				address:          0,
			},
			want:    "ccmn\twzr, #0xf, #0xd, cs",
			wantErr: false,
		},
		{
			name: "ccmn	x9, #31, #0, le",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xd9, 0x5f, 0xba}),
				address:          0,
			},
			want:    "ccmn\tx9, #0x1f, #0, le",
			wantErr: false,
		},
		{
			name: "ccmn	x3, #0, #15, gt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6f, 0xc8, 0x40, 0xba}),
				address:          0,
			},
			want:    "ccmn\tx3, #0, #0xf, gt",
			wantErr: false,
		},
		{
			name: "ccmn	xzr, #5, #7, ne",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe7, 0x1b, 0x45, 0xba}),
				address:          0,
			},
			want:    "ccmn\txzr, #0x5, #0x7, ne",
			wantErr: false,
		},
		{
			name: "ccmp	w1, wzr, #0, eq",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x00, 0x5f, 0x7a}),
				address:          0,
			},
			want:    "ccmp\tw1, wzr, #0, eq",
			wantErr: false,
		},
		{
			name: "ccmp	w3, w0, #15, hs",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6f, 0x20, 0x40, 0x7a}),
				address:          0,
			},
			want:    "ccmp\tw3, w0, #0xf, cs",
			wantErr: false,
		},
		{
			name: "ccmp	wzr, w15, #13, hs",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xed, 0x23, 0x4f, 0x7a}),
				address:          0,
			},
			want:    "ccmp\twzr, w15, #0xd, cs",
			wantErr: false,
		},
		{
			name: "ccmp	x9, xzr, #0, le",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xd1, 0x5f, 0xfa}),
				address:          0,
			},
			want:    "ccmp\tx9, xzr, #0, le",
			wantErr: false,
		},
		{
			name: "ccmp	x3, x0, #15, gt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6f, 0xc0, 0x40, 0xfa}),
				address:          0,
			},
			want:    "ccmp\tx3, x0, #0xf, gt",
			wantErr: false,
		},
		{
			name: "ccmp	xzr, x5, #7, ne",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe7, 0x13, 0x45, 0xfa}),
				address:          0,
			},
			want:    "ccmp\txzr, x5, #0x7, ne",
			wantErr: false,
		},
		{
			name: "ccmn	w1, wzr, #0, eq",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x00, 0x5f, 0x3a}),
				address:          0,
			},
			want:    "ccmn\tw1, wzr, #0, eq",
			wantErr: false,
		},
		{
			name: "ccmn	w3, w0, #15, hs",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6f, 0x20, 0x40, 0x3a}),
				address:          0,
			},
			want:    "ccmn\tw3, w0, #0xf, cs",
			wantErr: false,
		},
		{
			name: "ccmn	wzr, w15, #13, hs",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xed, 0x23, 0x4f, 0x3a}),
				address:          0,
			},
			want:    "ccmn\twzr, w15, #0xd, cs",
			wantErr: false,
		},
		{
			name: "ccmn	x9, xzr, #0, le",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xd1, 0x5f, 0xba}),
				address:          0,
			},
			want:    "ccmn\tx9, xzr, #0, le",
			wantErr: false,
		},
		{
			name: "ccmn	x3, x0, #15, gt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6f, 0xc0, 0x40, 0xba}),
				address:          0,
			},
			want:    "ccmn\tx3, x0, #0xf, gt",
			wantErr: false,
		},
		{
			name: "ccmn	xzr, x5, #7, ne",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe7, 0x13, 0x45, 0xba}),
				address:          0,
			},
			want:    "ccmn\txzr, x5, #0x7, ne",
			wantErr: false,
		},
		{
			name: "csel	w1, w0, w19, ne",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x01, 0x10, 0x93, 0x1a}),
				address:          0,
			},
			want:    "csel\tw1, w0, w19, ne",
			wantErr: false,
		},
		{
			name: "csel	wzr, w5, w9, eq",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x00, 0x89, 0x1a}),
				address:          0,
			},
			want:    "csel\twzr, w5, w9, eq",
			wantErr: false,
		},
		{
			name: "csel	w9, wzr, w30, gt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xc3, 0x9e, 0x1a}),
				address:          0,
			},
			want:    "csel\tw9, wzr, w30, gt",
			wantErr: false,
		},
		{
			name: "csel	w1, w28, wzr, mi",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x81, 0x43, 0x9f, 0x1a}),
				address:          0,
			},
			want:    "csel\tw1, w28, wzr, mi",
			wantErr: false,
		},
		{
			name: "csel	x19, x23, x29, lt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf3, 0xb2, 0x9d, 0x9a}),
				address:          0,
			},
			want:    "csel\tx19, x23, fp, lt",
			wantErr: false,
		},
		{
			name: "csel	xzr, x3, x4, ge",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0xa0, 0x84, 0x9a}),
				address:          0,
			},
			want:    "csel\txzr, x3, x4, ge",
			wantErr: false,
		},
		{
			name: "csel	x5, xzr, x6, hs",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe5, 0x23, 0x86, 0x9a}),
				address:          0,
			},
			want:    "csel\tx5, xzr, x6, cs",
			wantErr: false,
		},
		{
			name: "csel	x7, x8, xzr, lo",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x07, 0x31, 0x9f, 0x9a}),
				address:          0,
			},
			want:    "csel\tx7, x8, xzr, cc",
			wantErr: false,
		},
		{
			name: "csinc	w1, w0, w19, ne",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x01, 0x14, 0x93, 0x1a}),
				address:          0,
			},
			want:    "csinc\tw1, w0, w19, ne",
			wantErr: false,
		},
		{
			name: "csinc	wzr, w5, w9, eq",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x04, 0x89, 0x1a}),
				address:          0,
			},
			want:    "csinc\twzr, w5, w9, eq",
			wantErr: false,
		},
		{
			name: "csinc	w9, wzr, w30, gt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xc7, 0x9e, 0x1a}),
				address:          0,
			},
			want:    "csinc\tw9, wzr, w30, gt",
			wantErr: false,
		},
		{
			name: "csinc	w1, w28, wzr, mi",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x81, 0x47, 0x9f, 0x1a}),
				address:          0,
			},
			want:    "csinc\tw1, w28, wzr, mi",
			wantErr: false,
		},
		{
			name: "csinc	x19, x23, x29, lt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf3, 0xb6, 0x9d, 0x9a}),
				address:          0,
			},
			want:    "csinc\tx19, x23, fp, lt",
			wantErr: false,
		},
		{
			name: "csinc	xzr, x3, x4, ge",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0xa4, 0x84, 0x9a}),
				address:          0,
			},
			want:    "csinc\txzr, x3, x4, ge",
			wantErr: false,
		},
		{
			name: "csinc	x5, xzr, x6, hs",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe5, 0x27, 0x86, 0x9a}),
				address:          0,
			},
			want:    "csinc\tx5, xzr, x6, cs",
			wantErr: false,
		},
		{
			name: "csinc	x7, x8, xzr, lo",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x07, 0x35, 0x9f, 0x9a}),
				address:          0,
			},
			want:    "csinc\tx7, x8, xzr, cc",
			wantErr: false,
		},
		{
			name: "csinv	w1, w0, w19, ne",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x01, 0x10, 0x93, 0x5a}),
				address:          0,
			},
			want:    "csinv\tw1, w0, w19, ne",
			wantErr: false,
		},
		{
			name: "csinv	wzr, w5, w9, eq",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x00, 0x89, 0x5a}),
				address:          0,
			},
			want:    "csinv\twzr, w5, w9, eq",
			wantErr: false,
		},
		{
			name: "csinv	w9, wzr, w30, gt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xc3, 0x9e, 0x5a}),
				address:          0,
			},
			want:    "csinv\tw9, wzr, w30, gt",
			wantErr: false,
		},
		{
			name: "csinv	w1, w28, wzr, mi",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x81, 0x43, 0x9f, 0x5a}),
				address:          0,
			},
			want:    "csinv\tw1, w28, wzr, mi",
			wantErr: false,
		},
		{
			name: "csinv	x19, x23, x29, lt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf3, 0xb2, 0x9d, 0xda}),
				address:          0,
			},
			want:    "csinv\tx19, x23, fp, lt",
			wantErr: false,
		},
		{
			name: "csinv	xzr, x3, x4, ge",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0xa0, 0x84, 0xda}),
				address:          0,
			},
			want:    "csinv\txzr, x3, x4, ge",
			wantErr: false,
		},
		{
			name: "csinv	x5, xzr, x6, hs",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe5, 0x23, 0x86, 0xda}),
				address:          0,
			},
			want:    "csinv\tx5, xzr, x6, cs",
			wantErr: false,
		},
		{
			name: "csinv	x7, x8, xzr, lo",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x07, 0x31, 0x9f, 0xda}),
				address:          0,
			},
			want:    "csinv\tx7, x8, xzr, cc",
			wantErr: false,
		},
		{
			name: "csneg	w1, w0, w19, ne",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x01, 0x14, 0x93, 0x5a}),
				address:          0,
			},
			want:    "csneg\tw1, w0, w19, ne",
			wantErr: false,
		},
		{
			name: "csneg	wzr, w5, w9, eq",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x04, 0x89, 0x5a}),
				address:          0,
			},
			want:    "csneg\twzr, w5, w9, eq",
			wantErr: false,
		},
		{
			name: "csneg	w9, wzr, w30, gt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xc7, 0x9e, 0x5a}),
				address:          0,
			},
			want:    "csneg\tw9, wzr, w30, gt",
			wantErr: false,
		},
		{
			name: "csneg	w1, w28, wzr, mi",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x81, 0x47, 0x9f, 0x5a}),
				address:          0,
			},
			want:    "csneg\tw1, w28, wzr, mi",
			wantErr: false,
		},
		{
			name: "csneg	x19, x23, x29, lt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf3, 0xb6, 0x9d, 0xda}),
				address:          0,
			},
			want:    "csneg\tx19, x23, fp, lt",
			wantErr: false,
		},
		{
			name: "csneg	xzr, x3, x4, ge",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0xa4, 0x84, 0xda}),
				address:          0,
			},
			want:    "csneg\txzr, x3, x4, ge",
			wantErr: false,
		},
		{
			name: "csneg	x5, xzr, x6, hs",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe5, 0x27, 0x86, 0xda}),
				address:          0,
			},
			want:    "csneg\tx5, xzr, x6, cs",
			wantErr: false,
		},
		{
			name: "csneg	x7, x8, xzr, lo",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x07, 0x35, 0x9f, 0xda}),
				address:          0,
			},
			want:    "csneg\tx7, x8, xzr, cc",
			wantErr: false,
		},
		{
			name: "cset	w3, eq",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x17, 0x9f, 0x1a}),
				address:          0,
			},
			want:    "cset\tw3, eq",
			wantErr: false,
		},
		{
			name: "cset	x9, pl",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x47, 0x9f, 0x9a}),
				address:          0,
			},
			want:    "cset\tx9, pl",
			wantErr: false,
		},
		{
			name: "csetm	w20, ne",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x03, 0x9f, 0x5a}),
				address:          0,
			},
			want:    "csetm\tw20, ne",
			wantErr: false,
		},
		{
			name: "csetm	x30, ge",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfe, 0xb3, 0x9f, 0xda}),
				address:          0,
			},
			want:    "csetm\tlr, ge",
			wantErr: false,
		},
		{
			name: "cinc	w3, w5, gt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xd4, 0x85, 0x1a}),
				address:          0,
			},
			want:    "cinc\tw3, w5, gt",
			wantErr: false,
		},
		{
			name: "cinc	wzr, w4, le",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0xc4, 0x84, 0x1a}),
				address:          0,
			},
			want:    "cinc\twzr, w4, le",
			wantErr: false,
		},
		{
			name: "cset	w9, lt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xa7, 0x9f, 0x1a}),
				address:          0,
			},
			want:    "cset\tw9, lt",
			wantErr: false,
		},
		{
			name: "cinc	x3, x5, gt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xd4, 0x85, 0x9a}),
				address:          0,
			},
			want:    "cinc\tx3, x5, gt",
			wantErr: false,
		},
		{
			name: "cinc	xzr, x4, le",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0xc4, 0x84, 0x9a}),
				address:          0,
			},
			want:    "cinc\txzr, x4, le",
			wantErr: false,
		},
		{
			name: "cset	x9, lt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xa7, 0x9f, 0x9a}),
				address:          0,
			},
			want:    "cset\tx9, lt",
			wantErr: false,
		},
		{
			name: "cinv	w3, w5, gt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xd0, 0x85, 0x5a}),
				address:          0,
			},
			want:    "cinv\tw3, w5, gt",
			wantErr: false,
		},
		{
			name: "cinv	wzr, w4, le",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0xc0, 0x84, 0x5a}),
				address:          0,
			},
			want:    "cinv\twzr, w4, le",
			wantErr: false,
		},
		{
			name: "csetm	w9, lt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xa3, 0x9f, 0x5a}),
				address:          0,
			},
			want:    "csetm\tw9, lt",
			wantErr: false,
		},
		{
			name: "cinv	x3, x5, gt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xd0, 0x85, 0xda}),
				address:          0,
			},
			want:    "cinv\tx3, x5, gt",
			wantErr: false,
		},
		{
			name: "cinv	xzr, x4, le",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0xc0, 0x84, 0xda}),
				address:          0,
			},
			want:    "cinv\txzr, x4, le",
			wantErr: false,
		},
		{
			name: "csetm	x9, lt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xa3, 0x9f, 0xda}),
				address:          0,
			},
			want:    "csetm\tx9, lt",
			wantErr: false,
		},
		{
			name: "cneg	w3, w5, gt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xd4, 0x85, 0x5a}),
				address:          0,
			},
			want:    "cneg\tw3, w5, gt",
			wantErr: false,
		},
		{
			name: "cneg	wzr, w4, le",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0xc4, 0x84, 0x5a}),
				address:          0,
			},
			want:    "cneg\twzr, w4, le",
			wantErr: false,
		},
		{
			name: "cneg	w9, wzr, lt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xa7, 0x9f, 0x5a}),
				address:          0,
			},
			want:    "cneg\tw9, wzr, lt",
			wantErr: false,
		},
		{
			name: "cneg	x3, x5, gt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xd4, 0x85, 0xda}),
				address:          0,
			},
			want:    "cneg\tx3, x5, gt",
			wantErr: false,
		},
		{
			name: "cneg	xzr, x4, le",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0xc4, 0x84, 0xda}),
				address:          0,
			},
			want:    "cneg\txzr, x4, le",
			wantErr: false,
		},
		{
			name: "cneg	x9, xzr, lt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xa7, 0x9f, 0xda}),
				address:          0,
			},
			want:    "cneg\tx9, xzr, lt",
			wantErr: false,
		},
		{
			name: "rbit	w0, w7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x00, 0xc0, 0x5a}),
				address:          0,
			},
			want:    "rbit\tw0, w7",
			wantErr: false,
		},
		{
			name: "rbit	x18, x3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x72, 0x00, 0xc0, 0xda}),
				address:          0,
			},
			want:    "rbit\tx18, x3",
			wantErr: false,
		},
		{
			name: "rev16	w17, w1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x31, 0x04, 0xc0, 0x5a}),
				address:          0,
			},
			want:    "rev16\tw17, w1",
			wantErr: false,
		},
		{
			name: "rev16	x5, x2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x45, 0x04, 0xc0, 0xda}),
				address:          0,
			},
			want:    "rev16\tx5, x2",
			wantErr: false,
		},
		{
			name: "rev	w18, w0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x12, 0x08, 0xc0, 0x5a}),
				address:          0,
			},
			want:    "rev\tw18, w0",
			wantErr: false,
		},
		{
			name: "rev32	x20, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x34, 0x08, 0xc0, 0xda}),
				address:          0,
			},
			want:    "rev32\tx20, x1",
			wantErr: false,
		},
		{
			name: "rev32	x20, xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x0b, 0xc0, 0xda}),
				address:          0,
			},
			want:    "rev32\tx20, xzr",
			wantErr: false,
		},
		{
			name: "rev	x22, x2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x56, 0x0c, 0xc0, 0xda}),
				address:          0,
			},
			want:    "rev\tx22, x2",
			wantErr: false,
		},
		{
			name: "rev	x18, xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf2, 0x0f, 0xc0, 0xda}),
				address:          0,
			},
			want:    "rev\tx18, xzr",
			wantErr: false,
		},
		{
			name: "rev	w7, wzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe7, 0x0b, 0xc0, 0x5a}),
				address:          0,
			},
			want:    "rev\tw7, wzr",
			wantErr: false,
		},
		{
			name: "clz	w24, w3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x78, 0x10, 0xc0, 0x5a}),
				address:          0,
			},
			want:    "clz\tw24, w3",
			wantErr: false,
		},
		{
			name: "clz	x26, x4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9a, 0x10, 0xc0, 0xda}),
				address:          0,
			},
			want:    "clz\tx26, x4",
			wantErr: false,
		},
		{
			name: "cls	w3, w5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0x14, 0xc0, 0x5a}),
				address:          0,
			},
			want:    "cls\tw3, w5",
			wantErr: false,
		},
		{
			name: "cls	x20, x5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb4, 0x14, 0xc0, 0xda}),
				address:          0,
			},
			want:    "cls\tx20, x5",
			wantErr: false,
		},
		{
			name: "clz	w24, wzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf8, 0x13, 0xc0, 0x5a}),
				address:          0,
			},
			want:    "clz\tw24, wzr",
			wantErr: false,
		},
		{
			name: "rev	x22, xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf6, 0x0f, 0xc0, 0xda}),
				address:          0,
			},
			want:    "rev\tx22, xzr",
			wantErr: false,
		},
		{
			name: "rev	x13, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8d, 0x0d, 0xc0, 0xda}),
				address:          0,
			},
			want:    "rev\tx13, x12",
			wantErr: false,
		},
		{
			name: "udiv	w0, w7, w10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x08, 0xca, 0x1a}),
				address:          0,
			},
			want:    "udiv\tw0, w7, w10",
			wantErr: false,
		},
		{
			name: "udiv	x9, x22, x4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x0a, 0xc4, 0x9a}),
				address:          0,
			},
			want:    "udiv\tx9, x22, x4",
			wantErr: false,
		},
		{
			name: "sdiv	w12, w21, w0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x0e, 0xc0, 0x1a}),
				address:          0,
			},
			want:    "sdiv\tw12, w21, w0",
			wantErr: false,
		},
		{
			name: "sdiv	x13, x2, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4d, 0x0c, 0xc1, 0x9a}),
				address:          0,
			},
			want:    "sdiv\tx13, x2, x1",
			wantErr: false,
		},
		{
			name: "lsl	w11, w12, w13",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8b, 0x21, 0xcd, 0x1a}),
				address:          0,
			},
			want:    "lsl\tw11, w12, w13",
			wantErr: false,
		},
		{
			name: "lsl	x14, x15, x16",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xee, 0x21, 0xd0, 0x9a}),
				address:          0,
			},
			want:    "lsl\tx14, x15, x16",
			wantErr: false,
		},
		{
			name: "lsr	w17, w18, w19",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x51, 0x26, 0xd3, 0x1a}),
				address:          0,
			},
			want:    "lsr\tw17, w18, w19",
			wantErr: false,
		},
		{
			name: "lsr	x20, x21, x22",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb4, 0x26, 0xd6, 0x9a}),
				address:          0,
			},
			want:    "lsr\tx20, x21, x22",
			wantErr: false,
		},
		{
			name: "asr	w23, w24, w25",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x17, 0x2b, 0xd9, 0x1a}),
				address:          0,
			},
			want:    "asr\tw23, w24, w25",
			wantErr: false,
		},
		{
			name: "asr	x26, x27, x28",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7a, 0x2b, 0xdc, 0x9a}),
				address:          0,
			},
			want:    "asr\tx26, x27, x28",
			wantErr: false,
		},
		{
			name: "ror	w0, w1, w2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x2c, 0xc2, 0x1a}),
				address:          0,
			},
			want:    "ror\tw0, w1, w2",
			wantErr: false,
		},
		{
			name: "ror	x3, x4, x5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x83, 0x2c, 0xc5, 0x9a}),
				address:          0,
			},
			want:    "ror\tx3, x4, x5",
			wantErr: false,
		},
		{
			name: "lsl	w6, w7, w8",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe6, 0x20, 0xc8, 0x1a}),
				address:          0,
			},
			want:    "lsl\tw6, w7, w8",
			wantErr: false,
		},
		{
			name: "lsl	x9, x10, x11",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x21, 0xcb, 0x9a}),
				address:          0,
			},
			want:    "lsl\tx9, x10, x11",
			wantErr: false,
		},
		{
			name: "lsr	w12, w13, w14",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x25, 0xce, 0x1a}),
				address:          0,
			},
			want:    "lsr\tw12, w13, w14",
			wantErr: false,
		},
		{
			name: "lsr	x15, x16, x17",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0f, 0x26, 0xd1, 0x9a}),
				address:          0,
			},
			want:    "lsr\tx15, x16, x17",
			wantErr: false,
		},
		{
			name: "asr	w18, w19, w20",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x72, 0x2a, 0xd4, 0x1a}),
				address:          0,
			},
			want:    "asr\tw18, w19, w20",
			wantErr: false,
		},
		{
			name: "asr	x21, x22, x23",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd5, 0x2a, 0xd7, 0x9a}),
				address:          0,
			},
			want:    "asr\tx21, x22, x23",
			wantErr: false,
		},
		{
			name: "ror	w24, w25, w26",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x38, 0x2f, 0xda, 0x1a}),
				address:          0,
			},
			want:    "ror\tw24, w25, w26",
			wantErr: false,
		},
		{
			name: "ror	x27, x28, x29",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9b, 0x2f, 0xdd, 0x9a}),
				address:          0,
			},
			want:    "ror\tx27, x28, fp",
			wantErr: false,
		},
		{
			name: "madd	w1, w3, w7, w4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x61, 0x10, 0x07, 0x1b}),
				address:          0,
			},
			want:    "madd\tw1, w3, w7, w4",
			wantErr: false,
		},
		{
			name: "madd	wzr, w0, w9, w11",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x2c, 0x09, 0x1b}),
				address:          0,
			},
			want:    "madd\twzr, w0, w9, w11",
			wantErr: false,
		},
		{
			name: "madd	w13, wzr, w4, w4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xed, 0x13, 0x04, 0x1b}),
				address:          0,
			},
			want:    "madd\tw13, wzr, w4, w4",
			wantErr: false,
		},
		{
			name: "madd	w19, w30, wzr, w29",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd3, 0x77, 0x1f, 0x1b}),
				address:          0,
			},
			want:    "madd\tw19, w30, wzr, w29",
			wantErr: false,
		},
		{
			name: "mul	w4, w5, w6",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa4, 0x7c, 0x06, 0x1b}),
				address:          0,
			},
			want:    "mul\tw4, w5, w6",
			wantErr: false,
		},
		{
			name: "madd	x1, x3, x7, x4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x61, 0x10, 0x07, 0x9b}),
				address:          0,
			},
			want:    "madd\tx1, x3, x7, x4",
			wantErr: false,
		},
		{
			name: "madd	xzr, x0, x9, x11",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x2c, 0x09, 0x9b}),
				address:          0,
			},
			want:    "madd\txzr, x0, x9, x11",
			wantErr: false,
		},
		{
			name: "madd	x13, xzr, x4, x4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xed, 0x13, 0x04, 0x9b}),
				address:          0,
			},
			want:    "madd\tx13, xzr, x4, x4",
			wantErr: false,
		},
		{
			name: "madd	x19, x30, xzr, x29",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd3, 0x77, 0x1f, 0x9b}),
				address:          0,
			},
			want:    "madd\tx19, lr, xzr, fp",
			wantErr: false,
		},
		{
			name: "mul	x4, x5, x6",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa4, 0x7c, 0x06, 0x9b}),
				address:          0,
			},
			want:    "mul\tx4, x5, x6",
			wantErr: false,
		},
		{
			name: "msub	w1, w3, w7, w4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x61, 0x90, 0x07, 0x1b}),
				address:          0,
			},
			want:    "msub\tw1, w3, w7, w4",
			wantErr: false,
		},
		{
			name: "msub	wzr, w0, w9, w11",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0xac, 0x09, 0x1b}),
				address:          0,
			},
			want:    "msub\twzr, w0, w9, w11",
			wantErr: false,
		},
		{
			name: "msub	w13, wzr, w4, w4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xed, 0x93, 0x04, 0x1b}),
				address:          0,
			},
			want:    "msub\tw13, wzr, w4, w4",
			wantErr: false,
		},
		{
			name: "msub	w19, w30, wzr, w29",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd3, 0xf7, 0x1f, 0x1b}),
				address:          0,
			},
			want:    "msub\tw19, w30, wzr, w29",
			wantErr: false,
		},
		{
			name: "mneg	w4, w5, w6",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa4, 0xfc, 0x06, 0x1b}),
				address:          0,
			},
			want:    "mneg\tw4, w5, w6",
			wantErr: false,
		},
		{
			name: "msub	x1, x3, x7, x4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x61, 0x90, 0x07, 0x9b}),
				address:          0,
			},
			want:    "msub\tx1, x3, x7, x4",
			wantErr: false,
		},
		{
			name: "msub	xzr, x0, x9, x11",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0xac, 0x09, 0x9b}),
				address:          0,
			},
			want:    "msub\txzr, x0, x9, x11",
			wantErr: false,
		},
		{
			name: "msub	x13, xzr, x4, x4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xed, 0x93, 0x04, 0x9b}),
				address:          0,
			},
			want:    "msub\tx13, xzr, x4, x4",
			wantErr: false,
		},
		{
			name: "msub	x19, x30, xzr, x29",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd3, 0xf7, 0x1f, 0x9b}),
				address:          0,
			},
			want:    "msub\tx19, lr, xzr, fp",
			wantErr: false,
		},
		{
			name: "mneg	x4, x5, x6",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa4, 0xfc, 0x06, 0x9b}),
				address:          0,
			},
			want:    "mneg\tx4, x5, x6",
			wantErr: false,
		},
		{
			name: "smaddl	x3, w5, w2, x9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0x24, 0x22, 0x9b}),
				address:          0,
			},
			want:    "smaddl\tx3, w5, w2, x9",
			wantErr: false,
		},
		{
			name: "smaddl	xzr, w10, w11, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x31, 0x2b, 0x9b}),
				address:          0,
			},
			want:    "smaddl\txzr, w10, w11, x12",
			wantErr: false,
		},
		{
			name: "smaddl	x13, wzr, w14, x15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xed, 0x3f, 0x2e, 0x9b}),
				address:          0,
			},
			want:    "smaddl\tx13, wzr, w14, x15",
			wantErr: false,
		},
		{
			name: "smaddl	x16, w17, wzr, x18",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x30, 0x4a, 0x3f, 0x9b}),
				address:          0,
			},
			want:    "smaddl\tx16, w17, wzr, x18",
			wantErr: false,
		},
		{
			name: "smull	x19, w20, w21",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0x7e, 0x35, 0x9b}),
				address:          0,
			},
			want:    "smull\tx19, w20, w21",
			wantErr: false,
		},
		{
			name: "smsubl	x3, w5, w2, x9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xa4, 0x22, 0x9b}),
				address:          0,
			},
			want:    "smsubl\tx3, w5, w2, x9",
			wantErr: false,
		},
		{
			name: "smsubl	xzr, w10, w11, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0xb1, 0x2b, 0x9b}),
				address:          0,
			},
			want:    "smsubl\txzr, w10, w11, x12",
			wantErr: false,
		},
		{
			name: "smsubl	x13, wzr, w14, x15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xed, 0xbf, 0x2e, 0x9b}),
				address:          0,
			},
			want:    "smsubl\tx13, wzr, w14, x15",
			wantErr: false,
		},
		{
			name: "smsubl	x16, w17, wzr, x18",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x30, 0xca, 0x3f, 0x9b}),
				address:          0,
			},
			want:    "smsubl\tx16, w17, wzr, x18",
			wantErr: false,
		},
		{
			name: "smnegl	x19, w20, w21",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0xfe, 0x35, 0x9b}),
				address:          0,
			},
			want:    "smnegl\tx19, w20, w21",
			wantErr: false,
		},
		{
			name: "umaddl	x3, w5, w2, x9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0x24, 0xa2, 0x9b}),
				address:          0,
			},
			want:    "umaddl\tx3, w5, w2, x9",
			wantErr: false,
		},
		{
			name: "umaddl	xzr, w10, w11, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x31, 0xab, 0x9b}),
				address:          0,
			},
			want:    "umaddl\txzr, w10, w11, x12",
			wantErr: false,
		},
		{
			name: "umaddl	x13, wzr, w14, x15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xed, 0x3f, 0xae, 0x9b}),
				address:          0,
			},
			want:    "umaddl\tx13, wzr, w14, x15",
			wantErr: false,
		},
		{
			name: "umaddl	x16, w17, wzr, x18",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x30, 0x4a, 0xbf, 0x9b}),
				address:          0,
			},
			want:    "umaddl\tx16, w17, wzr, x18",
			wantErr: false,
		},
		{
			name: "umull	x19, w20, w21",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0x7e, 0xb5, 0x9b}),
				address:          0,
			},
			want:    "umull\tx19, w20, w21",
			wantErr: false,
		},
		{
			name: "umsubl	x3, w5, w2, x9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xa4, 0xa2, 0x9b}),
				address:          0,
			},
			want:    "umsubl\tx3, w5, w2, x9",
			wantErr: false,
		},
		{
			name: "umsubl	xzr, w10, w11, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0xb1, 0xab, 0x9b}),
				address:          0,
			},
			want:    "umsubl\txzr, w10, w11, x12",
			wantErr: false,
		},
		{
			name: "umsubl	x13, wzr, w14, x15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xed, 0xbf, 0xae, 0x9b}),
				address:          0,
			},
			want:    "umsubl\tx13, wzr, w14, x15",
			wantErr: false,
		},
		{
			name: "umsubl	x16, w17, wzr, x18",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x30, 0xca, 0xbf, 0x9b}),
				address:          0,
			},
			want:    "umsubl\tx16, w17, wzr, x18",
			wantErr: false,
		},
		{
			name: "umnegl	x19, w20, w21",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0xfe, 0xb5, 0x9b}),
				address:          0,
			},
			want:    "umnegl\tx19, w20, w21",
			wantErr: false,
		},
		{
			name: "smulh	x30, x29, x28",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbe, 0x7f, 0x5c, 0x9b}),
				address:          0,
			},
			want:    "smulh\tlr, fp, x28",
			wantErr: false,
		},
		{
			name: "smulh	xzr, x27, x26",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x7f, 0x5a, 0x9b}),
				address:          0,
			},
			want:    "smulh\txzr, x27, x26",
			wantErr: false,
		},
		{
			name: "smulh	x25, xzr, x24",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf9, 0x7f, 0x58, 0x9b}),
				address:          0,
			},
			want:    "smulh\tx25, xzr, x24",
			wantErr: false,
		},
		{
			name: "smulh	x23, x22, xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd7, 0x7e, 0x5f, 0x9b}),
				address:          0,
			},
			want:    "smulh\tx23, x22, xzr",
			wantErr: false,
		},
		{
			name: "umulh	x30, x29, x28",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbe, 0x7f, 0xdc, 0x9b}),
				address:          0,
			},
			want:    "umulh\tlr, fp, x28",
			wantErr: false,
		},
		{
			name: "umulh	xzr, x27, x26",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x7f, 0xda, 0x9b}),
				address:          0,
			},
			want:    "umulh\txzr, x27, x26",
			wantErr: false,
		},
		{
			name: "umulh	x25, xzr, x24",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf9, 0x7f, 0xd8, 0x9b}),
				address:          0,
			},
			want:    "umulh\tx25, xzr, x24",
			wantErr: false,
		},
		{
			name: "umulh	x23, x22, xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd7, 0x7e, 0xdf, 0x9b}),
				address:          0,
			},
			want:    "umulh\tx23, x22, xzr",
			wantErr: false,
		},
		{
			name: "mul	w3, w4, w5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x83, 0x7c, 0x05, 0x1b}),
				address:          0,
			},
			want:    "mul\tw3, w4, w5",
			wantErr: false,
		},
		{
			name: "mul	wzr, w6, w7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0x7c, 0x07, 0x1b}),
				address:          0,
			},
			want:    "mul\twzr, w6, w7",
			wantErr: false,
		},
		{
			name: "mul	w8, wzr, w9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe8, 0x7f, 0x09, 0x1b}),
				address:          0,
			},
			want:    "mul\tw8, wzr, w9",
			wantErr: false,
		},
		{
			name: "mul	w10, w11, wzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6a, 0x7d, 0x1f, 0x1b}),
				address:          0,
			},
			want:    "mul\tw10, w11, wzr",
			wantErr: false,
		},
		{
			name: "mul	x12, x13, x14",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x7d, 0x0e, 0x9b}),
				address:          0,
			},
			want:    "mul\tx12, x13, x14",
			wantErr: false,
		},
		{
			name: "mul	xzr, x15, x16",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x7d, 0x10, 0x9b}),
				address:          0,
			},
			want:    "mul\txzr, x15, x16",
			wantErr: false,
		},
		{
			name: "mul	x17, xzr, x18",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf1, 0x7f, 0x12, 0x9b}),
				address:          0,
			},
			want:    "mul\tx17, xzr, x18",
			wantErr: false,
		},
		{
			name: "mul	x19, x20, xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0x7e, 0x1f, 0x9b}),
				address:          0,
			},
			want:    "mul\tx19, x20, xzr",
			wantErr: false,
		},
		{
			name: "mneg	w21, w22, w23",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd5, 0xfe, 0x17, 0x1b}),
				address:          0,
			},
			want:    "mneg\tw21, w22, w23",
			wantErr: false,
		},
		{
			name: "mneg	wzr, w24, w25",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0xff, 0x19, 0x1b}),
				address:          0,
			},
			want:    "mneg\twzr, w24, w25",
			wantErr: false,
		},
		{
			name: "mneg	w26, wzr, w27",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfa, 0xff, 0x1b, 0x1b}),
				address:          0,
			},
			want:    "mneg\tw26, wzr, w27",
			wantErr: false,
		},
		{
			name: "mneg	w28, w29, wzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbc, 0xff, 0x1f, 0x1b}),
				address:          0,
			},
			want:    "mneg\tw28, w29, wzr",
			wantErr: false,
		},
		{
			name: "smull	x11, w13, w17",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xab, 0x7d, 0x31, 0x9b}),
				address:          0,
			},
			want:    "smull\tx11, w13, w17",
			wantErr: false,
		},
		{
			name: "umull	x11, w13, w17",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xab, 0x7d, 0xb1, 0x9b}),
				address:          0,
			},
			want:    "umull\tx11, w13, w17",
			wantErr: false,
		},
		{
			name: "smnegl	x11, w13, w17",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xab, 0xfd, 0x31, 0x9b}),
				address:          0,
			},
			want:    "smnegl\tx11, w13, w17",
			wantErr: false,
		},
		{
			name: "umnegl	x11, w13, w17",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xab, 0xfd, 0xb1, 0x9b}),
				address:          0,
			},
			want:    "umnegl\tx11, w13, w17",
			wantErr: false,
		},
		{
			name: "svc	#0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x01, 0x00, 0x00, 0xd4}),
				address:          0,
			},
			want:    "svc\t#0",
			wantErr: false,
		},
		{
			name: "svc	#0xffff",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe1, 0xff, 0x1f, 0xd4}),
				address:          0,
			},
			want:    "svc\t#0xffff",
			wantErr: false,
		},
		{
			name: "hvc	#0x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x22, 0x00, 0x00, 0xd4}),
				address:          0,
			},
			want:    "hvc\t#0x1",
			wantErr: false,
		},
		{
			name: "smc	#0x2ee0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x03, 0xdc, 0x05, 0xd4}),
				address:          0,
			},
			want:    "smc\t#0x2ee0",
			wantErr: false,
		},
		{
			name: "brk	#0xc",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x80, 0x01, 0x20, 0xd4}),
				address:          0,
			},
			want:    "brk\t#0xc",
			wantErr: false,
		},
		{
			name: "hlt	#0x7b",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x60, 0x0f, 0x40, 0xd4}),
				address:          0,
			},
			want:    "hlt\t#0x7b",
			wantErr: false,
		},
		{
			name: "dcps1	#0x2a",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x05, 0xa0, 0xd4}),
				address:          0,
			},
			want:    "dcps1\t#0x2a",
			wantErr: false,
		},
		{
			name: "dcps2	#0x9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x22, 0x01, 0xa0, 0xd4}),
				address:          0,
			},
			want:    "dcps2\t#0x9",
			wantErr: false,
		},
		{
			name: "dcps3	#0x3e8",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x03, 0x7d, 0xa0, 0xd4}),
				address:          0,
			},
			want:    "dcps3\t#0x3e8",
			wantErr: false,
		},
		{
			name: "dcps1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x01, 0x00, 0xa0, 0xd4}),
				address:          0,
			},
			want:    "dcps1",
			wantErr: false,
		},
		{
			name: "dcps2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x02, 0x00, 0xa0, 0xd4}),
				address:          0,
			},
			want:    "dcps2",
			wantErr: false,
		},
		{
			name: "dcps3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x03, 0x00, 0xa0, 0xd4}),
				address:          0,
			},
			want:    "dcps3",
			wantErr: false,
		},
		{
			name: "extr	w3, w5, w7, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0x00, 0x87, 0x13}),
				address:          0,
			},
			want:    "extr\tw3, w5, w7, #0",
			wantErr: false,
		},
		{
			name: "extr	w11, w13, w17, #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xab, 0x7d, 0x91, 0x13}),
				address:          0,
			},
			want:    "extr\tw11, w13, w17, #0x1f",
			wantErr: false,
		},
		{
			name: "extr	x3, x5, x7, #15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0x3c, 0xc7, 0x93}),
				address:          0,
			},
			want:    "extr\tx3, x5, x7, #0xf",
			wantErr: false,
		},
		{
			name: "extr	x11, x13, x17, #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xab, 0xfd, 0xd1, 0x93}),
				address:          0,
			},
			want:    "extr\tx11, x13, x17, #0x3f",
			wantErr: false,
		},
		{
			name: "ror	x19, x23, #24",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf3, 0x62, 0xd7, 0x93}),
				address:          0,
			},
			want:    "ror\tx19, x23, #0x18",
			wantErr: false,
		},
		{
			name: "ror	x29, xzr, #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfd, 0xff, 0xdf, 0x93}),
				address:          0,
			},
			want:    "ror\tfp, xzr, #0x3f",
			wantErr: false,
		},
		{
			name: "ror	w9, w13, #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x7d, 0x8d, 0x13}),
				address:          0,
			},
			want:    "ror\tw9, w13, #0x1f",
			wantErr: false,
		},
		{
			name: "fcmp	s3, s5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x60, 0x20, 0x25, 0x1e}),
				address:          0,
			},
			want:    "fcmp\ts3, s5",
			wantErr: false,
		},
		{
			name: "fcmp	s31, #0.0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe8, 0x23, 0x20, 0x1e}),
				address:          0,
			},
			want:    "fcmp\ts31, #0.00000000",
			wantErr: false,
		},
		{
			name: "fcmpe	s29, s30",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb0, 0x23, 0x3e, 0x1e}),
				address:          0,
			},
			want:    "fcmpe\ts29, s30",
			wantErr: false,
		},
		{
			name: "fcmpe	s15, #0.0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf8, 0x21, 0x20, 0x1e}),
				address:          0,
			},
			want:    "fcmpe\ts15, #0.00000000",
			wantErr: false,
		},
		{
			name: "fcmp	d4, d12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x80, 0x20, 0x6c, 0x1e}),
				address:          0,
			},
			want:    "fcmp\td4, d12",
			wantErr: false,
		},
		{
			name: "fcmp	d23, #0.0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe8, 0x22, 0x60, 0x1e}),
				address:          0,
			},
			want:    "fcmp\td23, #0.00000000",
			wantErr: false,
		},
		{
			name: "fcmpe	d26, d22",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x50, 0x23, 0x76, 0x1e}),
				address:          0,
			},
			want:    "fcmpe\td26, d22",
			wantErr: false,
		},
		{
			name: "fcmpe	d29, #0.0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb8, 0x23, 0x60, 0x1e}),
				address:          0,
			},
			want:    "fcmpe\td29, #0.00000000",
			wantErr: false,
		},
		{
			name: "fccmp	s1, s31, #0, eq",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x04, 0x3f, 0x1e}),
				address:          0,
			},
			want:    "fccmp\ts1, s31, #0, eq",
			wantErr: false,
		},
		{
			name: "fccmp	s3, s0, #15, hs",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6f, 0x24, 0x20, 0x1e}),
				address:          0,
			},
			want:    "fccmp\ts3, s0, #0xf, cs",
			wantErr: false,
		},
		{
			name: "fccmp	s31, s15, #13, hs",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xed, 0x27, 0x2f, 0x1e}),
				address:          0,
			},
			want:    "fccmp\ts31, s15, #0xd, cs",
			wantErr: false,
		},
		{
			name: "fccmp	d9, d31, #0, le",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0xd5, 0x7f, 0x1e}),
				address:          0,
			},
			want:    "fccmp\td9, d31, #0, le",
			wantErr: false,
		},
		{
			name: "fccmp	d3, d0, #15, gt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6f, 0xc4, 0x60, 0x1e}),
				address:          0,
			},
			want:    "fccmp\td3, d0, #0xf, gt",
			wantErr: false,
		},
		{
			name: "fccmp	d31, d5, #7, ne",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe7, 0x17, 0x65, 0x1e}),
				address:          0,
			},
			want:    "fccmp\td31, d5, #0x7, ne",
			wantErr: false,
		},
		{
			name: "fccmpe	s1, s31, #0, eq",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x30, 0x04, 0x3f, 0x1e}),
				address:          0,
			},
			want:    "fccmpe\ts1, s31, #0, eq",
			wantErr: false,
		},
		{
			name: "fccmpe	s3, s0, #15, hs",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x24, 0x20, 0x1e}),
				address:          0,
			},
			want:    "fccmpe\ts3, s0, #0xf, cs",
			wantErr: false,
		},
		{
			name: "fccmpe	s31, s15, #13, hs",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfd, 0x27, 0x2f, 0x1e}),
				address:          0,
			},
			want:    "fccmpe\ts31, s15, #0xd, cs",
			wantErr: false,
		},
		{
			name: "fccmpe	d9, d31, #0, le",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x30, 0xd5, 0x7f, 0x1e}),
				address:          0,
			},
			want:    "fccmpe\td9, d31, #0, le",
			wantErr: false,
		},
		{
			name: "fccmpe	d3, d0, #15, gt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0xc4, 0x60, 0x1e}),
				address:          0,
			},
			want:    "fccmpe\td3, d0, #0xf, gt",
			wantErr: false,
		},
		{
			name: "fccmpe	d31, d5, #7, ne",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf7, 0x17, 0x65, 0x1e}),
				address:          0,
			},
			want:    "fccmpe\td31, d5, #0x7, ne",
			wantErr: false,
		},
		{
			name: "fcsel	s3, s20, s9, pl",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x83, 0x5e, 0x29, 0x1e}),
				address:          0,
			},
			want:    "fcsel\ts3, s20, s9, pl",
			wantErr: false,
		},
		{
			name: "fcsel	d9, d10, d11, mi",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x4d, 0x6b, 0x1e}),
				address:          0,
			},
			want:    "fcsel\td9, d10, d11, mi",
			wantErr: false,
		},
		{
			name: "fmov	s0, s1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x40, 0x20, 0x1e}),
				address:          0,
			},
			want:    "fmov\ts0, s1",
			wantErr: false,
		},
		{
			name: "fabs	s2, s3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xc0, 0x20, 0x1e}),
				address:          0,
			},
			want:    "fabs\ts2, s3",
			wantErr: false,
		},
		{
			name: "fneg	s4, s5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa4, 0x40, 0x21, 0x1e}),
				address:          0,
			},
			want:    "fneg\ts4, s5",
			wantErr: false,
		},
		{
			name: "fsqrt	s6, s7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe6, 0xc0, 0x21, 0x1e}),
				address:          0,
			},
			want:    "fsqrt\ts6, s7",
			wantErr: false,
		},
		{
			name: "fcvt	d8, s9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x28, 0xc1, 0x22, 0x1e}),
				address:          0,
			},
			want:    "fcvt\td8, s9",
			wantErr: false,
		},
		{
			name: "fcvt	h10, s11",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6a, 0xc1, 0x23, 0x1e}),
				address:          0,
			},
			want:    "fcvt\th10, s11",
			wantErr: false,
		},
		{
			name: "frintn	s12, s13",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x41, 0x24, 0x1e}),
				address:          0,
			},
			want:    "frintn\ts12, s13",
			wantErr: false,
		},
		{
			name: "frintp	s14, s15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xee, 0xc1, 0x24, 0x1e}),
				address:          0,
			},
			want:    "frintp\ts14, s15",
			wantErr: false,
		},
		{
			name: "frintm	s16, s17",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x30, 0x42, 0x25, 0x1e}),
				address:          0,
			},
			want:    "frintm\ts16, s17",
			wantErr: false,
		},
		{
			name: "frintz	s18, s19",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x72, 0xc2, 0x25, 0x1e}),
				address:          0,
			},
			want:    "frintz\ts18, s19",
			wantErr: false,
		},
		{
			name: "frinta	s20, s21",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb4, 0x42, 0x26, 0x1e}),
				address:          0,
			},
			want:    "frinta\ts20, s21",
			wantErr: false,
		},
		{
			name: "frintx	s22, s23",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf6, 0x42, 0x27, 0x1e}),
				address:          0,
			},
			want:    "frintx\ts22, s23",
			wantErr: false,
		},
		{
			name: "frinti	s24, s25",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x38, 0xc3, 0x27, 0x1e}),
				address:          0,
			},
			want:    "frinti\ts24, s25",
			wantErr: false,
		},
		{
			name: "fmov	d0, d1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x40, 0x60, 0x1e}),
				address:          0,
			},
			want:    "fmov\td0, d1",
			wantErr: false,
		},
		{
			name: "fabs	d2, d3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0xc0, 0x60, 0x1e}),
				address:          0,
			},
			want:    "fabs\td2, d3",
			wantErr: false,
		},
		{
			name: "fneg	d4, d5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa4, 0x40, 0x61, 0x1e}),
				address:          0,
			},
			want:    "fneg\td4, d5",
			wantErr: false,
		},
		{
			name: "fsqrt	d6, d7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe6, 0xc0, 0x61, 0x1e}),
				address:          0,
			},
			want:    "fsqrt\td6, d7",
			wantErr: false,
		},
		{
			name: "fcvt	s8, d9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x28, 0x41, 0x62, 0x1e}),
				address:          0,
			},
			want:    "fcvt\ts8, d9",
			wantErr: false,
		},
		{
			name: "fcvt	h10, d11",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6a, 0xc1, 0x63, 0x1e}),
				address:          0,
			},
			want:    "fcvt\th10, d11",
			wantErr: false,
		},
		{
			name: "frintn	d12, d13",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x41, 0x64, 0x1e}),
				address:          0,
			},
			want:    "frintn\td12, d13",
			wantErr: false,
		},
		{
			name: "frintp	d14, d15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xee, 0xc1, 0x64, 0x1e}),
				address:          0,
			},
			want:    "frintp\td14, d15",
			wantErr: false,
		},
		{
			name: "frintm	d16, d17",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x30, 0x42, 0x65, 0x1e}),
				address:          0,
			},
			want:    "frintm\td16, d17",
			wantErr: false,
		},
		{
			name: "frintz	d18, d19",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x72, 0xc2, 0x65, 0x1e}),
				address:          0,
			},
			want:    "frintz\td18, d19",
			wantErr: false,
		},
		{
			name: "frinta	d20, d21",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb4, 0x42, 0x66, 0x1e}),
				address:          0,
			},
			want:    "frinta\td20, d21",
			wantErr: false,
		},
		{
			name: "frintx	d22, d23",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf6, 0x42, 0x67, 0x1e}),
				address:          0,
			},
			want:    "frintx\td22, d23",
			wantErr: false,
		},
		{
			name: "frinti	d24, d25",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x38, 0xc3, 0x67, 0x1e}),
				address:          0,
			},
			want:    "frinti\td24, d25",
			wantErr: false,
		},
		{
			name: "fcvt	s26, h27",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7a, 0x43, 0xe2, 0x1e}),
				address:          0,
			},
			want:    "fcvt\ts26, h27",
			wantErr: false,
		},
		{
			name: "fcvt	d28, h29",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbc, 0xc3, 0xe2, 0x1e}),
				address:          0,
			},
			want:    "fcvt\td28, h29",
			wantErr: false,
		},
		{
			name: "fmul	s20, s19, s17",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x74, 0x0a, 0x31, 0x1e}),
				address:          0,
			},
			want:    "fmul\ts20, s19, s17",
			wantErr: false,
		},
		{
			name: "fdiv	s1, s2, s3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x18, 0x23, 0x1e}),
				address:          0,
			},
			want:    "fdiv\ts1, s2, s3",
			wantErr: false,
		},
		{
			name: "fadd	s4, s5, s6",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa4, 0x28, 0x26, 0x1e}),
				address:          0,
			},
			want:    "fadd\ts4, s5, s6",
			wantErr: false,
		},
		{
			name: "fsub	s7, s8, s9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x07, 0x39, 0x29, 0x1e}),
				address:          0,
			},
			want:    "fsub\ts7, s8, s9",
			wantErr: false,
		},
		{
			name: "fmax	s10, s11, s12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6a, 0x49, 0x2c, 0x1e}),
				address:          0,
			},
			want:    "fmax\ts10, s11, s12",
			wantErr: false,
		},
		{
			name: "fmin	s13, s14, s15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcd, 0x59, 0x2f, 0x1e}),
				address:          0,
			},
			want:    "fmin\ts13, s14, s15",
			wantErr: false,
		},
		{
			name: "fmaxnm	s16, s17, s18",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x30, 0x6a, 0x32, 0x1e}),
				address:          0,
			},
			want:    "fmaxnm\ts16, s17, s18",
			wantErr: false,
		},
		{
			name: "fminnm	s19, s20, s21",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0x7a, 0x35, 0x1e}),
				address:          0,
			},
			want:    "fminnm\ts19, s20, s21",
			wantErr: false,
		},
		{
			name: "fnmul	s22, s23, s24",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf6, 0x8a, 0x38, 0x1e}),
				address:          0,
			},
			want:    "fnmul\ts22, s23, s24",
			wantErr: false,
		},
		{
			name: "fmul	d20, d19, d17",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x74, 0x0a, 0x71, 0x1e}),
				address:          0,
			},
			want:    "fmul\td20, d19, d17",
			wantErr: false,
		},
		{
			name: "fdiv	d1, d2, d3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0x18, 0x63, 0x1e}),
				address:          0,
			},
			want:    "fdiv\td1, d2, d3",
			wantErr: false,
		},
		{
			name: "fadd	d4, d5, d6",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa4, 0x28, 0x66, 0x1e}),
				address:          0,
			},
			want:    "fadd\td4, d5, d6",
			wantErr: false,
		},
		{
			name: "fsub	d7, d8, d9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x07, 0x39, 0x69, 0x1e}),
				address:          0,
			},
			want:    "fsub\td7, d8, d9",
			wantErr: false,
		},
		{
			name: "fmax	d10, d11, d12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6a, 0x49, 0x6c, 0x1e}),
				address:          0,
			},
			want:    "fmax\td10, d11, d12",
			wantErr: false,
		},
		{
			name: "fmin	d13, d14, d15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcd, 0x59, 0x6f, 0x1e}),
				address:          0,
			},
			want:    "fmin\td13, d14, d15",
			wantErr: false,
		},
		{
			name: "fmaxnm	d16, d17, d18",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x30, 0x6a, 0x72, 0x1e}),
				address:          0,
			},
			want:    "fmaxnm\td16, d17, d18",
			wantErr: false,
		},
		{
			name: "fminnm	d19, d20, d21",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0x7a, 0x75, 0x1e}),
				address:          0,
			},
			want:    "fminnm\td19, d20, d21",
			wantErr: false,
		},
		{
			name: "fnmul	d22, d23, d24",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf6, 0x8a, 0x78, 0x1e}),
				address:          0,
			},
			want:    "fnmul\td22, d23, d24",
			wantErr: false,
		},
		{
			name: "fmadd	s3, s5, s6, s31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0x7c, 0x06, 0x1f}),
				address:          0,
			},
			want:    "fmadd\ts3, s5, s6, s31",
			wantErr: false,
		},
		{
			name: "fmadd	d3, d13, d0, d23",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0x5d, 0x40, 0x1f}),
				address:          0,
			},
			want:    "fmadd\td3, d13, d0, d23",
			wantErr: false,
		},
		{
			name: "fmsub	s3, s5, s6, s31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xfc, 0x06, 0x1f}),
				address:          0,
			},
			want:    "fmsub\ts3, s5, s6, s31",
			wantErr: false,
		},
		{
			name: "fmsub	d3, d13, d0, d23",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xdd, 0x40, 0x1f}),
				address:          0,
			},
			want:    "fmsub\td3, d13, d0, d23",
			wantErr: false,
		},
		{
			name: "fnmadd	s3, s5, s6, s31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0x7c, 0x26, 0x1f}),
				address:          0,
			},
			want:    "fnmadd\ts3, s5, s6, s31",
			wantErr: false,
		},
		{
			name: "fnmadd	d3, d13, d0, d23",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0x5d, 0x60, 0x1f}),
				address:          0,
			},
			want:    "fnmadd\td3, d13, d0, d23",
			wantErr: false,
		},
		{
			name: "fnmsub	s3, s5, s6, s31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xfc, 0x26, 0x1f}),
				address:          0,
			},
			want:    "fnmsub\ts3, s5, s6, s31",
			wantErr: false,
		},
		{
			name: "fnmsub	d3, d13, d0, d23",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xdd, 0x60, 0x1f}),
				address:          0,
			},
			want:    "fnmsub\td3, d13, d0, d23",
			wantErr: false,
		},
		{
			name: "fcvtzs	w3, s5, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xfc, 0x18, 0x1e}),
				address:          0,
			},
			want:    "fcvtzs\tw3, s5, #0x1",
			wantErr: false,
		},
		{
			name: "fcvtzs	wzr, s20, #13",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0xce, 0x18, 0x1e}),
				address:          0,
			},
			want:    "fcvtzs\twzr, s20, #0xd",
			wantErr: false,
		},
		{
			name: "fcvtzs	w19, s0, #32",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x13, 0x80, 0x18, 0x1e}),
				address:          0,
			},
			want:    "fcvtzs\tw19, s0, #0x20",
			wantErr: false,
		},
		{
			name: "fcvtzs	x3, s5, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xfc, 0x18, 0x9e}),
				address:          0,
			},
			want:    "fcvtzs\tx3, s5, #0x1",
			wantErr: false,
		},
		{
			name: "fcvtzs	x12, s30, #45",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x4f, 0x18, 0x9e}),
				address:          0,
			},
			want:    "fcvtzs\tx12, s30, #0x2d",
			wantErr: false,
		},
		{
			name: "fcvtzs	x19, s0, #64",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x13, 0x00, 0x18, 0x9e}),
				address:          0,
			},
			want:    "fcvtzs\tx19, s0, #0x40",
			wantErr: false,
		},
		{
			name: "fcvtzs	w3, d5, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xfc, 0x58, 0x1e}),
				address:          0,
			},
			want:    "fcvtzs\tw3, d5, #0x1",
			wantErr: false,
		},
		{
			name: "fcvtzs	wzr, d20, #13",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0xce, 0x58, 0x1e}),
				address:          0,
			},
			want:    "fcvtzs\twzr, d20, #0xd",
			wantErr: false,
		},
		{
			name: "fcvtzs	w19, d0, #32",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x13, 0x80, 0x58, 0x1e}),
				address:          0,
			},
			want:    "fcvtzs\tw19, d0, #0x20",
			wantErr: false,
		},
		{
			name: "fcvtzs	x3, d5, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xfc, 0x58, 0x9e}),
				address:          0,
			},
			want:    "fcvtzs\tx3, d5, #0x1",
			wantErr: false,
		},
		{
			name: "fcvtzs	x12, d30, #45",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x4f, 0x58, 0x9e}),
				address:          0,
			},
			want:    "fcvtzs\tx12, d30, #0x2d",
			wantErr: false,
		},
		{
			name: "fcvtzs	x19, d0, #64",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x13, 0x00, 0x58, 0x9e}),
				address:          0,
			},
			want:    "fcvtzs\tx19, d0, #0x40",
			wantErr: false,
		},
		{
			name: "fcvtzu	w3, s5, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xfc, 0x19, 0x1e}),
				address:          0,
			},
			want:    "fcvtzu\tw3, s5, #0x1",
			wantErr: false,
		},
		{
			name: "fcvtzu	wzr, s20, #13",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0xce, 0x19, 0x1e}),
				address:          0,
			},
			want:    "fcvtzu\twzr, s20, #0xd",
			wantErr: false,
		},
		{
			name: "fcvtzu	w19, s0, #32",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x13, 0x80, 0x19, 0x1e}),
				address:          0,
			},
			want:    "fcvtzu\tw19, s0, #0x20",
			wantErr: false,
		},
		{
			name: "fcvtzu	x3, s5, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xfc, 0x19, 0x9e}),
				address:          0,
			},
			want:    "fcvtzu\tx3, s5, #0x1",
			wantErr: false,
		},
		{
			name: "fcvtzu	x12, s30, #45",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x4f, 0x19, 0x9e}),
				address:          0,
			},
			want:    "fcvtzu\tx12, s30, #0x2d",
			wantErr: false,
		},
		{
			name: "fcvtzu	x19, s0, #64",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x13, 0x00, 0x19, 0x9e}),
				address:          0,
			},
			want:    "fcvtzu\tx19, s0, #0x40",
			wantErr: false,
		},
		{
			name: "fcvtzu	w3, d5, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xfc, 0x59, 0x1e}),
				address:          0,
			},
			want:    "fcvtzu\tw3, d5, #0x1",
			wantErr: false,
		},
		{
			name: "fcvtzu	wzr, d20, #13",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0xce, 0x59, 0x1e}),
				address:          0,
			},
			want:    "fcvtzu\twzr, d20, #0xd",
			wantErr: false,
		},
		{
			name: "fcvtzu	w19, d0, #32",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x13, 0x80, 0x59, 0x1e}),
				address:          0,
			},
			want:    "fcvtzu\tw19, d0, #0x20",
			wantErr: false,
		},
		{
			name: "fcvtzu	x3, d5, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xfc, 0x59, 0x9e}),
				address:          0,
			},
			want:    "fcvtzu\tx3, d5, #0x1",
			wantErr: false,
		},
		{
			name: "fcvtzu	x12, d30, #45",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x4f, 0x59, 0x9e}),
				address:          0,
			},
			want:    "fcvtzu\tx12, d30, #0x2d",
			wantErr: false,
		},
		{
			name: "fcvtzu	x19, d0, #64",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x13, 0x00, 0x59, 0x9e}),
				address:          0,
			},
			want:    "fcvtzu\tx19, d0, #0x40",
			wantErr: false,
		},
		{
			name: "scvtf	s23, w19, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x77, 0xfe, 0x02, 0x1e}),
				address:          0,
			},
			want:    "scvtf\ts23, w19, #0x1",
			wantErr: false,
		},
		{
			name: "scvtf	s31, wzr, #20",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xb3, 0x02, 0x1e}),
				address:          0,
			},
			want:    "scvtf\ts31, wzr, #0x14",
			wantErr: false,
		},
		{
			name: "scvtf	s14, w0, #32",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0e, 0x80, 0x02, 0x1e}),
				address:          0,
			},
			want:    "scvtf\ts14, w0, #0x20",
			wantErr: false,
		},
		{
			name: "scvtf	s23, x19, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x77, 0xfe, 0x02, 0x9e}),
				address:          0,
			},
			want:    "scvtf\ts23, x19, #0x1",
			wantErr: false,
		},
		{
			name: "scvtf	s31, xzr, #20",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xb3, 0x02, 0x9e}),
				address:          0,
			},
			want:    "scvtf\ts31, xzr, #0x14",
			wantErr: false,
		},
		{
			name: "scvtf	s14, x0, #64",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0e, 0x00, 0x02, 0x9e}),
				address:          0,
			},
			want:    "scvtf\ts14, x0, #0x40",
			wantErr: false,
		},
		{
			name: "scvtf	d23, w19, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x77, 0xfe, 0x42, 0x1e}),
				address:          0,
			},
			want:    "scvtf\td23, w19, #0x1",
			wantErr: false,
		},
		{
			name: "scvtf	d31, wzr, #20",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xb3, 0x42, 0x1e}),
				address:          0,
			},
			want:    "scvtf\td31, wzr, #0x14",
			wantErr: false,
		},
		{
			name: "scvtf	d14, w0, #32",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0e, 0x80, 0x42, 0x1e}),
				address:          0,
			},
			want:    "scvtf\td14, w0, #0x20",
			wantErr: false,
		},
		{
			name: "scvtf	d23, x19, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x77, 0xfe, 0x42, 0x9e}),
				address:          0,
			},
			want:    "scvtf\td23, x19, #0x1",
			wantErr: false,
		},
		{
			name: "scvtf	d31, xzr, #20",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xb3, 0x42, 0x9e}),
				address:          0,
			},
			want:    "scvtf\td31, xzr, #0x14",
			wantErr: false,
		},
		{
			name: "scvtf	d14, x0, #64",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0e, 0x00, 0x42, 0x9e}),
				address:          0,
			},
			want:    "scvtf\td14, x0, #0x40",
			wantErr: false,
		},
		{
			name: "ucvtf	s23, w19, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x77, 0xfe, 0x03, 0x1e}),
				address:          0,
			},
			want:    "ucvtf\ts23, w19, #0x1",
			wantErr: false,
		},
		{
			name: "ucvtf	s31, wzr, #20",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xb3, 0x03, 0x1e}),
				address:          0,
			},
			want:    "ucvtf\ts31, wzr, #0x14",
			wantErr: false,
		},
		{
			name: "ucvtf	s14, w0, #32",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0e, 0x80, 0x03, 0x1e}),
				address:          0,
			},
			want:    "ucvtf\ts14, w0, #0x20",
			wantErr: false,
		},
		{
			name: "ucvtf	s23, x19, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x77, 0xfe, 0x03, 0x9e}),
				address:          0,
			},
			want:    "ucvtf\ts23, x19, #0x1",
			wantErr: false,
		},
		{
			name: "ucvtf	s31, xzr, #20",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xb3, 0x03, 0x9e}),
				address:          0,
			},
			want:    "ucvtf\ts31, xzr, #0x14",
			wantErr: false,
		},
		{
			name: "ucvtf	s14, x0, #64",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0e, 0x00, 0x03, 0x9e}),
				address:          0,
			},
			want:    "ucvtf\ts14, x0, #0x40",
			wantErr: false,
		},
		{
			name: "ucvtf	d23, w19, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x77, 0xfe, 0x43, 0x1e}),
				address:          0,
			},
			want:    "ucvtf\td23, w19, #0x1",
			wantErr: false,
		},
		{
			name: "ucvtf	d31, wzr, #20",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xb3, 0x43, 0x1e}),
				address:          0,
			},
			want:    "ucvtf\td31, wzr, #0x14",
			wantErr: false,
		},
		{
			name: "ucvtf	d14, w0, #32",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0e, 0x80, 0x43, 0x1e}),
				address:          0,
			},
			want:    "ucvtf\td14, w0, #0x20",
			wantErr: false,
		},
		{
			name: "ucvtf	d23, x19, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x77, 0xfe, 0x43, 0x9e}),
				address:          0,
			},
			want:    "ucvtf\td23, x19, #0x1",
			wantErr: false,
		},
		{
			name: "ucvtf	d31, xzr, #20",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xb3, 0x43, 0x9e}),
				address:          0,
			},
			want:    "ucvtf\td31, xzr, #0x14",
			wantErr: false,
		},
		{
			name: "ucvtf	d14, x0, #64",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0e, 0x00, 0x43, 0x9e}),
				address:          0,
			},
			want:    "ucvtf\td14, x0, #0x40",
			wantErr: false,
		},
		{
			name: "fcvtns	w3, s31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0x20, 0x1e}),
				address:          0,
			},
			want:    "fcvtns\tw3, s31",
			wantErr: false,
		},
		{
			name: "fcvtns	xzr, s12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x01, 0x20, 0x9e}),
				address:          0,
			},
			want:    "fcvtns\txzr, s12",
			wantErr: false,
		},
		{
			name: "fcvtnu	wzr, s12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x01, 0x21, 0x1e}),
				address:          0,
			},
			want:    "fcvtnu\twzr, s12",
			wantErr: false,
		},
		{
			name: "fcvtnu	x0, s0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x00, 0x21, 0x9e}),
				address:          0,
			},
			want:    "fcvtnu\tx0, s0",
			wantErr: false,
		},
		{
			name: "fcvtps	wzr, s9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0x01, 0x28, 0x1e}),
				address:          0,
			},
			want:    "fcvtps\twzr, s9",
			wantErr: false,
		},
		{
			name: "fcvtps	x12, s20",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0x02, 0x28, 0x9e}),
				address:          0,
			},
			want:    "fcvtps\tx12, s20",
			wantErr: false,
		},
		{
			name: "fcvtpu	w30, s23",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfe, 0x02, 0x29, 0x1e}),
				address:          0,
			},
			want:    "fcvtpu\tw30, s23",
			wantErr: false,
		},
		{
			name: "fcvtpu	x29, s3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7d, 0x00, 0x29, 0x9e}),
				address:          0,
			},
			want:    "fcvtpu\tfp, s3",
			wantErr: false,
		},
		{
			name: "fcvtms	w2, s3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x00, 0x30, 0x1e}),
				address:          0,
			},
			want:    "fcvtms\tw2, s3",
			wantErr: false,
		},
		{
			name: "fcvtms	x4, s5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa4, 0x00, 0x30, 0x9e}),
				address:          0,
			},
			want:    "fcvtms\tx4, s5",
			wantErr: false,
		},
		{
			name: "fcvtmu	w6, s7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe6, 0x00, 0x31, 0x1e}),
				address:          0,
			},
			want:    "fcvtmu\tw6, s7",
			wantErr: false,
		},
		{
			name: "fcvtmu	x8, s9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x28, 0x01, 0x31, 0x9e}),
				address:          0,
			},
			want:    "fcvtmu\tx8, s9",
			wantErr: false,
		},
		{
			name: "fcvtzs	w10, s11",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6a, 0x01, 0x38, 0x1e}),
				address:          0,
			},
			want:    "fcvtzs\tw10, s11",
			wantErr: false,
		},
		{
			name: "fcvtzs	x12, s13",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x01, 0x38, 0x9e}),
				address:          0,
			},
			want:    "fcvtzs\tx12, s13",
			wantErr: false,
		},
		{
			name: "fcvtzu	w14, s15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xee, 0x01, 0x39, 0x1e}),
				address:          0,
			},
			want:    "fcvtzu\tw14, s15",
			wantErr: false,
		},
		{
			name: "fcvtzu	x15, s16",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0f, 0x02, 0x39, 0x9e}),
				address:          0,
			},
			want:    "fcvtzu\tx15, s16",
			wantErr: false,
		},
		{
			name: "scvtf	s17, w18",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x51, 0x02, 0x22, 0x1e}),
				address:          0,
			},
			want:    "scvtf\ts17, w18",
			wantErr: false,
		},
		{
			name: "scvtf	s19, x20",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0x02, 0x22, 0x9e}),
				address:          0,
			},
			want:    "scvtf\ts19, x20",
			wantErr: false,
		},
		{
			name: "ucvtf	s21, w22",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd5, 0x02, 0x23, 0x1e}),
				address:          0,
			},
			want:    "ucvtf\ts21, w22",
			wantErr: false,
		},
		{
			name: "scvtf	s23, x24",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x17, 0x03, 0x22, 0x9e}),
				address:          0,
			},
			want:    "scvtf\ts23, x24",
			wantErr: false,
		},
		{
			name: "fcvtas	w25, s26",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x59, 0x03, 0x24, 0x1e}),
				address:          0,
			},
			want:    "fcvtas\tw25, s26",
			wantErr: false,
		},
		{
			name: "fcvtas	x27, s28",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9b, 0x03, 0x24, 0x9e}),
				address:          0,
			},
			want:    "fcvtas\tx27, s28",
			wantErr: false,
		},
		{
			name: "fcvtau	w29, s30",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdd, 0x03, 0x25, 0x1e}),
				address:          0,
			},
			want:    "fcvtau\tw29, s30",
			wantErr: false,
		},
		{
			name: "fcvtau	xzr, s0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x00, 0x25, 0x9e}),
				address:          0,
			},
			want:    "fcvtau\txzr, s0",
			wantErr: false,
		},
		{
			name: "fcvtns	w3, d31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0x60, 0x1e}),
				address:          0,
			},
			want:    "fcvtns\tw3, d31",
			wantErr: false,
		},
		{
			name: "fcvtns	xzr, d12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x01, 0x60, 0x9e}),
				address:          0,
			},
			want:    "fcvtns\txzr, d12",
			wantErr: false,
		},
		{
			name: "fcvtnu	wzr, d12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x01, 0x61, 0x1e}),
				address:          0,
			},
			want:    "fcvtnu\twzr, d12",
			wantErr: false,
		},
		{
			name: "fcvtnu	x0, d0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x00, 0x61, 0x9e}),
				address:          0,
			},
			want:    "fcvtnu\tx0, d0",
			wantErr: false,
		},
		{
			name: "fcvtps	wzr, d9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0x01, 0x68, 0x1e}),
				address:          0,
			},
			want:    "fcvtps\twzr, d9",
			wantErr: false,
		},
		{
			name: "fcvtps	x12, d20",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0x02, 0x68, 0x9e}),
				address:          0,
			},
			want:    "fcvtps\tx12, d20",
			wantErr: false,
		},
		{
			name: "fcvtpu	w30, d23",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfe, 0x02, 0x69, 0x1e}),
				address:          0,
			},
			want:    "fcvtpu\tw30, d23",
			wantErr: false,
		},
		{
			name: "fcvtpu	x29, d3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7d, 0x00, 0x69, 0x9e}),
				address:          0,
			},
			want:    "fcvtpu\tfp, d3",
			wantErr: false,
		},
		{
			name: "fcvtms	w2, d3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x00, 0x70, 0x1e}),
				address:          0,
			},
			want:    "fcvtms\tw2, d3",
			wantErr: false,
		},
		{
			name: "fcvtms	x4, d5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa4, 0x00, 0x70, 0x9e}),
				address:          0,
			},
			want:    "fcvtms\tx4, d5",
			wantErr: false,
		},
		{
			name: "fcvtmu	w6, d7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe6, 0x00, 0x71, 0x1e}),
				address:          0,
			},
			want:    "fcvtmu\tw6, d7",
			wantErr: false,
		},
		{
			name: "fcvtmu	x8, d9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x28, 0x01, 0x71, 0x9e}),
				address:          0,
			},
			want:    "fcvtmu\tx8, d9",
			wantErr: false,
		},
		{
			name: "fcvtzs	w10, d11",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6a, 0x01, 0x78, 0x1e}),
				address:          0,
			},
			want:    "fcvtzs\tw10, d11",
			wantErr: false,
		},
		{
			name: "fcvtzs	x12, d13",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x01, 0x78, 0x9e}),
				address:          0,
			},
			want:    "fcvtzs\tx12, d13",
			wantErr: false,
		},
		{
			name: "fcvtzu	w14, d15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xee, 0x01, 0x79, 0x1e}),
				address:          0,
			},
			want:    "fcvtzu\tw14, d15",
			wantErr: false,
		},
		{
			name: "fcvtzu	x15, d16",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0f, 0x02, 0x79, 0x9e}),
				address:          0,
			},
			want:    "fcvtzu\tx15, d16",
			wantErr: false,
		},
		{
			name: "scvtf	d17, w18",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x51, 0x02, 0x62, 0x1e}),
				address:          0,
			},
			want:    "scvtf\td17, w18",
			wantErr: false,
		},
		{
			name: "scvtf	d19, x20",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0x02, 0x62, 0x9e}),
				address:          0,
			},
			want:    "scvtf\td19, x20",
			wantErr: false,
		},
		{
			name: "ucvtf	d21, w22",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd5, 0x02, 0x63, 0x1e}),
				address:          0,
			},
			want:    "ucvtf\td21, w22",
			wantErr: false,
		},
		{
			name: "ucvtf	d23, x24",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x17, 0x03, 0x63, 0x9e}),
				address:          0,
			},
			want:    "ucvtf\td23, x24",
			wantErr: false,
		},
		{
			name: "fcvtas	w25, d26",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x59, 0x03, 0x64, 0x1e}),
				address:          0,
			},
			want:    "fcvtas\tw25, d26",
			wantErr: false,
		},
		{
			name: "fcvtas	x27, d28",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9b, 0x03, 0x64, 0x9e}),
				address:          0,
			},
			want:    "fcvtas\tx27, d28",
			wantErr: false,
		},
		{
			name: "fcvtau	w29, d30",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdd, 0x03, 0x65, 0x1e}),
				address:          0,
			},
			want:    "fcvtau\tw29, d30",
			wantErr: false,
		},
		{
			name: "fcvtau	xzr, d0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x00, 0x65, 0x9e}),
				address:          0,
			},
			want:    "fcvtau\txzr, d0",
			wantErr: false,
		},
		{
			name: "fmov	w3, s9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x23, 0x01, 0x26, 0x1e}),
				address:          0,
			},
			want:    "fmov\tw3, s9",
			wantErr: false,
		},
		{
			name: "fmov	s9, w3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0x00, 0x27, 0x1e}),
				address:          0,
			},
			want:    "fmov\ts9, w3",
			wantErr: false,
		},
		{
			name: "fmov	x20, d31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x03, 0x66, 0x9e}),
				address:          0,
			},
			want:    "fmov\tx20, d31",
			wantErr: false,
		},
		{
			name: "fmov	d1, x15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe1, 0x01, 0x67, 0x9e}),
				address:          0,
			},
			want:    "fmov\td1, x15",
			wantErr: false,
		},
		{
			name: "fmov	x3, v12.d[1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x83, 0x01, 0xae, 0x9e}),
				address:          0,
			},
			want:    "fmov\tx3, v12.d[1]",
			wantErr: false,
		},
		{
			name: "fmov	v1.d[1], x19",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x61, 0x02, 0xaf, 0x9e}),
				address:          0,
			},
			want:    "fmov\tv1.d[1], x19",
			wantErr: false,
		},
		{
			name: "fmov	v3.d[1], xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0xaf, 0x9e}),
				address:          0,
			},
			want:    "fmov\tv3.d[1], xzr",
			wantErr: false,
		},
		{
			name: "fmov	s2, #0.12500000",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x02, 0x10, 0x28, 0x1e}),
				address:          0,
			},
			want:    "fmov\ts2, #0.12500000",
			wantErr: false,
		},
		{
			name: "fmov	s3, #1.00000000",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x03, 0x10, 0x2e, 0x1e}),
				address:          0,
			},
			want:    "fmov\ts3, #1.00000000",
			wantErr: false,
		},
		{
			name: "fmov	d30, #16.00000000",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1e, 0x10, 0x66, 0x1e}),
				address:          0,
			},
			want:    "fmov\td30, #16.00000000",
			wantErr: false,
		},
		{
			name: "fmov	s4, #1.06250000",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x04, 0x30, 0x2e, 0x1e}),
				address:          0,
			},
			want:    "fmov\ts4, #1.06250000",
			wantErr: false,
		},
		{
			name: "fmov	d10, #1.93750000",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0a, 0xf0, 0x6f, 0x1e}),
				address:          0,
			},
			want:    "fmov\td10, #1.93750000",
			wantErr: false,
		},
		{
			name: "fmov	s12, #-1.00000000",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x10, 0x3e, 0x1e}),
				address:          0,
			},
			want:    "fmov\ts12, #-1.00000000",
			wantErr: false,
		},
		{
			name: "fmov	d16, #8.50000000",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x10, 0x30, 0x64, 0x1e}),
				address:          0,
			},
			want:    "fmov\td16, #8.50000000",
			wantErr: false,
		},
		{
			name: "ldr	w0, #1048572",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0xff, 0x7f, 0x18}),
				address:          0,
			},
			want:    "ldr\tw0, 0xffffc",
			wantErr: false,
		},
		{
			name: "ldr	x10, #-1048576",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0a, 0x00, 0x80, 0x58}),
				address:          0,
			},
			want:    "ldr\tx10, 0xfffffffffff00000",
			wantErr: false,
		},
		{
			name: "stxrb	w1, w2, [x3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x7c, 0x01, 0x08}),
				address:          0,
			},
			want:    "stxrb\tw1, w2, [x3]",
			wantErr: false,
		},
		{
			name: "stxrh	w2, w3, [x4]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x83, 0x7c, 0x02, 0x48}),
				address:          0,
			},
			want:    "stxrh\tw2, w3, [x4]",
			wantErr: false,
		},
		{
			name: "stxr	wzr, w4, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe4, 0x7f, 0x1f, 0x88}),
				address:          0,
			},
			want:    "stxr\twzr, w4, [sp]",
			wantErr: false,
		},
		{
			name: "stxr	w5, x6, [x7]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe6, 0x7c, 0x05, 0xc8}),
				address:          0,
			},
			want:    "stxr\tw5, x6, [x7]",
			wantErr: false,
		},
		{
			name: "ldxrb	w7, [x9]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x27, 0x7d, 0x5f, 0x08}),
				address:          0,
			},
			want:    "ldxrb\tw7, [x9]",
			wantErr: false,
		},
		{
			name: "ldxrh	wzr, [x10]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x7d, 0x5f, 0x48}),
				address:          0,
			},
			want:    "ldxrh\twzr, [x10]",
			wantErr: false,
		},
		{
			name: "ldxr	w9, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x7f, 0x5f, 0x88}),
				address:          0,
			},
			want:    "ldxr\tw9, [sp]",
			wantErr: false,
		},
		{
			name: "ldxr	x10, [x11]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6a, 0x7d, 0x5f, 0xc8}),
				address:          0,
			},
			want:    "ldxr\tx10, [x11]",
			wantErr: false,
		},
		{
			name: "stxp	w11, w12, w13, [x14]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x35, 0x2b, 0x88}),
				address:          0,
			},
			want:    "stxp\tw11, w12, w13, [x14]",
			wantErr: false,
		},
		{
			name: "stxp	wzr, x23, x14, [x15]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf7, 0x39, 0x3f, 0xc8}),
				address:          0,
			},
			want:    "stxp\twzr, x23, x14, [x15]",
			wantErr: false,
		},
		{
			name: "ldxp	w12, wzr, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x7f, 0x7f, 0x88}),
				address:          0,
			},
			want:    "ldxp\tw12, wzr, [sp]",
			wantErr: false,
		},
		{
			name: "ldxp	x13, x14, [x15]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xed, 0x39, 0x7f, 0xc8}),
				address:          0,
			},
			want:    "ldxp\tx13, x14, [x15]",
			wantErr: false,
		},
		{
			name: "stlxrb	w14, w15, [x16]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0f, 0xfe, 0x0e, 0x08}),
				address:          0,
			},
			want:    "stlxrb\tw14, w15, [x16]",
			wantErr: false,
		},
		{
			name: "stlxrh	w15, w16, [x17]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x30, 0xfe, 0x0f, 0x48}),
				address:          0,
			},
			want:    "stlxrh\tw15, w16, [x17]",
			wantErr: false,
		},
		{
			name: "stlxr	wzr, w17, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf1, 0xff, 0x1f, 0x88}),
				address:          0,
			},
			want:    "stlxr\twzr, w17, [sp]",
			wantErr: false,
		},
		{
			name: "stlxr	w18, x19, [x20]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0xfe, 0x12, 0xc8}),
				address:          0,
			},
			want:    "stlxr\tw18, x19, [x20]",
			wantErr: false,
		},
		{
			name: "ldaxrb	w19, [x21]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb3, 0xfe, 0x5f, 0x08}),
				address:          0,
			},
			want:    "ldaxrb\tw19, [x21]",
			wantErr: false,
		},
		{
			name: "ldaxrh	w20, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0xff, 0x5f, 0x48}),
				address:          0,
			},
			want:    "ldaxrh\tw20, [sp]",
			wantErr: false,
		},
		{
			name: "ldaxr	wzr, [x22]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0xfe, 0x5f, 0x88}),
				address:          0,
			},
			want:    "ldaxr\twzr, [x22]",
			wantErr: false,
		},
		{
			name: "ldaxr	x21, [x23]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf5, 0xfe, 0x5f, 0xc8}),
				address:          0,
			},
			want:    "ldaxr\tx21, [x23]",
			wantErr: false,
		},
		{
			name: "stlxp	wzr, w22, w23, [x24]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x16, 0xdf, 0x3f, 0x88}),
				address:          0,
			},
			want:    "stlxp\twzr, w22, w23, [x24]",
			wantErr: false,
		},
		{
			name: "stlxp	w25, x26, x27, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfa, 0xef, 0x39, 0xc8}),
				address:          0,
			},
			want:    "stlxp\tw25, x26, x27, [sp]",
			wantErr: false,
		},
		{
			name: "ldaxp	w26, wzr, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfa, 0xff, 0x7f, 0x88}),
				address:          0,
			},
			want:    "ldaxp\tw26, wzr, [sp]",
			wantErr: false,
		},
		{
			name: "ldaxp	x27, x28, [x30]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdb, 0xf3, 0x7f, 0xc8}),
				address:          0,
			},
			want:    "ldaxp\tx27, x28, [lr]",
			wantErr: false,
		},
		{
			name: "stlrb	w27, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfb, 0xff, 0x9f, 0x08}),
				address:          0,
			},
			want:    "stlrb\tw27, [sp]",
			wantErr: false,
		},
		{
			name: "stlrh	w28, [x0]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1c, 0xfc, 0x9f, 0x48}),
				address:          0,
			},
			want:    "stlrh\tw28, [x0]",
			wantErr: false,
		},
		{
			name: "stlr	wzr, [x1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0xfc, 0x9f, 0x88}),
				address:          0,
			},
			want:    "stlr\twzr, [x1]",
			wantErr: false,
		},
		{
			name: "stlr	x30, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5e, 0xfc, 0x9f, 0xc8}),
				address:          0,
			},
			want:    "stlr\tlr, [x2]",
			wantErr: false,
		},
		{
			name: "ldarb	w29, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfd, 0xff, 0xdf, 0x08}),
				address:          0,
			},
			want:    "ldarb\tw29, [sp]",
			wantErr: false,
		},
		{
			name: "ldarh	w30, [x0]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1e, 0xfc, 0xdf, 0x48}),
				address:          0,
			},
			want:    "ldarh\tw30, [x0]",
			wantErr: false,
		},
		{
			name: "ldar	wzr, [x1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0xfc, 0xdf, 0x88}),
				address:          0,
			},
			want:    "ldar\twzr, [x1]",
			wantErr: false,
		},
		{
			name: "ldar	x1, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0xfc, 0xdf, 0xc8}),
				address:          0,
			},
			want:    "ldar\tx1, [x2]",
			wantErr: false,
		},
		{
			name: "stlxp	wzr, w22, w23, [x24]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x16, 0xdf, 0x3f, 0x88}),
				address:          0,
			},
			want:    "stlxp\twzr, w22, w23, [x24]",
			wantErr: false,
		},
		{
			name: "sturb	w9, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x03, 0x00, 0x38}),
				address:          0,
			},
			want:    "sturb\tw9, [sp]",
			wantErr: false,
		},
		{
			name: "sturh	wzr, [x12, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0xf1, 0x0f, 0x78}),
				address:          0,
			},
			want:    "sturh\twzr, [x12, #0xff]",
			wantErr: false,
		},
		{
			name: "stur	w16, [x0, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x10, 0x00, 0x10, 0xb8}),
				address:          0,
			},
			want:    "stur\tw16, [x0, #-0x100]",
			wantErr: false,
		},
		{
			name: "stur	x28, [x14, #1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdc, 0x11, 0x00, 0xf8}),
				address:          0,
			},
			want:    "stur\tx28, [x14, #0x1]",
			wantErr: false,
		},
		{
			name: "ldurb	w1, [x20, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x81, 0xf2, 0x4f, 0x38}),
				address:          0,
			},
			want:    "ldurb\tw1, [x20, #0xff]",
			wantErr: false,
		},
		{
			name: "ldurh	w20, [x1, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x34, 0xf0, 0x4f, 0x78}),
				address:          0,
			},
			want:    "ldurh\tw20, [x1, #0xff]",
			wantErr: false,
		},
		{
			name: "ldur	w12, [sp, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0xf3, 0x4f, 0xb8}),
				address:          0,
			},
			want:    "ldur\tw12, [sp, #0xff]",
			wantErr: false,
		},
		{
			name: "ldur	xzr, [x12, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0xf1, 0x4f, 0xf8}),
				address:          0,
			},
			want:    "ldur\txzr, [x12, #0xff]",
			wantErr: false,
		},
		{
			name: "ldursb	x9, [x7, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x00, 0x90, 0x38}),
				address:          0,
			},
			want:    "ldursb\tx9, [x7, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldursh	x17, [x19, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x71, 0x02, 0x90, 0x78}),
				address:          0,
			},
			want:    "ldursh\tx17, [x19, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldursw	x20, [x15, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x01, 0x90, 0xb8}),
				address:          0,
			},
			want:    "ldursw\tx20, [x15, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldursw	x13, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4d, 0x00, 0x80, 0xb8}),
				address:          0,
			},
			want:    "ldursw\tx13, [x2]",
			wantErr: false,
		},
		{
			name: "prfum	pldl2keep, [sp, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe2, 0x03, 0x90, 0xf8}),
				address:          0,
			},
			want:    "prfum\tpldl2keep, [sp, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldursb	w19, [x1, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x33, 0x00, 0xd0, 0x38}),
				address:          0,
			},
			want:    "ldursb\tw19, [x1, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldursh	w15, [x21, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xaf, 0x02, 0xd0, 0x78}),
				address:          0,
			},
			want:    "ldursh\tw15, [x21, #-0x100]",
			wantErr: false,
		},
		{
			name: "stur	b0, [sp, #1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x13, 0x00, 0x3c}),
				address:          0,
			},
			want:    "stur\tb0, [sp, #0x1]",
			wantErr: false,
		},
		{
			name: "stur	h12, [x12, #-1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0xf1, 0x1f, 0x7c}),
				address:          0,
			},
			want:    "stur\th12, [x12, #-0x1]",
			wantErr: false,
		},
		{
			name: "stur	s15, [x0, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0f, 0xf0, 0x0f, 0xbc}),
				address:          0,
			},
			want:    "stur\ts15, [x0, #0xff]",
			wantErr: false,
		},
		{
			name: "stur	d31, [x5, #25]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x90, 0x01, 0xfc}),
				address:          0,
			},
			want:    "stur\td31, [x5, #0x19]",
			wantErr: false,
		},
		{
			name: "stur	q9, [x5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x00, 0x80, 0x3c}),
				address:          0,
			},
			want:    "stur\tq9, [x5]",
			wantErr: false,
		},
		{
			name: "ldur	b3, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0x40, 0x3c}),
				address:          0,
			},
			want:    "ldur\tb3, [sp]",
			wantErr: false,
		},
		{
			name: "ldur	h5, [x4, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x85, 0x00, 0x50, 0x7c}),
				address:          0,
			},
			want:    "ldur\th5, [x4, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldur	s7, [x12, #-1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x87, 0xf1, 0x5f, 0xbc}),
				address:          0,
			},
			want:    "ldur\ts7, [x12, #-0x1]",
			wantErr: false,
		},
		{
			name: "ldur	d11, [x19, #4]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6b, 0x42, 0x40, 0xfc}),
				address:          0,
			},
			want:    "ldur\td11, [x19, #0x4]",
			wantErr: false,
		},
		{
			name: "ldur	q13, [x1, #2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2d, 0x20, 0xc0, 0x3c}),
				address:          0,
			},
			want:    "ldur\tq13, [x1, #0x2]",
			wantErr: false,
		},
		{
			name: "ldr	x0, [x0]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x00, 0x40, 0xf9}),
				address:          0,
			},
			want:    "ldr\tx0, [x0]",
			wantErr: false,
		},
		{
			name: "ldr	x4, [x29]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa4, 0x03, 0x40, 0xf9}),
				address:          0,
			},
			want:    "ldr\tx4, [fp]",
			wantErr: false,
		},
		{
			name: "ldr	x30, [x12, #32760]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9e, 0xfd, 0x7f, 0xf9}),
				address:          0,
			},
			want:    "ldr\tlr, [x12, #0x7ff8]",
			wantErr: false,
		},
		{
			name: "ldr	x20, [sp, #8]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x07, 0x40, 0xf9}),
				address:          0,
			},
			want:    "ldr\tx20, [sp, #0x8]",
			wantErr: false,
		},
		{
			name: "ldr	xzr, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x40, 0xf9}),
				address:          0,
			},
			want:    "ldr\txzr, [sp]",
			wantErr: false,
		},
		{
			name: "ldr	w2, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe2, 0x03, 0x40, 0xb9}),
				address:          0,
			},
			want:    "ldr\tw2, [sp]",
			wantErr: false,
		},
		{
			name: "ldr	w17, [sp, #16380]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf1, 0xff, 0x7f, 0xb9}),
				address:          0,
			},
			want:    "ldr\tw17, [sp, #0x3ffc]",
			wantErr: false,
		},
		{
			name: "ldr	w13, [x2, #4]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4d, 0x04, 0x40, 0xb9}),
				address:          0,
			},
			want:    "ldr\tw13, [x2, #0x4]",
			wantErr: false,
		},
		{
			name: "ldrsw	x2, [x5, #4]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa2, 0x04, 0x80, 0xb9}),
				address:          0,
			},
			want:    "ldrsw\tx2, [x5, #0x4]",
			wantErr: false,
		},
		{
			name: "ldrsw	x23, [sp, #16380]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf7, 0xff, 0xbf, 0xb9}),
				address:          0,
			},
			want:    "ldrsw\tx23, [sp, #0x3ffc]",
			wantErr: false,
		},
		{
			name: "ldrh	w2, [x4]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x82, 0x00, 0x40, 0x79}),
				address:          0,
			},
			want:    "ldrh\tw2, [x4]",
			wantErr: false,
		},
		{
			name: "ldrsh	w23, [x6, #8190]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd7, 0xfc, 0xff, 0x79}),
				address:          0,
			},
			want:    "ldrsh\tw23, [x6, #0x1ffe]",
			wantErr: false,
		},
		{
			name: "ldrsh	wzr, [sp, #2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x07, 0xc0, 0x79}),
				address:          0,
			},
			want:    "ldrsh\twzr, [sp, #0x2]",
			wantErr: false,
		},
		{
			name: "ldrsh	x29, [x2, #2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5d, 0x04, 0x80, 0x79}),
				address:          0,
			},
			want:    "ldrsh\tfp, [x2, #0x2]",
			wantErr: false,
		},
		{
			name: "ldrb	w26, [x3, #121]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7a, 0xe4, 0x41, 0x39}),
				address:          0,
			},
			want:    "ldrb\tw26, [x3, #0x79]",
			wantErr: false,
		},
		{
			name: "ldrb	w12, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0x00, 0x40, 0x39}),
				address:          0,
			},
			want:    "ldrb\tw12, [x2]",
			wantErr: false,
		},
		{
			name: "ldrsb	w27, [sp, #4095]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfb, 0xff, 0xff, 0x39}),
				address:          0,
			},
			want:    "ldrsb\tw27, [sp, #0xfff]",
			wantErr: false,
		},
		{
			name: "ldrsb	xzr, [x15]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x01, 0x80, 0x39}),
				address:          0,
			},
			want:    "ldrsb\txzr, [x15]",
			wantErr: false,
		},
		{
			name: "str	x30, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfe, 0x03, 0x00, 0xf9}),
				address:          0,
			},
			want:    "str\tlr, [sp]",
			wantErr: false,
		},
		{
			name: "str	w20, [x4, #16380]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x94, 0xfc, 0x3f, 0xb9}),
				address:          0,
			},
			want:    "str\tw20, [x4, #0x3ffc]",
			wantErr: false,
		},
		{
			name: "strh	w20, [x10, #14]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x54, 0x1d, 0x00, 0x79}),
				address:          0,
			},
			want:    "strh\tw20, [x10, #0xe]",
			wantErr: false,
		},
		{
			name: "strh	w17, [sp, #8190]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf1, 0xff, 0x3f, 0x79}),
				address:          0,
			},
			want:    "strh\tw17, [sp, #0x1ffe]",
			wantErr: false,
		},
		{
			name: "strb	w23, [x3, #4095]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x77, 0xfc, 0x3f, 0x39}),
				address:          0,
			},
			want:    "strb\tw23, [x3, #0xfff]",
			wantErr: false,
		},
		{
			name: "strb	wzr, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x00, 0x00, 0x39}),
				address:          0,
			},
			want:    "strb\twzr, [x2]",
			wantErr: false,
		},
		{
			name: "prfm	pldl1keep, [sp, #8]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x07, 0x80, 0xf9}),
				address:          0,
			},
			want:    "prfm\tpldl1keep, [sp, #0x8]",
			wantErr: false,
		},
		{
			name: "prfm	pldl1strm, [x3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x61, 0x00, 0x80, 0xf9}),
				address:          0,
			},
			want:    "prfm\tpldl1strm, [x3]",
			wantErr: false,
		},
		{
			name: "prfm	pldl2keep, [x5, #16]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa2, 0x08, 0x80, 0xf9}),
				address:          0,
			},
			want:    "prfm\tpldl2keep, [x5, #0x10]",
			wantErr: false,
		},
		{
			name: "prfm	pldl2strm, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x43, 0x00, 0x80, 0xf9}),
				address:          0,
			},
			want:    "prfm\tpldl2strm, [x2]",
			wantErr: false,
		},
		{
			name: "prfm	pldl3keep, [x5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa4, 0x00, 0x80, 0xf9}),
				address:          0,
			},
			want:    "prfm\tpldl3keep, [x5]",
			wantErr: false,
		},
		{
			name: "prfm	pldl3strm, [x6]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc5, 0x00, 0x80, 0xf9}),
				address:          0,
			},
			want:    "prfm\tpldl3strm, [x6]",
			wantErr: false,
		},
		{
			name: "prfm	plil1keep, [sp, #8]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe8, 0x07, 0x80, 0xf9}),
				address:          0,
			},
			want:    "prfm\tplil1keep, [sp, #0x8]",
			wantErr: false,
		},
		{
			name: "prfm	plil1strm, [x3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0x00, 0x80, 0xf9}),
				address:          0,
			},
			want:    "prfm\tplil1strm, [x3]",
			wantErr: false,
		},
		{
			name: "prfm	plil2keep, [x5, #16]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xaa, 0x08, 0x80, 0xf9}),
				address:          0,
			},
			want:    "prfm\tplil2keep, [x5, #0x10]",
			wantErr: false,
		},
		{
			name: "prfm	plil2strm, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4b, 0x00, 0x80, 0xf9}),
				address:          0,
			},
			want:    "prfm\tplil2strm, [x2]",
			wantErr: false,
		},
		{
			name: "prfm	plil3keep, [x5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x00, 0x80, 0xf9}),
				address:          0,
			},
			want:    "prfm\tplil3keep, [x5]",
			wantErr: false,
		},
		{
			name: "prfm	plil3strm, [x6]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcd, 0x00, 0x80, 0xf9}),
				address:          0,
			},
			want:    "prfm\tplil3strm, [x6]",
			wantErr: false,
		},
		{
			name: "prfm	pstl1keep, [sp, #8]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf0, 0x07, 0x80, 0xf9}),
				address:          0,
			},
			want:    "prfm\tpstl1keep, [sp, #0x8]",
			wantErr: false,
		},
		{
			name: "prfm	pstl1strm, [x3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x71, 0x00, 0x80, 0xf9}),
				address:          0,
			},
			want:    "prfm\tpstl1strm, [x3]",
			wantErr: false,
		},
		{
			name: "prfm	pstl2keep, [x5, #16]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb2, 0x08, 0x80, 0xf9}),
				address:          0,
			},
			want:    "prfm\tpstl2keep, [x5, #0x10]",
			wantErr: false,
		},
		{
			name: "prfm	pstl2strm, [x2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x53, 0x00, 0x80, 0xf9}),
				address:          0,
			},
			want:    "prfm\tpstl2strm, [x2]",
			wantErr: false,
		},
		{
			name: "prfm	pstl3keep, [x5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb4, 0x00, 0x80, 0xf9}),
				address:          0,
			},
			want:    "prfm\tpstl3keep, [x5]",
			wantErr: false,
		},
		{
			name: "prfm	pstl3strm, [x6]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd5, 0x00, 0x80, 0xf9}),
				address:          0,
			},
			want:    "prfm\tpstl3strm, [x6]",
			wantErr: false,
		},
		{
			name: "prfm	#15, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xef, 0x03, 0x80, 0xf9}),
				address:          0,
			},
			want:    "prfm\t#15, [sp]",
			wantErr: false,
		},
		{
			name: "ldr	b31, [sp, #4095]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xff, 0x7f, 0x3d}),
				address:          0,
			},
			want:    "ldr\tb31, [sp, #0xfff]",
			wantErr: false,
		},
		{
			name: "ldr	h20, [x2, #8190]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x54, 0xfc, 0x7f, 0x7d}),
				address:          0,
			},
			want:    "ldr\th20, [x2, #0x1ffe]",
			wantErr: false,
		},
		{
			name: "ldr	s10, [x19, #16380]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6a, 0xfe, 0x7f, 0xbd}),
				address:          0,
			},
			want:    "ldr\ts10, [x19, #0x3ffc]",
			wantErr: false,
		},
		{
			name: "ldr	d3, [x10, #32760]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x43, 0xfd, 0x7f, 0xfd}),
				address:          0,
			},
			want:    "ldr\td3, [x10, #0x7ff8]",
			wantErr: false,
		},
		{
			name: "str	q12, [sp, #65520]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0xff, 0xbf, 0x3d}),
				address:          0,
			},
			want:    "str\tq12, [sp, #0xfff0]",
			wantErr: false,
		},
		{
			name: "ldrb	w3, [sp, x5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x6b, 0x65, 0x38}),
				address:          0,
			},
			want:    "ldrb\tw3, [sp, x5]",
			wantErr: false,
		},
		{
			name: "ldrb	w9, [x27, x6, lsl #0]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0x7b, 0x66, 0x38}),
				address:          0,
			},
			want:    "ldrb\tw9, [x27, x6, lsl #0]",
			wantErr: false,
		},
		{
			name: "ldrsb	w10, [x30, x7]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xca, 0x6b, 0xe7, 0x38}),
				address:          0,
			},
			want:    "ldrsb\tw10, [lr, x7]",
			wantErr: false,
		},
		{
			name: "ldrb	w11, [x29, x3, sxtx]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xab, 0xeb, 0x63, 0x38}),
				address:          0,
			},
			want:    "ldrb\tw11, [fp, x3, sxtx]",
			wantErr: false,
		},
		{
			name: "strb	w12, [x28, xzr, sxtx #0]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0xfb, 0x3f, 0x38}),
				address:          0,
			},
			want:    "strb\tw12, [x28, xzr, sxtx #0]",
			wantErr: false,
		},
		{
			name: "ldrb	w14, [x26, w6, uxtw]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4e, 0x4b, 0x66, 0x38}),
				address:          0,
			},
			want:    "ldrb\tw14, [x26, w6, uxtw]",
			wantErr: false,
		},
		{
			name: "ldrsb	w15, [x25, w7, uxtw #0]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2f, 0x5b, 0xe7, 0x38}),
				address:          0,
			},
			want:    "ldrsb\tw15, [x25, w7, uxtw #0]",
			wantErr: false,
		},
		{
			name: "ldrb	w17, [x23, w9, sxtw]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf1, 0xca, 0x69, 0x38}),
				address:          0,
			},
			want:    "ldrb\tw17, [x23, w9, sxtw]",
			wantErr: false,
		},
		{
			name: "ldrsb	x18, [x22, w10, sxtw #0]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd2, 0xda, 0xaa, 0x38}),
				address:          0,
			},
			want:    "ldrsb\tx18, [x22, w10, sxtw #0]",
			wantErr: false,
		},
		{
			name: "ldrsh	w3, [sp, x5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x6b, 0xe5, 0x78}),
				address:          0,
			},
			want:    "ldrsh\tw3, [sp, x5]",
			wantErr: false,
		},
		{
			name: "ldrsh	w9, [x27, x6]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0x6b, 0xe6, 0x78}),
				address:          0,
			},
			want:    "ldrsh\tw9, [x27, x6]",
			wantErr: false,
		},
		{
			name: "ldrh	w10, [x30, x7, lsl #1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xca, 0x7b, 0x67, 0x78}),
				address:          0,
			},
			want:    "ldrh\tw10, [lr, x7, lsl #0x1]",
			wantErr: false,
		},
		{
			name: "strh	w11, [x29, x3, sxtx]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xab, 0xeb, 0x23, 0x78}),
				address:          0,
			},
			want:    "strh\tw11, [fp, x3, sxtx]",
			wantErr: false,
		},
		{
			name: "ldrh	w12, [x28, xzr, sxtx]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0xeb, 0x7f, 0x78}),
				address:          0,
			},
			want:    "ldrh\tw12, [x28, xzr, sxtx]",
			wantErr: false,
		},
		{
			name: "ldrsh	x13, [x27, x5, sxtx #1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6d, 0xfb, 0xa5, 0x78}),
				address:          0,
			},
			want:    "ldrsh\tx13, [x27, x5, sxtx #0x1]",
			wantErr: false,
		},
		{
			name: "ldrh	w14, [x26, w6, uxtw]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4e, 0x4b, 0x66, 0x78}),
				address:          0,
			},
			want:    "ldrh\tw14, [x26, w6, uxtw]",
			wantErr: false,
		},
		{
			name: "ldrh	w15, [x25, w7, uxtw]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2f, 0x4b, 0x67, 0x78}),
				address:          0,
			},
			want:    "ldrh\tw15, [x25, w7, uxtw]",
			wantErr: false,
		},
		{
			name: "ldrsh	w16, [x24, w8, uxtw #1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x10, 0x5b, 0xe8, 0x78}),
				address:          0,
			},
			want:    "ldrsh\tw16, [x24, w8, uxtw #0x1]",
			wantErr: false,
		},
		{
			name: "ldrh	w17, [x23, w9, sxtw]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf1, 0xca, 0x69, 0x78}),
				address:          0,
			},
			want:    "ldrh\tw17, [x23, w9, sxtw]",
			wantErr: false,
		},
		{
			name: "ldrh	w18, [x22, w10, sxtw]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd2, 0xca, 0x6a, 0x78}),
				address:          0,
			},
			want:    "ldrh\tw18, [x22, w10, sxtw]",
			wantErr: false,
		},
		{
			name: "strh	w19, [x21, wzr, sxtw #1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb3, 0xda, 0x3f, 0x78}),
				address:          0,
			},
			want:    "strh\tw19, [x21, wzr, sxtw #0x1]",
			wantErr: false,
		},
		{
			name: "ldr	w3, [sp, x5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x6b, 0x65, 0xb8}),
				address:          0,
			},
			want:    "ldr\tw3, [sp, x5]",
			wantErr: false,
		},
		{
			name: "ldr	s9, [x27, x6]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0x6b, 0x66, 0xbc}),
				address:          0,
			},
			want:    "ldr\ts9, [x27, x6]",
			wantErr: false,
		},
		{
			name: "ldr	w10, [x30, x7, lsl #2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xca, 0x7b, 0x67, 0xb8}),
				address:          0,
			},
			want:    "ldr\tw10, [lr, x7, lsl #0x2]",
			wantErr: false,
		},
		{
			name: "ldr	w11, [x29, x3, sxtx]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xab, 0xeb, 0x63, 0xb8}),
				address:          0,
			},
			want:    "ldr\tw11, [fp, x3, sxtx]",
			wantErr: false,
		},
		{
			name: "str	s12, [x28, xzr, sxtx]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0xeb, 0x3f, 0xbc}),
				address:          0,
			},
			want:    "str\ts12, [x28, xzr, sxtx]",
			wantErr: false,
		},
		{
			name: "str	w13, [x27, x5, sxtx #2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6d, 0xfb, 0x25, 0xb8}),
				address:          0,
			},
			want:    "str\tw13, [x27, x5, sxtx #0x2]",
			wantErr: false,
		},
		{
			name: "str	w14, [x26, w6, uxtw]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4e, 0x4b, 0x26, 0xb8}),
				address:          0,
			},
			want:    "str\tw14, [x26, w6, uxtw]",
			wantErr: false,
		},
		{
			name: "ldr	w15, [x25, w7, uxtw]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2f, 0x4b, 0x67, 0xb8}),
				address:          0,
			},
			want:    "ldr\tw15, [x25, w7, uxtw]",
			wantErr: false,
		},
		{
			name: "ldr	w16, [x24, w8, uxtw #2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x10, 0x5b, 0x68, 0xb8}),
				address:          0,
			},
			want:    "ldr\tw16, [x24, w8, uxtw #0x2]",
			wantErr: false,
		},
		{
			name: "ldrsw	x17, [x23, w9, sxtw]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf1, 0xca, 0xa9, 0xb8}),
				address:          0,
			},
			want:    "ldrsw\tx17, [x23, w9, sxtw]",
			wantErr: false,
		},
		{
			name: "ldr	w18, [x22, w10, sxtw]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd2, 0xca, 0x6a, 0xb8}),
				address:          0,
			},
			want:    "ldr\tw18, [x22, w10, sxtw]",
			wantErr: false,
		},
		{
			name: "ldrsw	x19, [x21, wzr, sxtw #2]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb3, 0xda, 0xbf, 0xb8}),
				address:          0,
			},
			want:    "ldrsw\tx19, [x21, wzr, sxtw #0x2]",
			wantErr: false,
		},
		{
			name: "ldr	x3, [sp, x5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x6b, 0x65, 0xf8}),
				address:          0,
			},
			want:    "ldr\tx3, [sp, x5]",
			wantErr: false,
		},
		{
			name: "str	x9, [x27, x6]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0x6b, 0x26, 0xf8}),
				address:          0,
			},
			want:    "str\tx9, [x27, x6]",
			wantErr: false,
		},
		{
			name: "ldr	d10, [x30, x7, lsl #3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xca, 0x7b, 0x67, 0xfc}),
				address:          0,
			},
			want:    "ldr\td10, [lr, x7, lsl #0x3]",
			wantErr: false,
		},
		{
			name: "str	x11, [x29, x3, sxtx]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xab, 0xeb, 0x23, 0xf8}),
				address:          0,
			},
			want:    "str\tx11, [fp, x3, sxtx]",
			wantErr: false,
		},
		{
			name: "ldr	x12, [x28, xzr, sxtx]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0xeb, 0x7f, 0xf8}),
				address:          0,
			},
			want:    "ldr\tx12, [x28, xzr, sxtx]",
			wantErr: false,
		},
		{
			name: "ldr	x13, [x27, x5, sxtx #3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6d, 0xfb, 0x65, 0xf8}),
				address:          0,
			},
			want:    "ldr\tx13, [x27, x5, sxtx #0x3]",
			wantErr: false,
		},
		{
			name: "prfm	pldl1keep, [x26, w6, uxtw]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0x4b, 0xa6, 0xf8}),
				address:          0,
			},
			want:    "prfm\tpldl1keep, [x26, w6, uxtw]",
			wantErr: false,
		},
		{
			name: "ldr	x15, [x25, w7, uxtw]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2f, 0x4b, 0x67, 0xf8}),
				address:          0,
			},
			want:    "ldr\tx15, [x25, w7, uxtw]",
			wantErr: false,
		},
		{
			name: "ldr	x16, [x24, w8, uxtw #3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x10, 0x5b, 0x68, 0xf8}),
				address:          0,
			},
			want:    "ldr\tx16, [x24, w8, uxtw #0x3]",
			wantErr: false,
		},
		{
			name: "ldr	x17, [x23, w9, sxtw]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf1, 0xca, 0x69, 0xf8}),
				address:          0,
			},
			want:    "ldr\tx17, [x23, w9, sxtw]",
			wantErr: false,
		},
		{
			name: "ldr	x18, [x22, w10, sxtw]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd2, 0xca, 0x6a, 0xf8}),
				address:          0,
			},
			want:    "ldr\tx18, [x22, w10, sxtw]",
			wantErr: false,
		},
		{
			name: "str	d19, [x21, wzr, sxtw #3]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb3, 0xda, 0x3f, 0xfc}),
				address:          0,
			},
			want:    "str\td19, [x21, wzr, sxtw #0x3]",
			wantErr: false,
		},
		{
			name: "prfm	#6, [x0, x5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x06, 0x68, 0xa5, 0xf8}),
				address:          0,
			},
			want:    "prfm\t#6, [x0, x5]",
			wantErr: false,
		},
		{
			name: "ldr	q3, [sp, x5]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x6b, 0xe5, 0x3c}),
				address:          0,
			},
			want:    "ldr\tq3, [sp, x5]",
			wantErr: false,
		},
		{
			name: "ldr	q9, [x27, x6]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0x6b, 0xe6, 0x3c}),
				address:          0,
			},
			want:    "ldr\tq9, [x27, x6]",
			wantErr: false,
		},
		{
			name: "ldr	q10, [x30, x7, lsl #4]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xca, 0x7b, 0xe7, 0x3c}),
				address:          0,
			},
			want:    "ldr\tq10, [lr, x7, lsl #0x4]",
			wantErr: false,
		},
		{
			name: "str	q11, [x29, x3, sxtx]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xab, 0xeb, 0xa3, 0x3c}),
				address:          0,
			},
			want:    "str\tq11, [fp, x3, sxtx]",
			wantErr: false,
		},
		{
			name: "str	q12, [x28, xzr, sxtx]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0xeb, 0xbf, 0x3c}),
				address:          0,
			},
			want:    "str\tq12, [x28, xzr, sxtx]",
			wantErr: false,
		},
		{
			name: "str	q13, [x27, x5, sxtx #4]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6d, 0xfb, 0xa5, 0x3c}),
				address:          0,
			},
			want:    "str\tq13, [x27, x5, sxtx #0x4]",
			wantErr: false,
		},
		{
			name: "ldr	q14, [x26, w6, uxtw]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4e, 0x4b, 0xe6, 0x3c}),
				address:          0,
			},
			want:    "ldr\tq14, [x26, w6, uxtw]",
			wantErr: false,
		},
		{
			name: "ldr	q15, [x25, w7, uxtw]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2f, 0x4b, 0xe7, 0x3c}),
				address:          0,
			},
			want:    "ldr\tq15, [x25, w7, uxtw]",
			wantErr: false,
		},
		{
			name: "ldr	q16, [x24, w8, uxtw #4]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x10, 0x5b, 0xe8, 0x3c}),
				address:          0,
			},
			want:    "ldr\tq16, [x24, w8, uxtw #0x4]",
			wantErr: false,
		},
		{
			name: "ldr	q17, [x23, w9, sxtw]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf1, 0xca, 0xe9, 0x3c}),
				address:          0,
			},
			want:    "ldr\tq17, [x23, w9, sxtw]",
			wantErr: false,
		},
		{
			name: "str	q18, [x22, w10, sxtw]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd2, 0xca, 0xaa, 0x3c}),
				address:          0,
			},
			want:    "str\tq18, [x22, w10, sxtw]",
			wantErr: false,
		},
		{
			name: "ldr	q19, [x21, wzr, sxtw #4]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb3, 0xda, 0xff, 0x3c}),
				address:          0,
			},
			want:    "ldr\tq19, [x21, wzr, sxtw #0x4]",
			wantErr: false,
		},
		{
			name: "strb	w9, [x2], #255",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xf4, 0x0f, 0x38}),
				address:          0,
			},
			want:    "strb\tw9, [x2], #0xff",
			wantErr: false,
		},
		{
			name: "strb	w10, [x3], #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6a, 0x14, 0x00, 0x38}),
				address:          0,
			},
			want:    "strb\tw10, [x3], #0x1",
			wantErr: false,
		},
		{
			name: "strb	w10, [x3], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6a, 0x04, 0x10, 0x38}),
				address:          0,
			},
			want:    "strb\tw10, [x3], #-0x100",
			wantErr: false,
		},
		{
			name: "strh	w9, [x2], #255",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xf4, 0x0f, 0x78}),
				address:          0,
			},
			want:    "strh\tw9, [x2], #0xff",
			wantErr: false,
		},
		{
			name: "strh	w9, [x2], #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x14, 0x00, 0x78}),
				address:          0,
			},
			want:    "strh\tw9, [x2], #0x1",
			wantErr: false,
		},
		{
			name: "strh	w10, [x3], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6a, 0x04, 0x10, 0x78}),
				address:          0,
			},
			want:    "strh\tw10, [x3], #-0x100",
			wantErr: false,
		},
		{
			name: "str	w19, [sp], #255",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf3, 0xf7, 0x0f, 0xb8}),
				address:          0,
			},
			want:    "str\tw19, [sp], #0xff",
			wantErr: false,
		},
		{
			name: "str	w20, [x30], #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd4, 0x17, 0x00, 0xb8}),
				address:          0,
			},
			want:    "str\tw20, [lr], #0x1",
			wantErr: false,
		},
		{
			name: "str	w21, [x12], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x95, 0x05, 0x10, 0xb8}),
				address:          0,
			},
			want:    "str\tw21, [x12], #-0x100",
			wantErr: false,
		},
		{
			name: "str	xzr, [x9], #255",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0xf5, 0x0f, 0xf8}),
				address:          0,
			},
			want:    "str\txzr, [x9], #0xff",
			wantErr: false,
		},
		{
			name: "str	x2, [x3], #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x14, 0x00, 0xf8}),
				address:          0,
			},
			want:    "str\tx2, [x3], #0x1",
			wantErr: false,
		},
		{
			name: "str	x19, [x12], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0x05, 0x10, 0xf8}),
				address:          0,
			},
			want:    "str\tx19, [x12], #-0x100",
			wantErr: false,
		},
		{
			name: "ldrb	w9, [x2], #255",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xf4, 0x4f, 0x38}),
				address:          0,
			},
			want:    "ldrb\tw9, [x2], #0xff",
			wantErr: false,
		},
		{
			name: "ldrb	w10, [x3], #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6a, 0x14, 0x40, 0x38}),
				address:          0,
			},
			want:    "ldrb\tw10, [x3], #0x1",
			wantErr: false,
		},
		{
			name: "ldrb	w10, [x3], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6a, 0x04, 0x50, 0x38}),
				address:          0,
			},
			want:    "ldrb\tw10, [x3], #-0x100",
			wantErr: false,
		},
		{
			name: "ldrh	w9, [x2], #255",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xf4, 0x4f, 0x78}),
				address:          0,
			},
			want:    "ldrh\tw9, [x2], #0xff",
			wantErr: false,
		},
		{
			name: "ldrh	w9, [x2], #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x14, 0x40, 0x78}),
				address:          0,
			},
			want:    "ldrh\tw9, [x2], #0x1",
			wantErr: false,
		},
		{
			name: "ldrh	w10, [x3], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6a, 0x04, 0x50, 0x78}),
				address:          0,
			},
			want:    "ldrh\tw10, [x3], #-0x100",
			wantErr: false,
		},
		{
			name: "ldr	w19, [sp], #255",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf3, 0xf7, 0x4f, 0xb8}),
				address:          0,
			},
			want:    "ldr\tw19, [sp], #0xff",
			wantErr: false,
		},
		{
			name: "ldr	w20, [x30], #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd4, 0x17, 0x40, 0xb8}),
				address:          0,
			},
			want:    "ldr\tw20, [lr], #0x1",
			wantErr: false,
		},
		{
			name: "ldr	w21, [x12], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x95, 0x05, 0x50, 0xb8}),
				address:          0,
			},
			want:    "ldr\tw21, [x12], #-0x100",
			wantErr: false,
		},
		{
			name: "ldr	xzr, [x9], #255",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0xf5, 0x4f, 0xf8}),
				address:          0,
			},
			want:    "ldr\txzr, [x9], #0xff",
			wantErr: false,
		},
		{
			name: "ldr	x2, [x3], #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x14, 0x40, 0xf8}),
				address:          0,
			},
			want:    "ldr\tx2, [x3], #0x1",
			wantErr: false,
		},
		{
			name: "ldr	x19, [x12], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0x05, 0x50, 0xf8}),
				address:          0,
			},
			want:    "ldr\tx19, [x12], #-0x100",
			wantErr: false,
		},
		{
			name: "ldrsb	xzr, [x9], #255",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0xf5, 0x8f, 0x38}),
				address:          0,
			},
			want:    "ldrsb\txzr, [x9], #0xff",
			wantErr: false,
		},
		{
			name: "ldrsb	x2, [x3], #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x14, 0x80, 0x38}),
				address:          0,
			},
			want:    "ldrsb\tx2, [x3], #0x1",
			wantErr: false,
		},
		{
			name: "ldrsb	x19, [x12], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0x05, 0x90, 0x38}),
				address:          0,
			},
			want:    "ldrsb\tx19, [x12], #-0x100",
			wantErr: false,
		},
		{
			name: "ldrsh	xzr, [x9], #255",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0xf5, 0x8f, 0x78}),
				address:          0,
			},
			want:    "ldrsh\txzr, [x9], #0xff",
			wantErr: false,
		},
		{
			name: "ldrsh	x2, [x3], #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x14, 0x80, 0x78}),
				address:          0,
			},
			want:    "ldrsh\tx2, [x3], #0x1",
			wantErr: false,
		},
		{
			name: "ldrsh	x19, [x12], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0x05, 0x90, 0x78}),
				address:          0,
			},
			want:    "ldrsh\tx19, [x12], #-0x100",
			wantErr: false,
		},
		{
			name: "ldrsw	xzr, [x9], #255",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0xf5, 0x8f, 0xb8}),
				address:          0,
			},
			want:    "ldrsw\txzr, [x9], #0xff",
			wantErr: false,
		},
		{
			name: "ldrsw	x2, [x3], #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x14, 0x80, 0xb8}),
				address:          0,
			},
			want:    "ldrsw\tx2, [x3], #0x1",
			wantErr: false,
		},
		{
			name: "ldrsw	x19, [x12], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0x05, 0x90, 0xb8}),
				address:          0,
			},
			want:    "ldrsw\tx19, [x12], #-0x100",
			wantErr: false,
		},
		{
			name: "ldrsb	wzr, [x9], #255",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0xf5, 0xcf, 0x38}),
				address:          0,
			},
			want:    "ldrsb\twzr, [x9], #0xff",
			wantErr: false,
		},
		{
			name: "ldrsb	w2, [x3], #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x14, 0xc0, 0x38}),
				address:          0,
			},
			want:    "ldrsb\tw2, [x3], #0x1",
			wantErr: false,
		},
		{
			name: "ldrsb	w19, [x12], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0x05, 0xd0, 0x38}),
				address:          0,
			},
			want:    "ldrsb\tw19, [x12], #-0x100",
			wantErr: false,
		},
		{
			name: "ldrsh	wzr, [x9], #255",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0xf5, 0xcf, 0x78}),
				address:          0,
			},
			want:    "ldrsh\twzr, [x9], #0xff",
			wantErr: false,
		},
		{
			name: "ldrsh	w2, [x3], #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x14, 0xc0, 0x78}),
				address:          0,
			},
			want:    "ldrsh\tw2, [x3], #0x1",
			wantErr: false,
		},
		{
			name: "ldrsh	w19, [x12], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0x05, 0xd0, 0x78}),
				address:          0,
			},
			want:    "ldrsh\tw19, [x12], #-0x100",
			wantErr: false,
		},
		{
			name: "str	b0, [x0], #255",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0xf4, 0x0f, 0x3c}),
				address:          0,
			},
			want:    "str\tb0, [x0], #0xff",
			wantErr: false,
		},
		{
			name: "str	b3, [x3], #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x63, 0x14, 0x00, 0x3c}),
				address:          0,
			},
			want:    "str\tb3, [x3], #0x1",
			wantErr: false,
		},
		{
			name: "str	b5, [sp], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe5, 0x07, 0x10, 0x3c}),
				address:          0,
			},
			want:    "str\tb5, [sp], #-0x100",
			wantErr: false,
		},
		{
			name: "str	h10, [x10], #255",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4a, 0xf5, 0x0f, 0x7c}),
				address:          0,
			},
			want:    "str\th10, [x10], #0xff",
			wantErr: false,
		},
		{
			name: "str	h13, [x23], #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xed, 0x16, 0x00, 0x7c}),
				address:          0,
			},
			want:    "str\th13, [x23], #0x1",
			wantErr: false,
		},
		{
			name: "str	h15, [sp], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xef, 0x07, 0x10, 0x7c}),
				address:          0,
			},
			want:    "str\th15, [sp], #-0x100",
			wantErr: false,
		},
		{
			name: "str	s20, [x20], #255",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x94, 0xf6, 0x0f, 0xbc}),
				address:          0,
			},
			want:    "str\ts20, [x20], #0xff",
			wantErr: false,
		},
		{
			name: "str	s23, [x23], #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf7, 0x16, 0x00, 0xbc}),
				address:          0,
			},
			want:    "str\ts23, [x23], #0x1",
			wantErr: false,
		},
		{
			name: "str	s25, [x0], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x19, 0x04, 0x10, 0xbc}),
				address:          0,
			},
			want:    "str\ts25, [x0], #-0x100",
			wantErr: false,
		},
		{
			name: "str	d20, [x20], #255",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x94, 0xf6, 0x0f, 0xfc}),
				address:          0,
			},
			want:    "str\td20, [x20], #0xff",
			wantErr: false,
		},
		{
			name: "str	d23, [x23], #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf7, 0x16, 0x00, 0xfc}),
				address:          0,
			},
			want:    "str\td23, [x23], #0x1",
			wantErr: false,
		},
		{
			name: "str	d25, [x0], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x19, 0x04, 0x10, 0xfc}),
				address:          0,
			},
			want:    "str\td25, [x0], #-0x100",
			wantErr: false,
		},
		{
			name: "ldr	b0, [x0], #255",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0xf4, 0x4f, 0x3c}),
				address:          0,
			},
			want:    "ldr\tb0, [x0], #0xff",
			wantErr: false,
		},
		{
			name: "ldr	b3, [x3], #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x63, 0x14, 0x40, 0x3c}),
				address:          0,
			},
			want:    "ldr\tb3, [x3], #0x1",
			wantErr: false,
		},
		{
			name: "ldr	b5, [sp], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe5, 0x07, 0x50, 0x3c}),
				address:          0,
			},
			want:    "ldr\tb5, [sp], #-0x100",
			wantErr: false,
		},
		{
			name: "ldr	h10, [x10], #255",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4a, 0xf5, 0x4f, 0x7c}),
				address:          0,
			},
			want:    "ldr\th10, [x10], #0xff",
			wantErr: false,
		},
		{
			name: "ldr	h13, [x23], #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xed, 0x16, 0x40, 0x7c}),
				address:          0,
			},
			want:    "ldr\th13, [x23], #0x1",
			wantErr: false,
		},
		{
			name: "ldr	h15, [sp], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xef, 0x07, 0x50, 0x7c}),
				address:          0,
			},
			want:    "ldr\th15, [sp], #-0x100",
			wantErr: false,
		},
		{
			name: "ldr	s20, [x20], #255",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x94, 0xf6, 0x4f, 0xbc}),
				address:          0,
			},
			want:    "ldr\ts20, [x20], #0xff",
			wantErr: false,
		},
		{
			name: "ldr	s23, [x23], #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf7, 0x16, 0x40, 0xbc}),
				address:          0,
			},
			want:    "ldr\ts23, [x23], #0x1",
			wantErr: false,
		},
		{
			name: "ldr	s25, [x0], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x19, 0x04, 0x50, 0xbc}),
				address:          0,
			},
			want:    "ldr\ts25, [x0], #-0x100",
			wantErr: false,
		},
		{
			name: "ldr	d20, [x20], #255",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x94, 0xf6, 0x4f, 0xfc}),
				address:          0,
			},
			want:    "ldr\td20, [x20], #0xff",
			wantErr: false,
		},
		{
			name: "ldr	d23, [x23], #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf7, 0x16, 0x40, 0xfc}),
				address:          0,
			},
			want:    "ldr\td23, [x23], #0x1",
			wantErr: false,
		},
		{
			name: "ldr	d25, [x0], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x19, 0x04, 0x50, 0xfc}),
				address:          0,
			},
			want:    "ldr\td25, [x0], #-0x100",
			wantErr: false,
		},
		{
			name: "ldr	q20, [x1], #255",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x34, 0xf4, 0xcf, 0x3c}),
				address:          0,
			},
			want:    "ldr\tq20, [x1], #0xff",
			wantErr: false,
		},
		{
			name: "ldr	q23, [x9], #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x37, 0x15, 0xc0, 0x3c}),
				address:          0,
			},
			want:    "ldr\tq23, [x9], #0x1",
			wantErr: false,
		},
		{
			name: "ldr	q25, [x20], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x99, 0x06, 0xd0, 0x3c}),
				address:          0,
			},
			want:    "ldr\tq25, [x20], #-0x100",
			wantErr: false,
		},
		{
			name: "str	q10, [x1], #255",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2a, 0xf4, 0x8f, 0x3c}),
				address:          0,
			},
			want:    "str\tq10, [x1], #0xff",
			wantErr: false,
		},
		{
			name: "str	q22, [sp], #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf6, 0x17, 0x80, 0x3c}),
				address:          0,
			},
			want:    "str\tq22, [sp], #0x1",
			wantErr: false,
		},
		{
			name: "str	q21, [x20], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x95, 0x06, 0x90, 0x3c}),
				address:          0,
			},
			want:    "str\tq21, [x20], #-0x100",
			wantErr: false,
		},
		{
			name: "ldr	x3, [x4, #0]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x83, 0x0c, 0x40, 0xf8}),
				address:          0,
			},
			want:    "ldr\tx3, [x4, #0]!",
			wantErr: false,
		},
		{
			name: "ldr	xzr, [sp, #0]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x0f, 0x40, 0xf8}),
				address:          0,
			},
			want:    "ldr\txzr, [sp, #0]!",
			wantErr: false,
		},
		{
			name: "strb	w9, [x2, #255]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xfc, 0x0f, 0x38}),
				address:          0,
			},
			want:    "strb\tw9, [x2, #0xff]!",
			wantErr: false,
		},
		{
			name: "strb	w10, [x3, #1]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6a, 0x1c, 0x00, 0x38}),
				address:          0,
			},
			want:    "strb\tw10, [x3, #0x1]!",
			wantErr: false,
		},
		{
			name: "strb	w10, [x3, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6a, 0x0c, 0x10, 0x38}),
				address:          0,
			},
			want:    "strb\tw10, [x3, #-0x100]!",
			wantErr: false,
		},
		{
			name: "strh	w9, [x2, #255]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xfc, 0x0f, 0x78}),
				address:          0,
			},
			want:    "strh\tw9, [x2, #0xff]!",
			wantErr: false,
		},
		{
			name: "strh	w9, [x2, #1]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x1c, 0x00, 0x78}),
				address:          0,
			},
			want:    "strh\tw9, [x2, #0x1]!",
			wantErr: false,
		},
		{
			name: "strh	w10, [x3, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6a, 0x0c, 0x10, 0x78}),
				address:          0,
			},
			want:    "strh\tw10, [x3, #-0x100]!",
			wantErr: false,
		},
		{
			name: "str	w19, [sp, #255]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf3, 0xff, 0x0f, 0xb8}),
				address:          0,
			},
			want:    "str\tw19, [sp, #0xff]!",
			wantErr: false,
		},
		{
			name: "str	w20, [x30, #1]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd4, 0x1f, 0x00, 0xb8}),
				address:          0,
			},
			want:    "str\tw20, [lr, #0x1]!",
			wantErr: false,
		},
		{
			name: "str	w21, [x12, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x95, 0x0d, 0x10, 0xb8}),
				address:          0,
			},
			want:    "str\tw21, [x12, #-0x100]!",
			wantErr: false,
		},
		{
			name: "str	xzr, [x9, #255]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0xfd, 0x0f, 0xf8}),
				address:          0,
			},
			want:    "str\txzr, [x9, #0xff]!",
			wantErr: false,
		},
		{
			name: "str	x2, [x3, #1]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x1c, 0x00, 0xf8}),
				address:          0,
			},
			want:    "str\tx2, [x3, #0x1]!",
			wantErr: false,
		},
		{
			name: "str	x19, [x12, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0x0d, 0x10, 0xf8}),
				address:          0,
			},
			want:    "str\tx19, [x12, #-0x100]!",
			wantErr: false,
		},
		{
			name: "ldrb	w9, [x2, #255]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xfc, 0x4f, 0x38}),
				address:          0,
			},
			want:    "ldrb\tw9, [x2, #0xff]!",
			wantErr: false,
		},
		{
			name: "ldrb	w10, [x3, #1]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6a, 0x1c, 0x40, 0x38}),
				address:          0,
			},
			want:    "ldrb\tw10, [x3, #0x1]!",
			wantErr: false,
		},
		{
			name: "ldrb	w10, [x3, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6a, 0x0c, 0x50, 0x38}),
				address:          0,
			},
			want:    "ldrb\tw10, [x3, #-0x100]!",
			wantErr: false,
		},
		{
			name: "ldrh	w9, [x2, #255]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xfc, 0x4f, 0x78}),
				address:          0,
			},
			want:    "ldrh\tw9, [x2, #0xff]!",
			wantErr: false,
		},
		{
			name: "ldrh	w9, [x2, #1]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x1c, 0x40, 0x78}),
				address:          0,
			},
			want:    "ldrh\tw9, [x2, #0x1]!",
			wantErr: false,
		},
		{
			name: "ldrh	w10, [x3, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6a, 0x0c, 0x50, 0x78}),
				address:          0,
			},
			want:    "ldrh\tw10, [x3, #-0x100]!",
			wantErr: false,
		},
		{
			name: "ldr	w19, [sp, #255]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf3, 0xff, 0x4f, 0xb8}),
				address:          0,
			},
			want:    "ldr\tw19, [sp, #0xff]!",
			wantErr: false,
		},
		{
			name: "ldr	w20, [x30, #1]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd4, 0x1f, 0x40, 0xb8}),
				address:          0,
			},
			want:    "ldr\tw20, [lr, #0x1]!",
			wantErr: false,
		},
		{
			name: "ldr	w21, [x12, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x95, 0x0d, 0x50, 0xb8}),
				address:          0,
			},
			want:    "ldr\tw21, [x12, #-0x100]!",
			wantErr: false,
		},
		{
			name: "ldr	xzr, [x9, #255]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0xfd, 0x4f, 0xf8}),
				address:          0,
			},
			want:    "ldr\txzr, [x9, #0xff]!",
			wantErr: false,
		},
		{
			name: "ldr	x2, [x3, #1]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x1c, 0x40, 0xf8}),
				address:          0,
			},
			want:    "ldr\tx2, [x3, #0x1]!",
			wantErr: false,
		},
		{
			name: "ldr	x19, [x12, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0x0d, 0x50, 0xf8}),
				address:          0,
			},
			want:    "ldr\tx19, [x12, #-0x100]!",
			wantErr: false,
		},
		{
			name: "ldrsb	xzr, [x9, #255]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0xfd, 0x8f, 0x38}),
				address:          0,
			},
			want:    "ldrsb\txzr, [x9, #0xff]!",
			wantErr: false,
		},
		{
			name: "ldrsb	x2, [x3, #1]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x1c, 0x80, 0x38}),
				address:          0,
			},
			want:    "ldrsb\tx2, [x3, #0x1]!",
			wantErr: false,
		},
		{
			name: "ldrsb	x19, [x12, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0x0d, 0x90, 0x38}),
				address:          0,
			},
			want:    "ldrsb\tx19, [x12, #-0x100]!",
			wantErr: false,
		},
		{
			name: "ldrsh	xzr, [x9, #255]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0xfd, 0x8f, 0x78}),
				address:          0,
			},
			want:    "ldrsh\txzr, [x9, #0xff]!",
			wantErr: false,
		},
		{
			name: "ldrsh	x2, [x3, #1]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x1c, 0x80, 0x78}),
				address:          0,
			},
			want:    "ldrsh\tx2, [x3, #0x1]!",
			wantErr: false,
		},
		{
			name: "ldrsh	x19, [x12, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0x0d, 0x90, 0x78}),
				address:          0,
			},
			want:    "ldrsh\tx19, [x12, #-0x100]!",
			wantErr: false,
		},
		{
			name: "ldrsw	xzr, [x9, #255]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0xfd, 0x8f, 0xb8}),
				address:          0,
			},
			want:    "ldrsw\txzr, [x9, #0xff]!",
			wantErr: false,
		},
		{
			name: "ldrsw	x2, [x3, #1]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x1c, 0x80, 0xb8}),
				address:          0,
			},
			want:    "ldrsw\tx2, [x3, #0x1]!",
			wantErr: false,
		},
		{
			name: "ldrsw	x19, [x12, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0x0d, 0x90, 0xb8}),
				address:          0,
			},
			want:    "ldrsw\tx19, [x12, #-0x100]!",
			wantErr: false,
		},
		{
			name: "ldrsb	wzr, [x9, #255]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0xfd, 0xcf, 0x38}),
				address:          0,
			},
			want:    "ldrsb\twzr, [x9, #0xff]!",
			wantErr: false,
		},
		{
			name: "ldrsb	w2, [x3, #1]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x1c, 0xc0, 0x38}),
				address:          0,
			},
			want:    "ldrsb\tw2, [x3, #0x1]!",
			wantErr: false,
		},
		{
			name: "ldrsb	w19, [x12, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0x0d, 0xd0, 0x38}),
				address:          0,
			},
			want:    "ldrsb\tw19, [x12, #-0x100]!",
			wantErr: false,
		},
		{
			name: "ldrsh	wzr, [x9, #255]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0xfd, 0xcf, 0x78}),
				address:          0,
			},
			want:    "ldrsh\twzr, [x9, #0xff]!",
			wantErr: false,
		},
		{
			name: "ldrsh	w2, [x3, #1]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x1c, 0xc0, 0x78}),
				address:          0,
			},
			want:    "ldrsh\tw2, [x3, #0x1]!",
			wantErr: false,
		},
		{
			name: "ldrsh	w19, [x12, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0x0d, 0xd0, 0x78}),
				address:          0,
			},
			want:    "ldrsh\tw19, [x12, #-0x100]!",
			wantErr: false,
		},
		{
			name: "str	b0, [x0, #255]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0xfc, 0x0f, 0x3c}),
				address:          0,
			},
			want:    "str\tb0, [x0, #0xff]!",
			wantErr: false,
		},
		{
			name: "str	b3, [x3, #1]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x63, 0x1c, 0x00, 0x3c}),
				address:          0,
			},
			want:    "str\tb3, [x3, #0x1]!",
			wantErr: false,
		},
		{
			name: "str	b5, [sp, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe5, 0x0f, 0x10, 0x3c}),
				address:          0,
			},
			want:    "str\tb5, [sp, #-0x100]!",
			wantErr: false,
		},
		{
			name: "str	h10, [x10, #255]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4a, 0xfd, 0x0f, 0x7c}),
				address:          0,
			},
			want:    "str\th10, [x10, #0xff]!",
			wantErr: false,
		},
		{
			name: "str	h13, [x23, #1]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xed, 0x1e, 0x00, 0x7c}),
				address:          0,
			},
			want:    "str\th13, [x23, #0x1]!",
			wantErr: false,
		},
		{
			name: "str	h15, [sp, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xef, 0x0f, 0x10, 0x7c}),
				address:          0,
			},
			want:    "str\th15, [sp, #-0x100]!",
			wantErr: false,
		},
		{
			name: "str	s20, [x20, #255]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x94, 0xfe, 0x0f, 0xbc}),
				address:          0,
			},
			want:    "str\ts20, [x20, #0xff]!",
			wantErr: false,
		},
		{
			name: "str	s23, [x23, #1]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf7, 0x1e, 0x00, 0xbc}),
				address:          0,
			},
			want:    "str\ts23, [x23, #0x1]!",
			wantErr: false,
		},
		{
			name: "str	s25, [x0, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x19, 0x0c, 0x10, 0xbc}),
				address:          0,
			},
			want:    "str\ts25, [x0, #-0x100]!",
			wantErr: false,
		},
		{
			name: "str	d20, [x20, #255]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x94, 0xfe, 0x0f, 0xfc}),
				address:          0,
			},
			want:    "str\td20, [x20, #0xff]!",
			wantErr: false,
		},
		{
			name: "str	d23, [x23, #1]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf7, 0x1e, 0x00, 0xfc}),
				address:          0,
			},
			want:    "str\td23, [x23, #0x1]!",
			wantErr: false,
		},
		{
			name: "str	d25, [x0, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x19, 0x0c, 0x10, 0xfc}),
				address:          0,
			},
			want:    "str\td25, [x0, #-0x100]!",
			wantErr: false,
		},
		{
			name: "ldr	b0, [x0, #255]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0xfc, 0x4f, 0x3c}),
				address:          0,
			},
			want:    "ldr\tb0, [x0, #0xff]!",
			wantErr: false,
		},
		{
			name: "ldr	b3, [x3, #1]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x63, 0x1c, 0x40, 0x3c}),
				address:          0,
			},
			want:    "ldr\tb3, [x3, #0x1]!",
			wantErr: false,
		},
		{
			name: "ldr	b5, [sp, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe5, 0x0f, 0x50, 0x3c}),
				address:          0,
			},
			want:    "ldr\tb5, [sp, #-0x100]!",
			wantErr: false,
		},
		{
			name: "ldr	h10, [x10, #255]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4a, 0xfd, 0x4f, 0x7c}),
				address:          0,
			},
			want:    "ldr\th10, [x10, #0xff]!",
			wantErr: false,
		},
		{
			name: "ldr	h13, [x23, #1]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xed, 0x1e, 0x40, 0x7c}),
				address:          0,
			},
			want:    "ldr\th13, [x23, #0x1]!",
			wantErr: false,
		},
		{
			name: "ldr	h15, [sp, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xef, 0x0f, 0x50, 0x7c}),
				address:          0,
			},
			want:    "ldr\th15, [sp, #-0x100]!",
			wantErr: false,
		},
		{
			name: "ldr	s20, [x20, #255]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x94, 0xfe, 0x4f, 0xbc}),
				address:          0,
			},
			want:    "ldr\ts20, [x20, #0xff]!",
			wantErr: false,
		},
		{
			name: "ldr	s23, [x23, #1]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf7, 0x1e, 0x40, 0xbc}),
				address:          0,
			},
			want:    "ldr\ts23, [x23, #0x1]!",
			wantErr: false,
		},
		{
			name: "ldr	s25, [x0, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x19, 0x0c, 0x50, 0xbc}),
				address:          0,
			},
			want:    "ldr\ts25, [x0, #-0x100]!",
			wantErr: false,
		},
		{
			name: "ldr	d20, [x20, #255]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x94, 0xfe, 0x4f, 0xfc}),
				address:          0,
			},
			want:    "ldr\td20, [x20, #0xff]!",
			wantErr: false,
		},
		{
			name: "ldr	d23, [x23, #1]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf7, 0x1e, 0x40, 0xfc}),
				address:          0,
			},
			want:    "ldr\td23, [x23, #0x1]!",
			wantErr: false,
		},
		{
			name: "ldr	d25, [x0, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x19, 0x0c, 0x50, 0xfc}),
				address:          0,
			},
			want:    "ldr\td25, [x0, #-0x100]!",
			wantErr: false,
		},
		{
			name: "ldr	q20, [x1, #255]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x34, 0xfc, 0xcf, 0x3c}),
				address:          0,
			},
			want:    "ldr\tq20, [x1, #0xff]!",
			wantErr: false,
		},
		{
			name: "ldr	q23, [x9, #1]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x37, 0x1d, 0xc0, 0x3c}),
				address:          0,
			},
			want:    "ldr\tq23, [x9, #0x1]!",
			wantErr: false,
		},
		{
			name: "ldr	q25, [x20, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x99, 0x0e, 0xd0, 0x3c}),
				address:          0,
			},
			want:    "ldr\tq25, [x20, #-0x100]!",
			wantErr: false,
		},
		{
			name: "str	q10, [x1, #255]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2a, 0xfc, 0x8f, 0x3c}),
				address:          0,
			},
			want:    "str\tq10, [x1, #0xff]!",
			wantErr: false,
		},
		{
			name: "str	q22, [sp, #1]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf6, 0x1f, 0x80, 0x3c}),
				address:          0,
			},
			want:    "str\tq22, [sp, #0x1]!",
			wantErr: false,
		},
		{
			name: "str	q21, [x20, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x95, 0x0e, 0x90, 0x3c}),
				address:          0,
			},
			want:    "str\tq21, [x20, #-0x100]!",
			wantErr: false,
		},
		{
			name: "sttrb	w9, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x0b, 0x00, 0x38}),
				address:          0,
			},
			want:    "sttrb\tw9, [sp]",
			wantErr: false,
		},
		{
			name: "sttrh	wzr, [x12, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0xf9, 0x0f, 0x78}),
				address:          0,
			},
			want:    "sttrh\twzr, [x12, #0xff]",
			wantErr: false,
		},
		{
			name: "sttr	w16, [x0, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x10, 0x08, 0x10, 0xb8}),
				address:          0,
			},
			want:    "sttr\tw16, [x0, #-0x100]",
			wantErr: false,
		},
		{
			name: "sttr	x28, [x14, #1]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdc, 0x19, 0x00, 0xf8}),
				address:          0,
			},
			want:    "sttr\tx28, [x14, #0x1]",
			wantErr: false,
		},
		{
			name: "ldtrb	w1, [x20, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x81, 0xfa, 0x4f, 0x38}),
				address:          0,
			},
			want:    "ldtrb\tw1, [x20, #0xff]",
			wantErr: false,
		},
		{
			name: "ldtrh	w20, [x1, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x34, 0xf8, 0x4f, 0x78}),
				address:          0,
			},
			want:    "ldtrh\tw20, [x1, #0xff]",
			wantErr: false,
		},
		{
			name: "ldtr	w12, [sp, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0xfb, 0x4f, 0xb8}),
				address:          0,
			},
			want:    "ldtr\tw12, [sp, #0xff]",
			wantErr: false,
		},
		{
			name: "ldtr	xzr, [x12, #255]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0xf9, 0x4f, 0xf8}),
				address:          0,
			},
			want:    "ldtr\txzr, [x12, #0xff]",
			wantErr: false,
		},
		{
			name: "ldtrsb	x9, [x7, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x08, 0x90, 0x38}),
				address:          0,
			},
			want:    "ldtrsb\tx9, [x7, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldtrsh	x17, [x19, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x71, 0x0a, 0x90, 0x78}),
				address:          0,
			},
			want:    "ldtrsh\tx17, [x19, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldtrsw	x20, [x15, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x09, 0x90, 0xb8}),
				address:          0,
			},
			want:    "ldtrsw\tx20, [x15, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldtrsb	w19, [x1, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x33, 0x08, 0xd0, 0x38}),
				address:          0,
			},
			want:    "ldtrsb\tw19, [x1, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldtrsh	w15, [x21, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xaf, 0x0a, 0xd0, 0x78}),
				address:          0,
			},
			want:    "ldtrsh\tw15, [x21, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldp	w3, w5, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x17, 0x40, 0x29}),
				address:          0,
			},
			want:    "ldp\tw3, w5, [sp]",
			wantErr: false,
		},
		{
			name: "stp	wzr, w9, [sp, #252]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xa7, 0x1f, 0x29}),
				address:          0,
			},
			want:    "stp\twzr, w9, [sp, #0xfc]",
			wantErr: false,
		},
		{
			name: "ldp	w2, wzr, [sp, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe2, 0x7f, 0x60, 0x29}),
				address:          0,
			},
			want:    "ldp\tw2, wzr, [sp, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldp	w9, w10, [sp, #4]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xab, 0x40, 0x29}),
				address:          0,
			},
			want:    "ldp\tw9, w10, [sp, #0x4]",
			wantErr: false,
		},
		{
			name: "ldpsw	x9, x10, [sp, #4]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xab, 0x40, 0x69}),
				address:          0,
			},
			want:    "ldpsw\tx9, x10, [sp, #0x4]",
			wantErr: false,
		},
		{
			name: "ldpsw	x9, x10, [x2, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x28, 0x60, 0x69}),
				address:          0,
			},
			want:    "ldpsw\tx9, x10, [x2, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldpsw	x20, x30, [sp, #252]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0xfb, 0x5f, 0x69}),
				address:          0,
			},
			want:    "ldpsw\tx20, lr, [sp, #0xfc]",
			wantErr: false,
		},
		{
			name: "ldp	x21, x29, [x2, #504]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x55, 0xf4, 0x5f, 0xa9}),
				address:          0,
			},
			want:    "ldp\tx21, fp, [x2, #0x1f8]",
			wantErr: false,
		},
		{
			name: "ldp	x22, x23, [x3, #-512]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x76, 0x5c, 0x60, 0xa9}),
				address:          0,
			},
			want:    "ldp\tx22, x23, [x3, #-0x200]",
			wantErr: false,
		},
		{
			name: "ldp	x24, x25, [x4, #8]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x98, 0xe4, 0x40, 0xa9}),
				address:          0,
			},
			want:    "ldp\tx24, x25, [x4, #0x8]",
			wantErr: false,
		},
		{
			name: "ldp	s29, s28, [sp, #252]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfd, 0xf3, 0x5f, 0x2d}),
				address:          0,
			},
			want:    "ldp\ts29, s28, [sp, #0xfc]",
			wantErr: false,
		},
		{
			name: "stp	s27, s26, [sp, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfb, 0x6b, 0x20, 0x2d}),
				address:          0,
			},
			want:    "stp\ts27, s26, [sp, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldp	s1, s2, [x3, #44]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x61, 0x88, 0x45, 0x2d}),
				address:          0,
			},
			want:    "ldp\ts1, s2, [x3, #0x2c]",
			wantErr: false,
		},
		{
			name: "stp	d3, d5, [x9, #504]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x23, 0x95, 0x1f, 0x6d}),
				address:          0,
			},
			want:    "stp\td3, d5, [x9, #0x1f8]",
			wantErr: false,
		},
		{
			name: "stp	d7, d11, [x10, #-512]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x47, 0x2d, 0x20, 0x6d}),
				address:          0,
			},
			want:    "stp\td7, d11, [x10, #-0x200]",
			wantErr: false,
		},
		{
			name: "ldp	d2, d3, [x30, #-8]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc2, 0x8f, 0x7f, 0x6d}),
				address:          0,
			},
			want:    "ldp\td2, d3, [lr, #-0x8]",
			wantErr: false,
		},
		{
			name: "stp	q3, q5, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x17, 0x00, 0xad}),
				address:          0,
			},
			want:    "stp\tq3, q5, [sp]",
			wantErr: false,
		},
		{
			name: "stp	q17, q19, [sp, #1008]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf1, 0xcf, 0x1f, 0xad}),
				address:          0,
			},
			want:    "stp\tq17, q19, [sp, #0x3f0]",
			wantErr: false,
		},
		{
			name: "ldp	q23, q29, [x1, #-1024]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x37, 0x74, 0x60, 0xad}),
				address:          0,
			},
			want:    "ldp\tq23, q29, [x1, #-0x400]",
			wantErr: false,
		},
		{
			name: "ldp	w3, w5, [sp], #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x17, 0xc0, 0x28}),
				address:          0,
			},
			want:    "ldp\tw3, w5, [sp], #0",
			wantErr: false,
		},
		{
			name: "stp	wzr, w9, [sp], #252",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xa7, 0x9f, 0x28}),
				address:          0,
			},
			want:    "stp\twzr, w9, [sp], #0xfc",
			wantErr: false,
		},
		{
			name: "ldp	w2, wzr, [sp], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe2, 0x7f, 0xe0, 0x28}),
				address:          0,
			},
			want:    "ldp\tw2, wzr, [sp], #-0x100",
			wantErr: false,
		},
		{
			name: "ldp	w9, w10, [sp], #4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xab, 0xc0, 0x28}),
				address:          0,
			},
			want:    "ldp\tw9, w10, [sp], #0x4",
			wantErr: false,
		},
		{
			name: "ldpsw	x9, x10, [sp], #4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xab, 0xc0, 0x68}),
				address:          0,
			},
			want:    "ldpsw\tx9, x10, [sp], #0x4",
			wantErr: false,
		},
		{
			name: "ldpsw	x9, x10, [x2], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x28, 0xe0, 0x68}),
				address:          0,
			},
			want:    "ldpsw\tx9, x10, [x2], #-0x100",
			wantErr: false,
		},
		{
			name: "ldpsw	x20, x30, [sp], #252",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0xfb, 0xdf, 0x68}),
				address:          0,
			},
			want:    "ldpsw\tx20, lr, [sp], #0xfc",
			wantErr: false,
		},
		{
			name: "ldp	x21, x29, [x2], #504",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x55, 0xf4, 0xdf, 0xa8}),
				address:          0,
			},
			want:    "ldp\tx21, fp, [x2], #0x1f8",
			wantErr: false,
		},
		{
			name: "ldp	x22, x23, [x3], #-512",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x76, 0x5c, 0xe0, 0xa8}),
				address:          0,
			},
			want:    "ldp\tx22, x23, [x3], #-0x200",
			wantErr: false,
		},
		{
			name: "ldp	x24, x25, [x4], #8",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x98, 0xe4, 0xc0, 0xa8}),
				address:          0,
			},
			want:    "ldp\tx24, x25, [x4], #0x8",
			wantErr: false,
		},
		{
			name: "ldp	s29, s28, [sp], #252",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfd, 0xf3, 0xdf, 0x2c}),
				address:          0,
			},
			want:    "ldp\ts29, s28, [sp], #0xfc",
			wantErr: false,
		},
		{
			name: "stp	s27, s26, [sp], #-256",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfb, 0x6b, 0xa0, 0x2c}),
				address:          0,
			},
			want:    "stp\ts27, s26, [sp], #-0x100",
			wantErr: false,
		},
		{
			name: "ldp	s1, s2, [x3], #44",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x61, 0x88, 0xc5, 0x2c}),
				address:          0,
			},
			want:    "ldp\ts1, s2, [x3], #0x2c",
			wantErr: false,
		},
		{
			name: "stp	d3, d5, [x9], #504",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x23, 0x95, 0x9f, 0x6c}),
				address:          0,
			},
			want:    "stp\td3, d5, [x9], #0x1f8",
			wantErr: false,
		},
		{
			name: "stp	d7, d11, [x10], #-512",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x47, 0x2d, 0xa0, 0x6c}),
				address:          0,
			},
			want:    "stp\td7, d11, [x10], #-0x200",
			wantErr: false,
		},
		{
			name: "ldp	d2, d3, [x30], #-8",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc2, 0x8f, 0xff, 0x6c}),
				address:          0,
			},
			want:    "ldp\td2, d3, [lr], #-0x8",
			wantErr: false,
		},
		{
			name: "stp	q3, q5, [sp], #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x17, 0x80, 0xac}),
				address:          0,
			},
			want:    "stp\tq3, q5, [sp], #0",
			wantErr: false,
		},
		{
			name: "stp	q17, q19, [sp], #1008",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf1, 0xcf, 0x9f, 0xac}),
				address:          0,
			},
			want:    "stp\tq17, q19, [sp], #0x3f0",
			wantErr: false,
		},
		{
			name: "ldp	q23, q29, [x1], #-1024",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x37, 0x74, 0xe0, 0xac}),
				address:          0,
			},
			want:    "ldp\tq23, q29, [x1], #-0x400",
			wantErr: false,
		},
		{
			name: "ldp	w3, w5, [sp, #0]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x17, 0xc0, 0x29}),
				address:          0,
			},
			want:    "ldp\tw3, w5, [sp, #0]!",
			wantErr: false,
		},
		{
			name: "stp	wzr, w9, [sp, #252]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xa7, 0x9f, 0x29}),
				address:          0,
			},
			want:    "stp\twzr, w9, [sp, #0xfc]!",
			wantErr: false,
		},
		{
			name: "ldp	w2, wzr, [sp, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe2, 0x7f, 0xe0, 0x29}),
				address:          0,
			},
			want:    "ldp\tw2, wzr, [sp, #-0x100]!",
			wantErr: false,
		},
		{
			name: "ldp	w9, w10, [sp, #4]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xab, 0xc0, 0x29}),
				address:          0,
			},
			want:    "ldp\tw9, w10, [sp, #0x4]!",
			wantErr: false,
		},
		{
			name: "ldpsw	x9, x10, [sp, #4]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xab, 0xc0, 0x69}),
				address:          0,
			},
			want:    "ldpsw\tx9, x10, [sp, #0x4]!",
			wantErr: false,
		},
		{
			name: "ldpsw	x9, x10, [x2, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x28, 0xe0, 0x69}),
				address:          0,
			},
			want:    "ldpsw\tx9, x10, [x2, #-0x100]!",
			wantErr: false,
		},
		{
			name: "ldpsw	x20, x30, [sp, #252]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0xfb, 0xdf, 0x69}),
				address:          0,
			},
			want:    "ldpsw\tx20, lr, [sp, #0xfc]!",
			wantErr: false,
		},
		{
			name: "ldp	x21, x29, [x2, #504]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x55, 0xf4, 0xdf, 0xa9}),
				address:          0,
			},
			want:    "ldp\tx21, fp, [x2, #0x1f8]!",
			wantErr: false,
		},
		{
			name: "ldp	x22, x23, [x3, #-512]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x76, 0x5c, 0xe0, 0xa9}),
				address:          0,
			},
			want:    "ldp\tx22, x23, [x3, #-0x200]!",
			wantErr: false,
		},
		{
			name: "ldp	x24, x25, [x4, #8]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x98, 0xe4, 0xc0, 0xa9}),
				address:          0,
			},
			want:    "ldp\tx24, x25, [x4, #0x8]!",
			wantErr: false,
		},
		{
			name: "ldp	s29, s28, [sp, #252]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfd, 0xf3, 0xdf, 0x2d}),
				address:          0,
			},
			want:    "ldp\ts29, s28, [sp, #0xfc]!",
			wantErr: false,
		},
		{
			name: "stp	s27, s26, [sp, #-256]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfb, 0x6b, 0xa0, 0x2d}),
				address:          0,
			},
			want:    "stp\ts27, s26, [sp, #-0x100]!",
			wantErr: false,
		},
		{
			name: "ldp	s1, s2, [x3, #44]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x61, 0x88, 0xc5, 0x2d}),
				address:          0,
			},
			want:    "ldp\ts1, s2, [x3, #0x2c]!",
			wantErr: false,
		},
		{
			name: "stp	d3, d5, [x9, #504]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x23, 0x95, 0x9f, 0x6d}),
				address:          0,
			},
			want:    "stp\td3, d5, [x9, #0x1f8]!",
			wantErr: false,
		},
		{
			name: "stp	d7, d11, [x10, #-512]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x47, 0x2d, 0xa0, 0x6d}),
				address:          0,
			},
			want:    "stp\td7, d11, [x10, #-0x200]!",
			wantErr: false,
		},
		{
			name: "ldp	d2, d3, [x30, #-8]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc2, 0x8f, 0xff, 0x6d}),
				address:          0,
			},
			want:    "ldp\td2, d3, [lr, #-0x8]!",
			wantErr: false,
		},
		{
			name: "stp	q3, q5, [sp, #0]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x17, 0x80, 0xad}),
				address:          0,
			},
			want:    "stp\tq3, q5, [sp, #0]!",
			wantErr: false,
		},
		{
			name: "stp	q17, q19, [sp, #1008]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf1, 0xcf, 0x9f, 0xad}),
				address:          0,
			},
			want:    "stp\tq17, q19, [sp, #0x3f0]!",
			wantErr: false,
		},
		{
			name: "ldp	q23, q29, [x1, #-1024]!",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x37, 0x74, 0xe0, 0xad}),
				address:          0,
			},
			want:    "ldp\tq23, q29, [x1, #-0x400]!",
			wantErr: false,
		},
		{
			name: "ldnp	w3, w5, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x17, 0x40, 0x28}),
				address:          0,
			},
			want:    "ldnp\tw3, w5, [sp]",
			wantErr: false,
		},
		{
			name: "stnp	wzr, w9, [sp, #252]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xa7, 0x1f, 0x28}),
				address:          0,
			},
			want:    "stnp\twzr, w9, [sp, #0xfc]",
			wantErr: false,
		},
		{
			name: "ldnp	w2, wzr, [sp, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe2, 0x7f, 0x60, 0x28}),
				address:          0,
			},
			want:    "ldnp\tw2, wzr, [sp, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldnp	w9, w10, [sp, #4]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xab, 0x40, 0x28}),
				address:          0,
			},
			want:    "ldnp\tw9, w10, [sp, #0x4]",
			wantErr: false,
		},
		{
			name: "ldnp	x21, x29, [x2, #504]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x55, 0xf4, 0x5f, 0xa8}),
				address:          0,
			},
			want:    "ldnp\tx21, fp, [x2, #0x1f8]",
			wantErr: false,
		},
		{
			name: "ldnp	x22, x23, [x3, #-512]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x76, 0x5c, 0x60, 0xa8}),
				address:          0,
			},
			want:    "ldnp\tx22, x23, [x3, #-0x200]",
			wantErr: false,
		},
		{
			name: "ldnp	x24, x25, [x4, #8]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x98, 0xe4, 0x40, 0xa8}),
				address:          0,
			},
			want:    "ldnp\tx24, x25, [x4, #0x8]",
			wantErr: false,
		},
		{
			name: "ldnp	s29, s28, [sp, #252]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfd, 0xf3, 0x5f, 0x2c}),
				address:          0,
			},
			want:    "ldnp\ts29, s28, [sp, #0xfc]",
			wantErr: false,
		},
		{
			name: "stnp	s27, s26, [sp, #-256]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xfb, 0x6b, 0x20, 0x2c}),
				address:          0,
			},
			want:    "stnp\ts27, s26, [sp, #-0x100]",
			wantErr: false,
		},
		{
			name: "ldnp	s1, s2, [x3, #44]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x61, 0x88, 0x45, 0x2c}),
				address:          0,
			},
			want:    "ldnp\ts1, s2, [x3, #0x2c]",
			wantErr: false,
		},
		{
			name: "stnp	d3, d5, [x9, #504]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x23, 0x95, 0x1f, 0x6c}),
				address:          0,
			},
			want:    "stnp\td3, d5, [x9, #0x1f8]",
			wantErr: false,
		},
		{
			name: "stnp	d7, d11, [x10, #-512]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x47, 0x2d, 0x20, 0x6c}),
				address:          0,
			},
			want:    "stnp\td7, d11, [x10, #-0x200]",
			wantErr: false,
		},
		{
			name: "ldnp	d2, d3, [x30, #-8]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc2, 0x8f, 0x7f, 0x6c}),
				address:          0,
			},
			want:    "ldnp\td2, d3, [lr, #-0x8]",
			wantErr: false,
		},
		{
			name: "stnp	q3, q5, [sp]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x17, 0x00, 0xac}),
				address:          0,
			},
			want:    "stnp\tq3, q5, [sp]",
			wantErr: false,
		},
		{
			name: "stnp	q17, q19, [sp, #1008]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf1, 0xcf, 0x1f, 0xac}),
				address:          0,
			},
			want:    "stnp\tq17, q19, [sp, #0x3f0]",
			wantErr: false,
		},
		{
			name: "ldnp	q23, q29, [x1, #-1024]",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x37, 0x74, 0x60, 0xac}),
				address:          0,
			},
			want:    "ldnp\tq23, q29, [x1, #-0x400]",
			wantErr: false,
		},
		{
			name: "orr	w3, w9, #0xffff0000",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x23, 0x3d, 0x10, 0x32}),
				address:          0,
			},
			want:    "orr\tw3, w9, #0xffff0000",
			wantErr: false,
		},
		{
			name: "orr	wsp, w10, #0xe00000ff",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x29, 0x03, 0x32}),
				address:          0,
			},
			want:    "orr\twsp, w10, #0xe00000ff",
			wantErr: false,
		},
		{
			name: "orr	w9, w10, #0x3ff",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x25, 0x00, 0x32}),
				address:          0,
			},
			want:    "orr\tw9, w10, #0x3ff",
			wantErr: false,
		},
		{
			name: "and	w14, w15, #0x80008000",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xee, 0x81, 0x01, 0x12}),
				address:          0,
			},
			want:    "and\tw14, w15, #0x80008000",
			wantErr: false,
		},
		{
			name: "and	w12, w13, #0xffc3ffc3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0xad, 0x0a, 0x12}),
				address:          0,
			},
			want:    "and\tw12, w13, #0xffc3ffc3",
			wantErr: false,
		},
		{
			name: "and	w11, wzr, #0x30003",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xeb, 0x87, 0x00, 0x12}),
				address:          0,
			},
			want:    "and\tw11, wzr, #0x30003",
			wantErr: false,
		},
		{
			name: "eor	w3, w6, #0xe0e0e0e0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc3, 0xc8, 0x03, 0x52}),
				address:          0,
			},
			want:    "eor\tw3, w6, #0xe0e0e0e0",
			wantErr: false,
		},
		{
			name: "eor	wsp, wzr, #0x3030303",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xc7, 0x00, 0x52}),
				address:          0,
			},
			want:    "eor\twsp, wzr, #0x3030303",
			wantErr: false,
		},
		{
			name: "eor	w16, w17, #0x81818181",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x30, 0xc6, 0x01, 0x52}),
				address:          0,
			},
			want:    "eor\tw16, w17, #0x81818181",
			wantErr: false,
		},
		{
			name: "tst	w18, #0xcccccccc",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0xe6, 0x02, 0x72}),
				address:          0,
			},
			want:    "tst\tw18, #0xcccccccc",
			wantErr: false,
		},
		{
			name: "ands	w19, w20, #0x33333333",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0xe6, 0x00, 0x72}),
				address:          0,
			},
			want:    "ands\tw19, w20, #0x33333333",
			wantErr: false,
		},
		{
			name: "ands	w21, w22, #0x99999999",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd5, 0xe6, 0x01, 0x72}),
				address:          0,
			},
			want:    "ands\tw21, w22, #0x99999999",
			wantErr: false,
		},
		{
			name: "tst	w3, #0xaaaaaaaa",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0xf0, 0x01, 0x72}),
				address:          0,
			},
			want:    "tst\tw3, #0xaaaaaaaa",
			wantErr: false,
		},
		{
			name: "tst	wzr, #0x55555555",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xf3, 0x00, 0x72}),
				address:          0,
			},
			want:    "tst\twzr, #0x55555555",
			wantErr: false,
		},
		{
			name: "eor	x3, x5, #0xffffffffc000000",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0x84, 0x66, 0xd2}),
				address:          0,
			},
			want:    "eor\tx3, x5, #0xffffffffc000000",
			wantErr: false,
		},
		{
			name: "and	x9, x10, #0x7fffffffffff",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xb9, 0x40, 0x92}),
				address:          0,
			},
			want:    "and\tx9, x10, #0x7fffffffffff",
			wantErr: false,
		},
		{
			name: "orr	x11, x12, #0x8000000000000fff",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8b, 0x31, 0x41, 0xb2}),
				address:          0,
			},
			want:    "orr\tx11, x12, #0x8000000000000fff",
			wantErr: false,
		},
		{
			name: "orr	x3, x9, #0xffff0000ffff0000",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x23, 0x3d, 0x10, 0xb2}),
				address:          0,
			},
			want:    "orr\tx3, x9, #0xffff0000ffff0000",
			wantErr: false,
		},
		{
			name: "orr	sp, x10, #0xe00000ffe00000ff",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x29, 0x03, 0xb2}),
				address:          0,
			},
			want:    "orr\tsp, x10, #0xe00000ffe00000ff",
			wantErr: false,
		},
		{
			name: "orr	x9, x10, #0x3ff000003ff",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x25, 0x00, 0xb2}),
				address:          0,
			},
			want:    "orr\tx9, x10, #0x3ff000003ff",
			wantErr: false,
		},
		{
			name: "and	x14, x15, #0x8000800080008000",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xee, 0x81, 0x01, 0x92}),
				address:          0,
			},
			want:    "and\tx14, x15, #0x8000800080008000",
			wantErr: false,
		},
		{
			name: "and	x12, x13, #0xffc3ffc3ffc3ffc3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0xad, 0x0a, 0x92}),
				address:          0,
			},
			want:    "and\tx12, x13, #0xffc3ffc3ffc3ffc3",
			wantErr: false,
		},
		{
			name: "and	x11, xzr, #0x3000300030003",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xeb, 0x87, 0x00, 0x92}),
				address:          0,
			},
			want:    "and\tx11, xzr, #0x3000300030003",
			wantErr: false,
		},
		{
			name: "eor	x3, x6, #0xe0e0e0e0e0e0e0e0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc3, 0xc8, 0x03, 0xd2}),
				address:          0,
			},
			want:    "eor\tx3, x6, #0xe0e0e0e0e0e0e0e0",
			wantErr: false,
		},
		{
			name: "eor	sp, xzr, #0x303030303030303",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xc7, 0x00, 0xd2}),
				address:          0,
			},
			want:    "eor\tsp, xzr, #0x303030303030303",
			wantErr: false,
		},
		{
			name: "eor	x16, x17, #0x8181818181818181",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x30, 0xc6, 0x01, 0xd2}),
				address:          0,
			},
			want:    "eor\tx16, x17, #0x8181818181818181",
			wantErr: false,
		},
		{
			name: "tst	x18, #0xcccccccccccccccc",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0xe6, 0x02, 0xf2}),
				address:          0,
			},
			want:    "tst\tx18, #0xcccccccccccccccc",
			wantErr: false,
		},
		{
			name: "ands	x19, x20, #0x3333333333333333",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0xe6, 0x00, 0xf2}),
				address:          0,
			},
			want:    "ands\tx19, x20, #0x3333333333333333",
			wantErr: false,
		},
		{
			name: "ands	x21, x22, #0x9999999999999999",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd5, 0xe6, 0x01, 0xf2}),
				address:          0,
			},
			want:    "ands\tx21, x22, #0x9999999999999999",
			wantErr: false,
		},
		{
			name: "tst	x3, #0xaaaaaaaaaaaaaaaa",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0xf0, 0x01, 0xf2}),
				address:          0,
			},
			want:    "tst\tx3, #0xaaaaaaaaaaaaaaaa",
			wantErr: false,
		},
		{
			name: "tst	xzr, #0x5555555555555555",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xf3, 0x00, 0xf2}),
				address:          0,
			},
			want:    "tst\txzr, #0x5555555555555555",
			wantErr: false,
		},
		{
			name: "mov	w3, #983055",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x8f, 0x00, 0x32}),
				address:          0,
			},
			want:    "mov\tw3, #0xf000f",
			wantErr: false,
		},
		// TODO: ADD THIS BACK IN
		// {
		// 	name: "mov	x10, #-6148914691236517206",
		// 	args: args{
		// 		instructionValue: binary.LittleEndian.Uint32([]byte{0xea, 0xf3, 0x01, 0xb2}),
		// 		address:          0,
		// 	},
		// 	want: "mov	x10, #-6148914691236517206",
		// 	wantErr: false,
		// },
		{
			name: "and	w2, w3, #0xfffffffd",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x62, 0x78, 0x1e, 0x12}),
				address:          0,
			},
			want:    "and\tw2, w3, #0xfffffffd",
			wantErr: false,
		},
		{
			name: "orr	w0, w1, #0xfffffffd",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x78, 0x1e, 0x32}),
				address:          0,
			},
			want:    "orr\tw0, w1, #0xfffffffd",
			wantErr: false,
		},
		{
			name: "eor	w16, w17, #0xfffffff9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x30, 0x76, 0x1d, 0x52}),
				address:          0,
			},
			want:    "eor\tw16, w17, #0xfffffff9",
			wantErr: false,
		},
		{
			name: "ands	w19, w20, #0xfffffff0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x93, 0x6e, 0x1c, 0x72}),
				address:          0,
			},
			want:    "ands\tw19, w20, #0xfffffff0",
			wantErr: false,
		},
		{
			name: "and	w12, w23, w21",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x02, 0x15, 0x0a}),
				address:          0,
			},
			want:    "and\tw12, w23, w21",
			wantErr: false,
		},
		{
			name: "and	w16, w15, w1, lsl #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf0, 0x05, 0x01, 0x0a}),
				address:          0,
			},
			want:    "and\tw16, w15, w1, lsl #0x1",
			wantErr: false,
		},
		{
			name: "and	w9, w4, w10, lsl #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x7c, 0x0a, 0x0a}),
				address:          0,
			},
			want:    "and\tw9, w4, w10, lsl #0x1f",
			wantErr: false,
		},
		{
			name: "and	w3, w30, w11",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc3, 0x03, 0x0b, 0x0a}),
				address:          0,
			},
			want:    "and\tw3, w30, w11",
			wantErr: false,
		},
		{
			name: "and	x3, x5, x7, lsl #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xfc, 0x07, 0x8a}),
				address:          0,
			},
			want:    "and\tx3, x5, x7, lsl #0x3f",
			wantErr: false,
		},
		{
			name: "and	x5, x14, x19, asr #4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc5, 0x11, 0x93, 0x8a}),
				address:          0,
			},
			want:    "and\tx5, x14, x19, asr #0x4",
			wantErr: false,
		},
		{
			name: "and	w3, w17, w19, ror #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x23, 0x7e, 0xd3, 0x0a}),
				address:          0,
			},
			want:    "and\tw3, w17, w19, ror #0x1f",
			wantErr: false,
		},
		{
			name: "and	w0, w2, wzr, lsr #17",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0x44, 0x5f, 0x0a}),
				address:          0,
			},
			want:    "and\tw0, w2, wzr, lsr #0x11",
			wantErr: false,
		},
		{
			name: "and	w3, w30, w11, asr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc3, 0x03, 0x8b, 0x0a}),
				address:          0,
			},
			want:    "and\tw3, w30, w11, asr #0",
			wantErr: false,
		},
		{
			name: "and	xzr, x4, x26",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x00, 0x1a, 0x8a}),
				address:          0,
			},
			want:    "and\txzr, x4, x26",
			wantErr: false,
		},
		{
			name: "and	w3, wzr, w20, ror #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0xd4, 0x0a}),
				address:          0,
			},
			want:    "and\tw3, wzr, w20, ror #0",
			wantErr: false,
		},
		{
			name: "and	x7, x20, xzr, asr #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x87, 0xfe, 0x9f, 0x8a}),
				address:          0,
			},
			want:    "and\tx7, x20, xzr, asr #0x3f",
			wantErr: false,
		},
		{
			name: "bic	x13, x20, x14, lsl #47",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8d, 0xbe, 0x2e, 0x8a}),
				address:          0,
			},
			want:    "bic\tx13, x20, x14, lsl #0x2f",
			wantErr: false,
		},
		{
			name: "bic	w2, w7, w9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe2, 0x00, 0x29, 0x0a}),
				address:          0,
			},
			want:    "bic\tw2, w7, w9",
			wantErr: false,
		},
		{
			name: "orr	w2, w7, w0, asr #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe2, 0x7c, 0x80, 0x2a}),
				address:          0,
			},
			want:    "orr\tw2, w7, w0, asr #0x1f",
			wantErr: false,
		},
		{
			name: "orr	x8, x9, x10, lsl #12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x28, 0x31, 0x0a, 0xaa}),
				address:          0,
			},
			want:    "orr\tx8, x9, x10, lsl #0xc",
			wantErr: false,
		},
		{
			name: "orn	x3, x5, x7, asr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0x00, 0xa7, 0xaa}),
				address:          0,
			},
			want:    "orn\tx3, x5, x7, asr #0",
			wantErr: false,
		},
		{
			name: "orn	w2, w5, w29",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa2, 0x00, 0x3d, 0x2a}),
				address:          0,
			},
			want:    "orn\tw2, w5, w29",
			wantErr: false,
		},
		{
			name: "ands	w7, wzr, w9, lsl #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe7, 0x07, 0x09, 0x6a}),
				address:          0,
			},
			want:    "ands\tw7, wzr, w9, lsl #0x1",
			wantErr: false,
		},
		{
			name: "ands	x3, x5, x20, ror #63",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0xfc, 0xd4, 0xea}),
				address:          0,
			},
			want:    "ands\tx3, x5, x20, ror #0x3f",
			wantErr: false,
		},
		{
			name: "bics	w3, w5, w7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa3, 0x00, 0x27, 0x6a}),
				address:          0,
			},
			want:    "bics\tw3, w5, w7",
			wantErr: false,
		},
		{
			name: "bics	x3, xzr, x3, lsl #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x07, 0x23, 0xea}),
				address:          0,
			},
			want:    "bics\tx3, xzr, x3, lsl #0x1",
			wantErr: false,
		},
		{
			name: "tst	w3, w7, lsl #31",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x7c, 0x07, 0x6a}),
				address:          0,
			},
			want:    "tst\tw3, w7, lsl #0x1f",
			wantErr: false,
		},
		{
			name: "tst	x2, x20, asr #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x00, 0x94, 0xea}),
				address:          0,
			},
			want:    "tst\tx2, x20, asr #0",
			wantErr: false,
		},
		{
			name: "mov	x3, x6",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0x06, 0xaa}),
				address:          0,
			},
			want:    "mov\tx3, x6",
			wantErr: false,
		},
		{
			name: "mov	x3, xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0x1f, 0xaa}),
				address:          0,
			},
			want:    "mov\tx3, xzr",
			wantErr: false,
		},
		{
			name: "mov	wzr, w2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x03, 0x02, 0x2a}),
				address:          0,
			},
			want:    "mov\twzr, w2",
			wantErr: false,
		},
		{
			name: "mov	w3, w5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe3, 0x03, 0x05, 0x2a}),
				address:          0,
			},
			want:    "mov\tw3, w5",
			wantErr: false,
		},
		{
			name: "mov	w1, #65535",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe1, 0xff, 0x9f, 0x52}),
				address:          0,
			},
			want:    "mov\tw1, #0xffff",
			wantErr: false,
		},
		{
			name: "movz	w2, #0, lsl #16",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x02, 0x00, 0xa0, 0x52}),
				address:          0,
			},
			want:    "movz\tw2, #0, lsl #0x10",
			wantErr: false,
		},
		{
			name: "mov	w2, #-1235",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x42, 0x9a, 0x80, 0x12}),
				address:          0,
			},
			want:    "mov\tw2, #0xfffffb2d",
			wantErr: false,
		},
		{
			name: "mov	x2, #5299989643264",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x42, 0x9a, 0xc0, 0xd2}),
				address:          0,
			},
			want:    "mov\tx2, #0x4d200000000",
			wantErr: false,
		},
		{
			name: "movk	xzr, #4321, lsl #48",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0x1c, 0xe2, 0xf2}),
				address:          0,
			},
			want:    "movk\txzr, #0x10e1, lsl #0x30",
			wantErr: false,
		},
		{
			name: "adrp	x30, #4096",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1e, 0x00, 0x00, 0xb0}),
				address:          0,
			},
			want:    "adrp\tlr, 0x1000",
			wantErr: false,
		},
		{
			name: "adr	x20, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x14, 0x00, 0x00, 0x10}),
				address:          0,
			},
			want:    "adr\tx20, 0x0",
			wantErr: false,
		},
		{
			name: "adr	x9, #-1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xff, 0xff, 0x70}),
				address:          0,
			},
			want:    "adr\tx9, 0xffffffffffffffff",
			wantErr: false,
		},
		{
			name: "adr	x5, #1048575",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe5, 0xff, 0x7f, 0x70}),
				address:          0,
			},
			want:    "adr\tx5, 0xfffff",
			wantErr: false,
		},
		{
			name: "adr	x9, #1048575",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xff, 0x7f, 0x70}),
				address:          0,
			},
			want:    "adr\tx9, 0xfffff",
			wantErr: false,
		},
		{
			name: "adr	x2, #-1048576",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x02, 0x00, 0x80, 0x10}),
				address:          0,
			},
			want:    "adr\tx2, 0xfffffffffff00000",
			wantErr: false,
		},
		{
			name: "adrp	x9, #4294963200",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xff, 0x7f, 0xf0}),
				address:          0,
			},
			want:    "adrp\tx9, 0xfffff000",
			wantErr: false,
		},
		{
			name: "adrp	x20, #-4294967296",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x14, 0x00, 0x80, 0x90}),
				address:          0,
			},
			want:    "adrp\tx20, 0xffffffff00000000",
			wantErr: false,
		},
		{
			name: "nop",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x20, 0x03, 0xd5}),
				address:          0,
			},
			want:    "nop",
			wantErr: false,
		},
		{
			name: "hint	#127",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x2f, 0x03, 0xd5}),
				address:          0,
			},
			want:    "hint\t#0x7f",
			wantErr: false,
		},
		{
			name: "nop",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x20, 0x03, 0xd5}),
				address:          0,
			},
			want:    "nop",
			wantErr: false,
		},
		{
			name: "yield",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0x20, 0x03, 0xd5}),
				address:          0,
			},
			want:    "yield",
			wantErr: false,
		},
		{
			name: "wfe",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x20, 0x03, 0xd5}),
				address:          0,
			},
			want:    "wfe",
			wantErr: false,
		},
		{
			name: "wfi",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x7f, 0x20, 0x03, 0xd5}),
				address:          0,
			},
			want:    "wfi",
			wantErr: false,
		},
		{
			name: "sev",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x20, 0x03, 0xd5}),
				address:          0,
			},
			want:    "sev",
			wantErr: false,
		},
		{
			name: "sevl",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x20, 0x03, 0xd5}),
				address:          0,
			},
			want:    "sevl",
			wantErr: false,
		},
		{
			name: "dgh",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0x20, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dgh",
			wantErr: false,
		},
		{
			name: "clrex",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x3f, 0x03, 0xd5}),
				address:          0,
			},
			want:    "clrex",
			wantErr: false,
		},
		{
			name: "clrex	#0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x30, 0x03, 0xd5}),
				address:          0,
			},
			want:    "clrex\t#0",
			wantErr: false,
		},
		{
			name: "clrex	#7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x37, 0x03, 0xd5}),
				address:          0,
			},
			want:    "clrex\t#0x7",
			wantErr: false,
		},
		{
			name: "clrex",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0x3f, 0x03, 0xd5}),
				address:          0,
			},
			want:    "clrex",
			wantErr: false,
		},
		{
			name: "ssbb",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x30, 0x03, 0xd5}),
				address:          0,
			},
			want:    "ssbb",
			wantErr: false,
		},
		{
			name: "pssbb",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x34, 0x03, 0xd5}),
				address:          0,
			},
			want:    "pssbb",
			wantErr: false,
		},
		{
			name: "dsb	#12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x3c, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dsb\t#12",
			wantErr: false,
		},
		{
			name: "dsb	sy",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x3f, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dsb\tsy",
			wantErr: false,
		},
		{
			name: "dsb	oshld",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x31, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dsb\toshld",
			wantErr: false,
		},
		{
			name: "dsb	oshst",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x32, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dsb\toshst",
			wantErr: false,
		},
		{
			name: "dsb	osh",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x33, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dsb\tosh",
			wantErr: false,
		},
		{
			name: "dsb	nshld",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x35, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dsb\tnshld",
			wantErr: false,
		},
		{
			name: "dsb	nshst",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x36, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dsb\tnshst",
			wantErr: false,
		},
		{
			name: "dsb	nsh",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x37, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dsb\tnsh",
			wantErr: false,
		},
		{
			name: "dsb	ishld",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x39, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dsb\tishld",
			wantErr: false,
		},
		{
			name: "dsb	ishst",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x3a, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dsb\tishst",
			wantErr: false,
		},
		{
			name: "dsb	ish",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x3b, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dsb\tish",
			wantErr: false,
		},
		{
			name: "dsb	ld",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x3d, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dsb\tld",
			wantErr: false,
		},
		{
			name: "dsb	st",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x3e, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dsb\tst",
			wantErr: false,
		},
		{
			name: "dsb	sy",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x3f, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dsb\tsy",
			wantErr: false,
		},
		{
			name: "dmb	#0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x30, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dmb\t#0",
			wantErr: false,
		},
		{
			name: "dmb	#12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x3c, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dmb\t#12",
			wantErr: false,
		},
		{
			name: "dmb	sy",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x3f, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dmb\tsy",
			wantErr: false,
		},
		{
			name: "dmb	oshld",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x31, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dmb\toshld",
			wantErr: false,
		},
		{
			name: "dmb	oshst",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x32, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dmb\toshst",
			wantErr: false,
		},
		{
			name: "dmb	osh",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x33, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dmb\tosh",
			wantErr: false,
		},
		{
			name: "dmb	nshld",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x35, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dmb\tnshld",
			wantErr: false,
		},
		{
			name: "dmb	nshst",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x36, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dmb\tnshst",
			wantErr: false,
		},
		{
			name: "dmb	nsh",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x37, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dmb\tnsh",
			wantErr: false,
		},
		{
			name: "dmb	ishld",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x39, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dmb\tishld",
			wantErr: false,
		},
		{
			name: "dmb	ishst",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x3a, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dmb\tishst",
			wantErr: false,
		},
		{
			name: "dmb	ish",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x3b, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dmb\tish",
			wantErr: false,
		},
		{
			name: "dmb	ld",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x3d, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dmb\tld",
			wantErr: false,
		},
		{
			name: "dmb	st",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x3e, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dmb\tst",
			wantErr: false,
		},
		{
			name: "dmb	sy",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x3f, 0x03, 0xd5}),
				address:          0,
			},
			want:    "dmb\tsy",
			wantErr: false,
		},
		{
			name: "isb",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0x3f, 0x03, 0xd5}),
				address:          0,
			},
			want:    "isb",
			wantErr: false,
		},
		{
			name: "isb",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0x3f, 0x03, 0xd5}),
				address:          0,
			},
			want:    "isb",
			wantErr: false,
		},
		{
			name: "isb	#12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0x3c, 0x03, 0xd5}),
				address:          0,
			},
			want:    "isb\t#0xc",
			wantErr: false,
		},
		{
			name: "msr	spsel, #0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xbf, 0x40, 0x00, 0xd5}),
				address:          0,
			},
			want:    "msr\tspsel, #0",
			wantErr: false,
		},
		{
			name: "msr	daifset, #15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0x4f, 0x03, 0xd5}),
				address:          0,
			},
			want:    "msr\tdaifset, #0xf",
			wantErr: false,
		},
		{
			name: "msr	daifclr, #12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0x4c, 0x03, 0xd5}),
				address:          0,
			},
			want:    "msr\tdaifclr, #0xc",
			wantErr: false,
		},
		{
			name: "sys	#7, c5, c9, #7, x5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe5, 0x59, 0x0f, 0xd5}),
				address:          0,
			},
			want:    "sys\t#0x7, c5, c9, #0x7, x5",
			wantErr: false,
		},
		{
			name: "sys	#0, c15, c15, #2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5f, 0xff, 0x08, 0xd5}),
				address:          0,
			},
			want:    "sys\t#0, c15, c15, #0x2",
			wantErr: false,
		},
		{
			name: "sysl	x9, #7, c5, c9, #7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x59, 0x2f, 0xd5}),
				address:          0,
			},
			want:    "sysl\tx9, #0x7, c5, c9, #0x7",
			wantErr: false,
		},
		{
			name: "sysl	x1, #0, c15, c15, #2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x41, 0xff, 0x28, 0xd5}),
				address:          0,
			},
			want:    "sysl\tx1, #0, c15, c15, #0x2",
			wantErr: false,
		},
		{
			name: "ic	ialluis",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x71, 0x08, 0xd5}),
				address:          0,
			},
			want:    "ic\tialluis",
			wantErr: false,
		},
		{
			name: "ic	iallu",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x75, 0x08, 0xd5}),
				address:          0,
			},
			want:    "ic\tiallu",
			wantErr: false,
		},
		{
			name: "ic	ivau, x9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x75, 0x0b, 0xd5}),
				address:          0,
			},
			want:    "ic\tivau, x9",
			wantErr: false,
		},
		{
			name: "dc	zva, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x74, 0x0b, 0xd5}),
				address:          0,
			},
			want:    "dc\tZVA, x12",
			wantErr: false,
		},
		{
			name: "dc	ivac, xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x3f, 0x76, 0x08, 0xd5}),
				address:          0,
			},
			want:    "dc\tIVAC, xzr",
			wantErr: false,
		},
		{
			name: "dc	isw, x2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x42, 0x76, 0x08, 0xd5}),
				address:          0,
			},
			want:    "dc\tISW, x2",
			wantErr: false,
		},
		{
			name: "dc	cvac, x9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x7a, 0x0b, 0xd5}),
				address:          0,
			},
			want:    "dc\tCVAC, x9",
			wantErr: false,
		},
		{
			name: "dc	csw, x10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4a, 0x7a, 0x08, 0xd5}),
				address:          0,
			},
			want:    "dc\tCSW, x10",
			wantErr: false,
		},
		{
			name: "dc	cvau, x0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x20, 0x7b, 0x0b, 0xd5}),
				address:          0,
			},
			want:    "dc\tCVAU, x0",
			wantErr: false,
		},
		{
			name: "dc	civac, x3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x23, 0x7e, 0x0b, 0xd5}),
				address:          0,
			},
			want:    "dc\tCIVAC, x3",
			wantErr: false,
		},
		{
			name: "dc	cisw, x30",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x5e, 0x7e, 0x08, 0xd5}),
				address:          0,
			},
			want:    "dc\tCISW, lr",
			wantErr: false,
		},
		{
			name: "at	s1e1r, x19",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x13, 0x78, 0x08, 0xd5}),
				address:          0,
			},
			want:    "at\tS1E1R, x19",
			wantErr: false,
		},
		{
			name: "at	s1e2r, x19",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x13, 0x78, 0x0c, 0xd5}),
				address:          0,
			},
			want:    "at\tS1E2R, x19",
			wantErr: false,
		},
		{
			name: "at	s1e3r, x19",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x13, 0x78, 0x0e, 0xd5}),
				address:          0,
			},
			want:    "at\tS1E3R, x19",
			wantErr: false,
		},
		{
			name: "at	s1e1w, x19",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x33, 0x78, 0x08, 0xd5}),
				address:          0,
			},
			want:    "at\tS1E1W, x19",
			wantErr: false,
		},
		{
			name: "at	s1e2w, x19",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x33, 0x78, 0x0c, 0xd5}),
				address:          0,
			},
			want:    "at\tS1E2W, x19",
			wantErr: false,
		},
		{
			name: "at	s1e3w, x19",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x33, 0x78, 0x0e, 0xd5}),
				address:          0,
			},
			want:    "at\tS1E3W, x19",
			wantErr: false,
		},
		{
			name: "at	s1e0r, x19",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x53, 0x78, 0x08, 0xd5}),
				address:          0,
			},
			want:    "at\tS1E0R, x19",
			wantErr: false,
		},
		{
			name: "at	s1e0w, x19",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x73, 0x78, 0x08, 0xd5}),
				address:          0,
			},
			want:    "at\tS1E0W, x19",
			wantErr: false,
		},
		{
			name: "at	s12e1r, x20",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x94, 0x78, 0x0c, 0xd5}),
				address:          0,
			},
			want:    "at\tS12E1R, x20",
			wantErr: false,
		},
		{
			name: "at	s12e1w, x20",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb4, 0x78, 0x0c, 0xd5}),
				address:          0,
			},
			want:    "at\tS12E1W, x20",
			wantErr: false,
		},
		{
			name: "at	s12e0r, x20",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xd4, 0x78, 0x0c, 0xd5}),
				address:          0,
			},
			want:    "at\tS12E0R, x20",
			wantErr: false,
		},
		{
			name: "at	s12e0w, x20",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf4, 0x78, 0x0c, 0xd5}),
				address:          0,
			},
			want:    "at\tS12E0W, x20",
			wantErr: false,
		},
		{
			name: "tlbi	ipas2e1is, x4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x24, 0x80, 0x0c, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tipas2e1is, x4",
			wantErr: false,
		},
		{
			name: "tlbi	ipas2le1is, x9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x80, 0x0c, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tipas2le1is, x9",
			wantErr: false,
		},
		{
			name: "tlbi	vmalle1is",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x83, 0x08, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tvmalle1is",
			wantErr: false,
		},
		{
			name: "tlbi	alle2is",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x83, 0x0c, 0xd5}),
				address:          0,
			},
			want:    "tlbi\talle2is",
			wantErr: false,
		},
		{
			name: "tlbi	alle3is",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x83, 0x0e, 0xd5}),
				address:          0,
			},
			want:    "tlbi\talle3is",
			wantErr: false,
		},
		{
			name: "tlbi	vae1is, x1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x21, 0x83, 0x08, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tvae1is, x1",
			wantErr: false,
		},
		{
			name: "tlbi	vae2is, x2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x22, 0x83, 0x0c, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tvae2is, x2",
			wantErr: false,
		},
		{
			name: "tlbi	vae3is, x3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x23, 0x83, 0x0e, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tvae3is, x3",
			wantErr: false,
		},
		{
			name: "tlbi	aside1is, x5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x45, 0x83, 0x08, 0xd5}),
				address:          0,
			},
			want:    "tlbi\taside1is, x5",
			wantErr: false,
		},
		{
			name: "tlbi	vaae1is, x9",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0x83, 0x08, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tvaae1is, x9",
			wantErr: false,
		},
		{
			name: "tlbi	alle1is",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x83, 0x0c, 0xd5}),
				address:          0,
			},
			want:    "tlbi\talle1is",
			wantErr: false,
		},
		{
			name: "tlbi	vale1is, x10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xaa, 0x83, 0x08, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tvale1is, x10",
			wantErr: false,
		},
		{
			name: "tlbi	vale2is, x11",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xab, 0x83, 0x0c, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tvale2is, x11",
			wantErr: false,
		},
		{
			name: "tlbi	vale3is, x13",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xad, 0x83, 0x0e, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tvale3is, x13",
			wantErr: false,
		},
		{
			name: "tlbi	vmalls12e1is",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0x83, 0x0c, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tvmalls12e1is",
			wantErr: false,
		},
		{
			name: "tlbi	vaale1is, x14",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xee, 0x83, 0x08, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tvaale1is, x14",
			wantErr: false,
		},
		{
			name: "tlbi	ipas2e1, x15",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2f, 0x84, 0x0c, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tipas2e1, x15",
			wantErr: false,
		},
		{
			name: "tlbi	ipas2le1, x16",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb0, 0x84, 0x0c, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tipas2le1, x16",
			wantErr: false,
		},
		{
			name: "tlbi	vmalle1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x87, 0x08, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tvmalle1",
			wantErr: false,
		},
		{
			name: "tlbi	alle2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x87, 0x0c, 0xd5}),
				address:          0,
			},
			want:    "tlbi\talle2",
			wantErr: false,
		},
		{
			name: "tlbi	alle3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x1f, 0x87, 0x0e, 0xd5}),
				address:          0,
			},
			want:    "tlbi\talle3",
			wantErr: false,
		},
		{
			name: "tlbi	vae1, x17",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x31, 0x87, 0x08, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tvae1, x17",
			wantErr: false,
		},
		{
			name: "tlbi	vae2, x18",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x32, 0x87, 0x0c, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tvae2, x18",
			wantErr: false,
		},
		{
			name: "tlbi	vae3, x19",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x33, 0x87, 0x0e, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tvae3, x19",
			wantErr: false,
		},
		{
			name: "tlbi	aside1, x20",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x54, 0x87, 0x08, 0xd5}),
				address:          0,
			},
			want:    "tlbi\taside1, x20",
			wantErr: false,
		},
		{
			name: "tlbi	vaae1, x21",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x75, 0x87, 0x08, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tvaae1, x21",
			wantErr: false,
		},
		{
			name: "tlbi	alle1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x9f, 0x87, 0x0c, 0xd5}),
				address:          0,
			},
			want:    "tlbi\talle1",
			wantErr: false,
		},
		{
			name: "tlbi	vale1, x22",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb6, 0x87, 0x08, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tvale1, x22",
			wantErr: false,
		},
		{
			name: "tlbi	vale2, x23",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb7, 0x87, 0x0c, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tvale2, x23",
			wantErr: false,
		},
		{
			name: "tlbi	vale3, x24",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xb8, 0x87, 0x0e, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tvale3, x24",
			wantErr: false,
		},
		{
			name: "tlbi	vmalls12e1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xdf, 0x87, 0x0c, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tvmalls12e1",
			wantErr: false,
		},
		{
			name: "tlbi	vaale1, x25",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xf9, 0x87, 0x08, 0xd5}),
				address:          0,
			},
			want:    "tlbi\tvaale1, x25",
			wantErr: false,
		},
		{
			name: "msr	teecr32_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x00, 0x12, 0xd5}),
				address:          0,
			},
			want:    "msr\tteecr32_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	osdtrrx_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0x00, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tosdtrrx_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	mdccint_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x02, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tmdccint_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	mdscr_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0x02, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tmdscr_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	osdtrtx_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0x03, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tosdtrtx_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgdtr_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x04, 0x13, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgdtr_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgdtrtx_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x05, 0x13, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgdtrrx_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	oseccr_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0x06, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\toseccr_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgvcr32_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x07, 0x14, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgvcr32_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbvr0_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0x00, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbvr0_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbvr1_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0x01, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbvr1_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbvr2_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0x02, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbvr2_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbvr3_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0x03, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbvr3_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbvr4_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0x04, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbvr4_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbvr5_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0x05, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbvr5_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbvr6_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0x06, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbvr6_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbvr7_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0x07, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbvr7_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbvr8_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0x08, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbvr8_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbvr9_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0x09, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbvr9_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbvr10_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0x0a, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbvr10_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbvr11_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0x0b, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbvr11_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbvr12_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0x0c, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbvr12_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbvr13_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0x0d, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbvr13_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbvr14_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0x0e, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbvr14_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbvr15_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0x0f, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbvr15_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbcr0_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x00, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbcr0_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbcr1_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x01, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbcr1_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbcr2_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x02, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbcr2_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbcr3_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x03, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbcr3_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbcr4_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x04, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbcr4_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbcr5_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x05, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbcr5_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbcr6_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x06, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbcr6_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbcr7_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x07, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbcr7_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbcr8_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x08, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbcr8_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbcr9_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x09, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbcr9_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbcr10_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x0a, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbcr10_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbcr11_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x0b, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbcr11_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbcr12_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x0c, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbcr12_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbcr13_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x0d, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbcr13_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbcr14_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x0e, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbcr14_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgbcr15_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x0f, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgbcr15_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwvr0_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x00, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwvr0_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwvr1_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x01, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwvr1_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwvr2_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x02, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwvr2_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwvr3_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x03, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwvr3_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwvr4_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x04, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwvr4_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwvr5_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x05, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwvr5_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwvr6_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x06, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwvr6_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwvr7_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x07, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwvr7_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwvr8_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x08, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwvr8_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwvr9_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x09, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwvr9_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwvr10_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x0a, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwvr10_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwvr11_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x0b, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwvr11_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwvr12_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x0c, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwvr12_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwvr13_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x0d, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwvr13_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwvr14_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x0e, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwvr14_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwvr15_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x0f, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwvr15_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwcr0_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x00, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwcr0_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwcr1_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x01, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwcr1_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwcr2_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x02, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwcr2_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwcr3_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x03, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwcr3_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwcr4_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x04, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwcr4_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwcr5_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x05, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwcr5_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwcr6_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x06, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwcr6_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwcr7_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x07, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwcr7_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwcr8_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x08, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwcr8_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwcr9_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x09, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwcr9_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwcr10_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x0a, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwcr10_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwcr11_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x0b, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwcr11_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwcr12_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x0c, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwcr12_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwcr13_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x0d, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwcr13_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwcr14_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x0e, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwcr14_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgwcr15_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x0f, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgwcr15_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	teehbr32_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x10, 0x12, 0xd5}),
				address:          0,
			},
			want:    "msr\tteehbr32_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	oslar_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0x10, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\toslar_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	osdlr_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0x13, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tosdlr_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgprcr_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0x14, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgprcr_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgclaimset_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x78, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgclaimset_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	dbgclaimclr_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0x79, 0x10, 0xd5}),
				address:          0,
			},
			want:    "msr\tdbgclaimclr_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	csselr_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x00, 0x1a, 0xd5}),
				address:          0,
			},
			want:    "msr\tcsselr_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	vpidr_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x00, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tvpidr_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	vmpidr_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x00, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tvmpidr_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	sctlr_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x10, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tsctlr_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	sctlr_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x10, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tsctlr_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	sctlr_el3, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x10, 0x1e, 0xd5}),
				address:          0,
			},
			want:    "msr\tsctlr_el3, x12",
			wantErr: false,
		},
		{
			name: "msr	actlr_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x10, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tactlr_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	actlr_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x10, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tactlr_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	actlr_el3, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x10, 0x1e, 0xd5}),
				address:          0,
			},
			want:    "msr\tactlr_el3, x12",
			wantErr: false,
		},
		{
			name: "msr	cpacr_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0x10, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tcpacr_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	hcr_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x11, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\thcr_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	scr_el3, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x11, 0x1e, 0xd5}),
				address:          0,
			},
			want:    "msr\tscr_el3, x12",
			wantErr: false,
		},
		{
			name: "msr	mdcr_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x11, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tmdcr_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	sder32_el3, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x11, 0x1e, 0xd5}),
				address:          0,
			},
			want:    "msr\tsder32_el3, x12",
			wantErr: false,
		},
		{
			name: "msr	cptr_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0x11, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tcptr_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	cptr_el3, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0x11, 0x1e, 0xd5}),
				address:          0,
			},
			want:    "msr\tcptr_el3, x12",
			wantErr: false,
		},
		{
			name: "msr	hstr_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6c, 0x11, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\thstr_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	hacr_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0x11, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\thacr_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	mdcr_el3, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x13, 0x1e, 0xd5}),
				address:          0,
			},
			want:    "msr\tmdcr_el3, x12",
			wantErr: false,
		},
		{
			name: "msr	ttbr0_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x20, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tttbr0_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	ttbr0_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x20, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tttbr0_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	ttbr0_el3, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x20, 0x1e, 0xd5}),
				address:          0,
			},
			want:    "msr\tttbr0_el3, x12",
			wantErr: false,
		},
		{
			name: "msr	ttbr1_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x20, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tttbr1_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	tcr_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0x20, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\ttcr_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	tcr_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0x20, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ttcr_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	tcr_el3, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0x20, 0x1e, 0xd5}),
				address:          0,
			},
			want:    "msr\ttcr_el3, x12",
			wantErr: false,
		},
		{
			name: "msr	vttbr_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x21, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tvttbr_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	vtcr_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0x21, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tvtcr_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	dacr32_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x30, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tdacr32_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	spsr_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x40, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tspsr_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	spsr_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x40, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tspsr_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	spsr_el3, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x40, 0x1e, 0xd5}),
				address:          0,
			},
			want:    "msr\tspsr_el3, x12",
			wantErr: false,
		},
		{
			name: "msr	elr_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x40, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\telr_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	elr_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x40, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\telr_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	elr_el3, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x40, 0x1e, 0xd5}),
				address:          0,
			},
			want:    "msr\telr_el3, x12",
			wantErr: false,
		},
		{
			name: "msr	sp_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x41, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tsp_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	sp_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x41, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tsp_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	sp_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x41, 0x1e, 0xd5}),
				address:          0,
			},
			want:    "msr\tsp_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	spsel, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x42, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tspsel, x12",
			wantErr: false,
		},
		{
			name: "msr	nzcv, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x42, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tnzcv, x12",
			wantErr: false,
		},
		{
			name: "msr	daif, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x42, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tdaif, x12",
			wantErr: false,
		},
		{
			name: "msr	spsr_irq, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x43, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tspsr_irq, x12",
			wantErr: false,
		},
		{
			name: "msr	spsr_abt, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x43, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tspsr_abt, x12",
			wantErr: false,
		},
		{
			name: "msr	spsr_und, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0x43, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tspsr_und, x12",
			wantErr: false,
		},
		{
			name: "msr	spsr_fiq, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6c, 0x43, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tspsr_fiq, x12",
			wantErr: false,
		},
		{
			name: "msr	fpcr, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x44, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tfpcr, x12",
			wantErr: false,
		},
		{
			name: "msr	fpsr, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x44, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tfpsr, x12",
			wantErr: false,
		},
		{
			name: "msr	dspsr_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x45, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tdspsr, x12",
			wantErr: false,
		},
		{
			name: "msr	dlr_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x45, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tdlr, x12",
			wantErr: false,
		},
		{
			name: "msr	ifsr32_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x50, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tifsr32_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	afsr0_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x51, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tafsr0_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	afsr0_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x51, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tafsr0_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	afsr0_el3, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x51, 0x1e, 0xd5}),
				address:          0,
			},
			want:    "msr\tafsr0_el3, x12",
			wantErr: false,
		},
		{
			name: "msr	afsr1_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x51, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tafsr1_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	afsr1_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x51, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tafsr1_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	afsr1_el3, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x51, 0x1e, 0xd5}),
				address:          0,
			},
			want:    "msr\tafsr1_el3, x12",
			wantErr: false,
		},
		{
			name: "msr	esr_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x52, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tesr_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	esr_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x52, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tesr_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	esr_el3, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x52, 0x1e, 0xd5}),
				address:          0,
			},
			want:    "msr\tesr_el3, x12",
			wantErr: false,
		},
		{
			name: "msr	fpexc32_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x53, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tfpexc32_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	far_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x60, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tfar_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	far_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x60, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tfar_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	far_el3, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x60, 0x1e, 0xd5}),
				address:          0,
			},
			want:    "msr\tfar_el3, x12",
			wantErr: false,
		},
		{
			name: "msr	hpfar_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0x60, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\thpfar_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	par_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x74, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tpar_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	pmcr_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x9c, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmcr_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmcntenset_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x9c, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmcntenset_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmcntenclr_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0x9c, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmcntenclr_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmovsclr_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6c, 0x9c, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmovsclr_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmselr_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0x9c, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmselr_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmccntr_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x9d, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmccntr_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmxevtyper_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x9d, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmxevtyper_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmxevcntr_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0x9d, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmxevcntr_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmuserenr_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0x9e, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmuserenr_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmintenset_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0x9e, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmintenset_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	pmintenclr_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0x9e, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmintenclr_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	pmovsset_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6c, 0x9e, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmovsset_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	mair_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xa2, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tmair_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	mair_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xa2, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tmair_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	mair_el3, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xa2, 0x1e, 0xd5}),
				address:          0,
			},
			want:    "msr\tmair_el3, x12",
			wantErr: false,
		},
		{
			name: "msr	amair_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xa3, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tamair_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	amair_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xa3, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tamair_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	amair_el3, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xa3, 0x1e, 0xd5}),
				address:          0,
			},
			want:    "msr\tamair_el3, x12",
			wantErr: false,
		},
		{
			name: "msr	vbar_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xc0, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tvbar_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	vbar_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xc0, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tvbar_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	vbar_el3, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xc0, 0x1e, 0xd5}),
				address:          0,
			},
			want:    "msr\tvbar_el3, x12",
			wantErr: false,
		},
		{
			name: "msr	rmr_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0xc0, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\trmr_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	rmr_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0xc0, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\trmr_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	rmr_el3, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0xc0, 0x1e, 0xd5}),
				address:          0,
			},
			want:    "msr\trmr_el3, x12",
			wantErr: false,
		},
		{
			name: "msr	contextidr_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0xd0, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tcontextidr_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	tpidr_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0xd0, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\ttpidr_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	tpidr_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0xd0, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\ttpidr_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	tpidr_el3, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0xd0, 0x1e, 0xd5}),
				address:          0,
			},
			want:    "msr\ttpidr_el3, x12",
			wantErr: false,
		},
		{
			name: "msr	tpidrro_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6c, 0xd0, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\ttpidrro_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	tpidr_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0xd0, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\ttpidr_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	cntfrq_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xe0, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tcntfrq_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	cntvoff_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6c, 0xe0, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tcntvoff_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	cntkctl_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xe1, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\tcnthctl_el21, x12",
			wantErr: false,
		},
		{
			name: "msr	cnthctl_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xe1, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tcnthctl_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	cntp_tval_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xe2, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tcnthp_tval_el21, x12",
			wantErr: false,
		},
		{
			name: "msr	cnthp_tval_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xe2, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tcnthp_tval_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	cntps_tval_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xe2, 0x1f, 0xd5}),
				address:          0,
			},
			want:    "msr\tcntps_tval_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	cntp_ctl_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0xe2, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tcnthp_ctl_el21, x12",
			wantErr: false,
		},
		{
			name: "msr	cnthp_ctl_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0xe2, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tcnthp_ctl_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	cntps_ctl_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0xe2, 0x1f, 0xd5}),
				address:          0,
			},
			want:    "msr\tcntps_ctl_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	cntp_cval_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0xe2, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tcnthp_cval_el21, x12",
			wantErr: false,
		},
		{
			name: "msr	cnthp_cval_el2, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0xe2, 0x1c, 0xd5}),
				address:          0,
			},
			want:    "msr\tcnthp_cval_el2, x12",
			wantErr: false,
		},
		{
			name: "msr	cntps_cval_el1, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0xe2, 0x1f, 0xd5}),
				address:          0,
			},
			want:    "msr\tcntps_cval_el1, x12",
			wantErr: false,
		},
		{
			name: "msr	cntv_tval_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xe3, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tcnthv_tval_el21, x12",
			wantErr: false,
		},
		{
			name: "msr	cntv_ctl_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0xe3, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tcnthv_ctl_el21, x12",
			wantErr: false,
		},
		{
			name: "msr	cntv_cval_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0xe3, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tcnthv_cval_el21, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr0_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xe8, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr0_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr1_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0xe8, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr1_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr2_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0xe8, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr2_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr3_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6c, 0xe8, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr3_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr4_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0xe8, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr4_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr5_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0xe8, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr5_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr6_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0xe8, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr6_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr7_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0xe8, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr7_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr8_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xe9, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr8_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr9_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0xe9, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr9_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr10_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0xe9, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr10_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr11_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6c, 0xe9, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr11_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr12_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0xe9, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr12_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr13_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0xe9, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr13_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr14_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0xe9, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr14_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr15_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0xe9, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr15_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr16_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xea, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr16_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr17_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0xea, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr17_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr18_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0xea, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr18_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr19_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6c, 0xea, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr19_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr20_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0xea, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr20_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr21_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0xea, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr21_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr22_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0xea, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr22_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr23_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0xea, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr23_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr24_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xeb, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr24_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr25_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0xeb, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr25_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr26_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0xeb, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr26_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr27_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6c, 0xeb, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr27_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr28_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0xeb, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr28_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr29_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0xeb, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr29_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevcntr30_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0xeb, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevcntr30_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmccfiltr_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0xef, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmccfiltr_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper0_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xec, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper0_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper1_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0xec, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper1_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper2_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0xec, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper2_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper3_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6c, 0xec, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper3_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper4_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0xec, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper4_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper5_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0xec, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper5_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper6_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0xec, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper6_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper7_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0xec, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper7_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper8_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xed, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper8_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper9_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0xed, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper9_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper10_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0xed, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper10_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper11_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6c, 0xed, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper11_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper12_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0xed, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper12_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper13_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0xed, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper13_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper14_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0xed, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper14_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper15_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0xed, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper15_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper16_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xee, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper16_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper17_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0xee, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper17_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper18_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0xee, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper18_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper19_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6c, 0xee, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper19_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper20_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0xee, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper20_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper21_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0xee, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper21_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper22_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0xee, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper22_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper23_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xec, 0xee, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper23_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper24_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xef, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper24_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper25_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2c, 0xef, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper25_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper26_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x4c, 0xef, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper26_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper27_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x6c, 0xef, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper27_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper28_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x8c, 0xef, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper28_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper29_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0xef, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper29_el0, x12",
			wantErr: false,
		},
		{
			name: "msr	pmevtyper30_el0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xcc, 0xef, 0x1b, 0xd5}),
				address:          0,
			},
			want:    "msr\tpmevtyper30_el0, x12",
			wantErr: false,
		},
		{
			name: "mrs	x9, teecr32_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x00, 0x32, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, teecr32_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, osdtrrx_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x00, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, osdtrrx_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, mdccsr_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x01, 0x33, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, mdccsr_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, mdccint_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x02, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, mdccint_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, mdscr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x02, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, mdscr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, osdtrtx_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x03, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, osdtrtx_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgdtr_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x04, 0x33, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgdtr_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgdtrrx_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x05, 0x33, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgdtrrx_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, oseccr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x06, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, oseccr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgvcr32_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x07, 0x34, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgvcr32_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbvr0_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x00, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbvr0_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbvr1_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x01, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbvr1_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbvr2_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x02, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbvr2_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbvr3_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x03, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbvr3_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbvr4_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x04, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbvr4_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbvr5_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x05, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbvr5_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbvr6_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x06, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbvr6_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbvr7_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x07, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbvr7_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbvr8_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x08, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbvr8_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbvr9_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x09, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbvr9_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbvr10_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x0a, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbvr10_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbvr11_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x0b, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbvr11_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbvr12_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x0c, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbvr12_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbvr13_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x0d, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbvr13_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbvr14_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x0e, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbvr14_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbvr15_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x0f, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbvr15_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbcr0_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x00, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbcr0_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbcr1_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x01, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbcr1_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbcr2_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x02, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbcr2_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbcr3_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x03, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbcr3_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbcr4_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x04, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbcr4_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbcr5_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x05, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbcr5_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbcr6_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x06, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbcr6_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbcr7_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x07, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbcr7_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbcr8_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x08, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbcr8_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbcr9_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x09, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbcr9_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbcr10_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x0a, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbcr10_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbcr11_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x0b, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbcr11_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbcr12_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x0c, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbcr12_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbcr13_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x0d, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbcr13_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbcr14_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x0e, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbcr14_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgbcr15_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x0f, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgbcr15_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwvr0_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x00, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwvr0_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwvr1_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x01, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwvr1_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwvr2_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x02, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwvr2_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwvr3_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x03, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwvr3_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwvr4_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x04, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwvr4_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwvr5_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x05, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwvr5_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwvr6_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x06, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwvr6_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwvr7_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x07, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwvr7_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwvr8_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x08, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwvr8_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwvr9_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x09, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwvr9_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwvr10_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x0a, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwvr10_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwvr11_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x0b, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwvr11_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwvr12_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x0c, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwvr12_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwvr13_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x0d, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwvr13_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwvr14_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x0e, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwvr14_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwvr15_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x0f, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwvr15_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwcr0_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x00, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwcr0_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwcr1_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x01, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwcr1_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwcr2_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x02, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwcr2_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwcr3_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x03, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwcr3_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwcr4_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x04, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwcr4_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwcr5_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x05, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwcr5_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwcr6_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x06, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwcr6_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwcr7_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x07, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwcr7_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwcr8_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x08, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwcr8_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwcr9_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x09, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwcr9_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwcr10_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x0a, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwcr10_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwcr11_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x0b, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwcr11_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwcr12_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x0c, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwcr12_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwcr13_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x0d, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwcr13_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwcr14_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x0e, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwcr14_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgwcr15_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x0f, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgwcr15_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, mdrar_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x10, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, mdrar_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, teehbr32_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x10, 0x32, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, teehbr32_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, oslsr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x11, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, oslsr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, osdlr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x13, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, osdlr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgprcr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x14, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgprcr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgclaimset_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x78, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgclaimset_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgclaimclr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x79, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgclaimclr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dbgauthstat_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x7e, 0x30, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dbgauthstat_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, midr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x00, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, midr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, ccsidr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x00, 0x39, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, ccsidr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, csselr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x00, 0x3a, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, csselr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, vpidr_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x00, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, vpidr_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, clidr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x00, 0x39, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, clidr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, ctr_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x00, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, ctr_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, mpidr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x00, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, mpidr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, vmpidr_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x00, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, vmpidr_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, revidr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x00, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, revidr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, aidr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x00, 0x39, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, aidr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, dczid_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x00, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dczid_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_pfr0_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x01, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_pfr0_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_pfr1_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x01, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_pfr1_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_dfr0_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x01, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_dfr0_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_afr0_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0x01, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_afr0_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_mmfr0_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x01, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_mmfr0_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_mmfr1_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x01, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_mmfr1_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_mmfr2_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x01, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_mmfr2_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_mmfr3_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x01, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_mmfr3_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_mmfr4_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x02, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_mmfr4_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_mmfr5_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x03, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_aa32res6_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_isar0_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x02, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_isar0_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_isar1_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x02, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_isar1_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_isar2_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x02, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_isar2_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_isar3_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0x02, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_isar3_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_isar4_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x02, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_isar4_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_isar5_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x02, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_isar5_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, mvfr0_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x03, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, mvfr0_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, mvfr1_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x03, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, mvfr1_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, mvfr2_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x03, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, mvfr2_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_aa64pfr0_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x04, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_aa64pfr0_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_aa64pfr1_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x04, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_aa64pfr1_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_aa64dfr0_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x05, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_aa64dfr0_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_aa64dfr1_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x05, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_aa64dfr1_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_aa64afr0_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x05, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_aa64afr0_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_aa64afr1_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x05, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_aa64afr1_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_aa64isar0_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x06, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_aa64isar0_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_aa64isar1_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x06, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_aa64isar1_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_aa64mmfr0_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x07, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_aa64mmfr0_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, id_aa64mmfr1_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x07, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, id_aa64mmfr1_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, sctlr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x10, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, sctlr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, sctlr_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x10, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, sctlr_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, sctlr_el3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x10, 0x3e, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, sctlr_el3",
			wantErr: false,
		},
		{
			name: "mrs	x9, actlr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x10, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, actlr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, actlr_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x10, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, actlr_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, actlr_el3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x10, 0x3e, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, actlr_el3",
			wantErr: false,
		},
		{
			name: "mrs	x9, cpacr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x10, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, cpacr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, hcr_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x11, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, hcr_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, scr_el3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x11, 0x3e, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, scr_el3",
			wantErr: false,
		},
		{
			name: "mrs	x9, mdcr_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x11, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, mdcr_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, sder32_el3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x11, 0x3e, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, sder32_el3",
			wantErr: false,
		},
		{
			name: "mrs	x9, cptr_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x11, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, cptr_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, cptr_el3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x11, 0x3e, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, cptr_el3",
			wantErr: false,
		},
		{
			name: "mrs	x9, hstr_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0x11, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, hstr_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, hacr_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x11, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, hacr_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, mdcr_el3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x13, 0x3e, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, mdcr_el3",
			wantErr: false,
		},
		{
			name: "mrs	x9, ttbr0_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x20, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, ttbr0_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, ttbr0_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x20, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, ttbr0_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, ttbr0_el3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x20, 0x3e, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, ttbr0_el3",
			wantErr: false,
		},
		{
			name: "mrs	x9, ttbr1_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x20, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, ttbr1_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, tcr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x20, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, tcr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, tcr_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x20, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, tcr_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, tcr_el3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x20, 0x3e, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, tcr_el3",
			wantErr: false,
		},
		{
			name: "mrs	x9, vttbr_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x21, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, vttbr_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, vtcr_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x21, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, vtcr_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, dacr32_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x30, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dacr32_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, spsr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x40, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, spsr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, spsr_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x40, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, spsr_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, spsr_el3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x40, 0x3e, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, spsr_el3",
			wantErr: false,
		},
		{
			name: "mrs	x9, elr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x40, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, elr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, elr_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x40, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, elr_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, elr_el3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x40, 0x3e, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, elr_el3",
			wantErr: false,
		},
		{
			name: "mrs	x9, sp_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x41, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, sp_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, sp_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x41, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, sp_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, sp_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x41, 0x3e, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, sp_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, spsel",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x42, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, spsel",
			wantErr: false,
		},
		{
			name: "mrs	x9, nzcv",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x42, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, nzcv",
			wantErr: false,
		},
		{
			name: "mrs	x9, daif",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x42, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, daif",
			wantErr: false,
		},
		{
			name: "mrs	x9, currentel",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x42, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, currentel",
			wantErr: false,
		},
		{
			name: "mrs	x9, spsr_irq",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x43, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, spsr_irq",
			wantErr: false,
		},
		{
			name: "mrs	x9, spsr_abt",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x43, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, spsr_abt",
			wantErr: false,
		},
		{
			name: "mrs	x9, spsr_und",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x43, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, spsr_und",
			wantErr: false,
		},
		{
			name: "mrs	x9, spsr_fiq",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0x43, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, spsr_fiq",
			wantErr: false,
		},
		{
			name: "mrs	x9, fpcr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x44, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, fpcr",
			wantErr: false,
		},
		{
			name: "mrs	x9, fpsr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x44, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, fpsr",
			wantErr: false,
		},
		{
			name: "mrs	x9, dspsr_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x45, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dspsr",
			wantErr: false,
		},
		{
			name: "mrs	x9, dlr_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x45, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, dlr",
			wantErr: false,
		},
		{
			name: "mrs	x9, ifsr32_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x50, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, ifsr32_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, afsr0_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x51, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, afsr0_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, afsr0_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x51, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, afsr0_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, afsr0_el3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x51, 0x3e, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, afsr0_el3",
			wantErr: false,
		},
		{
			name: "mrs	x9, afsr1_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x51, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, afsr1_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, afsr1_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x51, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, afsr1_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, afsr1_el3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x51, 0x3e, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, afsr1_el3",
			wantErr: false,
		},
		{
			name: "mrs	x9, esr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x52, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, esr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, esr_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x52, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, esr_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, esr_el3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x52, 0x3e, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, esr_el3",
			wantErr: false,
		},
		{
			name: "mrs	x9, fpexc32_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x53, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, fpexc32_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, far_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x60, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, far_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, far_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x60, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, far_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, far_el3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x60, 0x3e, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, far_el3",
			wantErr: false,
		},
		{
			name: "mrs	x9, hpfar_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0x60, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, hpfar_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, par_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x74, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, par_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmcr_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x9c, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmcr_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmcntenset_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x9c, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmcntenset_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmcntenclr_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x9c, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmcntenclr_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmovsclr_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0x9c, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmovsclr_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmselr_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0x9c, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmselr_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmceid0_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0x9c, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, s3_3_c9_c12_6",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmceid1_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0x9c, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, s3_3_c9_c12_7",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmccntr_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x9d, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmccntr_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmxevtyper_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x9d, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmxevtyper_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmxevcntr_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x9d, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmxevcntr_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmuserenr_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0x9e, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmuserenr_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmintenset_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0x9e, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmintenset_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmintenclr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0x9e, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmintenclr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmovsset_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0x9e, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmovsset_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, mair_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xa2, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, mair_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, mair_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xa2, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, mair_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, mair_el3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xa2, 0x3e, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, mair_el3",
			wantErr: false,
		},
		{
			name: "mrs	x9, amair_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xa3, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, amair_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, amair_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xa3, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, amair_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, amair_el3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xa3, 0x3e, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, amair_el3",
			wantErr: false,
		},
		{
			name: "mrs	x9, vbar_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xc0, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, vbar_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, vbar_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xc0, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, vbar_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, vbar_el3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xc0, 0x3e, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, vbar_el3",
			wantErr: false,
		},
		{
			name: "mrs	x9, rvbar_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0xc0, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, rvbar_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, rvbar_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0xc0, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, rvbar_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, rvbar_el3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0xc0, 0x3e, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, rvbar_el3",
			wantErr: false,
		},
		{
			name: "mrs	x9, rmr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xc0, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, rmr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, rmr_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xc0, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, rmr_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, rmr_el3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xc0, 0x3e, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, rmr_el3",
			wantErr: false,
		},
		{
			name: "mrs	x9, isr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xc1, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, isr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, contextidr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0xd0, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, contextidr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, tpidr_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xd0, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, tpidr_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, tpidr_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xd0, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, tpidr_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, tpidr_el3",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xd0, 0x3e, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, tpidr_el3",
			wantErr: false,
		},
		{
			name: "mrs	x9, tpidrro_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0xd0, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, tpidrro_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, tpidr_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0xd0, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, tpidr_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, cntfrq_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xe0, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, cntfrq_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, cntpct_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0xe0, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, cntpct_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, cntvct_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xe0, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, cntvct_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, cntvoff_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0xe0, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, cntvoff_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, cntkctl_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xe1, 0x38, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, cnthctl_el21",
			wantErr: false,
		},
		{
			name: "mrs	x9, cnthctl_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xe1, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, cnthctl_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, cntp_tval_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xe2, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, cnthp_tval_el21",
			wantErr: false,
		},
		{
			name: "mrs	x9, cnthp_tval_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xe2, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, cnthp_tval_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, cntps_tval_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xe2, 0x3f, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, cntps_tval_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, cntp_ctl_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0xe2, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, cnthp_ctl_el21",
			wantErr: false,
		},
		{
			name: "mrs	x9, cnthp_ctl_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0xe2, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, cnthp_ctl_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, cntps_ctl_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0xe2, 0x3f, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, cntps_ctl_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, cntp_cval_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xe2, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, cnthp_cval_el21",
			wantErr: false,
		},
		{
			name: "mrs	x9, cnthp_cval_el2",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xe2, 0x3c, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, cnthp_cval_el2",
			wantErr: false,
		},
		{
			name: "mrs	x9, cntps_cval_el1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xe2, 0x3f, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, cntps_cval_el1",
			wantErr: false,
		},
		{
			name: "mrs	x9, cntv_tval_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xe3, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, cnthv_tval_el21",
			wantErr: false,
		},
		{
			name: "mrs	x9, cntv_ctl_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0xe3, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, cnthv_ctl_el21",
			wantErr: false,
		},
		{
			name: "mrs	x9, cntv_cval_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xe3, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, cnthv_cval_el21",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr0_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xe8, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr0_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr1_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0xe8, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr1_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr2_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xe8, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr2_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr3_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0xe8, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr3_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr4_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0xe8, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr4_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr5_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0xe8, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr5_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr6_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0xe8, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr6_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr7_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xe8, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr7_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr8_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xe9, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr8_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr9_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0xe9, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr9_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr10_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xe9, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr10_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr11_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0xe9, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr11_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr12_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0xe9, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr12_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr13_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0xe9, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr13_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr14_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0xe9, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr14_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr15_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xe9, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr15_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr16_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xea, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr16_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr17_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0xea, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr17_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr18_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xea, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr18_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr19_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0xea, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr19_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr20_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0xea, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr20_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr21_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0xea, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr21_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr22_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0xea, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr22_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr23_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xea, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr23_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr24_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xeb, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr24_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr25_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0xeb, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr25_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr26_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xeb, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr26_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr27_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0xeb, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr27_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr28_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0xeb, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr28_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr29_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0xeb, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr29_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevcntr30_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0xeb, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevcntr30_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmccfiltr_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xef, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmccfiltr_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper0_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xec, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper0_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper1_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0xec, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper1_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper2_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xec, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper2_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper3_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0xec, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper3_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper4_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0xec, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper4_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper5_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0xec, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper5_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper6_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0xec, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper6_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper7_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xec, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper7_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper8_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xed, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper8_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper9_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0xed, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper9_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper10_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xed, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper10_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper11_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0xed, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper11_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper12_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0xed, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper12_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper13_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0xed, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper13_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper14_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0xed, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper14_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper15_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xed, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper15_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper16_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xee, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper16_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper17_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0xee, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper17_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper18_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xee, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper18_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper19_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0xee, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper19_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper20_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0xee, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper20_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper21_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0xee, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper21_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper22_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0xee, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper22_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper23_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe9, 0xee, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper23_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper24_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x09, 0xef, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper24_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper25_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x29, 0xef, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper25_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper26_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x49, 0xef, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper26_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper27_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x69, 0xef, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper27_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper28_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x89, 0xef, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper28_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper29_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xa9, 0xef, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper29_el0",
			wantErr: false,
		},
		{
			name: "mrs	x9, pmevtyper30_el0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc9, 0xef, 0x3b, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx9, pmevtyper30_el0",
			wantErr: false,
		},
		{
			name: "mrs	x12, s3_7_c15_c1_5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xac, 0xf1, 0x3f, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx12, upmc9",
			wantErr: false,
		},
		{
			name: "mrs	x13, s3_2_c11_c15_7",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xed, 0xbf, 0x3a, 0xd5}),
				address:          0,
			},
			want:    "mrs\tx13, s3_2_c11_c15_7",
			wantErr: false,
		},
		{
			name: "sysl	x14, #3, c9, c2, #1",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x2e, 0x92, 0x2b, 0xd5}),
				address:          0,
			},
			want:    "sysl\tx14, #0x3, c9, c2, #0x1",
			wantErr: false,
		},
		{
			name: "msr	s3_0_c15_c0_0, x12",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x0c, 0xf0, 0x18, 0xd5}),
				address:          0,
			},
			want:    "msr\thid0, x12",
			wantErr: false,
		},
		{
			name: "msr	s3_7_c11_c13_7, x5",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe5, 0xbd, 0x1f, 0xd5}),
				address:          0,
			},
			want:    "msr\ts3_7_c11_c13_7, x5",
			wantErr: false,
		},
		{
			name: "sys	#3, c9, c2, #1, x4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x24, 0x92, 0x0b, 0xd5}),
				address:          0,
			},
			want:    "sys\t#0x3, c9, c2, #0x1, x4",
			wantErr: false,
		},
		{
			name: "b	#4",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x01, 0x00, 0x00, 0x14}),
				address:          0,
			},
			want:    "b\t0x4",
			wantErr: false,
		},
		{
			name: "bl	#0",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x00, 0x00, 0x94}),
				address:          0,
			},
			want:    "bl\t0x0",
			wantErr: false,
		},
		{
			name: "b	#134217724",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xff, 0xff, 0xff, 0x15}),
				address:          0,
			},
			want:    "b\t0x7fffffc",
			wantErr: false,
		},
		{
			name: "bl	#-134217728",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x00, 0x00, 0x00, 0x96}),
				address:          0,
			},
			want:    "bl\t0xfffffffff8000000",
			wantErr: false,
		},
		{
			name: "br	x20",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x80, 0x02, 0x1f, 0xd6}),
				address:          0,
			},
			want:    "br\tx20",
			wantErr: false,
		},
		{
			name: "blr	xzr",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x03, 0x3f, 0xd6}),
				address:          0,
			},
			want:    "blr\txzr",
			wantErr: false,
		},
		{
			name: "ret	x10",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0x40, 0x01, 0x5f, 0xd6}),
				address:          0,
			},
			want:    "ret\tx10",
			wantErr: false,
		},
		{
			name: "ret",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xc0, 0x03, 0x5f, 0xd6}),
				address:          0,
			},
			want:    "ret",
			wantErr: false,
		},
		{
			name: "eret",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x03, 0x9f, 0xd6}),
				address:          0,
			},
			want:    "eret",
			wantErr: false,
		},
		{
			name: "drps",
			args: args{
				instructionValue: binary.LittleEndian.Uint32([]byte{0xe0, 0x03, 0xbf, 0xd6}),
				address:          0,
			},
			want:    "drps",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decompose(tt.args.address, tt.args.instructionValue, &results)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decompose() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got.Disassembly, tt.want) {
				t.Errorf("Disassembly = %v, want %v", got.Disassembly, tt.want)
			}
		})
	}
}
