# courses-go-goroutine
## overview
Introduction
Goroutines
Channels
Waitgroup
version > 1.19
## buoi 1
- waitgroup
- channel basic
### note
thread main
thread receiverOrder

A: 8 CPU- 16ram < 1s
B: 0.2 CPU- 0.5ram > 1s

- func chưa chạy xong, return.
- waitgroup ko done -> deadlook
- không được push data vào channel khi channel nil

- insert user:
  - check user tôn tại
  - check password, 8sô, có ký tự đăc biet

wait: insert data


## Đọc thêm
- tại sao add waitgroup nhiều hơn số lượng goroutine -> deadlock!
- GC
- channel buffer và unbuffer
- error group