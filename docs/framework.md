app目录下包含了各个应用，其中api是对外的getway，对外提供restful api。api目录下新建的infra/rpc目录，处理所有rpc客户端（注册，初始化）。

rpc_gen目录下是rpc代码生成工具，用于生成rpc客户端代码。而app目录下其他的应用则是rpc服务端，提供rpc服务。

## 服务发现
  每个rpc服务端都需要在consul中注册服务，以便客户端发现。在app目录下的main.go中，会初始化consul客户端，然后注册服务。

  ```go
 	// consul	register
	r, err := consul.NewConsulRegister(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		klog.Fatal(err)
	}
	opts = append(opts, server.WithRegistry(r))
``` 

## 接口测试
### login
  ```
  curl -X POST http://localhost:8080/user/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "wrong_password"
  }'
  ```
预期
  ```
  {
  "base": {
    "status_code": 401,
    "status_message": "用户名或密码错误"
    }
  }
  ```
