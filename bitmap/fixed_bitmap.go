package bitmap

import (
	"errors"
)

// FixedBitmap 固定大小的位图，初始化时指定最大位数量，不自动扩容
type FixedBitmap struct {
	bits []byte
	size uint64 // 最大支持的位数量
}

// NewFixedBitmap 创建固定大小的Bitmap
// size: 最大支持的位数量（比如最大用户ID）
func NewFixedBitmap(size uint64) *FixedBitmap {
	if size == 0 {
		panic("size must be greater than 0")
	}
	byteLen := (size + 7) / 8
	return &FixedBitmap{
		bits: make([]byte, byteLen),
		size: size,
	}
}

// checkBounds 边界检查
func (f *FixedBitmap) checkBounds(offset uint64) error {
	if offset >= f.size {
		return errors.New("offset out of bitmap bounds")
	}
	return nil
}

// Set 设置offset位为1
func (f *FixedBitmap) Set(offset uint64) error {
	if err := f.checkBounds(offset); err != nil {
		return err
	}
	byteIdx := offset / 8
	bitOffset := offset % 8
	f.bits[byteIdx] |= 1 << bitOffset
	return nil
}

// Unset 设置offset位为0
func (f *FixedBitmap) Unset(offset uint64) error {
	if err := f.checkBounds(offset); err != nil {
		return err
	}
	byteIdx := offset / 8
	bitOffset := offset % 8
	f.bits[byteIdx] &^= 1 << bitOffset
	return nil
}

// IsSet 判断offset位是否为1
func (f *FixedBitmap) IsSet(offset uint64) (bool, error) {
	if err := f.checkBounds(offset); err != nil {
		return false, err
	}
	byteIdx := offset / 8
	bitOffset := offset % 8
	return (f.bits[byteIdx]>>bitOffset)&1 == 1, nil
}

// Count 统计1的位数
func (f *FixedBitmap) Count() uint64 {
	var count uint64
	for _, b := range f.bits {
		count += uint64(countOneBits(b))
	}
	return count
}

// Size 返回位图最大位数量
func (f *FixedBitmap) Size() uint64 {
	return f.size
}

// countOneBits 统计单个字节中1的个数
func countOneBits(b byte) int {
	cnt := 0
	for b != 0 {
		cnt++
		b &= b - 1
	}
	return cnt
}
