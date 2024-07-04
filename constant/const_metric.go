package constant

import "math"

//指标 常量和字典

var (
	// UnitConversion 维护的单位换算
	UnitConversion = map[string]float64{
		"Bytes":  1,
		"KBytes": 1e3,
		"MBytes": 1e6,
		"GBytes": 1e9,
		"TBytes": 1e12,

		"Bytes/s":  1,    //Bps 字节每秒
		"KBytes/s": 1e3,  //KBps 千字节每秒
		"MBytes/s": 1e6,  //MBps 兆字节每秒
		"GBytes/s": 1e9,  //GBps 吉字节每秒
		"TBytes/s": 1e12, //TBps 太字节每秒

		"KiBytes": math.Pow(2, 10),
		"MiBytes": math.Pow(2, 20),
		"GiBytes": math.Pow(2, 30),
		"TiBytes": math.Pow(2, 40),

		"KiBytes/s": math.Pow(2, 10), //KiBps 千字节每秒, 二进制
		"MiBytes/s": math.Pow(2, 20), //MiBps 兆字节每秒, 二进制
		"GiBytes/s": math.Pow(2, 30), //GiBps 吉字节每秒, 二进制
		"TiBytes/s": math.Pow(2, 40), //TiBps 太字节每秒, 二进制

		"Bit":  1, //1Byte=8bit
		"KBit": 1e3,
		"MBit": 1e6,
		"GBit": 1e9,
		"TBit": 1e12,

		"Bits/s":  1,    //bps  (比特每秒): 这是数据传输速率的基本单位，表示每秒传输的比特数。
		"KBits/s": 1e3,  //Kbps 千比特每秒
		"MBits/s": 1e6,  //Mbps 兆比特每秒
		"GBits/s": 1e9,  //Gbps 吉比特每秒
		"TBits/s": 1e12, //Tbps 太比特每秒

		"KiBits": math.Pow(2, 10), //kiBit 千比特, 二进制
		"MiBits": math.Pow(2, 20), //MiBit 兆比特, 二进制
		"GiBts":  math.Pow(2, 30), //GiBit 吉比特, 二进制
		"TiBits": math.Pow(2, 40), //TiBit 太比特, 二进制

		"KiBits/s": math.Pow(2, 10), //kibps 千比特每秒, 二进制
		"MiBits/s": math.Pow(2, 20), //Mibps 兆比特每秒, 二进制
		"GiBts/s":  math.Pow(2, 30), //Gibps 吉比特每秒, 二进制
		"TiBits/s": math.Pow(2, 40), //Tibps 太比特每秒, 二进制

		"s":   1, //基本单位
		"ms":  1e-3,
		"us":  1e-6,
		"ns":  1e-9,
		"min": 60,
		"h":   3600,
		// 添加其他单位...
	}
	// Categories 维护的字典
	Categories = map[string]string{
		// 定义更多的单位类别

		"Bytes":  "0",
		"KBytes": "0",
		"MBytes": "0",
		"GBytes": "0",
		"TBytes": "0",

		"Bytes/s":  "1",
		"KBytes/s": "1",
		"MBytes/s": "1",
		"GBytes/s": "1",
		"TBytes/s": "1",

		"KiBytes": "2",
		"MiBytes": "2",
		"GiBytes": "2",
		"TiBytes": "2",

		"KiBytes/s": "3",
		"MiBytes/s": "3",
		"GiBytes/s": "3",
		"TiBytes/s": "3",

		"Bit":  "4", //1Byte=8bit
		"KBit": "4",
		"MBit": "4",
		"GBit": "4",
		"TBit": "4",

		"Bits/s":  "5",
		"KBits/s": "5",
		"MBits/s": "5",
		"GBits/s": "5",
		"TBits/s": "5",

		"KiBits": "6",
		"MiBits": "6",
		"GiBts":  "6",
		"TiBits": "6",

		"KiBits/s": "7",
		"MiBits/s": "7",
		"GiBits/s": "7",
		"TiBits/s": "7",

		"s":   "8",
		"ms":  "8",
		"us":  "8",
		"ns":  "8",
		"min": "8",
		"h":   "8",

		// ...其他类别和单位
	}
)
