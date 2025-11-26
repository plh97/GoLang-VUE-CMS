export enum INCOME_TYPE {
  '无收入' = 1,
  '1-414万円' = 2,
  '415〜600万円' = 3,
  '601〜850万円' = 4,
  '851〜1000万円' = 5,
  '1001〜1500万円' = 6,
  '1501〜2000万円' = 7,
  '2000万円～   ' = 8,
}

// 活动状态
export enum ACTIVITY_STATUS {
  招募中 = 1,
  报名已截止 = 2,
  已结束 = 3,
}

export enum GENDER {
  男 = 1,
  女 = 2,
}

// education level
export enum EDUCATION_LEVEL {
  专科 = 5,
  大学本科 = 1,
  硕士 = 2,
  博士 = 3,
  其他 = 4,
}

// sign up 状态 => 预报名，已报名，取消报名
export enum SIGN_UP_STATUS {
  // 预报名 = 1,
  // 已报名 = 2,
  // 取消报名 = 3,

  '未报名' = 0,
  '待付款' = 1,
  '已报名(已付款)' = 2,
  '已签到' = 3,
  '已参加' = 4,
  '已取消' = 5,
  '已过期' = 6,
}

// 报名的付款状态
export enum SIGN_UP_PAY_STATUS {
  未付款 = 1,
  已付款 = 2,
}

// 是否退款
export enum YES_NO {
  是 = 1,
  否 = 2,
}

// 付款状态
export enum PAY_METHOD {
  微信转账 = 1,
  Paypay转账 = 4,
}

// 体型
export enum BODY_SHAPE {
  纤瘦 = 1,
  标准 = 2,
  微胖 = 3,
  健硕 = 4,
}
// age
export enum AGE_RANGE {
  '80前' = 1,
  '80-84' = 2,
  '85-89' = 3,
  '90-94' = 4,
  '95-99' = 5,
  '00-04' = 6,
  '05后' = 7,
}

// height range
export enum HEIGHT_RANGE {
  '150及以下' = 1,
  '151-160' = 2,
  '161-170' = 3,
  '171-175' = 4,
  '176-180' = 5,
  '181-185' = 6,
  '185以上' = 7,
}

export enum MBTI {
  INTJ = 1,
  INTP = 2,
  ENTJ = 3,
  ENTP = 4,
  INFJ = 5,
  INFP = 6,
  ENFJ = 7,
  ENFP = 8,
  ISTJ = 9,
  ISFJ = 10,
  ESTJ = 11,
  ESFJ = 12,
  ISTP = 13,
  ISFP = 14,
  ESTP = 15,
  ESFP = 16,
}

export enum LOGIN_TYPE {
  // LoginTypeWxMiniApp = 1, // 小程序.      LoginTypeWxMiniApp
  微信 = 2, // 微信服务号.      LoginTypeWxWeb
  // LoginTypeApple = 3, // apple.      LoginTypeApple
  // LoginTypeGoogle = 4, // google.      LoginTypeGoogle
  // LoginTypeLine = 5, // line.      LoginTypeLine
  邮箱 = 6, // email.      LoginTypeEmail
}
