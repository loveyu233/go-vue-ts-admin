import {defineStore} from "pinia";

export const layoutStore = defineStore("layout", {
    state: () => {
        return {
            fold: false,
            refresh: false,
            dark: false
        };
    },
    actions: {
        checkFold() {
            this.fold = !this.fold;
        },
        checkRefresh() {
            this.refresh = !this.refresh;
        },
        checkDark() {
            this.dark = !this.dark;
        }
    },
    getters: {}
});