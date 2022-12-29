package main

import (
	"fmt"
	"os"
)

//var testDownloader = func() (dl Downloader) {
//	dl.OutputDir = "download_test"
//	dl.Debug = true
//	return
//}()

func main() {
	exitOnError(rootCmd.Execute())
	//ctx := context.Background()

	//video, err := testDownloader.Client.GetVideoContext(ctx, "BaW_jenozKc")
	//require.NoError(err)

	//require.NoError(testDownloader.DownloadComposite(ctx, "", video, "hd1080", "mp4"))
}

func exitOnError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
