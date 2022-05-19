#! bin/bash

echo "====== Serial ======"

time go run serial.go

echo "===== Paralelo ====="

time go run parallel.go
