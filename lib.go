package stoverflow_data

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"strings"
	"regexp"
)

type Job struct {
	company_name string
	publish_date string
	location string
	tags []string
}

func GetData(tags []string) []Job {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(fmt.Sprintf("https://stackoverflow.com/jobs/feed?q=%v&l=Bridgewater%%2c+MA%%2c+USA&u=Miles&d=100", strings.Join(tags, "+")))
	// fmt.Printf("%+v", feed)
	re := regexp.MustCompile(`\(([^\)]+)\)$`)
	var ret []Job
	for _, item := range feed.Items {
		company_name := item.Extensions["a10"]["author"][0].Children["name"][0].Value
		publish_date := item.Published
		location := re.FindStringSubmatch(item.Title)
		job := Job{company_name, publish_date, location[1], tags}
		ret = append(ret, job)
	}
	return ret
}
