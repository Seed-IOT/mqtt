# lwm2m
> 记录协议相关信息

## LWM2M Security

path: 0

用来存储访问lwm2m服务器所需的密钥

## LwM2M Server

path: 1

存储服务器相关资料

## LwM2M Access Control	

path: 2

它用于检查LWM2M服务器是否具有执行操作的访问权限

## Device

path: 3

设备对象，同时提供重启和重置

## Connectivity Monitoring

path: 4

连通性对象

## Firmware Update

path: 5

支持OTA的资源

## Location

path: 6

Location 对象

## Connectivity Statistics

path: 7

存储连接统计信息

## Lock and Wipe

path: 8

提供锁和擦除功能

## LWM2M Software Management

path: 9

软件管理，存储版本号

## Cellular connectivity

path: 10

？

## APN connection profile

path: 11

APN配置文件？

## WLAN connectivity

path: 12

wifi连接对象

## 	Bearer selection

path: 13

存储令牌？

## LWM2M Software Component

path: 14

远程软件管理对象

## DevCapMgmt

path: 15

管理设备功能，传感器等

## Portfolio

path: 16

扩展存储对象，提供加密等功能。
可用于验证和保护隐私数据。

## Communications Characteristics

path: 17

通用通信参数

## Non-Access Stratum (NAS) Configuration

path: 18

?

## BinaryAppDataContainer

path: 19

应用数据

## Event Log

path: 20

日志

## LWM2M OSCORE

path: 21

?

## Virtual Observe Notify

path: 22

提供一个函数

用于在一个通知中将多个资源ID推送到lwm2m服务器

