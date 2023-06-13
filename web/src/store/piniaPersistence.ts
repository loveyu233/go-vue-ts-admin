import {PiniaPluginContext} from "pinia";
import {toRaw} from "vue";

type O = {
    key?: string
}
export const piniaPlugin = (options: O) => {
    return (content: PiniaPluginContext) => {
        const {store} = content;
        const data = getLocalStore(`${options?.key ?? "default"}-${store.$id}`);
        store.$subscribe(() => {
            setLocalStore(`${options?.key ?? "default"}-${store.$id}`, toRaw(store.$state));
        });
        return {
            ...data
        };
    };

};

const setLocalStore = (key: string, value: any) => {
    localStorage.setItem(key, JSON.stringify(value));
};

const getLocalStore = (key: string) => {
    return localStorage.getItem(key) ? JSON.parse(localStorage.getItem(key) as string) : {};
};