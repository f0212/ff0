# ff0

通过网络资产线索（如：域名，IP地址，资产名称等），利用FOFA访问网络空间测绘数据，对资产进行拓展查询

## 0x00 简介

ff0 是一款使用 Go 编写的命令行 FoFa 查询工具，可以对特定资产进行拓展搜索，主要的功能如下：
基本 FoFa 语法查询     Icon Hash 在线查询     URL 证书计算查询    排除蜜罐资产 	更多......

## 0x01 下载

点击 [releases](https://github.com/zhangyang9123/ff0/releases/tag/ff0)下载链接 ，按照自己的系统架构选择相应的发行版本下载。

## 0x02 配置

MacOS/Linux/Win

第一次运行程序会在同级目录	生成config.yml配置文件，只需要配置完 `email` 和 `key` 就可以使用

```
DefaultEmail  : 
DefaultAPIKey : 
DefaultSize   : 10000
DefaultOutput : data.xlsx
```

## 0x03 使用方法

直接运行程序会打印出提示信息

```
/ "... an experienced, industrious, ambitious, and often quite often \
| picturesque liar."                                                 |
\                 -- Mark Twain                                      /
 --------------------------------------------------------------------
  \                                  ,+*^^*+___+++_
   \                           ,*^^^^              )
    \                       _+*                     ^**+_
     \                    +^       _ _++*+_+++_,         )
              _+^^*+_    (     ,+*^ ^          \+_        )
             {       )  (    ,(    ,_+--+--,      ^)      ^\
            { (@)    } f   ,(  ,+-^ __*_*_  ^^\_   ^\       )
           {:;-/    (_+*-+^^^^^+*+*<_ _++_)_    )    )      /
          ( /  (    (        ,___    ^*+_+* )   <    <      \
           U _/     )    *--<  ) ^\-----++__)   )    )       )
            (      )  _(^)^^))  )  )\^^^^^))^*+/    /       /
          (      /  (_))_^)) )  )  ))^^^^^))^^^)__/     +^^
         (     ,/    (^))^))  )  ) ))^^^^^^^))^^)       _)
          *+__+*       (_))^)  ) ) ))^^^^^^))^^^^^)____*^
          \             \_)^)_)) ))^^^^^^^^^^))^^^^)
           (_             ^\__^^^^^^^^^^^^))^^^^^^^)
             ^\___            ^\__^^^^^^))^^^^^^^^)\\
                  ^^^^^\uuu/^^\uuu/^^^^\^\^\^\^\^\^\^\
                     ___) >____) >___   ^\_\_\_\_\_\_\)
                    ^^^//\\_^^//\\_^       ^(\_\_\_\)
                      ^^^ ^^ ^^^ ^

Usage:
  [!]./linux_ff0  -q 'host="sdxiehe.edu.cn"' -s 10000 -o data.xlsx
  [!]./linux_ff0  -f query_rules_file.txt -s 10000 -o data.xlsx

Options:
  -h, --help
  -q, --query       [string]          参数字符串 (默认: '')
  -f, --file        [filePath]        批量查询规则文件 (默认: '')
  -s, --size        [int]             到处数据量 (默认: 10000)
  -e, --is_honeypot                   排除蜜罐数据  (仅限FOFA高级会员使用 )
  -o, --output      [string]          输出文件名字 / 绝对路径 (默认: data.xlsx)
  -g, --grammar                       fofa搜索语法帮助表
  -t, --tip         [string]          fofa 搜索关键字提示列表
  -x, --xlsx        [文件路径]        读取生成xlsx文件 (配合 level 参数使用)
  -l, --level       [1/2/3]               读取文件细粒度 (配合 xlsx 参数使用)
  -ih, --iconhash   [string]          计算指定 URL favicon icon_hash


```

#### 查询fofa语法

使用 -g 参数，打印出fofa语法规则

```
./ff0 -g

/ "... an experienced, industrious, ambitious, and often quite often \
| picturesque liar."                                                 |
\                 -- Mark Twain                                      /
 --------------------------------------------------------------------
  \                                  ,+*^^*+___+++_
   \                           ,*^^^^              )
    \                       _+*                     ^**+_
     \                    +^       _ _++*+_+++_,         )
              _+^^*+_    (     ,+*^ ^          \+_        )
             {       )  (    ,(    ,_+--+--,      ^)      ^\
            { (@)    } f   ,(  ,+-^ __*_*_  ^^\_   ^\       )
           {:;-/    (_+*-+^^^^^+*+*<_ _++_)_    )    )      /
          ( /  (    (        ,___    ^*+_+* )   <    <      \
           U _/     )    *--<  ) ^\-----++__)   )    )       )
            (      )  _(^)^^))  )  )\^^^^^))^*+/    /       /
          (      /  (_))_^)) )  )  ))^^^^^))^^^)__/     +^^
         (     ,/    (^))^))  )  ) ))^^^^^^^))^^)       _)
          *+__+*       (_))^)  ) ) ))^^^^^^))^^^^^)____*^
          \             \_)^)_)) ))^^^^^^^^^^))^^^^)
           (_             ^\__^^^^^^^^^^^^))^^^^^^^)
             ^\___            ^\__^^^^^^))^^^^^^^^)\\
                  ^^^^^\uuu/^^\uuu/^^^^\^\^\^\^\^\^\^\
                     ___) >____) >___   ^\_\_\_\_\_\_\)
                    ^^^//\\_^^//\\_^       ^(\_\_\_\)
                      ^^^ ^^ ^^^ ^


[+]                Rule                                   Mark
 ------------------------------------------- --------------------------------
  domain="qq.com"                             搜索根域名带有qq.com的网站
  host=".gov.cn"                              从host中搜索".gov.cn"
  ip="1.1.1.1"                                从ip中搜索包含"1.1.1.1"的网站
  ip="220.181.111.1/24"                       查询IP为"220.181.111.1"的C段资产
  port="6379"                                 查找对应"6379"端口的资产
  title="beijing"                             从标题中搜索"北京"
  status_code="402"                           查询服务器状态为"402"的资产
  protocol="quic"                             查询quic协议资产
  header="elastic"                            从http头中搜索"elastic"
  body="网络空间测绘"                            从html正文中搜索"网络空间测绘"
  os="centos"                                 搜索CentOS资产
  server=="Microsoft-IIS/10"                  搜索IIS10服务器
  app="Microsoft-Exchange"                    搜索Microsoft-Exchange设备
  base_protocol="udp"                         搜索指定udp协议的资产
  banner=users && protocol=ftp                搜索FTP协议中带有users文本的资产
  icp="京ICP证030173号"                        查找备案号为"京ICP证030173号"的网站
  icon_hash="-247388890"                      搜索使用此icon的资产(VIP)
  js_name="js/jquery.js"                      查找网站正文中包含js/jquery.js的资产
  js_md5="82ac3f14327a8b7ba49baa208d4eaa15"   查找js源码与之匹配的资产
  type=service                                搜索所有协议资产，支持subdomain和service两种
  is_domain=true                              搜索域名的资产
  ip_ports="80,161"                           搜索同时开放80和161端口的ip
  port_size="6"                               查询开放端口数量等于"6"的资产(VIP)
  port_size_gt="6"                            查询开放端口数量大于"6"的资产(VIP)
  port_size_lt="12"                           查询开放端口数量小于"12"的资产(VIP)
  is_ipv6=true                                搜索ipv6的资产
  is_fraud=false                              排除仿冒/欺诈数据
  is_honeypot=false                           排除蜜罐数据(VIP)
  country="CN"                                搜索指定国家(编码)的资产
  region="Xinjiang"                           搜索指定行政区的资产
  city="Ürümqi"                               搜索指定城市的资产
  asn="19551"                                 搜索指定asn的资产
  org="Amazon.com,Inc."                       搜索指定org(组织)的资产
  cert="baidu"                                搜索证书(https或者imaps等)中带有baidu的资产
  cert.subject="Oracle Corporation"           搜索证书持有者是OracleCorporation的资产
  cert.issuer="DigiCert"                      搜索证书颁发者为DigiCertInc的资产
  cert.is_valid=true                          验证证书是否有效,true有效,false无效(VIP)
  after="2017" && before="2017-10-01"         限定时间范围

[+] 高级搜索：可以使用括号 () / 和 && / 或 || / 完全匹配 == / 不为 != 等逻辑运算符。
```

#### 基础查询

单条语句查询

```
./ff0  -q '填写相应规则即可' -s 10000  -o data.xlsx
```

多条语句进行文件查询

```
./ff0  -f query_rules_file.txt -s 10000  -o data.xlsx
```

query_rules_file 文件实例

```
  domain="qq.com" 
  host=".gov.cn"
  ip="1.1.1.1" 
  ip="220.181.111.1/24"
  port="6379"
  title="beijing"
  status_code="402"
  protocol="quic"
```

#### 排除蜜罐干扰

-e 参数可排除蜜罐干扰

```
./ff0  -q '填写相应规则即可' -s 10000  -e -o data.xlsx
```

#### iconhash 查询

支持URL favicon Hash 查询资产

```
./ff0  -ih 'https://www.baidu.com/favicon.ico' -s 10000  -o data.xlsx
```

## 0x04 资产数据查看

资产查询结束后，会在当前文件夹下生成xlsx文件

可以使用 -x 参数指定查看的 xlxs 资产文件，并配合 -l / --level 参数指定输出细粒度

**level 1**

```
+------------------------------------+-----------------+------+
|                Host                |       Ip        | Port |
+------------------------------------+-----------------+------+
|     https://vod.sdxiehe.edu.cn     | 222.194.130.201 | 443  |
|     http://www.sdxiehe.edu.cn      | 222.194.130.199 |  80  |
|     http://vod.sdxiehe.edu.cn      | 222.194.130.201 |  80  |
|      http://oa.sdxiehe.edu.cn      | 222.194.130.69  |  80  |
|   http://stu.sdxiehe.edu.cn:8080   | 222.194.130.71  | 8080 |
| http://nxpbyhyjpsy.sdxiehe.edu.cn  | 222.194.130.118 |  80  |
|   http://dpstart.sdxiehe.edu.cn    | 222.194.130.13  |  80  |
|     http://stu.sdxiehe.edu.cn      | 222.194.130.71  |  80  |
|     http://jlsf.sdxiehe.edu.cn     | 222.194.130.126 |  80  |
+------------------------------------+-----------------+------+
```

**level 2**

```
+------------------------------------+-----------------+------+----------------+-------------------------------------------+
|                Host                |       Ip        | Port |     Domain     |                   Title                   |
+------------------------------------+-----------------+------+----------------+-------------------------------------------+
|     https://vod.sdxiehe.edu.cn     | 222.194.130.201 | 443  | sdxiehe.edu.cn |         山东协和学院视频点播系统          |
|     http://www.sdxiehe.edu.cn      | 222.194.130.199 |  80  | sdxiehe.edu.cn |                 302 Found                 |
|     http://vod.sdxiehe.edu.cn      | 222.194.130.201 |  80  | sdxiehe.edu.cn |                 302 Found                 |
|      http://oa.sdxiehe.edu.cn      | 222.194.130.69  |  80  | sdxiehe.edu.cn |         协和学院办公平台 V8.0SP1          |
|       http://sdxiehe.edu.cn        | 222.194.130.199 |  80  | sdxiehe.edu.cn |                 302 Found                 |
|     http://ids.sdxiehe.edu.cn      | 222.194.130.13  |  80  | sdxiehe.edu.cn |       山东协和学院统一身份认证平台        |
|    https://mail.sdxiehe.edu.cn     | 222.194.130.203 | 443  | sdxiehe.edu.cn |         山东协和学院电子邮件系统          |
|      http://my.sdxiehe.edu.cn      | 222.194.130.13  |  80  | sdxiehe.edu.cn |                 302 Found                 |
|     http://mail.sdxiehe.edu.cn     | 222.194.130.203 |  80  | sdxiehe.edu.cn |                                           |
|     https://www.sdxiehe.edu.cn     | 222.194.130.199 | 443  | sdxiehe.edu.cn |  山东协和学院 - 教育部批准的普通本科高校  |
|     http://xhzs.sdxiehe.edu.cn     | 222.194.130.115 |  80  | sdxiehe.edu.cn |                                           |
|     http://jpkc.sdxiehe.edu.cn     | 222.194.130.208 |  80  | sdxiehe.edu.cn |                 302 Found                 |
|    https://jpkc.sdxiehe.edu.cn     | 222.194.130.208 | 443  | sdxiehe.edu.cn |          山东协和学院精品课程网           |
|       dns1.sdxiehe.edu.cn:53       | 222.194.130.198 |  53  | sdxiehe.edu.cn |                                           |
|       dns.sdxiehe.edu.cn:53        | 222.194.130.198 |  53  | sdxiehe.edu.cn |                                           |
|        lx.sdxiehe.edu.cn:22        |   52.36.10.21   |  22  | sdxiehe.edu.cn |                                           |
|     https://waf.sdxiehe.edu.cn     | 222.194.130.251 | 443  | sdxiehe.edu.cn |                             
+------------------------------------+-----------------+------+----------------+-------------------------------------------+

```

**level** **3**

```
+------------------------------------+-----------------+------+-----------------------+----------------+-------------------------------------------+---------+
|                Host                |       Ip        | Port |        Server         |     Domain     |                   Title                   | Country |
+------------------------------------+-----------------+------+-----------------------+----------------+-------------------------------------------+---------+
|     https://vod.sdxiehe.edu.cn     | 222.194.130.201 | 443  |        Apache         | sdxiehe.edu.cn |         山东协和学院视频点播系统          |   CN    |
|     http://www.sdxiehe.edu.cn      | 222.194.130.199 |  80  |        Apache         | sdxiehe.edu.cn |                 302 Found                 |   CN    |
|     http://vod.sdxiehe.edu.cn      | 222.194.130.201 |  80  |        Apache         | sdxiehe.edu.cn |                 302 Found                 |   CN    |
|      http://oa.sdxiehe.edu.cn      | 222.194.130.69  |  80  |        SY8045         | sdxiehe.edu.cn |         协和学院办公平台 V8.0SP1          |   CN    |
|       http://sdxiehe.edu.cn        | 222.194.130.199 |  80  |        Apache         | sdxiehe.edu.cn |                 302 Found                 |   CN    |
|     http://ids.sdxiehe.edu.cn      | 222.194.130.13  |  80  |       openresty       | sdxiehe.edu.cn |       山东协和学院统一身份认证平台        |   CN    |
|    https://mail.sdxiehe.edu.cn     | 222.194.130.203 | 443  |         nginx         | sdxiehe.edu.cn |         山东协和学院电子邮件系统          |   CN    |
|      http://my.sdxiehe.edu.cn      | 222.194.130.13  |  80  |       openresty       | sdxiehe.edu.cn |                 302 Found                 |   CN    |
|     http://mail.sdxiehe.edu.cn     | 222.194.130.203 |  80  |         nginx         | sdxiehe.edu.cn |                                           |   CN    |
|     https://www.sdxiehe.edu.cn     | 222.194.130.199 | 443  |        Apache         | sdxiehe.edu.cn |  山东协和学院 - 教育部批准的普通本科高校  |   CN    |
|     http://xhzs.sdxiehe.edu.cn     | 222.194.130.115 |  80  |   Apache-Coyote/1.1   | sdxiehe.edu.cn |   
+------------------------------------+-----------------+------+-----------------------+----------------+-------------------------------------------+---------+

```



## 0x05 结语

如果您有好的建议，就在 Issues 提出来吧！
