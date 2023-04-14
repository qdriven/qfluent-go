# `/configs`

Configuration file templates or default configs.

Put your `confd` or `consul-template` template files here.

## Functions and Demos

```toml
free_form="test"

[ws]
sh_home = "/bin/sh"
working_dir = "."
```

1. how to get a key:

```shell
	config := NewDefaultConfig()
	fmt.Println(config)
	assert.Equal(t, config.Get("free_form"), "test")
	assert.Equal(t, config.Get("ws.sh_home"), "/bin/sh")
```
2. use named config for additional configuration

```shell
	config := &AppConfig{}
	NewNamedConfig("app.toml", "named", config)
	fmt.Println(config)
	assert.Equal(t, config.Ws.ShHome, "/bin/sh")
	v := GetViperByName("named")
	//ws := v.Get("ws").(map[string]interface{})
	ws := v.Get("ws").(map[string]interface{})
	assert.Equal(t, ws["sh_home"], "/bin/sh")
```