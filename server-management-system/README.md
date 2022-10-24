# VCS_SMS
##1.Quản lý trạng thái On/Off của  server
##2.Installation
##3.Build
###1.Tạo file docker-compose.yml để 
###2.Tạo file app.env chứa thông tin đăng nhập vào postgres
###3.Khởi động Postgres Docker container với lệnh:      docker-compose up -d
###4.Tạo file connectDB.go để kết nối với Postgres server
###5. Xây dựng mô hình cơ sở dữ liệu, tạo User trong Postgres
###6. Tạo file migrate.go và chạy
###7. Tạo Golang server với gin gonic, package air để reload Gin server mỗi lần thay đổi
###8.Test Golang API server
- Register cho User
- Login User
- Lấy thông tin User đang Login
- Logout User
- Tạo Server
- View Server (kèm Sort và Filter theo trường nào đó)
- Update 1 sever theo id
- Xóa 1 Server theo id
###9.Export danh sách server từ db ra file excel, Import tạo danh sách server vào file excel 
##4.Technologies Used
- Viper package: github.com/spf13/viper
- Gorm package
- Gin framework

##5.Chưa hoàn thành
- Báo cáo định kỳ (gửi email)
- Unit Test
- Dùng Redis Cache để Optimize performance
- Dùng Elasticsearch để tính uptime của server