## laravel-redis-cluster 配置

最近在学习laravel(版本5.5)时发现laravel redis cluster配置有问题一直连接不上。laravel文档中给的redis集群配置说明[https://laravel.com/docs/5.6/redis#interacting-with-redis](https://laravel.com/docs/5.6/redis#interacting-with-redis) 按照以上配置，一直无法保存数据到集群，看看源码 Illuminate\Redis\RedisManager文件中找到 resolveCluster($name 默认是default)

    protected function resolveCluster($name)
    {
        $clusterOptions = $this->config['clusters']['options'] ?? [];

        return $this->connector()->connectToCluster(
            $this->config['clusters'][$name], $clusterOptions, $this->config['options'] ?? []
        );
    }

 再到 Illuminate\Redis\Connectors\PredisConnector 

	public function connectToCluster(array $config, array $clusterOptions, array $options)
    {
        $clusterSpecificOptions = Arr::pull($config, 'options', []);

        return new PredisClusterConnection(new Client(array_values($config), array_merge(
            $options, $clusterOptions, $clusterSpecificOptions
        )));
    }

在new Client()传入的是一个一维数组，而predis源码目录下的案例代码中custom_cluster_distributor.php 创建集群是用的是一个二维数组。
问题找到了，按照predis的配置要求将default 改为一个二维数组如下：


	'redis' => [
		'client' => 'predis',
		'clusters' => [
			'options' =>['cluster'=>'redis'],
			'default' => [
				'first' => [
					'host' => env('REDIS_HOST', '127.0.0.1'),
					'port' => env('REDIS_PORT', 6380),
				],
				'second' => [
					'host'=>env('REDIS_HOST','127.0.0.1'),
					'port'=>env('REDIS_EXTRA_PORT',6381)
				]
			],
		]
	],



修改配置文件后，再次测试问题解决。