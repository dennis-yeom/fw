# fw
fw able to read linode buckets

## creating a module...
create a module with path of git hub repo

```
go mod init github.com/dennis-yeom/fw
```

## installing cobra...
make sure you have downloaded necessary packages:

```
go get -u github.com/spf13/cobra
```

## testing cobra command lines...
to see possible entries:
```
go run main.go -h
```

## enabling versioning in linode bucket
dont forget to export your linode access, secret keys


enable versioning:
```
aws s3api put-bucket-versioning --bucket pcw-test --versioning-configuration Status=Enabled --endpoint-url https://us-east-1.linodeobjects.com
```

check:
```
aws s3api get-bucket-versioning --bucket pcw-test --endpoint-url https://us-east-1.linodeobjects.com
```