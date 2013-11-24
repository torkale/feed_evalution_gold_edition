package main

import (
    "flag"
    "github.com/SlyMarbo/rss"
    "log"
    "time"
)

func main() {
    var url string
    flag.StringVar(&url, "url", "", "url for feed")
    flag.Parse()
    log.Printf("**%s**", url)
    feed, err := rss.Fetch(url)
    if err != nil {
        log.Fatalf("%s", err)
    }
    dates := make([]time.Time, len(feed.Items))
    intervals := make([]time.Duration, len(feed.Items)-1)
    for i, item := range feed.Items {
        d := item.Date
        if d.After(time.Now()) {
            log.Printf("Item %s has time in the future: %s", item.Title, d)
        }
        dates[i] = d
    }
    for i, date := range dates {
        if i > 0 {
            intervals[i-1] = dates[i-1].Sub(date)
        }
    }
    var total time.Duration
    for i, duration := range intervals {
        if i == 0 {
            total = duration
        } else {
            total += duration
        }
    }
    log.Printf("Average time between items is %d minutes", int(total/time.Minute)/len(intervals))
}