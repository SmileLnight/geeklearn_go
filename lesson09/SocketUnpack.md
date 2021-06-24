# SocketUnpack

---

# fix length

## 总结

Fixed Length 是基于固定长度消息进行粘包拆包处理的

## 举例

如下所示，收到以下四个分片数据包：

```
 +---+----+------+----+
 | A | BC | DEFG | HI |
 +---+----+------+----+
```

可以通过**Fixed Length**将它们解码为以下三个固定长度的数据包：

```
 +-----+-----+-----+
 | ABC | DEF | GHI |
 +-----+-----+-----+
```

# delimiter based

## 总结

Delimiter Based 是基于消息边界方式进行粘包拆包处理的

## 举例

有以下数据：

```
 +------------+
 | ABC\nDEF\n |
 +------------+
```

**Delimiter Based**将选择**\n**作为第一个分隔符并产生两个帧：

```
 +-----+-----+
 | ABC | DEF |
 +-----+-----+
```

# length field based frame decoder

## 总结

length field based frame decoder是基于消息头指定消息长度进行粘包拆包处理的

## 举例

长度字段的值为`12（0x0C）`，代表"HELLO, WORLD"的长度。默认情况下，解码器假定长度字段表示长度字段后面的字节数:

解码器配置参数

```
lengthFieldOffset   = 0
lengthFieldLength   = 2
lengthAdjustment    = 0
initialBytesToStrip = 0 (不剥离标头)
```

解码前：

```
 BEFORE DECODE (14 bytes)      
 +--------+----------------+
 | Length | Actual Content |
 | 0x000C | "HELLO, WORLD" |
 +--------+----------------+
```

解码后：

```
AFTER DECODE (14 bytes)
+--------+----------------+
| Length | Actual Content |
| 0x000C | "HELLO, WORLD" |
+--------+----------------+
```