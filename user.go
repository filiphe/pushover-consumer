package main

type User struct {
	QuietHours bool   `json:"quiet_hours"`
	IsAndLic   bool   `json:"is_android_licensed"`
	IsIOSLic   bool   `json:"is_ios_licensed"`
	IsDeskLic  bool   `json:"is_desktop_licensed"`
	Email      string `json:"email"`
}
