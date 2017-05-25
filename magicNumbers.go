package hideFile

type magicNumber struct {
	Name        string
	Extension   string
	Description string
	Number      []byte
	Offset      offset
}

type offset struct {
	Count int64
	Value []byte
}

var magicNumberList = []magicNumber{
	{
		Name:        "JPEG",
		Extension:   "jpg",
		Description: "JPEG raw",
		Number:      []byte{0xFF, 0xD8, 0xFF, 0xDB},
	},
	{
		Name:        "7ZIP",
		Extension:   "7z",
		Description: "7-Zip File Format",
		Number:      []byte{0x37, 0x7A, 0xBC, 0xAF, 0x27, 0x1C},
	},
	{
		Name:        "GZIP",
		Extension:   "gz",
		Description: "GZIP",
		Number:      []byte{0x1F, 0x8B},
	},
	{
		Name:        "DMG",
		Extension:   "dmg",
		Description: "Apple Disk Image",
		Number:      []byte{0x78, 0x01, 0x73, 0x0D, 0x62, 0x62, 0x60},
	},
	{
		Name:        "PDF",
		Extension:   "pdf",
		Description: "PDF document",
		Number:      []byte{0x25, 0x50, 0x44, 0x46},
	},
}
