
STEP.1 -->   go mod init NAME_YOU_WANT_GIVE     ex. go mod init api
STEP.2 -->   go get -u github.com/gin-gonic/gin
STEP.3 -->   Create main.go
STEP.4 -->   Create Models, Controllers, Routes    -----folder
STEP.5 -->   make model file .....ex UserModel.go
                --------UserModel 
                        1. create struct    

STEP.6 -->  make controller file.... ex UserController.go 
                --------UserController
                        1. Write Defination of function 

STEP.7 -->  make route file ......ex UserRoute.go
                --------UserRoute
                        1. Create SetRoute Function which return *gin.Engine 
                        2. set paths or routes GET POST PUT DELATE
                        


