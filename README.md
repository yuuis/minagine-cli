# minagine cli

[minagine-kun](https://github.com/yuuis/minagine-kun)のcli。  
先にminagine-kunをデプロイする必要がある。デプロイ方法は、[minagine-kun/README](https://github.com/yuuis/minagine-kun/blob/master/README.md)を参照。

## how to use

```shell script
go get github.com/yuuis/minagine-cli/cmd/minagine
```

```shell script
export MINAGINE_TOKEN=xxxxxxxxx
export MINAGINE_PROJECT=xxxxxxxxx

minagine start # start working
minagine end # end working
```

```shell script
minagine start --project 'xxxxxxxxx' --token 'xxxxxxxxx' # start working
minagine end --project 'xxxxxxxxx' --token 'xxxxxxxxx' # end working
```
