package macho

const (
	// Magic32 is the Mach-O 32-bit magic number in big-endian
	Magic32 uint32 = 0xFEEDFACE
	// Cigam32 is the Mach-O 32-bit magic number in little-endian
	Cigam32 uint32 = 0xCEFAEDFE
	// Magic64 is the Mach-O 64-bit magic number in big-endian
	Magic64 uint32 = 0xFEEDFACF
	// Cigam64 is the 64-bit magic number in little-endian
	Cigam64 uint32 = 0xCFFAEDFE
)

type FileHeader struct {
	Magic      uint32
	CpuType    Cpu
	CpuSubType uint32
	Type       Type
	Ncmds      uint32
	SizeOfCmds uint32
	Flags      uint32
}

const (
	FileHeaderSize32 = 7 * 4 // 7 4-byte fields
	FileHeaderSize64 = 8 * 4 // 8 4-byte fields (additional field is reserved)
)

type Cpu uint32

const arch64 = 0x1000000

const (
	Cpu386   Cpu = 7
	CpuAmd64 Cpu = (Cpu386 | arch64)
	CpuArm   Cpu = 12
	CpuArm64 Cpu = (CpuArm | arch64)
	CpuPpc   Cpu = 18
	CpuPpc64 Cpu = (CpuPpc | arch64)
)

type Type uint32

const (
	TypeObject     Type = 1
	TypeExecutable Type = 2
	TypeDylib      Type = 6
	TypeBundle     Type = 8
)
