const second = 1000
const minute = 60 * second
const hour = 60 * minute
const day = 24 * hour

/**
 * 期間を人間が読める形式にフォーマットする
 * 
 * @param millis 期間（ミリ秒）
 * @returns フォーマットされた文字列
 */
export const durationHuman = (millis: number): string => {
  let remainMillis = millis
  const days = Math.floor(remainMillis / day)
  remainMillis -= days * day
  const hours = Math.floor(remainMillis / hour)
  remainMillis -= hours * hour
  const minutes = Math.floor(remainMillis / minute)
  remainMillis -= minutes * minute
  const seconds = Math.floor(remainMillis / second)
  remainMillis -= seconds * second
  if (days > 0) return `${days} day${days > 1 ? 's' : ''}`
  if (hours > 0) return `${hours} hour${hours > 1 ? 's' : ''}`
  if (minutes > 0) return `${minutes} min${minutes > 1 ? 's' : ''}`
  if (seconds > 0) return `${seconds} sec${seconds > 1 ? 's' : ''}`
  return `${remainMillis} ms`
}

/**
 * 日時の差分を人間が読める形式にフォーマットする
 * 
 * @param target 比較対象の日時
 * @returns 差分とロケール文字列
 */
export const diffHuman = (target: Date) => {
  const diff = new Date().getTime() - target.getTime()
  const suffix = diff > 0 ? 'ago' : 'from now'
  const human = durationHuman(Math.abs(diff))
  const localeString = target.toLocaleString()
  return {
    diff: `${human} ${suffix}`,
    localeString
  }
}
