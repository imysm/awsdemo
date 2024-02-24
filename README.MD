# fearture awsdemo

    简单实现一个读写aws rds for mysql 、elastic cache for redis的接口

## 说明

    1、mysql和redis配置信息从secret manager 读取
    2、secret mange的登录校验信息从环境变量或者aws的默认配置文件读取
    3、建立 rds for mysql 连接, 测试数据写入和读取
    4、建立 alastic cache for redis 链接 测试写入和读取
    5、简单验证 token 可在config中自定义 ,如未配置则使用默认tokn