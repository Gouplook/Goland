/********************************************

@Author :yinjinlin<yinjinlin_uplook@163.com>
@Time : 2021/2/3 13:43
@Description: 消息认证

*********************************************/
package messageauthentication

import (
	"crypto/hmac"
	"crypto/sha256"
)

//  弊端
//  有秘钥分发困难的问题
//  无法解决的问题:
//  1: 不能进行第三方证明
//  2: 不能防止否认

// 流程：
// 发送者：
// 1： 发送原始法消息
// 2： 将原始消息生成消息认证码
//	   ((原始消息) + 秘钥)  *   哈希函数 = 散列值(消息认证码)
// 3： 将消息认证码发送给对方

// 接收者：
// 1：接收原始数据
// 2：接收消息认证码
// 3：校验
//	 (接收的消息  +  秘钥 ) * 哈希函数  = 新的散列值
//	 通过新的散列值和接收的散列值进行比较


// 生成消息认证码
func GenerateHamc(plainText,key []byte) []byte {
	// 1. 创建哈希接口, 需要指定使用的哈希算法, 和秘钥
	hash := hmac.New(sha256.New,key)
	// 2. 给哈希对象添加数据
	hash.Write(plainText)
	// 3. 计算散列值
	message := hash.Sum(nil)
	return message
}

// 验证消息认证码
func VerifHamc(plainText,key, hashText []byte) bool {
	// 1.创建哈希接口, 需要指定使用的哈希算法, 和秘钥
	hash := hmac.New(sha256.New, key)
	// 2. 给哈希对象添加数据
	hash.Write(plainText)
	// 3. 计算散列值
	message := hash.Sum(nil )
	// 4. 两个散列值比较
	b := hmac.Equal(hashText, message)
	return b
}