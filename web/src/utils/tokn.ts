export const SET_TOKEN = (token: string) => {
    localStorage.setItem("token", token);
};

export const GET_TOKEN = (): string => {
    return localStorage.getItem("token") as string;
};

export const Del_TOKEN = () => {
    return localStorage.removeItem("token");
};