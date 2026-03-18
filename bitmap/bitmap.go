package bitmap

// Bitmap 位图结构，用[]byte存储位数据
type Bitmap struct {
	data []byte
}

// NewBitmap 创建一个空的Bitmap
func NewBitmap() *Bitmap {
	return &Bitmap{
		data: make([]byte, 0),
	}
}

// Set 将指定ID的位设为1（标记：在线/已打卡）
func (b *Bitmap) Set(id uint64) {
	// 1. 计算字节索引和位偏移
	byteIdx := id / 8
	bitOffset := id % 8

	// 2. 确保数组长度足够（避免索引越界）
	for uint64(len(b.data)) <= byteIdx {
		b.data = append(b.data, 0)
	}

	// 3. 位运算设为1
	b.data[byteIdx] |= 1 << bitOffset
}

// Unset 将指定ID的位设为0（标记：离线/未打卡）
func (b *Bitmap) Unset(id uint64) {
	byteIdx := id / 8
	bitOffset := id % 8
	// 数组长度不足时，无需操作（默认就是0）
	if uint64(len(b.data)) <= byteIdx {
		return
	}
	// 位运算设为0
	b.data[byteIdx] &= ^(1 << bitOffset)
}

// IsSet 查询指定ID的位是否为1（是否在线/已打卡）
func (b *Bitmap) IsSet(id uint64) bool {
	byteIdx := id / 8
	bitOffset := id % 8

	// 数组长度不足时，返回false
	if uint64(len(b.data)) <= byteIdx {
		return false
	}

	// 位运算查询
	return (b.data[byteIdx]>>bitOffset)&1 == 1
}

// Count 统计所有为1的位的数量（在线人数/打卡人数）
func (b *Bitmap) Count() uint64 {
	var count uint64
	// 遍历每个字节，统计其中1的个数
	for _, b := range b.data {
		// 快速统计一个字节中1的个数（Go内置技巧）
		count += uint64(popCount(b))
	}
	return count
}

// popCount 统计单个字节中1的位数（辅助函数）
func popCount(b byte) int {
	count := 0
	for b != 0 {
		count++
		b &= b - 1 // 清除最低位的1
	}
	return count
}
