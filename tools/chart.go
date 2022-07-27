package tools

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Item struct {
	Time time.Time
	Data bool
}

// 生成折线图
func GenerateLineGraph(data []Item) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		line := charts.NewLine()

		line.SetGlobalOptions(
			charts.WithYAxisOpts(opts.YAxis{
				SplitNumber: 1,
			}),
			charts.WithXAxisOpts(opts.XAxis{
				SplitNumber: len(data),
			}),
			charts.WithXAxisOpts(opts.XAxis{
				AxisLabel: &opts.AxisLabel{
					Show:         true,
					Interval:     "0",
					Rotate:       20, // 减小旋转角度
					ShowMinLabel: true,
					ShowMaxLabel: true,
					FontSize:     "10", // 减小标签字体大小
				},
			}),
		)

		// 横坐标,为时间
		var times []string
		// 纵坐标为是否限流
		items := make([]opts.LineData, 0)

		for _, v := range data {
			times = append(times, v.Time.Format("1"))
			items = append(items, opts.LineData{Value: v.Data})
		}

		line.SetXAxis(times)
		line.AddSeries("type1", items)
		line.Render(w)
	}
}

func Listen(port int, data []Item) {
	http.HandleFunc("/", GenerateLineGraph(data))
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
