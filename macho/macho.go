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
