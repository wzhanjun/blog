## linux 虚拟内存


问题： virtual memory exhausted: Cannot allocate memory

free -m  查看内存

mkdir /usr/local/temp

dd if=/dev/zero of=/usr/img/swap bs=1024 count=2048000     

mkswap /usr/local/temp/swap  

swapon /usr/local/temp/swap 

swapoff swap 

rm -f /usr/local/temp/swap 

参考： [http://www.cnblogs.com/chenpingzhao/p/4820814.html](http://www.cnblogs.com/chenpingzhao/p/4820814.html)
