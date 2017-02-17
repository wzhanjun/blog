## 编译一个动态库

libtest.so

```

	so_test.h：
	
	#include <stdio.h>
	#include <stdlib.h>
	void test_a();
	void test_b();
	void test_c();
	
	test_a.c：
	
	#include "so_test.h"
	void test_a()
	{
	    printf("this is in test_a...\n");
	}
	
	test_b.c：
	
	#include "so_test.h"
	void test_b()
	{
	    printf("this is in test_b...\n");
	}
	
	test_c.c：
	
	#include "so_test.h"
	void test_c()
	{
	    printf("this is in test_c...\n");
	}
	

	将这几个文件编译成一个动态库：libtest.so
	$ gcc test_a.c test_b.c test_c.c -fPIC -shared -o libtest.so


```

	test.c：

	#include "so_test.h"
	int main()
	{
	    test_a();
	    test_b();
	    test_c();
	    return 0;
	
	}
	将test.c与动态库libtest.so链接生成执行文件test：
	$ gcc test.c -L. -ltest -o test

####备注

- LD_LIBRARY_PATH

	$ export LD_LIBRARY_PATH=/opt/gtk/lib:$LD_LIBRARY_PATH
	
	可以用下面的命令查看 LD_LIBRAY_PATH 的设置内容：
	
	$ echo $LD_LIBRARY_PATH
	
	至此，库的两种设置就完成了。


参考： [http://blog.sina.com.cn/s/blog_4b9b714a0100ieam.html](http://blog.sina.com.cn/s/blog_4b9b714a0100ieam.html)