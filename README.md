# go mod init github.com/DavidAndersonAR/gopportunities.git
go build -o executavel .
go run main.go 
go mod tidy


# Em Go se a primeira letra de uma função nao estiver Maiuscula essa func nao é exportada
# Em Go tudo que esta no package é acessivel pelo package