import camelcaseKeys from "camelcase-keys"
import { coolcar } from "./service/proto_gen/trip_pb"
// app.ts
App<IAppOption>({
	globalData: {},
	onLaunch() {
		wx.request({
			url: "http://localhost:8080/trip/trip123",
			success(res) {

				const getTripRes = coolcar.GetTripResponse.fromObject(camelcaseKeys(res.data as object, {
					deep: true
				}));
				console.log(getTripRes)
				console.log(coolcar.TripStatus[getTripRes?.trip?.status ?? 0])
			}
		})
		// 登录
		wx.login({
			success: res => {
				console.log(res.code)
				// 发送 res.code 到后台换取 openId, sessionKey, unionId
			},
		})
	},
})