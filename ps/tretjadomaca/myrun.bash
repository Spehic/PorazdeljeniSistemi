#/bin/bash


# go run TretjaDomaca.go -id 1 &
# go run TretjaDomaca.go -id 2 &
# go run TretjaDomaca.go -id 3 &
# go run TretjaDomaca.go -id 4 &

go run TretjaDomaca.go -p 9300 -id 0 -n 3 -m 5 -k 2 &
go run TretjaDomaca.go -p 9300 -id 1 -n 3 -m 5 -k 2 &
go run TretjaDomaca.go -p 9300 -id 2 -n 3 -m 5 -k 2