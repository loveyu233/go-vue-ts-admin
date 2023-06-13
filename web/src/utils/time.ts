export const GetTime = () => {
    let time: number = new Date().getHours();
    let message: string = "";
    if (time > 6 && time < 9) {
        message = "早上";
    } else if (time >= 9 && time < 11) {
        message = "上午";
    } else if (time >= 11 && time < 13) {
        message = "中午";
    } else if (time >= 13 && time < 19) {
        message = "下午";
    } else {
        message = "晚上";
    }
    return message;
};