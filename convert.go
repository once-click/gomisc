package gomisc

import (
	"encoding/binary"
	"math"
	"strconv"
	"time"
)

func ByteToFloat32(b []byte) float32 {
	bits := binary.LittleEndian.Uint32(b)
	return math.Float32frombits(bits)
}

func ByteToFloat64(b []byte) float64 {
	bits := binary.LittleEndian.Uint64(b)
	return math.Float64frombits(bits)
}

func Float32ToByte(f float32) []byte {
	bits := math.Float32bits(f)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

func Float64ToByte(f float64) []byte {
	bits := math.Float64bits(f)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func TimeStr2StampStr(s string) string {
	//t, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	t, _ := time.Parse("2006-01-02 15:04:05", s)
	return strconv.FormatInt(t.Unix()-28800, 10)
}

func TimeStampStr2Str(s string) string {
	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return ""
	}
	tm := time.Unix(i64, 0)
	return tm.Format("2006-01-02 15:04:05")
}
