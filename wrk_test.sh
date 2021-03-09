#!/bin/zsh

: << EOF
API性能测试脚本，自动执行wrk命令，采集数据、分析数据并调用gnuplot画图

1. 启动tiny_http_server(8080端口)
2. 执行测试脚本: ./wrktest.sh

脚本会生成_wrk.dat数据文件，每列含义为：

并发数 QPS 平均响应时间 成功率

对比2次测试结果

1. 执行命令：./wrktest.sh diff tiny_http_server_wrk.dat http_wrk.dat

> Note: 确保安装了wrk和gnuplot
EOF

t1="tiny_http_server"
t2="http"
job_name="tiny_http_server"

## wrk参数配置
d="300s"
concurrent="200 500 1000 3000 5000 10000 15000 20000 25000 50000 100000 200000 500000 1000000"
threads=144

if [ "$1" != "" ];then
	url="$1"
else
	url="http://127.0.0.1:8080/health/check"
fi

cmd="wrk --latency -t$threads -d$d -T30s $url"
api_perf="${job_name}_perf.png"
api_success_rate="${job_name}_success_rate.png"
data_file="${job_name}_wrk.dat"

# functions
function convertPlotData() {
	echo "$1" | awk -v data_file="$data_file" ' {
		if ($0 ~ "Running") {
			common_time=$2
		}
		if ($0 ~ "connections") {
			connections=$4
			common_threads=$1
		}
		if ($0 ~ "Latency   ") {
			avg_latency=convertLatency($2)
		}
		if ($0 ~ "50%") {
			p50=convertLatency($2)
		}
		if ($0 ~ "75%") {
			p75=convertLatency($2)
		}
		if ($0 ~ "90%") {
			p90=convertLatency($2)
		}
		if ($0 ~ "99%") {
			p99=convertLatency($2)
		}
		if ($0 ~ "Requests/sec") {
			qps=$2
		}
		if ($0 ~ "requests in") {
			allrequest=$1
		}
		if ($0 ~ "Socket errors") {
			err=$4+$6+$8+$10
		}

	}
	END {
		rate=sprintf("%.2f", (allrequest-err)*100/allrequest)
		print connections,qps,avg_latency,rate >> data_file
	}

	function convertLatency(s) {
		if (s ~ "us") {
			sub("us", "", s)
			return s/1000
		}
		if (s ~ "ms") {
			sub("ms", "", s)
			return s
		}
		if (s ~ "s") {
			sub("s", "", s)
			return s * 1000
		}
	}
	
	'
}

function prepare() {
	rm -f $data_file
}

function plot() {
	gnuplot <<  EOF
#输出格式为png文件
set terminal png enhanced
#指定数据文件名称
set output "$api_perf"
set title "QPS & TTLB\nRunning: 300s\nThreads: $threads"
set ylabel 'QPS'
set xlabel 'Concurrent'
set y2label 'Average Latency (ms)'
set key top left vertical noreverse spacing 1.2 box
set tics out nomirror
set border 3 front
set style line 1 linecolor rgb '#00ff00' linewidth 2 linetype 3 pointtype 2
set style line 2 linecolor rgb '#ff0000' linewidth 1 linetype 3 pointtype 2
set style data linespoints

#显示网格
set grid
#只需要一个x轴
set xtics nomirror rotate #by 90
set mxtics 5
#可以增加分刻度
set mytics 5
set ytics nomirror
set y2tics

set autoscale y
set autoscale y2

plot "$data_file" using 2:xticlabels(1) w lp pt 7 ps 1 lc rgbcolor "#EE0000" axis x1y1 t "QPS","$datfile" using 3:xticlabels(1) w lp pt 5 ps 1 lc rgbcolor "#0000CD" axis x2y2 t "Avg Latency (ms)"

unset y2tics
unset y2label
set ytics nomirror
set yrange[0:100]
#指定数据文件名称
set output "$api_success_rate"
set title "Success Rate\nRunning: 300s\nThreads: $threads"
plot "$data_file" using 4:xticlabels(1) w lp pt 7 ps 1 lc rgbcolor "#F62817" t "Success Rate"
EOF
}

function plotDiff() {
	gnuplot <<  EOF
#输出格式为png文件
set terminal png enhanced
#指定数据文件名称
set output "${t1}_$t2.qps.diff.png"
set title "QPS & TTLB\nRunning: 300s\nThreads: $threads"
set xlabel 'Concurrent'
set ylabel 'QPS'
set y2label 'Average Latency (ms)'
set key below left vertical noreverse spacing 1.2 box autotitle columnheader
set tics out nomirror
set border 3 front
set style line 1 linecolor rgb '#00ff00' linewidth 2 linetype 3 pointtype 2
set style line 2 linecolor rgb '#ff0000' linewidth 1 linetype 3 pointtype 2
set style data linespoints

#这会让坐标图的border更好看
#set border 3 lt 3 lw 2
#显示网格
set grid
#只需要一个x轴
set xtics nomirror rotate #by 90
set mxtics 5
#可以增加分刻度
set mytics 5
set ytics nomirror
set y2tics

#点的像素大小
#set pointsize 0.4
#数据文件的字段用\t分开
#set datafile separator '\t'

set autoscale y
set autoscale y2

#设置图像的大小 为标准大小的2倍
#set size 2.3,2

plot "/tmp/plot_diff.dat" using 2:xticlabels(1) w lp pt 7 ps 1 lc rgbcolor "#EE0000" axis x1y1 t "$t1 QPS","/tmp/plot_diff.dat" using 5:xticlabels(1) w lp pt 7 ps 1 lc rgbcolor "#EE82EE" axis x1y1 t "$t2 QPS","/tmp/plot_diff.dat" using 3:xticlabels(1) w lp pt 5 ps 1 lc rgbcolor "#0000CD" axis x2y2 t "$t1 Avg Latency (ms)", "/tmp/plot_diff.dat" using 6:xticlabels(1) w lp pt 5 ps 1 lc rgbcolor "#6495ED" axis x2y2 t "$t2 Avg Latency (ms)"

unset y2tics
unset y2label
set ytics nomirror
set yrange[0:100]
set title "Success Rate\nRunning: 300s\nThreads: $threads"
set output "${t1}_$t2.success_rate.diff.png"  #指定数据文件名称
plot "/tmp/plot_diff.dat" using 4:xticlabels(1) w lp pt 7 ps 1 lc rgbcolor "#EE0000" t "$t1 Success Rate","/tmp/plot_diff.dat" using 7:xticlabels(1) w lp pt 7 ps 1 lc rgbcolor "#EE82EE" t "$t2 Success Rate"
EOF
}

if [ "$1" == "diff" ];then
	join $2 $3 > /tmp/plot_diff.dat
	plotDiff `basename $2` `basename $3`
	exit 0
fi

prepare

for c in $concurrent
do
	wrkcmd="$cmd -c $c"
	echo -e "\nRunning wrk command: $wrkcmd"
	result=`eval $wrkcmd`
	convertPlotData "$result"
done

echo -e "\nNow plot according to $data_file"
plot &> /dev/null
echo -e "QPS graphic file is: $api_perf\nSuccess rate graphic file is: $api_success_rate"
