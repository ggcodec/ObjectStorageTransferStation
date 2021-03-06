package aliyun

import (
	"fmt"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/k0kubun/go-ansi"
	progressbar "github.com/schollz/progressbar/v3"
)

// NewOssProgressListener todo
func NewOssProgressListener() *OssProgressListener {
	return &OssProgressListener{}
}

// OssProgressListener is the progress listener
type OssProgressListener struct {
	bar     *progressbar.ProgressBar
	startAt time.Time
}

// ProgressChanged todo
func (p *OssProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		p.bar = progressbar.NewOptions64(event.TotalBytes,
			progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
			progressbar.OptionEnableColorCodes(true),
			progressbar.OptionShowBytes(true),
			progressbar.OptionSetWidth(30),
			progressbar.OptionSetDescription("开始上传:"),
			progressbar.OptionSetTheme(progressbar.Theme{
				Saucer:        "=",
				SaucerHead:    ">",
				SaucerPadding: " ",
				BarStart:      "[",
				BarEnd:        "]",
			}),
		)
		p.startAt = time.Now()
		fmt.Printf("文件大小: %s\n", HumanBytesLoaded(event.TotalBytes))
	case oss.TransferDataEvent:
		p.bar.Add64(event.RwBytes)
	case oss.TransferCompletedEvent:
		fmt.Printf("\n上传完成: 耗时%d秒\n", int(time.Since(p.startAt).Seconds()))
	case oss.TransferFailedEvent:
		fmt.Printf("\n上传失败: \n")
	default:
	}
}

const (
	bu = 1 << 10
	kb = 1 << 20
	mb = 1 << 30
	gb = 1 << 40
	tb = 1 << 50
	eb = 1 << 60
)

// HumanBytesLoaded 单位转换
func HumanBytesLoaded(bytesLength int64) string {
	if bytesLength < bu {
		return fmt.Sprintf("%dB", bytesLength)
	} else if bytesLength < kb {
		return fmt.Sprintf("%.2fKB", float64(bytesLength)/float64(bu))
	} else if bytesLength < mb {
		return fmt.Sprintf("%.2fMB", float64(bytesLength)/float64(kb))
	} else if bytesLength < gb {
		return fmt.Sprintf("%.2fGB", float64(bytesLength)/float64(mb))
	} else if bytesLength < tb {
		return fmt.Sprintf("%.2fTB", float64(bytesLength)/float64(gb))
	} else {
		return fmt.Sprintf("%.2fEB", float64(bytesLength)/float64(tb))
	}
}
