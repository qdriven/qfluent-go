package feishu

import (
	"context"
	"fmt"
	"github.com/88250/lute"
	"github.com/chyroc/lark"
	"github.com/pkg/errors"
	"github.com/qdriven/qfluent-go/pkg/feishu"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func handleDumpCommand(url string) error {
	configPath, err := feishu.GetConfigFilePath()
	feishu.CheckErr(err)
	config, err := feishu.ReadConfigFromFile(configPath)
	feishu.CheckErr(err)

	reg := regexp.MustCompile("^https://[a-zA-Z0-9-]+.(feishu.cn|larksuite.com|f.mioffice.cn)/(docx|wiki)/([a-zA-Z0-9]+)")
	matchResult := reg.FindStringSubmatch(url)
	if matchResult == nil || len(matchResult) != 4 {
		return errors.Errorf("Invalid feishu/larksuite URL format")
	}

	domain := matchResult[1]
	docType := matchResult[2]
	docToken := matchResult[3]

	ctx := context.Background()

	client := feishu.NewClient(
		config.Feishu.AppId, config.Feishu.AppSecret, domain,
	)

	// for a wiki page, we need to renew docType and docToken first
	if docType == "wiki" {
		node, err := client.GetWikiNodeInfo(ctx, docToken)
		feishu.CheckErr(err)
		docType = node.ObjType
		docToken = node.ObjToken
	}

	docx, blocks, err := client.GetDocxContent(ctx, docToken)
	feishu.CheckErr(err)

	data := struct {
		Document *lark.DocxDocument `json:"document"`
		Blocks   []*lark.DocxBlock  `json:"blocks"`
	}{
		Document: docx,
		Blocks:   blocks,
	}
	pdata := feishu.PrettyPrint(data)
	fmt.Println(pdata)

	return nil
}

type DownloadOpts struct {
	allowHost string
	outputDir string
	dump      bool
}

var downloadOpts = DownloadOpts{}

func handleDownloadCommand(url string, opts *DownloadOpts) error {
	// Validate the url to download
	domain, docType, docToken, err := feishu.ValidateDownloadURL(url, opts.allowHost)
	feishu.CheckErr(err)
	fmt.Println("Captured document token:", docToken)

	// Load config
	configPath, err := feishu.GetConfigFilePath()
	feishu.CheckErr(err)
	config, err := feishu.ReadConfigFromFile(configPath)
	feishu.CheckErr(err)

	// Create client with context
	ctx := context.WithValue(context.Background(), "output", config.Output)

	client := feishu.NewClient(
		config.Feishu.AppId, config.Feishu.AppSecret, domain,
	)

	// for a wiki page, we need to renew docType and docToken first
	if docType == "wiki" {
		node, err := client.GetWikiNodeInfo(ctx, docToken)
		feishu.CheckErr(err)
		docType = node.ObjType
		docToken = node.ObjToken
	}
	if docType == "docs" {
		return errors.Errorf("Feishu Docs is no longer supported. Please refer to the Readme/Release for v1_support.")
	}

	// Process the download
	docx, blocks, err := client.GetDocxContent(ctx, docToken)
	feishu.CheckErr(err)

	parser := feishu.NewParser(ctx)

	title := docx.Title
	markdown := parser.ParseDocxContent(docx, blocks)

	if !config.Output.SkipImgDownload {
		for _, imgToken := range parser.ImgTokens {
			localLink, err := client.DownloadImage(
				ctx, imgToken, filepath.Join(opts.outputDir, config.Output.ImageDir),
			)
			feishu.CheckErr(err)
			markdown = strings.Replace(markdown, imgToken, localLink, 1)
		}
	}

	// Format the markdown document
	engine := lute.New(func(l *lute.Lute) {
		l.RenderOptions.AutoSpace = true
	})
	result := engine.FormatStr("md", markdown)

	// Handle the output directory and name
	if _, err := os.Stat(opts.outputDir); os.IsNotExist(err) {
		if err := os.MkdirAll(opts.outputDir, 0o755); err != nil {
			return err
		}
	}

	if opts.dump {
		jsonName := fmt.Sprintf("%s.json", docToken)
		outputPath := filepath.Join(opts.outputDir, jsonName)
		data := struct {
			Document *lark.DocxDocument `json:"document"`
			Blocks   []*lark.DocxBlock  `json:"blocks"`
		}{
			Document: docx,
			Blocks:   blocks,
		}
		pdata := feishu.PrettyPrint(data)

		if err = os.WriteFile(outputPath, []byte(pdata), 0o644); err != nil {
			return err
		}
		fmt.Printf("Dumped json response to %s\n", outputPath)
	}

	// Write to markdown file
	mdName := fmt.Sprintf("%s.md", docToken)
	if config.Output.TitleAsFilename {
		mdName = fmt.Sprintf("%s.md", title)
	}
	outputPath := filepath.Join(opts.outputDir, mdName)
	if err = os.WriteFile(outputPath, []byte(result), 0o644); err != nil {
		return err
	}
	fmt.Printf("Downloaded markdown file to %s\n", outputPath)

	return nil
}
