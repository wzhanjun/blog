### 前端下载文件

##### 同步下载

- 通过<a>标签的download

- open 或者 location.href

- 优缺：简单方便直接

- 缺点：不能进行鉴权

##### 异步下载

- 通过xhr利用blob下载

- 优点： 可设置header，也就可添加鉴权信息

- 缺点：兼容性问题
  
- 获取文件名
  - header：Content-Disposition

    ```
    Content-Disposition: attachment; filename=CMCoWork__________20200323151823_190342.xlsx; filename*=UTF-8''CMCoWork_%E4

    // xhr是XMLHttpRequest对象
    const content = xhr.getResponseHeader('content-disposition'); // 注意是全小写，自定义的header也是全小写
    if (content) {
        let name1 = content.match(/filename=(.*);/)[1]; // 获取filename的值
        let name2 = content.match(/filename\*=(.*)/)[1]; // 获取filename*的值
        name1 = decodeURIComponent(name1);
        name2 = decodeURIComponent(name2.substring(6)); // 这个下标6就是UTF-8''
    }

    filename* 是个现代浏览器支持的，为了解决filename的不足，一般是UTF-8
    ```

   - 自定义header

    ```
        Access-Control-Expose-Headers: Content-Disposition, custom-header
    ```


##### 参考

[https://juejin.cn/post/6844904069958467592#heading-12](https://juejin.cn/post/6844904069958467592#heading-12)