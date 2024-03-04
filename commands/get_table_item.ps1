# get table item
aws dynamodb get-item --table-name go-tutorial-1 --key '{\"pk\":{\"S\":\"USERNAME#eric\"},\"sk\":{\"S\":\"TYPE#loginDetails\"}}'