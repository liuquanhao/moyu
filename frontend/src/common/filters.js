/**
 * 转换成可读字节
 */
export function humanByte(size) {
    if (size < 1024) {
        return size + " Byte"
    } else if (size < 1024 * 1024) {
        return Math.trunc(size / 1024) + " KB"
    } else if (size < 1024 * 1024 * 1024) {
        return Math.trunc(size / 1024 / 1024) + " MB"
    } else {
        return (size / 1024 / 1024 / 1024).toFixed(2) + " GB"
    }
}

/**
 * 将秒数转换成可读日时分秒
 * 例：3600秒->1小时
 * 例：183秒->3分3秒
 */
export function humanSec(sec) {
    if (sec < 60) {
        return Math.trunc(sec) + " 秒";
    } else if (sec < 3600) {
        var sec = sec % 60
        var min = Math.trunc(sec / 60);
        return min + " 分 " + sec + " 秒"
    } else if (sec < 86400) {
        var min = Math.trunc(sec % 3600 / 60);
        var hour = Math.trunc(sec / 3600);
        return hour + " 小时 " + min + " 分"
    } else {
        var hour = Math.trunc(sec % 86400 / 3600);
        var day = Math.trunc(sec / 86400);
        return day + " 天 " + hour + " 小时"
    }
}

/**
 * 转换为num位百分比
 */
export function humanPerc(float, num) {
    if (!float || !num) {
        return 0
    }
    return float.toFixed(num)
}