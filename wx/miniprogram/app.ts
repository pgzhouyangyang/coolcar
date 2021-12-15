import camelcaseKeys from "camelcase-keys"
import { auth } from "./service/proto_gen/auth/auth_pb"

// app.ts
App<IAppOption>({
	globalData: {},
	onLaunch() {
		// 登录
		wx.login({
			success: res => {
				console.log(res.code)
				// 发送 res.code 到后台换取 openId, sessionKey, unionId

				wx.request({
					url: "http://localhost:8080/v1/auth/login",
					method: "POST",
					data: {
						code: res.code
					} as auth.v1.LoginRequest,
					success: res=> {
						const data: auth.v1.ILoginResponse = auth.v1.LoginResponse.fromObject(camelcaseKeys(res.data as object))

						console.log(data)
					}
				})
			},
		})
	},
})