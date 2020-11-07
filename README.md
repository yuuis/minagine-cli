# minagine cli

[minagine-kun](https://github.com/yuuis/minagine-kun)のcli。  
先にこちらをデプロイする必要がある。デプロイ方法は、[minagine-kunのREADME](https://github.com/yuuis/minagine-kun)を参照。

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
