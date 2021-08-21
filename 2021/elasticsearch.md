## Elasticsearch 基础



##### 查询集群健康状况
```
GET /_cat/health?v
```



##### 查询集群中有哪些索引
```
GET /_cat/indices?v
```



##### 新增索引
```
PUT /demo_08
```



##### 删除指定索引
```
DELETE /demo_08
```

##### 删除所有索引
```
DELETE /_all
```

##### 查询索引配置信息
```
GET /demo/_settings
```

##### 查询所有索引配置信息
```
GET /_all/_settings
```



##### 新增索引，并指定 primary shards 和 replica shards 数量。
```
PUT /demo_order
{
  "settings": {
    "index": {
      "number_of_shards": 3,
      "number_of_replicas": 1
    }
  }
}
```



##### 新增完索引后，更改 replica shards 数量
```
PUT /demo_order/_settings
{
  "number_of_replicas": 2
}
```



##### 新增索引并设置 mapping
```
PUT /demo_product
{
  "settings": {
    "refresh_interval": "5s",
    "number_of_shards": 5,
    "number_of_replicas": 1
  },
  "mappings": {
    "properties": {
      "id": {
        "type": "text",
        "index": false
      },
      "product_name": {
        "type": "text",
        "boost": 5,
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "product_desc": {
        "type": "text"
      },
      "price": {
        "type": "float"
      },
      "created_date_time": {
        "type": "date"
      },
      "last_modified_date_time": {
        "type": "date",
        "format": "basic_date_time"
      },
      "version": {
        "type": "long",
        "index": false
      }
    }
  }
}
```




##### 新增文档
```
PUT /demo_product/_doc/1
{
    "product_name" : "PHILIPS toothbrush HX6730/02",
    "product_desc" :  "【3?9 元，前 1000 名赠刷头，6 月 1 日 0 点火爆开抢，618 开门红巅峰 48 小时，抢先加入购物车】飞利浦畅销款，万千好评！深入净齿，智能美白！",
    "price" :  399.00
}
```




##### 通过 ID 查询
```
GET /demo_product/_doc/1
```



##### 通过 ID 查询 （返回指定元数据）
```
GET /demo_product/_doc/1?_source=product_name,price
```



##### 查询指定索引的所有数据
```
GET /demo_product/_search
```



##### 设置查询超时
```
GET /demo_product/_search?timeout=5s
```



##### 查询多个索引
```
GET /demo_product,demo_order/_search
```

##### 匹配符模糊查询多个索引

```
GET /demo_*/_search
```



##### 通过商品名搜索，并价格倒序
```
GET /demo_product/_search?q=product_name:toothbrush&sort=price:desc
```



##### 通过商品名搜索（商品名称不等于搜索词，这里用了减号），并价格倒序
```
GET /demo_product/_search?q=-product_name:toothbrush&sort=price:desc
```



##### 普通分页 (from 不是页数，是第几条数据开始)
```
GET /demo_product/_search?from=0&size=2
```



##### 深度分页（Deep Paging）
```
GET /demo_product/_search?scroll=1m
{
  "query": {
    "match_all": {}
  },
  "sort": [
    "_doc"
  ],
  "size": 2
}
```



##### 更新整个 Document（需要带上所有属性，注意细节，这里改了 product_name）

```
PUT /demo_product/_doc/3
{
    "product_name" : "星空太空 iphone7 plus 蓝紫色 6s 繁星 7plus 宇宙 se 原创保护苹果 5 包手机壳",
    "product_desc" :  "一说到星空，就有太多美好的记忆，美丽的浩瀚宇宙，有太多说不清的神秘之处，星空太美丽，太绚烂！",
    "price" :  36.00
}
```

  这种方式的本质是：软删除。把旧版本标记为 deleted，实际还没物理删除，该条数据的 _version 元数据其实会再 +1 的。如果你再 PUT 下还是这个 ID 数据进去，_version 还是会继续 +1。当 Elasticsearch 数据越来越多，会物理删除这些标记的数据。

##### 更新部门 Document
```
POST /demo_product/_doc/3/_update
{
  "doc": {
    "product_name": "星空太空 iphone7 蓝紫色 6s 繁星 iphone7 plus 宇宙 se 原创保护苹果 5 包手机壳"
  }
}
```



##### 删除 Document
```
DELETE /demo_product/_doc/3
```




### DSL语句

##### 校验查询语句
```
GET /demo_product/_doc/_validate/query?explain
{
  "query": {
    "match": {
      "product_name": "toothbrush"
    }
  }
}
```



##### 查询所有的商品
```
GET /demo_product/_search
{
  "query": {
    "match_all": {}
  }
}
```



##### 查询商品名称包含 toothbrush 的商品，同时按照价格降序排序
```
GET /demo_product/_doc/_search
{
  "query": {
    "match": {
      "product_name": "toothbrush"
    }
  },
  "sort": [
    {
      "price": "desc"
    }
  ]
}
```



##### 分页查询商品
```
GET /demo_product/_doc/_search
{
  "query": {
    "match_all": {}
  },
  "from": 0, 
  "size": 1 
}
```




##### 指定查询结果字段（field）
```
GET /demo_product/_doc/_search
{
  "query": {
    "match_all": {}
  },
  "_source": [
    "product_name",
    "price"
  ]
}
```



##### 搜索商品名称包含 toothbrush，而且售价大于 400 元，小于 700 的商品
```
GET /demo_product/_doc/_search
{
  "query": {
    "bool": {
      "must": {
        "match": {
          "product_name": "toothbrush"
        }
      },
      "filter": {
        "range": {
          "price": {
            "gt": 400,
            "lt": 700
          }
        }
      }
    }
  }
}
```



##### 全文检索 (索引中只要有任意一个匹配拆分后词就可以出现在结果中，只是匹配度越高的排越前面)
```
GET /demo_product/_doc/_search
{
  "query": {
    "match": {
      "product_name": "PHILIPS toothbrush"
    }
  }
}
```



##### phrase search 短语搜索(索引中必须有同时有这两个单词的才会在结果中)
```
GET /demo_product/_doc/_search
{
  "query": {
    "match_phrase": {
      "product_name": "PHILIPS toothbrush"
    }
  }
}
```




##### Highlight Search 高亮搜索
```
GET /demo_product/_doc/_search
{
  "query": {
    "match": {
      "product_name": "PHILIPS toothbrush"
    }
  },
  "highlight": {
    "fields": {
      "product_name": {}
    }
  }
}
```



##### range 用法，查询数值、时间区间
```
GET /demo_product/_doc/_search
{
  "query": {
    "range": {
      "price": {
        "gte": 30.00
      }
    }
  }
}
```




##### match operator
```
GET /demo_product/_doc/_search
{
  "query": {
    "match": {
      "product_name": "PHILIPS toothbrush"
    }
  }
}

GET /demo_product/_doc/_search
{
  "query": {
    "match": {
      "product_name": {
        "query": "PHILIPS toothbrush",
        "operator": "and"
      }
     }
   }
}
```



##### 满足分词结果中百分比的词
```
GET /demo_product/_doc/_search
{
  "query": {
    "match": {
      "product_name": {
        "query": "java 程序员 书 推荐",
        "minimum_should_match": "50%"
      }
    }
  }
}
```



##### multi_match 用法： 查询 product_name 和 product_desc 字段中，只要有：toothbrush 关键字的就查询出来。
```
GET /demo_product/_doc/_search
{
  "query": {
    "multi_match": {
      "query": "toothbrush",
      "fields": [
        "product_name",
        "product_desc"
      ]
    }
  }
}
```



##### multi_match 跨多个 field 查询，表示查询分词必须出现在相同字段中。
```
GET /demo_product/_doc/_search
{
  "query": {
    "multi_match": {
      "query": "PHILIPS toothbrush",
      "type": "cross_fields",
      "operator": "and",
      "fields": [
        "product_name",
        "product_desc"
      ]
    }
  }
}
```



##### match_phrase 用法（短语搜索）（与 match 进行对比）：对这个查询词不进行分词，必须完全匹配查询词才可以作为结果显示。
```
GET /demo_product/_doc/_search
{
  "query": {
    "match_phrase": {
      "product_name": "PHILIPS toothbrush"
    }
  }
}
```



##### match_phrase + slop （查询的关键字分词后移动多少位可以跟 doc 内容匹配，移动的次数就是 slop）
```
GET /demo_product/_doc/_search
{
  "query": {
    "match_phrase": {
      "product_name" : {
          "query" : "PHILIPS HX6730",
          "slop" : 1
      }
    }
  }
}
```



##### 多搜索条件组合查询（最常用）

##### bool 下包括：must（必须匹配，类似于数据库的 =），must_not（必须不匹配，类似于数据库的 !=），should（没有强制匹配，类似于数据库的 or），filter（过滤）

```
GET /demo_product/_doc/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "product_name": "PHILIPS toothbrush"
          }
        }
      ],
      "should": [
        {
          "match": {
            "product_desc": "刷头"
          }
        }
      ],
      "must_not": [
        {
          "match": {
            "product_name": "HX6730"
          }
        }
      ],
      "filter": {
        "range": {
          "price": {
            "gte": 33.00
          }
        }
      }
    }
  }
}

GET /demo_product/_doc/_search
{
  "query": {
    "bool": {
      "should": [
        {
          "term": {
            "product_name": "飞利浦"
          }
        },
        {
          "bool": {
            "must": [
              {
                "term": {
                  "product_desc": "保护牙"
                }
              }
            ]
          }
        }
      ]
    }
  }
}
```



##### should 有一个特殊性，如果组合查询中没有 must 条件，那么 should 中必须至少匹配一个。我们也还可以通过 minimum_should_match 来限制它匹配更多个

```
GET /demo_product/_doc/_search
{
  "query": {
    "bool": {
      "should": [
        {
          "match": {
            "product_name": "java"
          }
        },
        {
          "match": {
            "product_name": "程序员"
          }
        },
        {
          "match": {
            "product_name": "书"
          }
        },
        {
          "match": {
            "product_name": "推荐"
          }
        }
      ],
      "minimum_should_match": 3
    }
  }
}


GET /demo_product/_doc/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "product_name": "PHILIPS toothbrush"
          }
        }
      ],
      "should": [
        {
          "match": {
            "product_desc": "刷头"
          }
        }
      ],
      "filter": {
        "bool": {
          "must": [
            {
              "range": {
                "price": {
                  "gte": 33.00
                }
              }
            },
            {
              "range": {
                "price": {
                  "lte": 555.55
                }
              }
            }
          ],
          "must_not": [
            {
              "term": {
                "product_name": "HX6730"
              }
            }
          ]
        }
      }
    }
  },
  "sort": [
    {
      "price": {
        "order": "desc"
      }
    }
  ]
}
```





#####  boost 用法（默认是 1）。在搜索精准度的控制上，还有一种需求，比如搜索：PHILIPS toothbrush，要比：Braun toothbrush 更加优先，我们可以这样：

```
GET /demo_product/_doc/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "product_name": "toothbrush"
          }
        }
      ],
      "should": [
        {
          "match": {
            "product_name": {
              "query": "PHILIPS",
              "boost": 4
            }
          }
        },
        {
          "match": {
            "product_name": {
              "query": "Braun",
              "boost": 5
            }
          }
        }
      ]
    }
  }
}
```



##### dis_max 用法，也称作：best fields 策略。
```
GET /demo_product/_doc/_search
{
  "query": {
    "dis_max": {
      "queries": [
        {
          "match": {
            "product_name": "PHILIPS toothbrush"
          }
        },
        {
          "match": {
            "product_desc": "iphone shell"
          }
        }
      ],
      "tie_breaker": 0.2
    }
  }
}


GET /demo_product/_doc/_search
{
  "query": {
    "dis_max": {
      "queries": [
        {
          "match": {
            "product_name": {
              "query": "PHILIPS toothbrush",
              "minimum_should_match": "50%",
              "boost": 3
            }
          }
        },
        {
          "match": {
            "product_desc": {
              "query": "iphone shell",
              "minimum_should_match": "50%",
              "boost": 2
            }
          }
        }
      ],
      "tie_breaker": 0.3
    }
  }
}
```





##### prefix 前缀搜索（性能较差，扫描所有倒排索引）
```
GET /demo_product/_doc/_search
{
  "query": {
    "prefix": {
      "product_name": {
        "value": "iphone"
      }
    }
  }
}
```



##### 通配符搜索（性能较差，扫描所有倒排索引）
```
GET /demo_product/_doc/_search
{
  "query": {
    "wildcard": {
      "product_name": {
        "value": "ipho*"
      }
    }
  }
}
```




##### 正则搜索（性能较差，扫描所有倒排索引）
```
GET /demo_product/_doc/_search
{
  "query": {
    "regexp": {
      "product_name": "iphone[0-9]."
    }
  }
}
```



##### fuzzy 纠错查询(参数 fuzziness 默认是 2，表示最多可以纠错两次，但是这个值不能很大，不然没效果。一般 AUTO 是自动纠错。)
```
GET /demo_product/_doc/_search
{
  "query": {
    "match": {
      "product_name": {
        "query": "PHILIPS tothbrush",
        "fuzziness": "AUTO",
        "operator": "and"
      }
    }
  }
}
```



##### 自定义一个针对路径分词分词器（eg：/usr/local/nginx，会分词为：/usr/local/nginx，/usr/local，/usr）

```
POST /demo_order/_close

PUT /demo_order/_settings
{
  "settings": {
    "analysis": {
      "analyzer": {
        "custom_path_analyzer": {
          "tokenizer": "path_hierarchy"
        }
      }
    }
  }
}

POST /demo_order/_open
```



##### 校验下这个自定义分词器
```
GET /demo_order/_analyze
{
  "text": "/usr/local/nginx",
  "analyzer": "custom_path_analyzer"
}
```



##### 通过路径地址查询
```
GET /demo_order/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "order_num": "123456"
          }
        },
        {
          "constant_score": {
            "filter": {
              "term": {
                "order_image_path": "/usr/local/nginx"
              }
            }
          }
        }
      ]
    }
  }
}
```



##### 创建带有 Tags 的索引数据

```
PUT /demo_product_0815/_doc/3
{
  "product_name": "iphone7 shell",
  "product_desc": "一说到星空，就有太多美好的记忆，美丽的浩瀚宇宙，有太多说不清的神秘之处，星空太美丽，太绚烂！",
  "price": 36.00,
  "tags": [
    "iphone7",
    "phone",
    "shell"
  ]
}
```



##### 计算每个 tag 下的商品数量
```
GET /demo_product_0815/_search
{
  "aggs": {
    "product_group_by_tags": {
      "terms": {
        "field": "tags"
      }
    }
  }
}

POST /demo_product_0815/_mapping
{
  "properties": {
    "tags": { 
      "type":     "text",
      "fielddata": true
    }
  }
}
```



##### 不显示 hits 原数据，只显示聚合统计结果
```
GET /demo_product_0815/_search
{
  "size": 0, 
  "aggs": {
    "product_group_by_tags": {
      "terms": {
        "field": "tags"
      }
    }
  }
}
```



##### 搜索商品名称中包含 toothbrush 的商品结果中，计算每个 tag 下的商品数量
```
GET /demo_product_0815/_search
{
  "size": 0,
  "query": {
    "match": {
      "product_name": "toothbrush"
    }
  },
  "aggs": {
    "query_product_group_by_tags": {
      "terms": {
        "field": "tags"
      }
    }
  }
}
```




##### 搜索商品名称中包含 toothbrush 的商品结果中，先用 tags 字段进行分组，然后再计算每组中商品价格的平均值

```
GET /demo_product_0815/_search
{
  "size": 0,
  "query": {
    "match": {
      "product_name": "toothbrush"
    }
  },
  "aggs": {
    "query_product_group_by_tags_and_avg": {
      "terms": {
        "field": "tags"
      },
      "aggs": {
        "product_price_avg_price": {
          "avg": {
            "field": "price"
          }
        }
      }
    }
  }
}
```



##### 搜索商品名称中包含 toothbrush 的商品结果中，先用 tags 字段进行分组，然后再计算每组中商品价格的平均值，并按平均价格进行倒序

```
GET /demo_product_0815/_search
{
  "size": 0,
  "query": {
    "match": {
      "product_name": "toothbrush"
    }
  },
  "aggs": {
    "query_product_group_by_tags_and_avg": {
      "terms": {
        "field": "tags",
        "order": {
          "product_price_avg_price": "asc"
        }
      },
      "aggs": {
        "product_price_avg_price": {
          "avg": {
            "field": "price"
          }
        }
      }
    }
  }
}
```



##### 搜索商品名称中包含 toothbrush 的商品结果中，按照指定的价格范围区间进行分组聚合，然后再按 tag 进行分组，最后再计算每组的平均价格

```
GET /demo_product_0815/_search
{
  "size": 0,
  "query": {
    "match": {
      "product_name": "toothbrush"
    }
  },
  "aggs": {
    "proudct_group_by_price": {
      "range": {
        "field": "price",
        "ranges": [
          {
            "from": 0,
            "to": 300
          },
          {
            "from": 300,
            "to": 400
          },
          {
            "from": 400,
            "to": 1000
          }
        ]
      },
      "aggs": {
        "product_group_by_tags": {
          "terms": {
            "field": "tags"
          },
          "aggs": {
            "product_average_price": {
              "avg": {
                "field": "price"
              }
            }
          }
        }
      }
    }
  }
}
```




##### mget 批量查询

##### 根据 index 名称，ID 进行查询（可以是不同 index、id）：
```
GET /_mget
{
  "docs": [
    {
      "_index": "demo_product_0815",
      "_id": 1
    },
    {
      "_index": "demo_product",
      "_id": 1
    }
  ]
}
```




##### 根据 index 名称，ID 进行查询（可以是不同 index、id）：
```
GET /demo_product_0815/_mget
{
  "docs": [
    {
      "_id": 1
    },
    {
      "_id": 1
    }
  ]
}

GET /demo_product_0815/_mget
{
  "ids": [1, 2]
}
```




##### bulk 批量增删改
```
POST /_bulk
{"delete": {"_index": "demo_product_0815","_id": "1"}}

POST /_bulk
{ "create": { "_index": "demo_product_0815", "_id": "333" } }
{ "product_name": "iphone7 shell2", "product_desc": "一说到星空，就有太多美好的记忆，美丽的浩瀚宇宙，有太多说不清的神秘之处，星空太美丽，太绚烂！", "price": 36.00, "tags": [ "iphone7", "phone", "shell" ] }


POST /_bulk
{"update":{"_index": "demo_product_0815","_id": "333"}}
{"doc":{"product_name": "iphone7 shell2222"}}
```

 在调用 INDEX、UPDATE、DELETE api 时，通过在 url 传递参数 if_seq_no 和 if_primary_term 来分别指定 序列号 和 主项




##### 参考
- https://github.com/judasn/Elasticsearch-Tutorial-zh-CN/blob/master/elasticsearch-base
- https://stackoverflow.com/questions/19758335/error-when-trying-to-update-the-settings
- https://www.cnblogs.com/yjf512/p/4897294.html
- https://www.cyhone.com/articles/introduction-of-elasticsearch/

