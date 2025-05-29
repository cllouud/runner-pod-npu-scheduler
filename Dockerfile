# 构建阶段
FROM --platform=linux/arm64 golang:1.21 AS builder

WORKDIR /app

# 复制源代码和 Makefile
COPY . .

# 安装依赖并构建
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o build/runner-pod-npu-scheduler ./cmd/main.go

# 最终阶段
FROM --platform=linux/arm64 arm64v8/alpine:3.18

# 创建必要的目录
RUN mkdir -p /etc/kubernetes

# 设置工作目录
WORKDIR /app

# 从构建阶段复制编译好的主程序
COPY --from=builder /app/build/runner-pod-npu-scheduler .

# 设置入口点
ENTRYPOINT ["./runner-pod-npu-scheduler"]
