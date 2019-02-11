package stoverflow_data

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"strings"
	"regexp"
)

type Job struct {
	CompanyName string
	PublishDate string
	Location string
	Tags []string
}

func GetData() []Job {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://stackoverflow.com/jobs/feed")
	// fmt.Printf("%+v", feed)
	re := regexp.MustCompile(`\(([^\)]+)\)$`)
	var ret []Job
	for _, item := range feed.Items {
		company_name := item.Extensions["a10"]["author"][0].Children["name"][0].Value
		publish_date := item.Published
		location := re.FindStringSubmatch(item.Title)
		tags := item.Categories
		job := Job{company_name, publish_date, location[1], tags}
		ret = append(ret, job)
	}
	return ret
}

