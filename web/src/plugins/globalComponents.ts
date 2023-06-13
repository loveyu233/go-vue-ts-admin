import {App} from "vue";

import SvgIcon from "@/components/SvgIcon.vue";

import * as ElementPlusIconsVue from "@element-plus/icons-vue";


const components: any[] = [SvgIcon];

export default {
    install(app: App) {
        components.forEach((item) => {
            app.component(item.__name, item);
        });
        for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
            app.component(key, component);
        }
    }
};