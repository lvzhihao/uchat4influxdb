// Copyright © 2018 edwin <edwin.lzh@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"encoding/json"

	rmqtool "github.com/lvzhihao/go-rmqtool"
	"github.com/lvzhihao/goutils"
	"github.com/lvzhihao/uchat4influxdb/stats"
	"github.com/lvzhihao/uchatlib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

// memberJoinCmd represents the memberJoin command
var memberJoinCmd = &cobra.Command{
	Use:   "member_join 4 influxdb",
	Short: "member_join",
	Long:  `member_join for influxdb`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := GetLogger()
		defer logger.Sync()
		// rmqtool config
		rmqtool.DefaultConsumerToolName = viper.GetString("global_consumer_flag")
		rmqtool.Log.Set(GetZapLoggerWrapperForRmqtool(logger))
		//rmqtool.Log.Debug("logger warpper demo", "no key param", zap.Any("ccc", time.Now()), zap.Any("dddd", []string{"xx"}), "no key param again", zap.Error(errors.New("xx")))
		// load config
		config, err := LoadConfig("member_join_config")
		if err != nil {
			logger.Fatal("load config error", zap.Error(err))
		}
		logger.Debug("load member_join config success", zap.Any("config", config))

		queue, err := config.ConsumerQueue()
		if err != nil {
			logger.Fatal("migrate member_join_queue error", zap.Error(err))
		}

		influx, err := config.InfluxdbClient()
		if err != nil {
			logger.Fatal("call member_join_target_influxdb error", zap.Error(err))
		}
		defer influx.Close()

		queue.Consume(1, func(msg amqp.Delivery) {
			var rst uchatlib.UchatMemberJoin
			err := json.Unmarshal(msg.Body, &rst)
			if err != nil {
				msg.Ack(false) //先消费掉，避免队列堵塞
				rmqtool.Log.Error("process error", zap.Error(err), zap.Any("msg", msg))
			} else {
				tags := make(map[string]string, 0)
				fields := make(map[string]interface{}, 0)
				tags["room"] = rst.ChatRoomSerialNo
				tags["user"] = rst.WxUserSerialNo
				tags["father"] = rst.FatherWxUserSerialNo
				tags["type"] = goutils.ToString(rst.JoinChatRoomType)
				fields["count"] = 1
				err := stats.WriteMemberJoin(influx, config.Influxdb.Db, tags, fields, rst.JoinDate)
				if err != nil {
					logger.Error("write influxdb error", zap.Error(err))
				}
				msg.Ack(false)
			}
		}) //尽量保证聊天记录的时序，以api回调接口收到消息进入receive队列为准
	},
}

func init() {
	rootCmd.AddCommand(memberJoinCmd)
}
