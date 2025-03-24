package dify

type EnumResponseMode string

const (
	ResponseModeStreaming EnumResponseMode = "streaming" // 流式模式（推荐）。基于 SSE（Server-Sent Events）实现类似打字机输出方式的流式返回。
	ResponseModeBlocking  EnumResponseMode = "blocking"  // 阻塞模式，等待执行完毕后返回结果。（请求若流程较长可能会被中断）。
)

type EnumFileType string // 文件类型

const (
	FileTypeDocument EnumFileType = "document" // 具体类型包含：'TXT', 'MD', 'MARKDOWN', 'PDF', 'HTML', 'XLSX', 'XLS', 'DOCX', 'CSV', 'EML', 'MSG', 'PPTX', 'PPT', 'XML', 'EPUB'
	FileTypeImage    EnumFileType = "image"    // 具体类型包含：'JPG', 'JPEG', 'PNG', 'GIF', 'WEBP', 'SVG'
	FileTypeAudio    EnumFileType = "audio"    // 具体类型包含：'MP3', 'M4A', 'WAV', 'WEBM', 'AMR'
	FileTypeVideo    EnumFileType = "video"    // 具体类型包含：'MP4', 'MOV', 'MPEG', 'MPGA'
	FileTypeCustom   EnumFileType = "custom"   // 具体类型包含：其他文件类型
)

type FileTransferMethodEnum string // 传递方式

const (
	FileTransferMethodRemoteUrl FileTransferMethodEnum = "remote_url" // 图片地址
	FileTransferMethodLocalFile FileTransferMethodEnum = "local_file" // 上传文件
)
