<template>
    <div>
        <el-card>
            <div style="margin: 30px auto">
                <el-button type="success" @click="addImagesShow">添加轮播图</el-button>
                <el-button type="danger" @click="operateCarousel('1')">删除选中项</el-button>
                <el-button type="warning" @click="operateCarousel('2')">修改选中项isShow为true</el-button>
                <el-button type="warning" @click="operateCarousel('3')">修改选中项isShow为false</el-button>
            </div>
            <el-dialog
                    v-model="dialogTableVisible"
                    :title="dialogTitle"
            >
                <el-form
                        :model="newImage"
                        label-width="200px"
                        ref="formData"
                >
                    <el-form-item label="是否添加该图片到轮播图" prop="username">
                        <el-switch v-model="newImage.isShow"/>
                    </el-form-item>
                    <el-form-item label="图片" prop="icon">
                        <el-upload
                                class="avatar-uploader"
                                action=""
                                :http-request="upload"
                                :show-file-list="false"
                                :on-success="handleAvatarSuccess"
                                :before-upload="beforeAvatarUpload"
                                :headers="handler"
                        >
                            <img
                                    style="width: 200px;height: 100px"
                                    v-if="newImage.url"
                                    :src="newImage.url"
                                    class="avatar"
                                    alt="重复上传以最后一个为主"
                            />
                            <el-icon v-else class="avatar-uploader-icon">
                                <Plus/>
                            </el-icon>
                        </el-upload>
                    </el-form-item>
                </el-form>
                <template #footer>
                  <span class="dialog-footer">
                    <el-button
                            type="primary"
                            @click="dialogTableVisibleFalse"
                    >
                        取消
                    </el-button>
                    <el-button
                            type="danger"
                            @click="addCarousel"
                    >
                      确定
                    </el-button>
                  </span>
                </template>
            </el-dialog>
            <el-table
                    :data="imageData"
                    @selection-change="handleSelectionChange"
            >
                <el-table-column type="selection"/>
                <el-table-column label="id" prop="id" width="100px"></el-table-column>
                <el-table-column label="url" prop="url" width="400px">
                    <template #default="{row,$index}">
                        <img :src="row.url" style="width: 300px;height: 100px"/>
                    </template>
                </el-table-column>
                <el-table-column label="isShow" prop="isShow"></el-table-column>
                <el-table-column label="操作">
                    <template #default="{row,$index}">
                        <el-button
                                type="primary"
                                icon="Edit"
                                circle
                                @click="updateImage(row)"
                        />
                        <el-popconfirm
                                confirm-button-text="Yes"
                                cancel-button-text="No"
                                icon="Delete"
                                icon-color="#626AEF"
                                title="你确定要删除吗?"
                                width="200px"
                                @confirm="deleteImage(row)"
                        >
                            <template #reference>
                                <el-button
                                        type="danger"
                                        icon="Delete"
                                        circle
                                />
                            </template>
                        </el-popconfirm>
                    </template>
                </el-table-column>
            </el-table>
        </el-card>
    </div>
</template>

<script setup lang="ts">
import {onMounted, ref} from "vue";
import {AddCarousel, DelCarousel, OperateCarouselIn, GetCarouselAll, UpdateCarousel} from "@/api/carousel";
import {CarouselType, ResIconType} from "@/type/userApi.ts";
import {UpdateIcon} from "@/api/user";
import {ElMessage, UploadProps} from "element-plus";
import {rootStore} from "@/store";

let userStore = rootStore();

let imageData = ref<CarouselType[]>([]);

let dialogTableVisible = ref<boolean>(false);

let newImage = ref<CarouselType>({isShow: false, id: 0, url: ""});

let handler = "token:" + userStore.token;

let selected = ref<CarouselType[]>([]);

let dialogTitle = ref<string>("");

let flagIsUpdate = ref<boolean>(false);

onMounted(() => {
    GetCarouselAll().then(res => {
        imageData.value = res.data;
    });
});


const dialogTableVisibleFalse = () => {
    dialogTableVisible.value = false;
    newImage.value = {id: 0, isShow: false, url: ""};
};

const addImagesShow = () => {
    dialogTableVisible.value = true;
    flagIsUpdate.value = false;
    dialogTitle.value = "添加轮播图";
};
const upload = (param) => {
    const formData = new FormData();
    formData.append("file", param.file);
    UpdateIcon(formData).then((res) => {
        console.log(res);
        newImage.value.url = res.data.url;
    });
};
const handleAvatarSuccess: UploadProps["onSuccess"] = (
    response: ResIconType,
    uploadFile
) => {
    console.log(response);
    newImage.value.url = response.data.url;
};

// 上传前的回调
const beforeAvatarUpload: UploadProps["beforeUpload"] = (rawFile) => {
    if (rawFile.type === "image/jpeg" || rawFile.type === "image/png" || rawFile.type === "image/gif") {
        if (rawFile.size / 1024 / 1024 > 2) {
            ElMessage.error("图片大小超出限制,大小请在2mb内!");
            return false;
        }
        return true;
    } else {
        ElMessage.error("图片类型不符合");
        return false;
    }
};

const addCarousel = () => {
    if (!flagIsUpdate.value) {
        AddCarousel(newImage.value).then(res => {
            console.log(res);
            if (res.data === 1) {
                GetCarouselAll().then(res => {
                    imageData.value = res.data;
                });
                ElMessage.success("上传成功");
            } else {
                ElMessage.error("上传失败");
            }
            dialogTableVisible.value = false;
        });
    } else {
        console.log(newImage.value);
        UpdateCarousel(newImage.value).then(res => {
            if (res.data === 1) {
                ElMessage.success("修改成功");
            } else {
                ElMessage.error("修改失败");
            }
            dialogTableVisible.value = false;
        }).catch(err => {
            console.log(err);
        });
    }

};

const handleSelectionChange = (v: CarouselType[]) => {
    selected.value = v;
};

const updateImage = (row) => {
    dialogTableVisible.value = true;
    flagIsUpdate.value = true;
    dialogTitle.value = "修改轮播图";
    newImage.value = row;
};

const deleteImage = (row) => {
    DelCarousel(row.id).then(res => {
        if (res.data === 1) {
            GetCarouselAll().then(res => {
                imageData.value = res.data;
            });
            ElMessage.success("删除成功");
        } else {
            ElMessage.error("删除失败");
        }
    });
};

const getSelected = () => {
    let idArr = [];
    selected.value.forEach(item => {
        idArr.push(item.id);
    });
    return idArr;
};

const operateCarousel = (operate: string) => {
    OperateCarouselIn(getSelected(), operate).then(res => {
        if (res.data > 0) {
            GetCarouselAll().then(res => {
                imageData.value = res.data;
            });
            ElMessage.success("操作成功");
        } else {
            ElMessage.error("操作失败");
        }
    });
};
</script>

<style scoped lang="scss">
</style>