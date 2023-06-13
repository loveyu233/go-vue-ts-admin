import {defineConfig} from "vite";
import vue from "@vitejs/plugin-vue";
import * as path from "path";
import {createSvgIconsPlugin} from "vite-plugin-svg-icons";

export default defineConfig({
    plugins: [
        vue(),
        createSvgIconsPlugin({
            // Specify the icon folder to be cached
            iconDirs: [path.resolve(process.cwd(), "src/assets/icons")],
            // Specify symbolId format
            symbolId: "icon-[dir]-[name]",
            svgoOptions: {
                plugins: [
                    {
                        name: "removeAttrs",
                        params: {attrs: ["class", "data-name", "fill", "stroke"]}
                    }
                ]
            }
        })
    ],
    resolve: {
        alias: {
            "@": path.resolve("./src")
        }
    },
    // 配置scss全局变量
    css: {
        preprocessorOptions: {
            scss: {
                javascriptEnabled: true,
                additionalData: "@import \"./src/style/globalVar.scss\";"
            }
        }
    }
});


