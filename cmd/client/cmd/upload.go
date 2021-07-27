package cmd

import (
	"fmt"
	"path"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"

	"gitee.com/infraboard/go-course/day8/cloudstation/store"
	"gitee.com/infraboard/go-course/day8/cloudstation/store/provider/aliyun"
)

const (
	// BuckName todo
	defaultBuckName = "cloud-station"
	defaultEndpoint = "http://oss-cn-chengdu.aliyuncs.com"
	defaultALIAK    = "LTAI5tMvG5NA51eiH3ENZDaa"
	defaultALISK    = ""
)

var (
	buckName       string
	uploadFilePath string
	bucketEndpoint string
)

// uploadCmd represents the start command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "上传文件到中转站",
	Long:  `上传文件到中转站`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := getProvider()
		if err != nil {
			return err
		}

		if uploadFilePath == "" {
			return fmt.Errorf("upload file path is missing")
		}
		day := time.Now().Format("20060102")
		fn := path.Base(uploadFilePath)
		ok := fmt.Sprintf("%s/%s", day, fn)
		err = p.UploadFile(buckName, ok, uploadFilePath)
		if err != nil {
			return err
		}
		return nil
	},
}

func getAccessKeyFromInput() {
	fmt.Printf("请输入access key: ")
	fmt.Scanln(&aliAccessKey)
}

func getAccessKeyFromInputV2() {
	prompt := &survey.Password{
		Message: "请输入access key: ",
	}
	survey.AskOne(prompt, &aliAccessKey)
}

func getProvider() (p store.Uploader, err error) {
	switch ossProvider {
	case "aliyun":
		fmt.Printf("上传云商: 阿里云[%s]\n", defaultEndpoint)
		if aliAccessID == "" {
			aliAccessID = defaultALIAK
		}
		if aliAccessKey == "" {
			aliAccessKey = defaultALISK
		}
		fmt.Printf("上传用户: %s\n", aliAccessID)
		getAccessKeyFromInputV2()
		p, err = aliyun.NewUploader(bucketEndpoint, aliAccessID, aliAccessKey)
		return
	case "qcloud":
		return nil, fmt.Errorf("not impl")
	default:
		return nil, fmt.Errorf("unknown oss privier options [aliyun/qcloud]")
	}
}

func init() {
	uploadCmd.PersistentFlags().StringVarP(&uploadFilePath, "file_path", "f", "", "upload file path")
	uploadCmd.PersistentFlags().StringVarP(&buckName, "bucket_name", "b", defaultBuckName, "upload oss bucket name")
	uploadCmd.PersistentFlags().StringVarP(&bucketEndpoint, "bucket_endpoint", "e", defaultEndpoint, "upload oss endpoint")
	RootCmd.AddCommand(uploadCmd)
}
