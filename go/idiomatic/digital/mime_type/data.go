package mime_type

import "fmt"

var all = []MimeTypeModel{
	{
		Name:           "Portable Document Format",
		MimeTypes:      []string{"application/pdf"},
		FileExtensions: []string{".pdf"},
	},
	{
		Name:           "Plain Text",
		MimeTypes:      []string{"text/plain"},
		FileExtensions: []string{".txt", ".text"},
	},
	{
		Name:           "Microsoft Word",
		MimeTypes:      []string{"application/msword"},
		FileExtensions: []string{".doc", ".dot", ".docx", ".dotx", ".docm", ".dotm"},
	},
	{
		Name: "Microsoft Word",
		MimeTypes: []string{
			"application/msword",
			"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		},
		FileExtensions: []string{".doc", ".docx"},
	},
	{
		Name:           "Portable Network Graphics",
		Abbreviations:  []string{"PNG"},
		FileExtensions: []string{".png"},
		MimeTypes:      []string{"image/png"},
	},
	{
		Name:           "Joint Photographic Experts Group",
		Abbreviations:  []string{"JPEG"},
		FileExtensions: []string{".jpg", ".jpeg"},
		MimeTypes:      []string{"image/jpeg", "image/pjpeg"},
	},
	{
		Name:           "Graphics Interchange Format",
		Abbreviations:  []string{"GIF"},
		FileExtensions: []string{".gif"},
		MimeTypes:      []string{"image/gif"},
	},
	{
		Name:           "Scalable Vector Graphics",
		Abbreviations:  []string{"SVG"},
		FileExtensions: []string{".svg"},
		MimeTypes:      []string{"image/svg+xml"},
	},
	{
		Name:           "Tagged Image File Format",
		Abbreviations:  []string{"TIFF"},
		FileExtensions: []string{".tiff", ".tif"},
		MimeTypes:      []string{"image/tiff", "image/x-tiff"},
	},
	{
		Name:           "Microsoft Excel (Binary Workbook)",
		Abbreviations:  []string{"XLS"},
		FileExtensions: []string{".xls"},
		MimeTypes:      []string{"application/vnd.ms-excel", "application/x-msexcel"},
	},
	{
		Name:           "Microsoft Excel (Open XML Workbook)",
		Abbreviations:  []string{"XLSX"},
		FileExtensions: []string{".xlsx"},
		MimeTypes:      []string{"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"},
	},
	{

		Name:           "Secure Hash Algorithm 1 Checksum",
		Abbreviations:  []string{"SHA-1"},
		FileExtensions: []string{".sha1", ".sha1sum"},
		MimeTypes:      []string{"application/octet-stream", "application/x-sha-1"},
	},
	{
		Name:           "Secure Hash Algorithm 224 Checksum",
		Abbreviations:  []string{"SHA-224"},
		FileExtensions: []string{".sha224", ".sha224sum"},
		MimeTypes:      []string{"application/octet-stream", "application/x-sha-224"},
	},
	{
		Name:           "Secure Hash Algorithm 256 Checksum",
		Abbreviations:  []string{"SHA-256"},
		FileExtensions: []string{".sha256", ".sha256sum"},
		MimeTypes:      []string{"application/octet-stream", "application/x-sha-256"},
	},
	{
		Name:           "Secure Hash Algorithm 384 Checksum",
		Abbreviations:  []string{"SHA-384"},
		FileExtensions: []string{".sha384", ".sha384sum"},
		MimeTypes:      []string{"application/octet-stream", "application/x-sha-384"},
	},
	{
		Name:           "Secure Hash Algorithm 512 Checksum",
		Abbreviations:  []string{"SHA-512"},
		FileExtensions: []string{".sha512", ".sha512sum"},
		MimeTypes:      []string{"application/octet-stream", "application/x-sha-512"},
	},
	{
		Name:           "Microsoft PowerPoint (Binary Presentation)",
		Abbreviations:  []string{"PPT"},
		FileExtensions: []string{".ppt"},
		MimeTypes:      []string{"application/vnd.ms-powerpoint", "application/powerpoint", "application/x-mspowerpoint"},
	},
	{
		Name:           "Microsoft PowerPoint (Open XML Presentation)",
		Abbreviations:  []string{"PPTX"},
		FileExtensions: []string{".pptx"},
		MimeTypes:      []string{"application/vnd.openxmlformats-officedocument.presentationml.presentation"},
	},
	{
		Name:           "Microsoft PowerPoint Template (Binary)",
		Abbreviations:  []string{"POT"},
		FileExtensions: []string{".pot"},
		MimeTypes:      []string{"application/vnd.ms-powerpoint"},
	},
	{
		Name:           "Microsoft PowerPoint Template (Open XML)",
		Abbreviations:  []string{"POTX"},
		FileExtensions: []string{".potx"},
		MimeTypes:      []string{"application/vnd.openxmlformats-officedocument.presentationml.template"},
	},
	{
		Name:           "Microsoft PowerPoint (Open XML Presentation, Macro-Enabled)",
		Abbreviations:  []string{"PPTM"},
		FileExtensions: []string{".pptm"},
		MimeTypes:      []string{"application/vnd.ms-powerpoint.presentation.macroEnabled.12"},
	},
	{
		Name:           "Microsoft PowerPoint Template (Open XML, Macro-Enabled)",
		Abbreviations:  []string{"POTM"},
		FileExtensions: []string{".potm"},
		MimeTypes:      []string{"application/vnd.ms-powerpoint.template.macroEnabled.12"},
	},
	{
		Name:           "OpenDocument Text",
		Abbreviations:  []string{"ODT"},
		FileExtensions: []string{".odt"},
		MimeTypes:      []string{"application/vnd.oasis.opendocument.text"},
	},
	{
		Name:           "OpenDocument Text Template",
		Abbreviations:  []string{"OTT"},
		FileExtensions: []string{".ott"},
		MimeTypes:      []string{"application/vnd.oasis.opendocument.text-template"},
	},
	{
		Name:           "OpenDocument Spreadsheet",
		Abbreviations:  []string{"ODS"},
		FileExtensions: []string{".ods"},
		MimeTypes:      []string{"application/vnd.oasis.opendocument.spreadsheet"},
	},
	{
		Name:           "OpenDocument Spreadsheet Template",
		Abbreviations:  []string{"OTS"},
		FileExtensions: []string{".ots"},
		MimeTypes:      []string{"application/vnd.oasis.opendocument.spreadsheet-template"},
	},
	{
		Name:           "OpenDocument Presentation",
		Abbreviations:  []string{"ODP"},
		FileExtensions: []string{".odp"},
		MimeTypes:      []string{"application/vnd.oasis.opendocument.presentation"},
	},
	{
		Name:           "OpenDocument Presentation Template",
		Abbreviations:  []string{"OTP"},
		FileExtensions: []string{".otp"},
		MimeTypes:      []string{"application/vnd.oasis.opendocument.presentation-template"},
	},
	{
		Name:           "OpenDocument Graphics",
		Abbreviations:  []string{"ODG"},
		FileExtensions: []string{".odg"},
		MimeTypes:      []string{"application/vnd.oasis.opendocument.graphics"},
	},
	{
		Name:           "OpenDocument Graphics Template",
		Abbreviations:  []string{"OTG"},
		FileExtensions: []string{".otg"},
		MimeTypes:      []string{"application/vnd.oasis.opendocument.graphics-template"},
	},
	{
		Name:           "Markdown",
		Abbreviations:  []string{"MD"},
		FileExtensions: []string{".md", ".markdown", ".mdown", ".mkd", ".mkdn", ".mdwn", ".mdtxt"},
		MimeTypes:      []string{"text/markdown", "text/x-markdown"},
	},
	{
		Name:           "MD5 Checksum",
		Abbreviations:  []string{"MD5"},
		FileExtensions: []string{".md5", ".md5sum"},
		MimeTypes:      []string{"application/octet-stream", "application/x-md5"},
	},
	{
		Name:           "ZIP Archive",
		Abbreviations:  []string{"ZIP"},
		FileExtensions: []string{".zip"},
		MimeTypes:      []string{"application/zip"},
	},
	{
		Name:           "7-Zip Archive",
		Abbreviations:  []string{"7Z"},
		FileExtensions: []string{".7z"},
		MimeTypes:      []string{"application/x-7z-compressed", "application/7z"},
	},
	{
		Name:           "TAR Archive",
		Abbreviations:  []string{"TAR"},
		FileExtensions: []string{".tar"},
		MimeTypes:      []string{"application/x-tar"},
	},
	{
		Name:           "Gzip Compressed",
		Abbreviations:  []string{"GZIP"},
		FileExtensions: []string{".gz"},
		MimeTypes:      []string{"application/gzip", "application/x-gzip"},
	},
	{
		Name:           "Bzip2 Compressed",
		Abbreviations:  []string{"BZIP2"},
		FileExtensions: []string{".bz2", ".bzip2"},
		MimeTypes:      []string{"application/x-bzip2"},
	},
	{
		Name:           "XZ Compressed / LZMA2",
		Abbreviations:  []string{"XZ", "LZMA2"},
		FileExtensions: []string{".xz", ".lzma"},
		MimeTypes:      []string{"application/x-xz", "application/x-lzma"},
	},
	{
		Name:           "Tarball (gzip) / .tgz",
		Abbreviations:  []string{"TAR.GZ", "TGZ"},
		FileExtensions: []string{".tar.gz", ".tgz"},
		MimeTypes:      []string{"application/gzip", "application/x-gtar", "application/x-tgz"},
	},
	{
		Name:           "Tarball (bzip2) / .tbz2",
		Abbreviations:  []string{"TAR.BZ2", "TBZ2"},
		FileExtensions: []string{".tar.bz2", ".tbz2"},
		MimeTypes:      []string{"application/x-bzip2", "application/x-bzip"},
	},
	{
		Name:           "RAR Archive",
		Abbreviations:  []string{"RAR"},
		FileExtensions: []string{".rar"},
		MimeTypes:      []string{"application/x-rar-compressed", "application/vnd.rar"},
	},
	{
		Name:           "Unix compress (.Z)",
		Abbreviations:  []string{".Z"},
		FileExtensions: []string{".Z"},
		MimeTypes:      []string{"application/x-compress", "application/x-compressed"},
	},
	{
		Name:           "POSIX pax Archive",
		Abbreviations:  []string{"PAX"},
		FileExtensions: []string{".pax"},
		MimeTypes:      []string{"application/x-pax"},
	},
	{
		Name:           "AR Archive (Unix ar)",
		Abbreviations:  []string{"AR"},
		FileExtensions: []string{".ar"},
		MimeTypes:      []string{"application/x-archive"},
	},
	{
		Name:           "Windows Cabinet",
		Abbreviations:  []string{"CAB"},
		FileExtensions: []string{".cab"},
		MimeTypes:      []string{"application/vnd.ms-cab-compressed"},
	},
	{
		Name:           "ISO Disk Image",
		Abbreviations:  []string{"ISO"},
		FileExtensions: []string{".iso"},
		MimeTypes:      []string{"application/x-iso9660-image", "application/octet-stream"},
	},
	{
		Name:           "Apple Disk Image",
		Abbreviations:  []string{"DMG"},
		FileExtensions: []string{".dmg"},
		MimeTypes:      []string{"application/x-apple-diskimage", "application/octet-stream"},
	},
}

func init() {
	hrblock := MimeTypeModel{
		Name:      "H&R Block Data",
		MimeTypes: []string{"application/x-hrblock"},
	}

	turboTax := MimeTypeModel{
		Name:      "H&R Block Data",
		MimeTypes: []string{"application/x-hrblock"},
	}

	for i := 10; i <= 50; i++ {
		hrblock.FileExtensions = append(hrblock.FileExtensions, fmt.Sprintf(".T%d", i))
		turboTax.FileExtensions = append(turboTax.FileExtensions, fmt.Sprintf(".tax20%d", i))
	}

	all = append(all, hrblock)
	all = append(all, turboTax)
}
